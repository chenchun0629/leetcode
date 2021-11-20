package main

import "fmt"

func main() {
	var single = 0
	fmt.Println(single ^ 1)
	fmt.Println(single ^ 1 ^ 2)
	fmt.Println(single ^ 1 ^ 1)
	fmt.Println(single ^ 1 ^ 2 ^ 1 ^ 7)
	fmt.Println(single ^ 1 ^ 2 ^ 1 ^ 7 ^ 2)
}

//给定一个非空整数数组，除了某个元素只出现一次以外，其余每个元素均出现两次。找出那个只出现了一次的元素。
//
// 说明：
//
// 你的算法应该具有线性时间复杂度。 你可以不使用额外空间来实现吗？
//
// 示例 1:
//
// 输入: [2,2,1]
//输出: 1
//
//
// 示例 2:
//
// 输入: [4,1,2,1,2]
//输出: 4
// Related Topics 位运算 数组 👍 2112 👎 0

//leetcode submit region begin(Prohibit modification and deletion)
func singleNumber(nums []int) int {
	var d = 0

	for _, v := range nums {
		d ^= v
	}

	return d
}

//leetcode submit region end(Prohibit modification and deletion)

func singleNumber_A(nums []int) int {
	var m = make(map[int]int)

	for _, v := range nums {
		if _, has := m[v]; !has {
			m[v] = 0
		}
		m[v]++
	}

	for k, v := range m {
		if v == 1 {
			return k
		}
	}

	return 0
}
