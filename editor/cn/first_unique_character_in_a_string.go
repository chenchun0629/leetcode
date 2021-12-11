package main

import "fmt"

func main() {
	fmt.Println(firstUniqChar("leetcode"))
}

//给定一个字符串，找到它的第一个不重复的字符，并返回它的索引。如果不存在，则返回 -1。
//
//
//
// 示例：
//
// s = "leetcode"
//返回 0
//
//s = "loveleetcode"
//返回 2
//
//
//
//
// 提示：你可以假定该字符串只包含小写字母。
// Related Topics 队列 哈希表 字符串 计数 👍 479 👎 0

//leetcode submit region begin(Prohibit modification and deletion)
func firstUniqChar(s string) int {
	var mp = make(map[rune]int)
	for _, v := range s {
		mp[v-'a']++
	}

	for k, v := range s {
		if mp[v-'a'] == 1 {
			return k
		}
	}

	return -1
}

//leetcode submit region end(Prohibit modification and deletion)

func firstUniqChar_A(s string) int {
	var left, right = 0, 0
	var n = len(s)

	for left < n {
		if right >= n {
			return left
		}

		if left != right && s[left] == s[right] {
			left++
			right = 0
		} else {
			right++
		}
	}

	return -1
}
