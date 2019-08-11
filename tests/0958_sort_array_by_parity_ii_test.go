package tests

import (
	"reflect"
	"testing"
)

/**
 * [958] Sort Array By Parity II
 *
 * Given an array A of non-negative integers, half of the integers in A are odd, and half of the integers are even.
 *
 * Sort the array so that whenever A[i] is odd, i is odd; and whenever A[i] is even, i is even.
 *
 * You may return any answer array that satisfies this condition.
 *
 *
 *
 * Example 1:
 *
 *
 * Input: [4,2,5,7]
 * Output: [4,5,2,7]
 * Explanation: [4,7,2,5], [2,5,4,7], [2,7,4,5] would also have been accepted.
 *
 *
 *
 *
 * Note:
 *
 *
 * 	2 <= A.length <= 20000
 * 	A.length % 2 == 0
 * 	0 <= A[i] <= 1000
 *
 *
 *
 *
 *
 */

func TestSortArrayByParityII(t *testing.T) {
	var cases = []struct {
		input  []int
		output []int
	}{
		{
			input:  []int{4, 2, 5, 7},
			output: []int{4, 7, 2, 5},
		},
	}
	for _, c := range cases {
		x := sortArrayByParityII(c.input)
		if !reflect.DeepEqual(x, c.output) {
			t.Fail()
		}
	}
}

// submission codes start here

func sortArrayByParityII(A []int) []int {
	l := len(A)
    m := make([]int, l)
	i, j := 0, l-1
    for _, n := range A {
        if n % 2 == 0 {
            m[i] = n
            i+=2
        } else {
            m[j] = n
            j-=2
        }
    }
    return m
}

// submission codes end
