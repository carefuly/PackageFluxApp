// 导入二次封装axios
import sky from "@/utils/axios";

// 统一管理接口
enum API {
  ADD = "/application/app/file/add",
  DELETE = "/application/app/file/del/",
  BATCH_DELETE = "/application/app/file/batch/del",
  UPDATE = "/application/app/file/update/",
  LIST_PAGE = "/application/app/file/listPage",
  LIST_ALL = "/application/app/file/listAll",
  GET_BY_ID = "/application/app/file/getById/",
}

// 添加
export const add = (data: any) => sky.post(API.ADD, data);

// 删除
export const deleteById = (id: any) => sky.delete(API.DELETE + id);

// 批量删除
export const batchDelete = (ids: any) => sky.post(API.BATCH_DELETE, ids);

// 更新
export const updateById = (id: any, data: any) => sky.put(API.UPDATE + id, data);

// 分页列表查询
export const listPage = (data: any) => sky.get(API.LIST_PAGE, data);

// 列表查询
export const listAll = (data: any) => sky.get(API.LIST_ALL, data);

// 根据ID进行查询
export const getById = (id: any) => sky.get(API.GET_BY_ID + id);

