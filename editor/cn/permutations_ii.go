package main

import (
	"fmt"
	"runtime"
	"sort"
)

func main() {
	var ms_b runtime.MemStats
	var ms_e runtime.MemStats
	runtime.ReadMemStats(&ms_b)
	//fmt.Println(permuteUnique([]int{3, 3, 0, 3}))
	fmt.Println(permuteUnique([]int{1, 2, 3}))
	runtime.ReadMemStats(&ms_e)
	fmt.Printf("%d \n", ms_e.Alloc-ms_b.Alloc)
	//==== %+v[]
	//==== %+v[1]
	//==== %+v[2 1]
	//==== %+v[3 2 1]
	//==== %+v[3 1]
	//==== %+v[2 3 1]
	//==== %+v[2]
	//==== %+v[1 2]
	//==== %+v[3 1 2]
	//==== %+v[3 2]
	//==== %+v[1 3 2]
	//==== %+v[3]
	//==== %+v[1 3]
	//==== %+v[2 1 3]
	//==== %+v[2 3]
	//==== %+v[1 2 3]
}

//给定一个可包含重复数字的序列 nums ，按任意顺序 返回所有不重复的全排列。
//
//
//
// 示例 1：
//
//
//输入：nums = [1,1,2]
//输出：
//[[1,1,2],
// [1,2,1],
// [2,1,1]]
//
//
// 示例 2：
//
//
//输入：nums = [1,2,3]
//输出：[[1,2,3],[1,3,2],[2,1,3],[2,3,1],[3,1,2],[3,2,1]]
//
//
//
//
// 提示：
//
//
// 1 <= nums.length <= 8
// -10 <= nums[i] <= 10
//
// Related Topics 数组 回溯 👍 899 👎 0

//leetcode submit region begin(Prohibit modification and deletion)
func permuteUnique(nums []int) [][]int {
	var (
		ret     [][]int
		n       = len(nums)
		fn      func()
		visited = make([]bool, n)
		prefix  = make([]int, 0)
	)
	sort.Ints(nums)

	fn = func() {
		if len(prefix) == n {
			ret = append(ret, append([]int(nil), prefix...))
			return
		}
		for i := 0; i < n; i++ {
			if visited[i] || i > 0 && visited[i-1] && nums[i] == nums[i-1] {
				continue
			}

			prefix = append(prefix, nums[i])
			visited[i] = true
			fn()
			visited[i] = false
			prefix = prefix[:len(prefix)-1]
		}
	}

	fn()

	return ret
}

//leetcode submit region end(Prohibit modification and deletion)

// 慢 和 内存用的多
func permuteUnique_A(nums []int) [][]int {
	var (
		ret [][]int
		n   = len(nums)
		fn  func(prefix []int, visited []int)
		m   = make(map[string]struct{})
	)

	fn = func(prefix []int, visited []int) {
		var key = fmt.Sprint("%+v", prefix)
		if _, has := m[key]; has {
			return
		}
		m[key] = struct{}{}

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
