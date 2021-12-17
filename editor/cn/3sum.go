package main

import (
	"fmt"
	"sort"
)

func main() {
	//fmt.Printf("%#v \n", threeSum([]int{-1, 0, 1, 2, -1, -4}))
	//fmt.Printf("%#v \n", threeSum([]int{0, 0, 0, 0}))
	//fmt.Printf("%#v \n", threeSum([]int{-2, 0, 1, 1, 2}))
	fmt.Printf("%#v \n", threeSum([]int{3, 0, -2, -1, 1, 2}))
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

func threeSumx(nums []int) [][]int {
	n := len(nums)
	sort.Ints(nums)
	ans := make([][]int, 0)

	// 枚举 a
	for first := 0; first < n; first++ {
		// 需要和上一次枚举的数不相同
		if first > 0 && nums[first] == nums[first-1] {
			continue
		}
		// c 对应的指针初始指向数组的最右端
		third := n - 1
		target := -1 * nums[first]
		// 枚举 b
		for second := first + 1; second < n; second++ {
			// 需要和上一次枚举的数不相同
			if second > first+1 && nums[second] == nums[second-1] {
				continue
			}
			// 需要保证 b 的指针在 c 的指针的左侧
			for second < third && nums[second]+nums[third] > target {
				third--
			}
			// 如果指针重合，随着 b 后续的增加
			// 就不会有满足 a+b+c=0 并且 b<c 的 c 了，可以退出循环
			if second == third {
				break
			}
			if nums[second]+nums[third] == target {
				ans = append(ans, []int{nums[first], nums[second], nums[third]})
			}
		}
	}
	return ans
}

func threeSum(nums []int) [][]int {
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

//leetcode submit region end(Prohibit modification and deletion)
