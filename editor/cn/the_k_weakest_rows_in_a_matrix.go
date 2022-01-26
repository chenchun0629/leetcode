package main

import (
	"container/heap"
	"fmt"
)

func main() {
	fmt.Println(kWeakestRows([][]int{
		{1, 1, 0, 0, 0},
		{1, 1, 1, 1, 0},
		{1, 0, 0, 0, 0},
		{1, 1, 0, 0, 0},
		{1, 1, 1, 1, 1},
	}, 3))
}

//给你一个大小为 m * n 的矩阵 mat，矩阵由若干军人和平民组成，分别用 1 和 0 表示。
//
// 请你返回矩阵中战斗力最弱的 k 行的索引，按从最弱到最强排序。
//
// 如果第 i 行的军人数量少于第 j 行，或者两行军人数量相同但 i 小于 j，那么我们认为第 i 行的战斗力比第 j 行弱。
//
// 军人 总是 排在一行中的靠前位置，也就是说 1 总是出现在 0 之前。
//
//
//
// 示例 1：
//
//
//输入：mat =
//[[1,1,0,0,0],
// [1,1,1,1,0],
// [1,0,0,0,0],
// [1,1,0,0,0],
// [1,1,1,1,1]],
//k = 3
//输出：[2,0,3]
//解释：
//每行中的军人数目：
//行 0 -> 2
//行 1 -> 4
//行 2 -> 1
//行 3 -> 2
//行 4 -> 5
//从最弱到最强对这些行排序后得到 [2,0,3,1,4]
//
//
// 示例 2：
//
//
//输入：mat =
//[[1,0,0,0],
// [1,1,1,1],
// [1,0,0,0],
// [1,0,0,0]],
//k = 2
//输出：[0,2]
//解释：
//每行中的军人数目：
//行 0 -> 1
//行 1 -> 4
//行 2 -> 1
//行 3 -> 1
//从最弱到最强对这些行排序后得到 [0,2,3,1]
//
//
//
//
// 提示：
//
//
// m == mat.length
// n == mat[i].length
// 2 <= n, m <= 100
// 1 <= k <= m
// matrix[i][j] 不是 0 就是 1
//
// Related Topics 数组 二分查找 矩阵 排序 堆（优先队列） 👍 134 👎 0

//leetcode submit region begin(Prohibit modification and deletion)

type pair struct{ pow, idx int }
type hp []pair

func (h hp) Len() int { return len(h) }
func (h hp) Less(i, j int) bool {
	a, b := h[i], h[j]
	return a.pow < b.pow || a.pow == b.pow && a.idx < b.idx
}
func (h hp) Swap(i, j int)       { h[i], h[j] = h[j], h[i] }
func (h *hp) Push(v interface{}) { *h = append(*h, v.(pair)) }
func (h *hp) Pop() interface{}   { a := *h; v := a[len(a)-1]; *h = a[:len(a)-1]; return v }

func kWeakestRows(mat [][]int, k int) []int {
	var (
		h = hp{}
	)

	for k, row := range mat {
		var head, tail = 0, len(row)
		for head < tail {
			var mid = head + (tail-head)>>1
			if row[mid] == 1 {
				head = mid + 1
			} else {
				tail = mid
			}
		}
		h = append(h, pair{head, k})
	}

	heap.Init(&h)
	var ret = make([]int, k)
	for i := range ret {
		ret[i] = heap.Pop(&h).(pair).idx
	}
	return ret
}

//leetcode submit region end(Prohibit modification and deletion)
