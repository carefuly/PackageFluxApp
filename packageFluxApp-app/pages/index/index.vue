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
				visible: false,
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
			/** 添加 */
			handleAdd() {
				// 打开弹出框
				this.visible = true;
				// 标题
				this.title = "添加";
				// 重置该表单项
				this.resetForm();
			},
			/** 删除 */
			handleDelete(row) {
				const that = this;
				const id = row.recordId;
				if (id === null || id === "") {
					skyShowToast("请选中需要删除的数据");
					return;
				}
				this.$refs["confirm"]?.open({
					title: "温馨提示",
					message: `您确认需要删除名称` + row.appName + `么？`,
					async callback(action) {
						switch (action.action) {
							case "confirm":
								try {
									await deleteById(id);
									await that.handleListAll();
									that.handleOpen("success", "删除成功🌻");
								} catch (error) {
									console.log(error);
									that.handleOpen("cancel", "删除失败，请刷新重试🌻");
								}
								break;
							case "cancel":
								break;
							case "close":
								break;
						}
					}
				});
			},
			/** 修改 */
			handleUpdate(row) {
				// 标题
				this.title = "修改";
				// 文件数据表格
				// method.handleFileListAll();
				// 重置表单
				this.resetForm();
				const id = row.recordId;
				if (id == null || id === "") {
					skyShowToast("请选中需要修改的数据🌻");
				}
				// 回显数据
				this.handleEcho(id);
				// 打开弹出框
				this.visible = true;
			},
			/** 回显数据 */
			async handleEcho(id) {
				if (id === null || id === "") {
					skyShowToast("请选择需要修改的数据");
					return;
				}
				try {
					const res = await getById(id);
					this.form = res.data;
				} catch (error) {
					this.handleOpen("cancel", "数据获取失败，请刷新重试🌻");
				}
			},
			// 重置该表单项，将其值重置为初始值，并移除校验结果
			resetForm() {
				this.form = {
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
				};
			},
			/** 确定  */
			handleConfirm() {
				this.confirmLoading = true;
				this.$refs["formRef"].validate(async (valid) => {
					if (valid) {
						if (this.form.id) {
							try {
								await updateById(this.form.recordId, this.form);
								this.handleOpen("success", "修改成功🌻");
								this.confirmLoading = false;
								this.resetForm();
								await this.handleListAll();
								this.handleCancel();
							} catch (error) {
								this.confirmLoading = false;
								this.handleOpen("cancel", "修改失败，请刷新重试🌻");
							}
						} else {
							try {
								await add(this.form);
								this.handleOpen("success", "添加成功🌻");
								this.confirmLoading = false;
								this.resetForm();
								await this.handleListAll();
								this.handleCancel();
							} catch (error) {
								this.confirmLoading = false;
								this.handleOpen("cancel", "添加失败，请刷新重试🌻");
							}
						}
					} else {
						this.handleOpen("cancel", "验证失败，请检查填写内容🌻");
						this.confirmLoading = false;
					}
				});
			},
			/** 取消 */
			handleCancel() {
				this.visible = false;
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
		<cl-confirm ref="confirm"></cl-confirm>

		<cl-card label="温馨提示">
			<cl-text value="有任何问题请加作者沟通(备注)，避免造成应用发布后无法更新的损失" color="primary"></cl-text>
			<cl-text value="已上线储桶功能，每个账号存储桶暂无限制" color="warning"></cl-text><br />
			<cl-text :value="'当前账号使用:' + userinfo.usageNumber + 
			'次, 剩余次数:' + userinfo.total" color="warning"></cl-text>
		</cl-card>

		<cl-card label="操作">
			<view style="display: flex; align-items: center; justify-content: space-evenly;" class="">
				<cl-button icon="cl-icon-plus" size="mini" type="primary" @tap="handleAdd">新增</cl-button>
				<cl-button size="mini" type="error" @tap="handleListAll">刷新</cl-button>
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
						<cl-button size="mini" type="success" plain @tap="handleUpdate(item)">修改</cl-button>
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

		<cl-dialog :title="title" :visible.sync="visible" :close-on-click-modal="false">
			<cl-form ref="formRef" :model="form" :rules="rules" label-position="right" label-width="140rpx">
				<cl-form-item label="应用logo" prop="logoUrl">
					<cl-input v-model="form.logoUrl" placeholder="将图片url粘贴到此处" clearable></cl-input>
				</cl-form-item>
				<cl-form-item label="应用名称" prop="appName">
					<cl-input v-model="form.appName" placeholder="应用名请勿包含空格" clearable></cl-input>
				</cl-form-item>
				<cl-form-item label="appleId" prop="appleId">
					<cl-input v-model="form.appleId" placeholder="App Store下的app信息" clearable></cl-input>
				</cl-form-item>
				<cl-form-item label="应用描述" prop="description">
					<cl-textarea v-model="form.description" placeholder="请输入应用描述"></cl-textarea>
				</cl-form-item>
				<cl-form-item label="预览截图" prop="preview">
					<cl-select border v-model="form.preview" :options="fileList" placeholder="请选择预览图"></cl-select>
				</cl-form-item>
				<cl-form-item label="备注" prop="remark">
					<cl-textarea v-model="form.remark" placeholder="请输入备注"></cl-textarea>
				</cl-form-item>
			</cl-form>
			<cl-button type="error" @tap="handleCancel">取消</cl-button>
			<cl-button type="primary" @tap="handleConfirm">保存</cl-button>
		</cl-dialog>
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