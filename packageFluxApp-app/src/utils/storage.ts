// @ts-ignore
import cookies from "js-cookie";
import {PINIA_PREFIX} from "@/config";

/**
 * 封装获取用户信息的方法
 */
export const getToken = () => {
  const skyUser = window.localStorage.getItem(PINIA_PREFIX + "user");
  if (skyUser != null && skyUser != "" && skyUser != undefined) {
    const parseSkyUser = JSON.parse(skyUser);
    // console.log("parseSkyUser",parseSkyUser)
    const keys = Object.keys(parseSkyUser);
    if (keys.includes("token")) {
      return parseSkyUser.token;
    } else {
      return "";
    }
  } else {
    return "";
  }
};

/**
 * 各个小仓库之间进行数据交互使用 OR 页面获取storage值使用
 * window.sessionStorage
 * @method set 设置会话缓存
 * @method get 获取会话缓存
 * @method remove 移除会话缓存
 * @method clear 移除全部会话缓存
 */
export const skySessionStorage = {
  put(key: string, value: any) {
    window.sessionStorage.setItem(PINIA_PREFIX + key, value);
  },
  set(key: string, value: any) {
    window.sessionStorage.setItem(PINIA_PREFIX + key, value);
  },
  get(key: string) {
    const value: any = window.sessionStorage.getItem(PINIA_PREFIX + key);
    return value;
  },
  remove(key: string) {
    window.sessionStorage.removeItem(PINIA_PREFIX + key);
  },
  clear() {
    window.sessionStorage.clear();
  },
  putJSON(key: string, jsonValue: any) {
    window.sessionStorage.put(PINIA_PREFIX + key, JSON.stringify(jsonValue));
  },
  setJSON(key: string, jsonValue: any) {
    window.sessionStorage.put(PINIA_PREFIX + key, JSON.stringify(jsonValue));
  },
  getJSON(key: string) {
    const jsonValue: any = window.sessionStorage.get(PINIA_PREFIX + key);
    return JSON.parse(jsonValue);
  }
};

/**
 * window.localStorage
 * @method set 设置
 * @method get 获取
 * @method remove 移除
 * @method clear 移除全部
 */
export const skyLocalStorage = {
  put(key: string, value: any) {
    window.localStorage.setItem(PINIA_PREFIX + key, value);
  },
  set(key: string, value: any) {
    window.localStorage.setItem(PINIA_PREFIX + key, value);
  },
  get(key: string) {
    const value: any = window.localStorage.getItem(PINIA_PREFIX + key);
    return value;
  },
  remove(key: string) {
    window.localStorage.removeItem(PINIA_PREFIX + key);
  },
  clear() {
    window.localStorage.clear();
  },
  putJSON(key: string, jsonValue: any) {
    window.localStorage.put(PINIA_PREFIX + key, JSON.stringify(jsonValue));
  },
  setJSON(key: string, jsonValue: any) {
    window.localStorage.put(PINIA_PREFIX + key, JSON.stringify(jsonValue));
  },
  getJSON(key: string) {
    const jsonValue: any = window.localStorage.get(PINIA_PREFIX + key);
    return JSON.parse(jsonValue);
  }
};

/**
 * cookies
 * @method set 设置
 * @method get 获取
 * @method remove 移除
 */
export const skyCookie = {
  put(key: string, value: any) {
    cookies.set(PINIA_PREFIX + key, value);
  },
  set(key: string, value: any) {
    cookies.set(PINIA_PREFIX + key, value);
  },
  get(key: string) {
    return cookies.get(PINIA_PREFIX + key);
  },
  remove(key: string) {
    cookies.remove(PINIA_PREFIX + key);
  }
};
