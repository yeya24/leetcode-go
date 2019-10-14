package tests

import (
    "testing"
)

/**
 * [1331] Path with Maximum Gold
 *
 * In a gold mine grid of size m * n, each cell in this mine has an integer representing the amount of gold in that cell, 0 if it is empty.
 * 
 * Return the maximum amount of gold you can collect under the conditions:
 * 
 * 
 * 	Every time you are located in a cell you will collect all the gold in that cell.
 * 	From your position you can walk one step to the left, right, up or down.
 * 	You can't visit the same cell more than once.
 * 	Never visit a cell with 0 gold.
 * 	You can start and stop collecting gold from any position in the grid that has some gold.
 * 
 * 
 *  
 * Example 1:
 * 
 * 
 * Input: grid = [[0,6,0],[5,8,7],[0,9,0]]
 * Output: 24
 * Explanation:
 * [[0,6,0],
 *  [5,8,7],
 *  [0,9,0]]
 * Path to get the maximum gold, 9 -> 8 -> 7.
 * 
 * 
 * Example 2:
 * 
 * 
 * Input: grid = [[1,0,7],[2,0,6],[3,4,5],[0,3,0],[9,0,20]]
 * Output: 28
 * Explanation:
 * [[1,0,7],
 *  [2,0,6],
 *  [3,4,5],
 *  [0,3,0],
 *  [9,0,20]]
 * Path to get the maximum gold, 1 -> 2 -> 3 -> 4 -> 5 -> 6 -> 7.
 * 
 * 
 *  
 * Constraints:
 * 
 * 
 * 	1 <= grid.length, grid[i].length <= 15
 * 	0 <= grid[i][j] <= 100
 * 	There are at most 25 cells containing gold.
 * 
 */

func TestPathwithMaximumGold(t *testing.T) {
    var cases = []struct {
        input  [][]int
        output int
    }{
        {
            input:  [][]int{{0,6,0},{5,8,7},{0,9,0}},
            output: 24,
        },
        {
            input:  [][]int{{1,0,7},{2,0,6},{3,4,5},{0,3,0},{9,0,20}},
            output: 28,
        },
    }
    for _, c := range cases {
        x := getMaximumGold(c.input)
        if x != c.output {
            t.Fail()
        }
    }
}

// submission codes start here

func getMaximumGold(grid [][]int) int {
    m := len(grid)
    n := len(grid[0])
    var ans int
    for i := 0; i<m;i++ {
        for j := 0;j < n; j++ {
            ans = max(ans, dfsGold(grid, m, n, i, j, 0))
        }
    }
    return ans
}

func dfsGold(grid [][]int, m, n, x, y, cur int) int {
    if x < 0 || y < 0 || x >= m || y >= n || grid[x][y] == 0 {
        return 0
    }
    dx := []int{-1,0,1,0}
    dy := []int{0,1,0,-1}
    v := grid[x][y]
    cur += v
    ans := cur
    grid[x][y] = 0
    for i := 0; i < 4; i++ {
        tx, ty := x + dx[i], y + dy[i]
        ans = max(ans, dfsGold(grid, m, n, tx, ty, cur))
    }
    grid[x][y] = v
    return ans
}
// submission codes end
