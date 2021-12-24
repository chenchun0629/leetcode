package main

import (
	"fmt"
)

func main() {
	//fmt.Println(searchRange([]int{5,7,7,8,8,10}, 6))
	fmt.Println(searchRange([]int{1, 3}, 1))
}

//给定一个按照升序排列的整数数组 nums，和一个目标值 target。找出给定目标值在数组中的开始位置和结束位置。
//
// 如果数组中不存在目标值 target，返回 [-1, -1]。
//
// 进阶：
//
//
// 你可以设计并实现时间复杂度为 O(log n) 的算法解决此问题吗？
//
//
//
//
// 示例 1：
//
//
//输入：nums = [5,7,7,8,8,10], target = 8
//输出：[3,4]
//
// 示例 2：
//
//
//输入：nums = [5,7,7,8,8,10], target = 6
//输出：[-1,-1]
//
// 示例 3：
//
//
//输入：nums = [], target = 0
//输出：[-1,-1]
//
//
//
// 提示：
//
//
// 0 <= nums.length <= 10⁵
// -10⁹ <= nums[i] <= 10⁹
// nums 是一个非递减数组
// -10⁹ <= target <= 10⁹
//
// Related Topics 数组 二分查找 👍 1351 👎 0

//leetcode submit region begin(Prohibit modification and deletion)
func searchRange(nums []int, target int) []int {
	var n = len(nums)
	var left, right = 0, n - 1
	var begin, end = -1, -1
	for left <= right {
		var mid = (left + right) / 2
		if nums[mid] == target {
			var tmp = mid
			for tmp > 0 && nums[tmp] == nums[tmp-1] {
				tmp = tmp - 1
			}
			begin = tmp

			tmp = mid
			for tmp < n-1 && nums[tmp] == nums[tmp+1] {
				tmp = tmp + 1
			}
			end = tmp
			break
		}

		if nums[mid] < target {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}

	return []int{begin, end}
}

//leetcode submit region end(Prohibit modification and deletion)
