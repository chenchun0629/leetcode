package main

import (
	"fmt"
	"sort"
)

func main() {
	//fmt.Println(factorial(3)/(factorial(3)*factorial(3-3)))
	//fmt.Println(factorial(4)/(factorial(3)*factorial(4-3)))
	//fmt.Println(factorial(5)/(factorial(3)*factorial(5-3)))

	fmt.Println(triangleNumber([]int{2, 2, 3, 4, 7, 8, 9}))
}

//给定一个包含非负整数的数组，你的任务是统计其中可以组成三角形三条边的三元组个数。
//
// 示例 1:
//
//
//输入: [2,2,3,4]
//输出: 3
//解释:
//有效的组合是:
//2,3,4 (使用第一个 2)
//2,3,4 (使用第二个 2)
//2,2,3
//
//
// 注意:
//
//
// 数组长度不超过1000。
// 数组里整数的范围为 [0, 1000]。
//
// Related Topics 贪心 数组 双指针 二分查找 排序 👍 340 👎 0

//leetcode submit region begin(Prohibit modification and deletion)
func triangleNumber(nums []int) int {
	sort.Ints(nums)
	var (
		ans      = 0
		getIndex func(nums []int, a int) int
	)

	getIndex = func(nums []int, a int) int { //获得的索引为该索引值是小于a的最大索引
		if 0 == len(nums) {
			return -1
		}
		left, right := 0, len(nums)-1
		for left < right {
			mid := left + (right-left)/2
			if nums[mid] < a && nums[mid+1] < a {
				left = mid + 1
			} else if nums[mid] < a && nums[mid+1] >= a {
				return mid
			} else if nums[mid] >= a {
				right = mid - 1
			}
		}
		if nums[left] >= a { //没找到，返回-1
			return -1
		}
		return left
	}

	for i := 0; i < len(nums); i++ {
		for j := i + 1; j < len(nums); j++ { //因为是递增排序的，所以第三条边长度区间为，nums[j] ~ nums[i]+nums[j]，所以从j+1开始
			ans += getIndex(nums[j+1:], nums[i]+nums[j]) + 1
		}
	}
	return ans
}

//leetcode submit region end(Prohibit modification and deletion)
