# Jank 插件系统

基于 hashicorp/go-plugin 的进程隔离插件架构，支持自动编译和生命周期管理。

## 🎯 系统架构

```bash
HTTP API → PluginServiceImpl → PluginManagerImpl → hashicorp/go-plugin
                                      ↓
Plugin Process (gRPC) ←→ Main Process
```

**核心组件：**
- `PluginManagerImpl`: 插件生命周期管理
- `PluginServiceImpl`: HTTP API服务层
- `PluginInfo`: 插件元数据和运行时状态
- `CompileGoPlugin`: 自动编译工具

## 🚀 核心特性

### 进程隔离
每个插件运行在独立进程中，通过 gRPC 通信，插件崩溃不影响主进程。

### 自动编译
检测源码变化自动编译为二进制文件：
```bash
CGO_ENABLED=0 go build -o bin/plugin-name main.go
```

### 类型安全通信
基于 Protocol Buffers 的 gRPC 接口确保类型安全。

### 生命周期管理
支持插件的加载、执行、停止、卸载全流程管理。

## 📁 目录结构

```bash
internal/plugin/
├── impl/
│   └── plugin_manager.go      # 核心管理器实现
├── proto/
│   ├── plugin.proto          # gRPC接口定义
│   ├── plugin.pb.go          # 生成的protobuf代码
│   └── plugin_grpc.pb.go     # 生成的gRPC代码
└── README.md                 # 本文档

pkg/plugin/
├── consts/
│   └── plugin.go             # 常量定义
├── grpc.go                   # gRPC客户端/服务端
├── interface.go              # 插件接口定义
└── plugin.go                 # 插件实现

plugins/                      # 插件存放目录
└── hello-world/
    ├── main.go              # 插件代码
    ├── plugin.json          # 插件配置
    └── bin/                 # 编译输出
```

## ⚙️ 配置文件

### plugin.json 格式
```json
{
  "id": "dev.jank.plugins.hello-world",
  "name": "Hello World Plugin",
  "version": "1.0.0",
  "author": "Done-0",
  "type": "handler",
  "auto_start": true,
  "binary": "hello-world"
}
```

### 插件ID命名规范
采用反向域名格式：`dev.jank.plugins.plugin-name`

### 插件类型标识符
- `provider`: 数据提供者插件
- `filter`: 数据过滤插件
- `handler`: 业务处理插件
- `notifier`: 通知插件

## 🔧 插件开发示例

### 基本插件结构
```go
package main

import (
    "context"
    "github.com/hashicorp/go-plugin"
    "github.com/Done-0/jank/pkg/plugin/consts"
    jank "github.com/Done-0/jank/pkg/plugin"
)

type MyPlugin struct{}

func (p *MyPlugin) Execute(ctx context.Context, method string, args map[string]string) (map[string]string, error) {
    // 插件业务逻辑
    return map[string]string{"result": "success"}, nil
}

func (p *MyPlugin) HealthCheck(ctx context.Context) error {
    return nil
}

func main() {
    plugin.Serve(&plugin.ServeConfig{
        HandshakeConfig: jank.HandshakeConfig,
        Plugins: map[string]plugin.Plugin{
            consts.PluginTypeHandler: jank.NewGRPCPlugin(&MyPlugin{}),
        },
        GRPCServer: plugin.DefaultGRPCServer,
    })
}
```

## 🌐 HTTP API

### 插件列表 `GET /api/plugin/list`
返回所有插件（包括已注册和未注册）：
```json
{
  "code": 0,
  "data": {
    "registered": [...],
    "unregistered": [...]
  }
}
```

### 执行插件 `POST /api/plugin/execute`
```json
{
  "id": "dev.jank.plugins.hello-world",
  "method": "greet",
  "payload": "World"
}
```

## 🔄 插件状态

### 已注册插件状态
- `ready`: 插件就绪，可执行
- `running`: 插件正在运行
- `stopped`: 插件已停止
- `error`: 插件运行错误

### 未注册插件状态  
- `available`: 有二进制文件，可直接注册
- `source_only`: 仅有源码，需编译
- `incomplete`: 配置不完整

## 🛠️ 核心组件

### PluginManagerImpl
- 插件注册/注销管理
- 进程生命周期控制
- 自动编译和发现

### PluginInfo
```go
type PluginInfo struct {
    ID            string // 插件唯一标识
    Name          string // 显示名称
    Version       string // 版本号
    Type          string // 插件类型
    Status        string // 运行状态
    ProcessID     string // 进程ID
    IsExited      bool   // 是否已退出
}
```

### PluginDiscoveryInfo
用于插件发现和列表展示，嵌入 PluginInfo 并添加路径和注册状态信息。

## 🔒 安全特性

- 进程隔离：插件在独立进程中运行
- gRPC通信：类型安全的远程调用
- 超时控制：防止插件无响应
- 资源限制：可配置的资源约束

## 📊 性能特点

- 插件并发执行互不干扰
- 自动进程回收和资源清理
- 支持插件热加载和卸载
- 最大支持100个并发插件

## 🐛 错误处理

系统通过多层错误处理确保稳定性：
1. 插件进程崩溃自动重启
2. gRPC通信错误重试机制  
3. 超时和资源限制保护
4. 详细的错误日志记录