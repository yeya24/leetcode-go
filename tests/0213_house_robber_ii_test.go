package tests

import (
	"testing"
)

/**
 * [213] House Robber II
 *
 * You are a professional robber planning to rob houses along a street. Each house has a certain amount of money stashed. All houses at this place are arranged in a circle. That means the first house is the neighbor of the last one. Meanwhile, adjacent houses have security system connected and it will automatically contact the police if two adjacent houses were broken into on the same night.
 *
 * Given a list of non-negative integers representing the amount of money of each house, determine the maximum amount of money you can rob tonight without alerting the police.
 *
 * Example 1:
 *
 *
 * Input: [2,3,2]
 * Output: 3
 * Explanation: You cannot rob house 1 (money = 2) and then rob house 3 (money = 2),
 *              because they are adjacent houses.
 *
 *
 * Example 2:
 *
 *
 * Input: [1,2,3,1]
 * Output: 4
 * Explanation: Rob house 1 (money = 1) and then rob house 3 (money = 3).
 *              Total amount you can rob = 1 + 3 = 4.
 *
 */

func TestHouseRobberII(t *testing.T) {
	var cases = []struct {
		input  []int
		output int
	}{
		//{
		//	input:  []int{2, 3, 2},
		//	output: 3,
		//},
		{
			input:  []int{1, 2, 3, 1},
			output: 4,
		},
		{
			input:  []int{2, 1, 1, 2},
			output: 3,
		},
	}
	for _, c := range cases {
		x := robII(c.input)
		if x != c.output {
			t.Fail()
		}
	}
}

// submission codes start here

// use rolling arrays.
func robII(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	if len(nums) == 1 {
		return nums[0]
	}
	return max(robRow(nums[1:]), robRow(nums[:len(nums)-1]))
}

func robRow(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	if len(nums) == 1 {
		return nums[0]
	}
	dp := make([]int, 2)
	dp[0] = nums[0]
	dp[1] = max(nums[0], nums[1])
	for i := 2; i < len(nums); i++ {
		dp[i%2] = max(dp[(i-1)%2], dp[i%2] + nums[i])
	}
	return max(dp[0], dp[1])
}

// submission codes end
