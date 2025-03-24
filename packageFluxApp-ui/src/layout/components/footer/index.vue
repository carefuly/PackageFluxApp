<script setup>
const pageData = ref({
  url: "https://package.flechazo.net",
  copy: "https://beian.miit.gov.cn/#/Integrated/index",

});
const method = reactive({
  goto: () => {
    window.open(pageData.value.url);
  },
  copy: () => {
    window.open(pageData.value.copy);
  },
});
</script>

<template>
  <div class="copyright">
    本网站由
    <span @click="method.goto" style="color: #afaeae; cursor: pointer;">PackageFluxApp | </span>
    <span @click="method.copy()"
          style="color: #afaeae; cursor: pointer;">蜀ICP备2025122391号-2</span>
    提供技术支持
    <br>
  </div>
</template>

<style scoped>
.copyright {
  margin: 0 auto;
}
</style>
