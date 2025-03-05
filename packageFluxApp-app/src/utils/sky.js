// 工具类提示信息
import { ElNotification, ElMessageBox, ElMessage } from "element-plus";
/** 封装任意提示类型通知，默认info */
export function skyNotice(message, title = "温馨提示", duration = 2000, type = "info", parseHtml = false) {
    ElNotification.closeAll();
    ElNotification({
        message,
        title,
        type,
        duration: duration,
        showClose: true,
        dangerouslyUseHTMLString: parseHtml
    });
}
/** 封装提示通知，默认success */
export function skyNoticeSuccess(message, title = "温馨提示", duration = 2000, type = "success", parseHtml = false) {
    ElNotification.closeAll();
    ElNotification({
        message,
        type,
        title,
        duration: duration,
        showClose: true,
        dangerouslyUseHTMLString: parseHtml
    });
}
/** 封装提示通知，默认error */
export function skyNoticeError(message, title = "温馨提示", duration = 2000, type = "error", parseHtml = false) {
    ElNotification.closeAll();
    ElNotification({
        message,
        type,
        title,
        duration: duration,
        showClose: true,
        dangerouslyUseHTMLString: parseHtml
    });
}
/** 封装提示通知，默认warning */
export function skyNoticeWarning(message, title = "温馨提示", duration = 2000, type = "warning", parseHtml = false) {
    ElNotification.closeAll();
    ElNotification({
        message,
        title,
        type,
        duration: duration,
        showClose: true,
        dangerouslyUseHTMLString: parseHtml
    });
}
/** 封装提示通知，默认info */
export function skyNoticeInfo(message, title = "温馨提示", duration = 2000, type = "info", parseHtml = false) {
    ElNotification.closeAll();
    ElNotification({
        message,
        title,
        type,
        duration: duration,
        showClose: true,
        dangerouslyUseHTMLString: parseHtml
    });
}
/** 封装提示信息，默认info */
export function skyMsg(message, duration = 2000, type = "info", parseHtml = false) {
    ElMessage.closeAll();
    ElMessage({
        message,
        type,
        plain: true,
        duration: duration,
        showClose: true,
        dangerouslyUseHTMLString: parseHtml
    });
}
/** 封装提示信息，默认success */
export function skyMsgSuccess(message, duration = 2000, type = "success", parseHtml = false) {
    ElMessage.closeAll();
    ElMessage({
        message,
        type,
        plain: true,
        duration: duration,
        showClose: true,
        dangerouslyUseHTMLString: parseHtml
    });
}
/** 封装提示信息，默认error */
export function skyMsgError(message, duration = 2000, type = "error", parseHtml = false) {
    ElMessage.closeAll();
    ElMessage({
        message,
        type,
        plain: true,
        duration: duration,
        showClose: true,
        dangerouslyUseHTMLString: parseHtml
    });
}
/** 封装提示信息，默认warning */
export function skyMsgWarning(message, duration = 2000, type = "warning", parseHtml = false) {
    ElMessage.closeAll();
    ElMessage({
        message,
        type,
        plain: true,
        duration: duration,
        showClose: true,
        dangerouslyUseHTMLString: parseHtml
    });
}
/** 封装提示信息，默认info */
export function skyMsgInfo(message, duration = 2000, type = "info", parseHtml = false) {
    ElMessage.closeAll();
    ElMessage({
        message,
        type,
        plain: true,
        duration: duration,
        showClose: true,
        dangerouslyUseHTMLString: parseHtml
    });
}
/** 封装确认信息，默认warning */
export function skyMsgBox(message = "您确定进行关闭么？", title = "温馨提示：", confirmButtonText = "确定", cancelButtonText = "取消", type = "warning") {
    return new Promise((resolve, reject) => {
        ElMessageBox.confirm(message, title, {
            confirmButtonText,
            cancelButtonText,
            type,
            draggable: true,
            dangerouslyUseHTMLString: true
        })
            .then(() => {
            resolve(true);
        })
            .catch(() => {
            reject(false);
        });
    });
}
/** 封装确认信息，默认warning  */
export function skyMsgBoxHtml(message = `<p style="color: teal">您确定进行关闭么？</p>`, title = "温馨提示：", confirmButtonText = "确定", cancelButtonText = "取消", type = "warning") {
    return new Promise((resolve, reject) => {
        ElMessageBox.confirm(message, title, {
            confirmButtonText,
            cancelButtonText,
            type,
            draggable: true,
            dangerouslyUseHTMLString: true
        })
            .then(() => {
            resolve(true);
        })
            .catch(() => {
            reject(false);
        });
    });
}
/** Prompt 类型的消息框 */
export function skyMsgBoxPrompt(message = "请输入需要修改的数据？", title = "温馨提示：", confirmButtonText = "确定", cancelButtonText = "取消", type = "info", inputPattern = "", inputErrorMessage = "无效输入") {
    return new Promise((resolve, reject) => {
        ElMessageBox.prompt(message, title, {
            confirmButtonText: confirmButtonText,
            cancelButtonText: cancelButtonText,
            type,
            inputPattern: inputPattern,
            inputErrorMessage: inputErrorMessage,
            draggable: true
        })
            .then((res) => {
            // 返回值获取通过[res.value]
            resolve(res);
        })
            .catch(() => {
            reject(false);
        });
    });
}
/** Alert 类型的消息框 */
export function skyMsgBoxAlert(message = "请输入需要修改的数据？", title = "温馨提示：", confirmButtonText = "确定", type = "info") {
    return new Promise((resolve, reject) => {
        ElMessageBox.alert(message, title, {
            confirmButtonText: confirmButtonText,
            type,
            draggable: true
        })
            .then(() => {
            // 返回值获取通过[res.value]
            resolve(true);
        })
            .catch(() => {
            reject(false);
        });
    });
}
//# sourceMappingURL=sky.js.map