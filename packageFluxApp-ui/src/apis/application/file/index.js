// 导入二次封装axios
import sky from "@/utils/axios";
// 统一管理接口
var API;
(function (API) {
    API["ADD"] = "/application/app/file/add";
    API["DELETE"] = "/application/app/file/del/";
    API["BATCH_DELETE"] = "/application/app/file/batch/del";
    API["UPDATE"] = "/application/app/file/update/";
    API["LIST_PAGE"] = "/application/app/file/listPage";
    API["LIST_ALL"] = "/application/app/file/listAll";
    API["GET_BY_ID"] = "/application/app/file/getById/";
})(API || (API = {}));
// 添加
export const add = (data) => sky.post(API.ADD, data);
// 删除
export const deleteById = (id) => sky.delete(API.DELETE + id);
// 批量删除
export const batchDelete = (ids) => sky.post(API.BATCH_DELETE, ids);
// 更新
export const updateById = (id, data) => sky.put(API.UPDATE + id, data);
// 分页列表查询
export const listPage = (data) => sky.get(API.LIST_PAGE, data);
// 列表查询
export const listAll = (data) => sky.get(API.LIST_ALL, data);
// 根据ID进行查询
export const getById = (id) => sky.get(API.GET_BY_ID + id);
//# sourceMappingURL=index.js.map