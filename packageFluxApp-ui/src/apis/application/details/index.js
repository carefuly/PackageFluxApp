// 导入二次封装axios
import sky from "@/utils/axios";
// 统一管理接口
var API;
(function (API) {
    API["ADD"] = "/application/app/details/add";
    API["DELETE"] = "/application/app/details/del/";
    API["UPDATE"] = "/application/app/details/update/";
    API["LIST_ALL"] = "/application/app/details/listAll";
    API["GET_BY_ID"] = "/application/app/details/getById/";
})(API || (API = {}));
// 添加
export const add = (data) => sky.post(API.ADD, data);
// 删除
export const deleteById = (id) => sky.delete(API.DELETE + id);
// 更新
export const updateById = (id, data) => sky.put(API.UPDATE + id, data);
// 列表查询
export const listAll = (data) => sky.get(API.LIST_ALL, data);
// 根据ID进行查询
export const getById = (id) => sky.get(API.GET_BY_ID + id);
//# sourceMappingURL=index.js.map