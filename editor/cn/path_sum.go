package main

func main() {

}

//给你二叉树的根节点 root 和一个表示目标和的整数 targetSum ，判断该树中是否存在 根节点到叶子节点 的路径，这条路径上所有节点值相加等于目标和
// targetSum 。
//
// 叶子节点 是指没有子节点的节点。
//
//
//
// 示例 1：
//
//
//输入：root = [5,4,8,11,null,13,4,7,2,null,null,null,1], targetSum = 22
//输出：true
//
//
// 示例 2：
//
//
//输入：root = [1,2,3], targetSum = 5
//输出：false
//
//
// 示例 3：
//
//
//输入：root = [1,2], targetSum = 0
//输出：false
//
//
//
//
// 提示：
//
//
// 树中节点的数目在范围 [0, 5000] 内
// -1000 <= Node.val <= 1000
// -1000 <= targetSum <= 1000
//
// Related Topics 树 深度优先搜索 二叉树 👍 720 👎 0

//leetcode submit region begin(Prohibit modification and deletion)
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func hasPathSum(root *TreeNode, targetSum int) bool {
	if root == nil {
		return false
	}

	targetSum -= root.Val

	if root.Left == nil && root.Right == nil {
		return targetSum == 0
	}

	if root.Left != nil {
		if hasPathSum(root.Left, targetSum) {
			return true
		}
	}

	if root.Right != nil {
		if hasPathSum(root.Right, targetSum) {
			return true
		}
	}

	return false
}

//leetcode submit region end(Prohibit modification and deletion)
