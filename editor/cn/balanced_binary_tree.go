package main

import (
	"fmt"
)

func main() {
	fmt.Println(
		isBalanced(&TreeNode{
			Val: 1,
			Left: &TreeNode{Val: 2,
				Left:  &TreeNode{Val: 4, Left: &TreeNode{Val: 8}},
				Right: &TreeNode{Val: 5}},
			Right: &TreeNode{Val: 3,
				Left:  &TreeNode{Val: 6},
				Right: nil,
			},
		}),
	)
}

//给定一个二叉树，判断它是否是高度平衡的二叉树。
//
// 本题中，一棵高度平衡二叉树定义为：
//
//
// 一个二叉树每个节点 的左右两个子树的高度差的绝对值不超过 1 。
//
//
//
//
// 示例 1：
//
//
//输入：root = [3,9,20,null,null,15,7]
//输出：true
//
//
// 示例 2：
//
//
//输入：root = [1,2,2,3,3,null,null,4,4]
//输出：false
//
//
// 示例 3：
//
//
//输入：root = []
//输出：true
//
//
//
//
// 提示：
//
//
// 树中的节点数在范围 [0, 5000] 内
// -10⁴ <= Node.val <= 10⁴
//
// Related Topics 树 深度优先搜索 二叉树 👍 808 👎 0

//[3,9,20,null,null,15,7]

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
func isBalanced(root *TreeNode) bool {
	if root == nil {
		return true
	}

	var fn func(n *TreeNode) (int, bool)
	fn = func(n *TreeNode) (int, bool) {
		if n == nil {
			return 0, true
		}

		var ll, lb = fn(n.Left)
		var rl, rb = fn(n.Right)
		var sub = ll - rl

		//fmt.Printf("%#v \n", n)
		//fmt.Println(ll, lb, rl, rb, sub)

		if !lb || !rb || (sub > 1 || sub < -1) {
			return sub, false
		}

		return isBalancedMax(ll, rl) + 1, true
	}

	var _, b = fn(root)
	return b

	//var ll = fn(root.Left)
	//var rl = fn(root.Right)
	//var sub = ll - rl
	//
	//return !(sub > 1 || sub < -1)
}

func isBalancedMax(a, b int) int {
	if a > b {
		return a
	}
	return b
}

//             1,
//       2,         2,
//    3,  null,  null,3,
//4,null,            null,4

//leetcode submit region end(Prohibit modification and deletion)
