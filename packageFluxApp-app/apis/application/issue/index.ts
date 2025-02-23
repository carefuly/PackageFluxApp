// 导入二次封装axios
import sky from "@/utils/axios";

// 统一管理接口
enum API {
  GET_ISSUE = "/application/app/issue/",
}

// 获取应用发布信息
export const getIssue = (id: any) => sky.get(API.GET_ISSUE + id);

