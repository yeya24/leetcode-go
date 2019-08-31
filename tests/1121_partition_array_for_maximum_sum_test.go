package tests

import (
    "math"
    "testing"
)

/**
 * [1121] Partition Array for Maximum Sum
 *
 * Given an integer array A, you partition the array into (contiguous) subarrays of length at most K.  After partitioning, each subarray has their values changed to become the maximum value of that subarray.
 * 
 * Return the largest sum of the given array after partitioning.
 * 
 *  
 * 
 * Example 1:
 * 
 * 
 * Input: A = [1,15,7,9,2,5,10], K = 3
 * Output: 84
 * Explanation: A becomes [15,15,15,9,10,10,10]
 * 
 *  
 * 
 * Note:
 * 
 * 
 * 	1 <= K <= A.length <= 500
 * 	0 <= A[i] <= 10^6
 * 
 * 
 */

func TestPartitionArrayforMaximumSum(t *testing.T) {
    var cases = []struct {
        input  []int
        k int
        output int
    }{
        {
            input:  []int{1,15,7,9,2,5,10},
            k: 3,
            output: 84,
        },
    }
    for _, c := range cases {
        x := maxSumAfterPartitioning(c.input, c.k)
        if x != c.output {
            t.Fail()
        }
    }
}

// submission codes start here

func maxSumAfterPartitioning(A []int, K int) int {
    n := len(A)
    dp := make([]int, 1+n)
    dp[0] = 0
    for i := 1; i <= n; i++ {
        m := math.MinInt64
        for j := 1; j <= min(i, K); j++ {
            m = max(m, A[i-j])
            dp[i] = max(dp[i], dp[i-j] + m * j)
        }
    }
    return dp[n]
}

// submission codes end
