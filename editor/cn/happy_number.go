package main

import "fmt"

func main() {
	var a = 0
	fmt.Printf("%b \n", a|3)
	//fmt.Println(isHappy(2))
}

//编写一个算法来判断一个数 n 是不是快乐数。
//
// 「快乐数」定义为：
//
//
// 对于一个正整数，每一次将该数替换为它每个位置上的数字的平方和。
// 然后重复这个过程直到这个数变为 1，也可能是 无限循环 但始终变不到 1。
// 如果 可以变为 1，那么这个数就是快乐数。
//
//
// 如果 n 是快乐数就返回 true ；不是，则返回 false 。
//
//
//
// 示例 1：
//
//
//输入：n = 19
//输出：true
//解释：
//1² + 9² = 82
//8² + 2² = 68
//6² + 8² = 100
//1² + 0² + 0² = 1
//
//
// 示例 2：
//
//
//输入：n = 2
//输出：false
//
//
//
//
// 提示：
//
//
// 1 <= n <= 2³¹ - 1
//
// Related Topics 哈希表 数学 双指针 👍 730 👎 0

//leetcode submit region begin(Prohibit modification and deletion)
func isHappy(n int) bool {
	var m = make(map[int]bool)
	for {
		var is = isHappyNumToSlice(n)
		var sum = isHappySum(is...)
		if sum == 1 {
			return true
		}

		if _, has := m[sum]; has {
			break
		}
		m[sum] = true

		if sum*sum < 10 {
			break
		}

		n = sum
	}

	fmt.Println(m)

	return false
}

func isHappySum(is ...int) int {
	var sum int
	for _, v := range is {
		sum += isHappySquare(v)
	}

	return sum
}

func isHappyNumToSlice(n int) []int {
	var is = make([]int, 0)

	for {
		var mod = n % 10
		n = n / 10
		is = append([]int{mod}, is...)

		if n <= 0 {
			break
		}
	}

	return is
}

func isHappySquare(i int) int {
	return i * i
}

//leetcode submit region end(Prohibit modification and deletion)
