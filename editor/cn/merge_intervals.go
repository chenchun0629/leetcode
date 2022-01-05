package main

import "sort"

func main() {
}

//以数组 intervals 表示若干个区间的集合，其中单个区间为 intervals[i] = [starti, endi] 。请你合并所有重叠的区间，并返
//回一个不重叠的区间数组，该数组需恰好覆盖输入中的所有区间。
//
//
//
// 示例 1：
//
//
//输入：intervals = [[1,3],[2,6],[8,10],[15,18]]
//输出：[[1,6],[8,10],[15,18]]
//解释：区间 [1,3] 和 [2,6] 重叠, 将它们合并为 [1,6].
//
//
// 示例 2：
//
//
//输入：intervals = [[1,4],[4,5]]
//输出：[[1,5]]
//解释：区间 [1,4] 和 [4,5] 可被视为重叠区间。
//
//
//
// 提示：
//
//
// 1 <= intervals.length <= 10⁴
// intervals[i].length == 2
// 0 <= starti <= endi <= 10⁴
//
// Related Topics 数组 排序 👍 1253 👎 0

//leetcode submit region begin(Prohibit modification and deletion)
func merge1(intervals [][]int) [][]int {
	sort.Slice(intervals, func(i, j int) bool {
		//if intervals[i][0] == intervals[j][0] {
		//	return intervals[i][0] > intervals[j][0]
		//}

		return intervals[i][0] < intervals[j][0]
	})

	var (
		n   = len(intervals)
		ret = [][]int{intervals[0]}
		//idx              int
		//prevMax = intervals[0][1]
	)

	for i := 1; i < n; i++ {
		if intervals[i][0] <= ret[len(ret)-1][1] {
			// 更新
			if intervals[i][1] > ret[len(ret)-1][1] {
				//prevMax = intervals[i][1]
				ret[len(ret)-1][1] = intervals[i][1]
			}
		} else {
			// 新增
			ret = append(ret, intervals[i])
			//prevMax = intervals[i][1]
			//idx++
		}
	}

	return ret
}

//leetcode submit region end(Prohibit modification and deletion)
