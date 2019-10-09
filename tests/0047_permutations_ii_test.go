package tests

import (
    "reflect"
    "sort"
    "testing"
)

/**
 * [47] Permutations II
 *
 * Given a collection of numbers that might contain duplicates, return all possible unique permutations.
 * 
 * Example:
 * 
 * 
 * Input: [1,1,2]
 * Output:
 * [
 *   [1,1,2],
 *   [1,2,1],
 *   [2,1,1]
 * ]
 * 
 * 
 */

func TestPermutationsII(t *testing.T) {
    var cases = []struct {
        input  []int
        output [][]int
    }{
        {
            input:  []int{1, 1, 2},
            output: [][]int{{1, 1, 2}, {1, 2, 1}, {2, 1, 1}},
        },
    }
    for _, c := range cases {
        x := permuteUnique(c.input)
        if !reflect.DeepEqual(x, c.output) {
            t.Fail()
        }
    }
}

// submission codes start here

func permuteUnique(nums []int) [][]int {
    var res [][]int
    var cur []int
    used := make([]bool, len(nums))
    sort.Ints(nums)
    dfsPermuteUnique(nums, cur, used, &res)
    return res
}

func dfsPermuteUnique(nums, cur []int, used []bool, res *[][]int) {
    if len(cur) == len(nums) {
        *res = append(*res, cur)
        return
    }
    for i := 0; i< len(nums); i++ {
        if used[i] {
            continue
        }
        if i > 0 && nums[i] == nums[i-1] && !used[i-1] {
            continue
        }
        used[i] = true
        cur = append(cur, nums[i])
        dfsPermuteUnique(nums, append([]int{}, cur...), used, res)
        cur = cur[:len(cur)-1]
        used[i] = false
    }
}

// submission codes end
