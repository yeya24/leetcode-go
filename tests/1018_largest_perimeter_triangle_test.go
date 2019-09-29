package tests

import (
	"testing"
)

/**
 * [1018] Largest Perimeter Triangle
 *
 * Given an array A of positive lengths, return the largest perimeter of a triangle with non-zero area, formed from 3 of these lengths.
 *
 * If it is impossible to form any triangle of non-zero area, return 0.
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
 * Input: [2,1,2]
 * Output: 5
 *
 *
 *
 * Example 2:
 *
 *
 * Input: [1,2,1]
 * Output: 0
 *
 *
 *
 * Example 3:
 *
 *
 * Input: [3,2,3,4]
 * Output: 10
 *
 *
 *
 * Example 4:
 *
 *
 * Input: [3,6,2,3]
 * Output: 8
 *
 *
 *
 *
 * Note:
 *
 *
 * 	3 <= A.length <= 10000
 * 	1 <= A[i] <= 10^6
 *
 *
 *
 *
 *
 */

func TestLargestPerimeterTriangle(t *testing.T) {
	var cases = []struct {
		input  []int
		output int
	}{
		{
			input:  []int{1, 2, 1},
			output: 0,
		},
		{
			input:  []int{3, 2, 3, 4},
			output: 10,
		},
		{
			input:  []int{2, 1, 2},
			output: 5,
		},
		{
			input:  []int{3, 6, 2, 3},
			output: 8,
		},
	}
	for _, c := range cases {
		x := largestPerimeter(c.input)
		if x != c.output {
			t.Fail()
		}
	}
}

// submission codes start here

func largestPerimeter(A []int) int {
	l := len(A)
	quickSort(A, 0, l-1)
	for i := l - 1; i >= 2; i-- {
		if A[i-2]+A[i-1] > A[i] {
			return A[i-2] + A[i-1] + A[i]
		}
	}
	return 0
}

// submission codes end
