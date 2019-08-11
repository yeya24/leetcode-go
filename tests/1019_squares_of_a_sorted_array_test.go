package tests

import (
	"math"
	"reflect"
	"testing"
)

/**
 * [1019] Squares of a Sorted Array
 *
 * Given an array of integers A sorted in non-decreasing order, return an array of the squares of each number, also in sorted non-decreasing order.
 *
 *
 *
 *
 * Example 1:
 *
 *
 * Input: [-4,-1,0,3,10]
 * Output: [0,1,9,16,100]
 *
 *
 *
 * Example 2:
 *
 *
 * Input: [-7,-3,2,3,11]
 * Output: [4,9,9,49,121]
 *
 *
 *
 *
 * Note:
 *
 *
 * 	1 <= A.length <= 10000
 * 	-10000 <= A[i] <= 10000
 * 	A is sorted in non-decreasing order.
 *
 *
 *
 */

func TestSquaresofaSortedArray(t *testing.T) {
	var cases = []struct {
		input  []int
		output []int
	}{
		{
			input:  []int{-4, -1, 0, 3, 10},
			output: []int{0, 1, 9, 16, 100},
		},
		{
			input:  []int{-7, -3, 2, 3, 11},
			output: []int{4, 9, 9, 49, 121},
		},
	}
	for _, c := range cases {
		x := sortedSquares(c.input)
		if !reflect.DeepEqual(x, c.output) {
			t.Fail()
		}
	}
}

// submission codes start here

func sortedSquares(A []int) []int {
	l := len(A)
	m := make([]int, l)
	i, j := 0, l-1
	for c := l - 1; c >= 0; c-- {
		if math.Abs(float64(A[i])) > math.Abs(float64(A[j])) {
			m[c] = A[i] * A[i]
			i++
		} else {
			m[c] = A[j] * A[j]
			j--
		}
	}
	return m
}

// submission codes end
