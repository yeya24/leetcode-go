package tests

import (
	"testing"
)

/**
 * [882] Peak Index in a Mountain Array
 *
 * Let's call an array A a mountain if the following properties hold:
 *
 *
 * 	A.length >= 3
 * 	There exists some 0 < i < A.length - 1 such that A[0] < A[1] < ... A[i-1] < A[i] > A[i+1] > ... > A[A.length - 1]
 *
 *
 * Given an array that is definitely a mountain, return any i such that A[0] < A[1] < ... A[i-1] < A[i] > A[i+1] > ... > A[A.length - 1].
 *
 * Example 1:
 *
 *
 * Input: [0,1,0]
 * Output: 1
 *
 *
 *
 * Example 2:
 *
 *
 * Input: [0,2,1,0]
 * Output: 1
 *
 *
 * Note:
 *
 *
 * 	3 <= A.length <= 10000
 * 	0 <= A[i] <= 10^6
 * 	A is a mountain, as defined above.
 *
 *
 */

func TestPeakIndexinaMountainArray(t *testing.T) {
	var cases = []struct {
		input  []int
		output int
	}{
		{
			input:  []int{0,1,0},
			output: 1,
		},
		{
			input:  []int{0,2,1,0},
			output: 1,
		},
	}
	for _, c := range cases {
		x := peakIndexInMountainArray(c.input)
		if x != c.output {
			t.Fail()
		}
	}
}

// submission codes start here

func peakIndexInMountainArray(A []int) int {
	for i := 1; i < len(A)-1; i++ {
        if A[i] > A[i+1] {
            return i
        }
	}
	return -1
}

// submission codes end
