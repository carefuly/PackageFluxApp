<script setup lang="ts">
import {useRoute, useRouter} from "vue-router";
import {getNowDate} from "@/utils";
import {skyMsgBox, skyMsgError, skyMsgSuccess, skyMsgWarning, skyNoticeError, skyNoticeSuccess} from "@/utils/sky";
import {getById as getByIdApp} from "@/apis/application/details";
import {statistics} from "@/apis/application/statistic";
import {listPage} from "@/apis/application/log";
import {Warning} from "@element-plus/icons-vue";

const route = useRoute();
const router = useRouter();
const pageData = ref({
  id: null,
  details: {
    logoUrl: "",
    appName: "",
  },
  pageParams: {
    create_time: null,
    detailId: null,
    page: 1,
    pageSize: 10,
  },
  title: "应用统计",
  date: null,
  total: 0,
  loading: false,
  frequency: {
    openNum: 0,
    openTotalNum: 0,
    userNum: 0,
    userTotalNum: 0,
  },
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
  /** 选择日期 */
  handleChangeDate: (e) => {
    pageData.value.date = e;
    pageData.value.pageParams.create_time = e;
    method.handleFrequency();
    method.handleListPage();
  },
  /** 获取统计数据 */
  handleFrequency: async () => {
    try {
      pageData.value.frequency = {
        OpenNum: 0,
        userTotalNum: 0,
      };
      const res: any = await statistics({
        detailId: pageData.value.pageParams.detailId,
        create_time: pageData.value.date,
      });
      pageData.value.frequency = res.data;
    } catch (error) {
      skyMsgError("数据查询失败，请刷新重试🌻");
    }
  },
  /** 数据表格 */
  handleListPage: async () => {
    try {
      pageData.value.loading = true;
      pageData.value.tableList = [];
      const res: any = await listPage(pageData.value.pageParams);
      pageData.value.total = res.data.total;
      pageData.value.tableList = res.data.list;
      pageData.value.loading = false;
    } catch (error) {
      skyMsgError("数据查询失败，请刷新重试🌻");
    }
  },
});
onMounted(() => {
  pageData.value.date = getNowDate();
  pageData.value.pageParams.detailId = method.handleGetAppId();
  pageData.value.pageParams.create_time = getNowDate();
  method.handleGetByIdApp(method.handleGetAppId());
  method.handleFrequency();
  method.handleListPage();
});
</script>

<template>
  <div class="sky-flex">
    <SkyCard>
      <!-- 标题 -->
      <el-row :gutter="10">
        <div style="display: flex; align-items: center;">
          <el-col :span="1.5">
            <el-tooltip content="返回">
              <el-button size="small" type="primary" icon="ArrowLeftBold" circle plain
                         @click="router.push('/application')"></el-button>
            </el-tooltip>
          </el-col>
          <img style="width: 30px; height: 30px" :src="pageData.details.logoUrl" :alt="pageData.details.appName">
          <el-col :span="1.5">
            {{ pageData.details.appName }}
            |
            <span style="font-size: 13px">统计数据</span>
          </el-col>
        </div>
        <el-divider></el-divider>
      </el-row>
      <!-- 指标 -->
      <el-row :gutter="10">
        <el-col :span="1.5">
          <span class="mr4">访问指标概述</span>
          <el-date-picker
            style="width: 160px;"
            v-model="pageData.date"
            type="date"
            value-format="YYYY-MM-DD"
            placeholder="请选择日期"
            @change="method.handleChangeDate"
          />
        </el-col>
      </el-row>
      <br/>
      <!-- 数据 -->
      <div class="mb-4 flex-row">
        <div class="flex-1">
          <div class="statistic">
            <el-statistic :value="pageData.frequency.openNum">
              <template #suffix>
                <span style="font-size: 15px">次</span>
              </template>
              <template #title>
                <div style="display: inline-flex; align-items: center">
                  打开次数
                  <el-tooltip
                    placement="top"
                    effect="dark"
                    :content="pageData.date + '这天，用户打开APP的次数'"
                  >
                    <el-icon style="margin-left: 4px" :size="12">
                      <Warning/>
                    </el-icon>
                  </el-tooltip>
                </div>
              </template>
            </el-statistic>
          </div>
        </div>
        <div class="flex-1 flex-row">
          <el-divider style="height: 100%" direction="vertical"></el-divider>
          <div class="statistic">
            <el-statistic :value="pageData.frequency.userNum">
              <template #suffix>
                <span style="font-size: 15px">人</span>
              </template>
              <template #title>
                <div style="display: inline-flex; align-items: center">
                  访问人数
                  <el-tooltip
                    placement="top"
                    effect="dark"
                    :content="pageData.date + '这天有多少人使用了APP'"
                  >
                    <el-icon style="margin-left: 4px" :size="12">
                      <Warning/>
                    </el-icon>
                  </el-tooltip>
                </div>
              </template>
            </el-statistic>
          </div>
        </div>
      </div>
      <el-divider></el-divider>
      <div class="mb-4 flex-row">
        <div class="flex-1">
          <div class="statistic">
            <el-statistic :value="pageData.frequency.openTotalNum">
              <template #suffix>
                <span style="font-size: 15px">次</span>
              </template>
              <template #title>
                <div style="display: inline-flex; align-items: center">
                  累计打开次数
                  <el-tooltip
                    placement="top"
                    effect="dark"
                    :content="'累计至今天，所有用户打开APP的次数'"
                  >
                    <el-icon style="margin-left: 4px" :size="12">
                      <Warning/>
                    </el-icon>
                  </el-tooltip>
                </div>
              </template>
            </el-statistic>
          </div>
        </div>
        <div class="flex-1 flex-row">
          <el-divider style="height: 100%" direction="vertical"></el-divider>
          <div class="statistic">
            <el-statistic :value="pageData.frequency.userTotalNum">
              <template #suffix>
                <span style="font-size: 15px">次</span>
              </template>
              <template #title>
                <div style="display: inline-flex; align-items: center">
                  累计访问次数
                  <el-tooltip
                    placement="top"
                    effect="dark"
                    :content="'累计至今天，共有多少人使用了APP'"
                  >
                    <el-icon style="margin-left: 4px" :size="12">
                      <Warning/>
                    </el-icon>
                  </el-tooltip>
                </div>
              </template>
            </el-statistic>
          </div>
        </div>
      </div>
      <br/>
      <!-- 操作 -->
      <el-row :gutter="10" style="display: flex; align-items: center; justify-content: space-between">
        <el-col :span="1.5">
          <span class="mr4">今日请求日志</span>
        </el-col>
        <el-col :span="1.5">
          <el-tooltip content="刷新日志">
            <el-button size="small" type="info" icon="RefreshRight" circle plain @click="method.handleListPage"></el-button>
          </el-tooltip>
        </el-col>
      </el-row>
      <br/>
      <!-- 数据表格 -->
      <div style="height: 270px">
        <el-table
          v-loading="pageData.loading"
          height="270"
          border
          :data="pageData.tableList"
          empty-text="暂时没有数据哟🌻"
        >
          <el-table-column label="唯一Id" prop="uniqueId" align="center" :show-overflow-tooltip="true"/>
          <el-table-column label="应用标识Id" prop="appId" align="center" :show-overflow-tooltip="true"/>
          <el-table-column label="应用名称" prop="appName" align="center" :show-overflow-tooltip="true"/>
          <el-table-column label="应用版本(检查的版本)" prop="appVersion" width="200" align="center" :show-overflow-tooltip="true"/>
          <el-table-column label="设备品牌" prop="brand" align="center" :show-overflow-tooltip="true"/>
          <el-table-column label="系统名称" prop="systemName" align="center" :show-overflow-tooltip="true"/>
          <el-table-column label="应用标识备注" prop="uniRemark" align="center" :show-overflow-tooltip="true"/>
        </el-table>
      </div>
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
    </SkyCard>
  </div>
</template>

<style scoped lang="scss">
.sky-flex {

  .flex-1 {
    flex: 1 !important
  }

  .flex-row {
    display: flex;
    flex-direction: row;
  }

  .statistic {
    height: 100%;
    border-radius: 4px;
    background-color: var(--el-bg-color-overlay);
  }

  .el-statistic {
    --el-statistic-content-font-size: 28px;
  }
}
</style>
