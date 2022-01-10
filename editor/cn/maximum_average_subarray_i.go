package main

import "fmt"

func main() {
	fmt.Println(findMaxAverage([]int{1, 12, -5, -6, 50, 3}, 4))
}

//给你一个由 n 个元素组成的整数数组 nums 和一个整数 k 。
//
// 请你找出平均数最大且 长度为 k 的连续子数组，并输出该最大平均数。
//
// 任何误差小于 10⁻⁵ 的答案都将被视为正确答案。
//
//
//
// 示例 1：
//
//
//输入：nums = [1,12,-5,-6,50,3], k = 4
//输出：12.75
//解释：最大平均数 (12-5-6+50)/4 = 51/4 = 12.75
//
//
// 示例 2：
//
//
//输入：nums = [5], k = 1
//输出：5.00000
//
//
//
//
// 提示：
//
//
// n == nums.length
// 1 <= k <= n <= 10⁵
// -10⁴ <= nums[i] <= 10⁴
//
// Related Topics 数组 滑动窗口 👍 218 👎 0

//leetcode submit region begin(Prohibit modification and deletion)
func findMaxAverage(nums []int, k int) float64 {
	var (
		n   = len(nums)
		max = func(a, b int) int {
			if a > b {
				return a
			}
			return b
		}
		sum = func(begin, end int) (sum int) {
			for i := begin; i < end; i++ {
				sum += nums[i]
			}
			return
		}
		sumV = sum(0, k)
		maxV = sumV
	)

	for i := k; i < n; i++ {
		sumV = sumV - nums[i-k] + nums[i]
		maxV = max(maxV, sumV)
	}

	return float64(maxV) / float64(k)
}

//leetcode submit region end(Prohibit modification and deletion)
