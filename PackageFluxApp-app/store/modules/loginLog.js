const LOGINLOG_KEY = 'SKY-LoginLog';

const setLoginLog = (log) => {
	let loginLog = getLoginLog() || [];
	loginLog.push(log);
	uni.setStorageSync(LOGINLOG_KEY, loginLog);
}

const getLoginLog = () => {
	return uni.getStorageSync(LOGINLOG_KEY);
}

const clearLoginLog = () => {
	uni.removeStorageSync(LOGINLOG_KEY);
}

export {
	LOGINLOG_KEY,
	setLoginLog,
	getLoginLog,
	clearLoginLog
};