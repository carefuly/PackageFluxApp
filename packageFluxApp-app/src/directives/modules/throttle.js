const throttle = {
    mounted(el, binding) {
        let timer = null;
        const delay = parseInt(binding.arg) || 500;
        const handler = binding.value;
        let lastExecTime = 0;
        el.addEventListener("click", () => {
            const now = Date.now();
            if (now - lastExecTime > delay) {
                handler();
                lastExecTime = now;
            }
        }, { passive: false });
    }
};
export default throttle;
//# sourceMappingURL=throttle.js.map