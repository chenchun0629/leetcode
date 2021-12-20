package main

import (
	"fmt"
	"math"
	"sort"
)

func main() {
	fmt.Println(threeSumClosest([]int{0, 2, 1, -3}, 1))
}

//给你一个长度为 n 的整数数组 nums 和 一个目标值 target。请你从 nums 中选出三个整数，使它们的和与 target 最接近。
//
// 返回这三个数的和。
//
// 假定每组输入只存在恰好一个解。
//
//
//
// 示例 1：
//
//
//输入：nums = [-1,2,1,-4], target = 1
//输出：2
//解释：与 target 最接近的和是 2 (-1 + 2 + 1 = 2) 。
//
//
// 示例 2：
//
//
//输入：nums = [0,0,0], target = 1
//输出：0
//
//
//
//
// 提示：
//
//
// 3 <= nums.length <= 1000
// -1000 <= nums[i] <= 1000
// -10⁴ <= target <= 10⁴
//
// Related Topics 数组 双指针 排序 👍 980 👎 0

//leetcode submit region begin(Prohibit modification and deletion)
func threeSumClosest(nums []int, target int) int {
	var (
		closed int = math.MaxInt32
		n          = len(nums)
	)

	if n < 3 {
		return closed
	}

	sort.Ints(nums)
	var (
		first = 0
		abs   = func(x int) int {
			if x < 0 {
				return -1 * x
			}
			return x
		}
		update = func(x int) {
			a, b := abs(x-target), abs(closed-target)
			c := a < b
			_, _, _ = a, b, c
			if abs(x-target) < abs(closed-target) {
				closed = x
			}
		}
	)

	for ; first < n-2; first++ {
		var third = n - 1
		for second := first + 1; second < third; {
			var t = nums[first] + nums[second] + nums[third]
			if t == target {
				return target
			}

			update(t)

			if t > target {
				third--
			} else {
				second++
			}
		}
	}

	return closed
}

//leetcode submit region end(Prohibit modification and deletion)
