package main

import (
	"fmt"
)

func main() {
	fmt.Println(climbStairs(44))
}

//70. 爬楼梯
//假设你正在爬楼梯。需要 n 阶你才能到达楼顶。
//
//每次你可以爬 1 或 2 个台阶。你有多少种不同的方法可以爬到楼顶呢？
//
//注意：给定 n 是一个正整数。
//
//示例 1：
//
//输入： 2
//输出： 2
//解释： 有两种方法可以爬到楼顶。
//1.  1 阶 + 1 阶
//2.  2 阶
//示例 2：
//
//输入： 3
//输出： 3
//解释： 有三种方法可以爬到楼顶。
//1.  1 阶 + 1 阶 + 1 阶
//2.  1 阶 + 2 阶
//3.  2 阶 + 1 阶

func climbStairs(n int) int {
	return climbStairs_B(n)
}

func climbStairs_B(n int) int {
	if n == 1 {
		return 1
	} else if n == 2 {
		return 2
	}
	var (
		a    = 1 // n = 1时， 一种方式
		b    = 2 // n = 2时， 两种方式
		temp int
	)

	// 第三次起循环
	for i := 3; i <= n; i++ {
		temp = a     // 将上上次的结果存入temp
		a = b        // 将上次结果存入上上次结果
		b = temp + b // 本次结果 = 上次结果 + 上上次结果
	}

	return b
}

//climbStairs_A 超出时间限制
func climbStairs_A(n int) int {
	if n < 0 {
		return 0
	} else if n == 0 {
		return 1
	}

	var s = 0
	s += climbStairs_A(n - 1)
	s += climbStairs_A(n - 2)

	return s
}
