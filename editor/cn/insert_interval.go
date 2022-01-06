package main

import "fmt"

func main() {
	//fmt.Printf("%+v \n", insert([][]int{{1, 2}, {3, 5}, {6, 7}, {8, 10}, {12, 16}}, []int{4, 8}))
	//fmt.Printf("%+v \n", insert([][]int{{1, 3}, {6, 9}}, []int{2, 5}))
	fmt.Printf("%+v \n", insert([][]int{{1, 5}}, []int{2, 7}))
}

//给你一个 无重叠的 ，按照区间起始端点排序的区间列表。
//
// 在列表中插入一个新的区间，你需要确保列表中的区间仍然有序且不重叠（如果有必要的话，可以合并区间）。
//
//
//
// 示例 1：
//
//
//输入：intervals = [[1,3],[6,9]], newInterval = [2,5]
//输出：[[1,5],[6,9]]
//
//
// 示例 2：
//
//
//输入：intervals = [[1,2],[3,5],[6,7],[8,10],[12,16]], newInterval = [4,8]
//输出：[[1,2],[3,10],[12,16]]
//解释：这是因为新的区间 [4,8] 与 [3,5],[6,7],[8,10] 重叠。
//
// 示例 3：
//
//
//输入：intervals = [], newInterval = [5,7]
//输出：[[5,7]]
//
//
// 示例 4：
//
//
//输入：intervals = [[1,5]], newInterval = [2,3]
//输出：[[1,5]]
//
//
// 示例 5：
//
//
//输入：intervals = [[1,5]], newInterval = [2,7]
//输出：[[1,7]]
//
//
//
//
// 提示：
//
//
// 0 <= intervals.length <= 10⁴
// intervals[i].length == 2
// 0 <= intervals[i][0] <= intervals[i][1] <= 10⁵
// intervals 根据 intervals[i][0] 按 升序 排列
// newInterval.length == 2
// 0 <= newInterval[0] <= newInterval[1] <= 10⁵
//
// Related Topics 数组 👍 530 👎 0

//leetcode submit region begin(Prohibit modification and deletion)
func insert(intervals [][]int, newInterval []int) [][]int {
	var (
		ret      [][]int
		n        = len(intervals)
		inserted bool
	)

	for i := 0; i < n; i++ {
		if inserted {
			ret = append(ret, intervals[i])
			continue
		}

		if intervals[i][0] > newInterval[1] {
			ret = append(ret, newInterval, intervals[i])
			inserted = true
			continue
		}

		// n o
		if newInterval[0] <= intervals[i][0] {
			// n o n o => n o
			if newInterval[1] <= intervals[i][1] {
				ret = append(ret, []int{newInterval[0], intervals[i][1]})
				inserted = true
				continue
			}

			// n o o n => n o1 o1 ... ox n ox
			// n o o n => n o1 o1 ... ox ox n
			if intervals[i][1] < newInterval[1] {
				for i = i + 1; i < n; i++ {
					// n o o n => n o1 o1 ... ox ox n
					if newInterval[1] < intervals[i][0] {
						ret = append(ret, newInterval, intervals[i])
						inserted = true
						break
					}

					// n o o n => n o1 o1 ... ox n ox
					if newInterval[1] < intervals[i][1] {
						ret = append(ret, []int{newInterval[0], intervals[i][1]})
						inserted = true
						break
					}
				}

				if !inserted {
					ret = append(ret, newInterval)
					inserted = true
				}

				continue
			}

		}

		// o n
		if newInterval[0] >= intervals[i][0] && newInterval[0] <= intervals[i][1] {
			// o n n o
			if newInterval[1] <= intervals[i][1] {
				ret = append(ret, intervals[i])
				inserted = true
				continue
			}

			// o n o n => o n o ... ox n ox => o ox
			// o n o n => o n o ... ox  ox n => o n
			var beginO = intervals[i][0]
			if newInterval[1] > intervals[i][1] {
				for i = i + 1; i < n; i++ {
					// o n o n => o n o ... ox  ox n => o n
					if newInterval[1] < intervals[i][0] {
						ret = append(ret, []int{beginO, newInterval[1]}, intervals[i])
						inserted = true
						break
					}

					// o n o n => o n o ... ox n ox => o ox
					if newInterval[1] <= intervals[i][1] {
						ret = append(ret, []int{beginO, intervals[i][1]})
						inserted = true
						break
					}
				}

				if !inserted {
					ret = append(ret, []int{beginO, newInterval[1]})
					inserted = true
				}

				continue
			}
		}

		ret = append(ret, intervals[i])
	}

	if !inserted {
		ret = append(ret, newInterval)
		inserted = true
	}

	return ret
}

//leetcode submit region end(Prohibit modification and deletion)
