package tests

import (
	"testing"
)

/**
 * [200] Number of Islands
 *
 * Given a 2d grid map of '1's (land) and '0's (water), count the number of islands. An island is surrounded by water and is formed by connecting adjacent lands horizontally or vertically. You may assume all four edges of the grid are all surrounded by water.
 *
 * Example 1:
 *
 *
 * Input:
 * 11110
 * 11010
 * 11000
 * 00000
 *
 * Output: 1
 *
 *
 * Example 2:
 *
 *
 * Input:
 * 11000
 * 11000
 * 00100
 * 00011
 *
 * Output: 3
 *
 */

func TestNumberofIslands(t *testing.T) {
	var cases = []struct {
		input  [][]byte
		output int
	}{
		{
			input:  [][]byte{[]byte("11110"), []byte("11010"), []byte("11000"), []byte("00000")},
			output: 1,
		},
		{
			input:  [][]byte{[]byte("11000"), []byte("11000"), []byte("00100"), []byte("00011")},
			output: 3,
		},
	}
	for _, c := range cases {
		x := numIslands(c.input)
		if x != c.output {
			t.Fail()
		}
	}
}

// submission codes start here

func numIslands(grid [][]byte) int {
	m := len(grid)
	if m == 0 {
		return 0
	}
	n := len(grid[0])
	if n == 0 {
		return 0
	}
	c := 0
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
            if grid[i][j] == '1' {
                dfsIslands(grid, i, j)
                c++
            }
		}
	}
	return c
}

func dfsIslands(grid [][]byte, x, y int) {
    if x < 0 || y < 0 || x >= len(grid) || y >= len(grid[0]) || grid[x][y] == '0' {
        return
    }
    grid[x][y] = '0'
    dfsIslands(grid, x+1, y)
	dfsIslands(grid, x-1, y)
	dfsIslands(grid, x, y+1)
	dfsIslands(grid, x, y-1)
}

// submission codes end
