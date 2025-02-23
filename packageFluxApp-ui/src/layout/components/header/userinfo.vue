<script setup lang="ts">
import {useRouter} from "vue-router";
import {LOGIN_URL} from "@/config";
import {useUserStore} from "@/store";
import {getAssets} from "@/utils";
import {skyMsgError, skyMsgSuccess} from "@/utils/sky";
import {skySessionStorage} from "@/utils/storage";
import {logout, userinfo} from "@/apis/auth/login";
import {ArrowDown} from "@element-plus/icons-vue";

const router = useRouter();
const avatar = getAssets("images/auth/avatar.webp");
const errorAvatar = ref("");
const userStore = useUserStore();
const pageData = ref({
  userInfo: {
    email: "",
    avatar: "",
  },
});
const method = reactive({
  handleGetUserInfo: async () => {
    try {
      const res: any = await userinfo();
      userStore.setUser(res.data);
      pageData.value.userInfo = res.data;
    } catch (error) {
      skyMsgError("数据查询失败，请刷新重试🌻");
    }
  },
  handleLogout: async () => {
    try {
      const res: any = await logout();
      // 清除缓存
      skySessionStorage.clear();
      // 清除用户信息
      userStore.setToken("");
      userStore.setUser({});
      skyMsgSuccess("注销成功");
      // 退出登录。必须使用replace把页面缓存刷掉。
      // window.location.replace(LOGIN_URL);
      router.go(0);
    } catch (error) {
      skyMsgError("操作失败，请刷新重试🌻");
    }
  }
});
// 下拉折叠
const handleCommand = (command: any) => {
  switch (command) {
    case "skyMine":
      router.push("/system/personage");
      break;
    case "logout":
      method.handleLogout();
      break;
  }
};
onMounted(() => {
  method.handleGetUserInfo();
});
</script>

<template>
  <!-- 头像 -->
  <el-image class="w-34px h-34px rounded-full select-none user-avatar" :src="pageData.userInfo?.avatar || avatar">
    <template #error>
      <el-image class="w-34px h-34px rounded-full select-none user-avatar" :src="errorAvatar"></el-image>
    </template>
  </el-image>
  <el-dropdown class="m-l-10px" :hide-on-click="false" @command="handleCommand">
    <div class="sky-dropdown">
      <div class="max-w-113px text-14px m-r-6px line-clamp-1">{{ pageData.userInfo.username }}</div>
      <el-icon>
        <arrow-down/>
      </el-icon>
    </div>
    <template #dropdown>
      <el-dropdown-menu>
        <el-dropdown-item command="logout">退出登录</el-dropdown-item>
      </el-dropdown-menu>
    </template>
  </el-dropdown>
</template>

<style scoped lang="scss">
.sky-dropdown {
  cursor: pointer;
  color: var(--el-color-primary);
  font-size: 17px;
  white-space: nowrap; /* 不换行 */
  outline: none; // 去除伪元素
  display: flex;
  align-items: center;
}
</style>
