package main

func main() {

}

//给定一个非空二叉树, 返回一个由每层节点平均值组成的数组。
//
//
//
// 示例 1：
//
// 输入：
//    3
//   / \
//  9  20
//    /  \
//   15   7
//输出：[3, 14.5, 11]
//解释：
//第 0 层的平均值是 3 ,  第1层是 14.5 , 第2层是 11 。因此返回 [3, 14.5, 11] 。
//
//
//
//
// 提示：
//
//
// 节点值的范围在32位有符号整数范围内。
//
// Related Topics 树 深度优先搜索 广度优先搜索 二叉树 👍 305 👎 0

//leetcode submit region begin(Prohibit modification and deletion)
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func averageOfLevels(root *TreeNode) []float64 {
	var (
		ret []float64
		q   = []*TreeNode{root}
	)

	for len(q) > 0 {
		var (
			tmp []*TreeNode
			n   = len(q)
			sum int
		)

		for _, v := range q {
			sum += v.Val
			if v.Left != nil {
				tmp = append(tmp, v.Left)
			}
			if v.Right != nil {
				tmp = append(tmp, v.Right)
			}
		}

		ret = append(ret, float64(sum)/float64(n))
		q = tmp
	}

	return ret
}

//leetcode submit region end(Prohibit modification and deletion)
