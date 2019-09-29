package tests

import (
	"fmt"
	"testing"
)

/**
 * [932] Monotonic Array
 *
 * An array is monotonic if it is either monotone increasing or monotone decreasing.
 *
 * An array A is monotone increasing if for all i <= j, A[i] <= A[j].  An array A is monotone decreasing if for all i <= j, A[i] >= A[j].
 *
 * Return true if and only if the given array A is monotonic.
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
 * Input: [1,2,2,3]
 * Output: true
 *
 *
 *
 * Example 2:
 *
 *
 * Input: [6,5,4,4]
 * Output: true
 *
 *
 *
 * Example 3:
 *
 *
 * Input: [1,3,2]
 * Output: false
 *
 *
 *
 * Example 4:
 *
 *
 * Input: [1,2,4,5]
 * Output: true
 *
 *
 *
 * Example 5:
 *
 *
 * Input: [1,1,1]
 * Output: true
 *
 *
 *
 *
 * Note:
 *
 *
 * 	1 <= A.length <= 50000
 * 	-100000 <= A[i] <= 100000
 *
 *
 *
 *
 *
 *
 *
 */

func TestMonotonicArray(t *testing.T) {
	var cases = []struct {
		input  []int
		output bool
	}{
		{
			input:  []int{1, 2, 2, 3},
			output: true,
		},
		{
			input:  []int{6, 5, 4, 4},
			output: true,
		},
		{
			input:  []int{1, 3, 2},
			output: false,
		},
		{
			input:  []int{1, 2, 4, 5},
			output: true,
		},
	}
	for _, c := range cases {
		x := isMonotonic(c.input)
		fmt.Println(x)
		if x != c.output {
			t.Fail()
		}
	}
}

// submission codes start here

func isMonotonic(A []int) bool {
	l := len(A)
	if l == 1 {
		return true
	}
	var flag bool
	j := 0
	for j < l-1 {
		if A[j] != A[j+1] {
			flag = A[j] > A[j+1]
			break
		}
		j++
	}
	if j == l-1 {
		return true
	}
	for i := j + 1; i < l-1; i++ {
		if A[i] == A[i+1] {
			continue
		}
		if A[i] > A[i+1] != flag {
			return false
		}
	}
	return true
}

// Best way in leetcode solution.
func isMonotonic1(A []int) bool {
	asc := true
	dsc := true
	for i := 0; i < len(A) -1; i++ {
		if A[i+1] > A[i] {
			dsc = false
		} else if A[i+1] < A[i] {
			asc = false
		}
	}
	return asc || dsc
}

// submission codes end
