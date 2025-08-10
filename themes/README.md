# Jank 主题开发规范

基于主流分层架构和脚本化构建的主题开发规范。

## 🎯 系统架构

```bash
HTTP API → ThemeServiceImpl → ThemeManagerImpl → File System
                     ↓                    ↓
            Business Logic        Core Management
            (Build & Rebuild)     (Switch & Route)
```

**分层设计：**
- **Service 层**：处理业务逻辑，包括构建参数校验和 rebuild 逻辑
- **Manager 层**：纯粹的资源管理，接口统一为 `SwitchTheme(id string)`
- **Utils 层**：通用构建工具，支持脚本化构建流程

## 📁 目录结构

```
themes/theme-name/
├── theme.json            # 主题配置
├── package.json          # NPM依赖（现代前端主题）
├── index.html            # 源模板（现代前端主题）
├── src/                  # 源代码目录（现代前端主题）
├── scripts/              # 构建脚本（可选，无需构建的情况下不使用）
│   └── build.sh          # 构建脚本
└── dist/                 # 构建输出（服务器使用）
    ├── index.html        # 主题首页
    ├── assets/           # 静态资源
    └── vite.svg          # 主题图标
```

## ⚙️ theme.json 配置

```json
{
  "id": "theme-name",
  "name": "主题名称",
  "version": "1.0.0",
  "author": "作者",
  "description": "主题描述",
  "repository": "https://github.com/Done-0/jank-themes",
  "preview": "/assets/preview.png",
  "index_file_path": "/dist/index.html",
  "static_dir_path": "/dist/assets"
}
```

**配置字段说明：**
- `id`: 主题唯一标识
- `name`: 主题显示名称
- `version`: 版本号
- `author`: 作者
- `description`: 主题描述
- `repository`: 主题仓库地址（可选）
- `preview`: 主题预览图地址（可选）
- `index_file_path`: 主题入口文件路径（必须指向 `dist/` 目录）
- `static_dir_path`: 静态资源目录路径（通常为 `dist/assets`）

**配置要求：**
- **主题目录名必须与主题 ID 一致**：系统通过 ID 查找对应目录
- 路径必须指向 `dist/` 目录
- 主题 ID 必须唯一，不可与其他主题重复
- 所有路径使用相对路径格式

## 🔧 脚本化构建

### 构建约定
- 主题根目录存在 `scripts/build.sh` 时，系统将自动执行构建
- 构建脚本必须从 `theme.json` 读取所有配置，无硬编码路径
- 脚本在主题根目录下执行，可访问所有源文件和配置

### 构建触发时机
- 主题切换时，如果设置 `rebuild=true` 参数
- 主题初始化时，如果 `dist/` 目录不存在
- 开发者手动触发构建

## 🛠️ 开发流程

1. 创建主题目录
2. 编写 `theme.json` 配置
3. 开发主题文件：
   - **静态主题**：直接在 `dist/` 下开发
   - **现代前端主题**：创建 `scripts/build.sh` 构建脚本
4. 测试主题切换（自动执行构建）

## 📋 开发规范

- 所有主题文件放在 `dist/` 目录
- 静态资源放在 `dist/assets/`
- `theme.json` 路径必须指向 `dist/`
- 构建脚本完全配置驱动，无硬编码路径
- 修改源文件后必须重新构建
- 禁止直接编辑 `dist/` 目录内容

## 🌐 主题状态

### 主题状态标识符
- `ready`: 主题已加载，可以切换
- `active`: 当前激活主题
- `inactive`: 非激活状态
- `error`: 主题加载失败

## 🔧 API 接口

- `GET /api/theme/list` - 主题列表
- `POST /api/theme/switch` - 切换主题
- `GET /api/theme/get` - 当前主题
