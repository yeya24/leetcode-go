package tests

import (
	"testing"
)

/**
 * [64] Minimum Path Sum
 *
 * Given a m x n grid filled with non-negative numbers, find a path from top left to bottom right which minimizes the sum of all numbers along its path.
 *
 * Note: You can only move either down or right at any point in time.
 *
 * Example:
 *
 *
 * Input:
 * [
 *   [1,3,1],
 *   [1,5,1],
 *   [4,2,1]
 * ]
 * Output: 7
 * Explanation: Because the path 1&rarr;3&rarr;1&rarr;1&rarr;1 minimizes the sum.
 *
 *
 */

func TestMinimumPathSum(t *testing.T) {
	var cases = []struct {
		input  [][]int
		output int
	}{
		{
			input:  [][]int{{1, 3, 1}, {1, 5, 1}, {4, 2, 1}},
			output: 7,
		},
		{
			input:  [][]int{{1, 2, 5}, {3, 2, 1}},
			output: 6,
		},
	}
	for _, c := range cases {
		x := minPathSum(c.input)
		if x != c.output {
			t.Fail()
		}
	}
}

// submission codes start here

func minPathSum(grid [][]int) int {
	m := len(grid)
	n := len(grid[0])
	l := m * n
	if m == 0 || n == 0 {
		return 0
	}
	paths := make([]int, l)
	paths[0] = grid[0][0]
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if i != 0 || j != 0 {
				if i == 0 {
					paths[i*n+j] = paths[i*n+j-1] + grid[i][j]
				} else if j == 0 {
					paths[i*n+j] = paths[n*(i-1)+j] + grid[i][j]
				} else {
					paths[i*n+j] = min(paths[i*n+j-1], paths[n*(i-1)+j]) + grid[i][j]
				}
			}
		}
	}
	return paths[l-1]
}

// submission codes end
