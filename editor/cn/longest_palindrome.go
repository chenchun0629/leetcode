package main

import (
	"fmt"
)

func main() {
	//fmt.Println(len("123"))
	//fmt.Println(len("123abc"))
	//fmt.Println(len([]rune("123abc哈哈")))
	//fmt.Println(utf8.RuneCountInString("123abc哈哈"))
	fmt.Println(longestPalindrome("abccccdd"))
}

//给定一个包含大写字母和小写字母的字符串，找到通过这些字母构造成的最长的回文串。
//
// 在构造过程中，请注意区分大小写。比如 "Aa" 不能当做一个回文字符串。
//
// 注意:
//假设字符串的长度不会超过 1010。
//
// 示例 1:
//
//
//输入:
//"abccccdd"
//
//输出:
//7
//
//解释:
//我们可以构造的最长的回文串是"dccaccd", 它的长度是 7。
//
// Related Topics 贪心 哈希表 字符串 👍 362 👎 0

//leetcode submit region begin(Prohibit modification and deletion)
func longestPalindrome(s string) int {
	var (
		m      = make(map[byte]int)
		n      = len(s)
		hasOdd bool
		l      int
	)
	for i := 0; i < n; i++ {
		m[s[i]]++
	}

	for _, v := range m {
		if !hasOdd && v%2 == 1 {
			hasOdd = true
			l++
		}
		l += v / 2 * 2
	}

	return l
}

//leetcode submit region end(Prohibit modification and deletion)
