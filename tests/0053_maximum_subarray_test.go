package tests

import (
    "math"
    "testing"
)

/**
 * [53] Maximum Subarray
 *
 * Given an integer array nums, find the contiguous subarray (containing at least one number) which has the largest sum and return its sum.
 * 
 * Example:
 * 
 * 
 * Input: [-2,1,-3,4,-1,2,1,-5,4],
 * Output: 6
 * Explanation: [4,-1,2,1] has the largest sum = 6.
 * 
 * 
 * Follow up:
 * 
 * If you have figured out the O(n) solution, try coding another solution using the divide and conquer approach, which is more subtle.
 * 
 */

func TestMaximumSubarray(t *testing.T) {
    var cases = []struct {
        input  []int
        output int
    }{
        {
            input:  []int{-2, 1, -3, 4, -1, 2, 1, -5, 4},
            output: 6,
        },
    }
    for _, c := range cases {
        x := maxSubArray(c.input)
        if x != c.output {
            t.Fail()
        }

    }
}

// submission codes start here

func maxSubArray(nums []int) int {
    max := math.MinInt64
    var sum int
    for _, n := range nums {
        if sum <= 0 {
            sum = n
        } else {
            sum += n
        }
        if sum > max {
            max = sum
        }
    }
    return max
}

// submission codes end
