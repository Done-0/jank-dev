package impl

import (
	"context"
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"

	"github.com/cloudwego/hertz/pkg/app"

	"github.com/Done-0/jank/internal/global"
	"github.com/Done-0/jank/internal/plugin"
	"github.com/Done-0/jank/internal/plugin/impl"
	"github.com/Done-0/jank/pkg/plugin/consts"
	"github.com/Done-0/jank/pkg/serve/controller/dto"
	"github.com/Done-0/jank/pkg/serve/service"
	"github.com/Done-0/jank/pkg/vo"

	pluginUtils "github.com/Done-0/jank/internal/utils/plugin"
)

type PluginServiceImpl struct{}

func NewPluginService() service.PluginService {
	return &PluginServiceImpl{}
}

// RegisterPlugin 注册插件
func (s *PluginServiceImpl) RegisterPlugin(c *app.RequestContext, req *dto.RegisterPluginRequest) (*vo.RegisterPluginResponse, error) {
	entries, err := os.ReadDir(consts.PluginDir)
	if err != nil {
		return &vo.RegisterPluginResponse{Message: err.Error()}, fmt.Errorf("failed to read plugin directory: %w", err)
	}

	var pluginInfo *impl.PluginInfo
	for _, entry := range entries {
		if !entry.IsDir() {
			continue
		}

		pluginPath := filepath.Join(consts.PluginDir, entry.Name())
		configFile := filepath.Join(pluginPath, consts.PluginConfigFile)

		configData, err := os.ReadFile(configFile)
		if err != nil {
			continue
		}

		var config impl.PluginInfo
		if err := json.Unmarshal(configData, &config); err != nil {
			continue
		}

		if config.ID == req.ID {
			// 生成二进制路径
			binaryPath := pluginUtils.GenerateBinaryPath(pluginPath, config.ID, config.Binary)

			// 检查二进制文件是否存在，不存在则尝试编译
			if !pluginUtils.CheckBinaryExists(binaryPath) {
				if !pluginUtils.CheckMainFileExists(pluginPath) {
					return &vo.RegisterPluginResponse{Message: fmt.Sprintf("binary and %s not found", consts.PluginMainFile)},
						fmt.Errorf("binary and %s not found for plugin %s", consts.PluginMainFile, req.ID)
				}

				if err := pluginUtils.EnsureBinDirectory(pluginPath); err != nil {
					return &vo.RegisterPluginResponse{Message: err.Error()},
						fmt.Errorf("failed to create bin directory: %w", err)
				}

				if err := pluginUtils.RunGoModTidy(pluginPath); err != nil {
					return &vo.RegisterPluginResponse{Message: err.Error()},
						fmt.Errorf("go mod tidy failed: %w", err)
				}

				outputPath := pluginUtils.GenerateOutputPath(config.Binary, config.ID)
				if err := pluginUtils.CompileGoPlugin(pluginPath, outputPath); err != nil {
					return &vo.RegisterPluginResponse{Message: err.Error()},
						fmt.Errorf("compilation failed: %w", err)
				}

				global.SysLog.Infof("Plugin compiled successfully: %s -> %s", config.ID, outputPath)
			}

			config.Binary = binaryPath
			pluginInfo = &config
			break
		}
	}

	if pluginInfo == nil {
		return &vo.RegisterPluginResponse{Message: "plugin not found"},
			fmt.Errorf("plugin with ID %s not found", req.ID)
	}

	if err := plugin.GlobalPluginManager.RegisterPlugin(pluginInfo); err != nil {
		return &vo.RegisterPluginResponse{Message: err.Error()}, fmt.Errorf("failed to register plugin %s: %v", req.ID, err)
	}

	return &vo.RegisterPluginResponse{Message: "Plugin registered successfully"}, nil
}

// UnregisterPlugin 注销插件
func (s *PluginServiceImpl) UnregisterPlugin(c *app.RequestContext, req *dto.UnregisterPluginRequest) (*vo.UnregisterPluginResponse, error) {
	if err := plugin.GlobalPluginManager.UnregisterPlugin(req.ID); err != nil {
		return &vo.UnregisterPluginResponse{Message: err.Error()}, err
	}
	return &vo.UnregisterPluginResponse{Message: "Plugin unregistered successfully"}, nil
}

// ExecutePlugin 执行插件方法
func (s *PluginServiceImpl) ExecutePlugin(c *app.RequestContext, req *dto.ExecutePluginRequest) (*vo.ExecutePluginResponse, error) {
	payloadMap := map[string]string{consts.PluginPayloadKey: req.Payload}

	result, err := plugin.GlobalPluginManager.ExecutePlugin(context.Background(), req.ID, req.Method, payloadMap)
	if err != nil {
		return &vo.ExecutePluginResponse{}, fmt.Errorf("failed to execute plugin %s method %s: %v", req.ID, req.Method, err)
	}

	resultMap := make(map[string]any)
	for k, v := range result {
		resultMap[k] = v
	}

	return &vo.ExecutePluginResponse{
		Method: req.Method,
		Data:   resultMap,
	}, nil
}

// GetPlugin 获取插件信息
func (s *PluginServiceImpl) GetPlugin(c *app.RequestContext, req *dto.GetPluginRequest) (*vo.GetPluginResponse, error) {
	info, err := plugin.GlobalPluginManager.GetPlugin(req.ID)
	if err != nil {
		return &vo.GetPluginResponse{}, fmt.Errorf("failed to get plugin %s: %v", req.ID, err)
	}

	response := vo.GetPluginResponse{
		// 基本信息
		ID:          info.ID,
		Name:        info.Name,
		Version:     info.Version,
		Author:      info.Author,
		Description: info.Description,
		Repository:  info.Repository,
		Binary:      info.Binary,
		Type:        info.Type,

		// 配置信息
		AutoStart:    info.AutoStart,
		StartTimeout: info.StartTimeout,
		MinPort:      info.MinPort,
		MaxPort:      info.MaxPort,
		AutoMTLS:     info.AutoMTLS,
		Managed:      info.Managed,

		// 运行时信息
		Status:            info.Status,
		StartedAt:         info.StartedAt,
		ProcessID:         info.ProcessID,
		Protocol:          info.Protocol,
		IsExited:          info.IsExited,
		NegotiatedVersion: info.NegotiatedVersion,
		ProcessPID:        info.ProcessPID,
		ProtocolVersion:   info.ProtocolVersion,
		NetworkAddr:       info.NetworkAddr,
	}
	return &response, nil
}

// ListPlugins 列举所有插件（包括未注册的）
func (s *PluginServiceImpl) ListPlugins(c *app.RequestContext, req *dto.ListPluginsRequest) (*vo.ListPluginsResponse, error) {
	discoveredPlugins, err := plugin.GlobalPluginManager.ListPlugins()
	if err != nil {
		return &vo.ListPluginsResponse{}, fmt.Errorf("failed to list plugins: %w", err)
	}

	var allPlugins []vo.GetPluginResponse
	for _, discovered := range discoveredPlugins {
		pluginVO := vo.GetPluginResponse{
			// 基本信息
			ID:          discovered.ID,
			Name:        discovered.Name,
			Version:     discovered.Version,
			Author:      discovered.Author,
			Description: discovered.Description,
			Repository:  discovered.Repository,
			Binary:      discovered.Binary,
			Type:        discovered.Type,

			// 配置信息
			AutoStart:    discovered.AutoStart,
			StartTimeout: discovered.StartTimeout,
			MinPort:      discovered.MinPort,
			MaxPort:      discovered.MaxPort,
			AutoMTLS:     discovered.AutoMTLS,
			Managed:      discovered.Managed,

			// 运行时信息
			Status:            discovered.Status,
			StartedAt:         discovered.StartedAt,
			ProcessID:         discovered.ProcessID,
			Protocol:          discovered.Protocol,
			IsExited:          discovered.IsExited,
			NegotiatedVersion: discovered.NegotiatedVersion,
			ProcessPID:        discovered.ProcessPID,
			ProtocolVersion:   discovered.ProtocolVersion,
			NetworkAddr:       discovered.NetworkAddr,
		}
		allPlugins = append(allPlugins, pluginVO)
	}

	return &vo.ListPluginsResponse{
		Plugins: allPlugins,
		Total:   int64(len(allPlugins)),
	}, nil
}
