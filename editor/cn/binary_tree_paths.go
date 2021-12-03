package main

import (
	"fmt"
	"strings"
)

func main() {

}

//给你一个二叉树的根节点 root ，按 任意顺序 ，返回所有从根节点到叶子节点的路径。
//
// 叶子节点 是指没有子节点的节点。
//
//
// 示例 1：
//
//
//输入：root = [1,2,3,null,5]
//输出：["1->2->5","1->3"]
//
//
// 示例 2：
//
//
//输入：root = [1]
//输出：["1"]
//
//
//
//
// 提示：
//
//
// 树中节点的数目在范围 [1, 100] 内
// -100 <= Node.val <= 100
//
// Related Topics 树 深度优先搜索 字符串 二叉树 👍 612 👎 0

//leetcode submit region begin(Prohibit modification and deletion)
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func binaryTreePaths(root *TreeNode) []string {
	var (
		ps = make([]string, 0)
		fn func(root *TreeNode, path []string)
	)

	fn = func(root *TreeNode, path []string) {
		if root == nil {
			return
		}

		path = append(path, fmt.Sprintf("%d", root.Val))

		if root.Left != nil {
			fn(root.Left, path)
		}

		if root.Right != nil {
			fn(root.Right, path)
		}

		if root.Left == nil && root.Right == nil {
			ps = append(ps, strings.Join(path, "->"))
		}
	}

	fn(root, make([]string, 0))

	return ps
}

//leetcode submit region end(Prohibit modification and deletion)
