# Jank 主题开发规范

## 📁 目录结构

```
themes/theme-name/
├── theme.json
└── dist/                  # 构建输出（服务器使用）
    ├── index.html
    └── assets/
```

## ⚙️ theme.json 配置

```json
{
  "id": "theme-name",
  "name": "主题名称",
  "version": "1.0.0",
  "author": "作者",
  "description": "描述",
  "index_file_path": "/dist/index.html",
  "static_dir_path": "/dist/assets"
}
```

**要求**：
- 路径必须指向 `dist/` 目录
- 主题 ID 必须唯一，不可与其他主题重复

## 🛠️ 开发流程

1. 创建主题目录
2. 编写 `theme.json`
3. 在 `dist/` 下开发或构建主题文件
4. 实现主题切换时刷新页面的逻辑

## 📋 开发规范

- 所有主题文件放在 `dist/` 目录
- 静态资源放在 `dist/assets/`
- `theme.json` 路径必须指向 `dist/`

## 🔧 API 接口

- `GET /api/theme/list` - 主题列表
- `POST /api/theme/switch` - 切换主题
- `GET /api/theme/get` - 当前主题
