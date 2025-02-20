import {
	netConfig
} from "@/config/network.js";
const {
	baseURL,
	contentType,
	requestTimeout,
	successCode
} = netConfig;
import {
	skyShowToast,
	skyShowModal,
	skyShowLoading,
} from "@/utils/sky.js";

const sky = (params) => {
	const url = params.url;
	const method = params.method;
	const data = params.data || {};
	const header = {
		'Content-Type': contentType,
	};
	// 请求
	return new Promise((resolve, reject) => {
		uni.request({
			url: baseURL + url,
			method: method,
			header: header,
			data: data,
			timeout: requestTimeout,
			success(response) {
				const res = response;
				if (res.statusCode >= 0) {
					if (res.data.status >= 0) {
						// 返回数据
						resolve(res.data);
					} else {
						skyShowToast(res.data.msg, "error");
						reject(res.data);
					}
				} else {
					switch (res.statusCode) {
						case 401:
							showModal("Token为空，请先登录", false).then(res => {
								uni.reLaunch({
									url: '/pages/user/login/index',
								});
							});
							break;
						case 404:
							showToast("请求地址不存在...");
							break;
						default:
							showToast("请重试...");
							break;
					}
				}
			},
			fail(err) {
				console.log(err);
				if (err.errMsg.indexOf('request:fail') !== -1) {
					showToast("网络异常", "none");
				} else {
					showToast("未知异常", "none");
				}
				// 错误信息
				reject(err);
			},
			complete() {
				// 不管成功还是失败都会执行
				uni.hideLoading();
				uni.hideToast();
			}
		})
	});
};

export default sky;