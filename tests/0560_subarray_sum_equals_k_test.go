package tests

import (
    "testing"
)

/**
 * [560] Subarray Sum Equals K
 *
 * Given an array of integers and an integer k, you need to find the total number of continuous subarrays whose sum equals to k.
 * 
 * Example 1:
 * 
 * Input:nums = [1,1,1], k = 2
 * Output: 2
 * 
 * 
 * 
 * Note:
 * 
 * The length of the array is in range [1, 20,000].
 * The range of numbers in the array is [-1000, 1000] and the range of the integer k is [-1e7, 1e7].
 * 
 * 
 * 
 */

func TestSubarraySumEqualsK(t *testing.T) {
    var cases = []struct {
        input  []int
        k int
        output int
    }{
        {
            input:  []int{1,1,1},
            k: 2,
            output: 2,
        },
        {
            input:  []int{1,1,1,-1,2},
            k: 2,
            output: 5,
        },
    }
    for _, c := range cases {
        x := subarraySum2(c.input, c.k)
        if x != c.output {
            t.Fail()
        }
    }
}

// submission codes start here

// brute force.
func subarraySum(nums []int, k int) int {
    var cnt int
    for i := 0; i < len(nums); i++ {
        sum := 0
        for j := i; j < len(nums); j++ {
            sum += nums[j]
            if sum == k {
                cnt++
            }
        }
    }
    return cnt
}

// use a hashmap to store sum-k
func subarraySum2(nums []int, k int) int {
    var cnt int
    var sum int
    m := make(map[int]int)
    m[0] = 1
    for i := 0; i < len(nums); i++ {
        sum += nums[i]
        if _, ok := m[sum-k]; ok {
            cnt += m[sum-k]
        }
        m[sum]+=1
    }
    return cnt
}

// submission codes end
