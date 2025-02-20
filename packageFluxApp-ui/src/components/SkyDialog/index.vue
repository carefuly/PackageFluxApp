<script setup lang="ts">
import {skyMsgWarning, skyMsgBox} from "@/utils/sky";

const props = defineProps({
  title: {default: "朕很中意你SkyDialog"},
  width: {default: 650},
  height: {default: 300},
  visible: {default: false},
  confirmText: {default: "确定"},
  cancelText: {default: "取消"},
  destroyOnClose: {default: false},
  fullscreen: {default: false},
  loading: {default: false},
  footerHidden: {default: false},
});
// 开关变量
const visible = ref(false);
// 确定的loading，此处必须用toRefs，否则将失去响应式
const {loading} = toRefs(props);
const confirmLoading = ref(loading);
// 打开弹窗
const skyOpen = () => {
  visible.value = true;
};
// 取消
const skyClose = () => {
  skyMsgBox("您确认进行关闭么？")
    .then(() => {
      visible.value = false;
      skyMsgWarning("已关闭🌻");
    })
    .catch(() => {
      // 用户点击了取消按钮或关闭确认框
      // 执行取消操作或不做任何操作
      skyMsgWarning("已取消🌻");
    });
};
// 确认提交后关闭弹窗
const skyQuickClose = () => {
  visible.value = false;
  skyMsgWarning("已提交🌻");
};

// 当前组件获取父组件传递的事件方法
const emits = defineEmits(["skyConfirm", "skyCancel"]);
// 弹框的确定事件
const skyConfirm = () => {
  emits("skyConfirm");
};
// 弹框的取消事件
const skyCancel = () => {
  emits("skyCancel");
};

// 暴露给父组件方法
defineExpose({
  skyOpen,
  skyClose,
  skyQuickClose
});
</script>

<template>
  <el-dialog
    :model-value="visible"
    :title="props.title"
    :width="props.width"
    :close-on-click-modal="false"
    append-to-body
    draggable
    :destroy-on-close="props.destroyOnClose"
    :before-close="skyClose"
    :fullscreen="props.fullscreen"
    :loading="props.loading"
    :footerHidden="props.footerHidden"
  >
    <slot name="header"></slot>
    <div class="container" :style="{ height: props.height + 'px' }">
      <!-- 具名插槽 -->
      <slot name="content"></slot>
    </div>
    <template #footer v-if="!props.footerHidden">
      <span class="dialog-footer">
        <el-button type="primary" loading-icon="Eleme" :loading="confirmLoading" v-throttle="skyConfirm">
          {{ props.confirmText }}</el-button>
        <el-button type="danger" @click="skyCancel">{{ props.cancelText }}</el-button>
      </span>
    </template>
  </el-dialog>
</template>

<style scoped lang="scss">
.container {
  overflow-x: initial;
  overflow-y: auto; // 超出部分则滚动
}

.el-dialog {
  border-top-left-radius: 8px !important;
  border-top-right-radius: 8px !important;
  padding-top: 0;

  // 标题头部
  .el-dialog__header {
    display: flex;
    align-items: center;
    justify-content: center;
    height: 50px !important;
    padding: 0 0 6px 0;

    // background: #1e71ee;
    @apply dark:bg-#141414;
    .el-dialog__title {
      font-family: YouYuan;
      font-size: 18px;
      font-weight: 500;
    }
  }

  .el-dialog__body {
    padding: 0;
  }

  .el-dialog__footer {
    padding: 10px;
  }
}

.el-dialog__body {
  // 内容区域内边距
  padding: 10px;
}

.el-dialog__headerbtn {
  padding-bottom: 10px !important;

  .el-dialog__close {
    border: 1px solid;
    border-radius: 10px;
  }
}
</style>
