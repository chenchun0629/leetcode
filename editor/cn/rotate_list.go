package main

func main() {
	//var head = &ListNode{Val: 1}
	//head.AppendValue(2).AppendValue(3).AppendValue(4).AppendValue(5)
	//head.Print()
	rotateRight(nil, 7).Print()
}

//给你一个链表的头节点 head ，旋转链表，将链表每个节点向右移动 k 个位置。
//
//
//
// 示例 1：
//
//
//输入：head = [1,2,3,4,5], k = 2
//输出：[4,5,1,2,3]
//
//
// 示例 2：
//
//
//输入：head = [0,1,2], k = 4
//输出：[2,0,1]
//
//
//
//
// 提示：
//
//
// 链表中节点的数目在范围 [0, 500] 内
// -100 <= Node.val <= 100
// 0 <= k <= 2 * 10⁹
//
// Related Topics 链表 双指针 👍 699 👎 0
//type ListNode struct {
//	Val  int
//	Next *ListNode
//}
//
//func (n *ListNode) AppendValue(val int) *ListNode {
//	var c = &ListNode{Val: val}
//	n.Next = c
//	return c
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
func rotateRight(head *ListNode, k int) *ListNode {
	if head == nil {
		return head
	}

	var (
		fn   func(node *ListNode) int
		h    = &ListNode{Val: 0, Next: nil}
		last *ListNode
		n    int
	)

	for i := head; i != nil; i = i.Next {
		n++
	}

	n = k % n
	if n == 0 {
		return head
	}

	fn = func(node *ListNode) int {
		if node == nil {
			return 0
		}

		var cur = fn(node.Next) + 1
		if cur == 1 {
			last = node
		}

		if cur == n {
			h.Next = node
		} else if cur == n+1 {
			node.Next = nil
		}

		return cur
	}

	fn(head)
	last.Next = head

	return h.Next
}

//leetcode submit region end(Prohibit modification and deletion)
