package main

func main() {

}

//给你二叉树的根节点 root ，返回它节点值的 前序 遍历。
//
//
//
// 示例 1：
//
//
//输入：root = [1,null,2,3]
//输出：[1,2,3]
//
//
// 示例 2：
//
//
//输入：root = []
//输出：[]
//
//
// 示例 3：
//
//
//输入：root = [1]
//输出：[1]
//
//
// 示例 4：
//
//
//输入：root = [1,2]
//输出：[1,2]
//
//
// 示例 5：
//
//
//输入：root = [1,null,2]
//输出：[1,2]
//
//
//
//
// 提示：
//
//
// 树中节点数目在范围 [0, 100] 内
// -100 <= Node.val <= 100
//
//
//
//
// 进阶：递归算法很简单，你可以通过迭代算法完成吗？
// Related Topics 栈 树 深度优先搜索 二叉树 👍 681 👎 0

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

//leetcode submit region begin(Prohibit modification and deletion)
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func preorderTraversal(root *TreeNode) []int {
	if root == nil {
		return nil
	}

	var (
		r     = make([]int, 0)
		stack = make([]*TreeNode, 0)
	)

	stack = append(stack, root)

	for n := root; len(stack) > 0; n, stack = stack[len(stack)-1], stack[0:len(stack)-1] {
		if n == nil {
			continue
		}

		r = append(r, n.Val)

		stack = append(stack, n.Right, n.Left)
	}

	return r
}

//leetcode submit region end(Prohibit modification and deletion)

func preorderTraversal_A(root *TreeNode) []int {
	if root == nil {
		return nil
	}

	var r = make([]int, 0)

	r = append(r, root.Val)

	if c := preorderTraversal(root.Left); c != nil {
		r = append(r, c...)
	}

	if c := preorderTraversal(root.Right); c != nil {
		r = append(r, c...)
	}

	return r
}
