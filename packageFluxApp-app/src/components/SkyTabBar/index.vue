<script setup lang="ts">

const active = ref(0);
</script>

<template>
  <van-tabbar v-model="active">
    <van-tabbar-item icon="wap-home-o">详情</van-tabbar-item>
    <van-tabbar-item icon="down">存储桶</van-tabbar-item>
    <van-tabbar-item icon="manager-o">我的</van-tabbar-item>
  </van-tabbar>
</template>

<style scoped>

</style>
