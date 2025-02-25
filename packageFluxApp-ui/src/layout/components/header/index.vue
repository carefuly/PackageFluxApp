<script setup lang="ts">
import Upload from "./upload.vue";
import Bucket from "./bucket.vue";
import UserInfo from "./userinfo.vue";

const title = import.meta.env.VITE_WEB_TITLE;
</script>

<template>
  <header>
    <el-tooltip :content="title">
      <div class="application-logo">
        <img src="/favicon.png" alt="favicon">
        <span>{{ title }}</span>
      </div>
    </el-tooltip>
    <div class="application-list">
      <!-- 上传文件 -->
      <Upload></Upload>
      <!-- 存储桶 -->
      <Bucket></Bucket>
      <!-- 用户信息 -->
      <UserInfo></UserInfo>
    </div>
  </header>
</template>

<style scoped lang="scss">
header {
  margin: 0 auto;
  width: 1850px;
  height: 100%;
  display: flex;
  align-items: center;
  justify-content: space-between;

  .application-logo {
    height: 100%;
    display: flex;
    align-items: center;

    img {
      width: 36px;
      height: 36px;
      margin-right: 10px;
    }

    span {
      font-size: 20px;
      font-weight: 500;
    }
  }

  .application-list {
    height: 100%;
    display: flex;
    align-items: center;
    justify-content: center;
  }
}
</style>
