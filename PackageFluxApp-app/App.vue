<script>
	import {
		getToken,
	} from "./store/index.js";
	import {
		skyShowToast
	} from "@/utils/sky.js";
	export default {
		onLaunch: async function() {
			// console.log("App Launch");
			if (!getToken()) {
				uni.redirectTo({
					url: "/pages/auth/login/index",
				});
				skyShowToast("用户信息为空，请先登录");
			}
		},
		onShow: function() {
			console.log("App Show");
		},
		onHide: function() {
			console.log("App Hide");
		}
	};
</script>

<style lang="scss">
	@import "./static/css/index.scss";
	@import "./cl-uni/index.scss";

	// 设置整个项目的背景色
	page {
		font-family: "SkyFont", Helvetica, "PingFang SC", "Microsoft Yahei", sans-serif;
		background-color: #f5f5f5;
		word-break: break-word;
	}

	// 页面状态栏
	#statusBar {
		width: 100%;
		height: v-bind(statusBarHeight) + 'px';
		background-color: #ffffff;
	}

	/* 字体 */
	@font-face {
		font-family: "SkyFont";
		src: url("./static/fonts/Apple.ttf");
		// 字体加载时显示系统字体，字体加载完成后显示自定义字体
		font-display: swap;
	}

	.example-info {
		font-size: 14px;
		color: #333;
		padding: 10px;
	}
</style>