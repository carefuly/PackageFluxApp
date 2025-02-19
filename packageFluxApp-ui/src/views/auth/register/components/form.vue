<script setup lang="ts">
import settings from "@/settings";
import {useRouter} from "vue-router";
import {LOGIN_URL} from "@/config";
import {getAssets} from "@/utils";
import {skyMsgError, skyMsgSuccess} from "@/utils/sky";
import {captcha} from "@/apis/auth/captcha";
import {register} from "@/apis/auth/register";
import {User, Lock, Message, TurnOff} from "@element-plus/icons-vue";

const router = useRouter();
const logo = getAssets("images/logo/logo.png");
const formRef = ref();
const pageData = ref({
  isCountingDown: false,
  countdownSeconds: 60,
  loading: false,
  form: {
    username: "",
    email: "",
    password: "",
    code: "",
  },
  rules: {
    username: [{required: true, message: "请输入用户账号", trigger: "blur"}],
    email: [{required: true, message: "请输入登录的邮箱", trigger: "blur"}],
    password: [{required: true, message: "请输入密码", trigger: "blur"}],
    code: [{required: true, message: "请输入验证码", trigger: "blur"}],
  },
});
const method = reactive({
  // 发送验证码
  handleSendCaptcha: async () => {
    if (pageData.value.form.email.length <= 0) return skyMsgError("请输入邮箱🌻");
    if (pageData.value.isCountingDown) return; // 如果正在倒计时，则不执行发送操作
    const res: any = await captcha({email: pageData.value.form.email});
    try {
      // 发送成功
      skyMsgSuccess(res.msg + "🌻");
      // 开始倒计时
      pageData.value.isCountingDown = true;
      const intervalId = setInterval(() => {
        pageData.value.countdownSeconds--;
        if (pageData.value.countdownSeconds <= 0) {
          clearInterval(intervalId); // 清除定时器
          pageData.value.isCountingDown = false;
          pageData.value.countdownSeconds = 60; // 重置倒计时秒数
        }
      }, 1000);
    } catch (error) {
      skyMsgError("发送失败，请稍后重试🌻");
    }
  },
  // 注册
  skyRegister: async () => {
    if (!formRef.value) return;
    formRef.value.validate(async (valid: any) => {
      if (valid) {
        pageData.value.loading = true;
        try {
          await register(pageData.value.form);
          // 跳转到登录
          pageData.value.loading = false;
          skyMsgSuccess("注册成功🌻");
          await router.push(LOGIN_URL);
        } catch (error) {
          // 等待1秒关闭loading
          let loadingTime = 1;
          setInterval(() => {
            loadingTime--;
            if (loadingTime === 0) {
              pageData.value.loading = false;
            }
          }, 1000);
        }
      } else {
        skyMsgError("校验失败，信息填写有误🌻");
      }
    });
  },
});
const countdownText = computed(() => {
  return pageData.value.isCountingDown ? `${pageData.value.countdownSeconds}秒后重试` : '发送验证码';
});
</script>

<template>
  <div class="form center">
    <div class="flex items-center formRef">
      <el-image class="rounded-full w-36px h-36px" :src="logo"/>
      <div class="ml-6px font-bold text-xl">{{ settings.authTitle || "APP 热更新 管理平台" }}</div>
    </div>
    <div class="flex items-center space-x-3 text-gray-400 mt-16px mb-16px formRef">
      <span class="h-1px w-16 bg-gray-300 inline-block"></span>
      <span class="text-center">邮箱密码注册</span>
      <span class="h-1px w-16 bg-gray-300 inline-block"></span>
    </div>
    <!-- 输入框盒子 -->
    <el-form ref="formRef" class="w-300px formRef" :model="pageData.form" :rules="pageData.rules" label-width="auto">
      <el-row :gutter="24">
        <el-col :span="24">
          <el-form-item label="用户名" prop="username" required>
            <el-input type="text" v-model="pageData.form.username" placeholder="请输入用户名" :suffix-icon="User"/>
          </el-form-item>
        </el-col>
      </el-row>
      <el-row :gutter="24">
        <el-col :span="24">
          <el-form-item label="邮箱" prop="email" required>
            <el-input type="text" v-model="pageData.form.email" placeholder="请输入邮箱" :suffix-icon="Message"/>
          </el-form-item>
        </el-col>
      </el-row>
      <el-row :gutter="24">
        <el-col :span="24">
          <el-form-item label="密码" prop="password" required>
            <el-input type="password" v-model="pageData.form.password" placeholder="请输入密码" show-password
                      :suffix-icon="Lock" @keydown.enter="method.skyRegister"/>
          </el-form-item>
        </el-col>
      </el-row>
      <el-row :gutter="24">
        <el-col :span="17">
          <el-form-item label="验证码" prop="code" required>
            <el-input type="text" v-model="pageData.form.code" placeholder="请输入验证码" :suffix-icon="TurnOff"/>
          </el-form-item>
        </el-col>
        <el-col :span="7">
          <el-form-item>
            <el-button
              style="width: 80px"
              type="primary"
              text
              :disabled="pageData.isCountingDown"
              v-throttle="method.handleSendCaptcha"
            >
              {{ countdownText }}
            </el-button>
          </el-form-item>
        </el-col>
      </el-row>
      <el-row :gutter="24">
        <el-col :span="15">
          <el-form-item label="记住密码">
            <el-checkbox name="type"/>
          </el-form-item>
        </el-col>
        <el-col :span="8">
          <el-form-item>
            <span style="color: #1d189d; font-weight: bold;" @click="router.push('/login')">前往登录</span>
          </el-form-item>
        </el-col>
      </el-row>
      <el-row :gutter="24">
        <el-col :span="24">
          <!-- 注册按钮 -->
          <el-form-item>
            <el-button class="w-245px bg-[--el-color-primary]"
                       type="primary" v-if="!pageData.loading"
                       round @click="method.skyRegister">
              注册
            </el-button>
            <el-button class="w-245px bg-[--el-color-primary]" type="primary" v-if="pageData.loading" round
                       :loading="pageData.loading">
              拼命注册中...
            </el-button>
          </el-form-item>
        </el-col>
      </el-row>
    </el-form>
  </div>
</template>

<style scoped lang="scss">
.form {
  height: 100%;
  display: flex;
  flex-direction: column;
}

.center {
  display: flex;
  align-items: center;
  justify-content: center;
}

.formRef {
  animation: FadeInLeft 1s;
}
</style>
