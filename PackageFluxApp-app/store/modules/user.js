const USER_KEY = 'pageage_user';

const setUser = (token) => {
	uni.setStorageSync(USER_KEY, token);
}

const getUser = () => {
	return uni.getStorageSync(USER_KEY);
}

const clearUser = () => {
	uni.removeStorageSync(USER_KEY);
}

export {
	USER_KEY,
	setUser,
	getUser,
	clearUser
};