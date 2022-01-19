package main

import (
	"fmt"
)

func main() {
	//fmt.Println(minEatingSpeed([]int{3,6,7,11}, 8))
	//fmt.Println(minEatingSpeed([]int{30,11,23,4,20}, 6))
	//fmt.Println(minEatingSpeed([]int{312884470}, 312884469))
	fmt.Println(minEatingSpeed([]int{312884470}, 968709470))
}

//珂珂喜欢吃香蕉。这里有 N 堆香蕉，第 i 堆中有 piles[i] 根香蕉。警卫已经离开了，将在 H 小时后回来。
//
// 珂珂可以决定她吃香蕉的速度 K （单位：根/小时）。每个小时，她将会选择一堆香蕉，从中吃掉 K 根。如果这堆香蕉少于 K 根，她将吃掉这堆的所有香蕉，然后
//这一小时内不会再吃更多的香蕉。
//
// 珂珂喜欢慢慢吃，但仍然想在警卫回来前吃掉所有的香蕉。
//
// 返回她可以在 H 小时内吃掉所有香蕉的最小速度 K（K 为整数）。
//
//
//
//
//
//
// 示例 1：
//
// 输入: piles = [3,6,7,11], H = 8
//输出: 4
//
//
// 示例 2：
//
// 输入: piles = [30,11,23,4,20], H = 5
//输出: 30
//
//
// 示例 3：
//
// 输入: piles = [30,11,23,4,20], H = 6
//输出: 23
//
//
//
//
// 提示：
//
//
// 1 <= piles.length <= 10^4
// piles.length <= H <= 10^9
// 1 <= piles[i] <= 10^9
//
// Related Topics 数组 二分查找 👍 243 👎 0

//leetcode submit region begin(Prohibit modification and deletion)
func minEatingSpeed(piles []int, h int) int {
	//sort.Ints(piles)
	var (
		minV, maxV int
		fn         = func(k int) int {
			var t int
			for _, v := range piles {
				t = t + v/k
				if v%k > 0 {
					t++
				}
			}
			return t
		}
	)

	for _, v := range piles {
		if v > maxV {
			maxV = v
		}
	}

	for minV < maxV {
		var (
			mid = (minV + maxV) / 2
		)

		if 0 == mid {
			return 1
		}

		var tmp = fn(mid)

		//if tmp == h {
		//	return mid
		//}

		if tmp > h {
			minV = mid + 1
		} else {
			maxV = mid
		}
	}

	return minV
}

//leetcode submit region end(Prohibit modification and deletion)
