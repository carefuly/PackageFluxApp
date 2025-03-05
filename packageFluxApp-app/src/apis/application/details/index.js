// 导入二次封装axios
import sky from "@/utils/axios";
// 统一管理接口
var API;
(function (API) {
    API["ADD"] = "/v1/application/detail/create";
    API["DELETE"] = "/v1/application/detail/delete/";
    API["UPDATE"] = "/v1/application/detail/update/";
    API["GET_BY_ID"] = "/v1/application/detail/getById/";
    API["LIST_ALL"] = "/v1/application/detail/listAll";
})(API || (API = {}));
// 添加
export const add = (data) => sky.post(API.ADD, data);
// 删除
export const deleteById = (id) => sky.delete(API.DELETE + id);
// 更新
export const updateById = (id, data) => sky.put(API.UPDATE + id, data);
// 根据ID进行查询
export const getById = (id) => sky.get(API.GET_BY_ID + id);
// 列表查询
export const listAll = (data) => sky.get(API.LIST_ALL, data);
//# sourceMappingURL=index.js.map