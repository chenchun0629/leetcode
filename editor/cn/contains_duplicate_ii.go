package main

func main() {

}

//给定一个整数数组和一个整数 k，判断数组中是否存在两个不同的索引 i 和 j，使得 nums [i] = nums [j]，并且 i 和 j 的差的 绝对值
// 至多为 k。
//
//
//
// 示例 1:
//
// 输入: nums = [1,2,3,1], k = 3
//输出: true
//
// 示例 2:
//
// 输入: nums = [1,0,1,1], k = 1
//输出: true
//
// 示例 3:
//
// 输入: nums = [1,2,3,1,2,3], k = 2
//输出: false
// Related Topics 数组 哈希表 滑动窗口 👍 338 👎 0

//leetcode submit region begin(Prohibit modification and deletion)
func containsNearbyDuplicate(nums []int, k int) bool {
	var m = make(map[int]int)

	for ki, v := range nums {
		if i, has := m[v]; has && ki-i <= k {
			return true
		}
		m[v] = ki
	}
	return false
}

//leetcode submit region end(Prohibit modification and deletion)
