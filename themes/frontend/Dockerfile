# 构建阶段
FROM node:18-alpine AS builder

WORKDIR /app

# 配置国内镜像源
RUN npm config set registry https://registry.npmmirror.com/ && \
    npm install -g pnpm@latest && \
    pnpm config set registry https://registry.npmmirror.com/

# 复制package文件
COPY package.json pnpm-lock.yaml ./

# 安装依赖
RUN pnpm install --no-frozen-lockfile

# 复制源代码
COPY . .

# 复制生产环境配置
COPY .env.production .env

# 构建应用
RUN rm -rf .vite dist && pnpm build

# 生产阶段
FROM nginx:alpine

# 复制构建结果
COPY --from=builder /app/dist /usr/share/nginx/html

# 配置nginx支持SPA路由
RUN echo 'server {' > /etc/nginx/conf.d/default.conf && \
    echo '    listen 80;' >> /etc/nginx/conf.d/default.conf && \
    echo '    listen [::]:80;' >> /etc/nginx/conf.d/default.conf && \
    echo '    server_name localhost;' >> /etc/nginx/conf.d/default.conf && \
    echo '    root /usr/share/nginx/html;' >> /etc/nginx/conf.d/default.conf && \
    echo '    index index.html index.htm;' >> /etc/nginx/conf.d/default.conf && \
    echo '    location / {' >> /etc/nginx/conf.d/default.conf && \
    echo '        try_files $uri $uri/ /index.html;' >> /etc/nginx/conf.d/default.conf && \
    echo '    }' >> /etc/nginx/conf.d/default.conf && \
    echo '}' >> /etc/nginx/conf.d/default.conf

EXPOSE 80

CMD ["nginx", "-g", "daemon off;"]