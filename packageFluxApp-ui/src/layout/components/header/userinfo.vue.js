/// <reference types=".vue-global-types/vue_3.5_false.d.ts" />
import { useRouter } from "vue-router";
import { LOGIN_URL } from "@/config";
import { useUserStore } from "@/store";
import { getAssets } from "@/utils";
import { skyMsgError, skyMsgSuccess } from "@/utils/sky";
import { skySessionStorage } from "@/utils/storage";
import { logout, userinfo } from "@/apis/auth/login";
import { ArrowDown } from "@element-plus/icons-vue";
const { defineProps, defineSlots, defineEmits, defineExpose, defineModel, defineOptions, withDefaults, } = await import('vue');
const router = useRouter();
const avatar = getAssets("images/auth/wuwu.jpg");
const errorAvatar = ref("");
const userStore = useUserStore();
const pageData = ref({
    userInfo: {
        email: "",
        avatar: "",
    },
});
const method = reactive({
    handleGetUserInfo: async () => {
        try {
            const res = await userinfo();
            userStore.setUser(res.data);
            pageData.value.userInfo = res.data;
        }
        catch (error) {
            skyMsgError("数据查询失败，请刷新重试🌻");
        }
    },
    handleLogout: async () => {
        try {
            const res = await logout();
            // 清除缓存
            skySessionStorage.clear();
            // 清除用户信息
            userStore.setToken("");
            userStore.setUser({});
            skyMsgSuccess("注销成功");
            // 退出登录。必须使用replace把页面缓存刷掉。
            window.location.replace(LOGIN_URL);
        }
        catch (error) {
            skyMsgError("操作失败，请刷新重试🌻");
        }
    }
});
// 下拉折叠
const handleCommand = (command) => {
    switch (command) {
        case "skyMine":
            router.push("/system/personage");
            break;
        case "logout":
            method.handleLogout();
            break;
    }
};
onMounted(() => {
    method.handleGetUserInfo();
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
    const __VLS_0 = __VLS_resolvedLocalAndGlobalComponents.ElImage;
    /** @type { [typeof __VLS_components.ElImage, typeof __VLS_components.elImage, typeof __VLS_components.ElImage, typeof __VLS_components.elImage, ] } */
    // @ts-ignore
    const __VLS_1 = __VLS_asFunctionalComponent(__VLS_0, new __VLS_0({ ...{ class: ("w-34px h-34px rounded-full select-none user-avatar") }, src: ((__VLS_ctx.pageData.userInfo?.avatar || __VLS_ctx.avatar)), }));
    const __VLS_2 = __VLS_1({ ...{ class: ("w-34px h-34px rounded-full select-none user-avatar") }, src: ((__VLS_ctx.pageData.userInfo?.avatar || __VLS_ctx.avatar)), }, ...__VLS_functionalComponentArgsRest(__VLS_1));
    __VLS_elementAsFunction(__VLS_intrinsicElements.template, __VLS_intrinsicElements.template)({});
    {
        const { error: __VLS_thisSlot } = __VLS_nonNullable(__VLS_5.slots);
        const __VLS_6 = __VLS_resolvedLocalAndGlobalComponents.ElImage;
        /** @type { [typeof __VLS_components.ElImage, typeof __VLS_components.elImage, typeof __VLS_components.ElImage, typeof __VLS_components.elImage, ] } */
        // @ts-ignore
        const __VLS_7 = __VLS_asFunctionalComponent(__VLS_6, new __VLS_6({ ...{ class: ("w-34px h-34px rounded-full select-none user-avatar") }, src: ((__VLS_ctx.errorAvatar)), }));
        const __VLS_8 = __VLS_7({ ...{ class: ("w-34px h-34px rounded-full select-none user-avatar") }, src: ((__VLS_ctx.errorAvatar)), }, ...__VLS_functionalComponentArgsRest(__VLS_7));
    }
    const __VLS_5 = __VLS_pickFunctionalComponentCtx(__VLS_0, __VLS_2);
    const __VLS_12 = __VLS_resolvedLocalAndGlobalComponents.ElDropdown;
    /** @type { [typeof __VLS_components.ElDropdown, typeof __VLS_components.elDropdown, typeof __VLS_components.ElDropdown, typeof __VLS_components.elDropdown, ] } */
    // @ts-ignore
    const __VLS_13 = __VLS_asFunctionalComponent(__VLS_12, new __VLS_12({ ...{ 'onCommand': {} }, ...{ class: ("m-l-10px") }, hideOnClick: ((false)), }));
    const __VLS_14 = __VLS_13({ ...{ 'onCommand': {} }, ...{ class: ("m-l-10px") }, hideOnClick: ((false)), }, ...__VLS_functionalComponentArgsRest(__VLS_13));
    let __VLS_18;
    const __VLS_19 = {
        onCommand: (__VLS_ctx.handleCommand)
    };
    let __VLS_15;
    let __VLS_16;
    __VLS_elementAsFunction(__VLS_intrinsicElements.div, __VLS_intrinsicElements.div)({ ...{ class: ("sky-dropdown") }, });
    __VLS_elementAsFunction(__VLS_intrinsicElements.div, __VLS_intrinsicElements.div)({ ...{ class: ("max-w-113px text-14px m-r-6px line-clamp-1") }, });
    (__VLS_ctx.pageData.userInfo.email);
    const __VLS_20 = __VLS_resolvedLocalAndGlobalComponents.ElIcon;
    /** @type { [typeof __VLS_components.ElIcon, typeof __VLS_components.elIcon, typeof __VLS_components.ElIcon, typeof __VLS_components.elIcon, ] } */
    // @ts-ignore
    const __VLS_21 = __VLS_asFunctionalComponent(__VLS_20, new __VLS_20({}));
    const __VLS_22 = __VLS_21({}, ...__VLS_functionalComponentArgsRest(__VLS_21));
    const __VLS_26 = __VLS_resolvedLocalAndGlobalComponents.ArrowDown;
    /** @type { [typeof __VLS_components.ArrowDown, typeof __VLS_components.arrowDown, ] } */
    // @ts-ignore
    const __VLS_27 = __VLS_asFunctionalComponent(__VLS_26, new __VLS_26({}));
    const __VLS_28 = __VLS_27({}, ...__VLS_functionalComponentArgsRest(__VLS_27));
    __VLS_nonNullable(__VLS_25.slots).default;
    const __VLS_25 = __VLS_pickFunctionalComponentCtx(__VLS_20, __VLS_22);
    __VLS_elementAsFunction(__VLS_intrinsicElements.template, __VLS_intrinsicElements.template)({});
    {
        const { dropdown: __VLS_thisSlot } = __VLS_nonNullable(__VLS_17.slots);
        const __VLS_32 = __VLS_resolvedLocalAndGlobalComponents.ElDropdownMenu;
        /** @type { [typeof __VLS_components.ElDropdownMenu, typeof __VLS_components.elDropdownMenu, typeof __VLS_components.ElDropdownMenu, typeof __VLS_components.elDropdownMenu, ] } */
        // @ts-ignore
        const __VLS_33 = __VLS_asFunctionalComponent(__VLS_32, new __VLS_32({}));
        const __VLS_34 = __VLS_33({}, ...__VLS_functionalComponentArgsRest(__VLS_33));
        const __VLS_38 = __VLS_resolvedLocalAndGlobalComponents.ElDropdownItem;
        /** @type { [typeof __VLS_components.ElDropdownItem, typeof __VLS_components.elDropdownItem, typeof __VLS_components.ElDropdownItem, typeof __VLS_components.elDropdownItem, ] } */
        // @ts-ignore
        const __VLS_39 = __VLS_asFunctionalComponent(__VLS_38, new __VLS_38({ command: ("logout"), }));
        const __VLS_40 = __VLS_39({ command: ("logout"), }, ...__VLS_functionalComponentArgsRest(__VLS_39));
        __VLS_nonNullable(__VLS_43.slots).default;
        const __VLS_43 = __VLS_pickFunctionalComponentCtx(__VLS_38, __VLS_40);
        __VLS_nonNullable(__VLS_37.slots).default;
        const __VLS_37 = __VLS_pickFunctionalComponentCtx(__VLS_32, __VLS_34);
    }
    const __VLS_17 = __VLS_pickFunctionalComponentCtx(__VLS_12, __VLS_14);
    __VLS_styleScopedClasses['w-34px'];
    __VLS_styleScopedClasses['h-34px'];
    __VLS_styleScopedClasses['rounded-full'];
    __VLS_styleScopedClasses['select-none'];
    __VLS_styleScopedClasses['user-avatar'];
    __VLS_styleScopedClasses['w-34px'];
    __VLS_styleScopedClasses['h-34px'];
    __VLS_styleScopedClasses['rounded-full'];
    __VLS_styleScopedClasses['select-none'];
    __VLS_styleScopedClasses['user-avatar'];
    __VLS_styleScopedClasses['m-l-10px'];
    __VLS_styleScopedClasses['sky-dropdown'];
    __VLS_styleScopedClasses['max-w-113px'];
    __VLS_styleScopedClasses['text-14px'];
    __VLS_styleScopedClasses['m-r-6px'];
    __VLS_styleScopedClasses['line-clamp-1'];
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
            ArrowDown: ArrowDown,
            avatar: avatar,
            errorAvatar: errorAvatar,
            pageData: pageData,
            handleCommand: handleCommand,
        };
    },
});
export default (await import('vue')).defineComponent({
    setup() {
        return {};
    },
});
;
//# sourceMappingURL=userinfo.vue.js.map