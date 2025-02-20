<script setup lang="ts">
import SkyUploadFile from "@/components/SkyUpload/file.vue";

const skyFileRef = ref();
const method = reactive({
  /** 上传文件 */
  handleImportFile: () => {
    let params = {
      title: "上传文件🌻",
      isApi: true, // 是否后台上传
      importApi: "/application/app/file/batch/add"
    };
    skyFileRef.value.uploadParams(params);
  },
  /** 文件上传回调 */
  handleConfirmUpload: () => {
  },
});
</script>

<template>
  <div class="file">
    <el-button type="primary" icon="Upload" link @click="method.handleImportFile">上传文件</el-button>
    <!-- 文件上传 -->
    <SkyUploadFile ref="skyFileRef" @handleConfirmUpload="method.handleConfirmUpload"></SkyUploadFile>
  </div>
</template>

<style scoped lang="scss">
.file {
  margin-right: 15px;
}
</style>
