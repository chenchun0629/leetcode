package main

import "fmt"

func main() {
	fmt.Println(9 & 3)
	fmt.Println(10 & 3)
	fmt.Println(11 & 3)
	fmt.Println(12 & 3)
	fmt.Println(13 & 3)
	fmt.Println(14 & 3)
	fmt.Println(15 & 3)
}

//给你一个整数 n ，找出从 1 到 n 各个整数的 Fizz Buzz 表示，并用字符串数组 answer（下标从 1 开始）返回结果，其中：
//
//
// answer[i] == "FizzBuzz" 如果 i 同时是 3 和 5 的倍数。
// answer[i] == "Fizz" 如果 i 是 3 的倍数。
// answer[i] == "Buzz" 如果 i 是 5 的倍数。
// answer[i] == i （以字符串形式）如果上述条件全不满足。
//
//
//
//
// 示例 1：
//
//
//输入：n = 3
//输出：["1","2","Fizz"]
//
//
// 示例 2：
//
//
//输入：n = 5
//输出：["1","2","Fizz","4","Buzz"]
//
//
// 示例 3：
//
//
//输入：n = 15
//输出：["1","2","Fizz","4","Buzz","Fizz","7","8","Fizz","Buzz","11","Fizz","13","1
//4","FizzBuzz"]
//
//
//
// 提示：
//
//
// 1 <= n <= 10⁴
//
// Related Topics 数学 字符串 模拟 👍 159 👎 0

//leetcode submit region begin(Prohibit modification and deletion)
func fizzBuzz(n int) []string {
	var s []string

	for i := 1; i <= n; i++ {
		if i%15 == 0 {
			s = append(s, "FizzBuzz")
		} else if i%3 == 0 {
			s = append(s, "Fizz")
		} else if i%5 == 0 {
			s = append(s, "Buzz")
		} else {
			s = append(s, fmt.Sprint(i))
		}
	}

	return s
}

//leetcode submit region end(Prohibit modification and deletion)
