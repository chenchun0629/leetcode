package main

import "fmt"

func main() {
	var root = &TreeNode{Val: 5}
	root.Left = &TreeNode{Val: 3}
	root.Right = &TreeNode{Val: 7}
	var res = inorderTraversal(root)
	fmt.Printf("%#v \n", res)
}

//给定一个二叉树的根节点 root ，返回它的 中序 遍历。
//
//
//
// 示例 1：
//
//
//输入：root = [1,null,2,3]
//输出：[1,3,2]
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
//输出：[2,1]
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
// 进阶: 递归算法很简单，你可以通过迭代算法完成吗？
// Related Topics 栈 树 深度优先搜索 二叉树 👍 1156 👎 0

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
func inorderTraversal(root *TreeNode) []int {
	var is = make([]int, 0)
	if root == nil {
		return is
	}

	var fn func(tn *TreeNode, is *[]int)
	fn = func(tn *TreeNode, is *[]int) {
		if tn.Left != nil {
			fn(tn.Left, is)
		}

		*is = append(*is, tn.Val)

		if tn.Right != nil {
			fn(tn.Right, is)
		}
	}

	fn(root, &is)
	return is
}

//leetcode submit region end(Prohibit modification and deletion)

func inorderTraversal_A(root *TreeNode) []int {
	//var fn func(tn *TreeNode)  []int
	//fn = func(tn *TreeNode)  []int {
	//	var is = make([]int, 0)
	//	if tn.Left != nil {
	//		is = append(is, fn(tn.Left)...)
	//	}
	//
	//	is = append(is, tn.Val)
	//
	//	if tn.Right != nil {
	//		is = append(is, fn(tn.Right)...)
	//	}
	//
	//	return is
	//}

	var is = make([]int, 0)

	if root == nil {
		return is
	}

	if root.Left != nil {
		is = append(is, inorderTraversal(root.Left)...)
	}

	is = append(is, root.Val)

	if root.Right != nil {
		is = append(is, inorderTraversal(root.Right)...)
	}

	return is
}
