package main

import "math"

func main() {

}

//给定一个包含非负整数的 m x n 网格 grid ，请找出一条从左上角到右下角的路径，使得路径上的数字总和为最小。
//
// 说明：每次只能向下或者向右移动一步。
//
//
//
// 示例 1：
//
//
//输入：grid = [[1,3,1],[1,5,1],[4,2,1]]
//输出：7
//解释：因为路径 1→3→1→1→1 的总和最小。
//
//
// 示例 2：
//
//
//输入：grid = [[1,2,3],[4,5,6]]
//输出：12
//
//
//
//
// 提示：
//
//
// m == grid.length
// n == grid[i].length
// 1 <= m, n <= 200
// 0 <= grid[i][j] <= 100
//
// Related Topics 数组 动态规划 矩阵 👍 1118 👎 0

//leetcode submit region begin(Prohibit modification and deletion)
func minPathSum(grid [][]int) int {
	// 迭代
	var (
		m, n = len(grid), len(grid[0])
		min  = func(a, b int) int {
			if a > b {
				return b
			}
			return a
		}
		backlog = make([][]int, 0)
	)

	for i := 0; i < m; i++ {
		backlog = append(backlog, make([]int, n))
	}

	backlog[0][0] = grid[0][0]

	for i := 1; i < m; i++ {
		backlog[i][0] = backlog[i-1][0] + grid[i][0]
	}
	for i := 1; i < n; i++ {
		backlog[0][i] = backlog[0][i-1] + grid[0][i]
	}

	for i := 1; i < m; i++ {
		for j := 1; j < n; j++ {
			backlog[i][j] = min(backlog[i-1][j], backlog[i][j-1]) + grid[i][j]
		}
	}

	return backlog[m-1][n-1]
}

//leetcode submit region end(Prohibit modification and deletion)

// 递归
func minPathSum_A(grid [][]int) int {
	var (
		m, n = len(grid), len(grid[0])
		fn   func(x, y int) int
		min  = func(a, b int) int {
			if a > b {
				return b
			}
			return a
		}
		backlog = make(map[[2]int]int)
	)

	fn = func(x, y int) int {
		if x+1 == m && y+1 == n {
			return grid[x][y]
		}

		var key = [2]int{x, y}
		if _, has := backlog[key]; !has {
			var minV = math.MaxInt64

			if x+1 < m {
				minV = min(fn(x+1, y), minV)
			}
			if y+1 < n {
				minV = min(fn(x, y+1), minV)
			}

			backlog[key] = minV + grid[x][y]
		}

		return backlog[key]
	}

	return fn(0, 0)
}
