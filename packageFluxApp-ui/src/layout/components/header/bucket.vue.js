/// <reference types=".vue-global-types/vue_3.5_false.d.ts" />
import { skyMsgBox, skyMsgError, skyMsgInfo, skyMsgWarning, skyNoticeError, skyNoticeSuccess } from "@/utils/sky";
import { deleteById, batchDelete, updateById, listPage, getById } from "@/apis/application/file";
const { defineProps, defineSlots, defineEmits, defineExpose, defineModel, defineOptions, withDefaults, } = await import('vue');
const skyDialogRef = ref();
const formRef = ref();
const pageData = ref({
    title: "文件存储桶",
    pageParams: {
        creator: null,
        modifier: null,
        status: true,
        name: "",
        user_id: null,
        page: 1,
        pageSize: 10,
    },
    ids: [],
    multiple: false,
    visible: false,
    form: {
        id: null,
        name: "",
        size: 0,
        suffix: "",
        path: "",
        url: "",
        sort: 1,
        status: true,
        belong_dept: null,
        remark: null,
    },
    rules: {
        name: [{ required: true, message: "请输入文件名称", trigger: "blur" }],
    },
    loading: false,
    confirmLoading: false,
    total: 0,
    tableList: [],
});
const method = reactive({
    /** 打开 */
    fileOpen: () => {
        pageData.value.visible = true;
        method.handleListPage();
    },
    /** 删除 */
    handleDelete: (row) => {
        const id = row.id;
        if (id === null || id === "") {
            skyMsgWarning("请选中需要删除的数据🌻");
            return;
        }
        skyMsgBox("您确认需要删除名称[" + row.name + "]么？")
            .then(async () => {
            try {
                await deleteById(id);
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
    /** 批量删除 */
    handleBatchDelete: () => {
        if (pageData.value.ids.length == 0) {
            skyMsgInfo("请选择需要删除的数据🌻");
            return;
        }
        skyMsgBox("您确认需要进行批量删除么？")
            .then(async () => {
            try {
                const res = await batchDelete(pageData.value.ids);
                await method.handleListPage();
                skyNoticeSuccess(`${res.msg}🌻`);
            }
            catch (error) {
                await method.handleListPage();
                skyNoticeSuccess("批量删除失败，请刷新重试🌻");
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
        // 重置表单
        method.resetForm();
        pageData.value.title = "修改";
        const id = row ? row.id : pageData.value.ids[0];
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
            const res = await getById(id);
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
    /** 确定 */
    fileConfirm: () => {
    },
    /** 取消 */
    fileClose: () => {
        skyMsgBox("您确认进行关闭么？")
            .then(() => {
            pageData.value.visible = false;
            skyMsgWarning("已关闭🌻");
        })
            .catch(() => {
            // 用户点击了取消按钮或关闭确认框
            // 执行取消操作或不做任何操作
            skyMsgWarning("已取消🌻");
        });
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
            name: "",
            size: 0,
            suffix: "",
            path: "",
            url: "",
            sort: 1,
            status: true,
            belong_dept: null,
            remark: null,
        };
    },
    /** 是否多选 */
    handleSelectionChange: (selection) => {
        pageData.value.ids = selection.map((item) => item.id);
        pageData.value.multiple = !selection.length; // 多选
    },
    /** 数据表格 */
    handleListPage: async () => {
        try {
            pageData.value.loading = true;
            pageData.value.tableList = [];
            const res = await listPage(pageData.value.pageParams);
            pageData.value.tableList = res.data.items;
            pageData.value.total = res.data.total;
            pageData.value.loading = false;
        }
        catch (error) {
            skyMsgError("数据查询失败，请刷新重试🌻");
        }
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
    __VLS_elementAsFunction(__VLS_intrinsicElements.div, __VLS_intrinsicElements.div)({ ...{ class: ("bucket") }, });
    const __VLS_0 = __VLS_resolvedLocalAndGlobalComponents.ElButton;
    /** @type { [typeof __VLS_components.ElButton, typeof __VLS_components.elButton, typeof __VLS_components.ElButton, typeof __VLS_components.elButton, ] } */
    // @ts-ignore
    const __VLS_1 = __VLS_asFunctionalComponent(__VLS_0, new __VLS_0({ ...{ 'onClick': {} }, type: ("primary"), icon: ("TakeawayBox"), link: (true), }));
    const __VLS_2 = __VLS_1({ ...{ 'onClick': {} }, type: ("primary"), icon: ("TakeawayBox"), link: (true), }, ...__VLS_functionalComponentArgsRest(__VLS_1));
    let __VLS_6;
    const __VLS_7 = {
        onClick: (__VLS_ctx.method.fileOpen)
    };
    let __VLS_3;
    let __VLS_4;
    __VLS_nonNullable(__VLS_5.slots).default;
    const __VLS_5 = __VLS_pickFunctionalComponentCtx(__VLS_0, __VLS_2);
    const __VLS_8 = __VLS_resolvedLocalAndGlobalComponents.ElDialog;
    /** @type { [typeof __VLS_components.ElDialog, typeof __VLS_components.elDialog, typeof __VLS_components.ElDialog, typeof __VLS_components.elDialog, ] } */
    // @ts-ignore
    const __VLS_9 = __VLS_asFunctionalComponent(__VLS_8, new __VLS_8({ modelValue: ((__VLS_ctx.pageData.visible)), title: ((__VLS_ctx.pageData.title)), closeOnClickModal: ((false)), appendToBody: (true), draggable: (true), destroyOnClose: ((true)), beforeClose: ((__VLS_ctx.method.fileClose)), fullscreen: ((true)), }));
    const __VLS_10 = __VLS_9({ modelValue: ((__VLS_ctx.pageData.visible)), title: ((__VLS_ctx.pageData.title)), closeOnClickModal: ((false)), appendToBody: (true), draggable: (true), destroyOnClose: ((true)), beforeClose: ((__VLS_ctx.method.fileClose)), fullscreen: ((true)), }, ...__VLS_functionalComponentArgsRest(__VLS_9));
    __VLS_elementAsFunction(__VLS_intrinsicElements.div, __VLS_intrinsicElements.div)({ ...{ style: ({}) }, ...{ class: ("sky-flex") }, });
    const __VLS_14 = __VLS_resolvedLocalAndGlobalComponents.SkyCard;
    /** @type { [typeof __VLS_components.SkyCard, typeof __VLS_components.SkyCard, ] } */
    // @ts-ignore
    const __VLS_15 = __VLS_asFunctionalComponent(__VLS_14, new __VLS_14({}));
    const __VLS_16 = __VLS_15({}, ...__VLS_functionalComponentArgsRest(__VLS_15));
    const __VLS_20 = __VLS_resolvedLocalAndGlobalComponents.ElRow;
    /** @type { [typeof __VLS_components.ElRow, typeof __VLS_components.elRow, typeof __VLS_components.ElRow, typeof __VLS_components.elRow, ] } */
    // @ts-ignore
    const __VLS_21 = __VLS_asFunctionalComponent(__VLS_20, new __VLS_20({ gutter: ((10)), }));
    const __VLS_22 = __VLS_21({ gutter: ((10)), }, ...__VLS_functionalComponentArgsRest(__VLS_21));
    const __VLS_26 = __VLS_resolvedLocalAndGlobalComponents.ElCol;
    /** @type { [typeof __VLS_components.ElCol, typeof __VLS_components.elCol, typeof __VLS_components.ElCol, typeof __VLS_components.elCol, ] } */
    // @ts-ignore
    const __VLS_27 = __VLS_asFunctionalComponent(__VLS_26, new __VLS_26({ span: ((1.5)), }));
    const __VLS_28 = __VLS_27({ span: ((1.5)), }, ...__VLS_functionalComponentArgsRest(__VLS_27));
    const __VLS_32 = __VLS_resolvedLocalAndGlobalComponents.ElButton;
    /** @type { [typeof __VLS_components.ElButton, typeof __VLS_components.elButton, typeof __VLS_components.ElButton, typeof __VLS_components.elButton, ] } */
    // @ts-ignore
    const __VLS_33 = __VLS_asFunctionalComponent(__VLS_32, new __VLS_32({ ...{ 'onClick': {} }, type: ("danger"), icon: ("delete"), plain: (true), disabled: ((__VLS_ctx.pageData.multiple)), }));
    const __VLS_34 = __VLS_33({ ...{ 'onClick': {} }, type: ("danger"), icon: ("delete"), plain: (true), disabled: ((__VLS_ctx.pageData.multiple)), }, ...__VLS_functionalComponentArgsRest(__VLS_33));
    let __VLS_38;
    const __VLS_39 = {
        onClick: (...[$event]) => {
            __VLS_ctx.method.handleBatchDelete();
        }
    };
    let __VLS_35;
    let __VLS_36;
    __VLS_nonNullable(__VLS_37.slots).default;
    const __VLS_37 = __VLS_pickFunctionalComponentCtx(__VLS_32, __VLS_34);
    __VLS_nonNullable(__VLS_31.slots).default;
    const __VLS_31 = __VLS_pickFunctionalComponentCtx(__VLS_26, __VLS_28);
    __VLS_nonNullable(__VLS_25.slots).default;
    const __VLS_25 = __VLS_pickFunctionalComponentCtx(__VLS_20, __VLS_22);
    __VLS_elementAsFunction(__VLS_intrinsicElements.br)({});
    const __VLS_40 = __VLS_resolvedLocalAndGlobalComponents.ElTable;
    /** @type { [typeof __VLS_components.ElTable, typeof __VLS_components.elTable, typeof __VLS_components.ElTable, typeof __VLS_components.elTable, ] } */
    // @ts-ignore
    const __VLS_41 = __VLS_asFunctionalComponent(__VLS_40, new __VLS_40({ ...{ 'onSelectionChange': {} }, border: (true), data: ((__VLS_ctx.pageData.tableList)), emptyText: ("暂时没有数据哟🌻"), }));
    const __VLS_42 = __VLS_41({ ...{ 'onSelectionChange': {} }, border: (true), data: ((__VLS_ctx.pageData.tableList)), emptyText: ("暂时没有数据哟🌻"), }, ...__VLS_functionalComponentArgsRest(__VLS_41));
    __VLS_directiveAsFunction(__VLS_directives.vLoading)(null, { ...__VLS_directiveBindingRestFields, value: (__VLS_ctx.pageData.loading) }, null, null);
    let __VLS_46;
    const __VLS_47 = {
        onSelectionChange: (__VLS_ctx.method.handleSelectionChange)
    };
    let __VLS_43;
    let __VLS_44;
    const __VLS_48 = __VLS_resolvedLocalAndGlobalComponents.ElTableColumn;
    /** @type { [typeof __VLS_components.ElTableColumn, typeof __VLS_components.elTableColumn, ] } */
    // @ts-ignore
    const __VLS_49 = __VLS_asFunctionalComponent(__VLS_48, new __VLS_48({ type: ("selection"), width: ("55"), align: ("center"), fixed: ("left"), }));
    const __VLS_50 = __VLS_49({ type: ("selection"), width: ("55"), align: ("center"), fixed: ("left"), }, ...__VLS_functionalComponentArgsRest(__VLS_49));
    const __VLS_54 = __VLS_resolvedLocalAndGlobalComponents.ElTableColumn;
    /** @type { [typeof __VLS_components.ElTableColumn, typeof __VLS_components.elTableColumn, typeof __VLS_components.ElTableColumn, typeof __VLS_components.elTableColumn, ] } */
    // @ts-ignore
    const __VLS_55 = __VLS_asFunctionalComponent(__VLS_54, new __VLS_54({ label: ("文件标识"), prop: ("id"), width: ("110"), align: ("center"), showOverflowTooltip: ((true)), }));
    const __VLS_56 = __VLS_55({ label: ("文件标识"), prop: ("id"), width: ("110"), align: ("center"), showOverflowTooltip: ((true)), }, ...__VLS_functionalComponentArgsRest(__VLS_55));
    const __VLS_60 = __VLS_resolvedLocalAndGlobalComponents.ElTableColumn;
    /** @type { [typeof __VLS_components.ElTableColumn, typeof __VLS_components.elTableColumn, typeof __VLS_components.ElTableColumn, typeof __VLS_components.elTableColumn, ] } */
    // @ts-ignore
    const __VLS_61 = __VLS_asFunctionalComponent(__VLS_60, new __VLS_60({ label: ("文件名称"), prop: ("name"), width: ("180"), align: ("center"), showOverflowTooltip: ((true)), }));
    const __VLS_62 = __VLS_61({ label: ("文件名称"), prop: ("name"), width: ("180"), align: ("center"), showOverflowTooltip: ((true)), }, ...__VLS_functionalComponentArgsRest(__VLS_61));
    const __VLS_66 = __VLS_resolvedLocalAndGlobalComponents.ElTableColumn;
    /** @type { [typeof __VLS_components.ElTableColumn, typeof __VLS_components.elTableColumn, typeof __VLS_components.ElTableColumn, typeof __VLS_components.elTableColumn, ] } */
    // @ts-ignore
    const __VLS_67 = __VLS_asFunctionalComponent(__VLS_66, new __VLS_66({ label: ("文件大小(KB)"), prop: ("size"), width: ("160"), align: ("center"), showOverflowTooltip: ((true)), }));
    const __VLS_68 = __VLS_67({ label: ("文件大小(KB)"), prop: ("size"), width: ("160"), align: ("center"), showOverflowTooltip: ((true)), }, ...__VLS_functionalComponentArgsRest(__VLS_67));
    const __VLS_72 = __VLS_resolvedLocalAndGlobalComponents.ElTableColumn;
    /** @type { [typeof __VLS_components.ElTableColumn, typeof __VLS_components.elTableColumn, typeof __VLS_components.ElTableColumn, typeof __VLS_components.elTableColumn, ] } */
    // @ts-ignore
    const __VLS_73 = __VLS_asFunctionalComponent(__VLS_72, new __VLS_72({ label: ("文件后缀(类型)"), prop: ("suffix"), width: ("160"), align: ("center"), showOverflowTooltip: ((true)), }));
    const __VLS_74 = __VLS_73({ label: ("文件后缀(类型)"), prop: ("suffix"), width: ("160"), align: ("center"), showOverflowTooltip: ((true)), }, ...__VLS_functionalComponentArgsRest(__VLS_73));
    const __VLS_78 = __VLS_resolvedLocalAndGlobalComponents.ElTableColumn;
    /** @type { [typeof __VLS_components.ElTableColumn, typeof __VLS_components.elTableColumn, typeof __VLS_components.ElTableColumn, typeof __VLS_components.elTableColumn, ] } */
    // @ts-ignore
    const __VLS_79 = __VLS_asFunctionalComponent(__VLS_78, new __VLS_78({ label: ("文件地址"), prop: ("path"), width: ("480"), align: ("center"), showOverflowTooltip: ((true)), }));
    const __VLS_80 = __VLS_79({ label: ("文件地址"), prop: ("path"), width: ("480"), align: ("center"), showOverflowTooltip: ((true)), }, ...__VLS_functionalComponentArgsRest(__VLS_79));
    const __VLS_84 = __VLS_resolvedLocalAndGlobalComponents.ElTableColumn;
    /** @type { [typeof __VLS_components.ElTableColumn, typeof __VLS_components.elTableColumn, typeof __VLS_components.ElTableColumn, typeof __VLS_components.elTableColumn, ] } */
    // @ts-ignore
    const __VLS_85 = __VLS_asFunctionalComponent(__VLS_84, new __VLS_84({ label: ("完整URL"), prop: ("url"), align: ("center"), showOverflowTooltip: ((true)), }));
    const __VLS_86 = __VLS_85({ label: ("完整URL"), prop: ("url"), align: ("center"), showOverflowTooltip: ((true)), }, ...__VLS_functionalComponentArgsRest(__VLS_85));
    const __VLS_90 = __VLS_resolvedLocalAndGlobalComponents.ElTableColumn;
    /** @type { [typeof __VLS_components.ElTableColumn, typeof __VLS_components.elTableColumn, typeof __VLS_components.ElTableColumn, typeof __VLS_components.elTableColumn, ] } */
    // @ts-ignore
    const __VLS_91 = __VLS_asFunctionalComponent(__VLS_90, new __VLS_90({ label: ("操作"), align: ("center"), width: ("120"), fixed: ("right"), }));
    const __VLS_92 = __VLS_91({ label: ("操作"), align: ("center"), width: ("120"), fixed: ("right"), }, ...__VLS_functionalComponentArgsRest(__VLS_91));
    __VLS_elementAsFunction(__VLS_intrinsicElements.template, __VLS_intrinsicElements.template)({});
    {
        const { default: __VLS_thisSlot } = __VLS_nonNullable(__VLS_95.slots);
        const [{ row }] = __VLS_getSlotParams(__VLS_thisSlot);
        const __VLS_96 = __VLS_resolvedLocalAndGlobalComponents.ElTooltip;
        /** @type { [typeof __VLS_components.ElTooltip, typeof __VLS_components.elTooltip, typeof __VLS_components.ElTooltip, typeof __VLS_components.elTooltip, ] } */
        // @ts-ignore
        const __VLS_97 = __VLS_asFunctionalComponent(__VLS_96, new __VLS_96({ content: ("修改🌻"), placement: ("top"), }));
        const __VLS_98 = __VLS_97({ content: ("修改🌻"), placement: ("top"), }, ...__VLS_functionalComponentArgsRest(__VLS_97));
        const __VLS_102 = __VLS_resolvedLocalAndGlobalComponents.ElButton;
        /** @type { [typeof __VLS_components.ElButton, typeof __VLS_components.elButton, typeof __VLS_components.ElButton, typeof __VLS_components.elButton, ] } */
        // @ts-ignore
        const __VLS_103 = __VLS_asFunctionalComponent(__VLS_102, new __VLS_102({ ...{ 'onClick': {} }, type: ("primary"), icon: ("Edit"), circle: (true), plain: (true), }));
        const __VLS_104 = __VLS_103({ ...{ 'onClick': {} }, type: ("primary"), icon: ("Edit"), circle: (true), plain: (true), }, ...__VLS_functionalComponentArgsRest(__VLS_103));
        let __VLS_108;
        const __VLS_109 = {
            onClick: (...[$event]) => {
                __VLS_ctx.method.handleUpdate(row);
            }
        };
        let __VLS_105;
        let __VLS_106;
        const __VLS_107 = __VLS_pickFunctionalComponentCtx(__VLS_102, __VLS_104);
        __VLS_nonNullable(__VLS_101.slots).default;
        const __VLS_101 = __VLS_pickFunctionalComponentCtx(__VLS_96, __VLS_98);
        const __VLS_110 = __VLS_resolvedLocalAndGlobalComponents.ElTooltip;
        /** @type { [typeof __VLS_components.ElTooltip, typeof __VLS_components.elTooltip, typeof __VLS_components.ElTooltip, typeof __VLS_components.elTooltip, ] } */
        // @ts-ignore
        const __VLS_111 = __VLS_asFunctionalComponent(__VLS_110, new __VLS_110({ content: ("删除🌻"), placement: ("top"), }));
        const __VLS_112 = __VLS_111({ content: ("删除🌻"), placement: ("top"), }, ...__VLS_functionalComponentArgsRest(__VLS_111));
        const __VLS_116 = __VLS_resolvedLocalAndGlobalComponents.ElButton;
        /** @type { [typeof __VLS_components.ElButton, typeof __VLS_components.elButton, typeof __VLS_components.ElButton, typeof __VLS_components.elButton, ] } */
        // @ts-ignore
        const __VLS_117 = __VLS_asFunctionalComponent(__VLS_116, new __VLS_116({ ...{ 'onClick': {} }, type: ("danger"), icon: ("Delete"), circle: (true), plain: (true), }));
        const __VLS_118 = __VLS_117({ ...{ 'onClick': {} }, type: ("danger"), icon: ("Delete"), circle: (true), plain: (true), }, ...__VLS_functionalComponentArgsRest(__VLS_117));
        let __VLS_122;
        const __VLS_123 = {
            onClick: (...[$event]) => {
                __VLS_ctx.method.handleDelete(row);
            }
        };
        let __VLS_119;
        let __VLS_120;
        const __VLS_121 = __VLS_pickFunctionalComponentCtx(__VLS_116, __VLS_118);
        __VLS_nonNullable(__VLS_115.slots).default;
        const __VLS_115 = __VLS_pickFunctionalComponentCtx(__VLS_110, __VLS_112);
    }
    const __VLS_95 = __VLS_pickFunctionalComponentCtx(__VLS_90, __VLS_92);
    __VLS_nonNullable(__VLS_45.slots).default;
    const __VLS_45 = __VLS_pickFunctionalComponentCtx(__VLS_40, __VLS_42);
    __VLS_elementAsFunction(__VLS_intrinsicElements.br)({});
    const __VLS_124 = __VLS_resolvedLocalAndGlobalComponents.ElPagination;
    /** @type { [typeof __VLS_components.ElPagination, typeof __VLS_components.elPagination, ] } */
    // @ts-ignore
    const __VLS_125 = __VLS_asFunctionalComponent(__VLS_124, new __VLS_124({ ...{ 'onSizeChange': {} }, ...{ 'onCurrentChange': {} }, ...{ class: ("flex-center") }, currentPage: ((__VLS_ctx.pageData.pageParams.page)), pageSize: ((__VLS_ctx.pageData.pageParams.pageSize)), background: (true), pageSizes: (([10, 20, 50, 100, 200])), layout: ("total, sizes, prev, pager, next, jumper"), total: ((__VLS_ctx.pageData.total)), }));
    const __VLS_126 = __VLS_125({ ...{ 'onSizeChange': {} }, ...{ 'onCurrentChange': {} }, ...{ class: ("flex-center") }, currentPage: ((__VLS_ctx.pageData.pageParams.page)), pageSize: ((__VLS_ctx.pageData.pageParams.pageSize)), background: (true), pageSizes: (([10, 20, 50, 100, 200])), layout: ("total, sizes, prev, pager, next, jumper"), total: ((__VLS_ctx.pageData.total)), }, ...__VLS_functionalComponentArgsRest(__VLS_125));
    __VLS_directiveAsFunction(__VLS_directives.vShow)(null, { ...__VLS_directiveBindingRestFields, value: (__VLS_ctx.pageData.total > 0) }, null, null);
    let __VLS_130;
    const __VLS_131 = {
        onSizeChange: (__VLS_ctx.method.handleListPage)
    };
    const __VLS_132 = {
        onCurrentChange: (__VLS_ctx.method.handleListPage)
    };
    let __VLS_127;
    let __VLS_128;
    const __VLS_129 = __VLS_pickFunctionalComponentCtx(__VLS_124, __VLS_126);
    const __VLS_133 = __VLS_resolvedLocalAndGlobalComponents.SkyDialog;
    /** @type { [typeof __VLS_components.SkyDialog, typeof __VLS_components.SkyDialog, ] } */
    // @ts-ignore
    const __VLS_134 = __VLS_asFunctionalComponent(__VLS_133, new __VLS_133({ ...{ 'onSkyConfirm': {} }, ...{ 'onSkyCancel': {} }, ref: ("skyDialogRef"), title: ((__VLS_ctx.pageData.title)), loading: ((__VLS_ctx.pageData.confirmLoading)), width: ((700)), }));
    const __VLS_135 = __VLS_134({ ...{ 'onSkyConfirm': {} }, ...{ 'onSkyCancel': {} }, ref: ("skyDialogRef"), title: ((__VLS_ctx.pageData.title)), loading: ((__VLS_ctx.pageData.confirmLoading)), width: ((700)), }, ...__VLS_functionalComponentArgsRest(__VLS_134));
    // @ts-ignore navigation for `const skyDialogRef = ref()`
    __VLS_ctx.skyDialogRef;
    var __VLS_139 = {};
    let __VLS_140;
    const __VLS_141 = {
        onSkyConfirm: (__VLS_ctx.method.handleConfirm)
    };
    const __VLS_142 = {
        onSkyCancel: (__VLS_ctx.method.handleCancel)
    };
    let __VLS_136;
    let __VLS_137;
    __VLS_elementAsFunction(__VLS_intrinsicElements.template, __VLS_intrinsicElements.template)({});
    {
        const { content: __VLS_thisSlot } = __VLS_nonNullable(__VLS_138.slots);
        const __VLS_143 = __VLS_resolvedLocalAndGlobalComponents.ElForm;
        /** @type { [typeof __VLS_components.ElForm, typeof __VLS_components.elForm, typeof __VLS_components.ElForm, typeof __VLS_components.elForm, ] } */
        // @ts-ignore
        const __VLS_144 = __VLS_asFunctionalComponent(__VLS_143, new __VLS_143({ ref: ("formRef"), model: ((__VLS_ctx.pageData.form)), rules: ((__VLS_ctx.pageData.rules)), labelWidth: ("auto"), }));
        const __VLS_145 = __VLS_144({ ref: ("formRef"), model: ((__VLS_ctx.pageData.form)), rules: ((__VLS_ctx.pageData.rules)), labelWidth: ("auto"), }, ...__VLS_functionalComponentArgsRest(__VLS_144));
        // @ts-ignore navigation for `const formRef = ref()`
        __VLS_ctx.formRef;
        var __VLS_149 = {};
        const __VLS_150 = __VLS_resolvedLocalAndGlobalComponents.ElRow;
        /** @type { [typeof __VLS_components.ElRow, typeof __VLS_components.elRow, typeof __VLS_components.ElRow, typeof __VLS_components.elRow, ] } */
        // @ts-ignore
        const __VLS_151 = __VLS_asFunctionalComponent(__VLS_150, new __VLS_150({}));
        const __VLS_152 = __VLS_151({}, ...__VLS_functionalComponentArgsRest(__VLS_151));
        const __VLS_156 = __VLS_resolvedLocalAndGlobalComponents.ElCol;
        /** @type { [typeof __VLS_components.ElCol, typeof __VLS_components.elCol, typeof __VLS_components.ElCol, typeof __VLS_components.elCol, ] } */
        // @ts-ignore
        const __VLS_157 = __VLS_asFunctionalComponent(__VLS_156, new __VLS_156({ xs: (({ span: 24 })), sm: (({ span: 24 })), }));
        const __VLS_158 = __VLS_157({ xs: (({ span: 24 })), sm: (({ span: 24 })), }, ...__VLS_functionalComponentArgsRest(__VLS_157));
        const __VLS_162 = __VLS_resolvedLocalAndGlobalComponents.ElFormItem;
        /** @type { [typeof __VLS_components.ElFormItem, typeof __VLS_components.elFormItem, typeof __VLS_components.ElFormItem, typeof __VLS_components.elFormItem, ] } */
        // @ts-ignore
        const __VLS_163 = __VLS_asFunctionalComponent(__VLS_162, new __VLS_162({ label: ("文件名称"), prop: ("name"), }));
        const __VLS_164 = __VLS_163({ label: ("文件名称"), prop: ("name"), }, ...__VLS_functionalComponentArgsRest(__VLS_163));
        const __VLS_168 = __VLS_resolvedLocalAndGlobalComponents.ElInput;
        /** @type { [typeof __VLS_components.ElInput, typeof __VLS_components.elInput, ] } */
        // @ts-ignore
        const __VLS_169 = __VLS_asFunctionalComponent(__VLS_168, new __VLS_168({ modelValue: ((__VLS_ctx.pageData.form.name)), placeholder: ("请输入文件名称"), clearable: (true), }));
        const __VLS_170 = __VLS_169({ modelValue: ((__VLS_ctx.pageData.form.name)), placeholder: ("请输入文件名称"), clearable: (true), }, ...__VLS_functionalComponentArgsRest(__VLS_169));
        __VLS_nonNullable(__VLS_167.slots).default;
        const __VLS_167 = __VLS_pickFunctionalComponentCtx(__VLS_162, __VLS_164);
        __VLS_nonNullable(__VLS_161.slots).default;
        const __VLS_161 = __VLS_pickFunctionalComponentCtx(__VLS_156, __VLS_158);
        __VLS_nonNullable(__VLS_155.slots).default;
        const __VLS_155 = __VLS_pickFunctionalComponentCtx(__VLS_150, __VLS_152);
        const __VLS_174 = __VLS_resolvedLocalAndGlobalComponents.ElRow;
        /** @type { [typeof __VLS_components.ElRow, typeof __VLS_components.elRow, typeof __VLS_components.ElRow, typeof __VLS_components.elRow, ] } */
        // @ts-ignore
        const __VLS_175 = __VLS_asFunctionalComponent(__VLS_174, new __VLS_174({}));
        const __VLS_176 = __VLS_175({}, ...__VLS_functionalComponentArgsRest(__VLS_175));
        const __VLS_180 = __VLS_resolvedLocalAndGlobalComponents.ElCol;
        /** @type { [typeof __VLS_components.ElCol, typeof __VLS_components.elCol, typeof __VLS_components.ElCol, typeof __VLS_components.elCol, ] } */
        // @ts-ignore
        const __VLS_181 = __VLS_asFunctionalComponent(__VLS_180, new __VLS_180({ xs: (({ span: 24 })), sm: (({ span: 24 })), }));
        const __VLS_182 = __VLS_181({ xs: (({ span: 24 })), sm: (({ span: 24 })), }, ...__VLS_functionalComponentArgsRest(__VLS_181));
        const __VLS_186 = __VLS_resolvedLocalAndGlobalComponents.ElFormItem;
        /** @type { [typeof __VLS_components.ElFormItem, typeof __VLS_components.elFormItem, typeof __VLS_components.ElFormItem, typeof __VLS_components.elFormItem, ] } */
        // @ts-ignore
        const __VLS_187 = __VLS_asFunctionalComponent(__VLS_186, new __VLS_186({ label: ("文件大小"), prop: ("size"), }));
        const __VLS_188 = __VLS_187({ label: ("文件大小"), prop: ("size"), }, ...__VLS_functionalComponentArgsRest(__VLS_187));
        const __VLS_192 = __VLS_resolvedLocalAndGlobalComponents.ElInputNumber;
        /** @type { [typeof __VLS_components.ElInputNumber, typeof __VLS_components.elInputNumber, ] } */
        // @ts-ignore
        const __VLS_193 = __VLS_asFunctionalComponent(__VLS_192, new __VLS_192({ ...{ style: ({}) }, modelValue: ((__VLS_ctx.pageData.form.size)), disabled: (true), clearable: (true), }));
        const __VLS_194 = __VLS_193({ ...{ style: ({}) }, modelValue: ((__VLS_ctx.pageData.form.size)), disabled: (true), clearable: (true), }, ...__VLS_functionalComponentArgsRest(__VLS_193));
        __VLS_nonNullable(__VLS_191.slots).default;
        const __VLS_191 = __VLS_pickFunctionalComponentCtx(__VLS_186, __VLS_188);
        __VLS_nonNullable(__VLS_185.slots).default;
        const __VLS_185 = __VLS_pickFunctionalComponentCtx(__VLS_180, __VLS_182);
        __VLS_nonNullable(__VLS_179.slots).default;
        const __VLS_179 = __VLS_pickFunctionalComponentCtx(__VLS_174, __VLS_176);
        const __VLS_198 = __VLS_resolvedLocalAndGlobalComponents.ElRow;
        /** @type { [typeof __VLS_components.ElRow, typeof __VLS_components.elRow, typeof __VLS_components.ElRow, typeof __VLS_components.elRow, ] } */
        // @ts-ignore
        const __VLS_199 = __VLS_asFunctionalComponent(__VLS_198, new __VLS_198({}));
        const __VLS_200 = __VLS_199({}, ...__VLS_functionalComponentArgsRest(__VLS_199));
        const __VLS_204 = __VLS_resolvedLocalAndGlobalComponents.ElCol;
        /** @type { [typeof __VLS_components.ElCol, typeof __VLS_components.elCol, typeof __VLS_components.ElCol, typeof __VLS_components.elCol, ] } */
        // @ts-ignore
        const __VLS_205 = __VLS_asFunctionalComponent(__VLS_204, new __VLS_204({ xs: (({ span: 24 })), sm: (({ span: 24 })), }));
        const __VLS_206 = __VLS_205({ xs: (({ span: 24 })), sm: (({ span: 24 })), }, ...__VLS_functionalComponentArgsRest(__VLS_205));
        const __VLS_210 = __VLS_resolvedLocalAndGlobalComponents.ElFormItem;
        /** @type { [typeof __VLS_components.ElFormItem, typeof __VLS_components.elFormItem, typeof __VLS_components.ElFormItem, typeof __VLS_components.elFormItem, ] } */
        // @ts-ignore
        const __VLS_211 = __VLS_asFunctionalComponent(__VLS_210, new __VLS_210({ label: ("文件后缀"), prop: ("suffix"), }));
        const __VLS_212 = __VLS_211({ label: ("文件后缀"), prop: ("suffix"), }, ...__VLS_functionalComponentArgsRest(__VLS_211));
        const __VLS_216 = __VLS_resolvedLocalAndGlobalComponents.ElInput;
        /** @type { [typeof __VLS_components.ElInput, typeof __VLS_components.elInput, ] } */
        // @ts-ignore
        const __VLS_217 = __VLS_asFunctionalComponent(__VLS_216, new __VLS_216({ modelValue: ((__VLS_ctx.pageData.form.suffix)), placeholder: ("请输入文件后缀"), disabled: (true), clearable: (true), }));
        const __VLS_218 = __VLS_217({ modelValue: ((__VLS_ctx.pageData.form.suffix)), placeholder: ("请输入文件后缀"), disabled: (true), clearable: (true), }, ...__VLS_functionalComponentArgsRest(__VLS_217));
        __VLS_nonNullable(__VLS_215.slots).default;
        const __VLS_215 = __VLS_pickFunctionalComponentCtx(__VLS_210, __VLS_212);
        __VLS_nonNullable(__VLS_209.slots).default;
        const __VLS_209 = __VLS_pickFunctionalComponentCtx(__VLS_204, __VLS_206);
        __VLS_nonNullable(__VLS_203.slots).default;
        const __VLS_203 = __VLS_pickFunctionalComponentCtx(__VLS_198, __VLS_200);
        const __VLS_222 = __VLS_resolvedLocalAndGlobalComponents.ElRow;
        /** @type { [typeof __VLS_components.ElRow, typeof __VLS_components.elRow, typeof __VLS_components.ElRow, typeof __VLS_components.elRow, ] } */
        // @ts-ignore
        const __VLS_223 = __VLS_asFunctionalComponent(__VLS_222, new __VLS_222({}));
        const __VLS_224 = __VLS_223({}, ...__VLS_functionalComponentArgsRest(__VLS_223));
        const __VLS_228 = __VLS_resolvedLocalAndGlobalComponents.ElCol;
        /** @type { [typeof __VLS_components.ElCol, typeof __VLS_components.elCol, typeof __VLS_components.ElCol, typeof __VLS_components.elCol, ] } */
        // @ts-ignore
        const __VLS_229 = __VLS_asFunctionalComponent(__VLS_228, new __VLS_228({ xs: (({ span: 24 })), sm: (({ span: 24 })), }));
        const __VLS_230 = __VLS_229({ xs: (({ span: 24 })), sm: (({ span: 24 })), }, ...__VLS_functionalComponentArgsRest(__VLS_229));
        const __VLS_234 = __VLS_resolvedLocalAndGlobalComponents.ElFormItem;
        /** @type { [typeof __VLS_components.ElFormItem, typeof __VLS_components.elFormItem, typeof __VLS_components.ElFormItem, typeof __VLS_components.elFormItem, ] } */
        // @ts-ignore
        const __VLS_235 = __VLS_asFunctionalComponent(__VLS_234, new __VLS_234({ label: ("文件所在路径"), prop: ("path"), }));
        const __VLS_236 = __VLS_235({ label: ("文件所在路径"), prop: ("path"), }, ...__VLS_functionalComponentArgsRest(__VLS_235));
        const __VLS_240 = __VLS_resolvedLocalAndGlobalComponents.ElInput;
        /** @type { [typeof __VLS_components.ElInput, typeof __VLS_components.elInput, ] } */
        // @ts-ignore
        const __VLS_241 = __VLS_asFunctionalComponent(__VLS_240, new __VLS_240({ modelValue: ((__VLS_ctx.pageData.form.path)), placeholder: ("请输入文件所在路径"), disabled: (true), clearable: (true), }));
        const __VLS_242 = __VLS_241({ modelValue: ((__VLS_ctx.pageData.form.path)), placeholder: ("请输入文件所在路径"), disabled: (true), clearable: (true), }, ...__VLS_functionalComponentArgsRest(__VLS_241));
        __VLS_nonNullable(__VLS_239.slots).default;
        const __VLS_239 = __VLS_pickFunctionalComponentCtx(__VLS_234, __VLS_236);
        __VLS_nonNullable(__VLS_233.slots).default;
        const __VLS_233 = __VLS_pickFunctionalComponentCtx(__VLS_228, __VLS_230);
        __VLS_nonNullable(__VLS_227.slots).default;
        const __VLS_227 = __VLS_pickFunctionalComponentCtx(__VLS_222, __VLS_224);
        const __VLS_246 = __VLS_resolvedLocalAndGlobalComponents.ElRow;
        /** @type { [typeof __VLS_components.ElRow, typeof __VLS_components.elRow, typeof __VLS_components.ElRow, typeof __VLS_components.elRow, ] } */
        // @ts-ignore
        const __VLS_247 = __VLS_asFunctionalComponent(__VLS_246, new __VLS_246({}));
        const __VLS_248 = __VLS_247({}, ...__VLS_functionalComponentArgsRest(__VLS_247));
        const __VLS_252 = __VLS_resolvedLocalAndGlobalComponents.ElCol;
        /** @type { [typeof __VLS_components.ElCol, typeof __VLS_components.elCol, typeof __VLS_components.ElCol, typeof __VLS_components.elCol, ] } */
        // @ts-ignore
        const __VLS_253 = __VLS_asFunctionalComponent(__VLS_252, new __VLS_252({ xs: (({ span: 24 })), sm: (({ span: 24 })), }));
        const __VLS_254 = __VLS_253({ xs: (({ span: 24 })), sm: (({ span: 24 })), }, ...__VLS_functionalComponentArgsRest(__VLS_253));
        const __VLS_258 = __VLS_resolvedLocalAndGlobalComponents.ElFormItem;
        /** @type { [typeof __VLS_components.ElFormItem, typeof __VLS_components.elFormItem, typeof __VLS_components.ElFormItem, typeof __VLS_components.elFormItem, ] } */
        // @ts-ignore
        const __VLS_259 = __VLS_asFunctionalComponent(__VLS_258, new __VLS_258({ label: ("文件地址"), prop: ("url"), }));
        const __VLS_260 = __VLS_259({ label: ("文件地址"), prop: ("url"), }, ...__VLS_functionalComponentArgsRest(__VLS_259));
        const __VLS_264 = __VLS_resolvedLocalAndGlobalComponents.ElInput;
        /** @type { [typeof __VLS_components.ElInput, typeof __VLS_components.elInput, ] } */
        // @ts-ignore
        const __VLS_265 = __VLS_asFunctionalComponent(__VLS_264, new __VLS_264({ modelValue: ((__VLS_ctx.pageData.form.url)), placeholder: ("请输入文件地址"), disabled: (true), clearable: (true), }));
        const __VLS_266 = __VLS_265({ modelValue: ((__VLS_ctx.pageData.form.url)), placeholder: ("请输入文件地址"), disabled: (true), clearable: (true), }, ...__VLS_functionalComponentArgsRest(__VLS_265));
        __VLS_nonNullable(__VLS_263.slots).default;
        const __VLS_263 = __VLS_pickFunctionalComponentCtx(__VLS_258, __VLS_260);
        __VLS_nonNullable(__VLS_257.slots).default;
        const __VLS_257 = __VLS_pickFunctionalComponentCtx(__VLS_252, __VLS_254);
        __VLS_nonNullable(__VLS_251.slots).default;
        const __VLS_251 = __VLS_pickFunctionalComponentCtx(__VLS_246, __VLS_248);
        const __VLS_270 = __VLS_resolvedLocalAndGlobalComponents.ElRow;
        /** @type { [typeof __VLS_components.ElRow, typeof __VLS_components.elRow, typeof __VLS_components.ElRow, typeof __VLS_components.elRow, ] } */
        // @ts-ignore
        const __VLS_271 = __VLS_asFunctionalComponent(__VLS_270, new __VLS_270({}));
        const __VLS_272 = __VLS_271({}, ...__VLS_functionalComponentArgsRest(__VLS_271));
        const __VLS_276 = __VLS_resolvedLocalAndGlobalComponents.ElCol;
        /** @type { [typeof __VLS_components.ElCol, typeof __VLS_components.elCol, typeof __VLS_components.ElCol, typeof __VLS_components.elCol, ] } */
        // @ts-ignore
        const __VLS_277 = __VLS_asFunctionalComponent(__VLS_276, new __VLS_276({ xs: (({ span: 24 })), sm: (({ span: 24 })), }));
        const __VLS_278 = __VLS_277({ xs: (({ span: 24 })), sm: (({ span: 24 })), }, ...__VLS_functionalComponentArgsRest(__VLS_277));
        const __VLS_282 = __VLS_resolvedLocalAndGlobalComponents.ElFormItem;
        /** @type { [typeof __VLS_components.ElFormItem, typeof __VLS_components.elFormItem, typeof __VLS_components.ElFormItem, typeof __VLS_components.elFormItem, ] } */
        // @ts-ignore
        const __VLS_283 = __VLS_asFunctionalComponent(__VLS_282, new __VLS_282({ label: ("备注"), prop: ("url"), }));
        const __VLS_284 = __VLS_283({ label: ("备注"), prop: ("url"), }, ...__VLS_functionalComponentArgsRest(__VLS_283));
        const __VLS_288 = __VLS_resolvedLocalAndGlobalComponents.ElInput;
        /** @type { [typeof __VLS_components.ElInput, typeof __VLS_components.elInput, ] } */
        // @ts-ignore
        const __VLS_289 = __VLS_asFunctionalComponent(__VLS_288, new __VLS_288({ modelValue: ((__VLS_ctx.pageData.form.remark)), placeholder: ("请输入备注"), clearable: (true), }));
        const __VLS_290 = __VLS_289({ modelValue: ((__VLS_ctx.pageData.form.remark)), placeholder: ("请输入备注"), clearable: (true), }, ...__VLS_functionalComponentArgsRest(__VLS_289));
        __VLS_nonNullable(__VLS_287.slots).default;
        const __VLS_287 = __VLS_pickFunctionalComponentCtx(__VLS_282, __VLS_284);
        __VLS_nonNullable(__VLS_281.slots).default;
        const __VLS_281 = __VLS_pickFunctionalComponentCtx(__VLS_276, __VLS_278);
        __VLS_nonNullable(__VLS_275.slots).default;
        const __VLS_275 = __VLS_pickFunctionalComponentCtx(__VLS_270, __VLS_272);
        __VLS_nonNullable(__VLS_148.slots).default;
        const __VLS_148 = __VLS_pickFunctionalComponentCtx(__VLS_143, __VLS_145);
    }
    const __VLS_138 = __VLS_pickFunctionalComponentCtx(__VLS_133, __VLS_135);
    __VLS_elementAsFunction(__VLS_intrinsicElements.span, __VLS_intrinsicElements.span)({ ...{ class: ("dialog-footer") }, });
    const __VLS_294 = __VLS_resolvedLocalAndGlobalComponents.ElButton;
    /** @type { [typeof __VLS_components.ElButton, typeof __VLS_components.elButton, typeof __VLS_components.ElButton, typeof __VLS_components.elButton, ] } */
    // @ts-ignore
    const __VLS_295 = __VLS_asFunctionalComponent(__VLS_294, new __VLS_294({ ...{ 'onClick': {} }, type: ("primary"), loadingIcon: ("Eleme"), }));
    const __VLS_296 = __VLS_295({ ...{ 'onClick': {} }, type: ("primary"), loadingIcon: ("Eleme"), }, ...__VLS_functionalComponentArgsRest(__VLS_295));
    let __VLS_300;
    const __VLS_301 = {
        onClick: (__VLS_ctx.method.fileConfirm)
    };
    let __VLS_297;
    let __VLS_298;
    __VLS_nonNullable(__VLS_299.slots).default;
    const __VLS_299 = __VLS_pickFunctionalComponentCtx(__VLS_294, __VLS_296);
    const __VLS_302 = __VLS_resolvedLocalAndGlobalComponents.ElButton;
    /** @type { [typeof __VLS_components.ElButton, typeof __VLS_components.elButton, typeof __VLS_components.ElButton, typeof __VLS_components.elButton, ] } */
    // @ts-ignore
    const __VLS_303 = __VLS_asFunctionalComponent(__VLS_302, new __VLS_302({ ...{ 'onClick': {} }, type: ("danger"), }));
    const __VLS_304 = __VLS_303({ ...{ 'onClick': {} }, type: ("danger"), }, ...__VLS_functionalComponentArgsRest(__VLS_303));
    let __VLS_308;
    const __VLS_309 = {
        onClick: (__VLS_ctx.method.fileClose)
    };
    let __VLS_305;
    let __VLS_306;
    __VLS_nonNullable(__VLS_307.slots).default;
    const __VLS_307 = __VLS_pickFunctionalComponentCtx(__VLS_302, __VLS_304);
    __VLS_nonNullable(__VLS_19.slots).default;
    const __VLS_19 = __VLS_pickFunctionalComponentCtx(__VLS_14, __VLS_16);
    __VLS_nonNullable(__VLS_13.slots).default;
    const __VLS_13 = __VLS_pickFunctionalComponentCtx(__VLS_8, __VLS_10);
    __VLS_styleScopedClasses['bucket'];
    __VLS_styleScopedClasses['sky-flex'];
    __VLS_styleScopedClasses['flex-center'];
    __VLS_styleScopedClasses['dialog-footer'];
    var __VLS_slots;
    var __VLS_inheritedAttrs;
    const __VLS_refs = {
        "skyDialogRef": __VLS_139,
        "formRef": __VLS_149,
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
            skyDialogRef: skyDialogRef,
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
});
;
//# sourceMappingURL=bucket.vue.js.map