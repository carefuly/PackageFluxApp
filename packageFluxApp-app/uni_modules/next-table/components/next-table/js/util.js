/**
 * 获取滚动条宽度
 */
let cached = undefined;

export const getScrollbarSize = fresh => {
	// #ifdef H5
	
    if (fresh || cached === undefined) {
        let inner = document.createElement("div");
        let innerStyle = inner.style;

        innerStyle.width = "100%";
        innerStyle.height = "200px";

        let outer = document.createElement("div");
        let outerStyle = outer.style;

        outerStyle.position = "absolute";
        outerStyle.top = 0;
        outerStyle.left = 0;
        outerStyle.pointerEvents = "none";
        outerStyle.width = "200px";
        outerStyle.height = "150px";
        outerStyle.visibility = "hidden";

        outer.appendChild(inner);
        document.body.appendChild(outer);

        // 设置子元素超出部分隐藏
        outerStyle.overflow = "hidden";

        let width1 = inner.offsetWidth;

        // 设置子元素超出部分滚动
        outer.style.overflow = "scroll";

        let width2 = inner.offsetWidth;

        if (width1 === width2) {
            width2 = outer.clientWidth;
        }

        document.body.removeChild(outer);

        cached = width1 - width2;
    }
	//#endif

    return cached;
};

// 16进制转换rgba
export const colorHextoRGB = (val, opacity = 0.7) => {
	let color = val;
	const t = {},
					bits = (color.length == 4) ? 4 : 8,//假设是shorthand。 #fff, 那么bits为4位, 每一位代表的个属性, 其他的为8位 每两位代表一个属性 #ffffff00
					mask = (1 << bits) - 1; //表示字节占位符。 向左移4位或8位，var a = (1 << 4 ) - 1 -> 10000 - 1,  a.toString(2); // 1111。或者 8位的 1111 1111
	color = Number("0x" + color.substr(1)); //#ff0000 转变为16进制0xff0000;
	if(isNaN(color)){
			return null; // Color
	}
	["b", "g", "r"].forEach(function(x){
			const c = color & mask;
			color >>= bits;
			t[x] = bits == 4 ? 17 * c : c; // 0xfff ， 一个f应该代表 255, 应该当[0-255]，按15等份划分，每一等份间隔 17。 所以获得的值须要乘以17, 才干表示rgb中255的值
	});
	const rgba='rgba('+ t.r + ','  +t.g + ','+ t.b +',' + opacity+')'
	return rgba;   // Color
}

// rgba转换16进制
export const hexify = (color) => {
	const values = color
		.replace(/rgba?\(/, '')
		.replace(/\)/, '')
		.replace(/[\s+]/g, '')
		.split(',');
	const a = parseFloat(values[3] || 1),
			r = Math.floor(a * parseInt(values[0]) + (1 - a) * 255),
			g = Math.floor(a * parseInt(values[1]) + (1 - a) * 255),
			b = Math.floor(a * parseInt(values[2]) + (1 - a) * 255);
	return "#" +
		("0" + r.toString(16)).slice(-2) +
		("0" + g.toString(16)).slice(-2) +
		("0" + b.toString(16)).slice(-2);
}




