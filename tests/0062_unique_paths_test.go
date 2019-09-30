package tests

import (
	"testing"
)

/**
 * [62] Unique Paths
 *
 * A robot is located at the top-left corner of a m x n grid (marked 'Start' in the diagram below).
 *
 * The robot can only move either down or right at any point in time. The robot is trying to reach the bottom-right corner of the grid (marked 'Finish' in the diagram below).
 *
 * How many possible unique paths are there?
 *
 *
 * Above is a 7 x 3 grid. How many possible unique paths are there?
 *
 * Note: m and n will be at most 100.
 *
 * Example 1:
 *
 *
 * Input: m = 3, n = 2
 * Output: 3
 * Explanation:
 * From the top-left corner, there are a total of 3 ways to reach the bottom-right corner:
 * 1. Right -> Right -> Down
 * 2. Right -> Down -> Right
 * 3. Down -> Right -> Right
 *
 *
 * Example 2:
 *
 *
 * Input: m = 7, n = 3
 * Output: 28
 *
 */

func TestUniquePaths(t *testing.T) {
	var cases = []struct {
		m      int
		n      int
		output int
	}{
		{
			m:      3,
			n:      2,
			output: 3,
		},
		{
			m:      7,
			n:      3,
			output: 28,
		},
	}
	for _, c := range cases {
		x := uniquePaths(c.m, c.n)
		if x != c.output {
			t.Fail()
		}
	}
}

// submission codes start here

func uniquePaths(m int, n int) int {
	l := m * n
	paths := make([]int, l)
	paths[0] += 1
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if i != 0 || j != 0 {
				if i == 0 {
					paths[i*n+j] = 1
				} else if j == 0 {
					paths[i*n+j] = 1
				} else {
					paths[i*n+j] = paths[i*n+j-1] + paths[n*(i-1)+j]
				}
			}
		}
	}
	return paths[l-1]
}

// submission codes end
