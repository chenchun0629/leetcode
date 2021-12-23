package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println(divide(10, 3))
	//fmt.Println(divide(6, 3))
	//fmt.Println(divide(7, 3))
	//fmt.Println(divide(8, 3))
	//fmt.Println(divide(9, 3))
	//fmt.Println(divide(10, 3))
}

//给定两个整数，被除数 dividend 和除数 divisor。将两数相除，要求不使用乘法、除法和 mod 运算符。
//
// 返回被除数 dividend 除以除数 divisor 得到的商。
//
// 整数除法的结果应当截去（truncate）其小数部分，例如：truncate(8.345) = 8 以及 truncate(-2.7335) = -2
//
//
//
// 示例 1:
//
// 输入: dividend = 10, divisor = 3
//输出: 3
//解释: 10/3 = truncate(3.33333..) = truncate(3) = 3
//
// 示例 2:
//
// 输入: dividend = 7, divisor = -3
//输出: -2
//解释: 7/-3 = truncate(-2.33333..) = -2
//
//
//
// 提示：
//
//
// 被除数和除数均为 32 位有符号整数。
// 除数不为 0。
// 假设我们的环境只能存储 32 位有符号整数，其数值范围是 [−2³¹, 231 − 1]。本题中，如果除法结果溢出，则返回 231 − 1。
//
// Related Topics 位运算 数学 👍 789 👎 0

//leetcode submit region begin(Prohibit modification and deletion)
func divide(dividend int, divisor int) int {
	if dividend == math.MinInt32 {
		if divisor == 1 {
			return math.MinInt32
		}
		if divisor == -1 {
			return math.MaxInt32
		}
	}
	if divisor == math.MinInt32 {
		if dividend == math.MinInt32 {
			return 1
		}
		return 0
	}
	if dividend == 0 {
		return 0
	}

	var sym bool
	if divisor < 0 {
		divisor = -divisor
		sym = !sym
	}

	if dividend < 0 {
		dividend = -dividend
		sym = !sym
	}
	var candidates = []int{divisor}
	//for y := divisor; y >= dividend-y; {
	for y := divisor; dividend-y >= y; {
		y += y
		candidates = append(candidates, y)
	}

	var r = 0
	for i := len(candidates) - 1; i >= 0; i-- {
		if dividend >= candidates[i] {
			r |= 1 << i
			dividend -= candidates[i]
		}
	}

	if sym {
		r = -r
	}
	return r
}

//leetcode submit region end(Prohibit modification and deletion)

// 超时
func divide_A(dividend int, divisor int) int {
	if dividend == math.MinInt32 {
		if divisor == 1 {
			return math.MinInt32
		}
		if divisor == -1 {
			return math.MaxInt32
		}
	}
	if divisor == math.MinInt32 {
		if dividend == math.MinInt32 {
			return 1
		}
		return 0
	}
	if dividend == 0 {
		return 0
	}

	var sym bool
	if divisor < 0 {
		divisor = -divisor
		sym = !sym
	}

	if dividend < 0 {
		dividend = -dividend
		sym = !sym
	}
	var r = 0
	for ; dividend >= divisor; dividend = dividend - divisor {
		r++
	}

	if sym {
		r = -r
	}
	return r
}
