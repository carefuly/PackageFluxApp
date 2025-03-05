// 导入二次封装axios
import sky from "@/utils/axios";

// 统一管理接口
enum API {
  REGISTER = "/v1/auth/password-register",
}

// 注册
export const register = (data: any) => sky.post(API.REGISTER, data);

