// 导入二次封装axios
import sky from "@/utils/axios";
// 统一管理接口
var API;
(function (API) {
    API["CAPTCHA"] = "/v1/auth/send-email-captcha";
})(API || (API = {}));
// 发送验证码
export const captcha = (data) => sky.post(API.CAPTCHA, data);
//# sourceMappingURL=index.js.map