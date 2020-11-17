package main

import (
	"fmt"
	"strings"
)

//罗马数字包含以下七种字符: I， V， X， L，C，D 和 M。
//
// 字符          数值
//I             1
//V             5
//X             10
//L             50
//C             100
//D             500
//M             1000
//
// 例如， 罗马数字 2 写做 II ，即为两个并列的 1。12 写做 XII ，即为 X + II 。 27 写做 XXVII, 即为 XX + V + I
//I 。
//
// 通常情况下，罗马数字中小的数字在大的数字的右边。但也存在特例，例如 4 不写做 IIII，而是 IV。数字 1 在数字 5 的左边，所表示的数等于大数 5
// 减小数 1 得到的数值 4 。同样地，数字 9 表示为 IX。这个特殊的规则只适用于以下六种情况：
//
//
// I 可以放在 V (5) 和 X (10) 的左边，来表示 4 和 9。
// X 可以放在 L (50) 和 C (100) 的左边，来表示 40 和 90。
// C 可以放在 D (500) 和 M (1000) 的左边，来表示 400 和 900。
//
//
// 给定一个罗马数字，将其转换成整数。输入确保在 1 到 3999 的范围内。
//
//
//
// 示例 1:
//
// 输入: "III"
//输出: 3
//
// 示例 2:
//
// 输入: "IV"
//输出: 4
//
// 示例 3:
//
// 输入: "IX"
//输出: 9
//
// 示例 4:
//
// 输入: "LVIII"
//输出: 58
//解释: L = 50, V= 5, III = 3.
//
//
// 示例 5:
//
// 输入: "MCMXCIV"
//输出: 1994
//解释: M = 1000, CM = 900, XC = 90, IV = 4.
//
//
//
// 提示：
//
//
// 题目所给测试用例皆符合罗马数字书写规则，不会出现跨位等情况。
// IC 和 IM 这样的例子并不符合题目要求，49 应该写作 XLIX，999 应该写作 CMXCIX 。
// 关于罗马数字的详尽书写规则，可以参考 罗马数字 - Mathematics 。
//
// Related Topics 数学 字符串
// 👍 1119 👎 0

//leetcode submit region begin(Prohibit modification and deletion)
func romanToInt(s string) int {
	return romanToIntB(s)
}

// 方案B
// 我们可以找到新歌规则
// 前一个数大于后一个数时，我们可以剪掉后一个数
// 前一个数小于大于等于后一个数时，我们可以加上后一个数
// 例如：MCMXCIV => M > C < M > X < C > I < V
//                1000 - 100 + 1000 - 10 + 100 - 1 + 5 = 1994
func romanToIntB(s string) int {
	var (
		r    = []rune(s)
		l    = len(r)
		ret  int
		prev = m[string(r[0])]
	)

	for i := 1; i < l; i++ {
		var cur = m[string(r[i])]
		if prev >= cur {
			ret += prev
		} else {
			ret -= prev
		}
		prev = cur
	}

	ret += prev

	return ret
}

// 方案A
// 根据规则，可以分为三个步骤
// 1、循环遍历字符串， 确定一个span 如 X、IV、IXL 等
// 2、计算span结果， x => 10, IV => 4, IXL => 39
// 3、将所有span加在一起得到最后结果
// 例如： MCMXCIV 分为 M、 CM、 XC、 IV 4个span
// 最后结果为： 1000 + (-100 + 1000) + (-10 + 1000) + (-1 + 5) = 1000 + 900 + 90 + 4 = 1994
func romanToIntA(s string) int {
	var (
		r   = []rune(s)
		l   = len(r)
		ret int
	)

	for i := 0; i < l; i++ {
		var (
			t = []rune{r[i]}
			c int
		)
		for i+c+1 < l {
			if strings.Contains(n[string(r[i])], string(r[i+c+1])) {
				t = append(t, r[i+c+1])
				c++
			} else {
				break
			}
		}
		i = i + c

		var lt = len(t)
		for j := 0; j < lt; j++ {
			if j+1 == lt {
				ret += m[string(t[j])]
			} else {
				ret -= m[string(t[j])]
			}
		}
	}

	return ret
}

var (
	m = map[string]int{
		"I": 1,
		"V": 5,
		"X": 10,
		"L": 50,
		"C": 100,
		"D": 500,
		"M": 1000,
	}
	n = map[string]string{
		"I": "VX",
		"V": "XL",
		"X": "LC",
		"L": "CD",
		"C": "DM",
		"D": "M",
	}
)

//leetcode submit region end(Prohibit modification and deletion)

func main() {
	fmt.Println(romanToInt("MCMXCIV"))
}
