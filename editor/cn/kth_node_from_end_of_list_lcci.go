package main

func main() {

}

//实现一种算法，找出单向链表中倒数第 k 个节点。返回该节点的值。
//
// 注意：本题相对原题稍作改动
//
// 示例：
//
// 输入： 1->2->3->4->5 和 k = 2
//输出： 4
//
// 说明：
//
// 给定的 k 保证是有效的。
// Related Topics 链表 双指针 👍 90 👎 0

//leetcode submit region begin(Prohibit modification and deletion)
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func kthToLast(head *ListNode, k int) int {
	var (
		bn = head
		en = head
	)

	for i := 0; bn != nil; bn = bn.Next {
		if i < k {
			i++
		} else {
			en = en.Next
		}
	}

	return en.Val
}

//leetcode submit region end(Prohibit modification and deletion)
