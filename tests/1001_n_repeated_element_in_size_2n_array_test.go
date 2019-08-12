package tests

import (
	"testing"
)

/**
 * [1001] N-Repeated Element in Size 2N Array
 *
 * In a array A of size 2N, there are N+1 unique elements, and exactly one of these elements is repeated N times.
 *
 * Return the element repeated N times.
 *
 *
 *
 *
 *
 *
 *
 * Example 1:
 *
 *
 * Input: [1,2,3,3]
 * Output: 3
 *
 *
 *
 * Example 2:
 *
 *
 * Input: [2,1,2,5,3,2]
 * Output: 2
 *
 *
 *
 * Example 3:
 *
 *
 * Input: [5,1,5,2,5,3,5,4]
 * Output: 5
 *
 *
 *
 *
 * Note:
 *
 *
 * 	4 <= A.length <= 10000
 * 	0 <= A[i] < 10000
 * 	A.length is even
 *
 *
 *
 *
 *
 */

func TestNRepeatedElementinSize2NArray(t *testing.T) {
	var cases = []struct {
		input  []int
		output int
	}{
		{
			input:  []int{1, 2, 3, 3},
			output: 3,
		},
		{
			input:  []int{5, 1, 5, 2, 5, 3, 5, 4},
			output: 5,
		},
	}
	for _, c := range cases {
		x := repeatedNTimes(c.input)
		if x != c.output {
			t.Fail()
		}
	}
}

// submission codes start here

func repeatedNTimes(A []int) int {
	m := make(map[int]struct{})
	for _, v := range A {
		if _, ok := m[v]; ok {
			return v
		} else {
		    m[v] = struct{}{}
        }
	}
	return -1
}

// submission codes end
