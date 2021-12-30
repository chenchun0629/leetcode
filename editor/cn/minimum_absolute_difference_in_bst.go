package main

import "math"

func main() {

}

//给你一个二叉搜索树的根节点 root ，返回 树中任意两不同节点值之间的最小差值 。
//
// 差值是一个正数，其数值等于两值之差的绝对值。
//
//
//
// 示例 1：
//
//
//输入：root = [4,2,6,1,3]
//输出：1
//
//
// 示例 2：
//
//
//输入：root = [1,0,48,null,null,12,49]
//输出：1
//
//
//
//
// 提示：
//
//
// 树中节点的数目范围是 [2, 10⁴]
// 0 <= Node.val <= 10⁵
//
//
//
//
// 注意：本题与 783 https://leetcode-cn.com/problems/minimum-distance-between-bst-
//nodes/ 相同
// Related Topics 树 深度优先搜索 广度优先搜索 二叉搜索树 二叉树 👍 292 👎 0

//leetcode submit region begin(Prohibit modification and deletion)
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func getMinimumDifference(root *TreeNode) int {
	var (
		minV = math.MaxInt64
		pre  = -1
		//abs = func(a int) int {
		//	if a < 0 {
		//		return -a
		//	}
		//	return a
		//}
		//min = func(a, b int) int {
		//	if a < b {
		//		return a
		//	}
		//	return b
		//}
		fn func(n *TreeNode)
	)

	fn = func(n *TreeNode) {
		if n == nil {
			return
		}

		//if n.Left != nil {
		//	minV = min(minV, n.Val-n.Left.Val)
		//	fn(n.Left)
		//}

		//     5
		//  3     7
		// 1 2   6  8

		fn(n.Left)

		if pre != -1 && n.Val-pre < minV {
			minV = n.Val - pre
		}
		pre = n.Val

		fn(n.Right)

		//if n.Right != nil {
		//	minV = min(minV, n.Right.Val-n.Val)
		//	fn(n.Right)
		//}
	}

	fn(root)
	//      236,
	//    104,701,
	//null,227,null,911

	return minV
}

//leetcode submit region end(Prohibit modification and deletion)
