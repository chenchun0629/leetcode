package main

import "fmt"

func main() {
	fmt.Printf("%+v, \n", generateMatrix(3))
}

//给你一个正整数 n ，生成一个包含 1 到 n² 所有元素，且元素按顺时针顺序螺旋排列的 n x n 正方形矩阵 matrix 。
//
//
//
// 示例 1：
//
//
//输入：n = 3
//输出：[[1,2,3],[8,9,4],[7,6,5]]
//
//
// 示例 2：
//
//
//输入：n = 1
//输出：[[1]]
//
//
//
//
// 提示：
//
//
// 1 <= n <= 20
//
// Related Topics 数组 矩阵 模拟 👍 564 👎 0

//leetcode submit region begin(Prohibit modification and deletion)
func generateMatrix(n int) [][]int {
	var (
		matrix                   = [][]int{}
		direct                   int
		left, right, top, buttom = 0, n - 1, 0, n - 1
		row, col                 = 0, 0

		nn = n * n
	)

	for i := 0; i < n; i++ {
		matrix = append(matrix, make([]int, n))
	}

	for i := 0; i < nn; i++ {
		switch direct {
		case 0:
			matrix[row][col] = i + 1
			if col == right {
				direct = 1
				top++
				row++
			} else {
				col++
			}
		case 1:
			matrix[row][col] = i + 1
			if row == buttom {
				direct = 2
				right--
				col--
			} else {
				row++
			}
		case 2:
			matrix[row][col] = i + 1
			if col == left {
				direct = 3
				buttom--
				row--
			} else {
				col--
			}
		case 3:
			matrix[row][col] = i + 1
			if row == top {
				direct = 0
				left++
				col++
			} else {
				row--
			}
		}
	}

	return matrix
}

//leetcode submit region end(Prohibit modification and deletion)
