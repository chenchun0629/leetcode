package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println(minMoves([]int{1, 2, 3}))
}

//给你一个长度为 n 的整数数组，每次操作将会使 n - 1 个元素增加 1 。返回让数组所有元素相等的最小操作次数。
//
//
//
// 示例 1：
//
//
//输入：nums = [1,2,3]
//输出：3
//解释：
//只需要3次操作（注意每次操作会增加两个元素的值）：
//[1,2,3]  =>  [2,3,3]  =>  [3,4,3]  =>  [4,4,4]
//
//
// 示例 2：
//
//
//输入：nums = [1,1,1]
//输出：0
//
//
//
//
// 提示：
//
//
// n == nums.length
// 1 <= nums.length <= 10⁵
// -10⁹ <= nums[i] <= 10⁹
// 答案保证符合 32-bit 整数
//
// Related Topics 数组 数学 👍 405 👎 0

//leetcode submit region begin(Prohibit modification and deletion)
func minMoves(nums []int) int {
	// 0 0 1 => 1 1 1
	// 0 1 1 => 1 2 1 => 2 2 2
	// 0 1 3 => 1 2 3 => 2 3 3 => 3 4 3 => 4 4 4
	// 0 1 4 => 1 2 4 => 2 3 4 => 3 4 4 => 4 5 4 => 5 5 5
	// 0 2 3 => 1 3 3 => 2 4 3 => 3 4 4 => 4 5 4 => 5 5 5
	// 0 3 3 => 1 4 3 => 2 4 4 => 3 5 4 => 4 5 5 => 5 6 5 => 6 6 6
	var (
		n   = len(nums)
		min = math.MaxInt64
	)
	for i := 0; i < n; i++ {
		if min > nums[i] {
			min = nums[i]
		}
	}

	var t int
	for i := 0; i < n; i++ {
		t = t + nums[i] - min
	}

	return t
}

//leetcode submit region end(Prohibit modification and deletion)
