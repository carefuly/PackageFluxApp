// 导入二次封装axios
import sky from "@/utils/axios";
// 统一管理接口
var API;
(function (API) {
    API["ADD"] = "/v1/application/version/create";
    API["DELETE"] = "/v1/application/version/delete/";
    API["UPDATE"] = "/v1/application/version/update/";
    API["SET_FORMAL"] = "/v1/application/version/updateFormal/";
    API["LIST_ALL"] = "/v1/application/version/listAll";
    API["GET_BY_ID"] = "/v1/application/version/getById/";
})(API || (API = {}));
// 添加
export const add = (data) => sky.post(API.ADD, data);
// 删除
export const deleteById = (id, detailId) => sky.delete(API.DELETE + id + "?detailId=" + detailId);
// 更新
export const updateById = (id, data) => sky.put(API.UPDATE + id, data);
// 设置正式版
export const setFormal = (id, detailId) => sky.put(API.SET_FORMAL + id + "?detailId=" + detailId);
// 列表查询
export const listAll = (data) => sky.get(API.LIST_ALL, data);
// 根据ID进行查询
export const getById = (id, data) => sky.get(API.GET_BY_ID + id, data);
//# sourceMappingURL=index.js.map