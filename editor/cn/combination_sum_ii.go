package main

import (
	"fmt"
	"sort"
)

func main() {
	//fmt.Printf("%+v \n", combinationSum2([]int{10, 1, 2, 7, 6, 1, 5}, 8))
	fmt.Printf("%+v \n", combinationSum2([]int{1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1}, 30))
}

//给定一个数组 candidates 和一个目标数 target ，找出 candidates 中所有可以使数字和为 target 的组合。
//
// candidates 中的每个数字在每个组合中只能使用一次。
//
// 注意：解集不能包含重复的组合。
//
//
//
// 示例 1:
//
//
//输入: candidates = [10,1,2,7,6,1,5], target = 8,
//输出:
//[
//[1,1,6],
//[1,2,5],
//[1,7],
//[2,6]
//]
//
// 示例 2:
//
//
//输入: candidates = [2,5,2,1,2], target = 5,
//输出:
//[
//[1,2,2],
//[5]
//]
//
//
//
// 提示:
//
//
// 1 <= candidates.length <= 100
// 1 <= candidates[i] <= 50
// 1 <= target <= 30
//
// Related Topics 数组 回溯 👍 780 👎 0

//leetcode submit region begin(Prohibit modification and deletion)

// 超时	fmt.Printf("%+v \n", combinationSum2([]int{1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1}, 30))
func combinationSum2(candidates []int, target int) [][]int {
	var (
		n   = len(candidates)
		ret [][]int
		fn  func(prefix []int, idx int, target int)
		m   = make(map[string]struct{})
	)

	fn = func(prefix []int, idx int, target int) {
		for i := idx; i < n; i++ {
			if candidates[i] > target {
				return
			}

			//if i > 0 && candidates[i] == candidates[i-1] {
			//	continue
			//}

			if target == candidates[i] {
				var rt = append(prefix, candidates[i])
				var key = fmt.Sprintf("%+v", rt)
				if _, has := m[key]; !has {
					ret = append(ret, rt)
					m[key] = struct{}{}
				}

				//ret = append(ret, append(prefix, candidates[i]))
			}

			if target-candidates[i] > 0 {
				fn(append(prefix, candidates[i]), i+1, target-candidates[i])
			}
		}
	}

	sort.Ints(candidates)
	fn([]int{}, 0, target)

	return ret
}

//leetcode submit region end(Prohibit modification and deletion)
