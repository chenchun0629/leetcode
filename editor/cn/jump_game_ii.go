package main

import "fmt"

func main() {
	fmt.Println(jump([]int{2, 3, 1, 1, 4}))
	fmt.Println(jump([]int{1, 2, 3}))
	fmt.Println(jump([]int{1, 2, 1, 1, 1}))
}

//给你一个非负整数数组 nums ，你最初位于数组的第一个位置。
//
// 数组中的每个元素代表你在该位置可以跳跃的最大长度。
//
// 你的目标是使用最少的跳跃次数到达数组的最后一个位置。
//
// 假设你总是可以到达数组的最后一个位置。
//
//
//
// 示例 1:
//
//
//输入: nums = [2,3,1,1,4]
//输出: 2
//解释: 跳到最后一个位置的最小跳跃数是 2。
//     从下标为 0 跳到下标为 1 的位置，跳 1 步，然后跳 3 步到达数组的最后一个位置。
//
//
// 示例 2:
//
//
//输入: nums = [2,3,0,1,4]
//输出: 2
//
//
//
//
// 提示:
//
//
// 1 <= nums.length <= 10⁴
// 0 <= nums[i] <= 1000
//
// Related Topics 贪心 数组 动态规划 👍 1334 👎 0

//leetcode submit region begin(Prohibit modification and deletion)
func jump(nums []int) int {
	max := func(x, y int) int {
		if x > y {
			return x
		}
		return y
	}
	length := len(nums)
	end := 0
	maxPosition := 0
	steps := 0
	for i := 0; i < length-1; i++ {
		maxPosition = max(maxPosition, i+nums[i])
		if i == end {
			end = maxPosition
			steps++
		}
	}
	return steps
}

//leetcode submit region end(Prohibit modification and deletion)
