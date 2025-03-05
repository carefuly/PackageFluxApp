/**
 * 查询当天日期
 */
export const getNowDate = () => {
    const timeOne = new Date();
    const year = timeOne.getFullYear();
    let month = timeOne.getMonth() + 1;
    let day = timeOne.getDate();
    month = month < 10 ? '0' + month : month;
    day = day < 10 ? '0' + day : day;
    return `${year}-${month}-${day}`;
};
/**
 * @description 获取assets静态资源
 * @param url
 * @returns
 */
export const getAssets = (url) => {
    return new URL(`../assets/${url}`, import.meta.url).href;
};
//# sourceMappingURL=index.js.map