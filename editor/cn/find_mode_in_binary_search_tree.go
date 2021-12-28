package main

func main() {

}

//给定一个有相同值的二叉搜索树（BST），找出 BST 中的所有众数（出现频率最高的元素）。
//
// 假定 BST 有如下定义：
//
//
// 结点左子树中所含结点的值小于等于当前结点的值
// 结点右子树中所含结点的值大于等于当前结点的值
// 左子树和右子树都是二叉搜索树
//
//
// 例如：
//给定 BST [1,null,2,2],
//
//    1
//    \
//     2
//    /
//   2
//
//
// 返回[2].
//
// 提示：如果众数超过1个，不需考虑输出顺序
//
// 进阶：你可以不使用额外的空间吗？（假设由递归产生的隐式调用栈的开销不被计算在内）
// Related Topics 树 深度优先搜索 二叉搜索树 二叉树 👍 388 👎 0

//leetcode submit region begin(Prohibit modification and deletion)
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func findMode(root *TreeNode) []int {
	var cnt, val, maxCnt int
	var fn func(root *TreeNode)
	var ret []int

	fn = func(root *TreeNode) {
		if root == nil {
			return
		}

		fn(root.Left)

		if val == root.Val {
			cnt++
		} else {
			val, cnt = root.Val, 1
		}

		if cnt > maxCnt {
			maxCnt = cnt
			ret = []int{root.Val}
		} else if cnt == maxCnt {
			ret = append(ret, root.Val)
		}

		fn(root.Right)
	}

	fn(root)

	return ret
}

//leetcode submit region end(Prohibit modification and deletion)
