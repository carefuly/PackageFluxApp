<script>
	export default {
		onLaunch: function() {
			console.warn('当前组件仅支持 uni_modules 目录结构 ，请升级 HBuilderX 到 3.1.0 版本以上！')
			console.log('App Launch')
		},
		onShow: function() {
			console.log('App Show')
		},
		onHide: function() {
			console.log('App Hide')
		}
	}
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

	/* #endif */
	.example-info {
		font-size: 14px;
		color: #333;
		padding: 10px;
	}
</style>