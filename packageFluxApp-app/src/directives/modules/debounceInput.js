const debounce = {
    mounted(el, binding) {
        if (typeof binding.value !== "function") {
            throw "callback must be a function";
        }
        let timer = null;
        const delay = parseInt(binding.arg) || 500;
        const handler = binding.value;
        el.addEventListener("input", () => {
            clearTimeout(timer);
            timer = setTimeout(() => {
                handler();
            }, delay);
        }, { passive: false });
    }
};
export default debounce;
//# sourceMappingURL=debounceInput.js.map