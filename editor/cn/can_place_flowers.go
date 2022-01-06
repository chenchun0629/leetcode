package main

import "fmt"

func main() {
	fmt.Println(canPlaceFlowers([]int{0}, 1))
}

//假设有一个很长的花坛，一部分地块种植了花，另一部分却没有。可是，花不能种植在相邻的地块上，它们会争夺水源，两者都会死去。
//
// 给你一个整数数组 flowerbed 表示花坛，由若干 0 和 1 组成，其中 0 表示没种植花，1 表示种植了花。另有一个数 n ，能否在不打破种植规则
//的情况下种入 n 朵花？能则返回 true ，不能则返回 false。
//
//
//
// 示例 1：
//
//
//输入：flowerbed = [1,0,0,0,1], n = 1
//输出：true
//
//
// 示例 2：
//
//
//输入：flowerbed = [1,0,0,0,1], n = 2
//输出：false
//
//
//
//
// 提示：
//
//
// 1 <= flowerbed.length <= 2 * 10⁴
// flowerbed[i] 为 0 或 1
// flowerbed 中不存在相邻的两朵花
// 0 <= n <= flowerbed.length
//
// Related Topics 贪心 数组 👍 416 👎 0

//leetcode submit region begin(Prohibit modification and deletion)
func canPlaceFlowers(flowerbed []int, n int) bool {
	// 0 1 => 1 / 2 = 0
	// 0 0 1 => 2 / 2 = 1
	// 0 0 0 1 => 3 / 2 = 1
	// 0 1 1 => (2 - 1 - 2) / 2 = 0
	// 0 1 0 1 => (3 - 1 - 2) / 2 = 0
	// 0 1 0 0 1 =>  (4 - 1 - 2) / 2 = 0
	// 0 1 0 0 0 1 => (5 - 1 - 2) / 2 = 1

	var (
		l     = len(flowerbed)
		count int
		prev  = -1
	)

	for i := 0; i < l; i++ {
		if flowerbed[i] == 1 {
			if prev < 0 {
				count += (i / 2)
			} else {
				count += (i - prev - 2) / 2
			}
			if count >= n {
				return true
			}
			prev = i
		}
	}

	if prev < 0 {
		count += (l + 1) / 2
	} else {
		count += (l - prev - 1) / 2
	}

	return count >= n
}

//leetcode submit region end(Prohibit modification and deletion)

func canPlaceFlowers_A(flowerbed []int, n int) bool {
	var (
		l = len(flowerbed)
	)

	for i := 0; i < l && n > 0; i++ {
		if i == 0 && l >= 2 && flowerbed[0] == 0 && flowerbed[1] == 0 {
			n--
			flowerbed[i] = 1
		}

		if i > 0 && i < l-1 && flowerbed[i] == 0 && flowerbed[i-1] == 0 && flowerbed[i+1] == 0 {
			n--
			flowerbed[i] = 1
		}

		if i == l-1 && l >= 2 && flowerbed[i] == 0 && flowerbed[i-1] == 0 {
			n--
			flowerbed[i] = 1
		}

		if l == 1 && flowerbed[i] == 0 {
			n--
			flowerbed[i] = 1
		}
	}

	return n <= 0
}
