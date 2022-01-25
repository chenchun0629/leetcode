package main

import "fmt"

func main() {
	//fmt.Println(peakIndexInMountainArray([]int{3, 2, 1, 0}))
	fmt.Println(peakIndexInMountainArray([]int{0, 1, 2, 3}))
	//fmt.Println(peakIndexInMountainArray([]int{24, 69, 100, 99, 79, 78, 67, 36, 26, 19}))
}

//符合下列属性的数组 arr 称为 山脉数组 ：
//
// arr.length >= 3
// 存在 i（0 < i < arr.length - 1）使得：
//
// arr[0] < arr[1] < ... arr[i-1] < arr[i]
// arr[i] > arr[i+1] > ... > arr[arr.length - 1]
//
//
//
//
// 给你由整数组成的山脉数组 arr ，返回任何满足 arr[0] < arr[1] < ... arr[i - 1] < arr[i] > arr[i +
//1] > ... > arr[arr.length - 1] 的下标 i 。
//
//
//
// 示例 1：
//
//
//输入：arr = [0,1,0]
//输出：1
//
//
// 示例 2：
//
//
//输入：arr = [0,2,1,0]
//输出：1
//
//
// 示例 3：
//
//
//输入：arr = [0,10,5,2]
//输出：1
//
//
// 示例 4：
//
//
//输入：arr = [3,4,5,1]
//输出：2
//
//
// 示例 5：
//
//
//输入：arr = [24,69,100,99,79,78,67,36,26,19]
//输出：2
//
//
//
//
// 提示：
//
//
// 3 <= arr.length <= 10⁴
// 0 <= arr[i] <= 10⁶
// 题目数据保证 arr 是一个山脉数组
//
//
//
//
// 进阶：很容易想到时间复杂度 O(n) 的解决方案，你可以设计一个 O(log(n)) 的解决方案吗？
// Related Topics 数组 二分查找 👍 225 👎 0

//leetcode submit region begin(Prohibit modification and deletion)
func peakIndexInMountainArray(arr []int) int {
	var (
		n          = len(arr)
		head, tail = 0, n - 1
		mid        int
	)

	for head < tail {
		mid = head + (tail-head)>>1

		if mid == 0 {
			if arr[mid] > arr[mid+1] {
				return mid
			}
		} else if mid == n-1 {
			if arr[mid-1] < arr[mid] {
				return mid
			}
		} else {
			if arr[mid-1] < arr[mid] && arr[mid] > arr[mid+1] {
				return mid
			}
		}

		if arr[mid-1] < arr[mid] {
			head = mid + 1
		} else {
			tail = mid
		}
	}

	return head
}

//leetcode submit region end(Prohibit modification and deletion)
