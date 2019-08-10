package tests

import (
	"fmt"
	"math"
	"testing"
)

/**
 * [123] Best Time to Buy and Sell Stock III
 *
 * Say you have an array for which the i^th element is the price of a given stock on day i.
 *
 * Design an algorithm to find the maximum profit. You may complete at most two transactions.
 *
 * Note: You may not engage in multiple transactions at the same time (i.e., you must sell the stock before you buy again).
 *
 * Example 1:
 *
 *
 * Input: [3,3,5,0,0,3,1,4]
 * Output: 6
 * Explanation: Buy on day 4 (price = 0) and sell on day 6 (price = 3), profit = 3-0 = 3.
 *              Then buy on day 7 (price = 1) and sell on day 8 (price = 4), profit = 4-1 = 3.
 *
 * Example 2:
 *
 *
 * Input: [1,2,3,4,5]
 * Output: 4
 * Explanation: Buy on day 1 (price = 1) and sell on day 5 (price = 5), profit = 5-1 = 4.
 *              Note that you cannot buy on day 1, buy on day 2 and sell them later, as you are
 *              engaging multiple transactions at the same time. You must sell before buying again.
 *
 *
 * Example 3:
 *
 *
 * Input: [7,6,4,3,1]
 * Output: 0
 * Explanation: In this case, no transaction is done, i.e. max profit = 0.
 *
 */

func TestBestTimetoBuyandSellStockIII(t *testing.T) {
	var cases = []struct {
		input  []int
		output int
	}{
		{
			input:  []int{3, 3, 5, 0, 0, 3, 1, 4},
			output: 6,
		},
		{
			input:  []int{1, 2, 3, 4, 5},
			output: 4,
		},
		{
			input:  []int{7, 6, 4, 3, 1},
			output: 0,
		},
	}
	for _, c := range cases {
		x := maxProfitIII(c.input)
		if x != c.output {
			t.Fail()
		}
	}
}

// submission codes start here

func maxProfitIII(prices []int) int {
	oneBuy := math.MinInt64
	oneSell := 0
	twoBuy := math.MinInt64
	twoSell := 0
	for _, v := range prices {
		//we set prices to negative, so the calculation of profit will be convenient
		oneBuy = int(math.Max(float64(oneBuy), float64(-v)))
		oneSell = int(math.Max(float64(oneSell), float64(oneBuy + v)))
		//we can buy the second only after first is sold
		twoBuy = int(math.Max(float64(twoBuy), float64(oneSell-v)))
		twoSell = int(math.Max(float64(twoSell), float64(twoBuy + v)))
		fmt.Println(oneBuy, oneSell, twoBuy, twoSell)
	}
	return int(math.Max(float64(twoSell), float64(oneSell)))
}

// submission codes end
