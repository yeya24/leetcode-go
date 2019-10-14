package tests

import (
    "testing"
)

/**
 * [198] House Robber
 *
 * You are a professional robber planning to rob houses along a street. Each house has a certain amount of money stashed, the only constraint stopping you from robbing each of them is that adjacent houses have security system connected and it will automatically contact the police if two adjacent houses were broken into on the same night.
 * 
 * Given a list of non-negative integers representing the amount of money of each house, determine the maximum amount of money you can rob tonight without alerting the police.
 * 
 * Example 1:
 * 
 * 
 * Input: [1,2,3,1]
 * Output: 4
 * Explanation: Rob house 1 (money = 1) and then rob house 3 (money = 3).
 *              Total amount you can rob = 1 + 3 = 4.
 * 
 * Example 2:
 * 
 * 
 * Input: [2,7,9,3,1]
 * Output: 12
 * Explanation: Rob house 1 (money = 2), rob house 3 (money = 9) and rob house 5 (money = 1).
 *              Total amount you can rob = 2 + 9 + 1 = 12.
 * 
 * 
 */

func TestHouseRobber(t *testing.T) {
    var cases = []struct {
        input  []int
        output int
    }{
        {
            input:  []int{1,2,3,1},
            output: 4,
        },
        {
            input:  []int{2,7,9,3,1},
            output: 12,
        },
    }
    for _, c := range cases {
        x := rob(c.input)
        if x != c.output {
            t.Fail()
        }
    }
}

// submission codes start here

func rob(nums []int) int {
    if len(nums) == 0 {
       return 0
    }
    dp := make([]int, len(nums)+1)
    dp[0] = 0
    dp[1] = nums[0]
    for i := 1; i < len(nums); i++ {
       dp[i+1] = max(dp[i-1]+nums[i], dp[i])
    }
    return dp[len(nums)]
}

func rob1(nums []int) int {
    if len(nums) == 0 {
        return 0
    }
    if len(nums) == 1 {
        return nums[0]
    }
    dp := make([]int, len(nums))
    dp[0] = nums[0]
    dp[1] = max(nums[0],nums[1])
    for i := 2; i < len(nums); i++ {
        dp[i] = max(dp[i-1], dp[i-2]+nums[i])
    }
    return dp[len(nums)-1]
}

// use rolling array
func rob2(nums []int) int {
    if len(nums) == 0 {
        return 0
    }
    if len(nums) == 1 {
        return nums[0]
    }
    dp := make([]int, 2)
    dp[0] = nums[0]
    dp[1] = max(nums[0],nums[1])
    for i := 2; i < len(nums); i++ {
        dp[i%2] = max(dp[(i-1)%2], dp[(i-2)%2]+nums[i])
    }
    return max(dp[0], dp[1])
}

// submission codes end
