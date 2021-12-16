package main

import "fmt"

func main() {
	fmt.Println(searchInsert_B1([]int{1, 3, 5, 6}, 9))
	//fmt.Println(searchInsert_B1([]int{1, 2}, 0))
}

//给定一个排序数组和一个目标值，在数组中找到目标值，并返回其索引。如果目标值不存在于数组中，返回它将会被按顺序插入的位置。
//
// 请必须使用时间复杂度为 O(log n) 的算法。
//
//
//
// 示例 1:
//
//
//输入: nums = [1,3,5,6], target = 5
//输出: 2
//
//
// 示例 2:
//
//
//输入: nums = [1,3,5,6], target = 2
//输出: 1
//
//
// 示例 3:
//
//
//输入: nums = [1,3,5,6], target = 7
//输出: 4
//
//
// 示例 4:
//
//
//输入: nums = [1,3,5,6], target = 0
//输出: 0
//
//
// 示例 5:
//
//
//输入: nums = [1], target = 0
//输出: 0
//
//
//
//
// 提示:
//
//
// 1 <= nums.length <= 10⁴
// -10⁴ <= nums[i] <= 10⁴
// nums 为无重复元素的升序排列数组
// -10⁴ <= target <= 10⁴
//
// Related Topics 数组 二分查找 👍 1223 👎 0

//leetcode submit region begin(Prohibit modification and deletion)
func searchInsert(nums []int, target int) int {
	var (
		ret         = len(nums)
		left, right = 0, ret - 1
	)

	for left <= right {
		var mid = (left + right) / 2
		if nums[mid] > target {
			if mid-1 >= 0 && nums[mid-1] < target {
				return mid
			}
			right = mid - 1
		} else if nums[mid] < target {
			if mid+1 < ret && nums[mid+1] > target {
				return mid + 1
			}
			left = mid + 1
		} else {
			return mid
		}
	}

	return left
}

//leetcode submit region end(Prohibit modification and deletion)

func searchInsert_B2(nums []int, target int) int {
	var (
		ret         = len(nums)
		left, right = 0, ret - 1
	)

	for left <= right {
		var mid = (left + right) / 2
		if nums[mid] >= target {
			ret = mid
			right = mid - 1
		} else {
			left = mid + 1
		}
	}

	return ret
}

func searchInsert_A1(nums []int, target int) int {
	for k, v := range nums {
		if v >= target {
			return k
		}
	}

	return len(nums)
}
