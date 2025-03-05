// 导入二次封装axios
import sky from "@/utils/axios";
// 统一管理接口
var API;
(function (API) {
    API["LIST_PAGE"] = "/application/app/appLog/listPage";
})(API || (API = {}));
// 分页列表
export const listPage = (data) => sky.get(API.LIST_PAGE, data);
//# sourceMappingURL=index.js.map