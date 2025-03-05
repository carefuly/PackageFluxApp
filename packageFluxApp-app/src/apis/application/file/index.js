// 导入二次封装axios
import sky from "@/utils/axios";
// 统一管理接口
var API;
(function (API) {
    API["BATCH_UPLOAD"] = "/v1/application/file/batchUpload";
    API["BATCH_DELETE"] = "/v1/application/file/batchDelete";
    API["LIST_PAGE"] = "/v1/application/file/listPage";
    API["LIST_ALL"] = "/v1/application/file/listAll";
    API["GET_BY_ID"] = "/v1/application/file/getById/";
})(API || (API = {}));
// 批量删除
export const batchDelete = (ids) => sky.post(API.BATCH_DELETE, ids);
// 分页列表查询
export const listPage = (data) => sky.get(API.LIST_PAGE, data);
// 列表查询
export const listAll = (data) => sky.get(API.LIST_ALL, data);
// 根据ID进行查询
export const getById = (id) => sky.get(API.GET_BY_ID + id);
//# sourceMappingURL=index.js.map