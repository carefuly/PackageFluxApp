// 导入二次封装axios
import sky from "@/utils/axios";

// 统一管理接口
enum API {
  ADD = "/v1/application/details/create",
  DELETE = "/v1/application/details/delete/",
  UPDATE = "/v1/application/details/update/",
  GET_BY_ID = "/v1/application/details/getById/",
  LIST_ALL = "/v1/application/details/listAll",
}

// 添加
export const add = (data: any) => sky.post(API.ADD, data);

// 删除
export const deleteById = (id: any) => sky.delete(API.DELETE + id);

// 更新
export const updateById = (id: any, data: any) => sky.put(API.UPDATE + id, data);

// 根据ID进行查询
export const getById = (id: any) => sky.get(API.GET_BY_ID + id);

// 列表查询
export const listAll = (data: any) => sky.get(API.LIST_ALL, data);

