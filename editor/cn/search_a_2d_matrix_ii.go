package main

import "fmt"

func main() {
	fmt.Println(searchMatrix([][]int{
		{1, 4, 7, 11, 15},
		{2, 5, 8, 12, 19},
		{3, 6, 9, 16, 22},
		{10, 13, 14, 17, 24},
		{18, 21, 23, 26, 30},
	}, 20))
}

//编写一个高效的算法来搜索 m x n 矩阵 matrix 中的一个目标值 target 。该矩阵具有以下特性：
//
//
// 每行的元素从左到右升序排列。
// 每列的元素从上到下升序排列。
//
//
//
//
// 示例 1：
//
//
//输入：matrix = [[1,4,7,11,15],[2,5,8,12,19],[3,6,9,16,22],[10,13,14,17,24],[18,21
//,23,26,30]], target = 5
//输出：true
//
//
// 示例 2：
//
//
//输入：matrix = [[1,4,7,11,15],[2,5,8,12,19],[3,6,9,16,22],[10,13,14,17,24],[18,21
//,23,26,30]], target = 20
//输出：false
//
//
//
//
// 提示：
//
//
// m == matrix.length
// n == matrix[i].length
// 1 <= n, m <= 300
// -10⁹ <= matrix[i][j] <= 10⁹
// 每行的所有元素从左到右升序排列
// 每列的所有元素从上到下升序排列
// -10⁹ <= target <= 10⁹
//
// Related Topics 数组 二分查找 分治 矩阵 👍 908 👎 0

//leetcode submit region begin(Prohibit modification and deletion)
// z字形
func searchMatrix(matrix [][]int, target int) bool {
	var m, n = len(matrix), len(matrix[0])
	var x, y = 0, n - 1

	for y >= 0 && x < m {
		if matrix[x][y] == target {
			return true
		}

		if matrix[x][y] < target {
			x++
		} else {
			y--
		}
	}

	return false
}

//leetcode submit region end(Prohibit modification and deletion)

// 按行二分查找
func searchMatrix_A(matrix [][]int, target int) bool {
	for _, row := range matrix {
		var left, right = 0, len(row) - 1
		for left <= right {
			var mid = (left + right) / 2
			if row[mid] == target {
				return true
			}

			if row[mid] > target {
				right = mid - 1
			} else {
				left = mid + 1
			}
		}
	}

	return false
}
