import { ElMessage } from "element-plus";
const copy = {
    mounted(el, binding) {
        el.copyData = binding.value;
        el.handleClickEl = async function () {
            try {
                await navigator.clipboard.writeText(el.copyData.toString());
                ElMessage.success("复制成功🌻");
            }
            catch (error) {
                console.error("复制操作不被支持或失败: ", error);
                ElMessage.error("复制失败🌻");
            }
        };
        el.addEventListener("click", el.handleClickEl);
    },
    updated(el, binding) {
        el.copyData = binding.value;
    },
    beforeUnmount(el) {
        el.removeEventListener("click", el.handleClickEl);
    }
};
export default copy;
//# sourceMappingURL=copy.js.map