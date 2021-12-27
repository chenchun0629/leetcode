package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(findWords([]string{"Hello", "Alaska", "Dad", "Peace"}))
}

//给你一个字符串数组 words ，只返回可以使用在 美式键盘 同一行的字母打印出来的单词。键盘如下图所示。
//
// 美式键盘 中：
//
//
// 第一行由字符 "qwertyuiop" 组成。
// 第二行由字符 "asdfghjkl" 组成。
// 第三行由字符 "zxcvbnm" 组成。
//
//
//
//
//
//
// 示例 1：
//
//
//输入：words = ["Hello","Alaska","Dad","Peace"]
//输出：["Alaska","Dad"]
//
//
// 示例 2：
//
//
//输入：words = ["omk"]
//输出：[]
//
//
// 示例 3：
//
//
//输入：words = ["adsdf","sfd"]
//输出：["adsdf","sfd"]
//
//
//
//
// 提示：
//
//
// 1 <= words.length <= 20
// 1 <= words[i].length <= 100
// words[i] 由英文字母（小写和大写字母）组成
//
// Related Topics 数组 哈希表 字符串 👍 193 👎 0

//leetcode submit region begin(Prohibit modification and deletion)
func findWords(words []string) []string {
	var (
		//"qwertyuiop", "asdfghjkl", "zxcvbnm"
		lines = []map[byte]struct{}{
			{'q': struct{}{}, 'w': struct{}{}, 'e': struct{}{}, 'r': struct{}{}, 't': struct{}{}, 'y': struct{}{}, 'u': struct{}{}, 'i': struct{}{}, 'o': struct{}{}, 'p': struct{}{}},
			{'a': struct{}{}, 's': struct{}{}, 'd': struct{}{}, 'f': struct{}{}, 'g': struct{}{}, 'h': struct{}{}, 'j': struct{}{}, 'k': struct{}{}, 'l': struct{}{}},
			{'z': struct{}{}, 'x': struct{}{}, 'c': struct{}{}, 'v': struct{}{}, 'b': struct{}{}, 'n': struct{}{}, 'm': struct{}{}},
		}
		ret       []string
		inOneLine = func(word string, line int) bool {
			for k := range word {
				//fmt.Println(v, lines[line])
				if _, has := lines[line][word[k]]; !has {
					return false
				}
			}
			return true
		}
	)

	for _, word := range words {
		var line = -1
		for k, v := range lines {
			if _, has := v[strings.ToLower(word)[0]]; has {
				line = k
				break
			}
		}

		if line == -1 {
			continue
		}

		if inOneLine(strings.ToLower(word), line) {
			ret = append(ret, word)
		}
	}

	return ret
}

//leetcode submit region end(Prohibit modification and deletion)
