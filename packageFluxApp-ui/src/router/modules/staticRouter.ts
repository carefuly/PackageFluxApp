import Layout from "@/layout/index.vue";
import { RouteRecordRaw } from "vue-router";
import { HOME_URL, LOGIN_URL, REGISTER_URL } from "@/config";

/**
 * LayoutRouter (布局路由)
 */
export const layoutRouter: RouteRecordRaw[] = [
  {
    path: "/",
    name: "Layout",
    component: Layout,
    redirect: HOME_URL,
    children: [
      {
        path: HOME_URL, // [唯一]
        component: () => import("@/views/application/details/index.vue"),
        meta: {
          title: "应用详情", // 标题
        }
      },
      // {
      //   path: "/application/statistic/:id",
      //   name: 'Statistic',
      //   component: () => import("@/views/application/statistic/index.vue"),
      //   meta: {
      //     title: "应用统计",
      //   },
      // },
      {
        path: "/application/version/:id",
        name: 'Version',
        component: () => import("@/views/application/version/index.vue"),
        meta: {
          title: "应用版本",
        },
      },
      // {
      //   path: "/application/log/:id",
      //   name: 'Log',
      //   component: () => import("@/views/application/log/index.vue"),
      //   meta: {
      //     title: "应用版本日志",
      //   },
      // },
    ]
  },
  {
    path: REGISTER_URL,
    name: "Register",
    component: () => import("@/views/auth/register/index.vue"),
    meta: {
      title: "注册",
      requireAuth: false,
    }
  },
  {
    path: LOGIN_URL,
    name: "Login",
    component: () => import("@/views/auth/login/index.vue"),
    meta: {
      title: "登录",
      requireAuth: false,
    }
  },
  // {
  //   path: "/application/issue/:id",
  //   name: 'Issue',
  //   component: () => import("@/views/application/issue/index.vue"),
  //   meta: {
  //     title: "应用版本发布页",
  //   },
  // },
];
