export * from "./modules/user";

/**
 * 设置缓存信息
 */
const setStore = (state, data) => {
	uni.setStorageSync(state, data);
}

/**
 * 获取缓存信息
 */
const getStore = (state) => {
	return uni.getStorageSync(state);
}

/**
 * 清除缓存信息
 */
const clearStore = (state) => {
	uni.removeStorageSync(state);
}

export {
	getStore,
	setStore,
	clearStore,
};