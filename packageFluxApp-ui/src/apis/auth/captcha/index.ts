// 导入二次封装axios
import sky from "@/utils/axios";

// 统一管理接口
enum API {
  CAPTCHA = "/v1/auth/send-email-captcha",
}

// 发送验证码
export const captcha = (data: any) => sky.post(API.CAPTCHA, data);

