<template>
	<view class="container">
		<button class="mini-btn" type="primary" size="mini" @click="onHandleUpdate">检查更新</button>
		<!-- 插件 -->
		<package-version-update
			ref="updateRef"
			id="8d635dda-a234-49f3-b60c-667f7bcb6c77"
			:auto="auto"
			uniRemark="详情页面测试更新"
			:uniqueId="'13036436228'"
			@finish="onHandleFinish"
		/>
	</view>
</template>

<script>
export default {
	data() {
		return {
			auto: false,
			version: ''
		};
	},
	methods: {
		onHandleUpdate() {
			this.$refs.updateRef.check();
		},
		onHandleFinish(e) {
			uni.showToast({
				title: e.tip
			});
		}
	}
};
</script>

<style scoped lang="scss">
.container {
	width: 100%;
	height: 100%;
	display: flex;
	align-items: center;
	justify-content: center;

	.mini-btn {
		margin-top: 100rpx;
	}
}
</style>
