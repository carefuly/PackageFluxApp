import {
	setStore,
	getStore,
	clearStore
} from "../index.js";

const TOKEN_KEY = "package_token";
const USER_KEY = 'package_user';

const setToken = (token) => {
	setStore(TOKEN_KEY, token);
};

const setUser = (user) => {
	setStore(USER_KEY, user);
}

const getToken = () => {
	return getStore(TOKEN_KEY);
}

const getUser = () => {
	return getStore(USER_KEY);
}

const clearToken = () => {
	clearStore(TOKEN_KEY);
}

const clearUser = () => {
	clearStore(USER_KEY);
}

export {
	TOKEN_KEY,
	USER_KEY,
	setToken,
	setUser,
	getToken,
	getUser,
	clearToken,
	clearUser,
};