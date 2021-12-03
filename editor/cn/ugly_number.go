package main

import "fmt"

func main() {
	fmt.Printf("%b \n", 2)
	fmt.Printf("%b \n", 3)
	fmt.Printf("%b \n", 5)
	fmt.Printf("%b \n", 7)
	fmt.Printf("%b \n", 11)
	fmt.Printf("%b \n", 13)
	fmt.Printf("%b \n", 17)
	fmt.Printf("\n\n\n\n\n")

	fmt.Printf("%b, %b \n", 1, 1|1)
	fmt.Printf("%b, %b - \n", 2, 2|1)
	fmt.Printf("%b, %b - \n", 3, 3|1)
	fmt.Printf("%b, %b \n", 4, 4|1)
	fmt.Printf("%b, %b - \n", 5, 5|1)
	fmt.Printf("%b, %b \n", 6, 6|1)
	fmt.Printf("%b, %b - \n", 7, 7|1)
	fmt.Printf("%b, %b \n", 8, 8|1)
	fmt.Printf("%b, %b \n", 9, 9|1)
	fmt.Printf("%b, %b \n", 10, 0|1)
	fmt.Printf("%b, %b - \n", 11, 1|1)
	fmt.Printf("%b, %b \n", 12, 2|1)
	fmt.Printf("%b, %b - \n", 13, 3|1)
	fmt.Printf("%b, %b \n", 14, 4|1)
	fmt.Printf("%b, %b \n", 15, 5|1)
	fmt.Printf("%b, %b \n", 16, 6|1)
	fmt.Printf("%b, %b - \n", 17, 7|1)
}

//给你一个整数 n ，请你判断 n 是否为 丑数 。如果是，返回 true ；否则，返回 false 。
//
// 丑数 就是只包含质因数 2、3 和/或 5 的正整数。
//
//
//
// 示例 1：
//
//
//输入：n = 6
//输出：true
//解释：6 = 2 × 3
//
// 示例 2：
//
//
//输入：n = 8
//输出：true
//解释：8 = 2 × 2 × 2
//
//
// 示例 3：
//
//
//输入：n = 14
//输出：false
//解释：14 不是丑数，因为它包含了另外一个质因数 7 。
//
//
// 示例 4：
//
//
//输入：n = 1
//输出：true
//解释：1 通常被视为丑数。
//
//
//
//
// 提示：
//
//
// -2³¹ <= n <= 2³¹ - 1
//
// Related Topics 数学 👍 288 👎 0

//leetcode submit region begin(Prohibit modification and deletion)

func isUgly(n int) bool {
	var factors = []int{2, 3, 5}

	if n <= 0 {
		return false
	}
	for _, f := range factors {
		for n%f == 0 {
			n /= f
		}
	}

	return n == 1
}

//leetcode submit region end(Prohibit modification and deletion)
