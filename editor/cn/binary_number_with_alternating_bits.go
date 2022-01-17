package main

import "fmt"

func main() {
	// 1010 >> 1 = 101 ^ 1010 =>1111
	// 1111 & 10000 = 0

	// 1110 >> 1 = 111 ^ 1110 => 1001
	// 1001 & 1010 = 1000

	fmt.Printf("%b, %b, %t \n", 1, (1>>1)^1, hasAlternatingBits(1))
	fmt.Printf("%b, %b, %t \n", 2, (2>>1)^2, hasAlternatingBits(2))
	fmt.Printf("%b, %b, %t \n", 3, (3>>1)^3, hasAlternatingBits(3))
	fmt.Printf("%b, %b, %t \n", 4, (4>>1)^4, hasAlternatingBits(4))
	fmt.Printf("%b, %b, %t \n", 5, (5>>1)^5, hasAlternatingBits(5))
	fmt.Printf("%b, %b, %t \n", 6, (6>>1)^6, hasAlternatingBits(6))
	fmt.Printf("%b, %b, %t \n", 7, (7>>1)^7, hasAlternatingBits(7))
	fmt.Printf("%b, %b, %t \n", 8, (8>>1)^8, hasAlternatingBits(8))
	fmt.Printf("%b, %b, %t \n", 9, (9>>1)^9, hasAlternatingBits(9))
	fmt.Printf("%b, %b, %t \n", 10, ((10>>1)^10)&((10>>1)^10+1), hasAlternatingBits(10))
}

//给定一个正整数，检查它的二进制表示是否总是 0、1 交替出现：换句话说，就是二进制表示中相邻两位的数字永不相同。
//
//
//
// 示例 1：
//
//
//输入：n = 5
//输出：true
//解释：5 的二进制表示是：101
//
//
// 示例 2：
//
//
//输入：n = 7
//输出：false
//解释：7 的二进制表示是：111.
//
// 示例 3：
//
//
//输入：n = 11
//输出：false
//解释：11 的二进制表示是：1011.
//
// 示例 4：
//
//
//输入：n = 10
//输出：true
//解释：10 的二进制表示是：1010.
//
// 示例 5：
//
//
//输入：n = 3
//输出：false
//
//
//
//
// 提示：
//
//
// 1 <= n <= 2³¹ - 1
//
// Related Topics 位运算 👍 113 👎 0

//leetcode submit region begin(Prohibit modification and deletion)
func hasAlternatingBits(n int) bool {
	var m = (n >> 1) ^ n
	return m&(m+1) == 0
}

//leetcode submit region end(Prohibit modification and deletion)

func hasAlternatingBits_A(n int) bool {
	var first = true
	var next int

	for n >= 1 {
		var tmp = n & 1
		if first {
			first = false
		} else {
			if tmp != next {
				return false
			}
		}

		next = 1 - tmp
		n = n >> 1
	}
	return true
}
