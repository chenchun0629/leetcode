package main

func main() {

}

//计算给定二叉树的所有左叶子之和。
//
// 示例：
//
//
//    3
//   / \
//  9  20
//    /  \
//   15   7
//
//在这个二叉树中，有两个左叶子，分别是 9 和 15，所以返回 24
//
//
// Related Topics 树 深度优先搜索 广度优先搜索 二叉树 👍 370 👎 0

//leetcode submit region begin(Prohibit modification and deletion)
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func sumOfLeftLeaves(root *TreeNode) int {
	var s int

	if root == nil {
		return 0
	}

	if root.Left == nil && root.Right == nil {
		return 0
	}

	if root.Left != nil && root.Left.Left == nil && root.Left.Right == nil {
		s += root.Left.Val
	}

	if root.Left != nil {
		s += sumOfLeftLeaves(root.Left)
	}

	if root.Right != nil {
		s += sumOfLeftLeaves(root.Right)
	}

	return s
}

//leetcode submit region end(Prohibit modification and deletion)

func sumOfLeftLeaves_isLeafNode(node *TreeNode) bool {
	return node.Left == nil && node.Right == nil
}

func sumOfLeftLeaves_dfs(node *TreeNode) (ans int) {
	if node.Left != nil {
		if sumOfLeftLeaves_isLeafNode(node.Left) {
			ans += node.Left.Val
		} else {
			ans += sumOfLeftLeaves_dfs(node.Left)
		}
	}
	if node.Right != nil && !sumOfLeftLeaves_isLeafNode(node.Right) {
		ans += sumOfLeftLeaves_dfs(node.Right)
	}
	return
}

func sumOfLeftLeaves_B(root *TreeNode) int {
	if root == nil {
		return 0
	}
	return sumOfLeftLeaves_dfs(root)
}
