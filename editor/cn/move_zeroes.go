package main

func main() {

}

//给定一个数组 nums，编写一个函数将所有 0 移动到数组的末尾，同时保持非零元素的相对顺序。
//
// 示例:
//
// 输入: [0,1,0,3,12]
//输出: [1,3,12,0,0]
//
// 说明:
//
//
// 必须在原数组上操作，不能拷贝额外的数组。
// 尽量减少操作次数。
//
// Related Topics 数组 双指针 👍 1329 👎 0

//leetcode submit region begin(Prohibit modification and deletion)
func moveZeroes(nums []int) {
	var (
		left, right int
		n           = len(nums)
	)
	for right < n-1 {
		if nums[left] == 0 {
			right++
			if nums[right] != 0 {
				nums[left], nums[right] = nums[right], nums[left]
			}
		} else {
			left++
			if right < left {
				right = left
			}
		}
	}
}

//leetcode submit region end(Prohibit modification and deletion)
