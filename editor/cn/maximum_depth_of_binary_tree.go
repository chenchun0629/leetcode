package main

import "math"

func main() {

}

//给定一个二叉树，找出其最大深度。
//
// 二叉树的深度为根节点到最远叶子节点的最长路径上的节点数。
//
// 说明: 叶子节点是指没有子节点的节点。
//
// 示例：
//给定二叉树 [3,9,20,null,null,15,7]，
//
//     3
//   / \
//  9  20
//    /  \
//   15   7
//
// 返回它的最大深度 3 。
// Related Topics 树 深度优先搜索 广度优先搜索 二叉树 👍 1020 👎 0

//leetcode submit region begin(Prohibit modification and deletion)
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func maxDepth(root *TreeNode) int {
	var (
		fn func(n *TreeNode, level int) int
	)

	fn = func(n *TreeNode, level int) int {
		if n == nil {
			return level - 1
		}

		return int(math.Max(float64(fn(n.Left, level+1)), float64(fn(n.Right, level+1))))
	}

	return fn(root, 1)
}

//leetcode submit region end(Prohibit modification and deletion)

type LevelTreeNode struct {
	*TreeNode
	Level int
}

func maxDepth_A(root *TreeNode) int {
	var (
		q  = make([]*LevelTreeNode, 0)
		ml int
	)

	if root != nil {
		q = append(q, &LevelTreeNode{TreeNode: root, Level: 1})
	}

	for len(q) > 0 {
		var c = q[0]
		q = q[1:]

		if c.Level > ml {
			ml = c.Level
		}

		if c.Left != nil {
			q = append(q, &LevelTreeNode{TreeNode: c.Left, Level: c.Level + 1})
		}

		if c.Right != nil {
			q = append(q, &LevelTreeNode{TreeNode: c.Right, Level: c.Level + 1})
		}
	}

	return ml
}
