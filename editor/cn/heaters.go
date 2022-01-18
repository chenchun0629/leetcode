package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println(findRadius([]int{1, 1, 1, 1, 1, 1, 999, 999, 999, 999, 999}, []int{499, 500, 501}))
	//fmt.Println(findRadius([]int{1, 2, 3}, []int{4}))
	//fmt.Println(findRadius([]int{1, 2, 3, 4}, []int{1, 4}))
}

//冬季已经来临。 你的任务是设计一个有固定加热半径的供暖器向所有房屋供暖。
//
// 在加热器的加热半径范围内的每个房屋都可以获得供暖。
//
// 现在，给出位于一条水平线上的房屋 houses 和供暖器 heaters 的位置，请你找出并返回可以覆盖所有房屋的最小加热半径。
//
// 说明：所有供暖器都遵循你的半径标准，加热的半径也一样。
//
//
//
// 示例 1:
//
//
//输入: houses = [1,2,3], heaters = [2]
//输出: 1
//解释: 仅在位置2上有一个供暖器。如果我们将加热半径设为1，那么所有房屋就都能得到供暖。
//
//
// 示例 2:
//
//
//输入: houses = [1,2,3,4], heaters = [1,4]
//输出: 1
//解释: 在位置1, 4上有两个供暖器。我们需要将加热半径设为1，这样所有房屋就都能得到供暖。
//
//
// 示例 3：
//
//
//输入：houses = [1,5], heaters = [2]
//输出：3
//
//
//
//
// 提示：
//
//
// 1 <= houses.length, heaters.length <= 3 * 10⁴
// 1 <= houses[i], heaters[i] <= 10⁹
//
// Related Topics 数组 双指针 二分查找 排序 👍 362 👎 0

//leetcode submit region begin(Prohibit modification and deletion)
func findRadius(houses []int, heaters []int) int {
	// 二分查找
	sort.Ints(heaters)
	var (
		max = func(a, b int) int {
			if a > b {
				return a
			}
			return b
		}
		min = func(a, b int) int {
			if a > b {
				return b
			}
			return a
		}
		ret        int
		heatersLen = len(heaters)
	)

	for _, house := range houses {
		var left, right = 0, heatersLen - 1
		for left <= right {
			var mid = (left + right) / 2
			if heaters[mid] <= house {
				if mid < heatersLen-1 && heaters[mid+1] > house {
					ret = max(ret, min(heaters[mid+1]-house, house-heaters[mid]))
					break
				} else if mid == heatersLen-1 {
					ret = max(ret, house-heaters[mid])
					break
				}

				left = mid + 1
			} else {
				if mid == 0 {
					ret = max(ret, heaters[mid]-house)
					break
				} else if mid > 0 && heaters[mid-1] < house {
					ret = max(ret, min(house-heaters[mid-1], heaters[mid]-house))
					break
				}
				right = mid - 1
			}
		}
	}

	return ret
}

//leetcode submit region end(Prohibit modification and deletion)

func findRadius_A(houses []int, heaters []int) int {
	sort.Ints(houses)
	sort.Ints(heaters)
	var (
		max = func(a, b int) int {
			if a > b {
				return a
			}
			return b
		}
		min = func(a, b int) int {
			if a > b {
				return b
			}
			return a
		}
		ret         int
		heatersLen  = len(heaters)
		heaterIndex int
	)

	for _, house := range houses {
		for heaterIndex < heatersLen && heaters[heaterIndex] < house {
			// 供暖器在前面
			heaterIndex++
		}

		if heaterIndex == 0 {
			ret = max(ret, heaters[heaterIndex]-house)
		} else if heaterIndex == heatersLen {
			ret = max(ret, houses[len(houses)-1]-heaters[heaterIndex-1])
		} else {
			ret = max(ret, min(
				heaters[heaterIndex]-house,
				house-heaters[heaterIndex-1],
			))
		}
	}

	return ret
}
