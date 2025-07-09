import { defineConfig } from 'vite'
import react from '@vitejs/plugin-react'

export default defineConfig({
   plugins: [react()],
   server: {
      port: 5173,
      proxy: {
         // Proxy ALL /auth/* requests to your Go server on :8888
         '/auth': {
            target: 'http://localhost:8888',
            changeOrigin: true,
            secure: false,
         },
         // If you have other API paths like /v1/api/, add them here too:
         '/v1/api': {
            target: 'http://localhost:8888',
            changeOrigin: true,
            secure: false,
         },
      },
   },
})
