package tests

import (
	"testing"
)

/**
 * [1147] Flip Columns For Maximum Number of Equal Rows
 *
 * Given a matrix consisting of 0s and 1s, we may choose any number of columns in the matrix and flip every cell in that column.  Flipping a cell changes the value of that cell from 0 to 1 or from 1 to 0.
 *
 * Return the maximum number of rows that have all values equal after some number of flips.
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
 * Input: [[0,1],[1,1]]
 * Output: 1
 * Explanation: After flipping no values, 1 row has all values equal.
 *
 *
 *
 * Example 2:
 *
 *
 * Input: [[0,1],[1,0]]
 * Output: 2
 * Explanation: After flipping values in the first column, both rows have equal values.
 *
 *
 *
 * Example 3:
 *
 *
 * Input: [[0,0,0],[0,0,1],[1,1,0]]
 * Output: 2
 * Explanation: After flipping values in the first two columns, the last two rows have equal values.
 *
 *
 *
 *
 * Note:
 *
 *
 * 	1 <= matrix.length <= 300
 * 	1 <= matrix[i].length <= 300
 * 	All matrix[i].length's are equal
 * 	matrix[i][j] is 0 or 1
 *
 *
 *
 *
 */

func TestFlipColumnsForMaximumNumberofEqualRows(t *testing.T) {
	var cases = []struct {
		input  [][]int
		output int
	}{
		{
			input:  [][]int{{0, 1}, {1, 1}},
			output: 1,
		},
		{
			input:  [][]int{{0, 1}, {1, 0}},
			output: 2,
		},
		{
			input:  [][]int{{0, 0, 0}, {0, 0, 1}, {1, 1, 0}},
			output: 2,
		},
	}
	for _, c := range cases {
		x := maxEqualRowsAfterFlips(c.input)
		if x != c.output {
			t.Fail()
		}
	}
}

// submission codes start here

func maxEqualRowsAfterFlips(matrix [][]int) int {
	m := make(map[string]int)
	var res int
    for _, v := range matrix {
    	ones := ""
    	zeros := ""
    	for _, x := range v {
    		if x == 1 {
    			zeros += "1"
    			ones += "0"
			} else {
				zeros += "0"
				ones += "1"
			}
		}
    	if _, ok := m[zeros]; ok {
    		m[zeros] += 1
		} else {
			m[zeros] = 1
		}
		if _, ok := m[ones]; ok {
			m[ones] += 1
		} else {
			m[ones] = 1
		}
		res = max(res, max(m[zeros], m[ones]))
	}
    return res
}

// submission codes end
