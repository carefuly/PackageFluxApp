/**
 * @description 配置axios请求基础信息
 * @author 青姝 2224693191@qq.com
 * @type {{baseURL: string, cors: boolean, messageDuration: number, noPermissionCode: number, invalidCode: number, contentType: string, requestTimeout: number, successCode: number[]}}
 */
export const netConfig = {
	// axios 基础url地址
	// baseURL: baseUrl.CD_MZXDJX,
	baseURL: import.meta.env.VIITE_GLOBAL_URL + import.meta.env.VIITE_BASE_URL,
	// 为开发服务器配置 CORS。默认启用并允许任何源，传递一个 选项对象 来调整行为或设为 false 表示禁用
	cors: false,
	// 是否在请求中协带cookie
	withCredentials: false,
	// 根据后端定义配置
	contentType: 'application/json;charset=UTF-8',
	// 消息框消失时间
	messageDuration: 3000,
	// 最长请求时间
	requestTimeout: import.meta.env.VITE_TIMEOUT,
	// 操作正常code，支持String、Array、int多种类型
	successCode: [200],
	// 登录失效code
	invalidCode: 400,
	// 无token认证code
	noPermissionCode: 401,
};