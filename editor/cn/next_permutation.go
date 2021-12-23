package main

import (
	"fmt"
	"sort"
)

func main() {
	var x = []int{2, 2, 7, 5, 4, 3, 2, 2, 1}
	//var x = []int{4, 2, 0, 2, 3, 2, 0}
	//var x = []int{5,4,7,5,3,2}
	//var x = []int{3, 2, 1}
	nextPermutation(x)
	fmt.Println(x)
}

//实现获取 下一个排列 的函数，算法需要将给定数字序列重新排列成字典序中下一个更大的排列（即，组合出下一个更大的整数）。
//
// 如果不存在下一个更大的排列，则将数字重新排列成最小的排列（即升序排列）。
//
// 必须 原地 修改，只允许使用额外常数空间。
//
//
//
// 示例 1：
//
//
//输入：nums = [1,2,3]
//输出：[1,3,2]
//
//
// 示例 2：
//
//
//输入：nums = [3,2,1]
//输出：[1,2,3]
//
//
// 示例 3：
//
//
//输入：nums = [1,1,5]
//输出：[1,5,1]
//
//
// 示例 4：
//
//
//输入：nums = [1]
//输出：[1]
//
//
//
//
// 提示：
//
//
// 1 <= nums.length <= 100
// 0 <= nums[i] <= 100
//
// Related Topics 数组 双指针 👍 1454 👎 0

//leetcode submit region begin(Prohibit modification and deletion)
func nextPermutation(nums []int) {
	var n = len(nums)
	if n <= 1 {
		return
	}

	var i = n - 2
	for i >= 0 && nums[i] >= nums[i+1] {
		i--
	}

	var k = -1
	if i >= 0 {
		for j := i + 1; j < n; j++ {
			if nums[j] > nums[i] && (k == -1 || nums[j] < nums[k]) {
				k = j
			}
		}
	}

	//fmt.Println(nums, i, k)
	//return
	if i >= 0 && k > 0 {
		nums[i], nums[k] = nums[k], nums[i]
		sort.Ints(nums[i+1:])
	} else {
		sort.Ints(nums)
	}
}

//leetcode submit region end(Prohibit modification and deletion)
