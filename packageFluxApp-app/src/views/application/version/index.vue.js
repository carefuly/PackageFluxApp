import { useRoute, useRouter } from "vue-router";
import { skyMsgBox, skyMsgError, skyMsgSuccess, skyMsgWarning, skyNoticeError, skyNoticeSuccess } from "@/utils/sky";
import { getById as getByIdApp } from "@/apis/application/details";
import { listAll as fileListAll } from "@/apis/application/file";
import { add, deleteById, updateById, setFormal, listAll, getById, } from "@/apis/application/version";
import { Delete, Edit, Position, Search, Warning } from "@element-plus/icons-vue";
const route = useRoute();
const router = useRouter();
const skyDialogRef = ref();
const formRef = ref();
const pageData = ref({
    id: null,
    searchValue: "",
    details: {
        logoUrl: "",
        appName: "",
    },
    pageParams: {
        detailId: null,
    },
    title: "应用版本",
    confirmLoading: false,
    form: {
        id: null,
        versionCode: "",
        description: "",
        HBuilderUpdate: false,
        forceUpdate: false,
        formalVersion: false,
        apkUrl: "",
        wgtUrl: "",
        detail_id: null,
        remark: null,
    },
    rules: {
        versionCode: [{ required: true, message: "请输入版本号", trigger: "blur" }],
        apkUrl: [{ required: true, message: "请输入apkUrl地址", trigger: "blur" }],
        wgtUrl: [{ required: true, message: "请输入wgtUrl地址", trigger: "blur" }],
    },
    loading: false,
    total: 0,
    fileList: [],
    tableList: [],
});
const method = reactive({
    /** 获取id参数 */
    handleGetAppId: () => {
        pageData.value.id = route.path;
        return pageData.value.id.split('/')[pageData.value.id.split('/').length - 1];
    },
    /** 获取应用详情 */
    handleGetByIdApp: async (id) => {
        try {
            const res = await getByIdApp(id);
            pageData.value.details = res.data;
        }
        catch (error) {
            skyMsgError("数据查询失败，请刷新重试🌻");
        }
    },
    /** 重置 */
    resetSearch: () => {
        method.resetSearchParams();
        method.handleListPage();
    },
    /** 重置搜索参数 */
    resetSearchParams: () => {
        pageData.value.pageParams = {
            detailId: null,
        };
        pageData.value.pageParams.detailId = method.handleGetAppId();
    },
    /** 添加 */
    handleAdd: () => {
        // 打开弹出框
        skyDialogRef.value.skyOpen();
        // 重置表单
        method.resetForm();
        // 获取id
        pageData.value.form.detail_id = method.handleGetAppId();
        // 标题
        pageData.value.title = "添加";
        // 文件数据表格
        method.handleFileListAll();
    },
    /** 设置正式版 */
    handleFormal: async (row) => {
        try {
            const res = await setFormal(row.id, pageData.value.pageParams.detailId);
            await method.handleListPage();
            skyMsgSuccess("设置成功🌻");
        }
        catch (error) {
            skyMsgError(error);
        }
    },
    /** 删除 */
    handleDelete: (row) => {
        const id = row.id;
        if (id === null || id === "") {
            skyMsgWarning("请选中需要删除的数据🌻");
            return;
        }
        skyMsgBox("您确认需要删除名称[" + row.versionCode + "]么？")
            .then(async () => {
            try {
                await deleteById(id, pageData.value.pageParams.detailId);
                await method.handleListPage();
                skyNoticeSuccess("删除成功🌻");
            }
            catch (error) {
                await method.handleListPage();
                skyNoticeError("删除失败，请刷新重试🌻");
            }
        })
            .catch(() => {
            skyMsgError("已取消🌻");
        });
    },
    /** 修改 */
    handleUpdate: (row) => {
        // 打开弹出框
        skyDialogRef.value.skyOpen();
        // 文件数据表格
        method.handleFileListAll();
        // 重置表单
        method.resetForm();
        pageData.value.title = "修改";
        const id = row.id;
        if (id == null || id === "") {
            skyMsgError("请选中需要修改的数据🌻");
        }
        // 回显数据
        method.handleEcho(id);
    },
    /** 回显数据 */
    handleEcho: async (id) => {
        if (id === null || id === "") {
            skyMsgWarning("请选择需要修改的数据🌻");
            return;
        }
        try {
            const res = await getById(id, pageData.value.pageParams);
            pageData.value.form = res.data;
        }
        catch (error) {
            skyNoticeError("数据获取失败，请刷新重试🌻");
        }
    },
    /** 确定  */
    handleConfirm: () => {
        if (!formRef.value)
            return;
        pageData.value.confirmLoading = true;
        (formRef.value).validate(async (valid) => {
            if (valid) {
                if (pageData.value.form.id) {
                    try {
                        await updateById(pageData.value.form.id, pageData.value.form);
                        skyNoticeSuccess("修改成功🌻");
                        pageData.value.confirmLoading = false;
                        skyDialogRef.value.skyQuickClose();
                        method.resetForm();
                        await method.handleListPage();
                    }
                    catch (error) {
                        pageData.value.confirmLoading = false;
                        skyNoticeError("修改失败，请刷新重试🌻");
                    }
                }
                else {
                    try {
                        await add(pageData.value.form);
                        skyNoticeSuccess("添加成功🌻");
                        pageData.value.confirmLoading = false;
                        skyDialogRef.value.skyQuickClose();
                        method.resetForm();
                        await method.handleListPage();
                    }
                    catch (error) {
                        pageData.value.confirmLoading = false;
                        skyNoticeError("添加失败，请刷新重试🌻");
                    }
                }
            }
            else {
                skyMsgError("验证失败，请检查填写内容🌻");
                pageData.value.confirmLoading = false;
            }
        });
    },
    /** 取消 */
    handleCancel: () => {
        skyDialogRef.value.skyClose();
    },
    /** 清空表单数据 */
    resetForm: () => {
        nextTick(() => {
            if (formRef.value) {
                // 重置该表单项，将其值重置为初始值，并移除校验结果
                formRef.value.resetFields();
            }
        });
        pageData.value.form = {
            id: null,
            versionCode: "",
            description: "",
            HBuilderUpdate: false,
            forceUpdate: false,
            formalVersion: false,
            apkUrl: "",
            wgtUrl: "",
            detail_id: null,
            remark: null,
        };
    },
    /** 数据表格 */
    handleListPage: async () => {
        try {
            pageData.value.loading = true;
            pageData.value.tableList = [];
            const res = await listAll(pageData.value.pageParams);
            pageData.value.tableList = res.data;
            pageData.value.loading = false;
        }
        catch (error) {
            skyMsgError("数据查询失败，请刷新重试🌻");
        }
    },
    /** 文件数据表格 */
    handleFileListAll: async () => {
        try {
            pageData.value.fileList = [];
            const res = await fileListAll({});
            pageData.value.fileList = res.data;
        }
        catch (error) {
            skyMsgError("数据查询失败，请刷新重试🌻");
        }
    },
});
const filterList = computed(() => {
    if (pageData.value.searchValue) {
        return pageData.value.tableList.filter(item => item.versionCode.includes(pageData.value.searchValue));
    }
    else {
        return pageData.value.tableList;
    }
});
onMounted(() => {
    method.handleGetByIdApp(method.handleGetAppId());
    pageData.value.pageParams.detailId = method.handleGetAppId();
    method.handleListPage();
});
; /* PartiallyEnd: #3632/scriptSetup.vue */
const __VLS_ctx = {};
let __VLS_components;
let __VLS_directives;
// CSS variable injection 
// CSS variable injection end 
__VLS_asFunctionalElement(__VLS_intrinsicElements.div, __VLS_intrinsicElements.div)({
    ...{ class: ("sky-flex") },
});
const __VLS_0 = {}.SkyCard;
/** @type { [typeof __VLS_components.SkyCard, typeof __VLS_components.SkyCard, ] } */ ;
// @ts-ignore
const __VLS_1 = __VLS_asFunctionalComponent(__VLS_0, new __VLS_0({}));
const __VLS_2 = __VLS_1({}, ...__VLS_functionalComponentArgsRest(__VLS_1));
const __VLS_6 = {}.ElRow;
/** @type { [typeof __VLS_components.ElRow, typeof __VLS_components.elRow, typeof __VLS_components.ElRow, typeof __VLS_components.elRow, ] } */ ;
// @ts-ignore
const __VLS_7 = __VLS_asFunctionalComponent(__VLS_6, new __VLS_6({
    gutter: ((10)),
    ...{ style: ({}) },
}));
const __VLS_8 = __VLS_7({
    gutter: ((10)),
    ...{ style: ({}) },
}, ...__VLS_functionalComponentArgsRest(__VLS_7));
__VLS_asFunctionalElement(__VLS_intrinsicElements.div, __VLS_intrinsicElements.div)({
    ...{ style: ({}) },
});
const __VLS_12 = {}.ElCol;
/** @type { [typeof __VLS_components.ElCol, typeof __VLS_components.elCol, typeof __VLS_components.ElCol, typeof __VLS_components.elCol, ] } */ ;
// @ts-ignore
const __VLS_13 = __VLS_asFunctionalComponent(__VLS_12, new __VLS_12({
    span: ((1.5)),
}));
const __VLS_14 = __VLS_13({
    span: ((1.5)),
}, ...__VLS_functionalComponentArgsRest(__VLS_13));
const __VLS_18 = {}.ElTooltip;
/** @type { [typeof __VLS_components.ElTooltip, typeof __VLS_components.elTooltip, typeof __VLS_components.ElTooltip, typeof __VLS_components.elTooltip, ] } */ ;
// @ts-ignore
const __VLS_19 = __VLS_asFunctionalComponent(__VLS_18, new __VLS_18({
    content: ("返回"),
}));
const __VLS_20 = __VLS_19({
    content: ("返回"),
}, ...__VLS_functionalComponentArgsRest(__VLS_19));
const __VLS_24 = {}.ElButton;
/** @type { [typeof __VLS_components.ElButton, typeof __VLS_components.elButton, typeof __VLS_components.ElButton, typeof __VLS_components.elButton, ] } */ ;
// @ts-ignore
const __VLS_25 = __VLS_asFunctionalComponent(__VLS_24, new __VLS_24({
    ...{ 'onClick': {} },
    size: ("small"),
    type: ("primary"),
    icon: ("ArrowLeftBold"),
    circle: (true),
    plain: (true),
}));
const __VLS_26 = __VLS_25({
    ...{ 'onClick': {} },
    size: ("small"),
    type: ("primary"),
    icon: ("ArrowLeftBold"),
    circle: (true),
    plain: (true),
}, ...__VLS_functionalComponentArgsRest(__VLS_25));
let __VLS_30;
const __VLS_31 = {
    onClick: (...[$event]) => {
        __VLS_ctx.router.push('/application');
    }
};
let __VLS_27;
let __VLS_28;
var __VLS_29;
__VLS_23.slots.default;
var __VLS_23;
__VLS_17.slots.default;
var __VLS_17;
const __VLS_32 = {}.ElCol;
/** @type { [typeof __VLS_components.ElCol, typeof __VLS_components.elCol, typeof __VLS_components.ElCol, typeof __VLS_components.elCol, ] } */ ;
// @ts-ignore
const __VLS_33 = __VLS_asFunctionalComponent(__VLS_32, new __VLS_32({
    span: ((1.5)),
}));
const __VLS_34 = __VLS_33({
    span: ((1.5)),
}, ...__VLS_functionalComponentArgsRest(__VLS_33));
__VLS_asFunctionalElement(__VLS_intrinsicElements.img, __VLS_intrinsicElements.img)({
    ...{ style: ({}) },
    src: ((__VLS_ctx.pageData.details.logoUrl)),
    alt: ((__VLS_ctx.pageData.details.appName)),
});
__VLS_37.slots.default;
var __VLS_37;
const __VLS_38 = {}.ElCol;
/** @type { [typeof __VLS_components.ElCol, typeof __VLS_components.elCol, typeof __VLS_components.ElCol, typeof __VLS_components.elCol, ] } */ ;
// @ts-ignore
const __VLS_39 = __VLS_asFunctionalComponent(__VLS_38, new __VLS_38({
    span: ((1.5)),
}));
const __VLS_40 = __VLS_39({
    span: ((1.5)),
}, ...__VLS_functionalComponentArgsRest(__VLS_39));
(__VLS_ctx.pageData.details.appName);
__VLS_asFunctionalElement(__VLS_intrinsicElements.span, __VLS_intrinsicElements.span)({
    ...{ style: ({}) },
});
__VLS_43.slots.default;
var __VLS_43;
__VLS_asFunctionalElement(__VLS_intrinsicElements.div, __VLS_intrinsicElements.div)({});
const __VLS_44 = {}.ElCol;
/** @type { [typeof __VLS_components.ElCol, typeof __VLS_components.elCol, typeof __VLS_components.ElCol, typeof __VLS_components.elCol, ] } */ ;
// @ts-ignore
const __VLS_45 = __VLS_asFunctionalComponent(__VLS_44, new __VLS_44({
    span: ((1.5)),
}));
const __VLS_46 = __VLS_45({
    span: ((1.5)),
}, ...__VLS_functionalComponentArgsRest(__VLS_45));
const __VLS_50 = {}.ElButton;
/** @type { [typeof __VLS_components.ElButton, typeof __VLS_components.elButton, typeof __VLS_components.ElButton, typeof __VLS_components.elButton, ] } */ ;
// @ts-ignore
const __VLS_51 = __VLS_asFunctionalComponent(__VLS_50, new __VLS_50({
    ...{ 'onClick': {} },
    type: ("primary"),
    icon: ("plus"),
    plain: (true),
}));
const __VLS_52 = __VLS_51({
    ...{ 'onClick': {} },
    type: ("primary"),
    icon: ("plus"),
    plain: (true),
}, ...__VLS_functionalComponentArgsRest(__VLS_51));
let __VLS_56;
const __VLS_57 = {
    onClick: (...[$event]) => {
        __VLS_ctx.method.handleAdd();
    }
};
let __VLS_53;
let __VLS_54;
__VLS_55.slots.default;
var __VLS_55;
const __VLS_58 = {}.ElButton;
/** @type { [typeof __VLS_components.ElButton, typeof __VLS_components.elButton, typeof __VLS_components.ElButton, typeof __VLS_components.elButton, ] } */ ;
// @ts-ignore
const __VLS_59 = __VLS_asFunctionalComponent(__VLS_58, new __VLS_58({
    ...{ 'onClick': {} },
    type: ("danger"),
    icon: ("refresh"),
    plain: (true),
}));
const __VLS_60 = __VLS_59({
    ...{ 'onClick': {} },
    type: ("danger"),
    icon: ("refresh"),
    plain: (true),
}, ...__VLS_functionalComponentArgsRest(__VLS_59));
let __VLS_64;
const __VLS_65 = {
    onClick: (__VLS_ctx.method.resetSearch)
};
let __VLS_61;
let __VLS_62;
__VLS_63.slots.default;
var __VLS_63;
__VLS_49.slots.default;
var __VLS_49;
const __VLS_66 = {}.ElCol;
/** @type { [typeof __VLS_components.ElCol, typeof __VLS_components.elCol, typeof __VLS_components.ElCol, typeof __VLS_components.elCol, ] } */ ;
// @ts-ignore
const __VLS_67 = __VLS_asFunctionalComponent(__VLS_66, new __VLS_66({
    span: ((1.5)),
}));
const __VLS_68 = __VLS_67({
    span: ((1.5)),
}, ...__VLS_functionalComponentArgsRest(__VLS_67));
const __VLS_72 = {}.ElInput;
/** @type { [typeof __VLS_components.ElInput, typeof __VLS_components.elInput, typeof __VLS_components.ElInput, typeof __VLS_components.elInput, ] } */ ;
// @ts-ignore
const __VLS_73 = __VLS_asFunctionalComponent(__VLS_72, new __VLS_72({
    ...{ style: ({}) },
    placeholder: ("请输入应用版本号【例如：2.2.8】"),
    prefixIcon: ((__VLS_ctx.Search)),
    modelValue: ((__VLS_ctx.pageData.searchValue)),
    clearable: (true),
}));
const __VLS_74 = __VLS_73({
    ...{ style: ({}) },
    placeholder: ("请输入应用版本号【例如：2.2.8】"),
    prefixIcon: ((__VLS_ctx.Search)),
    modelValue: ((__VLS_ctx.pageData.searchValue)),
    clearable: (true),
}, ...__VLS_functionalComponentArgsRest(__VLS_73));
__VLS_71.slots.default;
var __VLS_71;
__VLS_11.slots.default;
var __VLS_11;
__VLS_asFunctionalElement(__VLS_intrinsicElements.br)({});
if (__VLS_ctx.filterList.length > 0) {
    const __VLS_78 = {}.ElTimeline;
    /** @type { [typeof __VLS_components.ElTimeline, typeof __VLS_components.elTimeline, typeof __VLS_components.ElTimeline, typeof __VLS_components.elTimeline, ] } */ ;
    // @ts-ignore
    const __VLS_79 = __VLS_asFunctionalComponent(__VLS_78, new __VLS_78({}));
    const __VLS_80 = __VLS_79({}, ...__VLS_functionalComponentArgsRest(__VLS_79));
    for (const [item] of __VLS_getVForSourceType((__VLS_ctx.filterList))) {
        const __VLS_84 = {}.ElTimelineItem;
        /** @type { [typeof __VLS_components.ElTimelineItem, typeof __VLS_components.elTimelineItem, typeof __VLS_components.ElTimelineItem, typeof __VLS_components.elTimelineItem, ] } */ ;
        // @ts-ignore
        const __VLS_85 = __VLS_asFunctionalComponent(__VLS_84, new __VLS_84({
            key: ((item.id)),
            type: ((item.formalVersion ? 'success' : 'info')),
        }));
        const __VLS_86 = __VLS_85({
            key: ((item.id)),
            type: ((item.formalVersion ? 'success' : 'info')),
        }, ...__VLS_functionalComponentArgsRest(__VLS_85));
        __VLS_asFunctionalElement(__VLS_intrinsicElements.div, __VLS_intrinsicElements.div)({
            ...{ class: ("title") },
        });
        const __VLS_90 = {}.ElText;
        /** @type { [typeof __VLS_components.ElText, typeof __VLS_components.elText, typeof __VLS_components.ElText, typeof __VLS_components.elText, ] } */ ;
        // @ts-ignore
        const __VLS_91 = __VLS_asFunctionalComponent(__VLS_90, new __VLS_90({
            ...{ class: ("mr-2") },
            size: ("large"),
            type: ("primary"),
        }));
        const __VLS_92 = __VLS_91({
            ...{ class: ("mr-2") },
            size: ("large"),
            type: ("primary"),
        }, ...__VLS_functionalComponentArgsRest(__VLS_91));
        (item.versionCode);
        __VLS_95.slots.default;
        var __VLS_95;
        if (!item.formalVersion) {
            const __VLS_96 = {}.ElPopconfirm;
            /** @type { [typeof __VLS_components.ElPopconfirm, typeof __VLS_components.elPopconfirm, typeof __VLS_components.ElPopconfirm, typeof __VLS_components.elPopconfirm, ] } */ ;
            // @ts-ignore
            const __VLS_97 = __VLS_asFunctionalComponent(__VLS_96, new __VLS_96({
                ...{ 'onConfirm': {} },
                width: ("220"),
                title: ("是否将当前版本设为正式版"),
                type: ("warning"),
                okText: ("确认"),
                cancelText: ("取消"),
            }));
            const __VLS_98 = __VLS_97({
                ...{ 'onConfirm': {} },
                width: ("220"),
                title: ("是否将当前版本设为正式版"),
                type: ("warning"),
                okText: ("确认"),
                cancelText: ("取消"),
            }, ...__VLS_functionalComponentArgsRest(__VLS_97));
            let __VLS_102;
            const __VLS_103 = {
                onConfirm: (...[$event]) => {
                    if (!((__VLS_ctx.filterList.length > 0)))
                        return;
                    if (!((!item.formalVersion)))
                        return;
                    __VLS_ctx.method.handleFormal(item);
                }
            };
            let __VLS_99;
            let __VLS_100;
            {
                const { reference: __VLS_thisSlot } = __VLS_101.slots;
                const __VLS_104 = {}.ElIcon;
                /** @type { [typeof __VLS_components.ElIcon, typeof __VLS_components.elIcon, typeof __VLS_components.ElIcon, typeof __VLS_components.elIcon, ] } */ ;
                // @ts-ignore
                const __VLS_105 = __VLS_asFunctionalComponent(__VLS_104, new __VLS_104({
                    ...{ class: ("mr-2") },
                }));
                const __VLS_106 = __VLS_105({
                    ...{ class: ("mr-2") },
                }, ...__VLS_functionalComponentArgsRest(__VLS_105));
                const __VLS_110 = {}.Position;
                /** @type { [typeof __VLS_components.Position, ] } */ ;
                // @ts-ignore
                const __VLS_111 = __VLS_asFunctionalComponent(__VLS_110, new __VLS_110({}));
                const __VLS_112 = __VLS_111({}, ...__VLS_functionalComponentArgsRest(__VLS_111));
                __VLS_109.slots.default;
                var __VLS_109;
            }
            __VLS_101.slots.default;
            var __VLS_101;
        }
        const __VLS_116 = {}.ElTooltip;
        /** @type { [typeof __VLS_components.ElTooltip, typeof __VLS_components.elTooltip, typeof __VLS_components.ElTooltip, typeof __VLS_components.elTooltip, ] } */ ;
        // @ts-ignore
        const __VLS_117 = __VLS_asFunctionalComponent(__VLS_116, new __VLS_116({
            content: ("编辑"),
        }));
        const __VLS_118 = __VLS_117({
            content: ("编辑"),
        }, ...__VLS_functionalComponentArgsRest(__VLS_117));
        const __VLS_122 = {}.ElIcon;
        /** @type { [typeof __VLS_components.ElIcon, typeof __VLS_components.elIcon, typeof __VLS_components.ElIcon, typeof __VLS_components.elIcon, ] } */ ;
        // @ts-ignore
        const __VLS_123 = __VLS_asFunctionalComponent(__VLS_122, new __VLS_122({
            ...{ 'onClick': {} },
            ...{ class: ("mr-2") },
        }));
        const __VLS_124 = __VLS_123({
            ...{ 'onClick': {} },
            ...{ class: ("mr-2") },
        }, ...__VLS_functionalComponentArgsRest(__VLS_123));
        let __VLS_128;
        const __VLS_129 = {
            onClick: (...[$event]) => {
                if (!((__VLS_ctx.filterList.length > 0)))
                    return;
                __VLS_ctx.method.handleUpdate(item);
            }
        };
        let __VLS_125;
        let __VLS_126;
        const __VLS_130 = {}.Edit;
        /** @type { [typeof __VLS_components.Edit, ] } */ ;
        // @ts-ignore
        const __VLS_131 = __VLS_asFunctionalComponent(__VLS_130, new __VLS_130({}));
        const __VLS_132 = __VLS_131({}, ...__VLS_functionalComponentArgsRest(__VLS_131));
        __VLS_127.slots.default;
        var __VLS_127;
        __VLS_121.slots.default;
        var __VLS_121;
        if (!item.formalVersion) {
            const __VLS_136 = {}.ElPopconfirm;
            /** @type { [typeof __VLS_components.ElPopconfirm, typeof __VLS_components.elPopconfirm, typeof __VLS_components.ElPopconfirm, typeof __VLS_components.elPopconfirm, ] } */ ;
            // @ts-ignore
            const __VLS_137 = __VLS_asFunctionalComponent(__VLS_136, new __VLS_136({
                ...{ 'onConfirm': {} },
                title: ("是否删除当前版本"),
                width: ("180"),
            }));
            const __VLS_138 = __VLS_137({
                ...{ 'onConfirm': {} },
                title: ("是否删除当前版本"),
                width: ("180"),
            }, ...__VLS_functionalComponentArgsRest(__VLS_137));
            let __VLS_142;
            const __VLS_143 = {
                onConfirm: (...[$event]) => {
                    if (!((__VLS_ctx.filterList.length > 0)))
                        return;
                    if (!((!item.formalVersion)))
                        return;
                    __VLS_ctx.method.handleDelete(item);
                }
            };
            let __VLS_139;
            let __VLS_140;
            {
                const { reference: __VLS_thisSlot } = __VLS_141.slots;
                const __VLS_144 = {}.ElIcon;
                /** @type { [typeof __VLS_components.ElIcon, typeof __VLS_components.elIcon, typeof __VLS_components.ElIcon, typeof __VLS_components.elIcon, ] } */ ;
                // @ts-ignore
                const __VLS_145 = __VLS_asFunctionalComponent(__VLS_144, new __VLS_144({}));
                const __VLS_146 = __VLS_145({}, ...__VLS_functionalComponentArgsRest(__VLS_145));
                const __VLS_150 = {}.Delete;
                /** @type { [typeof __VLS_components.Delete, ] } */ ;
                // @ts-ignore
                const __VLS_151 = __VLS_asFunctionalComponent(__VLS_150, new __VLS_150({}));
                const __VLS_152 = __VLS_151({}, ...__VLS_functionalComponentArgsRest(__VLS_151));
                __VLS_149.slots.default;
                var __VLS_149;
            }
            __VLS_141.slots.default;
            var __VLS_141;
        }
        __VLS_asFunctionalElement(__VLS_intrinsicElements.div, __VLS_intrinsicElements.div)({
            ...{ class: ("mt-2 mb-2") },
        });
        if (item.description) {
            for (const [desc] of __VLS_getVForSourceType((item.description.split('\n')))) {
                const __VLS_156 = {}.ElText;
                /** @type { [typeof __VLS_components.ElText, typeof __VLS_components.elText, typeof __VLS_components.ElText, typeof __VLS_components.elText, ] } */ ;
                // @ts-ignore
                const __VLS_157 = __VLS_asFunctionalComponent(__VLS_156, new __VLS_156({
                    ...{ class: ("mx-1") },
                    key: ((desc)),
                }));
                const __VLS_158 = __VLS_157({
                    ...{ class: ("mx-1") },
                    key: ((desc)),
                }, ...__VLS_functionalComponentArgsRest(__VLS_157));
                const __VLS_162 = {}.ElTooltip;
                /** @type { [typeof __VLS_components.ElTooltip, typeof __VLS_components.elTooltip, typeof __VLS_components.ElTooltip, typeof __VLS_components.elTooltip, ] } */ ;
                // @ts-ignore
                const __VLS_163 = __VLS_asFunctionalComponent(__VLS_162, new __VLS_162({
                    content: ((desc)),
                }));
                const __VLS_164 = __VLS_163({
                    content: ((desc)),
                }, ...__VLS_functionalComponentArgsRest(__VLS_163));
                (desc);
                __VLS_167.slots.default;
                var __VLS_167;
                __VLS_asFunctionalElement(__VLS_intrinsicElements.br)({});
                __VLS_161.slots.default;
                var __VLS_161;
            }
        }
        const __VLS_168 = {}.ElText;
        /** @type { [typeof __VLS_components.ElText, typeof __VLS_components.elText, typeof __VLS_components.ElText, typeof __VLS_components.elText, ] } */ ;
        // @ts-ignore
        const __VLS_169 = __VLS_asFunctionalComponent(__VLS_168, new __VLS_168({
            ...{ class: ("mx-1") },
            type: ("success"),
        }));
        const __VLS_170 = __VLS_169({
            ...{ class: ("mx-1") },
            type: ("success"),
        }, ...__VLS_functionalComponentArgsRest(__VLS_169));
        const __VLS_174 = {}.ElText;
        /** @type { [typeof __VLS_components.ElText, typeof __VLS_components.elText, typeof __VLS_components.ElText, typeof __VLS_components.elText, ] } */ ;
        // @ts-ignore
        const __VLS_175 = __VLS_asFunctionalComponent(__VLS_174, new __VLS_174({
            ...{ class: ("mx-1") },
        }));
        const __VLS_176 = __VLS_175({
            ...{ class: ("mx-1") },
        }, ...__VLS_functionalComponentArgsRest(__VLS_175));
        (item.apkUrl);
        __VLS_179.slots.default;
        var __VLS_179;
        __VLS_173.slots.default;
        var __VLS_173;
        __VLS_asFunctionalElement(__VLS_intrinsicElements.div, __VLS_intrinsicElements.div)({
            ...{ class: ("h-0.5em") },
        });
        const __VLS_180 = {}.ElText;
        /** @type { [typeof __VLS_components.ElText, typeof __VLS_components.elText, typeof __VLS_components.ElText, typeof __VLS_components.elText, ] } */ ;
        // @ts-ignore
        const __VLS_181 = __VLS_asFunctionalComponent(__VLS_180, new __VLS_180({
            ...{ class: ("mx-1") },
            type: ("warning"),
        }));
        const __VLS_182 = __VLS_181({
            ...{ class: ("mx-1") },
            type: ("warning"),
        }, ...__VLS_functionalComponentArgsRest(__VLS_181));
        const __VLS_186 = {}.ElText;
        /** @type { [typeof __VLS_components.ElText, typeof __VLS_components.elText, typeof __VLS_components.ElText, typeof __VLS_components.elText, ] } */ ;
        // @ts-ignore
        const __VLS_187 = __VLS_asFunctionalComponent(__VLS_186, new __VLS_186({
            ...{ class: ("mx-1") },
        }));
        const __VLS_188 = __VLS_187({
            ...{ class: ("mx-1") },
        }, ...__VLS_functionalComponentArgsRest(__VLS_187));
        (item.wgtUrl);
        __VLS_191.slots.default;
        var __VLS_191;
        __VLS_185.slots.default;
        var __VLS_185;
        __VLS_asFunctionalElement(__VLS_intrinsicElements.div, __VLS_intrinsicElements.div)({
            ...{ class: ("h-0.5em") },
        });
        if (item.HBuilderUpdate) {
            const __VLS_192 = {}.ElTag;
            /** @type { [typeof __VLS_components.ElTag, typeof __VLS_components.elTag, typeof __VLS_components.ElTag, typeof __VLS_components.elTag, ] } */ ;
            // @ts-ignore
            const __VLS_193 = __VLS_asFunctionalComponent(__VLS_192, new __VLS_192({
                ...{ class: ("mr-2") },
                type: ("primary"),
            }));
            const __VLS_194 = __VLS_193({
                ...{ class: ("mr-2") },
                type: ("primary"),
            }, ...__VLS_functionalComponentArgsRest(__VLS_193));
            __VLS_197.slots.default;
            var __VLS_197;
        }
        if (item.forceUpdate) {
            const __VLS_198 = {}.ElTag;
            /** @type { [typeof __VLS_components.ElTag, typeof __VLS_components.elTag, typeof __VLS_components.ElTag, typeof __VLS_components.elTag, ] } */ ;
            // @ts-ignore
            const __VLS_199 = __VLS_asFunctionalComponent(__VLS_198, new __VLS_198({
                type: ("success"),
            }));
            const __VLS_200 = __VLS_199({
                type: ("success"),
            }, ...__VLS_functionalComponentArgsRest(__VLS_199));
            __VLS_203.slots.default;
            var __VLS_203;
        }
        __VLS_asFunctionalElement(__VLS_intrinsicElements.div, __VLS_intrinsicElements.div)({
            ...{ class: ("h-0.5em") },
        });
        if (item.formalVersion) {
            const __VLS_204 = {}.ElTag;
            /** @type { [typeof __VLS_components.ElTag, typeof __VLS_components.elTag, typeof __VLS_components.ElTag, typeof __VLS_components.elTag, ] } */ ;
            // @ts-ignore
            const __VLS_205 = __VLS_asFunctionalComponent(__VLS_204, new __VLS_204({
                type: ("danger"),
            }));
            const __VLS_206 = __VLS_205({
                type: ("danger"),
            }, ...__VLS_functionalComponentArgsRest(__VLS_205));
            __VLS_209.slots.default;
            var __VLS_209;
        }
        __VLS_asFunctionalElement(__VLS_intrinsicElements.div, __VLS_intrinsicElements.div)({
            ...{ class: ("h-0.5em") },
        });
        const __VLS_210 = {}.ElText;
        /** @type { [typeof __VLS_components.ElText, typeof __VLS_components.elText, typeof __VLS_components.ElText, typeof __VLS_components.elText, ] } */ ;
        // @ts-ignore
        const __VLS_211 = __VLS_asFunctionalComponent(__VLS_210, new __VLS_210({
            ...{ class: ("mx-1") },
            type: ("danger"),
        }));
        const __VLS_212 = __VLS_211({
            ...{ class: ("mx-1") },
            type: ("danger"),
        }, ...__VLS_functionalComponentArgsRest(__VLS_211));
        const __VLS_216 = {}.ElText;
        /** @type { [typeof __VLS_components.ElText, typeof __VLS_components.elText, typeof __VLS_components.ElText, typeof __VLS_components.elText, ] } */ ;
        // @ts-ignore
        const __VLS_217 = __VLS_asFunctionalComponent(__VLS_216, new __VLS_216({
            ...{ class: ("mx-1") },
        }));
        const __VLS_218 = __VLS_217({
            ...{ class: ("mx-1") },
        }, ...__VLS_functionalComponentArgsRest(__VLS_217));
        (item.updateTime);
        __VLS_221.slots.default;
        var __VLS_221;
        __VLS_215.slots.default;
        var __VLS_215;
        __VLS_89.slots.default;
        var __VLS_89;
    }
    __VLS_83.slots.default;
    var __VLS_83;
}
else {
    __VLS_asFunctionalElement(__VLS_intrinsicElements.div, __VLS_intrinsicElements.div)({
        ...{ class: ("message center") },
    });
    __VLS_asFunctionalElement(__VLS_intrinsicElements.h1, __VLS_intrinsicElements.h1)({});
    __VLS_asFunctionalElement(__VLS_intrinsicElements.p, __VLS_intrinsicElements.p)({});
}
const __VLS_222 = {}.SkyDialog;
/** @type { [typeof __VLS_components.SkyDialog, typeof __VLS_components.SkyDialog, ] } */ ;
// @ts-ignore
const __VLS_223 = __VLS_asFunctionalComponent(__VLS_222, new __VLS_222({
    ...{ 'onSkyConfirm': {} },
    ...{ 'onSkyCancel': {} },
    ref: ("skyDialogRef"),
    title: ((__VLS_ctx.pageData.title)),
    loading: ((__VLS_ctx.pageData.confirmLoading)),
    width: ((750)),
    height: ((450)),
}));
const __VLS_224 = __VLS_223({
    ...{ 'onSkyConfirm': {} },
    ...{ 'onSkyCancel': {} },
    ref: ("skyDialogRef"),
    title: ((__VLS_ctx.pageData.title)),
    loading: ((__VLS_ctx.pageData.confirmLoading)),
    width: ((750)),
    height: ((450)),
}, ...__VLS_functionalComponentArgsRest(__VLS_223));
// @ts-ignore navigation for `const skyDialogRef = ref()`
/** @type { typeof __VLS_ctx.skyDialogRef } */ ;
var __VLS_228 = {};
let __VLS_229;
const __VLS_230 = {
    onSkyConfirm: (__VLS_ctx.method.handleConfirm)
};
const __VLS_231 = {
    onSkyCancel: (__VLS_ctx.method.handleCancel)
};
let __VLS_225;
let __VLS_226;
{
    const { content: __VLS_thisSlot } = __VLS_227.slots;
    const __VLS_232 = {}.ElForm;
    /** @type { [typeof __VLS_components.ElForm, typeof __VLS_components.elForm, typeof __VLS_components.ElForm, typeof __VLS_components.elForm, ] } */ ;
    // @ts-ignore
    const __VLS_233 = __VLS_asFunctionalComponent(__VLS_232, new __VLS_232({
        ref: ("formRef"),
        model: ((__VLS_ctx.pageData.form)),
        rules: ((__VLS_ctx.pageData.rules)),
        labelWidth: ("auto"),
    }));
    const __VLS_234 = __VLS_233({
        ref: ("formRef"),
        model: ((__VLS_ctx.pageData.form)),
        rules: ((__VLS_ctx.pageData.rules)),
        labelWidth: ("auto"),
    }, ...__VLS_functionalComponentArgsRest(__VLS_233));
    // @ts-ignore navigation for `const formRef = ref()`
    /** @type { typeof __VLS_ctx.formRef } */ ;
    var __VLS_238 = {};
    const __VLS_239 = {}.ElRow;
    /** @type { [typeof __VLS_components.ElRow, typeof __VLS_components.elRow, typeof __VLS_components.ElRow, typeof __VLS_components.elRow, ] } */ ;
    // @ts-ignore
    const __VLS_240 = __VLS_asFunctionalComponent(__VLS_239, new __VLS_239({}));
    const __VLS_241 = __VLS_240({}, ...__VLS_functionalComponentArgsRest(__VLS_240));
    const __VLS_245 = {}.ElCol;
    /** @type { [typeof __VLS_components.ElCol, typeof __VLS_components.elCol, typeof __VLS_components.ElCol, typeof __VLS_components.elCol, ] } */ ;
    // @ts-ignore
    const __VLS_246 = __VLS_asFunctionalComponent(__VLS_245, new __VLS_245({
        xs: (({ span: 24 })),
        sm: (({ span: 24 })),
    }));
    const __VLS_247 = __VLS_246({
        xs: (({ span: 24 })),
        sm: (({ span: 24 })),
    }, ...__VLS_functionalComponentArgsRest(__VLS_246));
    const __VLS_251 = {}.ElFormItem;
    /** @type { [typeof __VLS_components.ElFormItem, typeof __VLS_components.elFormItem, typeof __VLS_components.ElFormItem, typeof __VLS_components.elFormItem, ] } */ ;
    // @ts-ignore
    const __VLS_252 = __VLS_asFunctionalComponent(__VLS_251, new __VLS_251({
        label: ("应用版本号"),
        prop: ("versionCode"),
    }));
    const __VLS_253 = __VLS_252({
        label: ("应用版本号"),
        prop: ("versionCode"),
    }, ...__VLS_functionalComponentArgsRest(__VLS_252));
    const __VLS_257 = {}.ElInput;
    /** @type { [typeof __VLS_components.ElInput, typeof __VLS_components.elInput, ] } */ ;
    // @ts-ignore
    const __VLS_258 = __VLS_asFunctionalComponent(__VLS_257, new __VLS_257({
        modelValue: ((__VLS_ctx.pageData.form.versionCode)),
        placeholder: ("请输入应用版本号，例如：2.2.8"),
        disabled: ((__VLS_ctx.pageData.form.id !== null)),
        clearable: (true),
    }));
    const __VLS_259 = __VLS_258({
        modelValue: ((__VLS_ctx.pageData.form.versionCode)),
        placeholder: ("请输入应用版本号，例如：2.2.8"),
        disabled: ((__VLS_ctx.pageData.form.id !== null)),
        clearable: (true),
    }, ...__VLS_functionalComponentArgsRest(__VLS_258));
    __VLS_256.slots.default;
    var __VLS_256;
    __VLS_250.slots.default;
    var __VLS_250;
    __VLS_244.slots.default;
    var __VLS_244;
    const __VLS_263 = {}.ElRow;
    /** @type { [typeof __VLS_components.ElRow, typeof __VLS_components.elRow, typeof __VLS_components.ElRow, typeof __VLS_components.elRow, ] } */ ;
    // @ts-ignore
    const __VLS_264 = __VLS_asFunctionalComponent(__VLS_263, new __VLS_263({}));
    const __VLS_265 = __VLS_264({}, ...__VLS_functionalComponentArgsRest(__VLS_264));
    const __VLS_269 = {}.ElCol;
    /** @type { [typeof __VLS_components.ElCol, typeof __VLS_components.elCol, typeof __VLS_components.ElCol, typeof __VLS_components.elCol, ] } */ ;
    // @ts-ignore
    const __VLS_270 = __VLS_asFunctionalComponent(__VLS_269, new __VLS_269({
        xs: (({ span: 24 })),
        sm: (({ span: 24 })),
    }));
    const __VLS_271 = __VLS_270({
        xs: (({ span: 24 })),
        sm: (({ span: 24 })),
    }, ...__VLS_functionalComponentArgsRest(__VLS_270));
    const __VLS_275 = {}.ElFormItem;
    /** @type { [typeof __VLS_components.ElFormItem, typeof __VLS_components.elFormItem, typeof __VLS_components.ElFormItem, typeof __VLS_components.elFormItem, ] } */ ;
    // @ts-ignore
    const __VLS_276 = __VLS_asFunctionalComponent(__VLS_275, new __VLS_275({
        label: ("更新说明"),
        prop: ("description"),
    }));
    const __VLS_277 = __VLS_276({
        label: ("更新说明"),
        prop: ("description"),
    }, ...__VLS_functionalComponentArgsRest(__VLS_276));
    const __VLS_281 = {}.ElInput;
    /** @type { [typeof __VLS_components.ElInput, typeof __VLS_components.elInput, ] } */ ;
    // @ts-ignore
    const __VLS_282 = __VLS_asFunctionalComponent(__VLS_281, new __VLS_281({
        modelValue: ((__VLS_ctx.pageData.form.description)),
        rows: ((3)),
        type: ("textarea"),
        placeholder: ("\u4f8b\u5982\uff1a\u0031\u002e\u4fee\u590d\u4e86\u90e8\u5206\u0062\u0075\u0067\u000a\u0020\u0020\u0020\u0020\u0020\u0020\u0020\u0020\u0020\u0020\u0032\u002e\u589e\u52a0\u4e86\u67d0\u67d0\u529f\u80fd"),
        clearable: (true),
    }));
    const __VLS_283 = __VLS_282({
        modelValue: ((__VLS_ctx.pageData.form.description)),
        rows: ((3)),
        type: ("textarea"),
        placeholder: ("\u4f8b\u5982\uff1a\u0031\u002e\u4fee\u590d\u4e86\u90e8\u5206\u0062\u0075\u0067\u000a\u0020\u0020\u0020\u0020\u0020\u0020\u0020\u0020\u0020\u0020\u0032\u002e\u589e\u52a0\u4e86\u67d0\u67d0\u529f\u80fd"),
        clearable: (true),
    }, ...__VLS_functionalComponentArgsRest(__VLS_282));
    __VLS_280.slots.default;
    var __VLS_280;
    __VLS_274.slots.default;
    var __VLS_274;
    __VLS_268.slots.default;
    var __VLS_268;
    const __VLS_287 = {}.ElRow;
    /** @type { [typeof __VLS_components.ElRow, typeof __VLS_components.elRow, typeof __VLS_components.ElRow, typeof __VLS_components.elRow, ] } */ ;
    // @ts-ignore
    const __VLS_288 = __VLS_asFunctionalComponent(__VLS_287, new __VLS_287({}));
    const __VLS_289 = __VLS_288({}, ...__VLS_functionalComponentArgsRest(__VLS_288));
    const __VLS_293 = {}.ElCol;
    /** @type { [typeof __VLS_components.ElCol, typeof __VLS_components.elCol, typeof __VLS_components.ElCol, typeof __VLS_components.elCol, ] } */ ;
    // @ts-ignore
    const __VLS_294 = __VLS_asFunctionalComponent(__VLS_293, new __VLS_293({
        xs: (({ span: 24 })),
        sm: (({ span: 24 })),
    }));
    const __VLS_295 = __VLS_294({
        xs: (({ span: 24 })),
        sm: (({ span: 24 })),
    }, ...__VLS_functionalComponentArgsRest(__VLS_294));
    const __VLS_299 = {}.ElFormItem;
    /** @type { [typeof __VLS_components.ElFormItem, typeof __VLS_components.elFormItem, typeof __VLS_components.ElFormItem, typeof __VLS_components.elFormItem, ] } */ ;
    // @ts-ignore
    const __VLS_300 = __VLS_asFunctionalComponent(__VLS_299, new __VLS_299({
        label: ("HBuilderX是否更新"),
    }));
    const __VLS_301 = __VLS_300({
        label: ("HBuilderX是否更新"),
    }, ...__VLS_functionalComponentArgsRest(__VLS_300));
    {
        const { label: __VLS_thisSlot } = __VLS_304.slots;
        __VLS_asFunctionalElement(__VLS_intrinsicElements.div, __VLS_intrinsicElements.div)({
            ...{ style: ({}) },
        });
        __VLS_asFunctionalElement(__VLS_intrinsicElements.span, __VLS_intrinsicElements.span)({
            ...{ style: ({}) },
        });
        const __VLS_305 = {}.ElTooltip;
        /** @type { [typeof __VLS_components.ElTooltip, typeof __VLS_components.elTooltip, typeof __VLS_components.ElTooltip, typeof __VLS_components.elTooltip, ] } */ ;
        // @ts-ignore
        const __VLS_306 = __VLS_asFunctionalComponent(__VLS_305, new __VLS_305({}));
        const __VLS_307 = __VLS_306({}, ...__VLS_functionalComponentArgsRest(__VLS_306));
        {
            const { content: __VLS_thisSlot } = __VLS_310.slots;
        }
        const __VLS_311 = {}.ElIcon;
        /** @type { [typeof __VLS_components.ElIcon, typeof __VLS_components.elIcon, typeof __VLS_components.ElIcon, typeof __VLS_components.elIcon, ] } */ ;
        // @ts-ignore
        const __VLS_312 = __VLS_asFunctionalComponent(__VLS_311, new __VLS_311({}));
        const __VLS_313 = __VLS_312({}, ...__VLS_functionalComponentArgsRest(__VLS_312));
        const __VLS_317 = {}.Warning;
        /** @type { [typeof __VLS_components.Warning, ] } */ ;
        // @ts-ignore
        const __VLS_318 = __VLS_asFunctionalComponent(__VLS_317, new __VLS_317({}));
        const __VLS_319 = __VLS_318({}, ...__VLS_functionalComponentArgsRest(__VLS_318));
        __VLS_316.slots.default;
        var __VLS_316;
        __VLS_310.slots.default;
        var __VLS_310;
    }
    const __VLS_323 = {}.ElSwitch;
    /** @type { [typeof __VLS_components.ElSwitch, typeof __VLS_components.elSwitch, ] } */ ;
    // @ts-ignore
    const __VLS_324 = __VLS_asFunctionalComponent(__VLS_323, new __VLS_323({
        modelValue: ((__VLS_ctx.pageData.form.HBuilderUpdate)),
        inlinePrompt: (true),
        activeText: ("是"),
        inactiveText: ("否"),
    }));
    const __VLS_325 = __VLS_324({
        modelValue: ((__VLS_ctx.pageData.form.HBuilderUpdate)),
        inlinePrompt: (true),
        activeText: ("是"),
        inactiveText: ("否"),
    }, ...__VLS_functionalComponentArgsRest(__VLS_324));
    __VLS_304.slots.default;
    var __VLS_304;
    __VLS_298.slots.default;
    var __VLS_298;
    __VLS_292.slots.default;
    var __VLS_292;
    const __VLS_329 = {}.ElRow;
    /** @type { [typeof __VLS_components.ElRow, typeof __VLS_components.elRow, typeof __VLS_components.ElRow, typeof __VLS_components.elRow, ] } */ ;
    // @ts-ignore
    const __VLS_330 = __VLS_asFunctionalComponent(__VLS_329, new __VLS_329({}));
    const __VLS_331 = __VLS_330({}, ...__VLS_functionalComponentArgsRest(__VLS_330));
    const __VLS_335 = {}.ElCol;
    /** @type { [typeof __VLS_components.ElCol, typeof __VLS_components.elCol, typeof __VLS_components.ElCol, typeof __VLS_components.elCol, ] } */ ;
    // @ts-ignore
    const __VLS_336 = __VLS_asFunctionalComponent(__VLS_335, new __VLS_335({
        xs: (({ span: 24 })),
        sm: (({ span: 24 })),
    }));
    const __VLS_337 = __VLS_336({
        xs: (({ span: 24 })),
        sm: (({ span: 24 })),
    }, ...__VLS_functionalComponentArgsRest(__VLS_336));
    const __VLS_341 = {}.ElFormItem;
    /** @type { [typeof __VLS_components.ElFormItem, typeof __VLS_components.elFormItem, typeof __VLS_components.ElFormItem, typeof __VLS_components.elFormItem, ] } */ ;
    // @ts-ignore
    const __VLS_342 = __VLS_asFunctionalComponent(__VLS_341, new __VLS_341({
        label: ("是否强制更新"),
    }));
    const __VLS_343 = __VLS_342({
        label: ("是否强制更新"),
    }, ...__VLS_functionalComponentArgsRest(__VLS_342));
    const __VLS_347 = {}.ElSwitch;
    /** @type { [typeof __VLS_components.ElSwitch, typeof __VLS_components.elSwitch, ] } */ ;
    // @ts-ignore
    const __VLS_348 = __VLS_asFunctionalComponent(__VLS_347, new __VLS_347({
        modelValue: ((__VLS_ctx.pageData.form.forceUpdate)),
        inlinePrompt: (true),
        activeText: ("是"),
        inactiveText: ("否"),
    }));
    const __VLS_349 = __VLS_348({
        modelValue: ((__VLS_ctx.pageData.form.forceUpdate)),
        inlinePrompt: (true),
        activeText: ("是"),
        inactiveText: ("否"),
    }, ...__VLS_functionalComponentArgsRest(__VLS_348));
    {
        const { label: __VLS_thisSlot } = __VLS_346.slots;
        __VLS_asFunctionalElement(__VLS_intrinsicElements.div, __VLS_intrinsicElements.div)({
            ...{ style: ({}) },
        });
        __VLS_asFunctionalElement(__VLS_intrinsicElements.span, __VLS_intrinsicElements.span)({
            ...{ style: ({}) },
        });
        const __VLS_353 = {}.ElTooltip;
        /** @type { [typeof __VLS_components.ElTooltip, typeof __VLS_components.elTooltip, typeof __VLS_components.ElTooltip, typeof __VLS_components.elTooltip, ] } */ ;
        // @ts-ignore
        const __VLS_354 = __VLS_asFunctionalComponent(__VLS_353, new __VLS_353({}));
        const __VLS_355 = __VLS_354({}, ...__VLS_functionalComponentArgsRest(__VLS_354));
        {
            const { content: __VLS_thisSlot } = __VLS_358.slots;
        }
        const __VLS_359 = {}.ElIcon;
        /** @type { [typeof __VLS_components.ElIcon, typeof __VLS_components.elIcon, typeof __VLS_components.ElIcon, typeof __VLS_components.elIcon, ] } */ ;
        // @ts-ignore
        const __VLS_360 = __VLS_asFunctionalComponent(__VLS_359, new __VLS_359({}));
        const __VLS_361 = __VLS_360({}, ...__VLS_functionalComponentArgsRest(__VLS_360));
        const __VLS_365 = {}.Warning;
        /** @type { [typeof __VLS_components.Warning, ] } */ ;
        // @ts-ignore
        const __VLS_366 = __VLS_asFunctionalComponent(__VLS_365, new __VLS_365({}));
        const __VLS_367 = __VLS_366({}, ...__VLS_functionalComponentArgsRest(__VLS_366));
        __VLS_364.slots.default;
        var __VLS_364;
        __VLS_358.slots.default;
        var __VLS_358;
    }
    __VLS_346.slots.default;
    var __VLS_346;
    __VLS_340.slots.default;
    var __VLS_340;
    __VLS_334.slots.default;
    var __VLS_334;
    const __VLS_371 = {}.ElRow;
    /** @type { [typeof __VLS_components.ElRow, typeof __VLS_components.elRow, typeof __VLS_components.ElRow, typeof __VLS_components.elRow, ] } */ ;
    // @ts-ignore
    const __VLS_372 = __VLS_asFunctionalComponent(__VLS_371, new __VLS_371({}));
    const __VLS_373 = __VLS_372({}, ...__VLS_functionalComponentArgsRest(__VLS_372));
    const __VLS_377 = {}.ElCol;
    /** @type { [typeof __VLS_components.ElCol, typeof __VLS_components.elCol, typeof __VLS_components.ElCol, typeof __VLS_components.elCol, ] } */ ;
    // @ts-ignore
    const __VLS_378 = __VLS_asFunctionalComponent(__VLS_377, new __VLS_377({
        xs: (({ span: 24 })),
        sm: (({ span: 24 })),
    }));
    const __VLS_379 = __VLS_378({
        xs: (({ span: 24 })),
        sm: (({ span: 24 })),
    }, ...__VLS_functionalComponentArgsRest(__VLS_378));
    const __VLS_383 = {}.ElFormItem;
    /** @type { [typeof __VLS_components.ElFormItem, typeof __VLS_components.elFormItem, typeof __VLS_components.ElFormItem, typeof __VLS_components.elFormItem, ] } */ ;
    // @ts-ignore
    const __VLS_384 = __VLS_asFunctionalComponent(__VLS_383, new __VLS_383({
        label: ("apk下载地址"),
        prop: ("apkUrl"),
    }));
    const __VLS_385 = __VLS_384({
        label: ("apk下载地址"),
        prop: ("apkUrl"),
    }, ...__VLS_functionalComponentArgsRest(__VLS_384));
    const __VLS_389 = {}.ElSelect;
    /** @type { [typeof __VLS_components.ElSelect, typeof __VLS_components.elSelect, typeof __VLS_components.ElSelect, typeof __VLS_components.elSelect, ] } */ ;
    // @ts-ignore
    const __VLS_390 = __VLS_asFunctionalComponent(__VLS_389, new __VLS_389({
        modelValue: ((__VLS_ctx.pageData.form.apkUrl)),
        filterable: (true),
        allowCreate: (true),
        defaultFirstOption: (true),
        reserveKeyword: ((false)),
        placeholder: ("请输入或选择apk下载地址"),
        clearable: (true),
    }));
    const __VLS_391 = __VLS_390({
        modelValue: ((__VLS_ctx.pageData.form.apkUrl)),
        filterable: (true),
        allowCreate: (true),
        defaultFirstOption: (true),
        reserveKeyword: ((false)),
        placeholder: ("请输入或选择apk下载地址"),
        clearable: (true),
    }, ...__VLS_functionalComponentArgsRest(__VLS_390));
    for (const [item] of __VLS_getVForSourceType((__VLS_ctx.pageData.fileList))) {
        const __VLS_395 = {}.ElOption;
        /** @type { [typeof __VLS_components.ElOption, typeof __VLS_components.elOption, ] } */ ;
        // @ts-ignore
        const __VLS_396 = __VLS_asFunctionalComponent(__VLS_395, new __VLS_395({
            key: ((item.id)),
            label: ((item.name + `.${item.suffix}`)),
            value: ((item.url)),
        }));
        const __VLS_397 = __VLS_396({
            key: ((item.id)),
            label: ((item.name + `.${item.suffix}`)),
            value: ((item.url)),
        }, ...__VLS_functionalComponentArgsRest(__VLS_396));
    }
    __VLS_394.slots.default;
    var __VLS_394;
    __VLS_388.slots.default;
    var __VLS_388;
    __VLS_382.slots.default;
    var __VLS_382;
    __VLS_376.slots.default;
    var __VLS_376;
    const __VLS_401 = {}.ElRow;
    /** @type { [typeof __VLS_components.ElRow, typeof __VLS_components.elRow, typeof __VLS_components.ElRow, typeof __VLS_components.elRow, ] } */ ;
    // @ts-ignore
    const __VLS_402 = __VLS_asFunctionalComponent(__VLS_401, new __VLS_401({}));
    const __VLS_403 = __VLS_402({}, ...__VLS_functionalComponentArgsRest(__VLS_402));
    const __VLS_407 = {}.ElCol;
    /** @type { [typeof __VLS_components.ElCol, typeof __VLS_components.elCol, typeof __VLS_components.ElCol, typeof __VLS_components.elCol, ] } */ ;
    // @ts-ignore
    const __VLS_408 = __VLS_asFunctionalComponent(__VLS_407, new __VLS_407({
        xs: (({ span: 24 })),
        sm: (({ span: 24 })),
    }));
    const __VLS_409 = __VLS_408({
        xs: (({ span: 24 })),
        sm: (({ span: 24 })),
    }, ...__VLS_functionalComponentArgsRest(__VLS_408));
    const __VLS_413 = {}.ElFormItem;
    /** @type { [typeof __VLS_components.ElFormItem, typeof __VLS_components.elFormItem, typeof __VLS_components.ElFormItem, typeof __VLS_components.elFormItem, ] } */ ;
    // @ts-ignore
    const __VLS_414 = __VLS_asFunctionalComponent(__VLS_413, new __VLS_413({
        label: ("wgt下载地址"),
        prop: ("wgtUrl"),
    }));
    const __VLS_415 = __VLS_414({
        label: ("wgt下载地址"),
        prop: ("wgtUrl"),
    }, ...__VLS_functionalComponentArgsRest(__VLS_414));
    const __VLS_419 = {}.ElSelect;
    /** @type { [typeof __VLS_components.ElSelect, typeof __VLS_components.elSelect, typeof __VLS_components.ElSelect, typeof __VLS_components.elSelect, ] } */ ;
    // @ts-ignore
    const __VLS_420 = __VLS_asFunctionalComponent(__VLS_419, new __VLS_419({
        modelValue: ((__VLS_ctx.pageData.form.wgtUrl)),
        filterable: (true),
        allowCreate: (true),
        defaultFirstOption: (true),
        reserveKeyword: ((false)),
        placeholder: ("请输入或选择wgt下载地址"),
        clearable: (true),
    }));
    const __VLS_421 = __VLS_420({
        modelValue: ((__VLS_ctx.pageData.form.wgtUrl)),
        filterable: (true),
        allowCreate: (true),
        defaultFirstOption: (true),
        reserveKeyword: ((false)),
        placeholder: ("请输入或选择wgt下载地址"),
        clearable: (true),
    }, ...__VLS_functionalComponentArgsRest(__VLS_420));
    for (const [item] of __VLS_getVForSourceType((__VLS_ctx.pageData.fileList))) {
        const __VLS_425 = {}.ElOption;
        /** @type { [typeof __VLS_components.ElOption, typeof __VLS_components.elOption, ] } */ ;
        // @ts-ignore
        const __VLS_426 = __VLS_asFunctionalComponent(__VLS_425, new __VLS_425({
            key: ((item.id)),
            label: ((item.name + `.${item.suffix}`)),
            value: ((item.url)),
        }));
        const __VLS_427 = __VLS_426({
            key: ((item.id)),
            label: ((item.name + `.${item.suffix}`)),
            value: ((item.url)),
        }, ...__VLS_functionalComponentArgsRest(__VLS_426));
    }
    __VLS_424.slots.default;
    var __VLS_424;
    __VLS_418.slots.default;
    var __VLS_418;
    __VLS_412.slots.default;
    var __VLS_412;
    __VLS_406.slots.default;
    var __VLS_406;
    const __VLS_431 = {}.ElRow;
    /** @type { [typeof __VLS_components.ElRow, typeof __VLS_components.elRow, typeof __VLS_components.ElRow, typeof __VLS_components.elRow, ] } */ ;
    // @ts-ignore
    const __VLS_432 = __VLS_asFunctionalComponent(__VLS_431, new __VLS_431({}));
    const __VLS_433 = __VLS_432({}, ...__VLS_functionalComponentArgsRest(__VLS_432));
    const __VLS_437 = {}.ElCol;
    /** @type { [typeof __VLS_components.ElCol, typeof __VLS_components.elCol, typeof __VLS_components.ElCol, typeof __VLS_components.elCol, ] } */ ;
    // @ts-ignore
    const __VLS_438 = __VLS_asFunctionalComponent(__VLS_437, new __VLS_437({
        xs: (({ span: 24 })),
        sm: (({ span: 24 })),
    }));
    const __VLS_439 = __VLS_438({
        xs: (({ span: 24 })),
        sm: (({ span: 24 })),
    }, ...__VLS_functionalComponentArgsRest(__VLS_438));
    const __VLS_443 = {}.ElFormItem;
    /** @type { [typeof __VLS_components.ElFormItem, typeof __VLS_components.elFormItem, typeof __VLS_components.ElFormItem, typeof __VLS_components.elFormItem, ] } */ ;
    // @ts-ignore
    const __VLS_444 = __VLS_asFunctionalComponent(__VLS_443, new __VLS_443({
        label: ("备注"),
        prop: ("remark"),
    }));
    const __VLS_445 = __VLS_444({
        label: ("备注"),
        prop: ("remark"),
    }, ...__VLS_functionalComponentArgsRest(__VLS_444));
    const __VLS_449 = {}.ElInput;
    /** @type { [typeof __VLS_components.ElInput, typeof __VLS_components.elInput, ] } */ ;
    // @ts-ignore
    const __VLS_450 = __VLS_asFunctionalComponent(__VLS_449, new __VLS_449({
        modelValue: ((__VLS_ctx.pageData.form.remark)),
        rows: ((3)),
        type: ("textarea"),
        placeholder: ("请输入备注"),
    }));
    const __VLS_451 = __VLS_450({
        modelValue: ((__VLS_ctx.pageData.form.remark)),
        rows: ((3)),
        type: ("textarea"),
        placeholder: ("请输入备注"),
    }, ...__VLS_functionalComponentArgsRest(__VLS_450));
    __VLS_448.slots.default;
    var __VLS_448;
    __VLS_442.slots.default;
    var __VLS_442;
    __VLS_436.slots.default;
    var __VLS_436;
    __VLS_237.slots.default;
    var __VLS_237;
}
__VLS_227.slots.default;
var __VLS_227;
__VLS_5.slots.default;
var __VLS_5;
['sky-flex', 'title', 'mr-2', 'mr-2', 'mr-2', 'mt-2', 'mb-2', 'mx-1', 'mx-1', 'mx-1', 'h-0.5em', 'mx-1', 'mx-1', 'h-0.5em', 'mr-2', 'h-0.5em', 'h-0.5em', 'mx-1', 'mx-1', 'message', 'center',];
var __VLS_special;
const __VLS_self = (await import('vue')).defineComponent({
    setup() {
        return {
            Delete: Delete,
            Edit: Edit,
            Position: Position,
            Search: Search,
            Warning: Warning,
            router: router,
            skyDialogRef: skyDialogRef,
            formRef: formRef,
            pageData: pageData,
            method: method,
            filterList: filterList,
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