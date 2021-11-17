package main

func main() {

}

//给定一个二叉树，找出其最小深度。
//
// 最小深度是从根节点到最近叶子节点的最短路径上的节点数量。
//
// 说明：叶子节点是指没有子节点的节点。
//
//
//
// 示例 1：
//
//
//输入：root = [3,9,20,null,null,15,7]
//输出：2
//
//
// 示例 2：
//
//
//输入：root = [2,null,3,null,4,null,5,null,6]
//输出：5
//
//
//
//
// 提示：
//
//
// 树中节点数的范围在 [0, 10⁵] 内
// -1000 <= Node.val <= 1000
//
// Related Topics 树 深度优先搜索 广度优先搜索 二叉树 👍 617 👎 0

//leetcode submit region begin(Prohibit modification and deletion)
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func minDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}

	if root.Left == nil && root.Right == nil {
		return 1
	}

	if root.Left == nil {
		return minDepth(root.Right) + 1
	}

	if root.Right == nil {
		return minDepth(root.Left) + 1
	}

	return minDepthMin(minDepth(root.Left), minDepth(root.Right)) + 1
}

//		 3,
//        9,20,
//null,null,15,7

func minDepthMin(a, b int) int {
	if a < b {
		return a
	}
	return b
}

//leetcode submit region end(Prohibit modification and deletion)

func minDepth_A(root *TreeNode) int {
	if root == nil {
		return 0
	}

	var q = make([]*TreeNode, 0)
	var lvl = 0

	q = append(q, root)

	for len(q) > 0 {
		var c = q[:]
		q = make([]*TreeNode, 0)
		for _, v := range c {
			if v == nil {
				continue
			}

			if v.Left == nil && v.Right == nil {
				return lvl + 1
			}
			q = append(q, v.Left, v.Right)
		}

		lvl++
	}

	return lvl
}
