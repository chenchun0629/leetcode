package main

func main() {

}

//给定一个 n × n 的二维矩阵 matrix 表示一个图像。请你将图像顺时针旋转 90 度。
//
// 你必须在 原地 旋转图像，这意味着你需要直接修改输入的二维矩阵。请不要 使用另一个矩阵来旋转图像。
//
//
//
// 示例 1：
//
//
//输入：matrix = [[1,2,3],[4,5,6],[7,8,9]]
//输出：[[7,4,1],[8,5,2],[9,6,3]]
//
//
// 示例 2：
//
//
//输入：matrix = [[5,1,9,11],[2,4,8,10],[13,3,6,7],[15,14,12,16]]
//输出：[[15,13,2,5],[14,3,4,1],[12,6,8,9],[16,7,10,11]]
//
//
// 示例 3：
//
//
//输入：matrix = [[1]]
//输出：[[1]]
//
//
// 示例 4：
//
//
//输入：matrix = [[1,2],[3,4]]
//输出：[[3,1],[4,2]]
//
//
//
//
// 提示：
//
//
// matrix.length == n
// matrix[i].length == n
// 1 <= n <= 20
// -1000 <= matrix[i][j] <= 1000
//
// Related Topics 数组 数学 矩阵 👍 1105 👎 0

//leetcode submit region begin(Prohibit modification and deletion)
func rotate(matrix [][]int) {
	var (
		n = len(matrix)
	)

	for i := 0; i < n/2; i++ {
		for j := 0; j < (n+1)/2; j++ {
			//matrix[i][j], matrix[n-i-1][j], matrix[n-i-1][n-j-1], matrix[i][n-j-1] =
			//	matrix[n-i-1][j], matrix[n-i-1][n-j-1], matrix[i][n-j-1], matrix[i][j]

			matrix[i][j], matrix[n-j-1][i], matrix[n-i-1][n-j-1], matrix[j][n-i-1] =
				matrix[n-j-1][i], matrix[n-i-1][n-j-1], matrix[j][n-i-1], matrix[i][j]
		}
	}
}

//leetcode submit region end(Prohibit modification and deletion)

func rotate_B(matrix [][]int) {
	var (
		h  = len(matrix)
		w  = len(matrix[0])
		nm = make([][]int, 0)
	)

	for i := 0; i < w; i++ {
		var row = make([]int, 0)
		for j := 0; j < h; j++ {
			row = append(row, matrix[h-j-1][i])
		}

		nm = append(nm, row)
	}

	copy(matrix, nm)
}

func rotate_A(matrix [][]int) {
	var (
		h  = len(matrix)
		w  = len(matrix[0])
		nm = make([][]int, 0)
	)

	for i := 0; i < h; i++ {
		var tmp = make([]int, w)
		copy(tmp, matrix[i])
		nm = append(nm, tmp)
	}

	for i := 0; i < w; i++ {
		//var row = make([]int, 0)
		for j := 0; j < h; j++ {
			matrix[i][j] = nm[h-j-1][i]
			//row = append(row, matrix[h-j-1][i])
		}

		//nm = append(nm, row)
	}

	//fmt.Printf("%+v \n", nm)
}
