package tests

import (
	"reflect"
	"testing"
)

/**
 * [941] Sort Array By Parity
 *
 * Given an array A of non-negative integers, return an array consisting of all the even elements of A, followed by all the odd elements of A.
 *
 * You may return any answer array that satisfies this condition.
 *
 *
 *
 *
 * Example 1:
 *
 *
 * Input: [3,1,2,4]
 * Output: [2,4,3,1]
 * The outputs [4,2,3,1], [2,4,1,3], and [4,2,1,3] would also be accepted.
 *
 *
 *
 *
 * Note:
 *
 *
 * 	1 <= A.length <= 5000
 * 	0 <= A[i] <= 5000
 *
 *
 *
 */

func TestSortArrayByParity(t *testing.T) {
	var cases = []struct {
		input  []int
		output []int
	}{
		{
			input:  []int{3, 1, 2, 4},
			output: []int{2, 4, 1, 3},
		},
	}
	for _, c := range cases {
		x := sortArrayByParity(c.input)
		if !reflect.DeepEqual(x, c.output) {
			t.Fail()
		}
	}
}

// submission codes start here

func sortArrayByParity(A []int) []int {
	l := len(A)
	m := make([]int, l)
	i, j := 0, l-1
	for _, n := range A {
        if n % 2 == 0 {
            m[i] = n
            i++
        } else {
            m[j] = n
            j--
        }
	}
	return m
}

// submission codes end
