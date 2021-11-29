package main

import "fmt"

func main() {
	fmt.Println(isIsomorphic("addba", "eggag"))
}

//给定两个字符串 s 和 t，判断它们是否是同构的。
//
// 如果 s 中的字符可以按某种映射关系替换得到 t ，那么这两个字符串是同构的。
//
// 每个出现的字符都应当映射到另一个字符，同时不改变字符的顺序。不同字符不能映射到同一个字符上，相同字符只能映射到同一个字符上，字符可以映射到自己本身。
//
//
//
// 示例 1:
//
//
//输入：s = "egg", t = "add"
//输出：true
//
//
// 示例 2：
//
//
//输入：s = "foo", t = "bar"
//输出：false
//
// 示例 3：
//
//
//输入：s = "paper", t = "title"
//输出：true
//
//
//
// 提示：
//
//
// 可以假设 s 和 t 长度相同。
//
// Related Topics 哈希表 字符串 👍 407 👎 0

//leetcode submit region begin(Prohibit modification and deletion)
func isIsomorphic(s string, t string) bool {
	var (
		sm = make(map[byte]byte)
		tm = make(map[byte]byte)
	)

	for v := range s {
		var x, y = s[v], t[v]

		if sm[x] > 0 && sm[x] != y || tm[y] > 0 && tm[y] != x {
			return false
		}

		sm[x] = y
		tm[y] = x
	}

	return true
}

//leetcode submit region end(Prohibit modification and deletion)

func isIsomorphic_A(s string, t string) bool {
	var (
		l     = len(s)
		index = 0
		ss    = make([]uint8, l)
		sm    = make(map[uint8]uint8)
		si    uint8
		st    = make([]uint8, l)
		tm    = make(map[uint8]uint8)
		ti    uint8
	)

	for ; index < l; index++ {
		if _, has := sm[s[index]]; !has {
			sm[s[index]] = si
			si++
		}
		ss[index] = sm[s[index]]

		if _, has := tm[t[index]]; !has {
			tm[t[index]] = ti
			ti++
		}
		st[index] = tm[t[index]]

		if string(ss) != string(st) {
			return false
		}
	}

	return true
}
