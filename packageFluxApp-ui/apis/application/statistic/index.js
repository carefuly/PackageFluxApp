// 导入二次封装axios
import sky from "@/utils/axios";
// 统一管理接口
var API;
(function (API) {
    API["FREQUENCY"] = "/application/app/statistic/frequency";
})(API || (API = {}));
// 获取统计次数
export const frequency = (data) => sky.get(API.FREQUENCY, data);
//# sourceMappingURL=index.js.map