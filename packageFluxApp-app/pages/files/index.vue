<script>
export default {
	data() {
		return {
			visible: false,
			searchValue: '',
			pageParams: {
				creator: null,
				modifier: null,
				status: true,
				appName: ''
			},
			title: '文件管理',
			confirmLoading: false,
			form: {
				id: null,
				recordId: null,
				logoUrl: '',
				appName: '',
				appleId: null,
				description: null,
				preview: [],
				sort: 1,
				status: true,
				belong_dept: null,
				remark: null
			},
			rules: {
				appName: [
					{
						required: true,
						message: '请输入应用名称'
					}
				]
			},
			loading: false,
			total: 0,
			column: [
				{ type: 'selection', fixed: true, width: 60 },
				{ name: 'name', label: '文件名' },
				{ name: 'size', label: '文件大小' },
				{ name: 'address', label: '类型' },
				{ name: 'address', label: '项目根目录' },
				{ name: 'address', label: '文件地址' }
			],
			fileList: [],
			tableList: [{ name: '测试', size: '1.25', suffix: 'image/png', base_dir: '', url: '' }]
		};
	},
	methods: {
		/** 添加 */
		handleAdd() {
			// 打开弹出框
			this.visible = true;
			// 标题
			this.title = '上传';
			// 重置该表单项
			// this.resetForm();
		},

		/** 取消 */
		handleCancel() {
			this.visible = false;
		}
	},
	mounted() {}
};
</script>

<template>
	<view class="page">
		<cl-card label="文件操作">
			<cl-button icon="cl-icon-plus" size="mini" type="primary" @tap="handleAdd">上传文件</cl-button>
		</cl-card>

		<cl-card label="文件列表">
			<next-table
				style="border-left: 1px solid #e8e8e8"
				:show-header="true"
				:border="true"
				align="center"
				:columns="column"
				:stripe="true"
				:fit="false"
				:data="tableList"
			></next-table>
		</cl-card>

		<cl-dialog :title="title" :visible.sync="visible" :close-on-click-modal="false">
			<view class="example-body">
				<uni-file-picker limit="10" file-mediatype="all" title="最多选择10个文件"></uni-file-picker>
			</view>

			<cl-button type="error" @tap="handleCancel">取消</cl-button>
			<!-- <cl-button type="primary" @tap="handleConfirm">保存</cl-button> -->
		</cl-dialog>
	</view>
</template>

<style></style>
