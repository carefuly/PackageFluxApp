<template>
	<view
		class="cl-radio-group"
		:class="{
			'is-fill': fill,
		}"
	>
		<slot></slot>
	</view>
</template>

<script lang="ts">
import { defineComponent } from "vue";

export default defineComponent({
	name: "cl-radio-group",

	props: {
		// 绑定值
		modelValue: null,
		// 是否禁用
		disabled: Boolean,
		// 是否边框样式
		border: Boolean,
		// 是否带背景色
		bg: Boolean,
		// 是否宽度填充
		fill: Boolean,
		// 是否只显示文字
		text: Boolean,
		// 是否圆角
		round: Boolean,
		// 图标大小
		size: [String, Number],
	},

	emits: ["update:modelValue", "change"],

	setup(_, { emit }) {
		function change(label: any) {
			emit("update:modelValue", label);
			emit("change", label);
		}

		return {
			change,
		};
	},
});
</script>
