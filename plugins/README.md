# Jank 插件开发指南

快速开发和部署 Jank 插件的完整指南。

## 🚀 快速开始

### 1. 创建插件
```bash
mkdir plugins/com.example.my-plugin
cd plugins/com.example.my-plugin
go mod init com.example.my-plugin
go get github.com/Done-0/jank/pkg/plugin
go get github.com/Done-0/jank/pkg/plugin/consts
go get github.com/hashicorp/go-plugin
```

### 2. 编写代码
创建 `main.go`:
```go
package main

import (
    "context"
    "fmt"

    "github.com/hashicorp/go-plugin"
    "github.com/Done-0/jank/pkg/plugin/consts"
    jank "github.com/Done-0/jank/pkg/plugin"
)

type MyPlugin struct{}

func (p *MyPlugin) Execute(ctx context.Context, method string, args map[string]string) (map[string]string, error) {
    switch method {
    case "greet":
        name := args["name"]
        if name == "" {
            name = "World"
        }
        return map[string]string{"message": "Hello, " + name + "!"}, nil
    case "echo":
        return args, nil
    default:
        return nil, fmt.Errorf("unknown method: %s", method)
    }
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

### 3. 创建配置
创建 `plugin.json`:
```json
{
  "id": "com.example.my-plugin",
  "name": "My Plugin",
  "version": "1.0.0",
  "author": "Your Name",
  "description": "插件描述",
  "type": "handler",
  "auto_start": true,
  "binary": "my-plugin"
}
```

### 4. 测试插件
```bash
# 系统会自动编译和注册插件，然后测试：
curl -X POST http://localhost:8080/api/plugin/execute \
  -H "Content-Type: application/json" \
  -d '{"id": "com.example.my-plugin", "method": "greet", "payload": "Claude"}'
```

## 📁 目录结构
```
plugins/
└── com.example.my-plugin/
    ├── main.go           # 插件主代码
    ├── plugin.json       # 插件配置
    ├── go.mod           # Go模块文件
    └── bin/             # 自动编译生成的二进制文件目录
        └── my-plugin
```

## ⚙️ 配置文件

### plugin.json 格式
```json
{
  "id": "com.example.plugin-name",     // 插件唯一ID（反向域名）
  "name": "Plugin Name",               // 显示名称
  "version": "1.0.0",                 // 版本号
  "author": "Author",                 // 作者
  "description": "Plugin description", // 描述
  "type": "handler",                  // 类型：provider/filter/handler/notifier
  "auto_start": true,                 // 是否自动启动
  "binary": "plugin-name"             // 二进制文件名
}
```

### 插件类型
- `provider`: 数据提供者插件
- `filter`: 数据过滤插件  
- `handler`: 业务处理插件
- `notifier`: 通知插件

## 🔧 开发最佳实践

### 错误处理
```go
func (p *MyPlugin) Execute(ctx context.Context, method string, args map[string]string) (map[string]string, error) {
    // 验证输入
    if method == "" {
        return nil, fmt.Errorf("method cannot be empty")
    }
    
    // 返回错误信息到结果中（用户友好）
    return map[string]string{"error": "invalid input"}, nil
}
```

### 上下文处理
```go
func (p *MyPlugin) Execute(ctx context.Context, method string, args map[string]string) (map[string]string, error) {
    select {
    case <-ctx.Done():
        return nil, ctx.Err()
    default:
        // 执行业务逻辑
    }
}
```

## 🌐 插件状态

### 自动检测状态
- `available`: 有二进制文件，可直接注册
- `source_only`: 仅有源码，需要编译后注册
- `incomplete`: 配置或源码不完整

### 注册后状态
- `ready`: 插件就绪，等待执行
- `running`: 插件正在执行任务
- `stopped`: 插件已停止
- `error`: 插件运行错误

## 🔍 常用API

```bash
# 查看所有插件
curl http://localhost:8080/api/plugin/list

# 获取插件信息
curl "http://localhost:8080/api/plugin/get?id=com.example.my-plugin"

# 执行插件方法
curl -X POST http://localhost:8080/api/plugin/execute \
  -H "Content-Type: application/json" \
  -d '{"id": "com.example.my-plugin", "method": "greet", "payload": "Test"}'

# 注销插件
curl -X POST http://localhost:8080/api/plugin/unregister \
  -H "Content-Type: application/json" \
  -d '{"id": "com.example.my-plugin"}'
```

## ⚡ 自动编译

系统会在以下情况自动编译插件：
1. 检测到新的源码但没有对应二进制文件
2. 源码文件比二进制文件新

编译命令：`CGO_ENABLED=0 go build -o bin/plugin-name main.go`

## 🆘 常见问题

**Q: 插件编译失败？**  
A: 检查 `go.mod` 文件和依赖是否正确，确保代码符合插件接口规范。

**Q: 如何调试插件？**  
A: 可以直接运行 `go run main.go` 来调试，或查看系统日志。

**Q: 插件可以访问文件系统吗？**  
A: 可以，但建议限制访问范围，避免访问系统敏感目录。