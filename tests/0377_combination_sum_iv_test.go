package tests

import (
    "fmt"
    "testing"
)

/**
 * [377] Combination Sum IV
 *
 * Given an integer array with all positive numbers and no duplicates, find the number of possible combinations that add up to a positive integer target.
 * 
 * Example:
 * 
 * 
 * nums = [1, 2, 3]
 * target = 4
 * 
 * The possible combination ways are:
 * (1, 1, 1, 1)
 * (1, 1, 2)
 * (1, 2, 1)
 * (1, 3)
 * (2, 1, 1)
 * (2, 2)
 * (3, 1)
 * 
 * Note that different sequences are counted as different combinations.
 * 
 * Therefore the output is 7.
 * 
 * 
 *  
 * 
 * Follow up:
 * What if negative numbers are allowed in the given array?
 * How does it change the problem?
 * What limitation we need to add to the question to allow negative numbers?
 * 
 * Credits:
 * Special thanks to @pbrother for adding this problem and creating all test cases.
 * 
 */

func TestCombinationSumIV(t *testing.T) {
    var cases = []struct {
        nums   []int
        target  int
        output int
    }{
        {
           nums: []int{1,2,3},
           target: 4,
           output: 7,
        },
        {
            nums: []int{4,2,1},
            target: 32,
            output: 39882198,
        },
    }
    for _, c := range cases {
        x := combinationSum4DP(c.nums, c.target)
        fmt.Println(x)
        if x != c.output {
           t.Fail()
        }
    }
}

// submission codes start here

// Use recursion.
func combinationSum4(nums []int, target int) int {
    m := make([]int, target+1)
    for i := 1; i < target + 1; i++ {
        m[i] = -1
    }
    m[0] = 1
    return helperCombinationSum4(nums, m, target)
}

func helperCombinationSum4(nums, m []int, target int) int {
    if target < 0 {
        return 0
    }
    if m[target] != -1 {
        return m[target]
    }
    ans := 0
    for _, v := range nums {
        ans += helperCombinationSum4(nums, m, target - v)
    }
    m[target] = ans
    return ans
}

// Use dp
func combinationSum4DP(nums []int, target int) int {
    m := make([]int, target+1)
    m[0] = 1
    if target < 0 {
        return 0
    }
    for i := 0; i <= target; i++ {
        for _, v := range nums {
            if i - v >= 0 {
                m[i] += m[i - v]
            }
        }
    }
    return m[target]
}

// submission codes end
