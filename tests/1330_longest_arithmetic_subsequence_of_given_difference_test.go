package tests

import (
	"testing"
)

/**
 * [1330] Longest Arithmetic Subsequence of Given Difference
 *
 * Given an integer array arr and an integer arr which is an arithmetic sequence such that the difference between adjacent elements in the subsequence equals difference.
 *
 *
 * Example 1:
 *
 *
 * Input: arr = [1,2,3,4], difference = 1
 * Output: 4
 * Explanation: The longest arithmetic subsequence is [1,2,3,4].
 *
 * Example 2:
 *
 *
 * Input: arr = [1,3,5,7], difference = 1
 * Output: 1
 * Explanation: The longest arithmetic subsequence is any single element.
 *
 *
 * Example 3:
 *
 *
 * Input: arr = [1,5,7,8,5,3,4,2,1], difference = -2
 * Output: 4
 * Explanation: The longest arithmetic subsequence is [7,5,3,1].
 *
 *
 *
 * Constraints:
 *
 *
 * 	1 <= arr.length <= 10^5
 * 	-10^4 <= arr[i], difference <= 10^4
 *
 *
 */

func TestLongestArithmeticSubsequenceofGivenDifference(t *testing.T) {
	var cases = []struct {
		input  []int
		diff   int
		output int
	}{
		{
			input: []int{1,2,3,4},
			diff: 1,
			output: 4,
		},
		{
			input:  []int{1,3,5,7},
			diff: 1,
			output: 1,
		},
		{
			input:  []int{1,5,7,8,5,3,4,2,1},
			diff: -2,
			output: 4,
		},
	}
	for _, c := range cases {
		x := longestSubsequence(c.input, c.diff)
		if x != c.output {
			t.Fail()
		}
	}
}

// submission codes start here

func longestSubsequence(arr []int, difference int) int {
	dp := make(map[int]int)
	var ans int
	for _, v := range arr {
		dp[v] = max(dp[v], dp[v-difference]+1)
		ans = max(ans, dp[v])
	}
	return ans
}

// submission codes end
