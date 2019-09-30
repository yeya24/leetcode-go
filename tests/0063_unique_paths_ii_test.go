package tests

import (
	"testing"
)

/**
 * [63] Unique Paths II
 *
 * A robot is located at the top-left corner of a m x n grid (marked 'Start' in the diagram below).
 *
 * The robot can only move either down or right at any point in time. The robot is trying to reach the bottom-right corner of the grid (marked 'Finish' in the diagram below).
 *
 * Now consider if some obstacles are added to the grids. How many unique paths would there be?
 *
 *
 *
 * An obstacle and empty space is marked as 1 and 0 respectively in the grid.
 *
 * Note: m and n will be at most 100.
 *
 * Example 1:
 *
 *
 * Input:
 * [
 *   [0,0,0],
 *   [0,1,0],
 *   [0,0,0]
 * ]
 * Output: 2
 * Explanation:
 * There is one obstacle in the middle of the 3x3 grid above.
 * There are two ways to reach the bottom-right corner:
 * 1. Right -> Right -> Down -> Down
 * 2. Down -> Down -> Right -> Right
 *
 *
 */

func TestUniquePathsII(t *testing.T) {
	var cases = []struct {
		input  [][]int
		output int
	}{
		{
			input:  [][]int{{0, 0, 0}, {0, 1, 0}, {0, 0, 0}},
			output: 2,
		},
	}
	for _, c := range cases {
		x := uniquePathsWithObstacles(c.input)
		if x != c.output {
			t.Fail()
		}
	}
}

// submission codes start here

func uniquePathsWithObstacles(obstacleGrid [][]int) int {
	m := len(obstacleGrid)
	if m == 0 {
		return 0
	}
	n := len(obstacleGrid[0])
	if n == 0 {
		return 0
	}
	l := m * n
	paths := make([]int, l)
	if obstacleGrid[0][0] == 1 {
		return 0
	}
	paths[0] = 1
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if i != 0 || j != 0 {
				if obstacleGrid[i][j] == 1 {
					paths[i*n+j] = 0
				} else {
					if i == 0 {
						paths[i*n+j] = paths[i*n+j-1]
					} else if j == 0 {
						paths[i*n+j] = paths[n*(i-1)+j]
					} else {
						paths[i*n+j] = paths[i*n+j-1] + paths[n*(i-1)+j]
					}
				}
			}
		}
	}
	return paths[l-1]
}

// submission codes end
