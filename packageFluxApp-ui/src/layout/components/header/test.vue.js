/// <reference types=".vue-global-types/vue_3.5_false.d.ts" />
import sky from "@/utils/axios";
import { ElLoading, ElNotification } from "element-plus";
import { skyMsgSuccess } from "@/utils/sky";
const { defineProps, defineSlots, defineEmits, defineExpose, defineModel, defineOptions, withDefaults, } = await import('vue');
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
const excelParams = (params) => {
    skyParams.value = { ...skyParams.value, ...params };
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
const handleHttpUpload = async (param) => {
    if (skyParams.value?.isApi) {
        // 方法一：统一调用上传接口
        file.value = param.file;
    }
    else {
        // 方法二：返回上传方法，前端处理excel表
    }
};
/**
 * @description 文件上传之前判断
 * @param file 上传的文件
 * */
const beforeExcelUpload = (file) => {
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
        const res = await sky.upload(skyParams.value.importApi, fileFormData);
        if (res?.code != 200) {
            ElNotification({
                title: "温馨提示",
                message: "上传失败，请重试！",
                type: "error"
            });
        }
        else {
            skyMsgSuccess("文件上传成功");
            // visible.value = true;
            // importMsg.value = res.msg;
        }
        loading.close();
        emits("handleConfirmUpload");
    }
    else {
        // 不走上传接口抛出上传数据
        if (excelList.value.length > 0) {
            // 不走上传接口抛出上传数据
            emits("handleConfirmUpload", { tableCustom: tableCustom.value, tableList: excelList.value });
        }
        else {
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
// 接收父组件传递过来的参数
let __VLS_typeProps;
const props = withDefaults(defineProps(), {
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
const skyFileList = ref([]);
// 父组件传递回显数据
skyFileList.value = props.fileList;
// 修改进行回显的时候使用
watch(() => [props.fileList], () => {
    // 父组件传递回显数据
    skyFileList.value = props.fileList;
});
/**
 * 文件变化handleChange 这里监听上传文件的变化上传一个，执行一下后端上传单个文件请求方法。
 */
const handleChange = async (file) => {
    console.log(file);
};
// 移除文件
const removeFile = (url, index) => {
    skyFileList.value.splice(index, 1);
    emits("update:skyFileList", skyFileList.value);
    emits("fileRemove", url);
};
const __VLS_withDefaultsArg = (function (t) { return t; })({
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
const __VLS_fnComponent = (await import('vue')).defineComponent({
    emits: {},
});
;
let __VLS_functionalComponentProps;
function __VLS_template() {
    const __VLS_ctx = {};
    const __VLS_localComponents = {
        ...{},
        ...{},
        ...__VLS_ctx,
    };
    let __VLS_components;
    const __VLS_localDirectives = {
        ...{},
        ...__VLS_ctx,
    };
    let __VLS_directives;
    let __VLS_styleScopedClasses;
    // CSS variable injection 
    // CSS variable injection end 
    let __VLS_resolvedLocalAndGlobalComponents;
    const __VLS_0 = __VLS_resolvedLocalAndGlobalComponents.ElDialog;
    /** @type { [typeof __VLS_components.ElDialog, typeof __VLS_components.elDialog, typeof __VLS_components.ElDialog, typeof __VLS_components.elDialog, ] } */
    // @ts-ignore
    const __VLS_1 = __VLS_asFunctionalComponent(__VLS_0, new __VLS_0({ modelValue: ((__VLS_ctx.dialogVisible)), title: ((__VLS_ctx.skyParams.title)), destroyOnClose: ((true)), width: ("580px"), draggable: (true), }));
    const __VLS_2 = __VLS_1({ modelValue: ((__VLS_ctx.dialogVisible)), title: ((__VLS_ctx.skyParams.title)), destroyOnClose: ((true)), width: ("580px"), draggable: (true), }, ...__VLS_functionalComponentArgsRest(__VLS_1));
    __VLS_elementAsFunction(__VLS_intrinsicElements.div, __VLS_intrinsicElements.div)({});
    __VLS_nonNullable(__VLS_5.slots).default;
    const __VLS_5 = __VLS_pickFunctionalComponentCtx(__VLS_0, __VLS_2);
    var __VLS_slots;
    var __VLS_inheritedAttrs;
    const __VLS_refs = {};
    var $refs;
    return {
        slots: __VLS_slots,
        refs: $refs,
        attrs: {},
    };
}
;
const __VLS_self = (await import('vue')).defineComponent({
    setup() {
        return {
            dialogVisible: dialogVisible,
            skyParams: skyParams,
        };
    },
    emits: {},
    __typeProps: {},
    props: {},
});
export default (await import('vue')).defineComponent({
    setup() {
        return {};
    },
    emits: {},
    __typeProps: {},
    props: {},
});
;
//# sourceMappingURL=test.vue.js.map