package main

import (
	"encoding/json"
	"fmt"
)

//将两个升序链表合并为一个新的 升序 链表并返回。新链表是通过拼接给定的两个链表的所有节点组成的。
//
//
//
// 示例：
//
// 输入：1->2->4, 1->3->4
//输出：1->1->2->3->4->4
//
// Related Topics 链表
// 👍 1392 👎 0

//leetcode submit region begin(Prohibit modification and deletion)

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	return mergeTwoListsA(l1, l2)
}

// 方案A：递归
// l1和l2的值对比，那边的值小，就取那边的next继续递归下去
// 直至，l1 or l2 == nil
// fn([1, 3], [2, 4])
// 1.next = fn([3], [2, 4]) = 2
// 2.next = fn([3], [4]) = 3
// 3.next = fn(nil, [4]) = 4
func mergeTwoListsA(l1 *ListNode, l2 *ListNode) *ListNode {
	if l2 == nil {
		return l1
	} else if l1 == nil {
		return l2
	}

	if l1.Val <= l2.Val {
		l1.Next = mergeTwoListsA(l1.Next, l2)
		return l1
	} else {
		l2.Next = mergeTwoListsA(l1, l2.Next)
		return l2
	}
}

//leetcode submit region end(Prohibit modification and deletion)

type ListNode struct {
	Val  int
	Next *ListNode
}

func main() {
	//var (
	//	l1n3 = &ListNode{4, nil}
	//	l1n2 = &ListNode{2, l1n3}
	//	l1n1 = &ListNode{1, l1n2}
	//
	//	l2n3 = &ListNode{4, nil}
	//	l2n2 = &ListNode{3, l2n3}
	//	l2n1 = &ListNode{1, l2n2}
	//)
	//

	var (
		l1n3 = &ListNode{3, nil}
		l1n1 = &ListNode{-9, l1n3}

		l2n3 = &ListNode{7, nil}
		l2n1 = &ListNode{5, l2n3}
	)

	var x, err = json.Marshal(mergeTwoLists(l1n1, l2n1))
	fmt.Println(err)
	fmt.Println(string(x))
}
