// 导入二次封装axios
import sky from "@/utils/axios";

// 统一管理接口
enum API {
  ADD = "/application/app/version/add",
  DELETE = "/application/app/version/del/",
  UPDATE = "/application/app/version/update/",
  SET_FORMAL = "/application/app/version/appFormal",
  LIST_ALL = "/application/app/version/listAll",
  GET_BY_ID = "/application/app/version/getById/",
}

// 添加
export const add = (data: any) => sky.post(API.ADD, data);

// 删除
export const deleteById = (id: any) => sky.delete(API.DELETE + id);

// 更新
export const updateById = (id: any, data: any) => sky.put(API.UPDATE + id, data);

// 设置正式版
export const setFormal = (data: any) => sky.put(API.SET_FORMAL, data);

// 列表查询
export const listAll = (data: any) => sky.get(API.LIST_ALL, data);

// 根据ID进行查询
export const getById = (id: any) => sky.get(API.GET_BY_ID + id);

