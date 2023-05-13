npm i -g cordova

rimraf myApp

cordova create cordovacloud
cd cordovacloud
cordova platform add android --save
cordova platform add ios --save
cordova platform add browser  --save //创建浏览器文件

显示平台:
cordova platform ls

删除平台:

cordova platform remove ios

运行项目:
先编译vue-vite项目，其中需要修改vite.config.js，用来修改打包的路径
npm create vite@latest
```js
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

    build: {
        outDir: "../cordovacloud/www"
    }
})
```
还要修改index.html
```html
<body>
<div id="app"></div>
</body>
<script type="text/javascript" src="cordova.js"></script>
</html>
```

最后运行:
cordova run browser

http-server