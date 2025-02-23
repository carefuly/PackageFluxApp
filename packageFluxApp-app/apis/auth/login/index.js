// 导入二次封装axios
import sky from "@/utils/axios.js";

// 统一管理接口
const API = {
	CAPTCHA: "/v1/auth/send-login-captcha",
	LOGIN: "/v1/auth/email-login",
	LOGOUT: "/v1/auth/logout",
	USERINFO: "/v1/auth/userinfo",
}

// 发送验证码
export const captcha = (data) => sky({
	url: API.CAPTCHA,
	method: "post",
	data
});

// 登录
export const login = (data) => sky({
	url: API.LOGIN,
	method: "post",
	data
});

// 注销
export const logout = () => sky({
	url: API.LOGOUT,
	method: "post",
});

// 用户信息
export const userinfo = () => sky({
	url: API.USERINFO,
	method: "get",
});