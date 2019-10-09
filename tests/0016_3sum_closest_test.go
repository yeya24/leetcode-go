package tests

import (
    "math"
    "sort"
    "testing"
)

/**
 * [16] 3Sum Closest
 *
 * Given an array nums of n integers and an integer target, find three integers in nums such that the sum is closest to target. Return the sum of the three integers. You may assume that each input would have exactly one solution.
 * 
 * Example:
 * 
 * 
 * Given array nums = [-1, 2, 1, -4], and target = 1.
 * 
 * The sum that is closest to the target is 2. (-1 + 2 + 1 = 2).
 * 
 * 
 */

func Test3SumClosest(t *testing.T) {
    var cases = []struct {
        input  []int
        target int
        output int
    }{
        {
            input:  []int{-1, 2, 1, -4},
            target: 1,
            output: 2,
        },
    }
    for _, c := range cases {
        x := threeSumClosest(c.input, c.target)
        if x != c.output {
            t.Fail()
        }
    }
}

// submission codes start here

func threeSumClosest(nums []int, target int) int {
    sort.Ints(nums)
    closest := math.MaxInt64
    v := 0
    for i := 0; i < len(nums); i++ {
        if i > 0 && nums[i] == nums[i-1] {
            continue
        }
        t := target - nums[i]
        left := i+1
        right := len(nums)-1
        for left < right {
            sum := nums[left] + nums[right]
            if sum == t {
                return target
            } else if sum < t {
                left++
            } else {
                right--
            }
            g := diff(sum, t)
            if g < closest {
                closest = g
                v = sum + nums[i]
            }
        }
    }
    return v
}

// submission codes end
