// 导入二次封装axios
import sky from "@/utils/axios.js";
// 统一管理接口
var API;
(function(API) {
	API["ADD"] = "/v1/application/details/add";
	API["DELETE"] = "/v1/application/details/del/";
	API["UPDATE"] = "/v1/application/details/update/";
	API["LIST_ALL"] = "/v1/application/details/listAll";
	API["GET_BY_ID"] = "/v1/application/details/getById/";
})(API || (API = {}));

// 添加
export const add = (data) => sky({
	url: API.ADD,
	method: "post",
	data
});

// 删除
export const deleteById = (id) => sky({
	url: API.DELETE + id,
	method: "delete",
	data
});

// 更新
export const updateById = (id, data) => sky({
	url: API.DELETE + id,
	method: "pus",
	data
});

// 列表查询
export const listAll = (data) => sky({
	url: API.LIST_ALL,
	method: "get",
	data
});

// 根据ID进行查询
export const getById = (id) => sky({
	url: API.API.GET_BY_ID + id,
	method: "get",
	data
});
//# sourceMappingURL=index.js.map