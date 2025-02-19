import { App } from "vue";
import debounce from "./modules/debounce";
import debounceInput from "./modules/debounceInput";
import throttle from "./modules/throttle";
import throttleInput from "./modules/throttleInput";
import copy from "./modules/copy";

const directivesList: any = {
  debounce,
  debounceInput,
  throttle,
  throttleInput,
  copy
};

const directives = {
  install: function (app: App<Element>) {
    Object.keys(directivesList).forEach(key => {
      // 注册所有自定义指令
      app.directive(key, directivesList[key]);
    });
  }
};

export default directives;
