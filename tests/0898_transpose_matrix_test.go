package tests

import (
	"reflect"
	"testing"
)

/**
 * [898] Transpose Matrix
 *
 * Given a matrix A, return the transpose of A.
 *
 * The transpose of a matrix is the matrix flipped over it's main diagonal, switching the row and column indices of the matrix.
 *
 *
 *
 *
 * Example 1:
 *
 *
 * Input: [[1,2,3],[4,5,6],[7,8,9]]
 * Output: [[1,4,7],[2,5,8],[3,6,9]]
 *
 *
 *
 * Example 2:
 *
 *
 * Input: [[1,2,3],[4,5,6]]
 * Output: [[1,4],[2,5],[3,6]]
 *
 *
 *
 *
 * Note:
 *
 *
 * 	1 <= A.length <= 1000
 * 	1 <= A[0].length <= 1000
 *
 *
 *
 *
 */

func TestTransposeMatrix(t *testing.T) {
	var cases = []struct {
		input  [][]int
		output [][]int
	}{
		{
			input:  [][]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}},
			output: [][]int{{1, 4, 7}, {2, 5, 8}, {3, 6, 9}},
		},
		{
			input:  [][]int{{1, 2, 3}, {4, 5, 6}},
			output: [][]int{{1, 4}, {2, 5}, {3, 6}},
		},
	}
	for _, c := range cases {
		x := transpose(c.input)
		if !reflect.DeepEqual(x, c.output) {
			t.Fail()
		}
	}
}

// submission codes start here

func transpose(A [][]int) [][]int {
	m := len(A)
	n := len(A[0])
	newMatrix := make([][]int, n)
	for i := 0; i < n; i++ {
		newMatrix[i] = make([]int, m)
	}

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			newMatrix[j][i] = A[i][j]
		}
	}
	return newMatrix
}

// submission codes end
