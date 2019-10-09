package tests

import (
    "fmt"
    "testing"
)

/**
 * [39] Combination Sum
 *
 * Given a set of candidate numbers (candidates) (without duplicates) and a target number (target), find all unique combinations in candidates where the candidate numbers sums to target.
 * 
 * The same repeated number may be chosen from candidates unlimited number of times.
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
 * Input: candidates = [2,3,6,7], target = 7,
 * A solution set is:
 * [
 *   [7],
 *   [2,2,3]
 * ]
 * 
 * 
 * Example 2:
 * 
 * 
 * Input: candidates = [2,3,5], target = 8,
 * A solution set is:
 * [
 *   [2,2,2,2],
 *   [2,3,3],
 *   [3,5]
 * ]
 * 
 * 
 */

func TestCombinationSum(t *testing.T) {
    var cases = []struct {
        candidates  []int
        target int
        output [][]int
    }{
        {
            candidates: []int{2,3,6,7},
            target: 7,
            output: [][]int{{2,2,3},{7}},
        },
    }
    for _, c := range cases {
        x := combinationSum(c.candidates, c.target)
        fmt.Println(x)
        //if !reflect.DeepEqual(x, c.output) {
        //    t.Fail()
        //}
    }
}

// submission codes start here

func combinationSum(candidates []int, target int) [][]int {
    var ans [][]int
    var cur []int
    dfsCombinationSum(candidates, cur, target, 0, &ans)
    return ans
}

func dfsCombinationSum(candidates, cur []int, target, s int, ans *[][]int) {
    if target == 0 {
        *ans = append(*ans, cur)
        return
    }

    for i := s; i < len(candidates); i++ {
        cur = append(cur, candidates[i])
        v := target - candidates[i]
        if v >= 0 {
            dfsCombinationSum(candidates, append([]int{}, cur...), v, i, ans)
        }
        cur = cur[:len(cur)-1]
    }
}

// submission codes end
