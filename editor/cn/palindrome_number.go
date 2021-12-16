package main

import "fmt"

//判断一个整数是否是回文数。回文数是指正序（从左向右）和倒序（从右向左）读都是一样的整数。
//
// 示例 1:
//
// 输入: 121
//输出: true
//
//
// 示例 2:
//
// 输入: -121
//输出: false
//解释: 从左向右读, 为 -121 。 从右向左读, 为 121- 。因此它不是一个回文数。
//
//
// 示例 3:
//
// 输入: 10
//输出: false
//解释: 从右向左读, 为 01 。因此它不是一个回文数。
//
//
// 进阶:
//
// 你能不将整数转为字符串来解决这个问题吗？
// Related Topics 数学
// 👍 1307 👎 0

//leetcode submit region begin(Prohibit modification and deletion)
func isPalindrome_I(x int) bool {
	return isPalindromeB(x)
}

// 方案B
// 按数字来处理
// 回转一般数字进行比较
func isPalindromeB(x int) bool {
	// 负数、尾数为0的肯定不是回文数
	if x != 0 && (x < 0 || x%10 == 0) {
		return false
	}

	// 这里的只要目的是将x降位，并且将降位的数添加至y上
	// 例如：12321
	// 第一步， x = 1232 y = y*10+1 = 1
	// 第二步， x = 123  y = y*10+2 = 12
	// 第三步， x = 12   u = y*10+3 = 123
	// 第四步， x < y 退出循环
	var y int
	for x > y {
		var p = x % 10
		x = x / 10
		y = y*10 + p
	}

	// 两种情况：
	// x == y => 1221 => 12 & 21
	// x == y/10 => 12321 => 12 & 321 => 12 & 21
	return x == y || x == y/10
}

// 方案A
// 转换成文字后按字符串判断
func isPalindromeA(x int) bool {
	var (
		r = []rune(fmt.Sprint(x))
		l = len(r)
	)

	for i := 0; i < l/2; i++ {
		if r[i] != r[l-i-1] {
			return false
		}
	}

	return true
}

//leetcode submit region end(Prohibit modification and deletion)
