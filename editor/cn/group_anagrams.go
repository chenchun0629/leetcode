package main

import (
	"fmt"
	"sort"
)

func main() {

}

//给你一个字符串数组，请你将 字母异位词 组合在一起。可以按任意顺序返回结果列表。
//
// 字母异位词 是由重新排列源单词的字母得到的一个新单词，所有源单词中的字母通常恰好只用一次。
//
//
//
// 示例 1:
//
//
//输入: strs = ["eat", "tea", "tan", "ate", "nat", "bat"]
//输出: [["bat"],["nat","tan"],["ate","eat","tea"]]
//
// 示例 2:
//
//
//输入: strs = [""]
//输出: [[""]]
//
//
// 示例 3:
//
//
//输入: strs = ["a"]
//输出: [["a"]]
//
//
//
// 提示：
//
//
// 1 <= strs.length <= 10⁴
// 0 <= strs[i].length <= 100
// strs[i] 仅包含小写字母
//
// Related Topics 哈希表 字符串 排序 👍 953 👎 0

//leetcode submit region begin(Prohibit modification and deletion)
func groupAnagrams(strs []string) [][]string {
	var (
		ret [][]string
		idx int
		m   = make(map[string]int)
		fn  = func(b []byte) func(i, j int) bool {
			return func(i, j int) bool {
				return b[i] < b[j]
			}
		}
	)

	for _, v := range strs {
		var b = []byte(v)
		sort.Slice(b, fn(b))
		var k = fmt.Sprintf("%+v", b)

		if _, has := m[k]; !has {
			ret = append(ret, make([]string, 0))
			m[k] = idx
			idx++
		}

		ret[m[k]] = append(ret[m[k]], v)
	}

	return ret
}

//leetcode submit region end(Prohibit modification and deletion)
