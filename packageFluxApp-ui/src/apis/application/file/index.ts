// 导入二次封装axios
import sky from "@/utils/axios";

// 统一管理接口
enum API {
  BATCH_UPLOAD = "/v1/application/file/batchUpload",
  BATCH_DELETE = "/v1/application/file/batchDelete",
  LIST_PAGE = "/v1/application/file/listPage",
  LIST_ALL = "/v1/application/file/listAll",
  GET_BY_ID = "/v1/application/file/getById/",
}

// 批量删除
export const batchDelete = (ids: any) => sky.post(API.BATCH_DELETE, ids);

// 分页列表查询
export const listPage = (data: any) => sky.get(API.LIST_PAGE, data);

// 列表查询
export const listAll = (data: any) => sky.get(API.LIST_ALL, data);

// 根据ID进行查询
export const getById = (id: any) => sky.get(API.GET_BY_ID + id);

