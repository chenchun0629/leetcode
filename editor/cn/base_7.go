package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println(convertToBase7(0))
	fmt.Println(convertToBase7(1))
	fmt.Println(convertToBase7(2))
	fmt.Println(convertToBase7(3))
	fmt.Println(convertToBase7(4))
	fmt.Println(convertToBase7(5))
	fmt.Println(convertToBase7(6))
	fmt.Println(convertToBase7(7))
	fmt.Println(convertToBase7(8))
	fmt.Println(convertToBase7(9))

	fmt.Println(convertToBase7(100))
	fmt.Println(convertToBase7(101))
	fmt.Println(convertToBase7(102))
	fmt.Println(convertToBase7(103))
	fmt.Println(convertToBase7(104))
	fmt.Println(convertToBase7(105))
	fmt.Println(convertToBase7(106))
	fmt.Println(convertToBase7(107))
}

//给定一个整数 num，将其转化为 7 进制，并以字符串形式输出。
//
//
//
// 示例 1:
//
//
//输入: num = 100
//输出: "202"
//
//
// 示例 2:
//
//
//输入: num = -7
//输出: "-10"
//
//
//
//
// 提示：
//
//
// -10⁷ <= num <= 10⁷
//
// Related Topics 数学 👍 108 👎 0

//leetcode submit region begin(Prohibit modification and deletion)
func convertToBase7(num int) string {
	// 9 12
	// 9 / 7 = 1, 9 % 7 = 2
	var x, loop int
	var negative = 1
	if num < 0 {
		num = -num
		negative = -1
	}

	for num >= 7 {
		var mod = num % 7
		num = num / 7
		x = x + mod*int(math.Pow(10, float64(loop)))
		loop++
	}

	x = x + num*int(math.Pow(10, float64(loop)))
	x = x * negative

	return fmt.Sprintf("%d", x)
}

//leetcode submit region end(Prohibit modification and deletion)
