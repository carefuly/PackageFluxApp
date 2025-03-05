// 导入二次封装axios
import sky from "@/utils/axios";
// 统一管理接口
var API;
(function (API) {
    API["LOGIN"] = "/v1/auth/password-login";
    API["LOGOUT"] = "/v1/auth/logout";
    API["USERINFO"] = "/v1/auth/userinfo";
})(API || (API = {}));
// 登录
export const login = (data) => sky.post(API.LOGIN, data);
// 注销
export const logout = () => sky.post(API.LOGOUT);
// 用户信息
export const userinfo = () => sky.get(API.USERINFO);
//# sourceMappingURL=index.js.map