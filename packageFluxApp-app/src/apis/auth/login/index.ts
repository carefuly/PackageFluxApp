// 导入二次封装axios
import sky from "@/utils/axios";

// 统一管理接口
enum API {
  LOGIN = "/v1/auth/password-login",
  LOGOUT = "/v1/auth/logout",
  USERINFO = "/v1/auth/userinfo",
}

// 登录
export const login = (data: any) => sky.post(API.LOGIN, data);

// 注销
export const logout = () => sky.post(API.LOGOUT);

// 用户信息
export const userinfo = () => sky.get(API.USERINFO);
