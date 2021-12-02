package main

func main() {

}

//给定两个字符串 s 和 t ，编写一个函数来判断 t 是否是 s 的字母异位词。
//
// 注意：若 s 和 t 中每个字符出现的次数都相同，则称 s 和 t 互为字母异位词。
//
//
//
// 示例 1:
//
//
//输入: s = "anagram", t = "nagaram"
//输出: true
//
//
// 示例 2:
//
//
//输入: s = "rat", t = "car"
//输出: false
//
//
//
// 提示:
//
//
// 1 <= s.length, t.length <= 5 * 10⁴
// s 和 t 仅包含小写字母
//
//
//
//
// 进阶: 如果输入字符串包含 unicode 字符怎么办？你能否调整你的解法来应对这种情况？
// Related Topics 哈希表 字符串 排序 👍 459 👎 0

//leetcode submit region begin(Prohibit modification and deletion)
func isAnagram(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}

	var sm = make(map[rune]int)
	var tm = make(map[rune]int)

	for _, v := range s {
		sm[v]++
	}
	for _, v := range t {
		tm[v]++
	}

	for k, v := range sm {
		if tv, has := tm[k]; !has || v != tv {
			return false
		}

		delete(tm, k)
	}

	return len(tm) == 0

}

//leetcode submit region end(Prohibit modification and deletion)

func isAnagram_A(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}

	var sm = make(map[rune]int)
	var tm = make(map[rune]int)

	for _, v := range s {
		sm[v]++
	}
	for _, v := range t {
		tm[v]++
	}

	for k, v := range sm {
		if tv, has := tm[k]; !has || v != tv {
			return false
		}

		delete(tm, k)
	}

	return len(tm) == 0

}
