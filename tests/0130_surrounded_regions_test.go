package tests

import (
    "reflect"
    "testing"
)

/**
 * [130] Surrounded Regions
 *
 * Given a 2D board containing 'X' and 'O' (the letter O), capture all regions surrounded by 'X'.
 * 
 * A region is captured by flipping all 'O's into 'X's in that surrounded region.
 * 
 * Example:
 * 
 * 
 * X X X X
 * X O O X
 * X X O X
 * X O X X
 * 
 * 
 * After running your function, the board should be:
 * 
 * 
 * X X X X
 * X X X X
 * X X X X
 * X O X X
 * 
 * 
 * Explanation:
 * 
 * Surrounded regions shouldn&rsquo;t be on the border, which means that any 'O' on the border of the board are not flipped to 'X'. Any 'O' that is not on the border and it is not connected to an 'O' on the border will be flipped to 'X'. Two cells are connected if they are adjacent cells connected horizontally or vertically.
 * 
 */

func TestSurroundedRegions(t *testing.T) {
    var cases = []struct {
        input  [][]byte
        output [][]byte
    }{
        {
           input:  [][]byte{[]byte("XXXX"),[]byte("XOOX"),[]byte("XXOX"),[]byte("XOXX")},
           output: [][]byte{[]byte("XXXX"),[]byte("XXXX"),[]byte("XXXX"),[]byte("XOXX")},
        },
        {
            input: [][]byte{[]byte("XOXOXO"),[]byte("OXOXOX"),[]byte("XOXOXO"),[]byte("OXOXOX")},
            output: [][]byte{[]byte("XOXOXO"),[]byte("OXXXXX"),[]byte("XXXXXO"),[]byte("OXOXOX")},
        },
    }
    for _, c := range cases {
        solve(c.input)
        if !reflect.DeepEqual(c.input, c.output) {
            t.Fail()
        }
    }
}

// submission codes start here

func solve(board [][]byte)  {
    m := len(board)
    if m == 0 {
        return
    }
    n := len(board[0])
    if n == 0 {
        return
    }
    for i := 0; i < m; i ++ {
        dfsSolve(board, m, n, i, 0)
        dfsSolve(board, m, n, i, n-1)
    }
    for i := 0; i < n; i ++ {
        dfsSolve(board, m, n, 0, i)
        dfsSolve(board, m, n, m-1, i)
    }
    for i := 0; i < m; i++ {
        for j :=0; j < n; j++ {
            switch board[i][j]{
            case 'G':
                board[i][j] = 'O'
            case 'O':
                board[i][j] = 'X'
            }
        }
    }
}

func dfsSolve(board [][]byte, m, n, x, y int) {
    if x < 0 || x >= m || y < 0 || y >= n || board[x][y] != 'O' {
        return
    }
    board[x][y] = 'G'
    dfsSolve(board,m,n,x+1,y)
    dfsSolve(board,m,n,x-1,y)
    dfsSolve(board,m,n,x,y+1)
    dfsSolve(board,m,n,x,y-1)
}

// submission codes end
