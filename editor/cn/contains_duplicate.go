package main

import "fmt"

func main() {
	var t = 0
	fmt.Println(t & 1)
	fmt.Println(1 << 2)
	fmt.Println(1<<2 | 0)
	fmt.Println(1 << 2 & 1)
}

//给定一个整数数组，判断是否存在重复元素。
//
// 如果存在一值在数组中出现至少两次，函数返回 true 。如果数组中每个元素都不相同，则返回 false 。
//
//
//
// 示例 1:
//
//
//输入: [1,2,3,1]
//输出: true
//
// 示例 2:
//
//
//输入: [1,2,3,4]
//输出: false
//
// 示例 3:
//
//
//输入: [1,1,1,3,3,4,3,2,4,2]
//输出: true
// Related Topics 数组 哈希表 排序 👍 549 👎 0

//leetcode submit region begin(Prohibit modification and deletion)
func containsDuplicate(nums []int) bool {
	var m = make(map[int]bool)

	for _, v := range nums {
		if _, has := m[v]; has {
			return true
		}
		m[v] = true
	}

	//var flag int64
	//for _, v := range nums {
	//	if flag&(1<<v) > 0 {
	//		return true
	//	}
	//
	//	flag |= 1 << v
	//
	//}
	return false
}

//leetcode submit region end(Prohibit modification and deletion)
