// 导入二次封装axios
import sky from "@/utils/axios";
// 统一管理接口
var API;
(function (API) {
    API["GET_ISSUE"] = "/application/app/issue/";
})(API || (API = {}));
// 获取应用发布信息
export const getIssue = (id) => sky.get(API.GET_ISSUE + id);
//# sourceMappingURL=index.js.map