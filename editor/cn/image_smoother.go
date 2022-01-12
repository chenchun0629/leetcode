package main

func main() {

}

//包含整数的二维矩阵 M 表示一个图片的灰度。你需要设计一个平滑器来让每一个单元的灰度成为平均灰度 (向下舍入) ，平均灰度的计算是周围的8个单元和它本身的值
//求平均，如果周围的单元格不足八个，则尽可能多的利用它们。
//
// 示例 1:
//
//
//输入:
//[[1,1,1],
// [1,0,1],
// [1,1,1]]
//输出:
//[[0, 0, 0],
// [0, 0, 0],
// [0, 0, 0]]
//解释:
//对于点 (0,0), (0,2), (2,0), (2,2): 平均(3/4) = 平均(0.75) = 0
//对于点 (0,1), (1,0), (1,2), (2,1): 平均(5/6) = 平均(0.83333333) = 0
//对于点 (1,1): 平均(8/9) = 平均(0.88888889) = 0
//
//
// 注意:
//
//
// 给定矩阵中的整数范围为 [0, 255]。
// 矩阵的长和宽的范围均为 [1, 150]。
//
// Related Topics 数组 矩阵 👍 92 👎 0

//leetcode submit region begin(Prohibit modification and deletion)
func imageSmoother(img [][]int) [][]int {
	var (
		m, n = len(img), len(img[0])
		res  = make([][]int, m)
	)
	for i := range res {
		res[i] = make([]int, n)
	}

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			var count = 0
			for x := i - 1; x <= i+1; x++ {
				for y := j - 1; y <= j+1; y++ {
					if 0 <= x && x < m && 0 <= y && y < n {
						res[i][j] += img[x][y]
						count++
					}
				}
			}
			res[i][j] /= count
		}
	}
	return res
}

//leetcode submit region end(Prohibit modification and deletion)
