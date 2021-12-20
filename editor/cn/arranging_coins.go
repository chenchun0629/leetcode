package main

import "fmt"

func main() {
	fmt.Println(arrangeCoins(1))
}

//你总共有 n 枚硬币，并计划将它们按阶梯状排列。对于一个由 k 行组成的阶梯，其第 i 行必须正好有 i 枚硬币。阶梯的最后一行 可能 是不完整的。
//
// 给你一个数字 n ，计算并返回可形成 完整阶梯行 的总行数。
//
//
//
// 示例 1：
//
//
//输入：n = 5
//输出：2
//解释：因为第三行不完整，所以返回 2 。
//
//
// 示例 2：
//
//
//输入：n = 8
//输出：3
//解释：因为第四行不完整，所以返回 3 。
//
//
//
//
// 提示：
//
//
// 1 <= n <= 2³¹ - 1
//
// Related Topics 数学 二分查找 👍 183 👎 0

//leetcode submit region begin(Prohibit modification and deletion)
func arrangeCoins(n int) int {
	// (1 + n)n / 2
	// 1 = (1+1)1/2 = 1
	// 1 + 2 = (1+2)2/2 = 3
	// 1 + 2 + 3 = (1+3)3/2 =  6
	// 1 + ... + 10 = (1 + 10)10 / 2 = 55
	// 1 + ... + 100 = (1+100)100 / 2 = 5050

	var (
		left, right = 1, n
		mid         = n
	)

	for left <= right {
		var t = (1 + mid) * mid / 2
		if t == n {
			return mid
		}
		if t > n {
			if mid > 1 && n > (1+mid-1)*(mid-1)/2 {
				return mid - 1
			}
			right = mid - 1
		}
		if t < n {
			if n < (1+mid+1)*(mid+1)/2 {
				return mid
			}
			left = mid + 1
		}

		mid = (left + right) / 2
	}

	return 0
}

//leetcode submit region end(Prohibit modification and deletion)

func arrangeCoins_A(n int) int {
	// (1 + n)n / 2
	// 1 = (1+1)1/2 = 1
	// 1 + 2 = (1+2)2/2 = 3
	// 1 + 2 + 3 = (1+3)3/2 =  6
	// 1 + ... + 10 = (1 + 10)10 / 2 = 55
	// 1 + ... + 100 = (1+100)100 / 2 = 5050

	var (
		left, right = 1, n
		mid         = n
	)

	for left <= right {
		var t = (1 + mid) * mid / 2
		if t == n {
			return mid
		}
		if t > n {
			if mid > 1 && n > (1+mid-1)*(mid-1)/2 {
				return mid - 1
			}
			right = mid - 1
		}
		if t < n {
			if n < (1+mid+1)*(mid+1)/2 {
				return mid
			}
			left = mid + 1
		}

		mid = (left + right) / 2
	}

	return 0
}
