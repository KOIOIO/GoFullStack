import { defineConfig } from 'vite'
import vue from '@vitejs/plugin-vue'

export default defineConfig({
  plugins: [vue()]
  ,
  server: {
    proxy: {
      // 将 /login 和 /register 转发到后端 8888，避免浏览器跨域预检问题
      '/login': {
        target: 'http://localhost:8888',
        changeOrigin: true,
        secure: false
      },
      '/register': {
        target: 'http://localhost:8888',
        changeOrigin: true,
        secure: false
      }
    }
  }
})
