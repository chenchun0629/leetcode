package main

import "fmt"

func main() {
	//fmt.Println(removeNthFromEnd(&ListNode{1, &ListNode{2, &ListNode{3, &ListNode{4, &ListNode{5, nil}}}}}, 2))
	fmt.Println(removeNthFromEnd(&ListNode{1, nil}, 1))
}

//给你一个链表，删除链表的倒数第 n 个结点，并且返回链表的头结点。
//
//
//
// 示例 1：
//
//
//输入：head = [1,2,3,4,5], n = 2
//输出：[1,2,3,5]
//
//
// 示例 2：
//
//
//输入：head = [1], n = 1
//输出：[]
//
//
// 示例 3：
//
//
//输入：head = [1,2], n = 1
//输出：[1]
//
//
//
//
// 提示：
//
//
// 链表中结点的数目为 sz
// 1 <= sz <= 30
// 0 <= Node.val <= 100
// 1 <= n <= sz
//
//
//
//
// 进阶：你能尝试使用一趟扫描实现吗？
// Related Topics 链表 双指针 👍 1716 👎 0

//type ListNode struct {
//	Val  int
//	Next *ListNode
//}

//leetcode submit region begin(Prohibit modification and deletion)
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func removeNthFromEnd(head *ListNode, n int) *ListNode {
	head = &ListNode{Next: head}
	var fn func(head *ListNode) int

	fn = func(head *ListNode) int {
		if head == nil {
			return 0
		}

		var idx = fn(head.Next) + 1
		if idx == n+1 {
			head.Next = head.Next.Next
		}
		return idx
	}

	fn(head)
	return head.Next
}

//leetcode submit region end(Prohibit modification and deletion)
