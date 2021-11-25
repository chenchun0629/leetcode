package main

import (
	"fmt"
	"strconv"
)

func main() {
	//fmt.Printf("%032s", "10100101000001111010011100")
	fmt.Println(reverseBits(0b00000010100101000001111010011100))
	//var s = fmt.Sprintf("%b", 10)
	//fmt.Println(reflect.TypeOf(s).Name())
	//var i, err = strconv.ParseInt(s, 2, 64)
	//fmt.Println(i, err)
}

//颠倒给定的 32 位无符号整数的二进制位。
//
// 提示：
//
//
// 请注意，在某些语言（如 Java）中，没有无符号整数类型。在这种情况下，输入和输出都将被指定为有符号整数类型，并且不应影响您的实现，因为无论整数是有符号的
//还是无符号的，其内部的二进制表示形式都是相同的。
// 在 Java 中，编译器使用二进制补码记法来表示有符号整数。因此，在 示例 2 中，输入表示有符号整数 -3，输出表示有符号整数 -1073741825。
//
//
//
//
//
// 示例 1：
//
//
//输入：n = 00000010100101000001111010011100
//输出：964176192 (00111001011110000010100101000000)
//解释：输入的二进制串 00000010100101000001111010011100 表示无符号整数 43261596，
//     因此返回 964176192，其二进制表示形式为 00111001011110000010100101000000。
//
// 示例 2：
//
//
//输入：n = 11111111111111111111111111111101
//输出：3221225471 (10111111111111111111111111111111)
//解释：输入的二进制串 11111111111111111111111111111101 表示无符号整数 4294967293，
//     因此返回 3221225471 其二进制表示形式为 10111111111111111111111111111111 。
//
//
//
// 提示：
//
//
// 输入是一个长度为 32 的二进制字符串
//
//
//
//
// 进阶: 如果多次调用这个函数，你将如何优化你的算法？
// Related Topics 位运算 分治 👍 459 👎 0

//leetcode submit region begin(Prohibit modification and deletion)
// reverseBits 方法B 通过位运算
func reverseBits(num uint32) uint32 {
	num = ((num & 0xAAAAAAAA) >> 1) | ((num & 0x55555555) << 1)
	num = ((num & 0xCCCCCCCC) >> 2) | ((num & 0x33333333) << 2)
	num = ((num & 0xF0F0F0F0) >> 4) | ((num & 0x0F0F0F0F) << 4)
	num = ((num & 0xFF00FF00) >> 8) | ((num & 0x00FF00FF) << 8)
	num = ((num) >> 16) | ((num) << 16)

	return num
}

//leetcode submit region end(Prohibit modification and deletion)

// reverseBits_A 方法A 转字符串首位交换
func reverseBits_A(num uint32) uint32 {
	//fmt.Println(fmt.Sprintf("%b", num), "================a")
	var s = []rune(fmt.Sprintf("%032s", fmt.Sprintf("%b", num)))
	var b, e = 0, len(s) - 1
	for b < e {
		s[b], s[e] = s[e], s[b]
		b++
		e--
	}

	var i, _ = strconv.ParseInt(string(s), 2, 64)
	//fmt.Println(string(s), i, err, "================b")
	return uint32(i)
}
