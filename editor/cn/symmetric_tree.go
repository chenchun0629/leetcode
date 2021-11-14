package main

import "fmt"

func main() {
	var r = &TreeNode{
		Val:   1,
		Left:  &TreeNode{Val: 2, Left: &TreeNode{Val: 3}, Right: &TreeNode{Val: 4, Left: &TreeNode{Val: 1}}},
		Right: &TreeNode{Val: 2, Left: &TreeNode{Val: 4}, Right: &TreeNode{Val: 3}},
	}

	fmt.Println(isSymmetric(r))
}

//给定一个二叉树，检查它是否是镜像对称的。
//
//
//
// 例如，二叉树 [1,2,2,3,4,4,3] 是对称的。
//
//     1
//   / \
//  2   2
// / \ / \
//3  4 4  3
//
//
//
//
// 但是下面这个 [1,2,2,null,3,null,3] 则不是镜像对称的:
//
//     1
//   / \
//  2   2
//   \   \
//   3    3
//
//
//
//
// 进阶：
//
// 你可以运用递归和迭代两种方法解决这个问题吗？
// Related Topics 树 深度优先搜索 广度优先搜索 二叉树 👍 1600 👎 0

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
func isSymmetric(root *TreeNode) bool {
	if root == nil {
		return true
	}

	var queue = make([]*TreeNode, 0)
	queue = append(queue, root.Left, root.Right)

	for len(queue) > 0 {
		var v = queue[0]
		var u = queue[1]
		queue = queue[2:]
		if v == nil && u == nil {
			continue
		}

		if v == nil || u == nil {
			return false
		}

		if v.Val != u.Val {
			return false
		}

		queue = append(queue, v.Left, u.Right, v.Right, u.Left)
	}

	return true
}

//leetcode submit region end(Prohibit modification and deletion)

func isSymmetric_B(root *TreeNode) bool {
	if root == nil {
		return true
	}

	var fn func(left *TreeNode, right *TreeNode) bool
	fn = func(left *TreeNode, right *TreeNode) bool {
		if left == nil && right == nil {
			return true
		}

		if left == nil || right == nil ||
			left.Val != right.Val ||
			!fn(left.Left, right.Right) ||
			!fn(left.Right, right.Left) {
			return false
		}

		return true
	}

	return fn(root.Left, root.Right)
}

func isSymmetric_A(root *TreeNode) bool {
	if root == nil {
		return true
	}
	var queue = make([][2]*TreeNode, 0)
	queue = append(queue, [2]*TreeNode{root.Left, root.Right})

	for i := 0; i < len(queue); i++ {
		var v = queue[i]
		if v[0] == nil && v[1] != nil || v[0] != nil && v[1] == nil {
			return false
		}

		if v[0] == nil && v[1] == nil {

		} else {
			if v[0].Val != v[1].Val {
				return false
			}

			queue = append(queue, [2]*TreeNode{v[0].Left, v[1].Right}, [2]*TreeNode{v[0].Right, v[1].Left})
		}
	}

	return true
}
