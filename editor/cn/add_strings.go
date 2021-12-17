package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(addStrings("9", "99"))
}

//给定两个字符串形式的非负整数 num1 和num2 ，计算它们的和并同样以字符串形式返回。
//
// 你不能使用任何內建的用于处理大整数的库（比如 BigInteger）， 也不能直接将输入的字符串转换为整数形式。
//
//
//
// 示例 1：
//
//
//输入：num1 = "11", num2 = "123"
//输出："134"
//
//
// 示例 2：
//
//
//输入：num1 = "456", num2 = "77"
//输出："533"
//
//
// 示例 3：
//
//
//输入：num1 = "0", num2 = "0"
//输出："0"
//
//
//
//
//
//
// 提示：
//
//
// 1 <= num1.length, num2.length <= 10⁴
// num1 和num2 都只包含数字 0-9
// num1 和num2 都不包含任何前导零
//
// Related Topics 数学 字符串 模拟 👍 483 👎 0

//leetcode submit region begin(Prohibit modification and deletion)
func addStrings(num1 string, num2 string) string {
	var ptr1, ptr2 = len(num1) - 1, len(num2) - 1
	if ptr1 < ptr2 {
		ptr1, ptr2 = ptr2, ptr1
		num1, num2 = num2, num1
	}

	var (
		bu   = &strings.Builder{}
		bs   = make([]byte, 0)
		plus byte
	)
	for ptr1 >= 0 {
		var t = num1[ptr1] - '0' + plus
		if ptr2 >= 0 {
			t = t + num2[ptr2] - '0'
			ptr2--
		}
		ptr1--

		var tp = t / 10
		t = (t) % 10
		plus = tp
		bs = append(bs, t+'0')
	}

	if plus > 0 {
		bs = append(bs, plus+'0')
	}

	for i := len(bs) - 1; i >= 0; i-- {
		bu.WriteByte(bs[i])
	}

	return bu.String()
}

//leetcode submit region end(Prohibit modification and deletion)
