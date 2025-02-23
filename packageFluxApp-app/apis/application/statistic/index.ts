// 导入二次封装axios
import sky from "@/utils/axios";

// 统一管理接口
enum API {
  FREQUENCY = "/application/app/statistic/frequency",
}

// 获取统计次数
export const frequency = (data: any) => sky.get(API.FREQUENCY, data);

