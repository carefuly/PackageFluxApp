<script setup>
import Describe from "./components/describe.vue";
import Form from "./components/form.vue";
/** 适配移动端开始 */
import {useScreenStore} from "@/hooks/screen";
// 获取当前为[移动端、IPad、PC端]仓库，阔以使用 watchEffect(() => {}) 进行监听。
const {isScreen} = useScreenStore();
/** 适配移动端结束 */
</script>

<template>
  <el-row class="min-h-screen">
    <el-col v-if="!isScreen" :span="12">
      <Describe></Describe>
    </el-col>
    <el-col :span="isScreen ? 24 : 12">
      <Form></Form>
    </el-col>
  </el-row>
</template>

<style scoped>

</style>
