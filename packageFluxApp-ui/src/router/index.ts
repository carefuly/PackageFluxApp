import NProgress from "@/utils/nprogress";
import {createRouter, createWebHashHistory, createWebHistory} from "vue-router";
import {LOGIN_URL, ROUTER_WHITE_LIST} from "@/config";
import {layoutRouter} from "@/router/modules/staticRouter";
import {useUserStore} from "@/store";
import {skyMsgWarning} from "@/utils/sky";

// .env配置文件读取
const mode = import.meta.env.VITE_ROUTER_MODE;

// 路由访问两种模式：带#号的哈希模式，正常路径的web模式。
const routerMode: any = {
  hash: () => createWebHashHistory(),
  history: () => createWebHistory()
};

// 创建路由器对象
const router = createRouter({
  // 路由模式hash或者默认不带#
  history: routerMode[mode](),
  routes: [...layoutRouter],
  // 滚动行为
  scrollBehavior() {
    return {
      left: 0,
      top: 0
    };
  }
});

/**
 * @description 前置路由
 */
router.beforeEach(async (to, from, next) => {
  const userStore = useUserStore();
  // NProgress 开始
  NProgress.start();
  // 标题切换，没有防止后置路由，是因为页面路径不存在，title会变成undefined
  const title = import.meta.env.VITE_WEB_TITLE;
  document.title = (to.meta.title || title);
  if (userStore.token) {
    // 获取token
    const token = atob(userStore.token.split(".")[1]);
    // 获取token解析对象
    const token_str = JSON.parse(token);
    const time = Date.now();
    if (token_str.exp * 1000 <= time) {
      if (to.path === "/login") {
        next();
      } else {
        skyMsgWarning("Token已过期，请先登录！🌻");
        next(LOGIN_URL);
      }
    } else {
      next();
    }
  } else {
    // 判断访问页面是否在路由白名单地址[静态路由]中，如果存在直接放行。
    if (ROUTER_WHITE_LIST.includes(to.path)) {
      next();
    } else {
      skyMsgWarning("账号身份不存在，请先登录🌻");
      next(LOGIN_URL);
    }
  }
});

/**
 * @description 路由跳转错误
 */
router.onError(error => {
  // 结束全屏动画
  NProgress.done();
  console.warn("路由错误", error.message);
});

/**
 * @description 后置路由
 */
// @ts-ignore
router.afterEach((to, from) => {
  // console.log("后置守卫", to, from);
  // 结束全屏动画
  NProgress.done();
});

export default router;
