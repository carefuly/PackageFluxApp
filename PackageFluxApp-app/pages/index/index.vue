<script>
	import {
		getUser
	} from "@/store/index.js";
	import {
		add,
		deleteById,
		updateById,
		getById,
		listAll,
	} from "@/apis/application/details/index.js";
	import {
		skyShowToast,
	} from "@/utils/sky.js";

	export default {
		data() {
			return {
				userinfo: {},
				searchValue: "",
				pageParams: {
					creator: null,
					modifier: null,
					status: true,
					appName: "",
				},
				title: "应用详情",
				confirmLoading: false,
				form: {
					id: null,
					recordId: null,
					logoUrl: "",
					appName: "",
					appleId: null,
					description: null,
					preview: [],
					sort: 1,
					status: true,
					belong_dept: null,
					remark: null,
				},
				rules: {
					appName: [{
						required: true,
						message: "请输入应用名称",
						trigger: "blur"
					}],
				},
				loading: false,
				total: 0,
				fileList: [],
				tableList: [],
			};
		},
		methods: {
			/** 消息提示 */
			handleOpen(type, msg, options) {
				this.$refs["message"].open({
					type,
					message: msg,
					...options,
				});
			},

			/** 删除 */
			handleDelete(row) {
				const id = row.recordId;
				if (id === null || id === "") {
					skyShowToast("请选中需要删除的数据");
					return;
				}
				// skyMsgBox("您确认需要删除名称[" + row.appName + "]么？")
				//   .then(async () => {
				//     try {
				//       await deleteById(id);
				//       await method.handleListAll();
				//       skyNoticeSuccess("删除成功🌻");
				//     } catch (error) {
				//       skyNoticeError("删除失败，请刷新重试🌻");
				//     }
				//   })
				//   .catch(() => {
				//     skyMsgError("已取消🌻");
				//   });
			},

			/** 数据表格 */
			async handleListAll() {
				try {
					this.loading = true;
					this.tableList = [];
					const res = await listAll(this.pageParams);
					this.tableList = res.data;
				} catch (error) {
					skyShowToast("数据查询失败，请刷新重试");
				} finally {
					this.loading = false;
				}
			},
		},
		computed: {
			filterList() {
				if (this.searchValue) {
					return this.tableList.filter(item => item.appName.includes(this.searchValue));
				} else {
					return this.tableList;
				}
			}
		},
		mounted() {
			this.userinfo = JSON.parse(getUser());
			this.handleListAll();
		}
	};
</script>

<template>
	<view class="page">

		<cl-message ref="message"></cl-message>

		<cl-card label="温馨提示">
			<cl-text value="有任何问题请加作者沟通(备注)，避免造成应用发布后无法更新的损失" color="primary"></cl-text>
			<cl-text value="已上线储桶功能，每个账号存储桶暂无限制" color="warning"></cl-text><br />
			<cl-text :value="'当前账号使用:' + userinfo.usageNumber + 
			'次, 剩余次数:' + userinfo.total" color="warning"></cl-text>
		</cl-card>

		<cl-card label="操作">
			<view style="display: flex; align-items: center; justify-content: space-evenly;" class="">
				<cl-button icon="cl-icon-plus" size="mini" type="primary">新增</cl-button>
				<cl-search style="width: 75%;" v-model="searchValue" placeholder="请输入应用名称"
					:show-search-button="false"></cl-search>
			</view>
		</cl-card>

		<view class="space">
			<cl-loading-mask :loading="this.loading" text="数据加载中" />

			<cl-card class="cl-card" v-if="filterList.length > 0" :label="item.appName" v-for="item in filterList"
				:key="item.id">
				<view style="justify-content: space-between;" class="app-header center">
					<view class="center">
						<cl-image :src="item.logoUrl" :size="70"></cl-image>
						<cl-text :value="item.appName" :margin="[0, 0, 0, 10]"></cl-text>
					</view>
					<view class="operate">
						<cl-button size="mini" type="warning" plain>统计</cl-button>
						<cl-button size="mini" type="primary" plain>版本</cl-button>
						<cl-button size="mini" type="success" plain>修改</cl-button>
						<cl-button size="mini" type="error" plain @tap="handleDelete(item)">删除</cl-button>
					</view>
				</view>
				<view class="text">
					<cl-text :value="'应用id: ' + item.recordId" color="primary"></cl-text><br />
					<cl-text :value="'appleId: ' + item.appleId" color="success"></cl-text><br />
					<cl-text :value="'应用描述: ' + item.description" color="error"></cl-text><br />
					<cl-text :value="'备注: ' + item.remark" color="info"></cl-text><br />
					<cl-text :value="'修改时间: ' + item.updateTime" color="warning"></cl-text><br />
				</view>
			</cl-card>

			<cl-card v-else label="应用">
				<view class="center message">
					<h1>应用为空</h1>
					<p>请通过左上角的新增按钮添加应用</p>
				</view>
			</cl-card>
		</view>

		<view style="height: 10rpx;"></view>
	</view>
</template>

<style lang="scss">
	.space {
		margin-bottom: 10px;
		font-size: 28rpx;
		width: 100%;
		background-color: #f7f7f7;
		position: relative;
	}

	.app-header {}

	.text {
		margin-top: 10rpx;
	}

	.message {
		height: 700rpx;
		flex-direction: column;
	}

	.center {
		display: flex;
		align-items: center;
		justify-content: center;
	}
</style>