<script>
	import {
		skyShowToast,
	} from "@/utils/sky.js";
	import {
		captcha,
		register,
	} from "@/apis/auth/register/index.js";
	export default {
		data() {
			return {
				form: {
					username: "",
					email: "",
					password: "",
					code: "",
				},
				pwdShow: false,
				pactChecked: false,
				vcodeTime: 0,
				vcodeTimer: null,
			}
		},
		methods: {
			handleOpen(type, msg, options) {
				this.$refs["message"].open({
					type,
					message: msg,
					...options,
				});
			},
			handleSendCaptcha: async function() {
				if (this.form.email.length <= 0) return skyShowToast("请输入邮箱");
				try {
					const res = await captcha({
						email: this.form.email
					});
					this.handleOpen("success", res.msg);
				} catch (err) {
					this.handleOpen("cancel", JSON.stringify(err.msg));
				} finally {
					this.getVcode();
				}
			},
			handleRegister: async function() {
				try {
					const res = await register(this.form);
					this.handleOpen("success", res.msg);
					setTimeout(() => {
						this.toLogin();
					}, 1000);
				} catch (err) {
					this.handleOpen("cancel", JSON.stringify(err.msg));
				}
			},
			toLogin() {
				uni.redirectTo({
					url: "../login/index"
				});
			},
			getVcode() {
				this.vcodeTime = 60
				this.vcodeTimer = setInterval(() => {
					if (this.vcodeTime > 0) {
						this.vcodeTime--
					} else {
						clearInterval(this.vcodeTimer)
						this.vcodeTimer = null
					}
				}, 1000)
			},
			pactChange() {
				this.pactChecked = !this.pactChecked
			},
			toPact() {
				uni.showToast({
					title: "协议",
					icon: 'none'
				})
			}
		},
		onLoad() {

		},
	}
</script>


<template>
	<view class="page">
		<cl-message ref="message"></cl-message>
		<view class="title">
			<view>您好，一键注册！</view>
		</view>
		<view class="form">
			<view class="inputs">
				<view class="input">
					<input type="text" v-model="form.username" :adjust-position="false" placeholder="用户名" />
				</view>
				<view class="input">
					<input type="text" v-model="form.email" :adjust-position="false" placeholder="邮箱" />
				</view>
				<view class="input">
					<input type="safe-password" v-model="form.password" :adjust-position="false" placeholder="密码" />
				</view>
				<view class="vcode">
					<input type="number" v-model="form.code" :adjust-position="false" placeholder="验证码" />
					<text v-if="vcodeTime==0" style="color:#ffde66" @click="handleSendCaptcha">获取验证码</text>
					<text v-else style="color:#ccc">重新获取({{vcodeTime}}s)</text>
				</view>
			</view>
			<view class="button">
				<view @click="handleRegister">
					立即注册
				</view>
			</view>
			<view class="tip">
				<text @click="toLogin">已有帐号？去登录</text>
			</view>
		</view>
		<view class="pact">
			<checkbox :checked="pactChecked" activeBackgroundColor="#FFD640" activeBorderColor="#FFD640"
				borderColor="#000000" color="#fff" @click="pactChange" />
			<view>我已阅读并同意<text @click="toPact">《用户协议》</text>和<text @click="toPact">《隐私协议》</text></view>
		</view>
	</view>
</template>

<style lang="scss" scoped>
	.page {
		min-height: 100vh;
		display: flex;
		flex-direction: column;
		justify-content: space-between;

		.title {
			padding-top: 500rpx;
			padding-bottom: 120rpx;
			background: url("../../../static/images/pageBg.png") no-repeat top center;
			background-size: 100%;

			view {
				padding-left: 64rpx;
				font-size: 44rpx;
				font-weight: 700;
				line-height: 1;
			}
		}

		.form {
			flex: 1;
			padding: 0 64rpx;

			.inputs {
				margin-bottom: 100rpx;

				view {
					height: 75rpx;
					border-radius: 45rpx;
					border: 1rpx solid rgba(204, 204, 204, 1);
					padding: 0 48rpx;
					display: flex;
					align-items: center;

					input {
						flex: 1;
					}
				}

				.input {
					margin-bottom: 20rpx;
				}

				.vcode {
					text {
						text-wrap: nowrap;
						font-size: 26rpx;
					}
				}
			}

			.button {
				view {
					line-height: 96rpx;
					border-radius: 48rpx;
					text-align: center;
					font-size: 32rpx;
					background: linear-gradient(90deg, rgba(255, 222, 102, 1) 0%, rgba(202, 245, 253, 1) 100%);
				}
			}

			.tip {
				font-size: 28rpx;
				text-align: center;
				margin-top: 30rpx;
			}
		}

		.pact {
			display: flex;
			align-items: center;
			justify-content: center;
			padding-bottom: 70rpx;
			font-size: 24rpx;

			text {
				color: #ffde66;
			}

			checkbox {
				:deep(.uni-checkbox-input) {
					border: 1rpx solid #000000;
				}

				transform:scale(0.6)
			}
		}
	}
</style>