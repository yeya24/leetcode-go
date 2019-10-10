package tests

import (
    "testing"
)

/**
 * [674] Longest Continuous Increasing Subsequence
 *
 * 
 * Given an unsorted array of integers, find the length of longest continuous increasing subsequence (subarray).
 * 
 * 
 * Example 1:
 * 
 * Input: [1,3,5,4,7]
 * Output: 3
 * Explanation: The longest continuous increasing subsequence is [1,3,5], its length is 3. 
 * Even though [1,3,5,7] is also an increasing subsequence, it's not a continuous one where 5 and 7 are separated by 4. 
 * 
 * 
 * 
 * Example 2:
 * 
 * Input: [2,2,2,2,2]
 * Output: 1
 * Explanation: The longest continuous increasing subsequence is [2], its length is 1. 
 * 
 * 
 * 
 * Note:
 * Length of the array will not exceed 10,000.
 * 
 */

func TestLongestContinuousIncreasingSubsequence(t *testing.T) {
    var cases = []struct {
        input  []int
        output int
    }{
        {
            input:  []int{1,3,5,4,7},
            output: 3,
        },
        {
            input:  []int{2,2,2,2,2},
            output: 1,
        },
        {
            input:  []int{2,3,2,2,3},
            output: 2,
        },
    }
    for _, c := range cases {
        x := findLengthOfLCIS(c.input)
        if x != c.output {
            t.Fail()
        }
    }
}

// submission codes start here

func findLengthOfLCIS(nums []int) int {
    if len(nums) <= 1 {
        return len(nums)
    }
    cnt := 1
    max := 1
    for i := 1; i < len(nums); i++ {
        if nums[i] > nums[i-1] {
            cnt+=1
        } else {
            cnt=1
        }
        if cnt > max {
            max = cnt
        }
    }
    return max
}

// submission codes end
