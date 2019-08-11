package tests

import (
    "testing"
)

/**
 * [217] Contains Duplicate
 *
 * Given an array of integers, find if the array contains any duplicates.
 * 
 * Your function should return true if any value appears at least twice in the array, and it should return false if every element is distinct.
 * 
 * Example 1:
 * 
 * 
 * Input: [1,2,3,1]
 * Output: true
 * 
 * Example 2:
 * 
 * 
 * Input: [1,2,3,4]
 * Output: false
 * 
 * Example 3:
 * 
 * 
 * Input: [1,1,1,3,3,4,3,2,4,2]
 * Output: true
 * 
 */

func TestContainsDuplicate(t *testing.T) {
    var cases = []struct {
        input  []int
        output bool
    }{
        {
            input:  []int{1,2,3,1},
            output: true,
        },
        {
            input:  []int{1,2,3},
            output: false,
        },
    }
    for _, c := range cases {
        x := containsDuplicate(c.input)
        if x != c.output {
            t.Fail()
        }
    }
}

// submission codes start here

func containsDuplicate(nums []int) bool {
    m := make(map[int]struct{})
    for _, n := range nums {
        if _, ok := m[n]; ok {
            return true
        } else {
            m[n] = struct{}{}
        }
    }
    return false
}

// submission codes end
