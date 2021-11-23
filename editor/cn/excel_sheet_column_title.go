package main

import "fmt"

func main() {
	//fmt.Println(2147483647/26/26/26/26/26)
	fmt.Println(string(rune(0 + 65 - 1)))
	fmt.Println(701 / 26 / 26)
}

//给你一个整数 columnNumber ，返回它在 Excel 表中相对应的列名称。
//
// 例如：
//
//
//A -> 1
//B -> 2
//C -> 3
//...
//Z -> 26
//AA -> 27
//AB -> 28
//...
//
//
//
//
// 示例 1：
//
//
//输入：columnNumber = 1
//输出："A"
//
//
// 示例 2：
//
//
//输入：columnNumber = 28
//输出："AB"
//
//
// 示例 3：
//
//
//输入：columnNumber = 701
//输出："ZY"
//
//
// 示例 4：
//
//
//输入：columnNumber = 2147483647
//输出："FXSHRXW"
//
//
//
//
// 提示：
//
//
// 1 <= columnNumber <= 2³¹ - 1
//
// Related Topics 数学 字符串 👍 470 👎 0

//leetcode submit region begin(Prohibit modification and deletion)
func convertToTitle(columnNumber int) string {
	if columnNumber == 0 {
		return ""
	}

	var (
		mod int
	)

	mod = columnNumber % 26
	columnNumber = columnNumber / 26
	if mod == 0 {
		mod = 26
		columnNumber -= 1
	}

	return convertToTitle(columnNumber) + string(rune(mod+65-1))
}

//leetcode submit region end(Prohibit modification and deletion)

func convertToTitle_A(columnNumber int) string {
	var (
		s   = make([]rune, 0)
		mod int
	)
	for columnNumber != 0 {
		mod = columnNumber % 26
		columnNumber = columnNumber / 26
		if mod == 0 {
			mod = 26
			columnNumber -= 1
		}

		s = append([]rune{rune(mod + 65 - 1)}, s...)

	}
	return string(s)
}
