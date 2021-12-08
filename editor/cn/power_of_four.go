package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println(isPowerOfFour(1))
	fmt.Println(isPowerOfFour(2))
	fmt.Println(isPowerOfFour(3))
	fmt.Println(isPowerOfFour(4))
	fmt.Println(isPowerOfFour(5))
	fmt.Println(isPowerOfFour(6))
	fmt.Println(isPowerOfFour(7))
	fmt.Println(isPowerOfFour(8))
	fmt.Println(isPowerOfFour(12))
	fmt.Println(isPowerOfFour(16))
}

//给定一个整数，写一个函数来判断它是否是 4 的幂次方。如果是，返回 true ；否则，返回 false 。
//
// 整数 n 是 4 的幂次方需满足：存在整数 x 使得 n == 4ˣ
//
//
//
// 示例 1：
//
//
//输入：n = 16
//输出：true
//
//
// 示例 2：
//
//
//输入：n = 5
//输出：false
//
//
// 示例 3：
//
//
//输入：n = 1
//输出：true
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
// Related Topics 位运算 递归 数学 👍 276 👎 0

//leetcode submit region begin(Prohibit modification and deletion)
func isPowerOfFour(n int) bool {
	var m = math.Log(float64(n)) / math.Log(4)
	if m != math.Round(m) {
		return false
	}

	return n > 0 && int(math.Pow(4, m)) == n
}

//leetcode submit region end(Prohibit modification and deletion)
