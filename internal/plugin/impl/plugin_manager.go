package impl

import (
	"context"
	"encoding/json"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"sync"
	"time"

	"github.com/hashicorp/go-plugin"

	"github.com/Done-0/jank/internal/global"
	"github.com/Done-0/jank/pkg/plugin/consts"

	pluginUtils "github.com/Done-0/jank/internal/utils/plugin"
	jank "github.com/Done-0/jank/pkg/plugin"
)

// PluginInfo 插件元数据和运行时信息
type PluginInfo struct {
	// 基本信息
	ID          string `json:"id"`                    // 插件唯一标识
	Name        string `json:"name"`                  // 插件显示名称
	Version     string `json:"version"`               // 插件版本号
	Author      string `json:"author"`                // 插件作者
	Description string `json:"description,omitempty"` // 插件描述
	Repository  string `json:"repository"`            // 插件仓库地址
	Binary      string `json:"binary"`                // 二进制文件路径
	Type        string `json:"type"`                  // 插件类型（provider/filter/handler/notifier）

	// 配置信息
	AutoStart    bool  `json:"auto_start"`              // 是否自动启动
	StartTimeout int64 `json:"start_timeout,omitempty"` // 启动超时(毫秒)
	MinPort      uint  `json:"min_port,omitempty"`      // 最小端口
	MaxPort      uint  `json:"max_port,omitempty"`      // 最大端口
	AutoMTLS     bool  `json:"auto_mtls,omitempty"`     // 是否启用自动 MTLS
	Managed      bool  `json:"managed,omitempty"`       // 是否为托管模式

	// 运行时信息
	Status            string `json:"status"`                       // 当前状态
	StartedAt         int64  `json:"started_at,omitempty"`         // 启动时间戳
	ProcessID         string `json:"process_id,omitempty"`         // 进程标识
	Protocol          string `json:"protocol,omitempty"`           // 通信协议
	IsExited          bool   `json:"is_exited,omitempty"`          // 是否已退出
	NegotiatedVersion int    `json:"negotiated_version,omitempty"` // 协商的协议版本
	ProcessPID        int    `json:"process_pid,omitempty"`        // 系统进程 PID
	ProtocolVersion   int    `json:"protocol_version,omitempty"`   // 协议版本
	NetworkAddr       string `json:"network_addr,omitempty"`       // 网络地址
}

// PluginDiscoveryInfo 插件发现信息
type PluginDiscoveryInfo struct {
	*PluginInfo         // 嵌入插件基本信息
	Path         string `json:"path"`          // 插件路径
	IsRegistered bool   `json:"is_registered"` // 是否已注册
}

// PluginManagerImpl 插件管理器实现
type PluginManagerImpl struct {
	plugins map[string]*plugin.Client // 插件客户端映射
	infos   map[string]*PluginInfo    // 插件信息映射
	mu      sync.RWMutex              // 并发安全锁
}

// NewPluginManager 创建插件管理器实例
func NewPluginManager() *PluginManagerImpl {
	return &PluginManagerImpl{
		plugins: make(map[string]*plugin.Client),
		infos:   make(map[string]*PluginInfo),
	}
}

// RegisterPlugin 注册并启动插件
func (m *PluginManagerImpl) RegisterPlugin(info *PluginInfo) error {
	m.mu.Lock()
	defer m.mu.Unlock()

	if _, exists := m.infos[info.ID]; exists {
		return fmt.Errorf("plugin %s already registered", info.ID)
	}

	// 创建插件客户端并启动
	config := &plugin.ClientConfig{
		HandshakeConfig:  jank.HandshakeConfig,
		Plugins:          jank.PluginMap,
		Cmd:              exec.Command(info.Binary),
		AllowedProtocols: []plugin.Protocol{plugin.ProtocolGRPC},
		AutoMTLS:         info.AutoMTLS,
		Managed:          info.Managed,
		StartTimeout:     time.Duration(info.StartTimeout) * time.Millisecond,
		MinPort:          info.MinPort,
		MaxPort:          info.MaxPort,
	}

	client := plugin.NewClient(config)
	if _, err := client.Start(); err != nil {
		client.Kill()
		return fmt.Errorf("failed to start plugin %s: %v", info.ID, err)
	}

	// 更新运行时状态
	info.Status = consts.PluginStatusLoaded
	info.StartedAt = time.Now().Unix()
	m.refreshPluginInfo(info, client)

	// 保存到内存映射
	m.infos[info.ID] = info
	m.plugins[info.ID] = client

	global.SysLog.Infof("Plugin registered: %s (%s v%s) from %s, PID: %d",
		info.ID, info.Name, info.Version, info.Repository, info.ProcessPID)

	return nil
}

// UnregisterPlugin 注销并停止插件
func (m *PluginManagerImpl) UnregisterPlugin(id string) error {
	m.mu.Lock()
	defer m.mu.Unlock()

	client, exists := m.plugins[id]
	if !exists {
		return fmt.Errorf("plugin %s not found", id)
	}

	client.Kill()

	if info, ok := m.infos[id]; ok {
		info.Status = consts.PluginStatusStopped
		global.SysLog.Infof("Plugin unregistered: %s (%s v%s)",
			id, info.Name, info.Version)
	}

	delete(m.plugins, id)
	delete(m.infos, id)

	return nil
}

// ExecutePlugin 执行插件方法
func (m *PluginManagerImpl) ExecutePlugin(ctx context.Context, id, method string, args map[string]string) (map[string]string, error) {
	m.mu.RLock()
	client, exists := m.plugins[id]
	info, infoExists := m.infos[id]
	m.mu.RUnlock()

	if !exists || !infoExists {
		return nil, fmt.Errorf("plugin %s not found", id)
	}

	// 更新状态为执行中
	m.mu.Lock()
	info.Status = consts.PluginStatusRunning
	m.refreshPluginInfo(info, client)
	m.mu.Unlock()

	// 获取 RPC 客户端并执行
	rpcClient, err := client.Client()
	if err != nil {
		m.mu.Lock()
		info.Status = consts.PluginStatusError
		m.mu.Unlock()
		return nil, err
	}

	raw, err := rpcClient.Dispense(info.Type)
	if err != nil {
		m.mu.Lock()
		info.Status = consts.PluginStatusError
		m.mu.Unlock()
		return nil, err
	}

	// 执行插件方法
	result, err := raw.(jank.Plugin).Execute(ctx, method, args)

	// 更新执行结果状态
	m.mu.Lock()
	if err != nil {
		info.Status = consts.PluginStatusError
	} else {
		info.Status = consts.PluginStatusReady
	}
	m.refreshPluginInfo(info, client)
	m.mu.Unlock()

	return result, err
}

// GetPlugin 获取插件信息
func (m *PluginManagerImpl) GetPlugin(id string) (*PluginInfo, error) {
	m.mu.RLock()
	info, exists := m.infos[id]
	client, clientExists := m.plugins[id]
	m.mu.RUnlock()

	if !exists {
		return nil, fmt.Errorf("plugin %s not found", id)
	}

	// 刷新运行时信息
	if clientExists {
		m.mu.Lock()
		m.refreshPluginInfo(info, client)
		m.mu.Unlock()
	}

	return m.createPluginCopy(info), nil
}

// ListPlugins 列出所有插件（包括未注册的）
func (m *PluginManagerImpl) ListPlugins() ([]*PluginDiscoveryInfo, error) {
	entries, err := os.ReadDir(consts.PluginDir)
	if err != nil {
		return nil, fmt.Errorf("failed to read plugin directory: %w", err)
	}

	var discoveredPlugins []*PluginDiscoveryInfo

	// 获取已注册插件的ID列表
	m.mu.RLock()
	registeredPluginIDs := make(map[string]bool)
	registeredPluginsMap := make(map[string]*PluginInfo)
	for id, info := range m.infos {
		registeredPluginIDs[id] = true
		// 刷新运行时信息
		if client, exists := m.plugins[id]; exists {
			m.refreshPluginInfo(info, client)
		}
		registeredPluginsMap[id] = info
	}
	m.mu.RUnlock()

	// 遍历所有插件目录
	for _, entry := range entries {
		if !entry.IsDir() {
			continue
		}

		pluginPath := filepath.Join(consts.PluginDir, entry.Name())
		configFile := filepath.Join(pluginPath, consts.PluginConfigFile)

		configData, err := os.ReadFile(configFile)
		if err != nil {
			global.SysLog.Warnf("Skipping %s: cannot read %s", entry.Name(), consts.PluginConfigFile)
			continue
		}

		var config PluginInfo
		if err := json.Unmarshal(configData, &config); err != nil {
			global.SysLog.Warnf("Skipping %s: invalid %s", entry.Name(), consts.PluginConfigFile)
			continue
		}

		// 检查插件状态和运行时信息
		isRegistered := registeredPluginIDs[config.ID]
		var finalInfo *PluginInfo

		if isRegistered {
			// 使用已注册插件的运行时信息
			finalInfo = m.createPluginCopy(registeredPluginsMap[config.ID])
		} else {
			// 检查是否有二进制文件或源码
			binaryPath := pluginUtils.GenerateBinaryPath(pluginPath, config.ID, config.Binary)
			var status string
			if pluginUtils.CheckBinaryExists(binaryPath) {
				status = consts.PluginStatusAvailable
			} else if pluginUtils.CheckMainFileExists(pluginPath) {
				status = consts.PluginStatusSourceOnly
			} else {
				status = consts.PluginStatusIncomplete
			}
			// 使用配置文件信息
			finalInfo = &config
			finalInfo.Status = status
		}

		discoveryInfo := &PluginDiscoveryInfo{
			PluginInfo:   finalInfo,
			Path:         pluginPath,
			IsRegistered: isRegistered,
		}

		discoveredPlugins = append(discoveredPlugins, discoveryInfo)
	}

	return discoveredPlugins, nil
}

// Shutdown 关闭所有插件
func (m *PluginManagerImpl) Shutdown() {
	m.mu.Lock()
	defer m.mu.Unlock()

	for id, client := range m.plugins {
		client.Kill()
		if info, ok := m.infos[id]; ok {
			info.Status = consts.PluginStatusStopped
			info.IsExited = true
			global.SysLog.Infof("Plugin stopped: %s", id)
		}
	}

	m.plugins = make(map[string]*plugin.Client)
	m.infos = make(map[string]*PluginInfo)
}

// StartAutoPlugins 扫描并启动配置为自动启动的插件
func (m *PluginManagerImpl) StartAutoPlugins() error {
	entries, err := os.ReadDir(consts.PluginDir)
	if err != nil {
		if os.IsNotExist(err) {
			global.SysLog.Info("Plugin directory not found")
			return nil
		}
		return fmt.Errorf("failed to read plugin directory: %w", err)
	}

	for _, entry := range entries {
		if !entry.IsDir() {
			continue
		}

		// 构建插件路径
		pluginPath := filepath.Join(consts.PluginDir, entry.Name())
		configFile := filepath.Join(pluginPath, consts.PluginConfigFile)

		// 读取配置文件
		configData, err := os.ReadFile(configFile)
		if err != nil {
			global.SysLog.Warnf("Skipping %s: no %s found", entry.Name(), consts.PluginConfigFile)
			continue
		}

		var config PluginInfo
		if err := json.Unmarshal(configData, &config); err != nil {
			global.SysLog.Warnf("Skipping %s: invalid %s", entry.Name(), consts.PluginConfigFile)
			continue
		}

		// 只处理标记为自动启动的插件
		if !config.AutoStart {
			continue
		}

		// 生成二进制路径
		binaryPath := pluginUtils.GenerateBinaryPath(pluginPath, config.ID, config.Binary)

		// 检查二进制文件是否存在，不存在则尝试编译
		if !pluginUtils.CheckBinaryExists(binaryPath) {
			if !pluginUtils.CheckMainFileExists(pluginPath) {
				global.SysLog.Warnf("Plugin %s: binary and main.go not found in %s", config.ID, pluginPath)
				continue
			}

			if err := pluginUtils.EnsureBinDirectory(pluginPath); err != nil {
				global.SysLog.Warnf("Plugin %s: failed to create bin directory: %v", config.ID, err)
				continue
			}

			if err := pluginUtils.RunGoModTidy(pluginPath); err != nil {
				global.SysLog.Warnf("Plugin %s: go mod tidy failed: %v", config.ID, err)
				continue
			}

			outputPath := pluginUtils.GenerateOutputPath(config.Binary, config.ID)
			if err := pluginUtils.CompileGoPlugin(pluginPath, outputPath); err != nil {
				global.SysLog.Warnf("Plugin %s: compilation failed: %v", config.ID, err)
				continue
			}

			global.SysLog.Infof("Plugin compiled successfully: %s -> %s", config.ID, outputPath)
		}

		// 更新配置中的二进制路径
		config.Binary = binaryPath

		// 注册并启动插件
		if err := m.RegisterPlugin(&config); err != nil {
			global.SysLog.Errorf("Failed to auto-start plugin %s: %v", config.ID, err)
		}
	}

	return nil
}

// refreshPluginInfo 刷新插件运行时信息
func (m *PluginManagerImpl) refreshPluginInfo(info *PluginInfo, client *plugin.Client) {
	info.ProcessID = client.ID()
	info.Protocol = string(client.Protocol())
	info.IsExited = client.Exited()
	info.NegotiatedVersion = client.NegotiatedVersion()

	if config := client.ReattachConfig(); config != nil {
		info.ProcessPID = config.Pid
		info.ProtocolVersion = config.ProtocolVersion
		if config.Addr != nil {
			info.NetworkAddr = config.Addr.String()
		}
	}
}

// createPluginCopy 创建插件信息的深拷贝
func (m *PluginManagerImpl) createPluginCopy(info *PluginInfo) *PluginInfo {
	return &PluginInfo{
		ID: info.ID, Name: info.Name, Version: info.Version, Author: info.Author,
		Description: info.Description, Repository: info.Repository, Binary: info.Binary,
		Type:      info.Type,
		AutoStart: info.AutoStart, StartTimeout: info.StartTimeout,
		MinPort: info.MinPort, MaxPort: info.MaxPort,
		AutoMTLS: info.AutoMTLS, Managed: info.Managed,
		Status: info.Status, StartedAt: info.StartedAt,
		ProcessID: info.ProcessID, Protocol: info.Protocol,
		IsExited: info.IsExited, NegotiatedVersion: info.NegotiatedVersion,
		ProcessPID: info.ProcessPID, ProtocolVersion: info.ProtocolVersion,
		NetworkAddr: info.NetworkAddr,
	}
}
