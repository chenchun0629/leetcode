package main

import (
	"fmt"
	"sort"
)

func main() {
	//fmt.Println(checkIfExist([]int{10,2,5,3}))
	//fmt.Println(checkIfExist([]int{-10, 12, -20, -8, 15}))
	fmt.Println(checkIfExist([]int{-2, 0, 10, -19, 4, 6, -8}))
}

//给你一个整数数组 arr，请你检查是否存在两个整数 N 和 M，满足 N 是 M 的两倍（即，N = 2 * M）。
//
// 更正式地，检查是否存在两个下标 i 和 j 满足：
//
//
// i != j
// 0 <= i, j < arr.length
// arr[i] == 2 * arr[j]
//
//
//
//
// 示例 1：
//
// 输入：arr = [10,2,5,3]
//输出：true
//解释：N = 10 是 M = 5 的两倍，即 10 = 2 * 5 。
//
//
// 示例 2：
//
// 输入：arr = [7,1,14,11]
//输出：true
//解释：N = 14 是 M = 7 的两倍，即 14 = 2 * 7 。
//
//
// 示例 3：
//
// 输入：arr = [3,1,7,11]
//输出：false
//解释：在该情况下不存在 N 和 M 满足 N = 2 * M 。
//
//
//
//
// 提示：
//
//
// 2 <= arr.length <= 500
// -10^3 <= arr[i] <= 10^3
//
// Related Topics 数组 哈希表 双指针 二分查找 排序 👍 46 👎 0

//leetcode submit region begin(Prohibit modification and deletion)
func checkIfExist(arr []int) bool {
	sort.Ints(arr)
	var (
		n = len(arr)
	)

	for i := 0; i < n; i++ {
		var (
			head, tail = 0, n - 1
			target     = 2 * arr[i]
		)

		for head <= tail {
			var mid = head + (tail-head)>>1
			if arr[mid] == target && mid != i {
				return true
			}

			if arr[mid] < target {
				head = mid + 1
			} else {
				tail = mid - 1
			}
		}
	}

	return false
}

//leetcode submit region end(Prohibit modification and deletion)
