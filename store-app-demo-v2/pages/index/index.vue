<template>
	<view class="content">
		<image class="logo" src="/static/logo.png"></image>
		<!-- 标题 -->
		<h2>APP在线热更新插件</h2>
		<!-- 参数 -->
		<h4>更新测试参数</h4>
		<!-- 插件 -->
		<package-version-update id="8d635dda-a234-49f3-b60c-667f7bcb6c77" uniRemark="首页检查更新" :uniqueId="'13036436228'" @check="onHandleCheck" />
	</view>
</template>

<script>
export default {
	data() {
		return {
			title: 'Hello'
		};
	},
	onLoad() {},
	methods: {
		onHandleCheck() {}
	}
};
</script>

<style>
.content {
	display: flex;
	flex-direction: column;
	align-items: center;
	justify-content: center;
}

.logo {
	height: 200rpx;
	width: 200rpx;
	margin-top: 200rpx;
	margin-left: auto;
	margin-right: auto;
	margin-bottom: 50rpx;
}

.text-area {
	display: flex;
	justify-content: center;
}

.title {
	font-size: 36rpx;
	color: #8f8f94;
}
</style>
