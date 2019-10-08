package tests

import (
    "fmt"
    "reflect"
    "testing"
)

/**
 * [78] Subsets
 *
 * Given a set of distinct integers, nums, return all possible subsets (the power set).
 * 
 * Note: The solution set must not contain duplicate subsets.
 * 
 * Example:
 * 
 * 
 * Input: nums = [1,2,3]
 * Output:
 * [
 *   [3],
 *   [1],
 *   [2],
 *   [1,2,3],
 *   [1,3],
 *   [2,3],
 *   [1,2],
 *   []
 * ]
 * 
 */

func TestSubsets(t *testing.T) {
    var cases = []struct {
        input  []int
        output [][]int
    }{
        {
            input:  []int{1, 2, 3},
            output: [][]int{{3},{2},{1},{1,2,3},{1,3},{1,2},{2,3},{}},
        },
    }
    for _, c := range cases {
        x := subsets(c.input)
        fmt.Println(x)
        if !reflect.DeepEqual(x, c.output) {
            t.Fail()
        }
    }
}

// submission codes start here

func subsets(nums []int) [][]int {
    if len(nums) == 0 {
        return nil
    }
    var cur []int
    var ans [][]int
    for i := 0; i<= len(nums); i++ {
        dfsSubsets(nums, i, 0, cur, &ans)
    }
    return ans
}

func dfsSubsets(nums []int, n, s int, cur []int, ans *[][]int) {
    if len(cur) == n {
        *ans = append(*ans, cur)
        return
    }
    for i := s; i < len(nums); i++ {
        cur = append(cur, nums[i])
        dfsSubsets(nums, n, i+1, append([]int{}, cur...), ans)
        cur = cur[:len(cur)-1]
    }
}

// submission codes end
