<template>
  <div class="table-empty">
		<text>{{text}}</text>
  </div>
</template>
<script>
	export default {
	  props: {
			text: {
				type: String,
				default: '暂无数据'
			}
		}
	}
</script>
<style lang="scss" scoped>
.table-empty {
	width: 100%;
	height: 80rpx;
	display: flex;
	justify-content: center;
	align-items: center;
	border-bottom: 1px solid #e8e8e8;
}
</style>
