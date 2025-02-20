const PROXY_KEY = 'SKY-PROXY';

const setProxy = (proxy) => {
	uni.setStorageSync(PROXY_KEY, proxy);
}

const getProxy = () => {
	return uni.getStorageSync(PROXY_KEY);
}

const clearProxy = () => {
	uni.removeStorageSync(PROXY_KEY);
}

export {
	PROXY_KEY,
	setProxy,
	getProxy,
	clearProxy
};