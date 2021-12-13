package main

import (
	"fmt"
	"math/bits"
)

func main() {
	fmt.Println(bits.OnesCount8(1))
	fmt.Println(bits.OnesCount8(2))
	fmt.Println(bits.OnesCount8(3))
	fmt.Println(bits.OnesCount8(4))
	fmt.Println(bits.OnesCount8(5))
	fmt.Println(bits.OnesCount8(6))
	fmt.Println(bits.OnesCount8(7))
	fmt.Println(bits.OnesCount8(8))
	fmt.Println(bits.OnesCount8(9))
	fmt.Println(bits.OnesCount8(10))
	fmt.Println(bits.OnesCount8(11))
	fmt.Println(bits.OnesCount8(12))
	fmt.Println(bits.OnesCount8(13))
	fmt.Println(bits.OnesCount8(14))
	fmt.Println(bits.OnesCount8(15))
	fmt.Println(bits.OnesCount8(16))
	fmt.Println(bits.OnesCount8(17))

}

//二进制手表顶部有 4 个 LED 代表 小时（0-11），底部的 6 个 LED 代表 分钟（0-59）。每个 LED 代表一个 0 或 1，最低位在右侧。
//
//
//
// 例如，下面的二进制手表读取 "3:25" 。
//
//
//
//
// （图源：WikiMedia - Binary clock samui moon.jpg ，许可协议：Attribution-ShareAlike 3.0
//Unported (CC BY-SA 3.0) ）
//
// 给你一个整数 turnedOn ，表示当前亮着的 LED 的数量，返回二进制手表可以表示的所有可能时间。你可以 按任意顺序 返回答案。
//
// 小时不会以零开头：
//
//
// 例如，"01:00" 是无效的时间，正确的写法应该是 "1:00" 。
//
//
// 分钟必须由两位数组成，可能会以零开头：
//
//
// 例如，"10:2" 是无效的时间，正确的写法应该是 "10:02" 。
//
//
//
//
// 示例 1：
//
//
//输入：turnedOn = 1
//输出：["0:01","0:02","0:04","0:08","0:16","0:32","1:00","2:00","4:00","8:00"]
//
//
// 示例 2：
//
//
//输入：turnedOn = 9
//输出：[]
//
//
//
//
// 提示：
//
//
// 0 <= turnedOn <= 10
//
// Related Topics 位运算 回溯 👍 338 👎 0

//leetcode submit region begin(Prohibit modification and deletion)
func readBinaryWatch(turnedOn int) []string {
	var s []string
	var i, j uint8
	for i = 0; i < 12; i++ {
		for j = 0; j < 60; j++ {
			if bits.OnesCount8(i)+bits.OnesCount8(j) == turnedOn {
				s = append(s, fmt.Sprintf("%d:%02d", i, j))
			}
		}
	}

	return s
}

//leetcode submit region end(Prohibit modification and deletion)
