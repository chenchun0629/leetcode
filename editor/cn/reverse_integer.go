package main

import (
	"fmt"
	"math"
	"strconv"
)

//给出一个 32 位的有符号整数，你需要将这个整数中每位上的数字进行反转。
//
// 示例 1:
//
// 输入: 123
//输出: 321
//
//
// 示例 2:
//
// 输入: -123
//输出: -321
//
//
// 示例 3:
//
// 输入: 120
//输出: 21
//
//
// 注意:
//
// 假设我们的环境只能存储得下 32 位的有符号整数，则其数值范围为 [−231, 231 − 1]。请根据这个假设，如果反转后整数溢出那么就返回 0。
// Related Topics 数学
// 👍 2346 👎 0

//leetcode submit region begin(Prohibit modification and deletion)
func reverse(x int) int {
	return reverseB(x)
}

// 方案B
// 上面一个方案有些取巧，其实用的字符串反转，而非真正意义上的数字反值。
// 通过 p = x%10 取出 最后一位
// 通过 x = x/10 得到 取出最后一位后的值
// 通过 y = y*10 + 8 得到 反转的值
// 然后在for循环中，对每一步进行校验
func reverseB(x int) int {
	var (
		ul = int(math.Pow(2, 31) - 1)  // 2147483647
		ll = int(math.Pow(2, 31)) * -1 // -2147483648
		y  = 0
	)

	for x != 0 {
		if x > ul || x < ll {
			return 0
		}
		var p = x % 10
		x = x / 10

		// 提前判断
		if y > ul/10 || y < ll/10 {
			return 0
		}

		// 极限情况： y = 214748364 || y = -214748364
		// 所以 p 需要小于 7 以及 大于 -8
		if y == ul/10 && p > 7 || y < ll/10 && p < -8 {
			return 0
		}

		y = y*10 + p
	}

	return y
}

// 方案A
// 转换成字符串
// 需要注意的是不仅需要考虑输入值在范围中，输出值是不是也在范围中
func reverseA(x int) int {
	var (
		symbol int = 1
		s      string
		ul     = int32(math.Pow(2, 31) - 1)
		ll     = int32(math.Pow(2, 31)) * -1
	)

	if x > int(ul) || x < int(ll) {
		return 0
	}

	if x < 0 {
		symbol = -1
	} else {
		symbol = 1
	}

	s = fmt.Sprint(x * symbol)

	var (
		r = []rune(s)
		l = len(r)
		i int
	)

	for i = 0; i < l/2; i++ {
		r[i], r[l-i-1] = r[l-i-1], r[i]
	}

	x, _ = strconv.Atoi(string(r))

	if x > int(ul) {
		return 0
	}

	x = x * symbol

	if x < int(ll) {
		return 0
	}

	return x
}

//leetcode submit region end(Prohibit modification and deletion)

func main() {
	//2147483647
	//fmt.Println(reverseB(6463847412))
	//fmt.Println(reverseB(-123))
}
