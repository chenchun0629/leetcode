package main

func main() {

}

//给定两个二叉树，想象当你将它们中的一个覆盖到另一个上时，两个二叉树的一些节点便会重叠。
//
// 你需要将他们合并为一个新的二叉树。合并的规则是如果两个节点重叠，那么将他们的值相加作为节点合并后的新值，否则不为 NULL 的节点将直接作为新二叉树的节点
//。
//
// 示例 1:
//
//
//输入:
//	Tree 1                     Tree 2
//          1                         2
//         / \                       / \
//        3   2                     1   3
//       /                           \   \
//      5                             4   7
//输出:
//合并后的树:
//	     3
//	    / \
//	   4   5
//	  / \   \
//	 5   4   7
//
//
// 注意: 合并必须从两个树的根节点开始。
// Related Topics 树 深度优先搜索 广度优先搜索 二叉树 👍 846 👎 0

//leetcode submit region begin(Prohibit modification and deletion)
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func mergeTrees(root1 *TreeNode, root2 *TreeNode) *TreeNode {
	var (
		fn func(node1 *TreeNode, node2 *TreeNode) *TreeNode
	)

	fn = func(node1 *TreeNode, node2 *TreeNode) *TreeNode {
		if node1 == nil {
			return node2
		}

		if node2 == nil {
			return node1
		}

		node1.Val += node2.Val
		node1.Left = mergeTrees(node1.Left, node2.Left)
		node1.Right = mergeTrees(node1.Right, node2.Right)

		return node1
	}

	return fn(root1, root2)
}

//leetcode submit region end(Prohibit modification and deletion)

func mergeTrees_A(root1 *TreeNode, root2 *TreeNode) *TreeNode {
	var (
		fn func(node1 *TreeNode, node2 *TreeNode) *TreeNode
	)

	fn = func(node1 *TreeNode, node2 *TreeNode) *TreeNode {
		if node1 == nil && node2 == nil {
			return nil
		}

		var (
			root     = &TreeNode{}
			ln1, ln2 *TreeNode
			rn1, rn2 *TreeNode
		)

		if node1 != nil {
			ln1 = node1.Left
			rn1 = node1.Right
			root.Val += node1.Val
		}

		if node2 != nil {
			ln2 = node2.Left
			rn2 = node2.Right
			root.Val += node2.Val
		}

		root.Left = fn(ln1, ln2)
		root.Right = fn(rn1, rn2)

		return root
	}

	return fn(root1, root2)
}
