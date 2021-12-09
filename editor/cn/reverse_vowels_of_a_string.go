package main

func main() {

}

//给你一个字符串 s ，仅反转字符串中的所有元音字母，并返回结果字符串。
//
// 元音字母包括 'a'、'e'、'i'、'o'、'u'，且可能以大小写两种形式出现。
//
//
//
// 示例 1：
//
//
//输入：s = "hello"
//输出："holle"
//
//
// 示例 2：
//
//
//输入：s = "leetcode"
//输出："leotcede"
//
//
//
// 提示：
//
//
// 1 <= s.length <= 3 * 10⁵
// s 由 可打印的 ASCII 字符组成
//
// Related Topics 双指针 字符串 👍 230 👎 0

//leetcode submit region begin(Prohibit modification and deletion)
func reverseVowels(s string) string {
	var m = map[byte]bool{'a': true, 'e': true, 'i': true, 'o': true, 'u': true, 'A': true, 'E': true, 'I': true, 'O': true, 'U': true}
	var b = []byte(s)
	var left, right = 0, len(s) - 1
	for left < right {
		for left < right {
			if _, has := m[b[left]]; has {
				break
			}
			left++
		}

		for left < right {
			if _, has := m[b[right]]; has {
				break
			}
			right--
		}

		if left < right {
			b[left], b[right] = b[right], b[left]
			left++
			right--
		}
	}

	return string(b)
}

//leetcode submit region end(Prohibit modification and deletion)
