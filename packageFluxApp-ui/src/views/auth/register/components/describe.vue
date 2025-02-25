<script setup>
import settings from "@/settings";
import {getAssets} from "@/utils";

const logo = getAssets("images/logo/logo.png");
const bg = getAssets("images/auth/auth-box-bg.svg");
</script>

<template>
  <div class="text">
    <div class="title">
      <img :src="logo" alt="logo">
      <span>欢迎注册 {{ settings.authTitle || "APP 热更新 管理平台" }}</span>
    </div>
    <div class="describe">
      <img :src="bg" alt="register-box">
      <p>一款开箱即用的APP热更新管理系统！</p>
      <span>输入您的个人详细信息开始注册</span>
    </div>
  </div>
</template>

<style scoped lang="scss">
.text {

  &:before {
    content: "";
    position: absolute;
    top: 0;
    left: 0;
    width: 100%;
    height: 100%;
    margin-left: -48%;
    background-image: url("@/assets/images/auth/auth-bg.svg");
    background-repeat: no-repeat;
    background-position: 100%;
    background-size: auto 100%;
    z-index: -1;
  }

  .title {
    margin: 13px 0 0 250px;
    height: 80px;
    display: flex;
    align-items: center;
    animation: FadeInRight 1s;

    img {
      width: 48px;
      border-radius: 50%;
    }

    span {
      margin-left: 10px;
      font-size: 24px;
      font-weight: bold;
      color: #fff;
    }
  }

  .describe {
    margin: 200px 0 0 250px;
    animation: FadeInRight 1s;

    img {
      width: 350px;
      height: 180px;
      animation: FadeInRight 1s;
    }

    p {
      margin-top: 50px;
      font-size: 27px;
      color: #fff;
      animation: FadeInRight 1.5s;
    }

    span {
      font-size: 17px;
      color: #fff;
      animation: FadeInRight 2s;
    }
  }
}
</style>
