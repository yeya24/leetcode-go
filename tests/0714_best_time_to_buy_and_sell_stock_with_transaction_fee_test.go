package tests

import (
	"math"
	"testing"
)

/**
 * [714] Best Time to Buy and Sell Stock with Transaction Fee
 *
 * Your are given an array of integers prices, for which the i-th element is the price of a given stock on day i; and a non-negative integer fee representing a transaction fee.
 * You may complete as many transactions as you like, but you need to pay the transaction fee for each transaction.  You may not buy more than 1 share of a stock at a time (ie. you must sell the stock share before you buy again.)
 * Return the maximum profit you can make.
 *
 * Example 1:
 *
 * Input: prices = [1, 3, 2, 8, 4, 9], fee = 2
 * Output: 8
 * Explanation: The maximum profit can be achieved by:
 * Buying at prices[0] = 1
 * Selling at prices[3] = 8
 * Buying at prices[4] = 4
 * Selling at prices[5] = 9
 * The total profit is ((8 - 1) - 2) + ((9 - 4) - 2) = 8.
 *
 *
 *
 * Note:
 * 0 < prices.length <= 50000.
 * 0 < prices[i] < 50000.
 * 0 <= fee < 50000.
 *
 */

func TestBestTimetoBuyandSellStockwithTransactionFee(t *testing.T) {
	var cases = []struct {
		input  []int
		fee    int
		output int
	}{
		{
			input:  []int{1, 3, 2, 8, 4, 9},
			fee:    2,
			output: 8,
		},
	}
	for _, c := range cases {
		x := maxProfitFee(c.input, c.fee)
		if x != c.output {
			t.Fail()
		}
	}
}

// submission codes start here

func maxProfitFee(prices []int, fee int) int {
	p1, p2 := 0, math.MinInt32
	for _, v := range prices {
		p1old := p1
		if tmp := p2 + v; tmp > p1 {
			p1 = tmp
		}
		if tmp := p1old-v-fee; tmp > p2 {
			p2 = tmp
		}
	}
	return p1
}

// submission codes end
