import {defineConfig} from 'vite'
import vue from '@vitejs/plugin-vue'
import {resolve} from 'path';

// https://vitejs.dev/config/
export default defineConfig({
    plugins: [vue()],
    server: {
        host: '0.0.0.0'
    },
    resolve: {
        alias: {'@': resolve(__dirname, './src')}
    },
    define: {
        'process.env': {}
    },
    css: {
        // css预处理器
        preprocessorOptions: {
            scss: {
                additionalData: `@import '@/styles/variables.scss';`
            }
        }
    },
    build: {
        outDir: "../cordovacloud/www"
    }
})