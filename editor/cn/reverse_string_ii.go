package main

import "fmt"

func main() {
	fmt.Println(reverseStr("abcdefg", 8))
}

//给定一个字符串 s 和一个整数 k，从字符串开头算起，每计数至 2k 个字符，就反转这 2k 字符中的前 k 个字符。
//
//
// 如果剩余字符少于 k 个，则将剩余字符全部反转。
// 如果剩余字符小于 2k 但大于或等于 k 个，则反转前 k 个字符，其余字符保持原样。
//
//
//
//
// 示例 1：
//
//
//输入：s = "abcdefg", k = 2
//输出："bacdfeg"
//
//
// 示例 2：
//
//
//输入：s = "abcd", k = 2
//输出："bacd"
//
//
//
//
// 提示：
//
//
// 1 <= s.length <= 10⁴
// s 仅由小写英文组成
// 1 <= k <= 10⁴
//
// Related Topics 双指针 字符串 👍 228 👎 0

//leetcode submit region begin(Prohibit modification and deletion)
func reverseStr(s string, k int) string {
	var (
		bs         = []byte(s)
		n          = len(s)
		i          = 0
		front, end = i * k, (i+1)*k - 1
		reverse    = func(left, right int) {
			for left < right {
				bs[left], bs[right] = bs[right], bs[left]
				left++
				right--
			}
		}
	)

	for end < n {
		reverse(front, end)

		i = i + 2
		front, end = i*k, (i+1)*k-1
	}

	if front < n && n < end+1 {
		reverse(front, n-1)
	}

	return string(bs)
}

//leetcode submit region end(Prohibit modification and deletion)
