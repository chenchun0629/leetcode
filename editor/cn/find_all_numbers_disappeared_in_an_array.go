package main

import (
	"fmt"
	"sort"
)

func main() {
	// 1 3
	//fmt.Println(0 % 2)
	//fmt.Println(2 % 2)

	fmt.Println(findDisappearedNumbers([]int{1, 2, 2, 3, 3, 4, 7, 8}))
	//fmt.Println(findDisappearedNumbers([]int{4, 3, 2, 7, 8, 2, 3, 1}))
	//fmt.Println(findDisappearedNumbers([]int{1, 1, 4}))
	//fmt.Println(findDisappearedNumbers([]int{1, 1}))
}

//给你一个含 n 个整数的数组 nums ，其中 nums[i] 在区间 [1, n] 内。请你找出所有在 [1, n] 范围内但没有出现在 nums 中的数
//字，并以数组的形式返回结果。
//
//
//
// 示例 1：
//
//
//输入：nums = [4,3,2,7,8,2,3,1]
//输出：[5,6]
//
//
// 示例 2：
//
//
//输入：nums = [1,1]
//输出：[2]
//
//
//
//
// 提示：
//
//
// n == nums.length
// 1 <= n <= 10⁵
// 1 <= nums[i] <= n
//
//
// 进阶：你能在不使用额外空间且时间复杂度为 O(n) 的情况下解决这个问题吗? 你可以假定返回的数组不算在额外空间内。
// Related Topics 数组 哈希表 👍 865 👎 0

//leetcode submit region begin(Prohibit modification and deletion)
func findDisappearedNumbers(nums []int) []int {
	/// n = 8
	// [0:1  1:2   2:2   3:3   4:3  5:4  6:7   7:8]
	// [0:9, 1:18, 2:18, 3:11, 4:3, 5:4, 6:15, 7:16,]
	// (v - 1) % n
	// 1-1%8 = 0
	// 2-1%8 = 1

	var (
		n   = len(nums)
		ret []int
	)

	for _, v := range nums {
		v = (v - 1) % n // v 降至 index 范围 (0 ~ n-1)
		nums[v] += n    // 对应 v index 数值 + n， 对应 v index 肯定超过 n， 那么缺失index 数值 肯定小于等于 n
	}
	// 最后遍历，数值小于等于n的为确实值。
	// i+1 对应 v-1
	for i, v := range nums {
		if v <= n {
			ret = append(ret, i+1)
		}
	}

	return ret

}

//leetcode submit region end(Prohibit modification and deletion)

func findDisappearedNumbers_A(nums []int) []int {
	sort.Ints(nums)
	var (
		n   = len(nums)
		idx = 0
		ret []int
	)

	if n == 0 {
		return ret
	}

	nums = append(nums, n+1) // 哨兵模式

	// 4,3,2,7,8,2,3,1
	// 1 2 2 3 3 4 7 8
	for cur := 1; cur <= n && idx <= n; cur++ {
		if cur == nums[idx] {
			idx++
		} else if cur < nums[idx] {
			ret = append(ret, cur)
		}

		for idx < n && cur >= nums[idx] {
			idx++
		}

	}

	return ret

}
