package main

import "fmt"

func main() {
	var head = NewListNode(0)
	head.AppendValue(0).AppendValue(1).AppendValue(0)
	deleteDuplicates(head)
	head.Print()
}

//存在一个按升序排列的链表，给你这个链表的头节点 head ，请你删除所有重复的元素，使每个元素 只出现一次 。
//
//返回同样按升序排列的结果链表。
//
//
//
//示例 1：
//
//
//输入：head = [1,1,2]
//输出：[1,2]
//示例 2：
//
//
//输入：head = [1,1,2,3,3]
//输出：[1,2,3]
//
//提示：
//
//链表中节点数目在范围 [0, 300] 内
//-100 <= Node.val <= 100
//题目数据保证链表已经按升序排列

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func deleteDuplicates(head *ListNode) *ListNode {
	return deleteDuplicates_B(head)
}

// 题目数据保证链表已经按升序排列
func deleteDuplicates_B(head *ListNode) *ListNode {
	if head == nil {
		return head
	}

	var (
		prev *ListNode
		min  = head.Val - 1
	)
	for c := head; c != nil; c = c.Next {
		if c.Val <= min {
			var tmp = c.Next
			prev.Next = tmp
			c.Next = nil
			c = prev
		} else {
			prev = c
			min = c.Val
		}
	}

	return head
}

func deleteDuplicates_A(head *ListNode) *ListNode {
	var exists = make(map[int]bool)
	var prev *ListNode
	for c := head; c != nil; c = c.Next {
		if _, has := exists[c.Val]; has {
			var tmp = c.Next
			prev.Next = tmp
			c.Next = nil
			c = prev
		} else {
			prev = c
			exists[c.Val] = true
		}
	}

	return head
}

type ListNode struct {
	Val  int
	Next *ListNode
}

func (n *ListNode) AppendValue(val int) *ListNode {
	var c = NewListNode(val)
	c.Next = n.Next
	n.Next = c
	return c
}

func (n *ListNode) Print() {
	var i = make([]interface{}, 0)
	for c := n; c != nil; c = c.Next {
		i = append(i, c.Val)
	}

	fmt.Printf("%#v \n", i)
}

func NewListNode(val int) *ListNode {
	return &ListNode{Val: val}
}
