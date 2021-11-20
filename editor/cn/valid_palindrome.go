package main

import (
	"fmt"
	"regexp"
	"strings"
)

func main() {
	//fmt.Println(isPalindrome("A man, a plan, a canal: Panama"))
	//fmt.Println(isPalindrome(")(&@$%^"))
}

//给定一个字符串，验证它是否是回文串，只考虑字母和数字字符，可以忽略字母的大小写。
//
// 说明：本题中，我们将空字符串定义为有效的回文串。
//
//
//
// 示例 1:
//
//
//输入: "A man, a plan, a canal: Panama"
//输出: true
//解释："amanaplanacanalpanama" 是回文串
//
//
// 示例 2:
//
//
//输入: "race a car"
//输出: false
//解释："raceacar" 不是回文串
//
//
//
//
// 提示：
//
//
// 1 <= s.length <= 2 * 10⁵
// 字符串 s 由 ASCII 字符组成
//
// Related Topics 双指针 字符串 👍 437 👎 0

//leetcode submit region begin(Prohibit modification and deletion)
func isPalindrome(s string) bool {
	s = strings.ToLower(s)
	fmt.Println(s)
	var b, e = 0, len(s) - 1
	for b < e {
		for b < e && !isAlphabet(s[b]) {
			b++
		}

		for e > b && !isAlphabet(s[e]) {
			e--
		}

		if b > e || s[b] != s[e] {
			return false
		}

		b++
		e--
	}

	return true
}

func isAlphabet(b byte) bool {
	return b >= 'A' && b <= 'Z' || b >= 'a' && b <= 'z' || b >= '0' && b <= '9'
}

//leetcode submit region end(Prohibit modification and deletion)

func isPalindrome_II(s string) bool {
	s = strings.ToLower(s)
	var re, _ = regexp.Compile("[^a-zA-Z0-9]")
	s = re.ReplaceAllString(s, "")

	var b, e = 0, len(s) - 1
	for b < e {
		if s[b] != s[e] {
			return false
		}
		b++
		e--
	}
	return true
}
