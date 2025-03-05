import sky from "@/utils/axios";
import { skyMsgSuccess } from "@/utils/sky";
import { ElLoading, ElNotification } from "element-plus";
import { UploadFilled } from "@element-plus/icons-vue";
defineOptions({ name: "SkyUploadFile" });
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
const uploadParams = (params) => {
    skyParams.value = { ...skyParams.value, ...params };
    dialogVisible.value = true;
};
// 文件上传
const handleHttpUpload = async (param) => {
    if (skyParams.value?.isApi) {
        // 方法一：统一调用上传接口
        file.value.push(param.file);
    }
    else {
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
            fileFormData.append("files", f);
        }
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
            file.value = [];
        }
        loading.close();
        emits("handleConfirmUpload");
    }
    else {
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
const handleBeforeRemove = (f) => {
    const index = file.value.findIndex(item => item.name === f.name);
    file.value.splice(index, 1);
};
const __VLS_exposed = {
    uploadParams
};
defineExpose({
    uploadParams
});
; /* PartiallyEnd: #3632/scriptSetup.vue */
const __VLS_ctx = {};
let __VLS_components;
let __VLS_directives;
// CSS variable injection 
// CSS variable injection end 
const __VLS_0 = {}.ElDialog;
/** @type { [typeof __VLS_components.ElDialog, typeof __VLS_components.elDialog, typeof __VLS_components.ElDialog, typeof __VLS_components.elDialog, ] } */ ;
// @ts-ignore
const __VLS_1 = __VLS_asFunctionalComponent(__VLS_0, new __VLS_0({
    modelValue: ((__VLS_ctx.dialogVisible)),
    title: ((__VLS_ctx.skyParams.title)),
    destroyOnClose: ((true)),
    width: ("580px"),
    draggable: (true),
}));
const __VLS_2 = __VLS_1({
    modelValue: ((__VLS_ctx.dialogVisible)),
    title: ((__VLS_ctx.skyParams.title)),
    destroyOnClose: ((true)),
    width: ("580px"),
    draggable: (true),
}, ...__VLS_functionalComponentArgsRest(__VLS_1));
var __VLS_6 = {};
const __VLS_7 = {}.ElForm;
/** @type { [typeof __VLS_components.ElForm, typeof __VLS_components.elForm, typeof __VLS_components.ElForm, typeof __VLS_components.elForm, ] } */ ;
// @ts-ignore
const __VLS_8 = __VLS_asFunctionalComponent(__VLS_7, new __VLS_7({
    labelWidth: ("100px"),
}));
const __VLS_9 = __VLS_8({
    labelWidth: ("100px"),
}, ...__VLS_functionalComponentArgsRest(__VLS_8));
const __VLS_13 = {}.ElFormItem;
/** @type { [typeof __VLS_components.ElFormItem, typeof __VLS_components.elFormItem, typeof __VLS_components.ElFormItem, typeof __VLS_components.elFormItem, ] } */ ;
// @ts-ignore
const __VLS_14 = __VLS_asFunctionalComponent(__VLS_13, new __VLS_13({
    label: ("温馨提示："),
}));
const __VLS_15 = __VLS_14({
    label: ("温馨提示："),
}, ...__VLS_functionalComponentArgsRest(__VLS_14));
const __VLS_19 = {}.ElTag;
/** @type { [typeof __VLS_components.ElTag, typeof __VLS_components.elTag, typeof __VLS_components.ElTag, typeof __VLS_components.elTag, ] } */ ;
// @ts-ignore
const __VLS_20 = __VLS_asFunctionalComponent(__VLS_19, new __VLS_19({
    type: ("warning"),
}));
const __VLS_21 = __VLS_20({
    type: ("warning"),
}, ...__VLS_functionalComponentArgsRest(__VLS_20));
__VLS_24.slots.default;
var __VLS_24;
__VLS_18.slots.default;
var __VLS_18;
const __VLS_25 = {}.ElFormItem;
/** @type { [typeof __VLS_components.ElFormItem, typeof __VLS_components.elFormItem, typeof __VLS_components.ElFormItem, typeof __VLS_components.elFormItem, ] } */ ;
// @ts-ignore
const __VLS_26 = __VLS_asFunctionalComponent(__VLS_25, new __VLS_25({
    label: ("文件上传："),
}));
const __VLS_27 = __VLS_26({
    label: ("文件上传："),
}, ...__VLS_functionalComponentArgsRest(__VLS_26));
const __VLS_31 = {}.ElUpload;
/** @type { [typeof __VLS_components.ElUpload, typeof __VLS_components.elUpload, typeof __VLS_components.ElUpload, typeof __VLS_components.elUpload, ] } */ ;
// @ts-ignore
const __VLS_32 = __VLS_asFunctionalComponent(__VLS_31, new __VLS_31({
    action: ("#"),
    ...{ class: ("upload") },
    drag: ((true)),
    limit: ((__VLS_ctx.fileLimit)),
    multiple: ((true)),
    showFileList: ((true)),
    httpRequest: ((__VLS_ctx.handleHttpUpload)),
    onExceed: ((__VLS_ctx.handleExceed)),
    onError: ((__VLS_ctx.handleUploadError)),
    beforeRemove: ((__VLS_ctx.handleBeforeRemove)),
}));
const __VLS_33 = __VLS_32({
    action: ("#"),
    ...{ class: ("upload") },
    drag: ((true)),
    limit: ((__VLS_ctx.fileLimit)),
    multiple: ((true)),
    showFileList: ((true)),
    httpRequest: ((__VLS_ctx.handleHttpUpload)),
    onExceed: ((__VLS_ctx.handleExceed)),
    onError: ((__VLS_ctx.handleUploadError)),
    beforeRemove: ((__VLS_ctx.handleBeforeRemove)),
}, ...__VLS_functionalComponentArgsRest(__VLS_32));
var __VLS_37 = {};
const __VLS_38 = {}.ElIcon;
/** @type { [typeof __VLS_components.ElIcon, typeof __VLS_components.elIcon, typeof __VLS_components.ElIcon, typeof __VLS_components.elIcon, ] } */ ;
// @ts-ignore
const __VLS_39 = __VLS_asFunctionalComponent(__VLS_38, new __VLS_38({
    ...{ class: ("el-icon--upload") },
}));
const __VLS_40 = __VLS_39({
    ...{ class: ("el-icon--upload") },
}, ...__VLS_functionalComponentArgsRest(__VLS_39));
const __VLS_44 = {}.UploadFilled;
/** @type { [typeof __VLS_components.UploadFilled, typeof __VLS_components.uploadFilled, ] } */ ;
// @ts-ignore
const __VLS_45 = __VLS_asFunctionalComponent(__VLS_44, new __VLS_44({}));
const __VLS_46 = __VLS_45({}, ...__VLS_functionalComponentArgsRest(__VLS_45));
__VLS_43.slots.default;
var __VLS_43;
__VLS_asFunctionalElement(__VLS_intrinsicElements.div, __VLS_intrinsicElements.div)({
    ...{ class: ("el-upload__text") },
});
__VLS_asFunctionalElement(__VLS_intrinsicElements.em, __VLS_intrinsicElements.em)({});
{
    const { tip: __VLS_thisSlot } = __VLS_36.slots;
    var __VLS_50 = {};
    __VLS_asFunctionalElement(__VLS_intrinsicElements.div, __VLS_intrinsicElements.div)({
        ...{ class: ("el-upload__tip") },
    });
    (__VLS_ctx.skyParams.fileSize);
}
__VLS_36.slots.default;
var __VLS_36;
__VLS_30.slots.default;
var __VLS_30;
__VLS_12.slots.default;
var __VLS_12;
{
    const { footer: __VLS_thisSlot } = __VLS_5.slots;
    const __VLS_51 = {}.ElButton;
    /** @type { [typeof __VLS_components.ElButton, typeof __VLS_components.elButton, typeof __VLS_components.ElButton, typeof __VLS_components.elButton, ] } */ ;
    // @ts-ignore
    const __VLS_52 = __VLS_asFunctionalComponent(__VLS_51, new __VLS_51({
        ...{ 'onClick': {} },
        type: ("primary"),
    }));
    const __VLS_53 = __VLS_52({
        ...{ 'onClick': {} },
        type: ("primary"),
    }, ...__VLS_functionalComponentArgsRest(__VLS_52));
    let __VLS_57;
    const __VLS_58 = {
        onClick: (__VLS_ctx.handleConfirm)
    };
    let __VLS_54;
    let __VLS_55;
    __VLS_56.slots.default;
    var __VLS_56;
    const __VLS_59 = {}.ElButton;
    /** @type { [typeof __VLS_components.ElButton, typeof __VLS_components.elButton, typeof __VLS_components.ElButton, typeof __VLS_components.elButton, ] } */ ;
    // @ts-ignore
    const __VLS_60 = __VLS_asFunctionalComponent(__VLS_59, new __VLS_59({
        ...{ 'onClick': {} },
        type: ("danger"),
    }));
    const __VLS_61 = __VLS_60({
        ...{ 'onClick': {} },
        type: ("danger"),
    }, ...__VLS_functionalComponentArgsRest(__VLS_60));
    let __VLS_65;
    const __VLS_66 = {
        onClick: (__VLS_ctx.handleCancel)
    };
    let __VLS_62;
    let __VLS_63;
    __VLS_64.slots.default;
    var __VLS_64;
}
__VLS_5.slots.default;
var __VLS_5;
['upload', 'el-icon--upload', 'el-upload__text', 'el-upload__tip',];
var __VLS_special;
const __VLS_self = (await import('vue')).defineComponent({
    setup() {
        return {
            UploadFilled: UploadFilled,
            fileLimit: fileLimit,
            dialogVisible: dialogVisible,
            skyParams: skyParams,
            handleHttpUpload: handleHttpUpload,
            handleExceed: handleExceed,
            handleUploadError: handleUploadError,
            handleConfirm: handleConfirm,
            handleCancel: handleCancel,
            handleBeforeRemove: handleBeforeRemove,
        };
    },
    emits: {},
});
const __VLS_component = (await import('vue')).defineComponent({
    setup() {
        return {
            ...__VLS_exposed,
        };
    },
    emits: {},
    __typeEl: {},
});
export default {};
; /* PartiallyEnd: #4569/main.vue */
//# sourceMappingURL=file.vue.js.map