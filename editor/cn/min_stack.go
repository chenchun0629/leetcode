package main

import "fmt"

func main() {
	obj := Constructor()
	obj.Push(5)
	obj.Push(2)
	obj.Push(3)
	obj.Push(1)
	obj.Push(4)
	obj.Pop()
	param_3 := obj.Top()
	param_4 := obj.GetMin()
	fmt.Println(param_4, param_3)
}

//设计一个支持 push ，pop ，top 操作，并能在常数时间内检索到最小元素的栈。
//
//
// push(x) —— 将元素 x 推入栈中。
// pop() —— 删除栈顶的元素。
// top() —— 获取栈顶元素。
// getMin() —— 检索栈中的最小元素。
//
//
//
//
// 示例:
//
// 输入：
//["MinStack","push","push","push","getMin","pop","top","getMin"]
//[[],[-2],[0],[-3],[],[],[],[]]
//
//输出：
//[null,null,null,null,-3,null,0,-2]
//
//解释：
//MinStack minStack = new MinStack();
//minStack.push(-2);
//minStack.push(0);
//minStack.push(-3);
//minStack.getMin();   --> 返回 -3.
//minStack.pop();
//minStack.top();      --> 返回 0.
//minStack.getMin();   --> 返回 -2.
//
//
//
//
// 提示：
//
//
// pop、top 和 getMin 操作总是在 非空栈 上调用。
//
// Related Topics 栈 设计 👍 1087 👎 0

//leetcode submit region begin(Prohibit modification and deletion)
type MinStackNode struct {
	val  int
	min  int
	next *MinStackNode
}

type MinStack struct {
	head *MinStackNode
}

func Constructor() MinStack {
	return MinStack{}
}

func (this *MinStack) Push(val int) {
	if this.head == nil {
		this.head = &MinStackNode{
			val:  val,
			min:  val,
			next: nil,
		}
	} else {
		this.head = &MinStackNode{
			val:  val,
			min:  this.min(val, this.head.min),
			next: this.head,
		}
	}
}

func (s MinStack) min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func (this *MinStack) Pop() {
	if this.head == nil {
		return
	}
	var top = this.head
	this.head = top.next
	top.next = nil
}

func (this *MinStack) Top() int {
	return this.head.val
}

func (this *MinStack) GetMin() int {
	return this.head.min
}

/**
 * Your MinStack object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(val);
 * obj.Pop();
 * param_3 := obj.Top();
 * param_4 := obj.GetMin();
 */
//leetcode submit region end(Prohibit modification and deletion)
