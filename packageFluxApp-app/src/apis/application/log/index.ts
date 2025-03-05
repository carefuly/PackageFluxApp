// 导入二次封装axios
import sky from "@/utils/axios";

// 统一管理接口
enum API {
  LIST_PAGE = "/application/app/appLog/listPage",
}

// 分页列表
export const listPage = (data: any) => sky.get(API.LIST_PAGE, data);

