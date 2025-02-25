import path from "path";
import {defineConfig, loadEnv} from 'vite';
import createVitePlugins from './build/vite/plugin';

// https://vitejs.dev/config/
export default defineConfig(({command, mode}) => {
  // 获取配置文件别名配置
  const env = loadEnv(mode, process.cwd());

  return {
    // 开发或生产环境服务的公共基础路径
    base: env.VITE_BUILD_PATH,
    // 路径别名
    resolve: {
      alias: {
        "@": path.resolve("./src"), // 相对路径别名配置，使用 @ 代替 src
        "~": path.resolve("./src")
      }
    },
    // 引入sass全局样式变量
    css: {
      // css全局变量使用，@/styles/variable.scss文件
      preprocessorOptions: {
        scss: {
          javascriptEnabled: true,
          additionalData: '@import "./src/styles/variable.scss";'
        }
      }
    },
    // 服务配置跨域
    server: {
      host: "localhost", // 配置localhost、本机IP地址
      port: 9090, // 端口号
      hmr: true, // 热更新
      open: false, // 服务启动时是否自动打开浏览器
      proxy: {
        // 代理跨域
        [env.VITE_WEB_BASE_API]: {
          // 配置哪个环境下的
          target: env.VITE_SERVER,
          // 允许跨域
          changeOrigin: true,
          // 支持https
          secure: false,
          // 路径重写，例如：将路径中包含dev-api字段替换为空。注意：只有请求真实后端接口才会有用，使用mock接口还是得带koi
          rewrite: path => path.replace(new RegExp("^" + env.VITE_WEB_BASE_API), ""),
        }
      }
    },
    // 插件
    plugins: createVitePlugins(env, command === 'build'),
    // 构建
    build: {
      chunkSizeWarningLimit: 2000, // 消除打包大小超过500kb警告
      outDir: 'dist', // 指定打包路径，默认为项目根目录下的dist目录
      minify: 'terser', // Vite 2.6.x 以上需要配置 minify："terser"，terserOptions才能生效
      terserOptions: {
        compress: {
          keep_infinity: true, // 防止 Infinity 被压缩成 1/0，这可能会导致 Chrome 上的性能问题
          drop_console: env.VITE_DROP_CONSOLE, // 生产环境去除 console
          drop_debugger: env.VITE_DROP_CONSOLE // 生产环境去除 debugger
        },
        format: {
          comments: false // 删除注释
        }
      },
      // 静态资源打包到dist下的不同目录
      rollupOptions: {
        output: {
          chunkFileNames: 'static/js/[name]-[hash].js',
          entryFileNames: 'static/js/[name]-[hash].js',
          assetFileNames: 'static/[ext]/[name]-[hash].[ext]'
        }
      }
    },
    // 预编译，增加访问速度，针对node_modules
    optimizeDeps: {
      include: [
        "vue",
        "vue-router",
        "pinia",
        "axios",
        "@vueuse/core",
        "echarts",
        "vue-i18n",
        "element-plus/es/components/text/style/css",
        "element-plus/es/components/collapse-item/style/css",
        "element-plus/es/components/collapse/style/css",
        "element-plus/es/components/space/style/css",
        "element-plus/es/components/container/style/css",
        "element-plus/es/components/aside/style/css",
        "element-plus/es/components/main/style/css",
        "element-plus/es/components/header/style/css",
        "element-plus/es/components/button-group/style/css",
        "element-plus/es/components/radio-button/style/css",
        "element-plus/es/components/checkbox-group/style/css",
        "element-plus/es/components/form/style/css",
        "element-plus/es/components/form-item/style/css",
        "element-plus/es/components/button/style/css",
        "element-plus/es/components/input/style/css",
        "element-plus/es/components/input-number/style/css",
        "element-plus/es/components/switch/style/css",
        "element-plus/es/components/upload/style/css",
        "element-plus/es/components/menu/style/css",
        "element-plus/es/components/col/style/css",
        "element-plus/es/components/icon/style/css",
        "element-plus/es/components/row/style/css",
        "element-plus/es/components/tag/style/css",
        "element-plus/es/components/dialog/style/css",
        "element-plus/es/components/loading/style/css",
        "element-plus/es/components/radio/style/css",
        "element-plus/es/components/radio-group/style/css",
        "element-plus/es/components/popover/style/css",
        "element-plus/es/components/scrollbar/style/css",
        "element-plus/es/components/tooltip/style/css",
        "element-plus/es/components/dropdown/style/css",
        "element-plus/es/components/dropdown-menu/style/css",
        "element-plus/es/components/dropdown-item/style/css",
        "element-plus/es/components/sub-menu/style/css",
        "element-plus/es/components/menu-item/style/css",
        "element-plus/es/components/divider/style/css",
        "element-plus/es/components/card/style/css",
        "element-plus/es/components/link/style/css",
        "element-plus/es/components/breadcrumb/style/css",
        "element-plus/es/components/breadcrumb-item/style/css",
        "element-plus/es/components/table/style/css",
        "element-plus/es/components/tree-select/style/css",
        "element-plus/es/components/table-column/style/css",
        "element-plus/es/components/select/style/css",
        "element-plus/es/components/option/style/css",
        "element-plus/es/components/pagination/style/css",
        "element-plus/es/components/tree/style/css",
        "element-plus/es/components/alert/style/css",
        "element-plus/es/components/checkbox/style/css",
        "element-plus/es/components/date-picker/style/css",
        "element-plus/es/components/transfer/style/css",
        "element-plus/es/components/tabs/style/css",
        "element-plus/es/components/image/style/css",
        "element-plus/es/components/tab-pane/style/css"
      ]
    }
  };
});
