package main

func main() {

}

//给定一个 N 叉树，返回其节点值的 前序遍历 。
//
// N 叉树 在输入中按层序遍历进行序列化表示，每组子节点由空值 null 分隔（请参见示例）。
//
//
//
//
//
// 进阶：
//
// 递归法很简单，你可以使用迭代法完成此题吗?
//
//
//
// 示例 1：
//
//
//
//
//输入：root = [1,null,3,2,4,null,5,6]
//输出：[1,3,5,6,2,4]
//
//示例 2：
//
//
//
//
//输入：root = [1,null,2,3,4,5,null,null,6,7,null,8,null,9,10,null,null,11,null,12,
//null,13,null,null,14]
//输出：[1,2,3,6,7,11,14,4,8,12,5,9,13,10]
//
//
//
//
// 提示：
//
//
// N 叉树的高度小于或等于 1000
// 节点总数在范围 [0, 10^4] 内
//
//
//
// Related Topics 栈 树 深度优先搜索 👍 199 👎 0

type Node struct {
	Val      int
	Children []*Node
}

//leetcode submit region begin(Prohibit modification and deletion)
/**
 * Definition for a Node.
 * type Node struct {
 *     Val int
 *     Children []*Node
 * }
 */

func preorder(root *Node) []int {
	var ret []int

	if root == nil {
		return ret
	}

	var q = []*Node{root}

	for len(q) > 0 {
		var node = q[0]
		ret = append(ret, node.Val)

		q = q[1:]
		//var tmp = make([]*Node, 0)
		for k, v := range node.Children {
			//tmp = append(tmp, v)
			q = append(q[:k], append([]*Node{v}, q[k:]...)...)
		}
	}

	return ret
}

//leetcode submit region end(Prohibit modification and deletion)

func preorder_A(root *Node) []int {
	var ret []int
	var fn func(node *Node)

	fn = func(node *Node) {
		if node == nil {
			return
		}

		ret = append(ret, node.Val)
		for _, v := range node.Children {
			fn(v)
		}
	}

	fn(root)
	return ret
}
