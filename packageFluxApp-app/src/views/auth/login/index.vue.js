import settings from "@/settings";
import { useRouter } from "vue-router";
import { HOME_URL } from "@/config";
import { useUserStore } from "@/store";
import { getAssets } from "@/utils";
import { skyMsgError, skyMsgSuccess } from "@/utils/sky";
import { login } from "@/apis/auth/login";
import { Lock, Message } from "@element-plus/icons-vue";
const router = useRouter();
const logo = getAssets("images/logo/logo.png");
const formRef = ref();
const userStore = useUserStore();
const pageData = ref({
    isCountingDown: false,
    countdownSeconds: 60,
    loading: false,
    captcha: "",
    form: {
        email: "",
        password: "",
    },
    rules: {
        email: [{ required: true, message: "请输入登录的邮箱", trigger: "blur" }],
        password: [{ required: true, message: "请输入密码", trigger: "blur" }],
    },
});
const method = reactive({
    /** 登录 */
    skyLogin: async () => {
        if (!formRef.value)
            return;
        formRef.value.validate(async (valid) => {
            if (valid) {
                pageData.value.loading = true;
                try {
                    const res = await login(pageData.value.form);
                    userStore.setToken(res.data.token);
                    // 跳转到首页
                    pageData.value.loading = false;
                    skyMsgSuccess(res.msg);
                    await router.push(HOME_URL);
                }
                catch (error) {
                    // 等待1秒关闭loading
                    let loadingTime = 1;
                    setInterval(() => {
                        loadingTime--;
                        if (loadingTime === 0) {
                            pageData.value.loading = false;
                        }
                    }, 1000);
                }
            }
            else {
                skyMsgError("校验失败，信息填写有误🌻");
            }
        });
    },
});
const countdownText = computed(() => {
    return pageData.value.isCountingDown ? `${pageData.value.countdownSeconds}秒后重试` : '发送验证码';
});
; /* PartiallyEnd: #3632/scriptSetup.vue */
const __VLS_ctx = {};
let __VLS_components;
let __VLS_directives;
// CSS variable injection 
// CSS variable injection end 
__VLS_asFunctionalElement(__VLS_intrinsicElements.div, __VLS_intrinsicElements.div)({
    ...{ class: ("form center") },
});
__VLS_asFunctionalElement(__VLS_intrinsicElements.div, __VLS_intrinsicElements.div)({
    ...{ class: ("flex items-center formRef") },
});
const __VLS_0 = {}.ElImage;
/** @type { [typeof __VLS_components.ElImage, typeof __VLS_components.elImage, ] } */ ;
// @ts-ignore
const __VLS_1 = __VLS_asFunctionalComponent(__VLS_0, new __VLS_0({
    ...{ class: ("rounded-full w-36px h-36px") },
    src: ((__VLS_ctx.logo)),
}));
const __VLS_2 = __VLS_1({
    ...{ class: ("rounded-full w-36px h-36px") },
    src: ((__VLS_ctx.logo)),
}, ...__VLS_functionalComponentArgsRest(__VLS_1));
__VLS_asFunctionalElement(__VLS_intrinsicElements.div, __VLS_intrinsicElements.div)({
    ...{ class: ("ml-6px font-bold text-xl") },
});
(__VLS_ctx.settings.logoTitle || "APP 热更新 管理平台");
__VLS_asFunctionalElement(__VLS_intrinsicElements.div, __VLS_intrinsicElements.div)({
    ...{ class: ("flex items-center space-x-3 text-gray-400 mt-16px mb-16px formRef") },
});
__VLS_asFunctionalElement(__VLS_intrinsicElements.span, __VLS_intrinsicElements.span)({
    ...{ class: ("h-1px w-16 bg-gray-300 inline-block") },
});
__VLS_asFunctionalElement(__VLS_intrinsicElements.span, __VLS_intrinsicElements.span)({
    ...{ class: ("text-center") },
});
__VLS_asFunctionalElement(__VLS_intrinsicElements.span, __VLS_intrinsicElements.span)({
    ...{ class: ("h-1px w-16 bg-gray-300 inline-block") },
});
const __VLS_6 = {}.ElForm;
/** @type { [typeof __VLS_components.ElForm, typeof __VLS_components.elForm, typeof __VLS_components.ElForm, typeof __VLS_components.elForm, ] } */ ;
// @ts-ignore
const __VLS_7 = __VLS_asFunctionalComponent(__VLS_6, new __VLS_6({
    ref: ("formRef"),
    ...{ class: ("w-300px formRef") },
    model: ((__VLS_ctx.pageData.form)),
    rules: ((__VLS_ctx.pageData.rules)),
    labelWidth: ("auto"),
}));
const __VLS_8 = __VLS_7({
    ref: ("formRef"),
    ...{ class: ("w-300px formRef") },
    model: ((__VLS_ctx.pageData.form)),
    rules: ((__VLS_ctx.pageData.rules)),
    labelWidth: ("auto"),
}, ...__VLS_functionalComponentArgsRest(__VLS_7));
// @ts-ignore navigation for `const formRef = ref()`
/** @type { typeof __VLS_ctx.formRef } */ ;
var __VLS_12 = {};
const __VLS_13 = {}.ElRow;
/** @type { [typeof __VLS_components.ElRow, typeof __VLS_components.elRow, typeof __VLS_components.ElRow, typeof __VLS_components.elRow, ] } */ ;
// @ts-ignore
const __VLS_14 = __VLS_asFunctionalComponent(__VLS_13, new __VLS_13({
    gutter: ((24)),
}));
const __VLS_15 = __VLS_14({
    gutter: ((24)),
}, ...__VLS_functionalComponentArgsRest(__VLS_14));
const __VLS_19 = {}.ElCol;
/** @type { [typeof __VLS_components.ElCol, typeof __VLS_components.elCol, typeof __VLS_components.ElCol, typeof __VLS_components.elCol, ] } */ ;
// @ts-ignore
const __VLS_20 = __VLS_asFunctionalComponent(__VLS_19, new __VLS_19({
    span: ((24)),
}));
const __VLS_21 = __VLS_20({
    span: ((24)),
}, ...__VLS_functionalComponentArgsRest(__VLS_20));
const __VLS_25 = {}.ElFormItem;
/** @type { [typeof __VLS_components.ElFormItem, typeof __VLS_components.elFormItem, typeof __VLS_components.ElFormItem, typeof __VLS_components.elFormItem, ] } */ ;
// @ts-ignore
const __VLS_26 = __VLS_asFunctionalComponent(__VLS_25, new __VLS_25({
    label: ("邮箱"),
    prop: ("email"),
    required: (true),
}));
const __VLS_27 = __VLS_26({
    label: ("邮箱"),
    prop: ("email"),
    required: (true),
}, ...__VLS_functionalComponentArgsRest(__VLS_26));
const __VLS_31 = {}.ElInput;
/** @type { [typeof __VLS_components.ElInput, typeof __VLS_components.elInput, ] } */ ;
// @ts-ignore
const __VLS_32 = __VLS_asFunctionalComponent(__VLS_31, new __VLS_31({
    type: ("text"),
    modelValue: ((__VLS_ctx.pageData.form.email)),
    placeholder: ("请输入邮箱"),
    suffixIcon: ((__VLS_ctx.Message)),
}));
const __VLS_33 = __VLS_32({
    type: ("text"),
    modelValue: ((__VLS_ctx.pageData.form.email)),
    placeholder: ("请输入邮箱"),
    suffixIcon: ((__VLS_ctx.Message)),
}, ...__VLS_functionalComponentArgsRest(__VLS_32));
__VLS_30.slots.default;
var __VLS_30;
__VLS_24.slots.default;
var __VLS_24;
__VLS_18.slots.default;
var __VLS_18;
const __VLS_37 = {}.ElRow;
/** @type { [typeof __VLS_components.ElRow, typeof __VLS_components.elRow, typeof __VLS_components.ElRow, typeof __VLS_components.elRow, ] } */ ;
// @ts-ignore
const __VLS_38 = __VLS_asFunctionalComponent(__VLS_37, new __VLS_37({
    gutter: ((24)),
}));
const __VLS_39 = __VLS_38({
    gutter: ((24)),
}, ...__VLS_functionalComponentArgsRest(__VLS_38));
const __VLS_43 = {}.ElCol;
/** @type { [typeof __VLS_components.ElCol, typeof __VLS_components.elCol, typeof __VLS_components.ElCol, typeof __VLS_components.elCol, ] } */ ;
// @ts-ignore
const __VLS_44 = __VLS_asFunctionalComponent(__VLS_43, new __VLS_43({
    span: ((24)),
}));
const __VLS_45 = __VLS_44({
    span: ((24)),
}, ...__VLS_functionalComponentArgsRest(__VLS_44));
const __VLS_49 = {}.ElFormItem;
/** @type { [typeof __VLS_components.ElFormItem, typeof __VLS_components.elFormItem, typeof __VLS_components.ElFormItem, typeof __VLS_components.elFormItem, ] } */ ;
// @ts-ignore
const __VLS_50 = __VLS_asFunctionalComponent(__VLS_49, new __VLS_49({
    label: ("密码"),
    prop: ("password"),
    required: (true),
}));
const __VLS_51 = __VLS_50({
    label: ("密码"),
    prop: ("password"),
    required: (true),
}, ...__VLS_functionalComponentArgsRest(__VLS_50));
const __VLS_55 = {}.ElInput;
/** @type { [typeof __VLS_components.ElInput, typeof __VLS_components.elInput, ] } */ ;
// @ts-ignore
const __VLS_56 = __VLS_asFunctionalComponent(__VLS_55, new __VLS_55({
    ...{ 'onKeydown': {} },
    type: ("password"),
    modelValue: ((__VLS_ctx.pageData.form.password)),
    placeholder: ("请输入密码"),
    showPassword: (true),
    suffixIcon: ((__VLS_ctx.Lock)),
}));
const __VLS_57 = __VLS_56({
    ...{ 'onKeydown': {} },
    type: ("password"),
    modelValue: ((__VLS_ctx.pageData.form.password)),
    placeholder: ("请输入密码"),
    showPassword: (true),
    suffixIcon: ((__VLS_ctx.Lock)),
}, ...__VLS_functionalComponentArgsRest(__VLS_56));
let __VLS_61;
const __VLS_62 = {
    onKeydown: (__VLS_ctx.method.skyLogin)
};
let __VLS_58;
let __VLS_59;
var __VLS_60;
__VLS_54.slots.default;
var __VLS_54;
__VLS_48.slots.default;
var __VLS_48;
__VLS_42.slots.default;
var __VLS_42;
const __VLS_63 = {}.ElRow;
/** @type { [typeof __VLS_components.ElRow, typeof __VLS_components.elRow, typeof __VLS_components.ElRow, typeof __VLS_components.elRow, ] } */ ;
// @ts-ignore
const __VLS_64 = __VLS_asFunctionalComponent(__VLS_63, new __VLS_63({
    gutter: ((24)),
}));
const __VLS_65 = __VLS_64({
    gutter: ((24)),
}, ...__VLS_functionalComponentArgsRest(__VLS_64));
const __VLS_69 = {}.ElCol;
/** @type { [typeof __VLS_components.ElCol, typeof __VLS_components.elCol, typeof __VLS_components.ElCol, typeof __VLS_components.elCol, ] } */ ;
// @ts-ignore
const __VLS_70 = __VLS_asFunctionalComponent(__VLS_69, new __VLS_69({
    span: ((15)),
}));
const __VLS_71 = __VLS_70({
    span: ((15)),
}, ...__VLS_functionalComponentArgsRest(__VLS_70));
const __VLS_75 = {}.ElFormItem;
/** @type { [typeof __VLS_components.ElFormItem, typeof __VLS_components.elFormItem, typeof __VLS_components.ElFormItem, typeof __VLS_components.elFormItem, ] } */ ;
// @ts-ignore
const __VLS_76 = __VLS_asFunctionalComponent(__VLS_75, new __VLS_75({
    label: ("记住密码"),
}));
const __VLS_77 = __VLS_76({
    label: ("记住密码"),
}, ...__VLS_functionalComponentArgsRest(__VLS_76));
const __VLS_81 = {}.ElCheckbox;
/** @type { [typeof __VLS_components.ElCheckbox, typeof __VLS_components.elCheckbox, ] } */ ;
// @ts-ignore
const __VLS_82 = __VLS_asFunctionalComponent(__VLS_81, new __VLS_81({
    name: ("type"),
}));
const __VLS_83 = __VLS_82({
    name: ("type"),
}, ...__VLS_functionalComponentArgsRest(__VLS_82));
__VLS_80.slots.default;
var __VLS_80;
__VLS_74.slots.default;
var __VLS_74;
const __VLS_87 = {}.ElCol;
/** @type { [typeof __VLS_components.ElCol, typeof __VLS_components.elCol, typeof __VLS_components.ElCol, typeof __VLS_components.elCol, ] } */ ;
// @ts-ignore
const __VLS_88 = __VLS_asFunctionalComponent(__VLS_87, new __VLS_87({
    span: ((8)),
}));
const __VLS_89 = __VLS_88({
    span: ((8)),
}, ...__VLS_functionalComponentArgsRest(__VLS_88));
const __VLS_93 = {}.ElFormItem;
/** @type { [typeof __VLS_components.ElFormItem, typeof __VLS_components.elFormItem, typeof __VLS_components.ElFormItem, typeof __VLS_components.elFormItem, ] } */ ;
// @ts-ignore
const __VLS_94 = __VLS_asFunctionalComponent(__VLS_93, new __VLS_93({}));
const __VLS_95 = __VLS_94({}, ...__VLS_functionalComponentArgsRest(__VLS_94));
__VLS_asFunctionalElement(__VLS_intrinsicElements.span, __VLS_intrinsicElements.span)({
    ...{ onClick: (...[$event]) => {
            __VLS_ctx.router.push('/register');
        } },
    ...{ style: ({}) },
});
__VLS_98.slots.default;
var __VLS_98;
__VLS_92.slots.default;
var __VLS_92;
__VLS_68.slots.default;
var __VLS_68;
const __VLS_99 = {}.ElRow;
/** @type { [typeof __VLS_components.ElRow, typeof __VLS_components.elRow, typeof __VLS_components.ElRow, typeof __VLS_components.elRow, ] } */ ;
// @ts-ignore
const __VLS_100 = __VLS_asFunctionalComponent(__VLS_99, new __VLS_99({
    gutter: ((24)),
}));
const __VLS_101 = __VLS_100({
    gutter: ((24)),
}, ...__VLS_functionalComponentArgsRest(__VLS_100));
const __VLS_105 = {}.ElCol;
/** @type { [typeof __VLS_components.ElCol, typeof __VLS_components.elCol, typeof __VLS_components.ElCol, typeof __VLS_components.elCol, ] } */ ;
// @ts-ignore
const __VLS_106 = __VLS_asFunctionalComponent(__VLS_105, new __VLS_105({
    span: ((24)),
}));
const __VLS_107 = __VLS_106({
    span: ((24)),
}, ...__VLS_functionalComponentArgsRest(__VLS_106));
const __VLS_111 = {}.ElFormItem;
/** @type { [typeof __VLS_components.ElFormItem, typeof __VLS_components.elFormItem, typeof __VLS_components.ElFormItem, typeof __VLS_components.elFormItem, ] } */ ;
// @ts-ignore
const __VLS_112 = __VLS_asFunctionalComponent(__VLS_111, new __VLS_111({}));
const __VLS_113 = __VLS_112({}, ...__VLS_functionalComponentArgsRest(__VLS_112));
if (!__VLS_ctx.pageData.loading) {
    const __VLS_117 = {}.ElButton;
    /** @type { [typeof __VLS_components.ElButton, typeof __VLS_components.elButton, typeof __VLS_components.ElButton, typeof __VLS_components.elButton, ] } */ ;
    // @ts-ignore
    const __VLS_118 = __VLS_asFunctionalComponent(__VLS_117, new __VLS_117({
        ...{ 'onClick': {} },
        ...{ style: ({}) },
        ...{ class: ("bg-[--el-color-primary]") },
        type: ("primary"),
        round: (true),
    }));
    const __VLS_119 = __VLS_118({
        ...{ 'onClick': {} },
        ...{ style: ({}) },
        ...{ class: ("bg-[--el-color-primary]") },
        type: ("primary"),
        round: (true),
    }, ...__VLS_functionalComponentArgsRest(__VLS_118));
    let __VLS_123;
    const __VLS_124 = {
        onClick: (__VLS_ctx.method.skyLogin)
    };
    let __VLS_120;
    let __VLS_121;
    __VLS_122.slots.default;
    var __VLS_122;
}
if (__VLS_ctx.pageData.loading) {
    const __VLS_125 = {}.ElButton;
    /** @type { [typeof __VLS_components.ElButton, typeof __VLS_components.elButton, typeof __VLS_components.ElButton, typeof __VLS_components.elButton, ] } */ ;
    // @ts-ignore
    const __VLS_126 = __VLS_asFunctionalComponent(__VLS_125, new __VLS_125({
        ...{ style: ({}) },
        ...{ class: ("bg-[--el-color-primary]") },
        type: ("primary"),
        round: (true),
        loading: ((__VLS_ctx.pageData.loading)),
    }));
    const __VLS_127 = __VLS_126({
        ...{ style: ({}) },
        ...{ class: ("bg-[--el-color-primary]") },
        type: ("primary"),
        round: (true),
        loading: ((__VLS_ctx.pageData.loading)),
    }, ...__VLS_functionalComponentArgsRest(__VLS_126));
    __VLS_130.slots.default;
    var __VLS_130;
}
__VLS_116.slots.default;
var __VLS_116;
__VLS_110.slots.default;
var __VLS_110;
__VLS_104.slots.default;
var __VLS_104;
__VLS_11.slots.default;
var __VLS_11;
['form', 'center', 'flex', 'items-center', 'formRef', 'rounded-full', 'w-36px', 'h-36px', 'ml-6px', 'font-bold', 'text-xl', 'flex', 'items-center', 'space-x-3', 'text-gray-400', 'mt-16px', 'mb-16px', 'formRef', 'h-1px', 'w-16', 'bg-gray-300', 'inline-block', 'text-center', 'h-1px', 'w-16', 'bg-gray-300', 'inline-block', 'w-300px', 'formRef', 'bg-[--el-color-primary]', 'bg-[--el-color-primary]',];
var __VLS_special;
const __VLS_self = (await import('vue')).defineComponent({
    setup() {
        return {
            settings: settings,
            Lock: Lock,
            Message: Message,
            router: router,
            logo: logo,
            formRef: formRef,
            pageData: pageData,
            method: method,
        };
    },
});
export default (await import('vue')).defineComponent({
    setup() {
        return {};
    },
    __typeRefs: {},
    __typeEl: {},
});
; /* PartiallyEnd: #4569/main.vue */
//# sourceMappingURL=index.vue.js.map