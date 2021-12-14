package main

import "fmt"

func main() {
	var s = "bb"
	fmt.Println(longestPalindrome(s))
}

//给你一个字符串 s，找到 s 中最长的回文子串。
//
//
//
// 示例 1：
//
//
//输入：s = "babad"
//输出："bab"
//解释："aba" 同样是符合题意的答案。
//
//
// 示例 2：
//
//
//输入：s = "cbbd"
//输出："bb"
//
//
// 示例 3：
//
//
//输入：s = "a"
//输出："a"
//
//
// 示例 4：
//
//
//输入：s = "ac"
//输出："a"
//
//
//
//
// 提示：
//
//
// 1 <= s.length <= 1000
// s 仅由数字和英文字母（大写和/或小写）组成
//
// Related Topics 字符串 动态规划 👍 4451 👎 0

//leetcode submit region begin(Prohibit modification and deletion)
func longestPalindrome_1(s string) string {
	var fn func(s string, idx int, longest string) string
	fn = func(s string, idx int, longest string) string {
		var (
			sn = len(s)
			ln = len(longest)
		)

		if sn <= idx {
			return ""
		}

		var left, right = idx - 1, idx + 1
		for left >= 0 && right < sn && s[left] == s[right] {
			if right-left+1 > ln {
				longest = s[left : right+1]
			}

			left--
			right++
		}

		left, right = idx, idx+1
		for left >= 0 && right < sn && s[left] == s[right] {
			if right-left+1 > ln {
				longest = s[left : right+1]
			}

			left--
			right++
		}

		idx += 1
		if tmp := fn(s, idx, longest); len(longest) < len(tmp) {
			longest = tmp
		}

		return longest
	}

	var longest = fn(s, 0, "")
	if longest == "" && len(s) > 0 {
		longest = s[0:1]
	}

	return longest
}

//leetcode submit region end(Prohibit modification and deletion)
