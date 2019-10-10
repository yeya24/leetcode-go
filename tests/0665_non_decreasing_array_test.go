package tests

import (
    "testing"
)

/**
 * [665] Non-decreasing Array
 *
 * 
 * Given an array with n integers, your task is to check if it could become non-decreasing by modifying at most 1 element.
 * 
 * 
 * 
 * We define an array is non-decreasing if array[i] <= array[i + 1] holds for every i (1 <= i < n).
 * 
 * 
 * Example 1:
 * 
 * Input: [4,2,3]
 * Output: True
 * Explanation: You could modify the first 4 to 1 to get a non-decreasing array.
 * 
 * 
 * 
 * Example 2:
 * 
 * Input: [4,2,1]
 * Output: False
 * Explanation: You can't get a non-decreasing array by modify at most one element.
 * 
 * 
 * 
 * Note:
 * The n belongs to [1, 10,000].
 * 
 */

func TestNonDecreasingArray(t *testing.T) {
    var cases = []struct {
        input  []int
        output bool
    }{
        {
            input:  []int{4,2,3},
            output: true,
        },
        {
            input:  []int{4,2,1},
            output: false,
        },
    }
    for _, c := range cases {
        x := checkPossibility(c.input)
        if x != c.output {
            t.Fail()
        }
    }
}

// submission codes start here

// brute force. time O(n^2)
func checkPossibility(nums []int) bool {
	isMono := func (nums []int) bool {
        for i := 0; i < len(nums)-1; i++ {
            if nums[i] > nums[i+1] {
                return false
            }
        }
        return true
    }
    for i, _ := range nums {
        if i == 0 {
            if isMono(nums[1:]) {
                return true
            }
        } else if i == len(nums) - 1 {
            if isMono(nums[:i]) {
                return true
            }
        } else {
            if isMono(nums[:i]) && isMono(nums[i+1:]) && nums[i-1] <= nums[i+1] {
                return true
            }
        }
    }
    return false
}

// time O(n)
func checkPossibility2(nums []int) bool {
    if len(nums) <= 1 {
        return true
    }
    cnt := 0
    for i := 1; i < len(nums); i++ {
        if nums[i] < nums[i-1] {
            cnt++
            if cnt > 1 {
                return false
            }
            if i-2<0 || nums[i-2] <= nums[i] {
                nums[i-1] = nums[i]
            } else {
                nums[i] = nums[i-1]
            }
        }

    }
    return true
}

// submission codes end
