package main

import "fmt"

func main() {
	fmt.Printf("%+v \n", spiralOrder([][]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}}))
	fmt.Printf("%+v \n", spiralOrder([][]int{{1, 2, 3, 4}, {5, 6, 7, 8}, {9, 10, 11, 12}}))
	fmt.Printf("%+v \n", spiralOrder([][]int{{3}, {2}}))
}

//给你一个 m 行 n 列的矩阵 matrix ，请按照 顺时针螺旋顺序 ，返回矩阵中的所有元素。
//
//
//
// 示例 1：
//
//
//输入：matrix = [[1,2,3],[4,5,6],[7,8,9]]
//输出：[1,2,3,6,9,8,7,4,5]
//
//
// 示例 2：
//
//
//输入：matrix = [[1,2,3,4],[5,6,7,8],[9,10,11,12]]
//输出：[1,2,3,4,8,12,11,10,9,5,6,7]
//
//
//
//
// 提示：
//
//
// m == matrix.length
// n == matrix[i].length
// 1 <= m, n <= 10
// -100 <= matrix[i][j] <= 100
//
// Related Topics 数组 矩阵 模拟 👍 947 👎 0

//leetcode submit region begin(Prohibit modification and deletion)
func spiralOrder(matrix [][]int) []int {
	if len(matrix) == 0 || len(matrix[0]) == 0 {
		return []int{}
	}

	var (
		m                        = len(matrix)
		n                        = len(matrix[0])
		left, right, top, bottom = 0, n - 1, 0, m - 1
		row, col                 = left, top
		ret                      []int
	)

	for top <= bottom && left <= right {
		for col, row = left, top; col <= right; col++ {
			ret = append(ret, matrix[row][col])
		}

		for col, row = right, top+1; row <= bottom; row++ {
			ret = append(ret, matrix[row][col])
		}

		if left < right && top < bottom {
			for col, row = right-1, bottom; col >= left; col-- {
				ret = append(ret, matrix[row][col])
			}

			for col, row = left, bottom-1; row > top; row-- {
				ret = append(ret, matrix[row][col])
			}
		}

		top++
		right--
		bottom--
		left++
	}

	return ret
}

//leetcode submit region end(Prohibit modification and deletion)

func spiralOrder_A(matrix [][]int) []int {
	if len(matrix) == 0 || len(matrix[0]) == 0 {
		return []int{}
	}

	var (
		m                        = len(matrix)
		n                        = len(matrix[0])
		left, right, top, bottom = 0, n - 1, 0, m - 1
		//row, col                 = left, top
		ret []int
	)

	for top <= bottom && left <= right {
		for col := left; col <= right; col++ {
			ret = append(ret, matrix[top][col])
		}

		for row := top + 1; row <= bottom; row++ {
			ret = append(ret, matrix[row][right])
		}

		if left < right && top < bottom {
			for col := right - 1; col > left; col-- {
				ret = append(ret, matrix[bottom][col])
			}

			for row := bottom; row > top; row-- {
				ret = append(ret, matrix[row][left])
			}
		}

		top++
		right--
		bottom--
		left++

	}

	return ret
}
