package main

func main() {

}

//
//
// 给你两棵二叉树 root 和 subRoot 。检验 root 中是否包含和 subRoot 具有相同结构和节点值的子树。如果存在，返回 true ；否则
//，返回 false 。
//
// 二叉树 tree 的一棵子树包括 tree 的某个节点和这个节点的所有后代节点。tree 也可以看做它自身的一棵子树。
//
//
//
// 示例 1：
//
//
//输入：root = [3,4,5,1,2], subRoot = [4,1,2]
//输出：true
//
//
// 示例 2：
//
//
//输入：root = [3,4,5,1,2,null,null,null,null,0], subRoot = [4,1,2]
//输出：false
//
//
//
//
// 提示：
//
//
// root 树上的节点数量范围是 [1, 2000]
// subRoot 树上的节点数量范围是 [1, 1000]
// -10⁴ <= root.val <= 10⁴
// -10⁴ <= subRoot.val <= 10⁴
//
//
//
// Related Topics 树 深度优先搜索 二叉树 字符串匹配 哈希函数 👍 618 👎 0

//leetcode submit region begin(Prohibit modification and deletion)
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func isSubtree(root *TreeNode, subRoot *TreeNode) bool {
	var (
		fn   func(node *TreeNode) bool
		same func(a, b *TreeNode) bool
	)

	same = func(a, b *TreeNode) bool {
		if a == nil && b == nil {
			return true
		}

		if a == nil || b == nil {
			return false
		}

		if a.Val != b.Val {
			return false
		}

		if !same(a.Left, b.Left) || !same(a.Right, b.Right) {
			return false
		}

		return true
	}

	fn = func(node *TreeNode) bool {
		if node == nil {
			return false
		}

		return same(node, subRoot) || fn(node.Left) || fn(node.Right)
	}

	return fn(root)
}

//leetcode submit region end(Prohibit modification and deletion)

func isSubtree_A(root *TreeNode, subRoot *TreeNode) bool {
	var (
		fn   func(node *TreeNode) bool
		same func(a, b *TreeNode) bool
	)

	same = func(a, b *TreeNode) bool {
		if a == nil && b == nil {
			return true
		}

		if a == nil || b == nil {
			return false
		}

		if a.Val != b.Val {
			return false
		}

		if !same(a.Left, b.Left) || !same(a.Right, b.Right) {
			return false
		}

		return true
	}

	fn = func(node *TreeNode) bool {
		if node == nil {
			return false
		}

		if same(node, subRoot) {
			return true
		}

		if fn(node.Left) || fn(node.Right) {
			return true
		}

		return false
	}

	return fn(root)
}
