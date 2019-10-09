package tests

import (
    "fmt"
    "sort"
    "testing"
)

/**
 * [90] Subsets II
 *
 * Given a collection of integers that might contain duplicates, nums, return all possible subsets (the power set).
 * 
 * Note: The solution set must not contain duplicate subsets.
 * 
 * Example:
 * 
 * 
 * Input: [1,2,2]
 * Output:
 * [
 *   [2],
 *   [1],
 *   [1,2,2],
 *   [2,2],
 *   [1,2],
 *   []
 * ]
 * 
 * 
 */

func TestSubsetsII(t *testing.T) {
    var cases = []struct {
        input  []int
        output [][]int
    }{
        {
            input:  []int{1, 2, 2},
            output: [][]int{{},{1},{2},{1,2},{2,2},{1,2,2}},
        },
    }
    for _, c := range cases {
        x := subsetsWithDup(c.input)
        fmt.Println(x)
    }
}

// submission codes start here

func subsetsWithDup(nums []int) [][]int {
    var res [][]int
    var cur []int
    used := make([]bool, len(nums))
    sort.Ints(nums)
    for i := 0; i <= len(nums); i++ {
        dfsSubsetsWithDup(nums, cur, used, i, 0, &res)
    }
    return res
}

func dfsSubsetsWithDup(nums, cur []int, used []bool, n, s int, res *[][]int) {
    if len(cur) == n {
        *res = append(*res, cur)
        return
    }

    for i := s; i< len(nums); i++ {
        if used[i] || i > 0 && nums[i] == nums[i-1] && !used[i-1] {
            continue
        }
        used[i] = true
        cur = append(cur, nums[i])
        dfsSubsetsWithDup(nums, append([]int{}, cur...), used, n, i+1, res)
        cur = cur[:len(cur)-1]
        used[i] = false
    }
}

// submission codes end
