/**
 * 查询当天日期
 */
export const getNowDate = () => {
  const timeOne: Date = new Date()
  const year: number = timeOne.getFullYear()
  let month: any = timeOne.getMonth() + 1
  let day: any = timeOne.getDate()
  month = month < 10 ? '0' + month : month
  day = day < 10 ? '0' + day : day
  return `${year}-${month}-${day}`
}

/**
 * @description 获取assets静态资源
 * @param url
 * @returns
 */
export const getAssets = (url: string) => {
  return new URL(`../assets/${url}`, import.meta.url).href;
};
