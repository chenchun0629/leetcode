package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println(intersection([]int{2, 6, 2, 9, 1}, []int{7, 1}))
}

//给定两个数组，编写一个函数来计算它们的交集。
//
//
//
// 示例 1：
//
// 输入：nums1 = [1,2,2,1], nums2 = [2,2]
//输出：[2]
//
//
// 示例 2：
//
// 输入：nums1 = [4,9,5], nums2 = [9,4,9,8,4]
//输出：[9,4]
//
//
//
// 说明：
//
//
// 输出结果中的每个元素一定是唯一的。
// 我们可以不考虑输出结果的顺序。
//
// Related Topics 数组 哈希表 双指针 二分查找 排序 👍 449 👎 0

//leetcode submit region begin(Prohibit modification and deletion)
func intersection(nums1 []int, nums2 []int) []int {
	sort.Ints(nums1)
	sort.Ints(nums2)

	if len(nums2) > len(nums1) {
		nums1, nums2 = nums2, nums1
	}

	var r []int
	var l1, l2 = len(nums1), len(nums2)
	var i1, i2 = 0, 0

	for i1 < l1 && i2 < l2 {
		if i1 > 0 && nums1[i1] == nums1[i1-1] {
			i1++
			continue
		}

		if i2 > 0 && nums2[i2] == nums2[i2-1] {
			i2++
			continue
		}

		if nums1[i1] == nums2[i2] {
			r = append(r, nums1[i1])
			i1++
			i2++
		} else if nums1[i1] > nums2[i2] {
			i2++
		} else if nums1[i1] < nums2[i2] {
			i1++
		}

	}

	return r
}

//leetcode submit region end(Prohibit modification and deletion)
