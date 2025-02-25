<template>
	<!-- #ifdef VUE3 -->
	<view class="table-paging" :style="`--primaryColor: ${primaryColor}`">
	<!-- #endif -->
	<!-- #ifdef VUE2 -->
	<view class="table-paging" :style="{'--primaryColor': primaryColor}">
	<!-- #endif -->
    <view
      :class="{
        'table-paging-item': true,
        turner: true,
        disabled: pageIndex <= 1,
      }"
      @click="prePage"
    >
      <text>上一页</text>
    </view>
    <view class="table-paging-item page">
      <text class="table-paging-current">{{ pageIndex || "1" }}</text>
      <text class="table-paging-gap">/</text>
      <text>{{ pageTotal || "0" }}</text>
    </view>
    <view
      :class="{
        'table-paging-item': true,
        turner: true,
        disabled: pageIndex >= pageTotal,
      }"
      @click="nextPage"
    >
      
			<text>下一页</text>
    </view>
  </view>
</template>
<script>
export default {
	emits:['change'],
	props: {
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
		}
	},
	methods: {
		// 上一页
		prePage() {
		  if (this.pageIndex <= 1) {
		    return;
		  }
		  this.$emit("change", this.pageIndex - 1);
		},
		// 下一页
		nextPage() {
		  if (this.pageIndex >= this.pageTotal) {
		    return;
		  }
		  this.$emit("change", this.pageIndex + 1);
		}
	}
}

</script>
<style lang="scss" scoped>
$primary-color: var(--primaryColor);

.table-paging {
  height: 40px;
  display: flex;
  align-items: stretch;
  justify-content: space-between;
  font-size: 14px;
  color: $primary-color;
  line-height: 1;
}
.table-paging-item {
  display: flex;
  align-items: center;
  &.page {
    color: #999999ff;
  }
  &.turner {
    padding: 0 5px;
    transition: all 0.2s linear;

    &:hover {
      opacity: 0.8;
    }
  }
  &.disabled {
    cursor: not-allowed;
    color: #999999;
  }
}
.table-paging-current {
  color: $primary-color;
  font-size: 14px;
  font-weight: 500;
}
.table-paging-gap {
  margin: 0 5px;
}
</style>
