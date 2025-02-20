# lyzml-weixin
### 开发文档

### uvue页面中使用
```javascript
	import * as LyzmlWeiXin from "@/uni_modules/lyzml-weixin";
  import { WeiXinAuthParam } from "@/uni_modules/lyzml-weixin";
	
	LyzmlWeiXin.auth(authParam, (resp : UTSJSONObject) => {
		 
	});
	
	LyzmlWeiXin.stopAuth();
```

```uts
authParam的类型必须为WeiXinAuthParam,该参数需要通过后台接口生成返回
export type WeiXinAuthParam = {
	appId : string;
	scope : string;
	nonceStr : string;
	timeStamp : string;
	signature : string;
};
 
resp为授权返回结果
{
	"code": 0,
	"msg": "onAuthGotQrcode",
	"result": {
		"imgBase64Str": "base64格式图片字符串，可直接赋值给img.src"
	}
}

{
	"code": 1,
	"msg": "onQrcodeScanned"        //已使用微信扫码
}

{
	"code": 2,
	"msg": "onAuthFinish",
	"result": {
		"authCode": "微信授权成功后返回的授权码",
		"errCode": "授权成功时为0，其他为授权失败",
		"errMsg": "授权失败时返回的错误信息"
	}
}
```