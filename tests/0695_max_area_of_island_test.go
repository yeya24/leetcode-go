package tests

import (
	"testing"
)

/**
 * [695] Max Area of Island
 *
 * Given a non-empty 2D array grid of 0's and 1's, an island is a group of 1's (representing land) connected 4-directionally (horizontal or vertical.) You may assume all four edges of the grid are surrounded by water.
 *
 * Find the maximum area of an island in the given 2D array. (If there is no island, the maximum area is 0.)
 *
 * Example 1:
 *
 *
 * [[0,0,1,0,0,0,0,1,0,0,0,0,0],
 *  [0,0,0,0,0,0,0,1,1,1,0,0,0],
 *  [0,1,1,0,1,0,0,0,0,0,0,0,0],
 *  [0,1,0,0,1,1,0,0,1,0,1,0,0],
 *  [0,1,0,0,1,1,0,0,1,1,1,0,0],
 *  [0,0,0,0,0,0,0,0,0,0,1,0,0],
 *  [0,0,0,0,0,0,0,1,1,1,0,0,0],
 *  [0,0,0,0,0,0,0,1,1,0,0,0,0]]
 *
 * Given the above grid, return 6. Note the answer is not 11, because the island must be connected 4-directionally.
 *
 * Example 2:
 *
 *
 * [[0,0,0,0,0,0,0,0]]
 * Given the above grid, return 0.
 *
 * Note: The length of each dimension in the given grid does not exceed 50.
 *
 */

func TestMaxAreaofIsland(t *testing.T) {
	var cases = []struct {
		input  [][]int
		output int
	}{
		{
			input: [][]int{{0, 0, 1, 0, 0, 0, 0, 1, 0, 0, 0, 0, 0},
				{0, 0, 0, 0, 0, 0, 0, 1, 1, 1, 0, 0, 0},
				{0, 1, 1, 0, 1, 0, 0, 0, 0, 0, 0, 0, 0},
				{0, 1, 0, 0, 1, 1, 0, 0, 1, 0, 1, 0, 0},
				{0, 1, 0, 0, 1, 1, 0, 0, 1, 1, 1, 0, 0},
				{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1, 0, 0},
				{0, 0, 0, 0, 0, 0, 0, 1, 1, 1, 0, 0, 0},
				{0, 0, 0, 0, 0, 0, 0, 1, 1, 0, 0, 0, 0}},
			output: 6,
		},
		{
			input:  [][]int{{0, 0, 0, 0, 0, 0, 0, 0}},
			output: 0,
		},
	}
	for _, c := range cases {
		x := maxAreaOfIsland(c.input)
		if x != c.output {
			t.Fail()
		}
	}
}

// submission codes start here

func maxAreaOfIsland(grid [][]int) int {
	m := len(grid)
	if m == 0 {
		return 0
	}
	n := len(grid[0])
	if n == 0 {
		return 0
	}
	maxAreas := 0
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if grid[i][j] == 1 {
				maxAreas = max(maxAreas, dfsIslandsAreas(grid, i, j, 0))
			}
		}
	}
	return maxAreas
}

func dfsIslandsAreas(grid [][]int, x, y, areas int) int {
	if x < 0 || y < 0 || x >= len(grid) || y >= len(grid[0]) || grid[x][y] == 0 {
		return 0
	}
	grid[x][y] = 0
	areas++
	dx := [4]int{1, 0, -1, 0}
	dy := [4]int{0, 1, 0, -1}
	for i := 0; i < 4; i++ {
		v := dfsIslandsAreas(grid, x+dx[i], y+dy[i], areas)
		if v > areas {
			areas = v
		}
	}
	return areas
}

// submission codes end
