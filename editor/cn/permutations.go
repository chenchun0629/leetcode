package main

import "fmt"

func main() {

	fmt.Println(permute([]int{1, 2, 3, 4}))

	// 1234
	// 1243

	// 1324
	// 1342

	// 1423
	// 1432

	//1234
	//1324
	//1342
	//1243
	//1423
	//1432

	//[1 2 3 4] [1 3 2 4] [1 3 4 2] [1 3 2 4] [1 2 3 4]
	//
	//2134
	//2143
	//2314
	//2341
	//2413
	//2431
}

//给定一个不含重复数字的数组 nums ，返回其 所有可能的全排列 。你可以 按任意顺序 返回答案。
//
//
//
// 示例 1：
//
//
//输入：nums = [1,2,3]
//输出：[[1,2,3],[1,3,2],[2,1,3],[2,3,1],[3,1,2],[3,2,1]]
//
//
// 示例 2：
//
//
//输入：nums = [0,1]
//输出：[[0,1],[1,0]]
//
//
// 示例 3：
//
//
//输入：nums = [1]
//输出：[[1]]
//
//
//
//
// 提示：
//
//
// 1 <= nums.length <= 6
// -10 <= nums[i] <= 10
// nums 中的所有整数 互不相同
//
// Related Topics 数组 回溯 👍 1708 👎 0

//leetcode submit region begin(Prohibit modification and deletion)
func permute(nums []int) [][]int {
	var (
		ret [][]int
		n   = len(nums)
		fn  func(prefix []int, visited []int)
	)

	fn = func(prefix []int, visited []int) {
		if len(prefix) == n {
			ret = append(ret, prefix)
		}
		for i := 0; i < n; i++ {
			if visited[i] == 1 {
				continue
			}

			var tmp = append([]int{nums[i]}, prefix...)
			visited[i] = 1
			fn(tmp, visited)
			visited[i] = 0
		}
	}

	fn([]int{}, make([]int, n))

	return ret
}

//leetcode submit region end(Prohibit modification and deletion)
