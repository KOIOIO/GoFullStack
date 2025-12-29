前端（Vue 3 + Vite）快速启动说明

1. 进入 frontend 目录

```bash
cd frontend
```

2. 安装依赖

```bash
npm install
```

3. 启动开发服务器

```bash
npm run dev
```

默认开发服务器端口为 `5173`。后端 API 假定运行在 `http://localhost:8080`，如有不同请在源码中修改 `src/views/*.vue` 中的 URL。
如果后端运行在 `http://localhost:8888`（如你的情况），已在 `vite.config.js` 中配置代理将 `/login` 和 `/register` 转发到 `http://localhost:8888`，避免浏览器跨域预检（OPTIONS）返回 405 的问题。
