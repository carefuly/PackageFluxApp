/// <reference types=".vue-global-types/vue_3.5_false.d.ts" />
import SkyUploadFile from "@/components/SkyUpload/file.vue";
const { defineProps, defineSlots, defineEmits, defineExpose, defineModel, defineOptions, withDefaults, } = await import('vue');
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
const __VLS_fnComponent = (await import('vue')).defineComponent({});
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
    __VLS_elementAsFunction(__VLS_intrinsicElements.div, __VLS_intrinsicElements.div)({ ...{ class: ("file") }, });
    const __VLS_0 = __VLS_resolvedLocalAndGlobalComponents.ElButton;
    /** @type { [typeof __VLS_components.ElButton, typeof __VLS_components.elButton, typeof __VLS_components.ElButton, typeof __VLS_components.elButton, ] } */
    // @ts-ignore
    const __VLS_1 = __VLS_asFunctionalComponent(__VLS_0, new __VLS_0({ ...{ 'onClick': {} }, type: ("primary"), icon: ("Upload"), link: (true), }));
    const __VLS_2 = __VLS_1({ ...{ 'onClick': {} }, type: ("primary"), icon: ("Upload"), link: (true), }, ...__VLS_functionalComponentArgsRest(__VLS_1));
    let __VLS_6;
    const __VLS_7 = {
        onClick: (__VLS_ctx.method.handleImportFile)
    };
    let __VLS_3;
    let __VLS_4;
    __VLS_nonNullable(__VLS_5.slots).default;
    const __VLS_5 = __VLS_pickFunctionalComponentCtx(__VLS_0, __VLS_2);
    // @ts-ignore
    [SkyUploadFile, SkyUploadFile,];
    // @ts-ignore
    const __VLS_8 = __VLS_asFunctionalComponent(SkyUploadFile, new SkyUploadFile({ ...{ 'onHandleConfirmUpload': {} }, ref: ("skyFileRef"), }));
    const __VLS_9 = __VLS_8({ ...{ 'onHandleConfirmUpload': {} }, ref: ("skyFileRef"), }, ...__VLS_functionalComponentArgsRest(__VLS_8));
    // @ts-ignore navigation for `const skyFileRef = ref()`
    __VLS_ctx.skyFileRef;
    var __VLS_13 = {};
    let __VLS_14;
    const __VLS_15 = {
        onHandleConfirmUpload: (__VLS_ctx.method.handleConfirmUpload)
    };
    let __VLS_10;
    let __VLS_11;
    const __VLS_12 = __VLS_pickFunctionalComponentCtx(SkyUploadFile, __VLS_9);
    __VLS_styleScopedClasses['file'];
    var __VLS_slots;
    var __VLS_inheritedAttrs;
    const __VLS_refs = {
        "skyFileRef": __VLS_13,
    };
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
            SkyUploadFile: SkyUploadFile,
            skyFileRef: skyFileRef,
            method: method,
        };
    },
});
export default (await import('vue')).defineComponent({
    setup() {
        return {};
    },
});
;
//# sourceMappingURL=upload.vue.js.map