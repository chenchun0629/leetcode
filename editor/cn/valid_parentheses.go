package main

import "fmt"

//给定一个只包括 '('，')'，'{'，'}'，'['，']' 的字符串，判断字符串是否有效。
//
// 有效字符串需满足：
//
//
// 左括号必须用相同类型的右括号闭合。
// 左括号必须以正确的顺序闭合。
//
//
// 注意空字符串可被认为是有效字符串。
//
// 示例 1:
//
// 输入: "()"
//输出: true
//
//
// 示例 2:
//
// 输入: "()[]{}"
//输出: true
//
//
// 示例 3:
//
// 输入: "(]"
//输出: false
//
//
// 示例 4:
//
// 输入: "([)]"
//输出: false
//
//
// 示例 5:
//
// 输入: "{[]}"
//输出: true
// Related Topics 栈 字符串
// 👍 2001 👎 0

//leetcode submit region begin(Prohibit modification and deletion)
func isValid(s string) bool {
	return isValidA(s)
}

// 方案A：
// 括号的规则是其的对称性，只要不对称就返回false
// 通过栈， 判断是否字符 {[( 就进行 压栈
// 反之就是字符 }])，从栈中取值，
// 然后进行对比是否一堆括号
func isValidA(s string) bool {
	var l = len(s)
	if l == 0 || l%2 != 0 {
		return false
	}

	var (
		stack = make([]byte, 0)
	)

	for i := 0; i < l; i++ {
		if _, has := symmetry[s[i]]; has {
			stack = append(stack, s[i])
		} else {
			if len(stack) <= 0 {
				return false
			}
			var (
				ls   = len(stack) - 1
				left = stack[ls]
			)
			if symmetry[left] != s[i] {
				return false
			}

			stack = stack[:ls]
		}
	}

	if len(stack) > 0 {
		return false
	}

	return true
}

var symmetry = map[byte]byte{
	'(': ')',
	'[': ']',
	'{': '}',
}

//leetcode submit region end(Prohibit modification and deletion)

func main() {
	fmt.Println(isValid("(("))
	//fmt.Println(isValid("()[]"))
}
