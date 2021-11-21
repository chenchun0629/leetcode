package main

func main() {

}

//给定一个二叉树，返回它的 后序 遍历。
//
// 示例:
//
// 输入: [1,null,2,3]
//   1
//    \
//     2
//    /
//   3
//
//输出: [3,2,1]
//
// 进阶: 递归算法很简单，你可以通过迭代算法完成吗？
// Related Topics 栈 树 深度优先搜索 二叉树 👍 710 👎 0

//leetcode submit region begin(Prohibit modification and deletion)
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func postorderTraversal(root *TreeNode) []int {
	if root == nil {
		return nil
	}

	var (
		r     = make([]int, 0)
		stack = make([]*TreeNode, 0)
		prev  *TreeNode
	)

	for root != nil || len(stack) > 0 {
		for root != nil {
			stack = append(stack, root)
			root = root.Left
		}

		root = stack[len(stack)-1]
		stack = stack[:len(stack)-1]

		// 后续遍历的循环逻辑有点复杂
		if root.Right == nil || root.Right == prev {
			r = append(r, root.Val)
			prev = root
			root = nil
		} else {
			stack = append(stack, root)
			root = root.Right
		}

	}

	return r
}

//leetcode submit region end(Prohibit modification and deletion)

func postorderTraversal_A(root *TreeNode) []int {
	if root == nil {
		return nil
	}

	var r = make([]int, 0)

	if c := postorderTraversal(root.Left); c != nil {
		r = append(r, c...)
	}

	if c := postorderTraversal(root.Right); c != nil {
		r = append(r, c...)
	}

	r = append(r, root.Val)

	return r
}
