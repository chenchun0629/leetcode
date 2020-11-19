package main

import "fmt"

//编写一个函数来查找字符串数组中的最长公共前缀。
//
// 如果不存在公共前缀，返回空字符串 ""。
//
// 示例 1:
//
// 输入: ["flower","flow","flight"]
//输出: "fl"
//
//
// 示例 2:
//
// 输入: ["dog","racecar","car"]
//输出: ""
//解释: 输入不存在公共前缀。
//
//
// 说明:
//
// 所有输入只包含小写字母 a-z 。
// Related Topics 字符串
// 👍 1348 👎 0

//leetcode submit region begin(Prohibit modification and deletion)
func longestCommonPrefix(strs []string) string {
	return longestCommonPrefixA(strs)
}

// 方案A
func longestCommonPrefixA(strs []string) string {
	if len(strs) <= 0 {
		return ""
	}

	if len(strs) == 1 {
		return strs[0]
	}

	var (
		same  = make([]rune, 0)
		first = strs[0]
	)
	strs = strs[0:]
	for k, v := range first {
		for _, str := range strs {
			var strl = len(str)
			if k >= strl || string(str[k]) != string(v) {
				return string(same)
			}
		}

		same = append(same, v)
	}

	return string(same)
}

func longestCommonPrefixB(strs []string) string {
	if len(strs) == 0 {
		return ""
	}
	var lcp func(int, int) string
	lcp = func(start, end int) string {
		if start == end {
			return strs[start]
		}
		mid := (start + end) / 2
		lcpLeft, lcpRight := lcp(start, mid), lcp(mid+1, end)
		minLength := min(len(lcpLeft), len(lcpRight))
		for i := 0; i < minLength; i++ {
			if lcpLeft[i] != lcpRight[i] {
				return lcpLeft[:i]
			}
		}
		return lcpLeft[:minLength]
	}
	return lcp(0, len(strs)-1)
}

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}

func longestCommonPrefixC(strs []string) string {
	if len(strs) == 0 {
		return ""
	}

	isCommonPrefix := func(length int) bool {
		str0, count := strs[0][:length], len(strs)
		for i := 1; i < count; i++ {
			if strs[i][:length] != str0 {
				return false
			}
		}
		return true
	}
	minLength := len(strs[0])
	for _, s := range strs {
		if len(s) < minLength {
			minLength = len(s)
		}
	}
	low, high := 0, minLength
	for low < high {
		mid := (high-low+1)/2 + low
		if isCommonPrefix(mid) {
			low = mid
		} else {
			high = mid - 1
		}
	}
	return strs[0][:low]
}

//leetcode submit region end(Prohibit modification and deletion)

func main() {
	fmt.Println(longestCommonPrefixC([]string{"flower", "flow", "flight"}))
	//fmt.Println(longestCommonPrefix([]string{"flower", "flow", "flight"}))
}
