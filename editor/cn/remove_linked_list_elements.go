package main

func main() {
	var l = &ListNode{
		Val: 1,
		Next: &ListNode{
			Val: 2,
			Next: &ListNode{
				Val: 2,
				Next: &ListNode{
					Val:  1,
					Next: nil,
				},
			},
		},
	}

	l = removeElements(l, 2)
	l.Print()
}

//给你一个链表的头节点 head 和一个整数 val ，请你删除链表中所有满足 Node.val == val 的节点，并返回 新的头节点 。
//
//
// 示例 1：
//
//
//输入：head = [1,2,6,3,4,5,6], val = 6
//输出：[1,2,3,4,5]
//
//
// 示例 2：
//
//
//输入：head = [], val = 1
//输出：[]
//
//
// 示例 3：
//
//
//输入：head = [7,7,7,7], val = 7
//输出：[]
//
//
//
//
// 提示：
//
//
// 列表中的节点数目在范围 [0, 10⁴] 内
// 1 <= Node.val <= 50
// 0 <= val <= 50
//
// Related Topics 递归 链表 👍 741 👎 0

//type ListNode struct {
//	Val  int
//	Next *ListNode
//}
//
//func (n *ListNode) Print() {
//	var i = make([]interface{}, 0)
//	for c := n; c != nil; c = c.Next {
//		i = append(i, c.Val)
//	}
//
//	fmt.Printf("%#v \n", i)
//}

//leetcode submit region begin(Prohibit modification and deletion)
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func removeElements(head *ListNode, val int) *ListNode {
	if head == nil {
		return head
	}

	head.Next = removeElements(head.Next, val)

	if head.Val == val {
		return head.Next
	}

	return head
}

//leetcode submit region end(Prohibit modification and deletion)

func removeElements_A(head *ListNode, val int) *ListNode {
	var t = &ListNode{
		Next: head,
	}

	for h := t; h.Next != nil; {
		if h.Next.Val == val {
			h.Next = h.Next.Next
		} else {
			h = h.Next
		}
	}
	return t.Next
}
