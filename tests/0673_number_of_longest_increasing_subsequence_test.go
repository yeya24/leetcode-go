package tests

import (
    "testing"
)

/**
 * [673] Number of Longest Increasing Subsequence
 *
 * 
 * Given an unsorted array of integers, find the number of longest increasing subsequence.
 * 
 * 
 * Example 1:
 * 
 * Input: [1,3,5,4,7]
 * Output: 2
 * Explanation: The two longest increasing subsequence are [1, 3, 4, 7] and [1, 3, 5, 7].
 * 
 * 
 * 
 * Example 2:
 * 
 * Input: [2,2,2,2,2]
 * Output: 5
 * Explanation: The length of longest continuous increasing subsequence is 1, and there are 5 subsequences' length is 1, so output 5.
 * 
 * 
 * 
 * Note:
 * Length of the given array will be not exceed 2000 and the answer is guaranteed to be fit in 32-bit signed int.
 * 
 */

func TestNumberofLongestIncreasingSubsequence(t *testing.T) {
    var cases = []struct {
        input  []int
        output int
    }{
        {
            input:  []int{1,3,5,4,7},
            output: 2,
        },
        {
            input:  []int{2,2,2,2,2},
            output: 5,
        },
    }
    for _, c := range cases {
        x := findNumberOfLIS(c.input)
        if x != c.output {
            t.Fail()
        }
    }
}

// submission codes start here

func findNumberOfLIS(nums []int) int {
    if len(nums) <= 1 {
        return len(nums)
    }
    length := make([]int, len(nums))
    count := make([]int, len(nums))
    for i := 0; i < len(nums); i++ {
        length[i] = 1
        count[i] = 1
    }
    maxL := 1
    for i := 1; i < len(nums); i++ {
        for j := 0; j < i; j++ {
            if nums[i] > nums[j] {
                if length[j] >= length[i] {
                    length[i] = length[j] + 1
                    count[i] = count[j]
                } else if length[j] + 1 == length[i] {
                    count[i] += count[j]
                }
                if length[i] > maxL {
                    maxL = length[i]
                }
            }

        }
    }
    var cnt int
    for i := 0; i < len(nums); i++ {
        if length[i] == maxL {
            cnt += count[i]
        }
    }
    return cnt
}

// submission codes end
