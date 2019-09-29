package tests

import (
    "testing"
)

/**
 * [1062] Partition Array Into Three Parts With Equal Sum
 *
 * Given an array A of integers, return true if and only if we can partition the array into three non-empty parts with equal sums.
 * 
 * Formally, we can partition the array if we can find indexes i+1 < j with (A[0] + A[1] + ... + A[i] == A[i+1] + A[i+2] + ... + A[j-1] == A[j] + A[j-1] + ... + A[A.length - 1])
 * 
 *  
 * 
 * Example 1:
 * 
 * 
 * Input: [0,2,1,-6,6,-7,9,1,2,0,1]
 * Output: true
 * Explanation: 0 + 2 + 1 = -6 + 6 - 7 + 9 + 1 = 2 + 0 + 1
 * 
 * 
 * 
 * Example 2:
 * 
 * 
 * Input: [0,2,1,-6,6,7,9,-1,2,0,1]
 * Output: false
 * 
 * 
 * 
 * Example 3:
 * 
 * 
 * Input: [3,3,6,5,-2,2,5,1,-9,4]
 * Output: true
 * Explanation: 3 + 3 = 6 = 5 - 2 + 2 + 5 + 1 - 9 + 4
 * 
 * 
 * 
 * 
 *  
 * 
 * Note:
 * 
 * 
 * 	3 <= A.length <= 50000
 * 	-10000 <= A[i] <= 10000
 * 
 */

func TestPartitionArrayIntoThreePartsWithEqualSum(t *testing.T) {
    var cases = []struct {
        input  []int
        output bool
    }{
        {
            input:  []int{0,2,1,-6,6,-7,9,1,2,0,1},
            output: true,
        },
        {
            input:  []int{0,2,1,-6,6,7,9,-1,2,0,1},
            output: false,
        },
        {
            input: []int{3,3,6,5,-2,2,5,1,-9,4},
            output: true,
        },
    }
    for _, c := range cases {
        x := canThreePartsEqualSum(c.input)
        if x != c.output {
            t.Fail()
        }
    }
}

// submission codes start here

func canThreePartsEqualSum(A []int) bool {
    sum := 0
    for _, v := range A {
        sum += v
    }
    if sum % 3 != 0 {
        return false
    }
    target := sum / 3
    currentSum := 0
    for _, v := range A {
        currentSum += v
        if currentSum == target {
            currentSum = 0
        }
    }
    return currentSum == 0
}

// submission codes end
