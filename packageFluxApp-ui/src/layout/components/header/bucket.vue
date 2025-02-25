<script setup lang="ts">
import {skyMsgBox, skyMsgError, skyMsgInfo, skyMsgWarning, skyNoticeError, skyNoticeSuccess} from "@/utils/sky";
import {deleteById, batchDelete, updateById, listPage, getById} from "@/apis/application/file";

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
    name: [{required: true, message: "请输入文件名称", trigger: "blur"}],
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
  handleDelete: (row: any) => {
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
        } catch (error) {
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
          const res: any = await batchDelete(pageData.value.ids);
          await method.handleListPage();
          skyNoticeSuccess(`${res.msg}🌻`);
        } catch (error) {
          await method.handleListPage();
          skyNoticeSuccess("批量删除失败，请刷新重试🌻");
        }
      })
      .catch(() => {
        skyMsgError("已取消🌻");
      });
  },
  /** 修改 */
  handleUpdate: (row: any) => {
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
  handleEcho: async (id: any) => {
    if (id === null || id === "") {
      skyMsgWarning("请选择需要修改的数据🌻");
      return;
    }
    try {
      const res = await getById(id);
      pageData.value.form = res.data;
    } catch (error) {
      skyNoticeError("数据获取失败，请刷新重试🌻");
    }
  },
  /** 确定  */
  handleConfirm: () => {
    if (!formRef.value) return;
    pageData.value.confirmLoading = true;
    (formRef.value).validate(async (valid: any) => {
      if (valid) {
        if (pageData.value.form.id) {
          try {
            await updateById(pageData.value.form.id, pageData.value.form);
            skyNoticeSuccess("修改成功🌻");
            pageData.value.confirmLoading = false;
            skyDialogRef.value.skyQuickClose();
            method.resetForm();
            await method.handleListPage();
          } catch (error) {
            pageData.value.confirmLoading = false;
            skyNoticeError("修改失败，请刷新重试🌻");
          }
        }
      } else {
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
  handleSelectionChange: (selection: any) => {
    pageData.value.ids = selection.map((item: any) => item.id);
    pageData.value.multiple = !selection.length; // 多选
  },
  /** 数据表格 */
  handleListPage: async () => {
    try {
      pageData.value.loading = true;
      pageData.value.tableList = [];
      const res: any = await listPage(pageData.value.pageParams);
      pageData.value.tableList = res.data.items;
      pageData.value.total = res.data.total;
      pageData.value.loading = false;
    } catch (error) {
      skyMsgError("数据查询失败，请刷新重试🌻");
    }
  },
});
</script>

<template>
  <div class="bucket">
    <el-button type="primary" icon="TakeawayBox" link @click="method.fileOpen">存储桶</el-button>
    <!-- 文件存储桶列表 -->
    <el-dialog
      :model-value="pageData.visible"
      :title="pageData.title"
      :close-on-click-modal="false"
      append-to-body
      draggable
      :destroy-on-close="true"
      :before-close="method.fileClose"
      :fullscreen="true"
    >
      <div style="height: 90vh" class="sky-flex">
        <SkyCard>
          <!-- 表格头部按钮 -->
          <el-row :gutter="10">
            <el-col :span="1.5">
              <el-button type="danger" icon="delete" plain @click="method.handleBatchDelete()" :disabled="pageData.multiple">
                删除
              </el-button>
              <el-button type="success" icon="refresh" plain v-debounce="method.handleListPage">重置</el-button>
            </el-col>
          </el-row>
          <br/>
          <!-- 数据表格 -->
          <el-table
            v-loading="pageData.loading"
            border
            :data="pageData.tableList"
            empty-text="暂时没有数据哟🌻"
            @selection-change="method.handleSelectionChange"
          >
            <el-table-column type="selection" width="55" align="center" fixed="left"/>
            <el-table-column label="文件标识" prop="id" width="110" align="center"
                             :show-overflow-tooltip="true"></el-table-column>
            <el-table-column label="文件名称" prop="name" width="180" align="center"
                             :show-overflow-tooltip="true">
              <template #default="scope">
                <a :href="scope.row.url" target="_blank">
                  {{ scope.row.name }}
                </a>
              </template>
            </el-table-column>
            <el-table-column label="文件大小(KB)" prop="size" width="160" align="center"
                             :show-overflow-tooltip="true">
              <template #default="scope">
                {{ (scope.row.size / 1024).toFixed(2) }} KB
              </template>
            </el-table-column>
            <el-table-column label="文件后缀(类型)" prop="suffix" width="160" align="center"
                             :show-overflow-tooltip="true"></el-table-column>
            <el-table-column label="项目根目录" prop="baseDir" align="center"
                             :show-overflow-tooltip="true"></el-table-column>
            <el-table-column label="文件绝对路径" prop="fileDir" align="center"
                             :show-overflow-tooltip="true"></el-table-column>
            <el-table-column label="文件地址" prop="path" align="center"
                             :show-overflow-tooltip="true"></el-table-column>
            <el-table-column label="完整URL" prop="url" align="center" :show-overflow-tooltip="true"></el-table-column>
            <el-table-column label="操作" align="center" width="120" fixed="right">
              <template #default="{ row }">
                <el-tooltip content="修改🌻" placement="top">
                  <el-button
                    type="primary"
                    icon="Edit"
                    circle
                    plain
                    @click="method.handleUpdate(row)"
                  ></el-button>
                </el-tooltip>
                <el-tooltip content="删除🌻" placement="top">
                  <el-button
                    type="danger"
                    icon="Delete"
                    circle
                    plain
                    @click="method.handleDelete(row)"
                  ></el-button>
                </el-tooltip>
              </template>
            </el-table-column>
          </el-table>
          <br/>
          <!-- 分页 -->
          <el-pagination
            class="flex-center"
            v-model:current-page="pageData.pageParams.page"
            v-model:page-size="pageData.pageParams.pageSize"
            v-show="pageData.total > 0"
            background
            :page-sizes="[10, 20, 50, 100, 200]"
            layout="total, sizes, prev, pager, next, jumper"
            :total="pageData.total"
            @size-change="method.handleListPage"
            @current-change="method.handleListPage"
          />
          <!-- 添加 OR 修改 -->
          <SkyDialog
            ref="skyDialogRef"
            :title="pageData.title"
            @skyConfirm="method.handleConfirm"
            @skyCancel="method.handleCancel"
            :loading="pageData.confirmLoading"
            :width="700"
          >
            <template #content>
              <el-form ref="formRef" :model="pageData.form" :rules="pageData.rules" label-width="auto">
                <el-row>
                  <el-col :xs="{ span: 24 }" :sm="{ span: 24 }">
                    <el-form-item label="文件名称" prop="name">
                      <el-input v-model="pageData.form.name" placeholder="请输入文件名称" clearable/>
                    </el-form-item>
                  </el-col>
                </el-row>
                <el-row>
                  <el-col :xs="{ span: 24 }" :sm="{ span: 24 }">
                    <el-form-item label="文件大小" prop="size">
                      <el-input-number style="width: 100%" v-model="pageData.form.size" disabled clearable/>
                    </el-form-item>
                  </el-col>
                </el-row>
                <el-row>
                  <el-col :xs="{ span: 24 }" :sm="{ span: 24 }">
                    <el-form-item label="文件后缀" prop="suffix">
                      <el-input v-model="pageData.form.suffix" placeholder="请输入文件后缀" disabled clearable/>
                    </el-form-item>
                  </el-col>
                </el-row>
                <el-row>
                  <el-col :xs="{ span: 24 }" :sm="{ span: 24 }">
                    <el-form-item label="文件所在路径" prop="path">
                      <el-input v-model="pageData.form.path" placeholder="请输入文件所在路径" disabled clearable/>
                    </el-form-item>
                  </el-col>
                </el-row>
                <el-row>
                  <el-col :xs="{ span: 24 }" :sm="{ span: 24 }">
                    <el-form-item label="文件地址" prop="url">
                      <el-input v-model="pageData.form.url" placeholder="请输入文件地址" disabled clearable/>
                    </el-form-item>
                  </el-col>
                </el-row>
                <el-row>
                  <el-col :xs="{ span: 24 }" :sm="{ span: 24 }">
                    <el-form-item label="备注" prop="url">
                      <el-input v-model="pageData.form.remark" placeholder="请输入备注" clearable/>
                    </el-form-item>
                  </el-col>
                </el-row>
              </el-form>
            </template>
          </SkyDialog>
          <!-- 操作 -->
          <span class="dialog-footer">
            <el-button type="primary" loading-icon="Eleme" @click="method.fileConfirm">
              确定
            </el-button>
            <el-button type="danger" @click="method.fileClose">取消</el-button>
          </span>
        </SkyCard>
      </div>
    </el-dialog>
  </div>
</template>

<style scoped lang="scss">
.bucket {
  margin-right: 15px;
}

.flex-center {
  width: 100%;
  display: flex;
  justify-content: center;
}

.dialog-footer {
  width: 100%;
  display: flex;
  justify-content: end;
}
</style>
