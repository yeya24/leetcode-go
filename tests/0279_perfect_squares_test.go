package tests

import (
	"math"
	"testing"
)

/**
 * [279] Perfect Squares
 *
 * Given a positive integer n, find the least number of perfect square numbers (for example, 1, 4, 9, 16, ...) which sum to n.
 *
 * Example 1:
 *
 *
 * Input: n = 12
 * Output: 3
 * Explanation: 12 = 4 + 4 + 4.
 *
 * Example 2:
 *
 *
 * Input: n = 13
 * Output: 2
 * Explanation: 13 = 4 + 9.
 */

func TestPerfectSquares(t *testing.T) {
	var cases = []struct {
		input  int
		output int
	}{
		{
			input:  12,
			output: 3,
		},
		{
			input:  13,
			output: 2,
		},
	}
	for _, c := range cases {
		x := numSquares(c.input)
		if x != c.output {
			t.Fail()
		}
	}
}

// submission codes start here

func numSquares(n int) int {
	dp := make([]int, n+1)
	for i := 0; i < len(dp); i++ {
		dp[i] = math.MaxInt64
	}
	dp[0] = 0
	for i := 1; i <= n; i++ {
		for j := 1; j * j <= i; j++ {
			dp[i] = min(dp[i], dp[i-j*j]+1)
		}
	}
	return dp[n]
}

// submission codes end
