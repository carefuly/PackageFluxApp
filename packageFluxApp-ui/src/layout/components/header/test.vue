<script setup lang="ts">
import sky from "@/utils/axios";
import {ElLoading, ElNotification} from "element-plus";
import {Close, Download, UploadFilled} from "@element-plus/icons-vue";
import {skyMsgSuccess} from "@/utils/sky";

// defineOptions({name: "SkyUploadFile"});
// 当前组件获取父组件传递的事件方法，然后点击确认和提交是触发父组件传递过来的事件
// const emits = defineEmits(["handleTemplateExcel", "handleConfirmUpload"]);

// 最大文件上传数
const excelLimit = ref(1);
// dialog状态
const dialogVisible = ref(false);
//限制.xls文件
const xlsFile = "application/vnd.ms-excel";
//限制.xlsx文件
const xlsxFile = "application/vnd.openxmlformats-officedocument.spreadsheetml.sheet";

// 父组件传过来的参数
const skyParams = ref({
  title: "上传文件",
  isApi: false,
  fileSize: 5,
  templeApi: "",
  importApi: ""
});
// 接收父组件参数
const excelParams = (params: any) => {
  skyParams.value = {...skyParams.value, ...params};
  dialogVisible.value = true;
};
// 下载文件
const handleTemplateExcel = () => {
  emits("handleTemplateExcel");
};
// 上传的excel表数据
const tableCustom = ref([]);
const excelList = ref([]);
const file = ref();
// 文件上传
const handleHttpUpload = async (param: any) => {
  if (skyParams.value?.isApi) {
    // 方法一：统一调用上传接口
    file.value = param.file;
  } else {
    // 方法二：返回上传方法，前端处理excel表
  }
};
/**
 * @description 文件上传之前判断
 * @param file 上传的文件
 * */
const beforeExcelUpload = (file: any) => {
  const fileSize = file.size / 1024 / 1024 < skyParams.value.fileSize;
  if (!fileSize)
    setTimeout(() => {
      ElNotification({
        title: "温馨提示",
        message: `上传文件大小不能超过 ${skyParams.value.fileSize}MB！`,
        type: "warning"
      });
    }, 0);
  return fileSize;
};
// 文件数超出提示
const handleExceed = () => {
  ElNotification({
    title: "温馨提示",
    message: "最多只能上传一个文件！",
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
// 上传成功提示
const handleUploadSuccess = () => {
  ElNotification({
    title: "温馨提示",
    message: `${skyParams.value.title}文件上传成功！`,
    type: "success"
  });
};
// 导入结果
const visible = ref(false);
// 导入信息
const importMsg = ref();
// 关闭信息
const closeDrawer = () => {
  visible.value = false;
  importMsg.value = null;
};
// 上传
const handleConfirm = async () => {
  if (skyParams.value.isApi) {
    const loading = ElLoading.service({
      lock: true,
      text: '数据导入中...',
      background: 'rgba(0, 0, 0, 0.7)',
    });
    // 统一调用上传接口
    let fileFormData = new FormData();
    fileFormData.append("file", file.value);
    const res: any = await sky.upload(skyParams.value.importApi, fileFormData);
    if (res?.code != 200) {
      ElNotification({
        title: "温馨提示",
        message: "上传失败，请重试！",
        type: "error"
      });
    } else {
      skyMsgSuccess("文件上传成功");
      // visible.value = true;
      // importMsg.value = res.msg;
    }
    loading.close();
    emits("handleConfirmUpload");
  } else {
    // 不走上传接口抛出上传数据
    if (excelList.value.length > 0) {
      // 不走上传接口抛出上传数据
      emits("handleConfirmUpload", {tableCustom: tableCustom.value, tableList: excelList.value});
    } else {
      ElNotification({
        title: "温馨提示",
        message: `当前上传数据为空，自动取消上传！`,
        type: "error"
      });
    }
  }
  dialogVisible.value = false;
};
// 取消
const handleCancel = () => {
  tableCustom.value = [];
  excelList.value = [];
  dialogVisible.value = false;
};
// defineExpose({
//   excelParams
// });

// 当前组件获取父组件传递的事件方法，然后点击确认和提交是触发父组件传递过来的事件
const emits = defineEmits(["fileSuccess", "fileRemove", "update:skyFileList"]);

interface Props {
  acceptType?: string; // 上传文件类型
  acceptTypes?: string; // 描述 - 上传文件类型
  isMultiple?: boolean; //   是否可批量上传
  limit?: number; // 允许上传文件的最大数量
  fileDisabled?: boolean; // 是否禁用上传
  fileSize?: number; // 文件大小
  action?: string;
  fileList?: any; // 回显的文件
  isDownLoad?: boolean; // 是否可以下载
  folderName?: string; // 后端文件夹名称
}

// 接收父组件传递过来的参数
const props = withDefaults(defineProps<Props>(), {
  acceptType: ".xls,.xlsx,.pdf,.doc,.docx,.txt,.jar,.zip",
  acceptTypes: "xls/xlsx/pdf/doc/docx/txt/jar/zip",
  isMultiple: true,
  limit: 5,
  fileDisabled: false,
  fileSize: 10,
  action: "/application/app/file/batch/add",
  fileList: [],
  isDownLoad: false,
  folderName: "files"
});
// 上传文件数组
const skyFileList = ref<any>([]);
// 父组件传递回显数据
skyFileList.value = props.fileList;
// 修改进行回显的时候使用
watch(
  () => [props.fileList],
  () => {
    // 父组件传递回显数据
    skyFileList.value = props.fileList;
  }
);
/**
 * 文件变化handleChange 这里监听上传文件的变化上传一个，执行一下后端上传单个文件请求方法。
 */
const handleChange = async (file: any) => {
  console.log(file)
};
// 移除文件
const removeFile = (url: string, index: number) => {
  skyFileList.value.splice(index, 1);
  emits("update:skyFileList", skyFileList.value);
  emits("fileRemove", url);
};
</script>

<template>
  <!-- 上传文件 -->
  <el-dialog v-model="dialogVisible" :title="skyParams.title" :destroy-on-close="true" width="580px" draggable>
    <div>
      <!--      <el-upload-->
      <!--        :file-list="skyFileList"-->
      <!--        :multiple="props.isMultiple"-->
      <!--        :limit="props.limit"-->
      <!--        :accept="props.acceptType"-->
      <!--        :auto-upload="false"-->
      <!--        :show-file-list="false"-->
      <!--        :disabled="props.fileDisabled"-->
      <!--        :on-change="handleChange"-->
      <!--        :folderName="folderName"-->
      <!--      >-->
      <!--        <div class="el-upload-text">-->
      <!--          <el-icon size="16">-->
      <!--            <Upload/>-->
      <!--          </el-icon>-->
      <!--          <span>上传文件</span>-->
      <!--        </div>-->
      <!--      </el-upload>-->
      <!--      <div class="template-list">-->
      <!--        <div class="template" v-for="(item, index) in skyFileList" :key="index">-->
      <!--          <el-icon size="16">-->
      <!--            <Link/>-->
      <!--          </el-icon>-->
      <!--          <span class="document-name">{{ item.name }}</span>-->
      <!--          <span v-if="!props.fileDisabled">-->
      <!--          <el-icon size="16" @click="removeFile(item.url, index)">-->
      <!--            <Close/>-->
      <!--          </el-icon>-->
      <!--        </span>-->
      <!--          &lt;!&ndash; 默认不显示下载 &ndash;&gt;-->
      <!--          <span v-if="isDownLoad" style="padding-left: 5px">-->
      <!--          <el-icon size="16" @click="handleDownLoad(item.url)"><Download/></el-icon>-->
      <!--        </span>-->
      <!--        </div>-->
      <!--      </div>-->
      <!--      <span class="file-tips">-->
      <!--      <slot name="tip"> 支持{{ acceptTypes }}；文件大小不能超过{{ props.fileSize }}M；最多上传{{ props.limit }}个； </slot>-->
      <!--    </span>-->
    </div>
  </el-dialog>
</template>

<style scoped lang="scss">
.el-upload-text {
  display: flex;
  align-items: center;
  justify-content: center;
  width: 106px;
  height: 32px;
  color: var(--el-color-primary);

  /* 设置用户禁止选中 */
  user-select: none;
  border: 2px dashed var(--el-color-primary);
  border-radius: 6px;

  img {
    display: block;
    width: 14px;
    height: 14px;
  }

  span {
    padding-left: 6px;
    font-family: "PingFangSC-Regular, PingFang SC";
    font-size: 14px;
    font-weight: 500;
  }
}

.template-list {
  padding-bottom: 4px;
}

.template {
  display: flex;
  align-items: center;
  padding: 5px 0;

  span {
    line-height: 10px;
  }

  img {
    width: 16px;
    height: 16px;
    margin-right: 8px;
  }

  .document-name {
    margin-right: 12px;
    font-size: 14px;
  }
}

.file-tips {
  font-family: "PingFangSC-Regular, PingFang SC";
  font-size: 12px;
  color: var(--el-color-primary);
}
</style>
