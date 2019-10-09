package tests

import (
    "fmt"
    "testing"
)

/**
 * [77] Combinations
 *
 * Given two integers n and k, return all possible combinations of k numbers out of 1 ... n.
 * 
 * Example:
 * 
 * 
 * Input: n = 4, k = 2
 * Output:
 * [
 *   [2,4],
 *   [3,4],
 *   [2,3],
 *   [1,2],
 *   [1,3],
 *   [1,4],
 * ]
 * 
 * 
 */

func TestCombinations(t *testing.T) {
    var cases = []struct {
        n  int
        k  int
        output [][]int
    }{
        {
            n: 4,
            k: 2,
            output: [][]int{{2,4},{3,4},{2,3},{1,2},{1,3},{1,4}},
        },
    }
    for _, c := range cases {
        x := combine(c.n, c.k)
        fmt.Println(x)
        //if !reflect.DeepEqual(x, c.output) {
        //    t.Fail()
        //}
    }
}

// submission codes start here

// This way uses dfs.
func combine(n int, k int) [][]int {
    var cur []int
    var res [][]int
    dfsCombine(&res, cur, n, k, 1)
    return res
}

func dfsCombine(res *[][]int, cur []int, n, k, s int) {
    for k == 0 {
        *res = append(*res, cur)
        return
    }

    for i := s; i <= n - k + 1; i++ {
        cur = append(cur, i)
        dfsCombine(res, append([]int{}, cur...), n, k - 1, i+1)
        cur = cur[:len(cur)-1]
    }
}

// submission codes end
