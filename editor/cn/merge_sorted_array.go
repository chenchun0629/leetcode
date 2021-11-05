package main

import (
	"fmt"
)

func main() {
	var (
		num1 = []int{1, 2, 3, 0, 0, 0}
		//num1 = []int{0}
		m = 3
		//num2 = []int{1}
		num2 = []int{2, 5, 6}
		n    = 3
	)
	merge(num1, m, num2, n)
	fmt.Println(num1)
}

//给你两个按 非递减顺序 排列的整数数组 nums1 和 nums2，另有两个整数 m 和 n ，分别表示 nums1 和 nums2 中的元素数目。
//
// 请你 合并 nums2 到 nums1 中，使合并后的数组同样按 非递减顺序 排列。
//
// 注意：最终，合并后数组不应由函数返回，而是存储在数组 nums1 中。为了应对这种情况，nums1 的初始长度为 m + n，其中前 m 个元素表示应合并
//的元素，后 n 个元素为 0 ，应忽略。nums2 的长度为 n 。
//
//
//
// 示例 1：
//
//
//输入：nums1 = [1,2,3,0,0,0], m = 3, nums2 = [2,5,6], n = 3
//输出：[1,2,2,3,5,6]
//解释：需要合并 [1,2,3] 和 [2,5,6] 。
//合并结果是 [1,2,2,3,5,6] ，其中斜体加粗标注的为 nums1 中的元素。
//
//
// 示例 2：
//
//
//输入：nums1 = [1], m = 1, nums2 = [], n = 0
//输出：[1]
//解释：需要合并 [1] 和 [] 。
//合并结果是 [1] 。
//
//
// 示例 3：
//
//
//输入：nums1 = [0], m = 0, nums2 = [1], n = 1
//输出：[1]
//解释：需要合并的数组是 [] 和 [1] 。
//合并结果是 [1] 。
//注意，因为 m = 0 ，所以 nums1 中没有元素。nums1 中仅存的 0 仅仅是为了确保合并结果可以顺利存放到 nums1 中。
//
//
//
//
// 提示：
//
//
// nums1.length == m + n
// nums2.length == n
// 0 <= m, n <= 200
// 1 <= m + n <= 200
// -10⁹ <= nums1[i], nums2[j] <= 10⁹
//
//
//
//
// 进阶：你可以设计实现一个时间复杂度为 O(m + n) 的算法解决此问题吗？
// Related Topics 数组 双指针 排序 👍 1167 👎 0

//leetcode submit region begin(Prohibit modification and deletion)
func merge(nums1 []int, m int, nums2 []int, n int) {
	merge_A(nums1, m, nums2, n)
}

func merge_A(nums1 []int, m int, nums2 []int, n int) {
	var seek1, seek2 int
	var insert = func(num []int, i int, v int) {
		//fmt.Println(nums1, i, v)
		copy(num[i+1:], num[i:])
		num[i] = v
	}
	for seek1 < m || seek2 < n {
		if seek1 < m && seek2 < n {
			if nums1[seek1] >= nums2[seek2] {
				insert(nums1, seek1, nums2[seek2])
				seek2++
				m++
			} else {
				seek1++
			}
		} else if seek2 < n {
			insert(nums1, seek1, nums2[seek2])
			seek2++
			m++
		} else {
			seek1++
		}
		//fmt.Println(seek1, m, seek2, n)
		//time.Sleep(time.Second)
	}
}

//leetcode submit region end(Prohibit modification and deletion)
