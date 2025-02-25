import {
	getUser
} from "@/store/index.js";

/**
 * @description 封装统一公共接口参数方法
 * @param version 版本号
 * @param method 方法
 * @param param 参数
 * @returns {{method, param, jsonrpc: string, tags: {usertoken: UnwrapRef<string>, userid: UnwrapRef<string>}}}
 */
export function apiParam(version = "2.0", method, param) {
	const userStore = getUser();
	return {
		jsonrpc: version,
		method: method,
		params: param,
		tags: {
			userid: userStore.userid,
			usertoken: userStore.usertoken,
		},
	}
}

/**
 * 获取当前时间，并根据指定的格式字符串格式化输出。
 * @param {string} format - 输出日期的格式字符串（默认为 'yyyy/MM/dd hh:mm:ss'）。
 * @returns {string} 格式化后的当前时间字符串。
 */
export const getTime = (format = 'yyyy/MM/dd hh:mm:ss') => {
	const date = new Date();
	const year = date.getFullYear();
	const month = ('0' + (date.getMonth() + 1)).slice(-2);
	const day = ('0' + date.getDate()).slice(-2);
	const hours = ('0' + date.getHours()).slice(-2);
	const minutes = ('0' + date.getMinutes()).slice(-2);
	const seconds = ('0' + date.getSeconds()).slice(-2);

	return format
		.replace('yyyy', year)
		.replace('MM', month)
		.replace('dd', day)
		.replace('hh', hours)
		.replace('mm', minutes)
		.replace('ss', seconds);
};