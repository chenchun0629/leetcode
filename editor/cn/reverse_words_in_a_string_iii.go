package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(reverseWords("ab  bc"))
}

//给定一个字符串，你需要反转字符串中每个单词的字符顺序，同时仍保留空格和单词的初始顺序。
//
//
//
// 示例：
//
// 输入："Let's take LeetCode contest"
//输出："s'teL ekat edoCteeL tsetnoc"
//
//
//
//
// 提示：
//
//
// 在字符串中，每个单词由单个空格分隔，并且字符串中不会有任何额外的空格。
//
// Related Topics 双指针 字符串 👍 385 👎 0

//leetcode submit region begin(Prohibit modification and deletion)
func reverseWords(s string) string {
	var (
		ret []byte
		n   = len(s)
	)

	for i := 0; i < n; {
		var b = i
		for i < n && s[i] != ' ' {
			i++
		}

		for j := i - 1; j >= b; j-- {
			ret = append(ret, s[j])
		}

		for i < n && s[i] == ' ' {
			i++
			ret = append(ret, ' ')
		}

	}

	return string(ret)
}

//leetcode submit region end(Prohibit modification and deletion)

func reverseWords_A(s string) string {
	var (
		ret  = strings.Builder{}
		n    = len(s)
		b, e = 0, 0
		fn   = func(b, e int) {
			for i := e; i >= b; i-- {
				ret.WriteByte(s[i])
			}
		}
	)

	for b < n && e < n {
		if s[e] != ' ' {
			e++
			continue
		}

		if b < e {
			fn(b, e-1)
			ret.WriteByte(' ')
		}
		for e < n && s[e] == ' ' {
			e++
			b = e
		}
	}

	if b < e {
		fn(b, e-1)
	}

	return ret.String()
}
