<script setup lang="ts">

</script>

<template>
  <div class="sky-card">
    <slot></slot>
  </div>
</template>

<style scoped>
.sky-card {
  padding: 20px;
  background-color: #fff;
  border-radius: 6px;
  border: 1px solid #e4e7ed;
  color: #303133;
  transition: 0.3;
  @apply dark:bg-#1d1e1f dark:c-#cfd3dc dark:b-#414243 flex flex-col flex-1;
  overflow: hidden;
}
</style>
