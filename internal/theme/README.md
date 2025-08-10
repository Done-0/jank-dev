# Jank 主题系统

基于文件系统的主题管理架构，支持动态切换、现代前端构建工具和统一资源路由。

## 🎯 系统架构

```bash
HTTP API → ThemeServiceImpl → ThemeManagerImpl → File System
                                      ↓
Theme Files (dist/) ←→ Unified Static Route Handler
```

**核心组件：**
- `ThemeManagerImpl`: 主题生命周期管理
- `ThemeServiceImpl`: HTTP API服务层
- `ThemeInfo`: 主题元数据和运行时状态
- `Unified Route Handler`: 配置驱动的静态资源路由

## 🚀 核心特性

### 动态切换
支持运行时无重启切换主题，自动更新路由和静态资源映射。

### 现代前端支持
完整支持 React、Vite、Webpack 等现代构建工具：
```bash
themes/theme-name/src/ → npm run build → dist/ → 自动路由
```

### 统一资源路由
极简化的配置驱动路由，所有静态资源自动映射到主题构建目录。

### 配置持久化
主题切换状态自动持久化到配置文件，重启后恢复上次状态。

## 📁 目录结构

```bash
internal/theme/
├── impl/
│   ├── theme_manager.go       # 核心管理器实现
│   └── theme_info.go          # 主题信息结构
├── theme.go                   # 接口定义
└── README.md                  # 本文档

pkg/router/routes/
└── theme.go                   # 统一路由处理器

themes/                        # 主题存放目录
├── default/                   # 静态主题示例
│   ├── theme.json            # 主题配置
│   └── dist/                 # 构建输出目录
│       ├── index.html        # 主题首页
│       └── assets/           # 静态资源
└── moon/                     # React主题示例
    ├── theme.json            # 主题配置
    ├── package.json          # NPM依赖
    ├── index.html            # 源模板
    ├── src/                  # React源代码
    │   ├── main.tsx
    │   └── App.tsx
    └── dist/                 # Vite构建输出
        ├── index.html        # 构建后的HTML
        ├── assets/           # 构建后的资源
        └── vite.svg          # 主题图标
```

## ⚙️ 配置文件

### theme.json 统一格式
```json
{
  "id": "moon",
  "name": "Moon主题",
  "version": "1.0.0",
  "author": "Done-0",
  "description": "基于 React + Vite 的现代化主题",
  "index_file_path": "/dist/index.html",
  "static_dir_path": "/dist/assets"
}
```

**配置说明：**
- `index_file_path`: 主题入口文件，必须指向 `dist/` 目录
- `static_dir_path`: 静态资源目录，通常为 `dist/assets`
- **重要**: 所有路径都应指向构建输出目录，不是源文件目录

### 主题类型和结构

#### 静态主题（如 default）
```bash
themes/default/
├── theme.json               # 配置文件
└── dist/                    # 静态文件目录
    ├── index.html          # 主题首页
    └── assets/             # CSS/JS/图片等
```

#### 现代前端主题（如 moon）
```bash
themes/moon/
├── theme.json               # 配置文件
├── package.json            # NPM依赖
├── index.html              # 源模板
├── src/                    # 源代码
└── dist/                   # 构建输出（服务器使用）
```

### 主题状态标识符
- `ready`: 主题已加载，可以切换
- `active`: 当前激活主题
- `inactive`: 非激活状态
- `error`: 主题加载失败

## 🔧 核心接口

### ThemeManager 接口
```go
type ThemeManager interface {
    LoadTheme(themeID string) error
    SwitchTheme(themeID string) error
    GetActiveTheme() *ThemeInfo
    ListThemes() []*ThemeInfo
    Shutdown() error
}
```

### ThemeInfo 结构
```go
type ThemeInfo struct {
    ID              string `json:"id"`
    Name            string `json:"name"`
    Version         string `json:"version"`
    Author          string `json:"author"`
    Description     string `json:"description"`
    IndexFilePath   string `json:"index_file_path"`
    StaticDirPath   string `json:"static_dir_path"`
    IsActive        bool   `json:"is_active"`
    Status          string `json:"status"`
    ThemePath       string `json:"theme_path"`
}
```

## 🌟 API接口

### 切换主题
```bash
POST /api/theme/switch
Content-Type: application/json

{
  "id": "moon"
}
```

### 获取当前主题
```bash
GET /api/theme/get
```

### 列举所有主题
```bash
GET /api/theme/list?page_no=1&page_size=100
```

## 🔄 主题切换流程

1. **接收切换请求**: HTTP API接收主题切换请求
2. **验证主题存在**: 检查目标主题是否存在和有效
3. **动态加载**: 如果主题未加载，从文件系统动态加载
4. **更新状态**: 停用当前主题，激活新主题
5. **更新路由**: 重新映射静态资源路由到新主题的 `dist/` 目录
6. **持久化配置**: 保存新的激活主题到配置文件
7. **返回响应**: 返回切换成功响应
8. **前端刷新**: 前端接收成功响应后自动刷新页面

## 🛠️ 开发流程

### 静态主题开发
1. 创建主题目录：`themes/your-theme/`
2. 编写 `theme.json` 配置文件
3. 在 `dist/` 目录下放置 HTML、CSS、JS 文件
4. 测试主题切换功能

### 现代前端主题开发
1. 创建主题目录：`themes/your-theme/`
2. 初始化前端项目：`npm init` + 安装依赖
3. 编写源代码在 `src/` 目录
4. 配置构建脚本：`npm run build`
5. **重要**: 修改源文件后必须重新构建
6. 编写 `theme.json`，路径指向 `dist/` 目录
7. 测试主题切换功能

### Favicon 最佳实践

#### 推荐方案：内联 SVG
```html
<link rel="icon" href="data:image/svg+xml,<svg xmlns='http://www.w3.org/2000/svg'>...</svg>">
```

#### 外部文件方案
```html
<!-- 添加版本参数避免缓存问题 -->
<link rel="icon" type="image/svg+xml" href="/vite.svg?v=1" />
```

## 🛡️ 错误处理

系统提供完善的错误处理机制：

1. 主题加载失败自动回退到默认主题
2. 配置文件损坏时使用内置默认配置
3. 主题切换失败时保持当前主题不变
4. 详细的错误日志记录
5. 构建文件缺失时提示重新构建

## 🔧 路由实现

### 统一资源路由逻辑
```go
// 极简化的配置驱动路由
requestedFile := strings.TrimPrefix(string(c.Path()), "/")
buildDir := filepath.Dir(activeTheme.IndexFilePath)
fullPath := filepath.Join(activeTheme.Path, buildDir, requestedFile)
c.File(fullPath)
```

**路由特性：**
- 配置驱动：根据 `theme.json` 的 `index_file_path` 确定构建目录
- 统一映射：所有请求自动映射到主题的构建目录
- 无硬编码：不依赖特定的目录名称或路径
- 现代兼容：完整支持 Vite、Webpack 等构建工具的输出

## ⚠️ 重要注意事项

1. **禁止直接编辑 `dist/` 目录**：构建产物应通过构建命令生成
2. **修改源文件后必须重新构建**：确保 `dist/` 目录是最新的
3. **路径配置必须指向 `dist/`**：`theme.json` 中的路径应指向构建输出
