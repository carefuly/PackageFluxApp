import settings from "@/settings";
import { useRouter } from "vue-router";
import { LOGIN_URL } from "@/config";
import { getAssets } from "@/utils";
import { skyMsgError, skyMsgSuccess } from "@/utils/sky";
import { register } from "@/apis/auth/register";
import { User, Lock, Message } from "@element-plus/icons-vue";
const router = useRouter();
const logo = getAssets("images/logo/logo.png");
const formRef = ref();
const pageData = ref({
    isCountingDown: false,
    countdownSeconds: 60,
    loading: false,
    form: {
        username: "",
        email: "",
        password: "",
    },
    rules: {
        username: [{ required: true, message: "请输入用户账号", trigger: "blur" }],
        email: [{ required: true, message: "请输入登录的邮箱", trigger: "blur" }],
        password: [{ required: true, message: "请输入密码", trigger: "blur" }],
    },
});
const method = reactive({
    // 注册
    skyRegister: async () => {
        if (!formRef.value)
            return;
        formRef.value.validate(async (valid) => {
            if (valid) {
                pageData.value.loading = true;
                try {
                    await register(pageData.value.form);
                    // 跳转到登录
                    pageData.value.loading = false;
                    skyMsgSuccess("注册成功🌻");
                    await router.push(LOGIN_URL);
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
    label: ("用户名"),
    prop: ("username"),
    required: (true),
}));
const __VLS_27 = __VLS_26({
    label: ("用户名"),
    prop: ("username"),
    required: (true),
}, ...__VLS_functionalComponentArgsRest(__VLS_26));
const __VLS_31 = {}.ElInput;
/** @type { [typeof __VLS_components.ElInput, typeof __VLS_components.elInput, ] } */ ;
// @ts-ignore
const __VLS_32 = __VLS_asFunctionalComponent(__VLS_31, new __VLS_31({
    type: ("text"),
    modelValue: ((__VLS_ctx.pageData.form.username)),
    placeholder: ("请输入用户名"),
    suffixIcon: ((__VLS_ctx.User)),
}));
const __VLS_33 = __VLS_32({
    type: ("text"),
    modelValue: ((__VLS_ctx.pageData.form.username)),
    placeholder: ("请输入用户名"),
    suffixIcon: ((__VLS_ctx.User)),
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
    label: ("邮箱"),
    prop: ("email"),
    required: (true),
}));
const __VLS_51 = __VLS_50({
    label: ("邮箱"),
    prop: ("email"),
    required: (true),
}, ...__VLS_functionalComponentArgsRest(__VLS_50));
const __VLS_55 = {}.ElInput;
/** @type { [typeof __VLS_components.ElInput, typeof __VLS_components.elInput, ] } */ ;
// @ts-ignore
const __VLS_56 = __VLS_asFunctionalComponent(__VLS_55, new __VLS_55({
    type: ("text"),
    modelValue: ((__VLS_ctx.pageData.form.email)),
    placeholder: ("请输入邮箱"),
    suffixIcon: ((__VLS_ctx.Message)),
}));
const __VLS_57 = __VLS_56({
    type: ("text"),
    modelValue: ((__VLS_ctx.pageData.form.email)),
    placeholder: ("请输入邮箱"),
    suffixIcon: ((__VLS_ctx.Message)),
}, ...__VLS_functionalComponentArgsRest(__VLS_56));
__VLS_54.slots.default;
var __VLS_54;
__VLS_48.slots.default;
var __VLS_48;
__VLS_42.slots.default;
var __VLS_42;
const __VLS_61 = {}.ElRow;
/** @type { [typeof __VLS_components.ElRow, typeof __VLS_components.elRow, typeof __VLS_components.ElRow, typeof __VLS_components.elRow, ] } */ ;
// @ts-ignore
const __VLS_62 = __VLS_asFunctionalComponent(__VLS_61, new __VLS_61({
    gutter: ((24)),
}));
const __VLS_63 = __VLS_62({
    gutter: ((24)),
}, ...__VLS_functionalComponentArgsRest(__VLS_62));
const __VLS_67 = {}.ElCol;
/** @type { [typeof __VLS_components.ElCol, typeof __VLS_components.elCol, typeof __VLS_components.ElCol, typeof __VLS_components.elCol, ] } */ ;
// @ts-ignore
const __VLS_68 = __VLS_asFunctionalComponent(__VLS_67, new __VLS_67({
    span: ((24)),
}));
const __VLS_69 = __VLS_68({
    span: ((24)),
}, ...__VLS_functionalComponentArgsRest(__VLS_68));
const __VLS_73 = {}.ElFormItem;
/** @type { [typeof __VLS_components.ElFormItem, typeof __VLS_components.elFormItem, typeof __VLS_components.ElFormItem, typeof __VLS_components.elFormItem, ] } */ ;
// @ts-ignore
const __VLS_74 = __VLS_asFunctionalComponent(__VLS_73, new __VLS_73({
    label: ("密码"),
    prop: ("password"),
    required: (true),
}));
const __VLS_75 = __VLS_74({
    label: ("密码"),
    prop: ("password"),
    required: (true),
}, ...__VLS_functionalComponentArgsRest(__VLS_74));
const __VLS_79 = {}.ElInput;
/** @type { [typeof __VLS_components.ElInput, typeof __VLS_components.elInput, ] } */ ;
// @ts-ignore
const __VLS_80 = __VLS_asFunctionalComponent(__VLS_79, new __VLS_79({
    ...{ 'onKeydown': {} },
    type: ("password"),
    modelValue: ((__VLS_ctx.pageData.form.password)),
    placeholder: ("请输入密码"),
    showPassword: (true),
    suffixIcon: ((__VLS_ctx.Lock)),
}));
const __VLS_81 = __VLS_80({
    ...{ 'onKeydown': {} },
    type: ("password"),
    modelValue: ((__VLS_ctx.pageData.form.password)),
    placeholder: ("请输入密码"),
    showPassword: (true),
    suffixIcon: ((__VLS_ctx.Lock)),
}, ...__VLS_functionalComponentArgsRest(__VLS_80));
let __VLS_85;
const __VLS_86 = {
    onKeydown: (__VLS_ctx.method.skyRegister)
};
let __VLS_82;
let __VLS_83;
var __VLS_84;
__VLS_78.slots.default;
var __VLS_78;
__VLS_72.slots.default;
var __VLS_72;
__VLS_66.slots.default;
var __VLS_66;
const __VLS_87 = {}.ElRow;
/** @type { [typeof __VLS_components.ElRow, typeof __VLS_components.elRow, typeof __VLS_components.ElRow, typeof __VLS_components.elRow, ] } */ ;
// @ts-ignore
const __VLS_88 = __VLS_asFunctionalComponent(__VLS_87, new __VLS_87({
    gutter: ((24)),
}));
const __VLS_89 = __VLS_88({
    gutter: ((24)),
}, ...__VLS_functionalComponentArgsRest(__VLS_88));
const __VLS_93 = {}.ElCol;
/** @type { [typeof __VLS_components.ElCol, typeof __VLS_components.elCol, typeof __VLS_components.ElCol, typeof __VLS_components.elCol, ] } */ ;
// @ts-ignore
const __VLS_94 = __VLS_asFunctionalComponent(__VLS_93, new __VLS_93({
    span: ((15)),
}));
const __VLS_95 = __VLS_94({
    span: ((15)),
}, ...__VLS_functionalComponentArgsRest(__VLS_94));
const __VLS_99 = {}.ElFormItem;
/** @type { [typeof __VLS_components.ElFormItem, typeof __VLS_components.elFormItem, typeof __VLS_components.ElFormItem, typeof __VLS_components.elFormItem, ] } */ ;
// @ts-ignore
const __VLS_100 = __VLS_asFunctionalComponent(__VLS_99, new __VLS_99({
    label: ("记住密码"),
}));
const __VLS_101 = __VLS_100({
    label: ("记住密码"),
}, ...__VLS_functionalComponentArgsRest(__VLS_100));
const __VLS_105 = {}.ElCheckbox;
/** @type { [typeof __VLS_components.ElCheckbox, typeof __VLS_components.elCheckbox, ] } */ ;
// @ts-ignore
const __VLS_106 = __VLS_asFunctionalComponent(__VLS_105, new __VLS_105({
    name: ("type"),
}));
const __VLS_107 = __VLS_106({
    name: ("type"),
}, ...__VLS_functionalComponentArgsRest(__VLS_106));
__VLS_104.slots.default;
var __VLS_104;
__VLS_98.slots.default;
var __VLS_98;
const __VLS_111 = {}.ElCol;
/** @type { [typeof __VLS_components.ElCol, typeof __VLS_components.elCol, typeof __VLS_components.ElCol, typeof __VLS_components.elCol, ] } */ ;
// @ts-ignore
const __VLS_112 = __VLS_asFunctionalComponent(__VLS_111, new __VLS_111({
    span: ((8)),
}));
const __VLS_113 = __VLS_112({
    span: ((8)),
}, ...__VLS_functionalComponentArgsRest(__VLS_112));
const __VLS_117 = {}.ElFormItem;
/** @type { [typeof __VLS_components.ElFormItem, typeof __VLS_components.elFormItem, typeof __VLS_components.ElFormItem, typeof __VLS_components.elFormItem, ] } */ ;
// @ts-ignore
const __VLS_118 = __VLS_asFunctionalComponent(__VLS_117, new __VLS_117({}));
const __VLS_119 = __VLS_118({}, ...__VLS_functionalComponentArgsRest(__VLS_118));
__VLS_asFunctionalElement(__VLS_intrinsicElements.span, __VLS_intrinsicElements.span)({
    ...{ onClick: (...[$event]) => {
            __VLS_ctx.router.push('/login');
        } },
    ...{ style: ({}) },
});
__VLS_122.slots.default;
var __VLS_122;
__VLS_116.slots.default;
var __VLS_116;
__VLS_92.slots.default;
var __VLS_92;
const __VLS_123 = {}.ElRow;
/** @type { [typeof __VLS_components.ElRow, typeof __VLS_components.elRow, typeof __VLS_components.ElRow, typeof __VLS_components.elRow, ] } */ ;
// @ts-ignore
const __VLS_124 = __VLS_asFunctionalComponent(__VLS_123, new __VLS_123({
    gutter: ((24)),
}));
const __VLS_125 = __VLS_124({
    gutter: ((24)),
}, ...__VLS_functionalComponentArgsRest(__VLS_124));
const __VLS_129 = {}.ElCol;
/** @type { [typeof __VLS_components.ElCol, typeof __VLS_components.elCol, typeof __VLS_components.ElCol, typeof __VLS_components.elCol, ] } */ ;
// @ts-ignore
const __VLS_130 = __VLS_asFunctionalComponent(__VLS_129, new __VLS_129({
    span: ((24)),
}));
const __VLS_131 = __VLS_130({
    span: ((24)),
}, ...__VLS_functionalComponentArgsRest(__VLS_130));
const __VLS_135 = {}.ElFormItem;
/** @type { [typeof __VLS_components.ElFormItem, typeof __VLS_components.elFormItem, typeof __VLS_components.ElFormItem, typeof __VLS_components.elFormItem, ] } */ ;
// @ts-ignore
const __VLS_136 = __VLS_asFunctionalComponent(__VLS_135, new __VLS_135({}));
const __VLS_137 = __VLS_136({}, ...__VLS_functionalComponentArgsRest(__VLS_136));
if (!__VLS_ctx.pageData.loading) {
    const __VLS_141 = {}.ElButton;
    /** @type { [typeof __VLS_components.ElButton, typeof __VLS_components.elButton, typeof __VLS_components.ElButton, typeof __VLS_components.elButton, ] } */ ;
    // @ts-ignore
    const __VLS_142 = __VLS_asFunctionalComponent(__VLS_141, new __VLS_141({
        ...{ 'onClick': {} },
        ...{ style: ({}) },
        type: ("primary"),
        round: (true),
    }));
    const __VLS_143 = __VLS_142({
        ...{ 'onClick': {} },
        ...{ style: ({}) },
        type: ("primary"),
        round: (true),
    }, ...__VLS_functionalComponentArgsRest(__VLS_142));
    let __VLS_147;
    const __VLS_148 = {
        onClick: (__VLS_ctx.method.skyRegister)
    };
    let __VLS_144;
    let __VLS_145;
    __VLS_146.slots.default;
    var __VLS_146;
}
if (__VLS_ctx.pageData.loading) {
    const __VLS_149 = {}.ElButton;
    /** @type { [typeof __VLS_components.ElButton, typeof __VLS_components.elButton, typeof __VLS_components.ElButton, typeof __VLS_components.elButton, ] } */ ;
    // @ts-ignore
    const __VLS_150 = __VLS_asFunctionalComponent(__VLS_149, new __VLS_149({
        ...{ style: ({}) },
        type: ("primary"),
        round: (true),
        loading: ((__VLS_ctx.pageData.loading)),
    }));
    const __VLS_151 = __VLS_150({
        ...{ style: ({}) },
        type: ("primary"),
        round: (true),
        loading: ((__VLS_ctx.pageData.loading)),
    }, ...__VLS_functionalComponentArgsRest(__VLS_150));
    __VLS_154.slots.default;
    var __VLS_154;
}
__VLS_140.slots.default;
var __VLS_140;
__VLS_134.slots.default;
var __VLS_134;
__VLS_128.slots.default;
var __VLS_128;
__VLS_11.slots.default;
var __VLS_11;
['form', 'center', 'flex', 'items-center', 'formRef', 'rounded-full', 'w-36px', 'h-36px', 'ml-6px', 'font-bold', 'text-xl', 'flex', 'items-center', 'space-x-3', 'text-gray-400', 'mt-16px', 'mb-16px', 'formRef', 'h-1px', 'w-16', 'bg-gray-300', 'inline-block', 'text-center', 'h-1px', 'w-16', 'bg-gray-300', 'inline-block', 'w-300px', 'formRef',];
var __VLS_special;
const __VLS_self = (await import('vue')).defineComponent({
    setup() {
        return {
            settings: settings,
            User: User,
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