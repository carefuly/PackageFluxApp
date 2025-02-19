# package-version-update

~~~
--- 全新升级
~~~

> 1.一行代码实现APP版本更新，更提供后台管理让版本的发布和迭代再也不用愁了
>
> 2.支持安卓apk安装更新、支持安卓和ios热更新、支持ios跳转到appStore更新
>
> 3.支持配置非强制更新，暴露绝大部分参数自定义性强
>
> 4.封装了大部分js方法，让向完全自定义界面的开发者更加便捷
>
> 5.提供后台管理 [PackageHotApp](https://package.coisink.com/)，轻松完成版本更新
>
> 6.提供发布页，效果可查看案例 [Package发布页](https://package.coisink.com/#/app/1)
>
> 7.提供后台有偿私有化部署，具体可加群联系群主


## 一、后台管理添加应用

### 1.登录网址

[点我前往PackageAppStore后台管理](https://package.coisink.com/)，后台提供版本更新服务，注册登录即可。

### 2.创建应用

![](https://package.coisink.com/media/upload/admin@qq.com/添加应用.png)

### 3.添加版本

创建完应用后，点击菜单栏的版本管理，点击添加版本即可。其中HBuilderX是否更新和是否强制更新底部皆有解释，若还不明白的可添加QQ群讨论咨询。

![](https://package.coisink.com/media/upload/admin@qq.com/新增版本.png)

### 4.设置正式版

在添加的版本菜单栏，选择设为正式版即可。

![](https://package.coisink.com/media/upload/admin@qq.com/确定发布正式版.png)

## 二、使用

### 1.获取应用id

![](https://package.coisink.com/media/upload/admin@qq.com/获取应用id.png)

### 2.引入组件，修改id值

将插件导入项目，在首页引入组件，id**记得**替换为上一步获取到的

```vue
<!-- 插件 -->
remark: 备注个人建议写路由地址，方便排查文件
<package-version-update id="1" remark="测试更新APP" @check="onHandleCheck" @finish="onHandleFinish"></package-version-update>
// 方法
onHandleCheck(e) {
	this.version = e;
	if (e.needUpdate) {
		// 隐藏tabBar
		uni.hideTabBar();
	}
}
// 手动检查更新
onHandleUpdate() {
	this.$refs.updateRef.check();
},
onHandleFinish(e) {
	uni.showToast({
		title: e.tip
	})
}
```

### 3.监听事件

当需要更新会自动弹窗，当无需更新会触发`finish`事件，此时可表示应用为最新版，可由此事件继续应用的相关业务逻辑。

- @check (version)：无论是否需要更新都会触发事件，传入版本参数
- @finish (version): 当无需更新会触发事件，传入版本参数
- @error (err): 当发生错误会触发事件，传入错误信息