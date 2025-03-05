/** 全局默认配置项 */

// 首页地址[默认]
export const HOME_URL: string = "/application";

// 登录页地址[默认]
export const LOGIN_URL: string = "/login";

// 注册页地址
export const REGISTER_URL: string = "/register";

// pinia仓库前缀
export const PINIA_PREFIX: string = "package-";

// 路由白名单地址[本地存在的路由 staticRouter.ts 中]
export const ROUTER_WHITE_LIST: string[] = ["/login", "/register", "/application/issue"];
