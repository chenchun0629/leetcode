package main

import "fmt"

func main() {
	countBits(3)
	//fmt.Println()
}

//给你一个整数 n ，对于 0 <= i <= n 中的每个 i ，计算其二进制表示中 1 的个数 ，返回一个长度为 n + 1 的数组 ans 作为答案。
//
//
//
//
//
//
// 示例 1：
//
//
//输入：n = 2
//输出：[0,1,1]
//解释：
//0 --> 0
//1 --> 1
//2 --> 10
//
//
// 示例 2：
//
//
//输入：n = 5
//输出：[0,1,1,2,1,2]
//解释：
//0 --> 0
//1 --> 1
//2 --> 10
//3 --> 11
//4 --> 100
//5 --> 101
//
//
//
//
// 提示：
//
//
// 0 <= n <= 10⁵
//
//
//
//
// 进阶：
//
//
// 很容易就能实现时间复杂度为 O(n log n) 的解决方案，你可以在线性时间复杂度 O(n) 内用一趟扫描解决此问题吗？
// 你能不使用任何内置函数解决此问题吗？（如，C++ 中的 __builtin_popcount ）
//
//
//
// Related Topics 位运算 动态规划 👍 851 👎 0

//leetcode submit region begin(Prohibit modification and deletion)
func countBits(n int) []int {
	var bits = make([]int, n+1)
	var highBit = 0
	for i := 1; i <= n; i++ {
		if i&(i-1) == 0 {
			highBit = i
		}
		fmt.Printf("i: %d, %b, hb: %d, idx: %d, b: %d, bd: %d \n", i&(i-1), i, highBit, i-highBit, bits[bits[i-highBit]], bits[i-highBit]+1)
		bits[i] = bits[i-highBit] + 1
	}

	// 1, 10, 11, 100, 101, 110, 111, 1000, 1001, 1010, 1011, 1100

	return bits
}

//leetcode submit region end(Prohibit modification and deletion)

func countBits_A(n int) []int {
	var is = make([]int, 0)

	for i := 0; i <= n; i++ {
		is = append(is, countBitsHasOneCount(i))
	}

	return is
}

func countBitsHasOneCount(a int) int {
	var count = 0
	for a != 0 {
		a = a & (a - 1)
		count++
	}

	return count
}
