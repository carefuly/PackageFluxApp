import types from "./types";

class accomplish implements types {
  /**
   * 重置表单
   * @param formRef
   * @param object
   */
  resetForm(formRef: any, object: any): void {
    if (formRef) {
      nextTick(() => {
        if (formRef.value) {
          // 重置该表单项，将其值重置为初始值，并移除校验结果
          formRef.value.resetFields();
        }
      }).then(r => {
        JSON.parse(JSON.stringify(object));
      });
    } else {
      JSON.parse(JSON.stringify(object));
    }
  }
}

export default new accomplish();
