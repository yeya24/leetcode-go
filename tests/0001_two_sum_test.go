package tests

import (
    "testing"
)

/**
 * [1] Two Sum
 *
 * Given an array of integers, return indices of the two numbers such that they add up to a specific target.
 * 
 * You may assume that each input would have exactly one solution, and you may not use the same element twice.
 * 
 * Example:
 * 
 * 
 * Given nums = [2, 7, 11, 15], target = 9,
 * 
 * Because nums[0] + nums[1] = 2 + 7 = 9,
 * return [0, 1].
 * 
 * 
 */

func TestTwoSum(t *testing.T) {
    nums := []int{2, 7, 11, 15}
    output := twoSum(nums, 9)
    if output[0] != 0 || output[1] != 1 {
        t.Fail()
    }
}

// submission codes start here

func twoSum(nums []int, target int) []int {
    m := map[int]int{}

    for i, v := range nums {
        if _, ok := m[target-v]; ok {
            return []int{m[target-v], i}
        }
        m[v] = i
    }

    return []int{}
}

// submission codes end
