package main

func main() {

}

//给定一个已按照 非递减顺序排列 的整数数组 numbers ，请你从数组中找出两个数满足相加之和等于目标数 target 。
//
// 函数应该以长度为 2 的整数数组的形式返回这两个数的下标值。numbers 的下标 从 1 开始计数 ，所以答案数组应当满足 1 <= answer[0]
// < answer[1] <= numbers.length 。
//
// 你可以假设每个输入 只对应唯一的答案 ，而且你 不可以 重复使用相同的元素。
//
//
// 示例 1：
//
//
//输入：numbers = [2,7,11,15], target = 9
//输出：[1,2]
//解释：2 与 7 之和等于目标数 9 。因此 index1 = 1, index2 = 2 。
//
//
// 示例 2：
//
//
//输入：numbers = [2,3,4], target = 6
//输出：[1,3]
//
//
// 示例 3：
//
//
//输入：numbers = [-1,0], target = -1
//输出：[1,2]
//
//
//
//
// 提示：
//
//
// 2 <= numbers.length <= 3 * 10⁴
// -1000 <= numbers[i] <= 1000
// numbers 按 非递减顺序 排列
// -1000 <= target <= 1000
// 仅存在一个有效答案
//
// Related Topics 数组 双指针 二分查找 👍 622 👎 0

//leetcode submit region begin(Prohibit modification and deletion)
func twoSum(numbers []int, target int) []int {
	var b, e = 0, len(numbers) - 1

	for b < e {
		var v = numbers[b] + numbers[e]
		if v == target {
			return []int{b + 1, e + 1}
		} else if v < target {
			b++
		} else {
			e--
		}
	}

	return nil
}

//leetcode submit region end(Prohibit modification and deletion)

func twoSum_II_A(numbers []int, target int) []int {
	var patch = make(map[int]int) // map[target-value]index
	for k, v := range numbers {
		if pi, has := patch[v]; has {
			return []int{pi + 1, k + 1}
		}
		patch[target-v] = k
	}

	return nil
}
