package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Printf("%#v \n", threeSum([]int{-1, 0, 1, 2, -1, -4}))
	//fmt.Printf("%#v \n", threeSum([]int{0, 0, 0, 0}))
	//fmt.Printf("%#v \n", threeSum([]int{-2, 0, 1, 1, 2}))
	//fmt.Printf("%#v \n", threeSum([]int{3, 0, -2, -1, 1, 2}))
}

//给你一个包含 n 个整数的数组 nums，判断 nums 中是否存在三个元素 a，b，c ，使得 a + b + c = 0 ？请你找出所有和为 0 且不重
//复的三元组。
//
// 注意：答案中不可以包含重复的三元组。
//
//
//
// 示例 1：
//
//
//输入：nums = [-1,0,1,2,-1,-4]
//输出：[[-1,-1,2],[-1,0,1]]
//
//
// 示例 2：
//
//
//输入：nums = []
//输出：[]
//
//
// 示例 3：
//
//
//输入：nums = [0]
//输出：[]
//
//
//
//
// 提示：
//
//
// 0 <= nums.length <= 3000
// -10⁵ <= nums[i] <= 10⁵
//
// Related Topics 数组 双指针 排序 👍 4095 👎 0

//leetcode submit region begin(Prohibit modification and deletion)

func threeSum(nums []int) [][]int {
	var (
		n = len(nums)
		r = make([][]int, 0)
	)

	if n < 3 {
		return r
	}

	sort.Ints(nums)

	var left = 0

	for ; left < n; left++ {
		if left > 0 && nums[left] == nums[left-1] {
			continue
		}

		var right = n - 1
		var target = nums[left] * -1

		for mid := left + 1; mid < n; mid++ {
			if mid > left+1 && nums[mid] == nums[mid-1] {
				continue
			}

			for mid < right && nums[right]+nums[mid] > target {
				right--
			}

			if mid == right {
				break
			}

			if nums[right]+nums[mid] == target {
				r = append(r, []int{nums[left], nums[mid], nums[right]})
			}
		}
	}

	return r
}

//leetcode submit region end(Prohibit modification and deletion)

func threeSum_A(nums []int) [][]int {
	var (
		n = len(nums)
		r = make([][]int, 0)
	)

	if n < 3 {
		return r
	}

	sort.Ints(nums)

	for first := 0; first < n; first++ {
		// 过滤相同数
		if first > 0 && nums[first] == nums[first-1] {
			continue
		}

		var (
			third  = n - 1
			target = -1 * nums[first]
		)

		for second := first + 1; second < n; second++ {
			// 过滤相同数
			if second > first+1 && nums[second] == nums[second-1] {
				continue
			}

			for second < third && nums[second]+nums[third] > target {
				third--
			}

			if second == third {
				break
			}

			if nums[second]+nums[third] == target {
				r = append(r, []int{nums[first], nums[second], nums[third]})
			}
		}
	}

	return r
}
