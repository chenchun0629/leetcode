package main

import "fmt"

func main() {
	//var s = "abcd"
	//fmt.Println(s[:0])
	//fmt.Println(s[:1])
	//fmt.Println(s[0:2])
	//fmt.Println(s[2:4])
	fmt.Println(repeatedSubstringPattern("abaababaab"))
	fmt.Println(repeatedSubstringPattern("abab"))
	fmt.Println(repeatedSubstringPattern("abcabcabcabc"))
	fmt.Println(repeatedSubstringPattern("aba"))
	fmt.Println(repeatedSubstringPattern("abac"))
}

//给定一个非空的字符串，判断它是否可以由它的一个子串重复多次构成。给定的字符串只含有小写英文字母，并且长度不超过10000。
//
// 示例 1:
//
//
//输入: "abab"
//
//输出: True
//
//解释: 可由子字符串 "ab" 重复两次构成。
//
//
// 示例 2:
//
//
//输入: "aba"
//
//输出: False
//
//
// 示例 3:
//
//
//输入: "abcabcabcabc"
//
//输出: True
//
//解释: 可由子字符串 "abc" 重复四次构成。 (或者子字符串 "abcabc" 重复两次构成。)
//
// Related Topics 字符串 字符串匹配 👍 582 👎 0

//leetcode submit region begin(Prohibit modification and deletion)
func repeatedSubstringPattern(s string) bool {
	var (
		n          = len(s)
		isRepeated = func(pattern string) bool {
			var pn = len(pattern)
			if n%pn != 0 {
				return false
			}

			for j, i := pn, pn+pn; i <= n; j, i = i, i+pn {
				if pattern != s[j:i] {
					return false
				}
			}

			return true
		}
	)

	for i := 1; i <= n/2; i++ {
		if isRepeated(s[:i]) {
			return true
		}
	}

	return false
}

//leetcode submit region end(Prohibit modification and deletion)
