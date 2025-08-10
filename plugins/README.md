# Jank 插件开发规范

基于主流分层架构和脚本化构建的插件开发规范。

## 🎯 系统架构

```bash
HTTP API → PluginServiceImpl → PluginManagerImpl → hashicorp/go-plugin
                     ↓                    ↓
            Business Logic        Core Management
            (Build & Rebuild)     (Register & Switch)
```

**分层设计：**
- **Service 层**：处理业务逻辑，包括构建参数校验和 rebuild 逻辑
- **Manager 层**：纯粹的资源管理，接口统一为 `RegisterPlugin(id string)`
- **Utils 层**：通用构建工具，支持脚本化构建流程

## 📁 目录结构

```
plugins/plugin-name/
├── main.go           # 插件主代码
├── plugin.json       # 插件配置
├── go.mod           # Go模块文件
├── scripts/
│   └── build.sh     # 构建脚本
└── bin/             # 编译生成的二进制文件目录
    └── plugin-name
```

## ⚙️ plugin.json 配置

```json
{
  "id": "com.example.plugin-name",
  "name": "Plugin Name",
  "version": "1.0.0",
  "author": "Author",
  "description": "Plugin description",
  "repository": "https://github.com/Done-0/example-plugin",
  "binary": "./bin/plugin-name",
  "type": "handler",
  "auto_start": true,
  "start_timeout": 30000,
  "min_port": 10000,
  "max_port": 25000,
  "auto_mtls": true,
  "managed": true
}
```

**配置字段说明：**
- `id`: 插件唯一标识（反向域名格式）
- `name`: 插件显示名称
- `version`: 版本号
- `author`: 作者
- `description`: 插件描述
- `repository`: 插件仓库地址
- `binary`: 二进制文件路径（相对于插件根目录）
- `type`: 插件类型（provider/filter/handler/notifier）
- `auto_start`: 是否自动启动
- `start_timeout`: 启动超时时间（毫秒）
- `min_port`: 最小端口号
- `max_port`: 最大端口号
- `auto_mtls`: 是否自动启用 mTLS
- `managed`: 是否由系统管理

**插件类型：**
- `provider`: 数据提供者插件
- `filter`: 数据过滤插件
- `handler`: 业务处理插件
- `notifier`: 通知插件

## 🔧 脚本化构建

### 构建约定
- 插件根目录存在 `scripts/build.sh` 时，系统将自动执行构建
- 构建脚本必须从 `plugin.json` 读取所有配置，无硬编码路径
- 脚本在插件根目录下执行，可访问所有源文件和配置

### 构建触发时机
- 插件注册时，如果设置 `rebuild=true` 参数
- 插件初始化时，如果二进制文件不存在
- 开发者手动触发构建

### 构建脚本示例
创建 `scripts/build.sh`：
```bash
#!/bin/bash
set -e

# 从 plugin.json 读取配置
PLUGIN_ID=$(jq -r '.id' plugin.json)
BINARY_PATH=$(jq -r '.binary' plugin.json)

echo "Building plugin: $PLUGIN_ID"
echo "Output binary: $BINARY_PATH"

# 确保输出目录存在
mkdir -p "$(dirname "$BINARY_PATH")"

# 编译插件
CGO_ENABLED=0 go build -o "$BINARY_PATH" main.go

echo "Build completed: $BINARY_PATH"
```

## 🛠️ 开发流程

1. 创建插件目录
2. 编写 `plugin.json` 配置
3. 实现插件接口（Execute、HealthCheck）
4. 创建 `scripts/build.sh` 构建脚本
5. 测试插件功能

## 📋 开发规范

- **插件目录名必须与插件 ID 一致**：系统通过 ID 查找对应目录
- 插件 ID 建议使用简洁的命名（如 `hello-world`）
- 二进制文件路径必须相对于插件根目录
- 构建脚本完全配置驱动，无硬编码路径
- 插件必须实现 Execute 和 HealthCheck 方法
- 支持 `map[string]any` 参数类型

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

## 🔧 API 接口

- `GET /api/plugin/list` - 插件列表
- `POST /api/plugin/register` - 注册插件
- `POST /api/plugin/unregister` - 注销插件
- `POST /api/plugin/execute` - 执行插件方法
- `GET /api/plugin/get` - 获取插件信息