<script setup lang="ts">
import Header from "./components/header/index.vue";
import Footer from "./components/footer/index.vue";
</script>

<template>
  <div class="common-layout">
    <header class="layout-header">
      <Header></Header>
    </header>
    <article class="layout-main container">
      <router-view></router-view>
    </article>
    <footer class="layout-footer">
      <Footer></Footer>
    </footer>
  </div>
</template>

<style scoped lang="scss">
.common-layout {
  height: 100vh;
  background: #f2f3f5;
  overflow-y: auto;

  .layout-header {
    position: sticky;
    top: 0;
    height: 64px;
    line-height: 64px;
    background: #ffffff;
    z-index: 999;
  }

  .layout-main {
    //color: var(--color-text-2);
    font-size: 16px;
    font-weight: 400;
    font-stretch: condensed;
  }

  .container {
    margin: 15px auto;
    padding: 10px;
    width: 1000px;
    //background: var(--color-bg-3);
  }

  .layout-footer {
    width: 100%;
    height: 64px;
    //color: var(--color-text-2);
    font-weight: 400;
    font-size: 13px;
    line-height: 50px;
    text-align: center;
  }
}
</style>
