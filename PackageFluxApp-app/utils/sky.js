// 封装showModal
export const skyShowModal = (content, showCancel = false, title = "提示", duration = 3000) => {
	return new Promise((resolve, reject) => {
		uni.showModal({
			title: title,
			content: content,
			showCancel: showCancel,
			success: (res) => {
				resolve(res);
			},
			fail: (err) => {
				reject(err);
			},
			duration: duration,
		});
	});
};
// 封装showToast
export const skyShowToast = (content, icon = "none", duration = 3000) => {
	return new Promise((resolve, reject) => {
		uni.showToast({
			title: content + "🌻",
			icon: icon,
			success: (res) => {
				resolve(res);
			},
			fail: (err) => {
				reject(err);
			},
			duration: duration,
		});
	});
};
// 封装showLoading
export const skyShowLoading = (content, mask = true, duration = 3000) => {
	return new Promise((resolve, reject) => {
		uni.showLoading({
			title: content + "🌻",
			mask: mask,
			success: (res) => {
				resolve(res);
			},
			fail: (err) => {
				reject(err);
			},
			duration: duration,
		});
	});
};

// 封装showActionSheet
export const skyShowActionSheet = (itemList, itemColor, duration = 3000) => {
	return new Promise((resolve, reject) => {
		uni.showActionSheet({
			itemList: itemList,
			itemColor: itemColor,
			success: (res) => {
				resolve(res);
			},
			fail: (err) => {
				reject(err);
			},
			duration: duration,
		});
	});
};