package main

import (
	"fmt"
)

func main() {
	fmt.Println(uniquePaths(3, 7))
}

//一个机器人位于一个 m x n 网格的左上角 （起始点在下图中标记为 “Start” ）。
//
// 机器人每次只能向下或者向右移动一步。机器人试图达到网格的右下角（在下图中标记为 “Finish” ）。
//
// 问总共有多少条不同的路径？
//
//
//
// 示例 1：
//
//
//输入：m = 3, n = 7
//输出：28
//
// 示例 2：
//
//
//输入：m = 3, n = 2
//输出：3
//解释：
//从左上角开始，总共有 3 条路径可以到达右下角。
//1. 向右 -> 向下 -> 向下
//2. 向下 -> 向下 -> 向右
//3. 向下 -> 向右 -> 向下
//
//
// 示例 3：
//
//
//输入：m = 7, n = 3
//输出：28
//
//
// 示例 4：
//
//
//输入：m = 3, n = 3
//输出：6
//
//
//
// 提示：
//
//
// 1 <= m, n <= 100
// 题目数据保证答案小于等于 2 * 10⁹
//
// Related Topics 数学 动态规划 组合数学 👍 1244 👎 0

//leetcode submit region begin(Prohibit modification and deletion)
func uniquePaths(m int, n int) int {
	// m = 3, n = 7
	// 从左上角到右下角的过程中，我们需要移动 m+n-2 次，其中有 m-1 次向下移动，n-1 次向右移动
	// 路径的总数，就等于从 m+n-2 次移动中选择 m-1 次向下移动的方案数
	var ret = 1
	for x, y := n, 1; y < m; x, y = x+1, y+1 {
		ret = ret * x / y
	}
	return ret
	//return int(new(big.Int).Binomial(int64(m+n-2), int64(n-1)).Int64())
}

//leetcode submit region end(Prohibit modification and deletion)

func uniquePaths_A(m int, n int) int {
	var (
		fn      func(x, y int) int
		backlog = make(map[string]int)
	)

	fn = func(x, y int) int {
		if x+1 == m && y+1 == n {
			return 1
		}

		var ret int

		if x+1 < m {
			var key = fmt.Sprintf("%d_%d", x+1, y)
			if _, has := backlog[key]; !has {
				backlog[key] = fn(x+1, y)
			}
			ret += backlog[key]
		}

		if y+1 < n {
			var key = fmt.Sprintf("%d_%d", x, y+1)
			if _, has := backlog[key]; !has {
				backlog[key] = fn(x, y+1)
			}
			ret += backlog[key]
		}

		return ret
	}

	return fn(0, 0)
}
