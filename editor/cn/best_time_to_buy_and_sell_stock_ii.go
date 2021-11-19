package main

import "fmt"

func main() {
	var prices = []int{7, 1, 5, 3, 6, 4}
	fmt.Println(maxProfit(prices))
}

//给定一个数组 prices ，其中 prices[i] 是一支给定股票第 i 天的价格。
//
// 设计一个算法来计算你所能获取的最大利润。你可以尽可能地完成更多的交易（多次买卖一支股票）。
//
// 注意：你不能同时参与多笔交易（你必须在再次购买前出售掉之前的股票）。
//
//
//
// 示例 1:
//
//
//输入: prices = [7,1,5,3,6,4]
//输出: 7
//解释: 在第 2 天（股票价格 = 1）的时候买入，在第 3 天（股票价格 = 5）的时候卖出, 这笔交易所能获得利润 = 5-1 = 4 。
//     随后，在第 4 天（股票价格 = 3）的时候买入，在第 5 天（股票价格 = 6）的时候卖出, 这笔交易所能获得利润 = 6-3 = 3 。
//
//
// 示例 2:
//
//
//输入: prices = [1,2,3,4,5]
//输出: 4
//解释: 在第 1 天（股票价格 = 1）的时候买入，在第 5 天 （股票价格 = 5）的时候卖出, 这笔交易所能获得利润 = 5-1 = 4 。
//     注意你不能在第 1 天和第 2 天接连购买股票，之后再将它们卖出。因为这样属于同时参与了多笔交易，你必须在再次购买前出售掉之前的股票。
//
//
// 示例 3:
//
//
//输入: prices = [7,6,4,3,1]
//输出: 0
//解释: 在这种情况下, 没有交易完成, 所以最大利润为 0。
//
//
//
// 提示：
//
//
// 1 <= prices.length <= 3 * 10⁴
// 0 <= prices[i] <= 10⁴
//
// Related Topics 贪心 数组 动态规划 👍 1454 👎 0

//leetcode submit region begin(Prohibit modification and deletion)
func maxProfit(prices []int) int {
	var l = len(prices)

	if l <= 1 {
		return 0
	}

	// 只需要关注前一天，所以无需记录所有情况，去掉table历史记录，只记录前一天就就可以了
	var (
		dp0 = 0
		dp1 = -prices[0] // 第一天买入股票，利润负数。
	)

	for i := 1; i < l; i++ {
		// 前一天没有股票的利润 > 今天买出后的利润 ？ 前一天的利润 ：今天买出后的利润  ->  有利润则卖出，反之废弃昨日买出的决定
		// 相当于今天卖出没有利润，则废弃昨日买入的决定
		dp0 = maxProfitMax_II(dp0, dp1+prices[i])
		// 前一天买入后的利润 > 今天买入后的利润 ？ 前一天买入后的利润 : 当天买入股票后的利润
		// 今日买入后的利润 大于 昨日的利润，则 废弃昨日买入的决定
		dp1 = maxProfitMax_II(dp1, dp0-prices[i])
	}

	//fmt.Printf("%#v \n", table)

	return dp0
}

func maxProfitMax_II(a, b int) int {
	if a > b {
		return a
	}
	return b
}

//leetcode submit region end(Prohibit modification and deletion)

// maxProfit_II_B 动态规划
func maxProfit_II_B(prices []int) int {
	var l = len(prices)

	if l <= 1 {
		return 0
	}

	var table = make([][2]int, l) // [2]int[0] 当天没有股票的利润 [2]int[1] 当天买入股票后的利润
	table[0][1] = -prices[0]      // 第一天买入股票，利润负数。
	for i := 1; i < l; i++ {
		// 前一天没有股票的利润 > 今天买出后的利润 ？ 前一天的利润 ：今天买出后的利润  ->  有利润则卖出，反之废弃昨日买出的决定
		// 相当于今天卖出没有利润，则废弃昨日买入的决定
		table[i][0] = maxProfitMax_II(table[i-1][0], table[i-1][1]+prices[i])
		// 前一天买入后的利润 > 今天买入后的利润 ？ 前一天买入后的利润 : 当天买入股票后的利润
		// 今日买入后的利润 大于 昨日的利润，则 废弃昨日买入的决定
		table[i][1] = maxProfitMax_II(table[i-1][1], table[i-1][0]-prices[i])
	}

	//fmt.Printf("%#v \n", table)

	return table[l-1][0]
}

// maxProfit_II_A 贪心
func maxProfit_II_A(prices []int) int {
	var l = len(prices)

	if l <= 1 {
		return 0
	}

	var profit = 0
	for i := 0; i < l-1; i++ {
		if prices[i] < prices[i+1] {
			profit += (prices[i+1] - prices[i])
		}
	}

	return profit
}
