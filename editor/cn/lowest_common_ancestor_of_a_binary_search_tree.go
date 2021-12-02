package main

import (
	"fmt"
)

func main() {
	var root = &TreeNode{Val: 6,
		Left:  &TreeNode{Val: 2, Left: &TreeNode{Val: 0}, Right: &TreeNode{Val: 4, Left: &TreeNode{Val: 3}, Right: &TreeNode{Val: 5}}},
		Right: &TreeNode{Val: 8, Left: &TreeNode{Val: 7}, Right: &TreeNode{Val: 9}},
	}
	var p = &TreeNode{Val: 2}
	var q = &TreeNode{Val: 4}
	var n = lowestCommonAncestor(root, p, q)
	fmt.Println(n)
}

//给定一个二叉搜索树, 找到该树中两个指定节点的最近公共祖先。
//
// 百度百科中最近公共祖先的定义为：“对于有根树 T 的两个结点 p、q，最近公共祖先表示为一个结点 x，满足 x 是 p、q 的祖先且 x 的深度尽可能大（
//一个节点也可以是它自己的祖先）。”
//
// 例如，给定如下二叉搜索树: root = [6,2,8,0,4,7,9,null,null,3,5]
//
//
//
//
//
// 示例 1:
//
// 输入: root = [6,2,8,0,4,7,9,null,null,3,5], p = 2, q = 8
//输出: 6
//解释: 节点 2 和节点 8 的最近公共祖先是 6。
//
//
// 示例 2:
//
// 输入: root = [6,2,8,0,4,7,9,null,null,3,5], p = 2, q = 4
//输出: 2
//解释: 节点 2 和节点 4 的最近公共祖先是 2, 因为根据定义最近公共祖先节点可以为节点本身。
//
//
//
// 说明:
//
//
// 所有节点的值都是唯一的。
// p、q 为不同节点且均存在于给定的二叉搜索树中。
//
// Related Topics 树 深度优先搜索 二叉搜索树 二叉树 👍 707 👎 0

//type TreeNode struct {
//	Val   int
//	Left  *TreeNode
//	Right *TreeNode
//}

//leetcode submit region begin(Prohibit modification and deletion)
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val   int
 *     Left  *TreeNode
 *     Right *TreeNode
 * }
 */

func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
	var r = root

	for r != nil {
		if p.Val < r.Val && q.Val < r.Val {
			r = r.Left
		} else if p.Val > r.Val && q.Val > r.Val {
			r = r.Right
		} else {
			return r
		}
	}

	return r
}

//leetcode submit region end(Prohibit modification and deletion)

func lowestCommonAncestor_A(root, p, q *TreeNode) *TreeNode {
	var (
		pp = lowestCommonAncestorGetPath(root, p)
		qp = lowestCommonAncestorGetPath(root, q)
		lp = len(pp)
		lq = len(qp)
		r  *TreeNode
	)
	for idx := 0; idx < lp && idx < lq; idx++ {
		if pp[idx].Val == qp[idx].Val {
			r = pp[idx]
		} else {
			break
		}
	}

	return r
}

func lowestCommonAncestorGetPath(root, p *TreeNode) []*TreeNode {
	if root == nil {
		return nil
	}

	if root.Val == p.Val {
		return []*TreeNode{root}
	}

	var l = lowestCommonAncestorGetPath(root.Left, p)

	if l != nil {
		l = append([]*TreeNode{root}, l...)
		return l
	}

	l = lowestCommonAncestorGetPath(root.Right, p)
	if l != nil {
		l = append([]*TreeNode{root}, l...)
		return l
	}

	return l
}
