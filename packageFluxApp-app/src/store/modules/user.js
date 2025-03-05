import { defineStore } from "pinia";
import { PINIA_PREFIX } from "@/config";
export const useUserStore = defineStore("user", {
    state: () => {
        return {
            token: "",
            user: {},
        };
    },
    actions: {
        setToken(token) {
            this.token = token;
        },
        setUser(user) {
            this.user = user;
        },
    },
    getters: {},
    persist: {
        key: PINIA_PREFIX + "user", // 默认会以 store 的 id 作为 key
        storage: localStorage
    },
});
//# sourceMappingURL=user.js.map