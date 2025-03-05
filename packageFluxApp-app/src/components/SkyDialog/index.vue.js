import { skyMsgWarning, skyMsgBox } from "@/utils/sky";
const props = defineProps({
    title: { default: "朕很中意你SkyDialog" },
    width: { default: 650 },
    height: { default: 300 },
    visible: { default: false },
    confirmText: { default: "确定" },
    cancelText: { default: "取消" },
    destroyOnClose: { default: false },
    fullscreen: { default: false },
    loading: { default: false },
    footerHidden: { default: false },
});
// 开关变量
const visible = ref(false);
// 确定的loading，此处必须用toRefs，否则将失去响应式
const { loading } = toRefs(props);
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
const __VLS_exposed = {
    skyOpen,
    skyClose,
    skyQuickClose
};
defineExpose({
    skyOpen,
    skyClose,
    skyQuickClose
});
; /* PartiallyEnd: #3632/scriptSetup.vue */
const __VLS_ctx = {};
let __VLS_components;
let __VLS_directives;
['el-dialog__body',];
// CSS variable injection 
// CSS variable injection end 
const __VLS_0 = {}.ElDialog;
/** @type { [typeof __VLS_components.ElDialog, typeof __VLS_components.elDialog, typeof __VLS_components.ElDialog, typeof __VLS_components.elDialog, ] } */ ;
// @ts-ignore
const __VLS_1 = __VLS_asFunctionalComponent(__VLS_0, new __VLS_0({
    modelValue: ((__VLS_ctx.visible)),
    title: ((props.title)),
    width: ((props.width)),
    closeOnClickModal: ((false)),
    appendToBody: (true),
    draggable: (true),
    destroyOnClose: ((props.destroyOnClose)),
    beforeClose: ((__VLS_ctx.skyClose)),
    fullscreen: ((props.fullscreen)),
    loading: ((props.loading)),
    footerHidden: ((props.footerHidden)),
}));
const __VLS_2 = __VLS_1({
    modelValue: ((__VLS_ctx.visible)),
    title: ((props.title)),
    width: ((props.width)),
    closeOnClickModal: ((false)),
    appendToBody: (true),
    draggable: (true),
    destroyOnClose: ((props.destroyOnClose)),
    beforeClose: ((__VLS_ctx.skyClose)),
    fullscreen: ((props.fullscreen)),
    loading: ((props.loading)),
    footerHidden: ((props.footerHidden)),
}, ...__VLS_functionalComponentArgsRest(__VLS_1));
var __VLS_6 = {};
var __VLS_7 = {};
__VLS_asFunctionalElement(__VLS_intrinsicElements.div, __VLS_intrinsicElements.div)({
    ...{ class: ("container") },
    ...{ style: (({ height: props.height + 'px' })) },
});
var __VLS_8 = {};
if (!props.footerHidden) {
    {
        const { footer: __VLS_thisSlot } = __VLS_5.slots;
        __VLS_asFunctionalElement(__VLS_intrinsicElements.span, __VLS_intrinsicElements.span)({
            ...{ class: ("dialog-footer") },
        });
        const __VLS_9 = {}.ElButton;
        /** @type { [typeof __VLS_components.ElButton, typeof __VLS_components.elButton, typeof __VLS_components.ElButton, typeof __VLS_components.elButton, ] } */ ;
        // @ts-ignore
        const __VLS_10 = __VLS_asFunctionalComponent(__VLS_9, new __VLS_9({
            type: ("primary"),
            loadingIcon: ("Eleme"),
            loading: ((__VLS_ctx.confirmLoading)),
        }));
        const __VLS_11 = __VLS_10({
            type: ("primary"),
            loadingIcon: ("Eleme"),
            loading: ((__VLS_ctx.confirmLoading)),
        }, ...__VLS_functionalComponentArgsRest(__VLS_10));
        __VLS_asFunctionalDirective(__VLS_directives.vThrottle)(null, { ...__VLS_directiveBindingRestFields, value: (__VLS_ctx.skyConfirm) }, null, null);
        (props.confirmText);
        __VLS_14.slots.default;
        var __VLS_14;
        const __VLS_15 = {}.ElButton;
        /** @type { [typeof __VLS_components.ElButton, typeof __VLS_components.elButton, typeof __VLS_components.ElButton, typeof __VLS_components.elButton, ] } */ ;
        // @ts-ignore
        const __VLS_16 = __VLS_asFunctionalComponent(__VLS_15, new __VLS_15({
            ...{ 'onClick': {} },
            type: ("danger"),
        }));
        const __VLS_17 = __VLS_16({
            ...{ 'onClick': {} },
            type: ("danger"),
        }, ...__VLS_functionalComponentArgsRest(__VLS_16));
        let __VLS_21;
        const __VLS_22 = {
            onClick: (__VLS_ctx.skyCancel)
        };
        let __VLS_18;
        let __VLS_19;
        (props.cancelText);
        __VLS_20.slots.default;
        var __VLS_20;
    }
}
__VLS_5.slots.default;
var __VLS_5;
['container', 'dialog-footer',];
var __VLS_special;
const __VLS_self = (await import('vue')).defineComponent({
    setup() {
        return {
            visible: visible,
            confirmLoading: confirmLoading,
            skyClose: skyClose,
            skyConfirm: skyConfirm,
            skyCancel: skyCancel,
        };
    },
    emits: {},
    props: {
        title: { default: "朕很中意你SkyDialog" },
        width: { default: 650 },
        height: { default: 300 },
        visible: { default: false },
        confirmText: { default: "确定" },
        cancelText: { default: "取消" },
        destroyOnClose: { default: false },
        fullscreen: { default: false },
        loading: { default: false },
        footerHidden: { default: false },
    },
});
const __VLS_component = (await import('vue')).defineComponent({
    setup() {
        return {
            ...__VLS_exposed,
        };
    },
    emits: {},
    props: {
        title: { default: "朕很中意你SkyDialog" },
        width: { default: 650 },
        height: { default: 300 },
        visible: { default: false },
        confirmText: { default: "确定" },
        cancelText: { default: "取消" },
        destroyOnClose: { default: false },
        fullscreen: { default: false },
        loading: { default: false },
        footerHidden: { default: false },
    },
    __typeEl: {},
});
export default {};
; /* PartiallyEnd: #4569/main.vue */
//# sourceMappingURL=index.vue.js.map