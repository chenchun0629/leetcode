package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Printf("%d, \n", int(math.Log(9)/math.Log(3)))
}

//给定一个整数，写一个函数来判断它是否是 3 的幂次方。如果是，返回 true ；否则，返回 false 。
//
// 整数 n 是 3 的幂次方需满足：存在整数 x 使得 n == 3ˣ
//
//
//
// 示例 1：
//
//
//输入：n = 27
//输出：true
//
//
// 示例 2：
//
//
//输入：n = 0
//输出：false
//
//
// 示例 3：
//
//
//输入：n = 9
//输出：true
//
//
// 示例 4：
//
//
//输入：n = 45
//输出：false
//
//
//
//
// 提示：
//
//
// -2³¹ <= n <= 2³¹ - 1
//
//
//
//
// 进阶：你能不使用循环或者递归来完成本题吗？
// Related Topics 递归 数学 👍 236 👎 0

//leetcode submit region begin(Prohibit modification and deletion)
func isPowerOfThree(n int) bool {
	// math.Pow(3, 19) 表示 int 范围内3的最大幂等结果
	// math.Pow(3, 19) % n 表示 3的幂次 模 3的幂次 等于 0
	// 如 3^5 / 3^3 = 3^(5-2) = 3^2 余数等于0
	return (n > 0) && int(math.Pow(3, 19))%n == 0
}

//leetcode submit region end(Prohibit modification and deletion)

func isPowerOfThree_B(n int) bool {
	// ln e^2 / ln e = 2
	// ln 9 / ln 3 = 2
	// ln 27 / ln 3 = 3
	return (n > 0) && (int(math.Pow(3, math.Log(float64(n))/math.Log(3))) == n)
}

func isPowerOfThree_A(n int) bool {
	if n <= 0 {
		return false
	}
	for n != 1 {
		if n%3 != 0 {
			return false
		}

		n = n / 3
	}

	return true
}
