<script setup lang="ts">
import sky from "@/utils/axios";
import {skyMsgSuccess} from "@/utils/sky";
import {ElLoading, ElNotification} from "element-plus";
import {Close, Download, UploadFilled} from "@element-plus/icons-vue";

defineOptions({name: "SkyUploadFile"});
// 当前组件获取父组件传递的事件方法，然后点击确认和提交是触发父组件传递过来的事件
const emits = defineEmits(["handleTemplateExcel", "handleConfirmUpload"]);
// 最大文件上传数
const fileLimit = ref(10);
// dialog状态
const dialogVisible = ref(false);
// 父组件传过来的参数
const skyParams = ref({
  title: "上传文件🌻",
  isApi: false,
  fileSize: 5,
  importApi: ""
});
// 上传的文件
const file = ref([]);
// 接收父组件参数
const uploadParams = (params: any) => {
  skyParams.value = {...skyParams.value, ...params};
  dialogVisible.value = true;
};
// 文件上传
const handleHttpUpload = async (param: any) => {
  if (skyParams.value?.isApi) {
    // 方法一：统一调用上传接口
    file.value.push(param.file);
  } else {
    // 方法二：返回上传方法，前端处理上传
  }
};
// 文件数超出提示
const handleExceed = () => {
  ElNotification({
    title: "温馨提示",
    message: "最多只能上传十个文件！",
    type: "warning"
  });
};
// 上传错误提示
const handleUploadError = () => {
  ElNotification({
    title: "温馨提示",
    message: `${skyParams.value.title}上传失败，请您重新上传！`,
    type: "error"
  });
};
// 上传
const handleConfirm = async () => {
  if (skyParams.value.isApi) {
    const loading = ElLoading.service({
      lock: true,
      text: '文件上传中，请稍等...',
      background: 'rgba(0, 0, 0, 0.7)',
    });
    // 统一调用上传接口
    let fileFormData = new FormData();
    for (const f of file.value) {
      fileFormData.append("file", f);
    }
    const res: any = await sky.upload(skyParams.value.importApi, fileFormData);
    if (res?.code != 200) {
      ElNotification({
        title: "温馨提示",
        message: "上传失败，请重试！",
        type: "error"
      });
    } else {
      skyMsgSuccess("文件上传成功");
      file.value = [];
    }
    loading.close();
    emits("handleConfirmUpload");
  } else {
    // 不走上传接口抛出上传数据
  }
  dialogVisible.value = false;
};
// 取消
const handleCancel = () => {
  file.value = [];
  dialogVisible.value = false;
};
// 移除文件
const handleBeforeRemove = (f: any) => {
  const index = file.value.findIndex(item => item.name === f.name);
  file.value.splice(index, 1);
};
defineExpose({
  uploadParams
});
</script>

<template>
  <el-dialog v-model="dialogVisible" :title="skyParams.title" :destroy-on-close="true" width="580px" draggable>
    <el-form label-width="100px">
      <el-form-item label="温馨提示：">
        <el-tag type="warning">任何类型文件都可以上传，但是请加上后缀，支持批量上传</el-tag>
      </el-form-item>
      <el-form-item label="文件上传：">
        <el-upload
          action="#"
          class="upload"
          :drag="true"
          :limit="fileLimit"
          :multiple="true"
          :show-file-list="true"
          :http-request="handleHttpUpload"
          :on-exceed="handleExceed"
          :on-error="handleUploadError"
          :before-remove="handleBeforeRemove"
        >
          <slot name="content">
            <el-icon class="el-icon--upload">
              <upload-filled/>
            </el-icon>
            <div class="el-upload__text">将文件拖到此处，或<em>点击上传</em></div>
          </slot>
          <template #tip>
            <slot name="tip">
              <div class="el-upload__tip">请上传标准格式文件，文件最大为 {{ skyParams.fileSize }}M</div>
            </slot>
          </template>
        </el-upload>
      </el-form-item>
    </el-form>
    <template #footer>
      <el-button type="primary" @click="handleConfirm">上传</el-button>
      <el-button type="danger" @click="handleCancel">取消</el-button>
    </template>
  </el-dialog>
</template>

<style scoped lang="scss">
:deep(.el-upload-dragger) {
  border: 2px dashed var(--el-color-primary);

  &:hover {
    border: 2px dashed var(--el-color-primary);
  }
}

.upload {
  width: 80%;
}
</style>
