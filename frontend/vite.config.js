import { defineConfig } from 'vite'

export default defineConfig({
  root: 'assets',
  build: {
    outDir: '../assets/dist',
    emptyOutDir: true,
    rollupOptions: {
      output: {
        manualChunks(id) {
          if (id.includes('node_modules')) {
            return 'vendor';
          }
        }
      }
    }
  },
  plugins: [
  ]
})