package tests

import (
    "fmt"
    "sort"
    "testing"
)

/**
 * [40] Combination Sum II
 *
 * Given a collection of candidate numbers (candidates) and a target number (target), find all unique combinations in candidates where the candidate numbers sums to target.
 * 
 * Each number in candidates may only be used once in the combination.
 * 
 * Note:
 * 
 * 
 * 	All numbers (including target) will be positive integers.
 * 	The solution set must not contain duplicate combinations.
 * 
 * 
 * Example 1:
 * 
 * 
 * Input: candidates = [10,1,2,7,6,1,5], target = 8,
 * A solution set is:
 * [
 *   [1, 7],
 *   [1, 2, 5],
 *   [2, 6],
 *   [1, 1, 6]
 * ]
 * 
 * 
 * Example 2:
 * 
 * 
 * Input: candidates = [2,5,2,1,2], target = 5,
 * A solution set is:
 * [
 *   [1,2,2],
 *   [5]
 * ]
 * 
 * 
 */

func TestCombinationSumII(t *testing.T) {
    var cases = []struct {
        candidates  []int
        target int
        output [][]int
    }{
        {
            candidates: []int{10,1,2,7,6,1,5},
            target: 8,
            output: [][]int{{1,7},{1,2,5},{2,6},{1,1,6}},
        },
    }
    for _, c := range cases {
        x := combinationSum2(c.candidates, c.target)
        fmt.Println(x)
        //if !reflect.DeepEqual(x, c.output) {
        //    t.Fail()
        //}
    }
}

// submission codes start here

func combinationSum2(candidates []int, target int) [][]int {
    var ans [][]int
    var cur []int
    sort.Ints(candidates)
    dfsCombinationSum2(candidates, cur, target, 0, &ans)
    return ans
}

func dfsCombinationSum2(candidates, cur []int, target, s int, ans *[][]int) {
    if target == 0 {
        *ans = append(*ans, cur)
        return
    }

    for i := s; i < len(candidates); i++ {
        if i > s && candidates[i] == candidates[i-1] {
            continue
        }
        cur = append(cur, candidates[i])
        v := target - candidates[i]
        if v >= 0 {
            dfsCombinationSum2(candidates, append([]int{}, cur...), v, i+1, ans)
        }
        cur = cur[:len(cur)-1]
    }
}

// submission codes end
