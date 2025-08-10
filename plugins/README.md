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

func (p *MyPlugin) Execute(ctx context.Context, method string, args map[string]any) (map[string]any, error) {
    switch method {
    case "greet":
        name, _ := args["name"].(string)
        return map[string]any{"message": "Hello, " + name + "!"}, nil
    case "calculate":
        a, _ := args["a"].(float64)
        b, _ := args["b"].(float64)
        return map[string]any{
            "result": a + b,
            "type": "addition",
        }, nil
    case "info":
        return map[string]any{
            "version": "1.0.0",
            "features": []string{"greet", "calculate", "info"},
            "config": map[string]any{
                "max_connections": 100,
                "timeout": "30s",
            },
        }, nil
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
  "name": "Example Plugin",
  "version": "1.0.0",
  "author": "Done-0",
  "description": "A simple hello world plugin for demonstration",
  "repository": "https://github.com/Done-0/jank-hello-world-plugin",
  "binary": "./bin/hello-world",
  "type": "handler",
  "auto_start": true,
  "start_timeout": 30000,
  "min_port": 10000,
  "max_port": 25000,
  "auto_mtls": true,
  "managed": true
}
```

### 4. 测试插件
```bash
# 测试插件
curl -X POST http://localhost:8080/api/plugin/execute \
  -H "Content-Type: application/json" \
  -d '{
    "id": "com.example.my-plugin",
    "method": "greet",
    "args": {"name": "Claude", "age": 25}
  }'
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
  "version": "1.0.0",                  // 版本号
  "author": "Author",                  // 作者
  "description": "Plugin description", // 描述
  "repository": "https://github.com/Done-0/example-plugin", // 插件仓库地址
  "binary": "plugin-name",             // 二进制文件地址
  "type": "handler",                   // 类型：provider/filter/handler/notifier
  "auto_start": true,                  // 是否自动启动
  "start_timeout": 30000,              // 启动超时时间
  "min_port": 10000,                   // 最小端口
  "max_port": 25000,                   // 最大端口
  "auto_mtls": true,                   // 是否自动启用 mTLS
  "managed": true                      // 是否由系统管理
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
func (p *MyPlugin) Execute(ctx context.Context, method string, args map[string]any) (map[string]any, error) {
    // 验证输入
    if method == "" {
        return nil, fmt.Errorf("method cannot be empty")
    }
    
    // 类型安全的参数获取
    name, ok := args["name"].(string)
    if !ok {
        return map[string]any{"error": "name must be a string"}, nil
    }
    
    // 处理可选参数
    age := 0
    if ageVal, exists := args["age"]; exists {
        if ageFloat, ok := ageVal.(float64); ok {
            age = int(ageFloat)
        }
    }
    
    return map[string]any{
        "name": name,
        "age": age,
        "timestamp": time.Now().Unix(),
    }, nil
}
```

### 复杂数据类型处理
```go
func (p *MyPlugin) Execute(ctx context.Context, method string, args map[string]any) (map[string]any, error) {
    switch method {
    case "process_user":
        // 处理嵌套对象
        user, ok := args["user"].(map[string]any)
        if !ok {
            return nil, fmt.Errorf("user must be an object")
        }
        
        // 处理数组
        tags, _ := args["tags"].([]any)
        tagStrings := make([]string, 0, len(tags))
        for _, tag := range tags {
            if tagStr, ok := tag.(string); ok {
                tagStrings = append(tagStrings, tagStr)
            }
        }
        
        return map[string]any{
            "processed_user": user,
            "tag_count": len(tagStrings),
            "tags": tagStrings,
        }, nil
    }
    
    return nil, fmt.Errorf("unknown method: %s", method)
}
```

### 上下文处理
```go
func (p *MyPlugin) Execute(ctx context.Context, method string, args map[string]any) (map[string]any, error) {
    select {
    case <-ctx.Done():
        return nil, ctx.Err()
    default:
        // 执行业务逻辑
    }
    
    // 支持超时控制的长时间操作
    timeout := time.Second * 10
    if timeoutVal, exists := args["timeout"]; exists {
        if timeoutStr, ok := timeoutVal.(string); ok {
            if parsedTimeout, err := time.ParseDuration(timeoutStr); err == nil {
                timeout = parsedTimeout
            }
        }
    }
    
    ctx, cancel := context.WithTimeout(ctx, timeout)
    defer cancel()
    
    // 实际业务逻辑...
    return map[string]any{"status": "completed"}, nil
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
curl "http://localhost:8080/api/plugin/get?id=dev.jank.plugins.hello-world"

# 执行插件方法
curl -X POST "http://127.0.0.1:8080/api/plugin/execute" \
  -H "Content-Type: application/json" \
  -d '{
    "id": "dev.jank.plugins.hello-world",
    "method": "greet", 
    "args": {
      "name": "World", 
      "age": 25, 
      "city": "Beijing"
    }
  }'

# 执行复杂数据处理
curl -X POST "http://127.0.0.1:8080/api/plugin/execute" \
  -H "Content-Type: application/json" \
  -d '{
    "id": "dev.jank.plugins.hello-world",
    "method": "process_user",
    "args": {
      "user": {
        "id": 123,
        "name": "John Doe",
        "active": true
      },
      "tags": ["developer", "go", "backend"],
      "timeout": "30s"
    }
  }'

curl -X POST "http://127.0.0.1:8080/api/plugin/execute" \
  -H "Content-Type: application/json" \
  -d '{
    "id": "dev.jank.plugins.hello-world",
    "method": "info"
  }'

# 注销插件
curl -X POST http://localhost:8080/api/plugin/unregister \
  -H "Content-Type: application/json" \
  -d '{"id": "dev.jank.plugins.hello-world"}'
```

## 🔧 新特性说明

### 1. 灵活的数据类型支持
- **之前**: 只支持 `map[string]string`
- **现在**: 支持 `map[string]any`，可以传递复杂的 JSON 对象、数组、数字、布尔值等

### 2. 高性能类型转换
- 使用 `google.protobuf.Any` 类型进行数据传输
- 内置 converter 包实现高效的类型转换
- 基于 `structpb` 优化性能

### 3. 更好的开发体验
- 支持嵌套对象和数组
- 类型安全的参数访问
- 更灵活的返回值结构

## ⚡ 自动编译

系统会在以下情况自动编译插件：
1. 检测到新的源码但没有对应二进制文件
2. 源码文件比二进制文件新

编译命令：`CGO_ENABLED=0 go build -o bin/plugin-name main.go`

## 🆘 常见问题

**Q: 插件编译失败？**  
A: 检查 `go.mod` 文件和依赖是否正确，确保代码符合插件接口规范。

**Q: 如何处理复杂的数据类型？**  
A: 使用类型断言安全地处理 `map[string]any` 中的各种数据类型，参考上面的复杂数据类型处理示例。

**Q: 插件可以访问文件系统吗？**  
A: 可以，但建议限制访问范围，避免访问系统敏感目录。

**Q: 如何处理 JSON 数字精度问题？**  
A: JSON 数字会被解析为 `float64`，如需整数请使用类型转换：`int(val.(float64))`。