package main

import "fmt"

func main() {
	fmt.Println(diameterOfBinaryTree(&TreeNode{Val: 1, Right: &TreeNode{Val: 2}}))
}

//给定一棵二叉树，你需要计算它的直径长度。一棵二叉树的直径长度是任意两个结点路径长度中的最大值。这条路径可能穿过也可能不穿过根结点。
//
//
//
// 示例 :
//给定二叉树
//
//           1
//         / \
//        2   3
//       / \
//      4   5
//
//
// 返回 3, 它的长度是路径 [4,2,1,3] 或者 [5,2,1,3]。
//
//
//
// 注意：两结点之间的路径长度是以它们之间边的数目表示。
// Related Topics 树 深度优先搜索 二叉树 👍 864 👎 0

//type TreeNode struct {
//	Val   int
//	Left  *TreeNode
//	Right *TreeNode
//}

//leetcode submit region begin(Prohibit modification and deletion)
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func diameterOfBinaryTree(root *TreeNode) int {
	var (
		max = func(a, b int) int {
			if a > b {
				return a
			}
			return b
		}

		fn  func(node *TreeNode) int
		ret int
	)

	fn = func(node *TreeNode) int {
		if node == nil {
			return 0
		}

		var (
			lv = fn(node.Left)
			rv = fn(node.Right)
		)

		ret = max(ret, lv+rv)

		return max(lv, rv) + 1
	}

	fn(root)

	return ret
}

//leetcode submit region end(Prohibit modification and deletion)
