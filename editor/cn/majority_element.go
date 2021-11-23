package main

func main() {
	majorityElement([]int{1, 2, 3, 2, 3, 2})
}

//给定一个大小为 n 的数组，找到其中的多数元素。多数元素是指在数组中出现次数 大于 ⌊ n/2 ⌋ 的元素。
//
// 你可以假设数组是非空的，并且给定的数组总是存在多数元素。
//
//
//
// 示例 1：
//
//
//输入：[3,2,3]
//输出：3
//
// 示例 2：
//
//
//输入：[2,2,1,1,1,2,2]
//输出：2
//
//
//
//
// 进阶：
//
//
// 尝试设计时间复杂度为 O(n)、空间复杂度为 O(1) 的算法解决此问题。
//
// Related Topics 数组 哈希表 分治 计数 排序 👍 1204 👎 0

//leetcode submit region begin(Prohibit modification and deletion)
func majorityElement(nums []int) int {
	//Boyer-Moore 算法的本质和方法四中的分治十分类似。我们首先给出 Boyer-Moore 算法的详细步骤：
	//
	//- 我们维护一个候选众数 `candidate` 和它出现的次数 `count`。初始时 `candidate` 可以为任意值，`count` 为 `0`；
	//
	//- 我们遍历数组 `nums` 中的所有元素，对于每个元素 `x`，在判断 `x` 之前，如果 `count` 的值为 `0`，我们先将 `x` 的值赋予 `candidate`，随后我们判断 `x`：
	//
	//- 如果 `x` 与 `candidate` 相等，那么计数器 `count` 的值增加 `1`；
	//
	//- 如果 `x` 与 `candidate` 不等，那么计数器 `count` 的值减少 `1`。
	//- 在遍历完成后，`candidate` 即为整个数组的众数。

	var cnt, candidate = 0, 0

	for _, num := range nums {
		if cnt == 0 {
			candidate = num
		}

		if num == candidate {
			cnt += 1
		} else {
			cnt -= 1
		}
	}

	return candidate
}

//leetcode submit region end(Prohibit modification and deletion)

func majorityElement_A(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	var m = make(map[int]int)

	for _, v := range nums {
		m[v]++
	}

	var maxNum = 0
	var maxCnt = 0
	for num, cnt := range m {
		if cnt > maxCnt {
			maxNum = num
			maxCnt = cnt
		}
	}

	return maxNum
}
