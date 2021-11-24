package main

import "fmt"

func main() {
	fmt.Println("A -> ", titleToNumber("A"))
	fmt.Println("Z -> ", titleToNumber("Z"))
	fmt.Println("AA -> ", titleToNumber("AA"))
	fmt.Println("AB -> ", titleToNumber("AB"))
	fmt.Println("AZ -> ", titleToNumber("AZ"))
	fmt.Println("BA -> ", titleToNumber("BA"))
	fmt.Println("ZY -> ", titleToNumber("ZY"))
	fmt.Println("ZZ -> ", titleToNumber("ZZ"))
	fmt.Println("AAA -> ", titleToNumber("AAA"))
	fmt.Println("AAB -> ", titleToNumber("AAB"))
	fmt.Println("FXSHRXW -> ", titleToNumber("FXSHRXW"))
}

//给你一个字符串 columnTitle ，表示 Excel 表格中的列名称。返回该列名称对应的列序号。
//
//
//
// 例如，
//
//
//    A -> 1
//    B -> 2
//    C -> 3
//    ...
//    Z -> 26
//    AA -> 27
//    AB -> 28
//    ...
//
//
//
//
// 示例 1:
//
//
//输入: columnTitle = "A"
//输出: 1
//
//
// 示例 2:
//
//
//输入: columnTitle = "AB"
//输出: 28
//
//
// 示例 3:
//
//
//输入: columnTitle = "ZY"
//输出: 701
//
// 示例 4:
//
//
//输入: columnTitle = "FXSHRXW"
//输出: 2147483647
//
//
//
//
// 提示：
//
//
// 1 <= columnTitle.length <= 7
// columnTitle 仅由大写英文组成
// columnTitle 在范围 ["A", "FXSHRXW"] 内
//
// Related Topics 数学 字符串 👍 295 👎 0

//leetcode submit region begin(Prohibit modification and deletion)
func titleToNumber(columnTitle string) int {
	var r int
	for i, multiple := len(columnTitle)-1, 1; i >= 0; i-- {
		var n = columnTitle[i] - 'A' + 1
		r += int(n) * multiple
		multiple *= 26
	}

	return r
}

func pow(num, t int) int {
	var r = 26
	for i := 0; i < t; i++ {
		if i == t-1 {
			r *= num
		} else {
			r *= 26
		}
	}
	return r
}

// A 1
// AA 27 1*26 + 1,
// ZY 701 26*26 + 26 = 676 + 25
// AAA 703 26*26*1 + 1*26 + 1

//leetcode submit region end(Prohibit modification and deletion)

func titleToNumber_B(columnTitle string) int {
	var l = len(columnTitle)
	if l == 1 {
		return int(columnTitle[0] - 64)
	}

	var i = titleToNumber(columnTitle[1:])

	return pow(int(columnTitle[0]-64), l-1) + i
}

func titleToNumber_A(columnTitle string) int {
	var (
		l     = len(columnTitle)
		toNum = func(i uint8) int {
			return int(i - 64)
		}
		pow = func(num, t int) int {
			var r = 26
			for i := 0; i < t; i++ {
				if i == t-1 {
					r *= num
				} else {
					r *= 26
				}
			}
			return r
		}
		r int
	)

	for i := 0; i < l; i++ {
		var n = toNum(columnTitle[i])
		//fmt.Printf("%d-%d-%d \n", l-i-1, pow(26, l-i-1), n)
		if i == l-1 {
			r += n
		} else {
			r += pow(n, l-i-1)
		}
	}

	return r
}
