package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Printf("%#v \n", fourSum([]int{1, 0, -1, 0, -2, 2}, 0))
	fmt.Printf("%#v \n", fourSum([]int{2, 2, 2, 2}, 8))
	fmt.Printf("%#v \n", fourSum([]int{2, 2, 2, 2, 2}, 8))
}

//给你一个由 n 个整数组成的数组 nums ，和一个目标值 target 。请你找出并返回满足下述全部条件且不重复的四元组 [nums[a], nums[
//b], nums[c], nums[d]] （若两个四元组元素一一对应，则认为两个四元组重复）：
//
//
// 0 <= a, b, c, d < n
// a、b、c 和 d 互不相同
// nums[a] + nums[b] + nums[c] + nums[d] == target
//
//
// 你可以按 任意顺序 返回答案 。
//
//
//
// 示例 1：
//
//
//输入：nums = [1,0,-1,0,-2,2], target = 0
//输出：[[-2,-1,1,2],[-2,0,0,2],[-1,0,0,1]]
//
//
// 示例 2：
//
//
//输入：nums = [2,2,2,2,2], target = 8
//输出：[[2,2,2,2]]
//
//
//
//
// 提示：
//
//
// 1 <= nums.length <= 200
// -10⁹ <= nums[i] <= 10⁹
// -10⁹ <= target <= 10⁹
//
// Related Topics 数组 双指针 排序 👍 1044 👎 0

//leetcode submit region begin(Prohibit modification and deletion)
func fourSum(nums []int, target int) [][]int {
	var (
		n = len(nums)
		r [][]int
	)

	if n < 4 {
		return r
	}

	sort.Ints(nums)

	for first := 0; first < n-3; first++ {
		if first > 0 && nums[first] == nums[first-1] {
			continue
		}

		for second := first + 1; second < n-2; second++ {
			if second > first+1 && nums[second] == nums[second-1] {
				continue
			}
			var (
				t      = target + (-1 * (nums[first] + nums[second]))
				third  = second + 1
				fourth = n - 1
			)

			for ; third < n-1; third++ {
				if third > second+1 && nums[third] == nums[third-1] {
					continue
				}

				for third < fourth && t < nums[third]+nums[fourth] {
					fourth--
				}

				if third == fourth {
					break
				}

				if t == nums[third]+nums[fourth] {
					r = append(r, []int{nums[first], nums[second], nums[third], nums[fourth]})
				}
			}
		}

	}
	return r
}

//leetcode submit region end(Prohibit modification and deletion)
