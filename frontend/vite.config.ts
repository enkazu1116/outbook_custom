import { defineConfig } from 'vite'
import react from '@vitejs/plugin-react'

// https://vite.dev/config/
export default defineConfig({
  plugins: [react()],
  server: {
    proxy: {
      // フロント(5173)からバック(1322)へのCORS回避用（dev）
      '/users': {
        target: 'http://localhost:1322',
        changeOrigin: true,
      },
    },
  },
})
