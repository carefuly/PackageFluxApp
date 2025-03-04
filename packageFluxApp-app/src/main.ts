import {createApp} from "vue";
import App from "./App.vue";
// 引入路由
// import router from "./router/index";
// 引入仓库pinia
// import store from "./store";
// 引入ElementPlus所有图标
import * as ElementPlusIconsVue from "@element-plus/icons-vue";
// 引入全局自定义指令
import SkyDirectives from "@/directives";

// SVG插件配置
import "virtual:svg-icons-register";
// 解决chrome下的warning问题
import "default-passive-events";

// 引入ElementPlus
import ElementPlus from "element-plus";
// 引入ElementPlus的css
import "element-plus/dist/index.css";
// @ts-ignore忽略当前文件ts类型的检测否则有红色提示(打包会失败)
import zhCn from "element-plus/dist/locale/zh-cn.mjs";
// 引入ElementPlus的暗黑模式css
import "element-plus/theme-chalk/dark/css-vars.css";

// 引入css
import "@/styles/index.scss";
import "@/styles/animation.scss";
// 引入animate
import "animate.css";
// UnoCss防止覆盖ElementPlus 的 el-button的css，所有放置最下方
import "uno.css";
import "amfe-flexible";

// 引入js

// 创建app
const app = createApp(App);
// 注册ElementPlus
app.use(ElementPlus, {
  locale: zhCn,
});
// 注册ElementPlus所有图标
for (const [key, component] of Object.entries(ElementPlusIconsVue)) {
  app.component(key, component);
}

// 注册路由
// app.use(router);
// 注册pinia
// app.use(store);
// 注册全局自定义指令
app.use(SkyDirectives);
// 挂载
app.mount("#app");
