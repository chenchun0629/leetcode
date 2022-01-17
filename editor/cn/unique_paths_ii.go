package main

import "fmt"

func main() {
	fmt.Println(uniquePathsWithObstacles([][]int{
		{0, 0, 0, 0},
		{0, 1, 0, 0},
		{0, 0, 0, 0},
		{0, 0, 1, 0},
		{0, 0, 0, 0},
	}))
}

//一个机器人位于一个 m x n 网格的左上角 （起始点在下图中标记为“Start” ）。
//
// 机器人每次只能向下或者向右移动一步。机器人试图达到网格的右下角（在下图中标记为“Finish”）。
//
// 现在考虑网格中有障碍物。那么从左上角到右下角将会有多少条不同的路径？
//
//
//
// 网格中的障碍物和空位置分别用 1 和 0 来表示。
//
//
//
// 示例 1：
//
//
//输入：obstacleGrid = [[0,0,0],[0,1,0],[0,0,0]]
//输出：2
//解释：
//3x3 网格的正中间有一个障碍物。
//从左上角到右下角一共有 2 条不同的路径：
//1. 向右 -> 向右 -> 向下 -> 向下
//2. 向下 -> 向下 -> 向右 -> 向右
//
//
// 示例 2：
//
//
//输入：obstacleGrid = [[0,1],[0,0]]
//输出：1
//
//
//
//
// 提示：
//
//
// m == obstacleGrid.length
// n == obstacleGrid[i].length
// 1 <= m, n <= 100
// obstacleGrid[i][j] 为 0 或 1
//
// Related Topics 数组 动态规划 矩阵 👍 704 👎 0

//leetcode submit region begin(Prohibit modification and deletion)
func uniquePathsWithObstacles(obstacleGrid [][]int) int {
	var (
		backlog = make(map[string]int)
		m, n    = len(obstacleGrid), len(obstacleGrid[0])
		fn      func(x, y int) int
	)

	fn = func(x, y int) int {
		if obstacleGrid[x][y] == 1 {
			return 0
		}

		if x+1 == m && y+1 == n {
			return 1
		}

		var cnt int
		if x+1 < m && obstacleGrid[x+1][y] != 1 {
			var key = fmt.Sprintf("%d_%d", x+1, y)
			if _, has := backlog[key]; !has {
				backlog[key] = fn(x+1, y)
			}
			cnt += backlog[key]
		}

		if y+1 < n && obstacleGrid[x][y+1] != 1 {
			var key = fmt.Sprintf("%d_%d", x, y+1)
			if _, has := backlog[key]; !has {
				backlog[key] = fn(x, y+1)
			}
			cnt += backlog[key]
		}

		return cnt
	}

	return fn(0, 0)
}

//leetcode submit region end(Prohibit modification and deletion)
