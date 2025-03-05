// 导入二次封装axios
import sky from "@/utils/axios";
// 统一管理接口
var API;
(function (API) {
    API["REGISTER"] = "/v1/auth/password-register";
})(API || (API = {}));
// 注册
export const register = (data) => sky.post(API.REGISTER, data);
//# sourceMappingURL=index.js.map