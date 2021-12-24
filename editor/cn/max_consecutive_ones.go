package main

func main() {

}

//给定一个二进制数组， 计算其中最大连续 1 的个数。
//
//
//
// 示例：
//
//
//输入：[1,1,0,1,1,1]
//输出：3
//解释：开头的两位和最后的三位都是连续 1 ，所以最大连续 1 的个数是 3.
//
//
//
//
// 提示：
//
//
// 输入的数组只包含 0 和 1 。
// 输入数组的长度是正整数，且不超过 10,000。
//
// Related Topics 数组 👍 284 👎 0

//leetcode submit region begin(Prohibit modification and deletion)
func findMaxConsecutiveOnes(nums []int) int {
	var max = 0
	var cnt = 0
	for _, v := range nums {
		if v == 1 {
			cnt++
		} else {
			if cnt > max {
				max = cnt
			}
			cnt = 0
		}
	}
	if cnt > max {
		max = cnt
	}

	return max
}

//leetcode submit region end(Prohibit modification and deletion)

// 内存消耗多
func findMaxConsecutiveOnes_A(nums []int) int {
	var max, begin, end = 0, -1, -1
	for k, v := range nums {
		if v == 1 {
			if begin == -1 {
				begin = k
			}
			end = k
			if end-begin+1 > max {
				max = end - begin + 1
			}
		} else {
			begin = -1
			end = -1
		}
	}
	return max
}
