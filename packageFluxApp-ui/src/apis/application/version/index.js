// 导入二次封装axios
import sky from "@/utils/axios";
// 统一管理接口
var API;
(function (API) {
    API["ADD"] = "/application/app/version/add";
    API["DELETE"] = "/application/app/version/del/";
    API["UPDATE"] = "/application/app/version/update/";
    API["SET_FORMAL"] = "/application/app/version/appFormal";
    API["LIST_ALL"] = "/application/app/version/listAll";
    API["GET_BY_ID"] = "/application/app/version/getById/";
})(API || (API = {}));
// 添加
export const add = (data) => sky.post(API.ADD, data);
// 删除
export const deleteById = (id) => sky.delete(API.DELETE + id);
// 更新
export const updateById = (id, data) => sky.put(API.UPDATE + id, data);
// 设置正式版
export const setFormal = (data) => sky.put(API.SET_FORMAL, data);
// 列表查询
export const listAll = (data) => sky.get(API.LIST_ALL, data);
// 根据ID进行查询
export const getById = (id) => sky.get(API.GET_BY_ID + id);
//# sourceMappingURL=index.js.map