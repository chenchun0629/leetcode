package main

import "fmt"

func main() {
	var head = &ListNode{
		Val: 1,
		Next: &ListNode{
			Val:  2,
			Next: nil,
		},
	}

	reverseList(head).Print()
}

//给你单链表的头节点 head ，请你反转链表，并返回反转后的链表。
//
//
//
//
// 示例 1：
//
//
//输入：head = [1,2,3,4,5]
//输出：[5,4,3,2,1]
//
//
// 示例 2：
//
//
//输入：head = [1,2]
//输出：[2,1]
//
//
// 示例 3：
//
//
//输入：head = []
//输出：[]
//
//
//
//
// 提示：
//
//
// 链表中节点的数目范围是 [0, 5000]
// -5000 <= Node.val <= 5000
//
//
//
//
// 进阶：链表可以选用迭代或递归方式完成反转。你能否用两种方法解决这道题？
//
//
// Related Topics 递归 链表 👍 2119 👎 0

type ListNode struct {
	Val  int
	Next *ListNode
}

func (n *ListNode) Print() {
	var i = make([]interface{}, 0)
	for c := n; c != nil; c = c.Next {
		i = append(i, c.Val)
	}

	fmt.Printf("%#v \n", i)
}

//leetcode submit region begin(Prohibit modification and deletion)
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func reverseList(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}

	var (
		prev, curr *ListNode = nil, head
	)

	for curr != nil {
		var next = curr.Next
		curr.Next = prev
		prev = curr
		curr = next
	}
	return prev
}

//leetcode submit region end(Prohibit modification and deletion)

func reverseList_A(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}

	var last = reverseList(head.Next)

	head.Next.Next = head
	head.Next = nil
	return last
}
