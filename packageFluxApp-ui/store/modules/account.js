import {
	setStore,
} from "../index.js";

const ACCOUNT_KEY = 'SKY-ACCOUNT';

const setAccount = (account) => {
	let accountList = getAccount() || [];
	const flag = accountList.some(item => item.BookName === account.BookName);
	if (!flag) {
		accountList.push(account);
	}
	setStore(ACCOUNT_KEY, accountList);
};

const getAccount = () => {
	return uni.getStorageSync(ACCOUNT_KEY);
};

const clearAccount = () => {
	uni.removeStorageSync(ACCOUNT_KEY);
};

export {
	ACCOUNT_KEY,
	setAccount,
	getAccount,
	clearAccount
}