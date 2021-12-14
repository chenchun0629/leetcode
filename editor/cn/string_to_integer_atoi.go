package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(myAtoi("-000000000000001"))
	//fmt.Println(myAtoi("91283472332"))
	//fmt.Println(myAtoi("words and 987"))
	//fmt.Println(myAtoi("9223372036854775808"))
	//fmt.Println(1<<31 - 1)
	//fmt.Println(-1 << 31)
}

//请你来实现一个 myAtoi(string s) 函数，使其能将字符串转换成一个 32 位有符号整数（类似 C/C++ 中的 atoi 函数）。
//
// 函数 myAtoi(string s) 的算法如下：
//
//
// 读入字符串并丢弃无用的前导空格
// 检查下一个字符（假设还未到字符末尾）为正还是负号，读取该字符（如果有）。 确定最终结果是负数还是正数。 如果两者都不存在，则假定结果为正。
// 读入下一个字符，直到到达下一个非数字字符或到达输入的结尾。字符串的其余部分将被忽略。
// 将前面步骤读入的这些数字转换为整数（即，"123" -> 123， "0032" -> 32）。如果没有读入数字，则整数为 0 。必要时更改符号（从步骤
//2 开始）。
// 如果整数数超过 32 位有符号整数范围 [−2³¹, 231 − 1] ，需要截断这个整数，使其保持在这个范围内。具体来说，小于 −2³¹ 的整数应该被固
//定为 −2³¹ ，大于 231 − 1 的整数应该被固定为 231 − 1 。
// 返回整数作为最终结果。
//
//
// 注意：
//
//
// 本题中的空白字符只包括空格字符 ' ' 。
// 除前导空格或数字后的其余字符串外，请勿忽略 任何其他字符。
//
//
//
//
// 示例 1：
//
//
//输入：s = "42"
//输出：42
//解释：加粗的字符串为已经读入的字符，插入符号是当前读取的字符。
//第 1 步："42"（当前没有读入字符，因为没有前导空格）
//         ^
//第 2 步："42"（当前没有读入字符，因为这里不存在 '-' 或者 '+'）
//         ^
//第 3 步："42"（读入 "42"）
//           ^
//解析得到整数 42 。
//由于 "42" 在范围 [-2³¹, 2³¹ - 1] 内，最终结果为 42 。
//
// 示例 2：
//
//
//输入：s = "   -42"
//输出：-42
//解释：
//第 1 步："   -42"（读入前导空格，但忽视掉）
//            ^
//第 2 步："   -42"（读入 '-' 字符，所以结果应该是负数）
//             ^
//第 3 步："   -42"（读入 "42"）
//               ^
//解析得到整数 -42 。
//由于 "-42" 在范围 [-2³¹, 2³¹ - 1] 内，最终结果为 -42 。
//
//
// 示例 3：
//
//
//输入：s = "4193 with words"
//输出：4193
//解释：
//第 1 步："4193 with words"（当前没有读入字符，因为没有前导空格）
//         ^
//第 2 步："4193 with words"（当前没有读入字符，因为这里不存在 '-' 或者 '+'）
//         ^
//第 3 步："4193 with words"（读入 "4193"；由于下一个字符不是一个数字，所以读入停止）
//             ^
//解析得到整数 4193 。
//由于 "4193" 在范围 [-2³¹, 2³¹ - 1] 内，最终结果为 4193 。
//
//
// 示例 4：
//
//
//输入：s = "words and 987"
//输出：0
//解释：
//第 1 步："words and 987"（当前没有读入字符，因为没有前导空格）
//         ^
//第 2 步："words and 987"（当前没有读入字符，因为这里不存在 '-' 或者 '+'）
//         ^
//第 3 步："words and 987"（由于当前字符 'w' 不是一个数字，所以读入停止）
//         ^
//解析得到整数 0 ，因为没有读入任何数字。
//由于 0 在范围 [-2³¹, 2³¹ - 1] 内，最终结果为 0 。
//
// 示例 5：
//
//
//输入：s = "-91283472332"
//输出：-2147483648
//解释：
//第 1 步："-91283472332"（当前没有读入字符，因为没有前导空格）
//         ^
//第 2 步："-91283472332"（读入 '-' 字符，所以结果应该是负数）
//          ^
//第 3 步："-91283472332"（读入 "91283472332"）
//                     ^
//解析得到整数 -91283472332 。
//由于 -91283472332 小于范围 [-2³¹, 2³¹ - 1] 的下界，最终结果被截断为 -2³¹ = -2147483648 。
//
//
//
// 提示：
//
//
// 0 <= s.length <= 200
// s 由英文字母（大写和小写）、数字（0-9）、' '、'+'、'-' 和 '.' 组成
//
// Related Topics 字符串 👍 1286 👎 0

//leetcode submit region begin(Prohibit modification and deletion)
func myAtoi(s string) int {
	var (
		n        = len(s)
		ns       = &strings.Builder{}
		num      int64
		begin    bool
		zero     bool
		negative int64 = 1
	)

	for i := 0; i < n; i++ {
		if begin {
			if s[i] >= '0' && s[i] <= '9' {
				if (zero || ns.Len() == 0) && s[i] == '0' {

				} else {
					ns.Write([]byte{s[i]})
					zero = false
				}
			} else {
				break
			}
		} else {
			if s[i] == ' ' {
			} else if s[i] == '0' {
				zero = true
				begin = true
			} else if s[i] == '+' {
				begin = true
			} else if s[i] == '-' {
				negative = -1
				begin = true
			} else if s[i] > '0' && s[i] <= '9' {
				begin = true
				ns.Write([]byte{s[i]})
			} else {
				break
			}
		}
	}

	var (
		max int64 = 1<<31 - 1
		min int64 = -1 << 31
	)

	var nss, nssl = ns.String(), ns.Len()
	_ = nss
	if nssl > 0 && negative == -1 && nssl > len(fmt.Sprintf("%d", min))-1 {
		return int(min)
	} else if nssl > 0 && negative != -1 && nssl > len(fmt.Sprintf("%d", max)) {
		return int(max)
	}

	//var nssl = int64(ns.Len())
	for nss, i := ns.String(), 0; i < ns.Len(); i++ {
		//var a1 = nss[i] - '0'
		//var a2 = num * 10
		//var a3 = int64(a1) + a2
		//_, _, _ = a1, a2, a3
		num = int64(nss[i]-'0') + num*10
	}

	num *= negative

	if num > max {
		num = max
	} else if num < min {
		num = min
	}

	return int(num)
}

//leetcode submit region end(Prohibit modification and deletion)
