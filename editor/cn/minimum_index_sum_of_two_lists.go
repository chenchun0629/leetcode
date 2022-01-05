package main

import "fmt"

func main() {
	//fmt.Println(findRestaurant([]string{"Shogun","Tapioca Express","Burger King","KFC"}, []string{"Piatti","The Grill at Torrey Pines","Hungry Hunter Steakhouse","Shogun"}))
	fmt.Println(findRestaurant([]string{"KFC"}, []string{"KFC"}))
}

//假设Andy和Doris想在晚餐时选择一家餐厅，并且他们都有一个表示最喜爱餐厅的列表，每个餐厅的名字用字符串表示。
//
// 你需要帮助他们用最少的索引和找出他们共同喜爱的餐厅。 如果答案不止一个，则输出所有答案并且不考虑顺序。 你可以假设总是存在一个答案。
//
// 示例 1:
//
// 输入:
//["Shogun", "Tapioca Express", "Burger King", "KFC"]
//["Piatti", "The Grill at Torrey Pines", "Hungry Hunter Steakhouse", "Shogun"]
//输出: ["Shogun"]
//解释: 他们唯一共同喜爱的餐厅是“Shogun”。
//
//
// 示例 2:
//
// 输入:
//["Shogun", "Tapioca Express", "Burger King", "KFC"]
//["KFC", "Shogun", "Burger King"]
//输出: ["Shogun"]
//解释: 他们共同喜爱且具有最小索引和的餐厅是“Shogun”，它有最小的索引和1(0+1)。
//
//
// 提示:
//
//
// 两个列表的长度范围都在 [1, 1000]内。
// 两个列表中的字符串的长度将在[1，30]的范围内。
// 下标从0开始，到列表的长度减1。
// 两个列表都没有重复的元素。
//
// Related Topics 数组 哈希表 字符串 👍 131 👎 0

//leetcode submit region begin(Prohibit modification and deletion)
func findRestaurant(list1 []string, list2 []string) []string {
	var (
		n1  = len(list1)
		n2  = len(list2)
		m   = make(map[string]int)
		ret []string
		fn  = func(a []string, b []string) {
			for k, v := range a {
				m[v] = k
			}
			for k, v := range b {
				if _, has := m[v]; has {
					m[v] = n1 + n2 + k + m[v] + 2
				}
			}
		}
	)

	if n1 < n2 {
		fn(list1, list2)
	} else {
		fn(list2, list1)
	}

	var min = (n1 + n2) * 2
	for s, t := range m {
		if t < n1+n2 {
			continue
		}
		if t < min {
			ret = []string{s}
			min = t
		} else if t == min {
			ret = append(ret, s)
		}
	}

	return ret
}

//leetcode submit region end(Prohibit modification and deletion)
