package main

func main() {

}

//给定一个非负整数 numRows，生成「杨辉三角」的前 numRows 行。
//
// 在「杨辉三角」中，每个数是它左上方和右上方的数的和。
//
//
//
//
//
// 示例 1:
//
//
//输入: numRows = 5
//输出: [[1],[1,1],[1,2,1],[1,3,3,1],[1,4,6,4,1]]
//
//
// 示例 2:
//
//
//输入: numRows = 1
//输出: [[1]]
//
//
//
//
// 提示:
//
//
// 1 <= numRows <= 30
//
// Related Topics 数组 动态规划 👍 628 👎 0

//leetcode submit region begin(Prohibit modification and deletion)
func generate(numRows int) [][]int {
	var triangle = make([][]int, 0)
	for i := 0; i < numRows; i++ {
		if i == 0 {
			triangle = append(triangle, []int{1})
		} else if i == 1 {
			triangle = append(triangle, []int{1, 1})
		} else {
			var line = make([]int, 0)
			for j := 0; j <= i; j++ {
				if j == 0 || j == i {
					line = append(line, 1)
				} else {
					line = append(line, triangle[i-1][j]+triangle[i-1][j-1])
				}
			}
			triangle = append(triangle, line)
		}
	}

	return triangle
}

//leetcode submit region end(Prohibit modification and deletion)
