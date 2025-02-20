<template>
	<view class="container">
		<cl-button type="primary">主要</cl-button>
		<cl-button type="success">成功</cl-button>
		<cl-button type="error">错误</cl-button>
		<cl-button type="warning">警告</cl-button>
		<cl-button type="info">信息</cl-button>
		<cl-button type="text">文本</cl-button>
	</view>
</template>

<script>
	export default {
		data() {
			return {

			}
		},
		methods: {

		}
	}
</script>

<style>
	.container {
		padding: 20px;
		font-size: 14px;
		line-height: 24px;
	}
</style>