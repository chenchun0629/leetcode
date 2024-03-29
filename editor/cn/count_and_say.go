package main

import (
	"fmt"
	"strings"
)

func main() {
	//fmt.Println('0')
	//fmt.Println(string(byte(49)))
	fmt.Println(countAndSay(4))
}

//解释下。
//
//1
//11
//21
//1211
//111221
//一步一步来
//
//给一个数，这个数是1
//描述上一步的数，这个数是 1 即一个1，故写作11
//描述上一步的数，这个数是11即两个1，故写作21
//描述上一步的数，这个数是21即一个2一个1，故写作12-11
//描述上一步的数，这个数是1211即一个1一个2两个1，故写作11-12-21

//给定一个正整数 n ，输出外观数列的第 n 项。
//
// 「外观数列」是一个整数序列，从数字 1 开始，序列中的每一项都是对前一项的描述。
//
// 你可以将其视作是由递归公式定义的数字字符串序列：
//
//
// countAndSay(1) = "1"
// countAndSay(n) 是对 countAndSay(n-1) 的描述，然后转换成另一个数字字符串。
//
//
// 前五项如下：
//
//
//1.     1
//2.     11
//3.     21
//4.     1211
//5.     111221
//第一项是数字 1
//描述前一项，这个数是 1 即 “ 一 个 1 ”，记作 "11"
//描述前一项，这个数是 11 即 “ 二 个 1 ” ，记作 "21"
//描述前一项，这个数是 21 即 “ 一 个 2 + 一 个 1 ” ，记作 "1211"
//描述前一项，这个数是 1211 即 “ 一 个 1 + 一 个 2 + 二 个 1 ” ，记作 "111221"
//
//
// 要 描述 一个数字字符串，首先要将字符串分割为 最小 数量的组，每个组都由连续的最多 相同字符 组成。然后对于每个组，先描述字符的数量，然后描述字符，形成
//一个描述组。要将描述转换为数字字符串，先将每组中的字符数量用数字替换，再将所有描述组连接起来。
//
// 例如，数字字符串 "3322251" 的描述如下图：
//
//
//
//
//
//
// 示例 1：
//
//
//输入：n = 1
//输出："1"
//解释：这是一个基本样例。
//
//
// 示例 2：
//
//
//输入：n = 4
//输出："1211"
//解释：
//countAndSay(1) = "1"
//countAndSay(2) = 读 "1" = 一 个 1 = "11"
//countAndSay(3) = 读 "11" = 二 个 1 = "21"
//countAndSay(4) = 读 "21" = 一 个 2 + 一 个 1 = "12" + "11" = "1211"
//
//
//
//
// 提示：
//
//
// 1 <= n <= 30
//
// Related Topics 字符串 👍 850 👎 0

//leetcode submit region begin(Prohibit modification and deletion)
func countAndSay(n int) string {
	var (
		s  = "1"
		fn = func(s string) string {
			var (
				n  = len(s)
				ns = strings.Builder{}
			)

			var cnt, char = byte(0), byte(0)
			for i := 0; i < n; i++ {
				if char != s[i] {
					if char != 0 {
						ns.Write([]byte{cnt + '0', char})
					}

					char = s[i]
					cnt = 0
				}

				cnt++
			}

			if cnt > 0 {
				ns.Write([]byte{cnt + '0', char})
			}

			return ns.String()
		}
	)

	for i := 1; i < n; i++ {
		s = fn(s)
	}

	return s
}

//leetcode submit region end(Prohibit modification and deletion)
