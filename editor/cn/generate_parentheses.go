package main

import "fmt"

func main() {

	fmt.Println(generateParenthesis(3))
}

//数字 n 代表生成括号的对数，请你设计一个函数，用于能够生成所有可能的并且 有效的 括号组合。
//
//
//
// 示例 1：
//
//
//输入：n = 3
//输出：["((()))","(()())","(())()","()(())","()()()"]
//
//
// 示例 2：
//
//
//输入：n = 1
//输出：["()"]
//
//
//
//
// 提示：
//
//
// 1 <= n <= 8
//
// Related Topics 字符串 动态规划 回溯 👍 2226 👎 0

//leetcode submit region begin(Prohibit modification and deletion)
func generateParenthesis(n int) []string {
	var (
		ret []string
		fn  func(prefix []byte, left, right int)
	)

	fn = func(prefix []byte, left, right int) {
		if left == right && left == n {
			ret = append(ret, string(prefix))
		}

		if left < n {
			fn(append(prefix, '('), left+1, right)
		}

		if left > right {
			fn(append(prefix, ')'), left, right+1)
		}
	}

	if n > 0 {
		fn([]byte{'('}, 1, 0)
	}

	return ret
}

//leetcode submit region end(Prohibit modification and deletion)

func generateParenthesis_A(n int) []string {
	var (
		ret []string
		fn  func(prefix string, left, right int)
	)

	fn = func(prefix string, left, right int) {
		if left == right && left == n {
			ret = append(ret, prefix)
		}

		if left < n {
			fn(prefix+"(", left+1, right)
		}

		if left > right {
			fn(prefix+")", left, right+1)
		}
	}

	if n > 0 {
		fn("(", 1, 0)
	}

	return ret
}
