<script setup lang="ts">
import {useRouter} from "vue-router";
import accomplish from "@/global/method";
import {useUserStore} from "@/store";
import {skyNoticeSuccess, skyNoticeError, skyMsgError, skyMsgWarning, skyMsgBox} from "@/utils/sky.js";
import {listAll as fileListAll} from "@/apis/application/file";
import {
  add,
  deleteById,
  updateById,
  getById,
  listAll,
} from "@/apis/application/details";
import {Search, Warning} from "@element-plus/icons-vue";

const author = import.meta.env.VITE_AUTHOR;
const qq = import.meta.env.VITE_QQ;

const router = useRouter();
const userStore = useUserStore();
const skyDialogRef = ref();
const formRef = ref();
const pageData = ref({
  searchValue: "",
  pageParams: {
    creator: null,
    modifier: null,
    status: true,
    appName: "",
  },
  title: "应用详情",
  confirmLoading: false,
  form: {
    id: null,
    recordId: null,
    logoUrl: "",
    appName: "",
    appleId: null,
    description: null,
    preview: [],
    sort: 1,
    status: true,
    belong_dept: null,
    remark: null,
  },
  rules: {
    appName: [{required: true, message: "请输入应用名称", trigger: "blur"}],
  },
  loading: false,
  total: 0,
  fileList: [],
  tableList: [],
});
const method = reactive({
  /** 重置 */
  resetSearch: () => {
    method.resetSearchParams();
    method.handleListAll();
  },
  /** 重置搜索参数 */
  resetSearchParams: () => {
    accomplish.resetForm(null, pageData.value.pageParams);
  },
  /** 添加 */
  handleAdd: () => {
    // 标题
    pageData.value.title = "添加";
    // 文件数据表格
    // method.handleFileListAll();
    // 重置表单
    accomplish.resetForm(formRef, pageData.value.form);
    // 打开弹出框
    skyDialogRef.value.skyOpen();
  },
  /** 查看统计 */
  handleSkipLog: (row: any) => {
    router.push(`/application/statistic/${row.id}`);
  },
  /** 版本管理 */
  handleSkipVersion: (row: any) => {
    router.push(`/application/version/${row.id}`);
  },
  /** 删除 */
  handleDelete: (row: any) => {
    const id = row.recordId;
    if (id === null || id === "") {
      skyMsgWarning("请选中需要删除的数据🌻");
      return;
    }
    skyMsgBox("您确认需要删除名称[" + row.appName + "]么？")
      .then(async () => {
        try {
          await deleteById(id);
          await method.handleListAll();
          skyNoticeSuccess("删除成功🌻");
        } catch (error) {
          skyNoticeError("删除失败，请刷新重试🌻");
        }
      })
      .catch(() => {
        skyMsgError("已取消🌻");
      });
  },
  /** 修改 */
  handleUpdate: (row: any) => {
    // 标题
    pageData.value.title = "修改";
    // 文件数据表格
    // method.handleFileListAll();
    // 重置表单
    accomplish.resetForm(formRef, pageData.value.form);
    const id = row.recordId;
    if (id == null || id === "") {
      skyMsgError("请选中需要修改的数据🌻");
    }
    // 回显数据
    method.handleEcho(id);
    // 打开弹出框
    skyDialogRef.value.skyOpen();
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
            await updateById(pageData.value.form.recordId, pageData.value.form);
            skyNoticeSuccess("修改成功🌻");
            pageData.value.confirmLoading = false;
            skyDialogRef.value.skyQuickClose();
            method.resetForm();
            await method.handleListAll();
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
            await method.handleListAll();
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
  /** 数据表格 */
  handleListAll: async () => {
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
    return pageData.value.tableList.filter(item => item.appName.includes(pageData.value.searchValue));
  } else {
    return pageData.value.tableList;
  }
});

onMounted(() => {
  // method.handleFileListAll();
  method.handleListAll();
});
</script>

<template>
  <div class="sky-flex">
    <SkyCard>
      <!-- 网站通知 -->
      <el-alert type="success" show-icon>
        <template #title>
          作者昵称：{{ author }}，
          作者QQ：{{ qq }}，
          有任何问题请加作者沟通(备注)，避免造成应用发布后无法更新的损失
        </template>
      </el-alert>
      <br/>
      <el-alert type="warning" show-icon>
        <template #title>
          已上线储桶功能，每个账号存储桶暂无限制，可以自定义上传logo、apk和wgt文件, <br/>
          后续可能因为文件过多，会自动限制账号存储空间，如需要更多的空间请联系管理员扩容！<br/>
          <!--          已上线计费功能，可在右上角头像菜单前往计费套餐查看和购买次数，当前账号使用次数：-->
          当前账号使用：{{ userStore.user?.usageNumber }} 次，
          剩余次数： {{ userStore.user?.total - userStore.user?.usageNumber }}
        </template>
      </el-alert>
      <br/>
      <!-- 详情操作 -->
      <el-row :gutter="10" style="display: flex; justify-content: space-between">
        <el-col :span="1.5">
          <el-button type="primary" icon="plus" plain @click="method.handleAdd()">新增应用</el-button>
          <el-button type="danger" icon="refresh" plain v-debounce="method.resetSearch">重置</el-button>
        </el-col>
        <el-col :span="1.5">
          <el-input
            style="width: 240px"
            placeholder="请输入应用名称"
            :prefix-icon="Search"
            v-model="pageData.searchValue"
            clearable
          ></el-input>
        </el-col>
      </el-row>
      <!-- 应用列表 -->
      <el-descriptions
        style="margin-top: 20px"
        v-if="filterList.length > 0"
        v-for="item in filterList"
        :key="item.id"
        :title="item.appName"
        border
      >
        <el-descriptions-item
          :rowspan="2"
          :width="100"
          label="Logo"
          align="center"
        >
          <el-image
            style="width: 100px; height: 100px"
            :src="item.logoUrl"
          />
        </el-descriptions-item>
        <el-descriptions-item label="应用名称">{{ item.appName }}</el-descriptions-item>
        <el-descriptions-item label="当前正式版">
          <el-tag v-if="item?.versionCode" type="success">
            <el-tooltip :content="item?.versionCode">
              版本号：{{ item?.versionCode }}
            </el-tooltip>
          </el-tag>
          <el-tag v-else type="danger">当前暂无正式版哦，快去发布吧！</el-tag>
        </el-descriptions-item>
        <el-descriptions-item label="跳转发布页">
          <el-tooltip content="前往发布页">
            <el-button size="small" type="success" icon="Share" circle plain
                       @click="router.push(`/application/issue/${item.id}`)"></el-button>
          </el-tooltip>
        </el-descriptions-item>
        <el-descriptions-item label="其他操作">
          <el-tooltip content="查看统计🌻" placement="top">
            <el-button
              type="warning"
              icon="Histogram"
              circle
              plain
              text
              @click="method.handleSkipLog(item)"
            ></el-button>
          </el-tooltip>
          <el-tooltip content="版本管理🌻" placement="top">
            <el-button
              type="success"
              icon="Menu"
              circle
              plain
              text
              @click="method.handleSkipVersion(item)"
            ></el-button>
          </el-tooltip>
          <el-tooltip content="修改🌻" placement="top">
            <el-button
              type="primary"
              icon="Edit"
              circle
              plain
              text
              @click="method.handleUpdate(item)"
            ></el-button>
          </el-tooltip>
          <el-tooltip content="删除🌻" placement="top">
            <el-button
              type="danger"
              icon="Delete"
              circle
              plain
              text
              @click="method.handleDelete(item)"
            ></el-button>
          </el-tooltip>
        </el-descriptions-item>
        <el-descriptions-item label="其他详情">
          应用id：<span v-copy="item.recordId">{{ item.recordId }}</span> <br/>
          appleId：{{ item.appleId }} <br/>
          应用描述：{{ item.description }} <br/>
          备注：{{ item.remark }} <br/>
          修改时间：{{ item.updateTime }} <br/>
        </el-descriptions-item>
      </el-descriptions>
      <div v-else class="message center">
        <h1>应用为空</h1>
        <p>请通过左上角的新增按钮添加应用</p>
      </div>
      <!-- 添加 OR 修改 -->
      <SkyDialog
        ref="skyDialogRef"
        :title="pageData.title"
        @skyConfirm="method.handleConfirm"
        @skyCancel="method.handleCancel"
        :loading="pageData.confirmLoading"
        :width="700"
        :height="430"
      >
        <template #content>
          <el-form ref="formRef" :model="pageData.form" :rules="pageData.rules" label-width="auto">
            <el-row>
              <el-col :xs="{ span: 24 }" :sm="{ span: 24 }">
                <el-form-item label="应用logo" prop="logoUrl">
                  <el-input v-model="pageData.form.logoUrl"
                            placeholder="请手动上传logo图片到自己的云存储或图床等，将图片url粘贴到此处" clearable/>
                </el-form-item>
              </el-col>
            </el-row>
            <el-row>
              <el-col :xs="{ span: 24 }" :sm="{ span: 24 }">
                <el-form-item label="应用名称" prop="appName">
                  <el-input v-model="pageData.form.appName" placeholder="应用名请勿包含空格" clearable/>
                </el-form-item>
              </el-col>
            </el-row>
            <el-row>
              <el-col :xs="{ span: 24 }" :sm="{ span: 24 }">
                <el-form-item prop="appleId">
                  <template #label>
                    <div style="display: flex; align-items: center">
                      <span style="margin-right: 5px">appleId</span>
                      <el-tooltip>
                        <template #content>
                          在App Store Connect中的App Store下的app信息，可找到appleId
                        </template>
                        <el-icon>
                          <Warning/>
                        </el-icon>
                      </el-tooltip>
                    </div>
                  </template>
                  <el-input v-model="pageData.form.appleId"
                            placeholder="ios跳转到商店更新用，若不懂或不需要发布ios平台可不填"
                            clearable/>
                </el-form-item>
              </el-col>
            </el-row>
            <el-row>
              <el-col :xs="{ span: 24 }" :sm="{ span: 24 }">
                <el-form-item label="应用描述" prop="description">
                  <el-input v-model="pageData.form.description" :rows="3" type="textarea" placeholder="请输入应用描述"/>
                </el-form-item>
              </el-col>
            </el-row>
            <el-row>
              <el-col :xs="{ span: 24 }" :sm="{ span: 24 }">
                <el-form-item label="预览截图" prop="preview">
                  <el-select
                    v-model="pageData.form.preview"
                    multiple
                    filterable
                    allow-create
                    default-first-option
                    :reserve-keyword="false"
                    placeholder="请输入或选择展示图片【回车保存】"
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
  height: 500px;
  flex-direction: column;
}

.center {
  display: flex;
  align-items: center;
  justify-content: center;
}
</style>
