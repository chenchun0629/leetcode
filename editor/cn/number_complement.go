package main

func main() {
	findComplement(1)
	findComplement(2)
	findComplement(3)
	findComplement(4)
	findComplement(5)

	//fmt.Printf("%b, \n", 5)
	//fmt.Printf("%b, \n", 7)
	//fmt.Printf("%b, %d, \n", 5^7, 5^7)
	//fmt.Println(findComplement(1))
	//fmt.Println(findComplement(2))
	//fmt.Println(findComplement(3))
	//fmt.Println(findComplement(4))
	//fmt.Println(findComplement(5))
	//fmt.Println(findComplement(7))
}

//对整数的二进制表示取反（0 变 1 ，1 变 0）后，再转换为十进制表示，可以得到这个整数的补数。
//
//
// 例如，整数 5 的二进制表示是 "101" ，取反后得到 "010" ，再转回十进制表示得到补数 2 。
//
//
// 给你一个整数 num ，输出它的补数。
//
//
//
//
//
//
// 示例 1：
//
//
//输入：num = 5
//输出：2
//解释：5 的二进制表示为 101（没有前导零位），其补数为 010。所以你需要输出 2 。
//
//
// 示例 2：
//
//
//输入：num = 1
//输出：0
//解释：1 的二进制表示为 1（没有前导零位），其补数为 0。所以你需要输出 0 。
//
//
//
//
// 提示：
//
//
// 1 <= num < 2³¹
//
//
//
//
// 注意：本题与 1009 https://leetcode-cn.com/problems/complement-of-base-10-integer/ 相
//同
// Related Topics 位运算 👍 284 👎 0

//leetcode submit region begin(Prohibit modification and deletion)

func findComplement(num int) int {
	// 101 ^ 111 = 010 使用异或操作取反

	var (
		mask = 1
		hb   int
	)
	for hb = 0; mask <= num; hb++ {
		mask = 1 << hb
	}

	//fmt.Printf("%d, %d, %d, %d \n", num, mask-1, hb-1, num ^ (mask - 1))
	//fmt.Printf("%b, %b, %b \n", num, mask-1, num^(mask-1))

	return num ^ (mask - 1)
}

//leetcode submit region end(Prohibit modification and deletion)
