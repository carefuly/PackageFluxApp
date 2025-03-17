// 导入二次封装axios
import sky from "@/utils/axios";

// 统一管理接口
enum API {
  STATISTICS = "/v1/application/logger/statistics",
}

// 获取统计次数
export const statistics = (data: any) => sky.get(API.STATISTICS, data);

