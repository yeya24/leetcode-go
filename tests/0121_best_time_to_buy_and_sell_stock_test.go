package tests

import (
	"math"
	"testing"
)

/**
 * [121] Best Time to Buy and Sell Stock
 *
 * Say you have an array for which the i^th element is the price of a given stock on day i.
 *
 * If you were only permitted to complete at most one transaction (i.e., buy one and sell one share of the stock), design an algorithm to find the maximum profit.
 *
 * Note that you cannot sell a stock before you buy one.
 *
 * Example 1:
 *
 *
 * Input: [7,1,5,3,6,4]
 * Output: 5
 * Explanation: Buy on day 2 (price = 1) and sell on day 5 (price = 6), profit = 6-1 = 5.
 *              Not 7-1 = 6, as selling price needs to be larger than buying price.
 *
 *
 * Example 2:
 *
 *
 * Input: [7,6,4,3,1]
 * Output: 0
 * Explanation: In this case, no transaction is done, i.e. max profit = 0.
 *
 *
 */

func TestBestTimetoBuyandSellStock(t *testing.T) {
	var cases = []struct {
		input  []int
		output int
	}{
		{
			input:  []int{7, 1, 5, 3, 6, 4},
			output: 5,
		},
		{
			input:  []int{7, 6, 4, 3, 1},
			output: 0,
		},
		{
			input:  []int{},
			output: 0,
		},
	}
	for _, c := range cases {
		x := maxProfit(c.input)
		if x != c.output {
			t.Fail()
		}
	}
}

// submission codes start here

// reduces to max subarray
func maxProfit(prices []int) int {
	if len(prices) == 0 {
		return 0
	}
	l := make([]int, len(prices)-1)
	for i := 0; i < len(prices)-1; i++ {
		l[i] = prices[i+1] - prices[i]
	}
	max := 0
	sum := 0
	for _, v := range l {
		if sum < 0 {
			sum = v
		} else {
			sum += v
		}
		if sum > max {
			max = sum
		}
	}
	return max
}

func maxProfit2(prices []int) int {
	n := len(prices)
	if n == 0 {
		return 0
	}
	minPrice := make([]int, n)
	maxProfit := make([]int, n)
	minPrice[0] = prices[0]
	maxProfit[0] = 0
	for i := 1; i < n; i++ {
		minPrice[i] = int(math.Min(float64(prices[i]), float64(minPrice[i-1])))
		maxProfit[i] = int(math.Max(float64(prices[i]-minPrice[i-1]), float64(maxProfit[i-1])))
	}
	return maxProfit[n-1]
}

// submission codes end
