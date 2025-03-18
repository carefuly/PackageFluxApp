<script setup lang="ts">
import {useRoute} from "vue-router";
import {getIssue} from "@/apis/application/issue";
import {getById} from "@/apis/application/details";
import {getById as getVersion} from "@/apis/application/version";
import {skyMsgError} from "@/utils/sky";

const route = useRoute();
const pageData = ref({
  url: "https://package.coisink.com",
  copy: "https://beian.miit.gov.cn/#/Integrated/index",
  type: true,
  error: "",
  equipment: "",
  form: {
    logo: "",
    appName: "",
    appVersion: {
      description: "",
    },
    preview: [],
  },
  versionForm: {},
  description: "",
});
const method = reactive({
  handleGetAPPID: (route: any) => {
    const id = route.path;
    return id.split('/')[id.split('/').length - 1];
  },
  handleGetIssueObj: (id: any) => {
    getById(id).then((res: any) => {
      if (res.code === 200) {
        pageData.value.form = res.data;
        if (!pageData.value.form['version_id']) {
          pageData.value.type = false;
          pageData.value.error = "当前应用未发布正式版";
        } else {
          getVersion(pageData.value.form['version_id'], {
            detailId: pageData.value.form["id"],
          }).then(res => {
            if (res['code'] === 200) {
              pageData.value.versionForm = res.data;
              pageData.value.description = res.data.description;
            }
          });
        }
      }
    }).catch(err => {
      skyMsgError(err + "🌻");
      pageData.value.type = false;
      pageData.value.error = err;
    });
  },
  handleDownloadFile: () => {
  },
  goto: () => {
    window.open(pageData.value.url);
  },
  copy: () => {
    window.open(pageData.value.copy);
  },
});
onMounted(() => {
  const id = method.handleGetAPPID(route);
  method.handleGetIssueObj(id);
});
</script>

<template>
  <div v-if="pageData.type">
    <div class="app-box">
      <div class="flex-row">
        <div class="logo-box">
          <img class="logo-img" :src="pageData.form['logoUrl']" :alt="pageData.form.appName">
        </div>
        <div style="display: flex" class="info-box flex-column flex-1 justify-between">
          <div>
            <div class="flex-row justify-between">
              <div class="info-name">{{ pageData.form.appName }}</div>
              <el-button>在手机端打开</el-button>
            </div>
          </div>
          <div>
            <a :href="pageData.versionForm['apkUrl']">
              <el-button type="primary" style="border-radius: 10px" @click="method.handleDownloadFile">
                立即下载
              </el-button>
            </a>
          </div>
        </div>
      </div>
      <el-divider></el-divider>
      <div class="title">更新内容</div>
      <div class="flex-row justify-between">
        <ol style="list-style: none;" class="lineH-right">
          <li v-if="!pageData.description">
            <span>暂无描述信息</span>
          </li>
          <li v-else v-for="item in pageData.description.split('\n')">
            <span>{{ item }}</span>
          </li>
        </ol>
      </div>
      <el-divider></el-divider>
      <div class="title">预览</div>
      <div class="preview">
        <el-image
          class="mr5"
          style="width: 220px; height: 370px;"
          v-for="item in pageData.form.preview"
          :key="item"
          :src="item"
          :zoom-rate="1.2"
          :max-scale="7"
          :min-scale="0.2"
          :initial-index="4"
          :preview-src-list="pageData.form.preview"
          fit="cover"
        />
      </div>
      <el-divider></el-divider>
      <div class="copyright">
        本发布页
        <span @click="method.goto" style="color: #afaeae; cursor: pointer;">PackageHotApp | </span>
        <span @click="method.copy()"
              style="color: #afaeae; cursor: pointer;">蜀ICP备2023008183号-4</span>
        提供技术支持
        <br>
      </div>
    </div>
  </div>
  <div v-else class="flex error">
    <h1>{{ pageData.error }}</h1>
  </div>
</template>

<style scoped lang="scss">
.flex {
  display: flex;
  align-items: center;
  justify-content: center;
}

.app-box {
  --app-box-padding: 26px;
  --logo-size: 150px;
  --info-padding: 20px;
  --app-name-size: 24px;
  --app-slogan-size: 16px;
  --app-slogan-pt: 8px;
  --title-size: 20px;
  --screenshot-width: 340px;
  --screenshot-height: 400px;
  margin: 0 auto;
  padding: var(--app-box-padding);
  max-width: 800px;

  .flex-1 {
    flex: 1 !important;
  }

  .justify-between {
    justify-content: space-between !important;
  }

  .flex-column {
    flex-direction: column !important;
  }

  .flex-row {
    display: flex;

    .logo-box {
      margin: var(--info-padding) 0 var(--info-padding) var(--info-padding);
      display: flex;

      .logo-img {
        border-radius: 20px;
        width: var(--logo-size);
        height: var(--logo-size);
      }
    }

    .info-name {
      font-size: var(--app-name-size);
      font-weight: 700;
    }
  }

  .title {
    margin-bottom: 17px;
    font-size: var(--title-size);
    font-weight: 700;
  }

  .preview {
    width: 100%;
    height: 390px;
    white-space: nowrap;
    overflow-x: auto;

    img {
      white-space: nowrap;
      display: inline-block;
      width: 200px;
      height: 95%;
    }
  }

  .copyright {
    margin-top: 24px;
    padding-top: 8px;
    padding-bottom: 8px;
    text-align: center;
    font-size: 12px;
    color: var(--color-text-4);
  }
}

.error {
  width: 100%;
  height: 100vh;

  h1 {
    font-size: 45px;
  }
}
</style>
