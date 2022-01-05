package main

func main() {

}

//假设你正在爬楼梯。需要 n 阶你才能到达楼顶。
//
// 每次你可以爬 1 或 2 个台阶。你有多少种不同的方法可以爬到楼顶呢？
//
// 注意：给定 n 是一个正整数。
//
// 示例 1：
//
// 输入： 2
//输出： 2
//解释： 有两种方法可以爬到楼顶。
//1.  1 阶 + 1 阶
//2.  2 阶
//
// 示例 2：
//
// 输入： 3
//输出： 3
//解释： 有三种方法可以爬到楼顶。
//1.  1 阶 + 1 阶 + 1 阶
//2.  1 阶 + 2 阶
//3.  2 阶 + 1 阶
//
// Related Topics 记忆化搜索 数学 动态规划 👍 2100 👎 0

//leetcode submit region begin(Prohibit modification and deletion)

var mem = make(map[int]int)

func climbStairs(n int) int {
	if n <= 2 {
		return n
	}

	if d, has := mem[n]; has {
		return d
	}

	var s = 0
	s += climbStairs(n - 1)
	s += climbStairs(n - 2)
	mem[n] = s

	return s
}

//leetcode submit region end(Prohibit modification and deletion)
func climbStairs_A(n int) int {
	if n == 1 {
		return 1
	}
	if n == 2 {
		return 2
	}

	var a, b, tmp = 1, 2, 0
	for i := 3; i <= n; i++ {
		tmp = b
		b = a + b
		a = tmp
	}

	return b
}
