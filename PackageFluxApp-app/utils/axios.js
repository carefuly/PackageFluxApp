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
	getToken
} from "@/store/index.js";
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
		'Authorization': getToken(),
	};
	// 请求
	return new Promise((resolve, reject) => {
		skyShowLoading("加载中...");
		uni.request({
			url: baseURL + url,
			method: method,
			header: header,
			data: data,
			timeout: requestTimeout,
			success(res) {
				if (res.statusCode === 200) {
					if (res.data.code == 200) {
						resolve(res.data);
					} else {
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
							skyShowToast("请求地址不存在...");
							break;
						default:
							skyShowToast("请重试...");
							break;
					}
				}
			},
			fail(err) {
				if (err.errMsg.indexOf('request:fail') !== -1) {
					skyShowToast("网络异常", "none");
				} else {
					skyShowToast("未知异常", "none");
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