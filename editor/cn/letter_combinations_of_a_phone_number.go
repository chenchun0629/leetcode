package main

import "fmt"

func main() {
	fmt.Printf("%#v \n", letterCombinations("23"))
}

//给定一个仅包含数字 2-9 的字符串，返回所有它能表示的字母组合。答案可以按 任意顺序 返回。
//
// 给出数字到字母的映射如下（与电话按键相同）。注意 1 不对应任何字母。
//
//
//
//
//
// 示例 1：
//
//
//输入：digits = "23"
//输出：["ad","ae","af","bd","be","bf","cd","ce","cf"]
//
//
// 示例 2：
//
//
//输入：digits = ""
//输出：[]
//
//
// 示例 3：
//
//
//输入：digits = "2"
//输出：["a","b","c"]
//
//
//
//
// 提示：
//
//
// 0 <= digits.length <= 4
// digits[i] 是范围 ['2', '9'] 的一个数字。
//
// Related Topics 哈希表 字符串 回溯 👍 1653 👎 0

//leetcode submit region begin(Prohibit modification and deletion)
func letterCombinations(digits string) []string {
	var (
		m = map[byte][]string{
			'2': {"a", "b", "c"},
			'3': {"d", "e", "f"},
			'4': {"g", "h", "i"},
			'5': {"j", "k", "l"},
			'6': {"m", "n", "o"},
			'7': {"p", "q", "r", "s"},
			'8': {"t", "u", "v"},
			'9': {"w", "x", "y", "z"},
		}
		fn func(digits string, idx int, prefix []string) []string
	)

	if len(digits) == 0 {
		return []string{}
	}

	fn = func(digits string, idx int, prefix []string) []string {
		if idx == len(digits) {
			return prefix
		}

		var ss []string

		for _, pv := range prefix {
			for _, sv := range m[digits[idx]] {
				ss = append(ss, pv+sv)
			}
		}

		return fn(digits, idx+1, ss)
	}

	return fn(digits, 1, m[digits[0]])
}

//leetcode submit region end(Prohibit modification and deletion)

func letterCombinations_A(digits string) []string {
	var (
		m = map[byte][]string{
			'2': {"a", "b", "c"},
			'3': {"d", "e", "f"},
			'4': {"g", "h", "i"},
			'5': {"j", "k", "l"},
			'6': {"m", "n", "o"},
			'7': {"p", "q", "r", "s"},
			'8': {"t", "u", "v"},
			'9': {"w", "x", "y", "z"},
		}
		fn  func(digits string, idx int, prefix string)
		ret []string
	)

	if len(digits) == 0 {
		return ret
	}

	fn = func(digits string, idx int, prefix string) {
		if len(digits) == idx {
			ret = append(ret, prefix)
		} else {
			for _, s := range m[digits[idx]] {
				fn(digits, idx+1, prefix+s)
			}
		}

	}

	fn(digits, 0, "")

	return ret
}
