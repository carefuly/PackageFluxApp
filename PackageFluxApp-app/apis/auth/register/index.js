// 导入二次封装axios
import sky from "@/utils/axios.js";

// 统一管理接口
const API = {
	CAPTCHA: "/v1/auth/send-register-captcha",
	REGISTER: "/v1/auth/email-register",
};

// 发送验证码
export const captcha = (data) => sky({
	url: API.CAPTCHA,
	method: "post",
	data
});

// 注册
export const register = (data) => sky({
	url: API.REGISTER,
	method: "post",
	data
});