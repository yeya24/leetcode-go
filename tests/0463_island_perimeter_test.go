package tests

import (
	"reflect"
	"testing"
)

/**
 * [463] Island Perimeter
 *
 * You are given a map in form of a two-dimensional integer grid where 1 represents land and 0 represents water.
 *
 * Grid cells are connected horizontally/vertically (not diagonally). The grid is completely surrounded by water, and there is exactly one island (i.e., one or more connected land cells).
 *
 * The island doesn't have "lakes" (water inside that isn't connected to the water around the island). One cell is a square with side length 1. The grid is rectangular, width and height don't exceed 100. Determine the perimeter of the island.
 *
 *
 *
 * Example:
 *
 *
 * Input:
 * [[0,1,0,0],
 *  [1,1,1,0],
 *  [0,1,0,0],
 *  [1,1,0,0]]
 *
 * Output: 16
 *
 * Explanation: The perimeter is the 16 yellow stripes in the image below:
 *
 *
 *
 *
 */

func TestIslandPerimeter(t *testing.T) {
	var cases = []struct {
		input  [][]int
		output int
	}{
		{
			input:  [][]int{{0,1,0,0},{1,1,1,0},{0,1,0,0},{1,1,0,0}},
			output: 16,
		},
	}
	for _, c := range cases {
		x := islandPerimeter(c.input)
		if !reflect.DeepEqual(x, c.output) {
			t.Fail()
		}
	}
}

// submission codes start here

func islandPerimeter(grid [][]int) int {
	var island, neighbor int
	if grid == nil {
		return 0
	}
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[0]); j++ {
			if grid[i][j] == 1 {
				island++
				if j < len(grid[0])-1 && grid[i][j+1] == 1 {
					neighbor++
				}
				if i < len(grid)-1 && grid[i+1][j] == 1 {
					neighbor++
				}
			}
		}
	}
	return island*4 - neighbor*2
}

// submission codes end
