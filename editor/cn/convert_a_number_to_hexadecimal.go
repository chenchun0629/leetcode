package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(8&0xf, 8%16)
	fmt.Println(13&0xf, 13%16)
	fmt.Println(18&0xf, 18%16)
	fmt.Println(23&0xf, 23%16)
	fmt.Println(35&0xf, 35%16)
}

//给定一个整数，编写一个算法将这个数转换为十六进制数。对于负整数，我们通常使用 补码运算 方法。
//
// 注意:
//
//
// 十六进制中所有字母(a-f)都必须是小写。
// 十六进制字符串中不能包含多余的前导零。如果要转化的数为0，那么以单个字符'0'来表示；对于其他情况，十六进制字符串中的第一个字符将不会是0字符。
// 给定的数确保在32位有符号整数范围内。
// 不能使用任何由库提供的将数字直接转换或格式化为十六进制的方法。
//
//
// 示例 1：
//
//
//输入:
//26
//
//输出:
//"1a"
//
//
// 示例 2：
//
//
//输入:
//-1
//
//输出:
//"ffffffff"
//
// Related Topics 位运算 数学 👍 221 👎 0

//leetcode submit region begin(Prohibit modification and deletion)
func toHex(num int) string {
	if num == 0 {
		return "0"
	}

	var sb = &strings.Builder{}
	for i := 7; i >= 0; i-- {
		var val = (num >> (4 * i)) & 0xf // 这部最关键
		// num & 0xf 相当于 num % 16
		// num >> 4 相当于 num / 16
		if val > 0 || sb.Len() > 0 {
			var digit byte
			if val < 10 {
				digit = '0' + byte(val)
			} else {
				digit = 'a' + byte(val-10)
			}
			sb.WriteByte(digit)
		}
	}

	return sb.String()
}

//leetcode submit region end(Prohibit modification and deletion)
