package main

import "fmt"

func main() {
	// 1  0001
	// 2  0010
	// 3  0011
	// 4  0100
	// 5  0101
	// 11 1011
	// 13 1101
	// 15 1111

	fmt.Println(hammingDistance(1, 4))

	//fmt.Println(1 & 1)
	//fmt.Println(2 & 1)
	//fmt.Println(3 & 1)
	//fmt.Println(4 & 1)
	//fmt.Println(5 & 1)
	//fmt.Println(6 & 1)
	//fmt.Println(7 & 1)
	//fmt.Println("================")
	//
	fmt.Println(1 >> 1)
	fmt.Println(1 >> 1 >> 1)
	fmt.Println(4 >> 1)
	fmt.Println(4 >> 1 >> 1)
	fmt.Println(4 >> 1 >> 1 >> 1)
	fmt.Println(4 >> 1 >> 1 >> 1 >> 1)
	fmt.Println(4 >> 1 >> 1 >> 1 >> 1 >> 1)
	fmt.Println("================")

}

//两个整数之间的 汉明距离 指的是这两个数字对应二进制位不同的位置的数目。
//
// 给你两个整数 x 和 y，计算并返回它们之间的汉明距离。
//
//
//
// 示例 1：
//
//
//输入：x = 1, y = 4
//输出：2
//解释：
//1   (0 0 0 1)
//4   (0 1 0 0)
//       ↑   ↑
//上面的箭头指出了对应二进制位不同的位置。
//
//
// 示例 2：
//
//
//输入：x = 3, y = 1
//输出：1
//
//
//
//
// 提示：
//
//
// 0 <= x, y <= 2³¹ - 1
//
// Related Topics 位运算 👍 547 👎 0

//leetcode submit region begin(Prohibit modification and deletion)
func hammingDistance(x int, y int) int {
	var c int

	for x > 0 || y > 0 {
		if x == 0 {
			if 1 == y&1 {
				c++
			}
			y = y >> 1
		} else if y == 0 {
			if 1 == x&1 {
				c++
			}
			x = x >> 1
		} else {
			if x&1 != y&1 {
				c++
			}
			x = x >> 1
			y = y >> 1
		}
	}

	return c
}

//leetcode submit region end(Prohibit modification and deletion)
