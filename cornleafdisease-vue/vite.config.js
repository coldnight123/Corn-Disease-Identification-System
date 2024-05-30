import { defineConfig } from 'vite'
import vue from '@vitejs/plugin-vue'
import path from 'path' //导入 node.js path

// https://vitejs.dev/config/
export default defineConfig({
  plugins: [vue()],
  server: {
    port: 8082, // 设置端口为3001
    proxy: {
      '/api': {
        target: 'http://127.0.0.1:8081',	//实际请求地址
        changeOrigin: true,
        rewrite: (path) => path.replace(/^\/api/, '')
      },
    }
  },
  resolve: {
    alias: { //配置路径别名
      '@': path.resolve(__dirname, 'src')
    }
  }
})
