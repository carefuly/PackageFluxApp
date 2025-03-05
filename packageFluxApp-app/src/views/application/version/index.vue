<script setup lang="ts">
import {useRoute, useRouter} from "vue-router";
import {skyMsgBox, skyMsgError, skyMsgSuccess, skyMsgWarning, skyNoticeError, skyNoticeSuccess} from "@/utils/sky";
import {getById as getByIdApp} from "@/apis/application/details";
import {listAll as fileListAll} from "@/apis/application/file";
import {
  add,
  deleteById,
  updateById,
  setFormal,
  listAll,
  getById,
} from "@/apis/application/version";
import {Delete, Edit, Position, Search, Warning} from "@element-plus/icons-vue";

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
    versionCode: [{required: true, message: "请输入版本号", trigger: "blur"}],
    apkUrl: [{required: true, message: "请输入apkUrl地址", trigger: "blur"}],
    wgtUrl: [{required: true, message: "请输入wgtUrl地址", trigger: "blur"}],
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
  handleGetByIdApp: async (id: any) => {
    try {
      const res: any = await getByIdApp(id);
      pageData.value.details = res.data;
    } catch (error) {
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
  handleFormal: async (row: any) => {
    try {
      const res: any = await setFormal(row.id, pageData.value.pageParams.detailId);
      await method.handleListPage();
      skyMsgSuccess("设置成功🌻");
    } catch (error) {
      skyMsgError(error);
    }
  },
  /** 删除 */
  handleDelete: (row: any) => {
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
        } catch (error) {
          await method.handleListPage();
          skyNoticeError("删除失败，请刷新重试🌻");
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
  handleEcho: async (id: any) => {
    if (id === null || id === "") {
      skyMsgWarning("请选择需要修改的数据🌻");
      return;
    }
    try {
      const res = await getById(id, pageData.value.pageParams);
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
        } else {
          try {
            await add(pageData.value.form);
            skyNoticeSuccess("添加成功🌻");
            pageData.value.confirmLoading = false;
            skyDialogRef.value.skyQuickClose();
            method.resetForm();
            await method.handleListPage();
          } catch (error) {
            pageData.value.confirmLoading = false;
            skyNoticeError("添加失败，请刷新重试🌻");
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
      const res: any = await listAll(pageData.value.pageParams);
      pageData.value.tableList = res.data;
      pageData.value.loading = false;
    } catch (error) {
      skyMsgError("数据查询失败，请刷新重试🌻");
    }
  },
  /** 文件数据表格 */
  handleFileListAll: async () => {
    try {
      pageData.value.fileList = [];
      const res: any = await fileListAll({});
      pageData.value.fileList = res.data;
    } catch (error) {
      skyMsgError("数据查询失败，请刷新重试🌻");
    }
  },
});
const filterList = computed(() => {
  if (pageData.value.searchValue) {
    return pageData.value.tableList.filter(item => item.versionCode.includes(pageData.value.searchValue));
  } else {
    return pageData.value.tableList;
  }
});
onMounted(() => {
  method.handleGetByIdApp(method.handleGetAppId());
  pageData.value.pageParams.detailId = method.handleGetAppId();
  method.handleListPage();
});
</script>

<template>
  <div class="sky-flex">
    <SkyCard>
      <!-- 详情操作 -->
      <el-row :gutter="10" style="display: flex; justify-content: space-between">
        <div style="display: flex; align-items: center;">
          <el-col :span="1.5">
            <el-tooltip content="返回">
              <el-button size="small" type="primary" icon="ArrowLeftBold" circle plain
                         @click="router.push('/application')"></el-button>
            </el-tooltip>
          </el-col>
          <el-col :span="1.5">
            <img style="width: 25px; height: 25px" :src="pageData.details.logoUrl" :alt="pageData.details.appName">
          </el-col>
          <el-col :span="1.5">
            {{ pageData.details.appName }}
            |
            <span style="font-size: 13px">版本管理</span>
          </el-col>
        </div>
        <div>
          <el-col :span="1.5">
            <el-button type="primary" icon="plus" plain @click="method.handleAdd()">新增版本</el-button>
            <el-button type="danger" icon="refresh" plain @click="method.resetSearch">重置</el-button>
          </el-col>
        </div>
        <el-col :span="1.5">
          <el-input
            style="width: 270px"
            placeholder="请输入应用版本号【例如：2.2.8】"
            :prefix-icon="Search"
            v-model="pageData.searchValue"
            clearable
          ></el-input>
        </el-col>
      </el-row>
      <br/>
      <!-- 应用列表 -->
      <el-timeline v-if="filterList.length > 0">
        <el-timeline-item
          v-for="item in filterList"
          :key="item.id"
          :type="item.formalVersion ? 'success': 'info'"
        >
          <!-- 版本操作 -->
          <div class="title">
            <el-text class="mr-2" size="large" type="primary">{{ item.versionCode }}</el-text>
            <!-- 设为正式版 -->
            <el-popconfirm
              v-if="!item.formalVersion"
              width="220"
              title="是否将当前版本设为正式版"
              type="warning"
              ok-text="确认"
              cancel-text="取消"
              @confirm="method.handleFormal(item)"
            >
              <template #reference>
                <el-icon class="mr-2">
                  <Position/>
                </el-icon>
              </template>
            </el-popconfirm>
            <!-- 编辑 -->
            <el-tooltip content="编辑">
              <el-icon class="mr-2" @click="method.handleUpdate(item)">
                <Edit/>
              </el-icon>
            </el-tooltip>
            <!-- 删除 -->
            <el-popconfirm
              v-if="!item.formalVersion"
              title="是否删除当前版本"
              width="180"
              @confirm="method.handleDelete(item)"
            >
              <template #reference>
                <el-icon>
                  <Delete/>
                </el-icon>
              </template>
            </el-popconfirm>
          </div>
          <!-- 版本信息 -->
          <div class="mt-2 mb-2">
            <el-text
              v-if="item.description"
              class="mx-1"
              v-for="desc in item.description.split('\n')"
              :key="desc"
            >
              <el-tooltip :content="desc">
                {{ desc }}
              </el-tooltip>
              <br/>
            </el-text>
          </div>
          <el-text class="mx-1" type="success">
            Apk下载地址:
            <el-text class="mx-1">{{ item.apkUrl }}</el-text>
          </el-text>
          <div class="h-0.5em"></div>
          <el-text class="mx-1" type="warning">
            Wgt下载地址:
            <el-text class="mx-1">{{ item.wgtUrl }}</el-text>
          </el-text>
          <div class="h-0.5em"></div>
          <el-tag v-if="item.HBuilderUpdate" class="mr-2" type="primary">HBuilderX强制更新</el-tag>
          <el-tag v-if="item.forceUpdate" type="success">强制更新</el-tag>
          <div class="h-0.5em"></div>
          <el-tag v-if="item.formalVersion" type="danger">正式版本</el-tag>
          <div class="h-0.5em"></div>
          <el-text class="mx-1" type="danger">
            上次更新时间：
            <el-text class="mx-1">{{ item.updateTime }}</el-text>
          </el-text>
        </el-timeline-item>
      </el-timeline>
      <div v-else class="message center">
        <h1>版本为空</h1>
        <p>请通过中上角的新增按钮添加版本</p>
      </div>
      <!-- 添加 OR 修改 -->
      <SkyDialog
        ref="skyDialogRef"
        :title="pageData.title"
        @skyConfirm="method.handleConfirm"
        @skyCancel="method.handleCancel"
        :loading="pageData.confirmLoading"
        :width="750"
        :height="450"
      >
        <template #content>
          <el-form ref="formRef" :model="pageData.form" :rules="pageData.rules" label-width="auto">
            <el-row>
              <el-col :xs="{ span: 24 }" :sm="{ span: 24 }">
                <el-form-item label="应用版本号" prop="versionCode">
                  <el-input v-model="pageData.form.versionCode"
                            placeholder="请输入应用版本号，例如：2.2.8" :disabled="pageData.form.id !== null" clearable/>
                </el-form-item>
              </el-col>
            </el-row>
            <el-row>
              <el-col :xs="{ span: 24 }" :sm="{ span: 24 }">
                <el-form-item label="更新说明" prop="description">
                  <el-input v-model="pageData.form.description" :rows="3" type="textarea" placeholder="例如：1.修复了部分bug
          2.增加了某某功能" clearable/>
                </el-form-item>
              </el-col>
            </el-row>
            <el-row>
              <el-col :xs="{ span: 24 }" :sm="{ span: 24 }">
                <el-form-item label="HBuilderX是否更新">
                  <template #label>
                    <div style="display: flex; align-items: center">
                      <span style="margin-right: 5px">HBuilderX是否更新</span>
                      <el-tooltip>
                        <template #content>
                          是的话该版本只会进行apk全量更新
                        </template>
                        <el-icon>
                          <Warning/>
                        </el-icon>
                      </el-tooltip>
                    </div>
                  </template>
                  <el-switch v-model="pageData.form.HBuilderUpdate" inline-prompt active-text="是" inactive-text="否"/>
                </el-form-item>
              </el-col>
            </el-row>
            <el-row>
              <el-col :xs="{ span: 24 }" :sm="{ span: 24 }">
                <el-form-item label="是否强制更新">
                  <el-switch v-model="pageData.form.forceUpdate" inline-prompt active-text="是" inactive-text="否"/>
                  <template #label>
                    <div style="display: flex; align-items: center">
                      <span style="margin-right: 5px">是否强制更新</span>
                      <el-tooltip>
                        <template #content>
                          否的话更新弹窗底部会有[暂不更新]按钮
                        </template>
                        <el-icon>
                          <Warning/>
                        </el-icon>
                      </el-tooltip>
                    </div>
                  </template>
                </el-form-item>
              </el-col>
            </el-row>
            <el-row>
              <el-col :xs="{ span: 24 }" :sm="{ span: 24 }">
                <el-form-item label="apk下载地址" prop="apkUrl">
                  <el-select
                    v-model="pageData.form.apkUrl"
                    filterable
                    allow-create
                    default-first-option
                    :reserve-keyword="false"
                    placeholder="请输入或选择apk下载地址"
                    clearable
                  >
                    <el-option
                      v-for="item in pageData.fileList"
                      :key="item.id"
                      :label="item.name + `.${item.suffix}`"
                      :value="item.url"
                    />
                  </el-select>
                </el-form-item>
              </el-col>
            </el-row>
            <el-row>
              <el-col :xs="{ span: 24 }" :sm="{ span: 24 }">
                <el-form-item label="wgt下载地址" prop="wgtUrl">
                  <el-select
                    v-model="pageData.form.wgtUrl"
                    filterable
                    allow-create
                    default-first-option
                    :reserve-keyword="false"
                    placeholder="请输入或选择wgt下载地址"
                    clearable
                  >
                    <el-option
                      v-for="item in pageData.fileList"
                      :key="item.id"
                      :label="item.name + `.${item.suffix}`"
                      :value="item.url"
                    />
                  </el-select>
                </el-form-item>
              </el-col>
            </el-row>
            <el-row>
              <el-col :xs="{ span: 24 }" :sm="{ span: 24 }">
                <el-form-item label="备注" prop="remark">
                  <el-input v-model="pageData.form.remark" :rows="3" type="textarea" placeholder="请输入备注"/>
                </el-form-item>
              </el-col>
            </el-row>
          </el-form>
        </template>
      </SkyDialog>
    </SkyCard>
  </div>
</template>

<style scoped lang="scss">
.message {
  height: 600px;
  flex-direction: column;
}

.center {
  display: flex;
  align-items: center;
  justify-content: center;
}

.title {
  display: flex;
  align-items: center;
}

.url {
  margin: 10px 0;
}
</style>
