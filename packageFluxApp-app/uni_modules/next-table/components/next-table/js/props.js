export default {
	props: {
		highlight: {
			type: Boolean,
			default: false
		},
		itemDate: {
			type: Object,
			default: () => {}
		},
		columns: {
			type: Array,
			default: () => []
		},
		showSummary: {
			type: Boolean,
			default: false
		},
		isShowLoadMore: {
			type: Boolean,
			default: false
		},
		data: {
			type: Array,
			default: () => []
		},
		sumText: {
			type: String,
			default: '合计'
		},
		showHeader: {
			type: Boolean,
			default: true
		},
		border: {
			type: Boolean,
			default: false
		},
		stripe: {
			type: Boolean,
			default: true
		},
		fit: {
			type: Boolean,
			default: false
		},
		showPaging: {
			type: Boolean,
			default: false
		},
		pageIndex: {
			type: Number,
			default: 1
		},
		pageTotal: {
			type: Number,
			default: 0
		},
		primaryColor: {
			type: String,
			default: '#f0ad4e'
		},
		rowKey: [String, Function],
		summaryMethod: Function,
		pullUpLoading: Function,
		formatter: Function,
		cellStyle: Function,
		cellHeaderStyle: Function,
		permissionBtn: Function,
	}
}
