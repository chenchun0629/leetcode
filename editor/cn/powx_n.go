package main

import "fmt"

func main() {
	fmt.Println(myPow(2, -2))
}

//实现 pow(x, n) ，即计算 x 的 n 次幂函数（即，x⁴
//
// Related Topics 递归 数学 👍 830 👎 0

//leetcode submit region begin(Prohibit modification and deletion)
func myPow(x float64, n int) float64 {
	if n >= 0 {
		return myPowQuickMul(x, n)
	}
	return 1.0 / myPowQuickMul(x, -n)
}

// 2 ^ 9
// 2 ^ 4 * 2 ^ 4 * 2
// 2 ^ 4
// 2 ^ 2 * 2 ^ 2
// 2 ^ 5
// 2 ^ 2 * 2 ^ 2 * 2
func myPowQuickMul(x float64, n int) float64 {
	if n == 0 {
		return 1
	}
	y := myPowQuickMul(x, n/2)
	if n%2 == 0 {
		return y * y
	}
	return y * y * x
}

//leetcode submit region end(Prohibit modification and deletion)
