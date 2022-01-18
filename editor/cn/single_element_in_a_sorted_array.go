package main

import "fmt"

func main() {
	fmt.Println(singleNonDuplicate([]int{0, 1, 1, 2, 2, 3, 3, 4, 4}))
	fmt.Println(singleNonDuplicate([]int{1, 1, 2, 3, 3, 4, 4, 8, 8}))
	fmt.Println(singleNonDuplicate([]int{1, 1, 2, 2, 3, 3, 4, 4, 8}))
	fmt.Println(singleNonDuplicate([]int{0, 1, 1, 2, 2, 3, 3, 4, 4, 5, 5}))
	fmt.Println(singleNonDuplicate([]int{0, 0, 1, 1, 2, 2, 3, 3, 4, 4, 5}))
}

//给你一个仅由整数组成的有序数组，其中每个元素都会出现两次，唯有一个数只会出现一次。
//
// 请你找出并返回只出现一次的那个数。
//
// 你设计的解决方案必须满足 O(log n) 时间复杂度和 O(1) 空间复杂度。
//
//
//
// 示例 1:
//
//
//输入: nums = [1,1,2,3,3,4,4,8,8]
//输出: 2
//
//
// 示例 2:
//
//
//输入: nums =  [3,3,7,7,10,11,11]
//输出: 10
//
//
//
//
//
//
// 提示:
//
//
// 1 <= nums.length <= 10⁵
// 0 <= nums[i] <= 10⁵
//
// Related Topics 数组 二分查找 👍 305 👎 0

//leetcode submit region begin(Prohibit modification and deletion)
func singleNonDuplicate(nums []int) int {
	//           ↓
	// 0 1 1 2 2 3 3 4 4 5 5
	// 0 0 1 1 2 2 3 3 4 4 5

	//         ↓
	// 0 1 1 2 2 3 3 4 4
	// 0 0 1 1 2 2 3 3 4
	//
	// 0 0 1 2 2 3 3 4 4
	// 0 0 1 1 2 2 3 4 4
	var left, right = 0, len(nums) - 1
	for left < right {
		var mid = left + (right-left)/2
		var flag = (right-left+2)/2%2 == 0

		if flag {
			if nums[mid] == nums[mid+1] {
				right = mid - 1
			} else if nums[mid] == nums[mid-1] {
				left = mid + 1
			} else {
				return nums[mid]
			}
		} else {
			if nums[mid] == nums[mid+1] {
				left = mid
			} else if nums[mid] == nums[mid-1] {
				right = mid + 1
			} else {
				return nums[mid]
			}
		}
	}

	return nums[left]
}

//leetcode submit region end(Prohibit modification and deletion)

func singleNonDuplicate_A(nums []int) int {
	lth := len(nums)
	if lth == 1 {
		return nums[0]
	}
	left, right := 0, lth-1
	for left <= right {
		mid := (left + right) >> 1

		// 如果一个元素是目标元素，则其值与左右两边的值不相等
		if (mid-1 < 0 || nums[mid] != nums[mid-1]) && (mid+1 > lth-1 || nums[mid] != nums[mid+1]) {
			return nums[mid]
		}
		// 如果包含mid，左侧刚好偶数个，那么mid和mid-1的值要相等；如果包含mid为奇数个，则mid和mid+1要相等。
		if (mid%2 == 1 && nums[mid] == nums[mid-1]) || (mid%2 == 0 && nums[mid] == nums[mid+1]) {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}
	return nums[left]
}
