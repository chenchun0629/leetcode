package main

import "fmt"

//给定一个整数数组 nums 和一个目标值 target，请你在该数组中找出和为目标值的那 两个 整数，并返回他们的数组下标。
//
// 你可以假设每种输入只会对应一个答案。但是，数组中同一个元素不能使用两遍。
//
//
//
// 示例:
//
// 给定 nums = [2, 7, 11, 15], target = 9
//
//因为 nums[0] + nums[1] = 2 + 7 = 9
//所以返回 [0, 1]
//
// Related Topics 数组 哈希表
// 👍 9624 👎 0

//leetcode submit region begin(Prohibit modification and deletion)
func twoSum_I(nums []int, target int) []int {
	return twoSumC(nums, target)
}

// 总结方案B（twoSumB）
// 发现取反值的时候，可以不用全部遍历完，在进行比较。可以边遍历边取值
func twoSumC(nums []int, target int) []int {
	var data = make(map[int]int)
	for k, v := range nums {
		if index, has := data[target-v]; has {
			return []int{index, k}
		}

		data[v] = k
	}
	return nil
}

// 空间换时间
// 缺点： 内存使用大，速度也不快
func twoSumB(nums []int, target int) []int {
	type (
		node struct {
			num   int
			diff  int
			index []int
		}
	)

	var data = make(map[int]*node)
	for k, v := range nums {
		if _, has := data[v]; !has {
			data[v] = &node{
				num:   v,
				diff:  target - v,
				index: make([]int, 0),
			}
		}

		data[v].index = append(data[v].index, k)
	}

	for _, node := range data {
		var diffNode, has = data[node.diff]
		if !has {
			continue
		}

		if diffNode.num != node.num {
			return []int{node.index[0], diffNode.index[0]}
		}

		if len(node.index) >= 2 {
			return node.index[:2]
		}
	}

	return nil
}

// 遍历
// 缺点：慢，虽然可以通过单元测试，但是提交时提示超时
func twoSumA(nums []int, target int) []int {
	var i, j, l = 0, 0, len(nums)
	for i = 0; i < l-1; i++ {
		for j = i + 1; j < l; j++ {
			if nums[i]+nums[j] == target {
				return []int{i, j}
			}
		}
	}
	return nil
}

//leetcode submit region end(Prohibit modification and deletion)

func main() {
	fmt.Println(twoSum([]int{1, 3, 3, 4}, 6))
}
