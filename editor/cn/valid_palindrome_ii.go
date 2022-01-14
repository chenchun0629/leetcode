package main

import "fmt"

func main() {
	// qmgmlcupuuuupuculmgmq
	fmt.Println(validPalindrome("aguokepatgbnvfqmgmlcupuufxoohdfpgjdmysgvhmvffcnqxjjxqncffvmhvgsymdjgpfdhooxfuupuculmgmqfvnbgtapekouga"))
	fmt.Println(validPalindrome("qmgmlcupuuuupuculmgmq"))
	fmt.Println(validPalindrome("ababaa"))
	fmt.Println(validPalindrome("aabaaa"))
}

//给定一个非空字符串 s，最多删除一个字符。判断是否能成为回文字符串。
//
//
//
// 示例 1:
//
//
//输入: s = "aba"
//输出: true
//
//
// 示例 2:
//
//
//输入: s = "abca"
//输出: true
//解释: 你可以删除c字符。
//
//
// 示例 3:
//
//
//输入: s = "abc"
//输出: false
//
//
//
// 提示:
//
//
// 1 <= s.length <= 10⁵
// s 由小写英文字母组成
//
// Related Topics 贪心 双指针 字符串 👍 440 👎 0

//leetcode submit region begin(Prohibit modification and deletion)
func validPalindrome(s string) bool {
	var (
		n           = len(s)
		left, right = 0, n - 1
		fn          func(left, right int, fixed bool) bool
	)

	fn = func(left, right int, fixed bool) bool {
		for left < right {
			if s[left] == s[right] {
				left++
				right--
				continue
			}
			break
		}

		if left < right {
			if fixed {
				return false
			}

			return fn(left+1, right, true) || fn(left, right-1, true)
		}

		return true
	}

	return fn(left, right, false)
}

//leetcode submit region end(Prohibit modification and deletion)

//aguokepatgbnvfqmgml cupuufxoohdfpgjdmysgvhmvffcnqxj
//aguokepatgbnvfqmgmlucupuufxoohdfppjdmysgvhmvffcnqxj
