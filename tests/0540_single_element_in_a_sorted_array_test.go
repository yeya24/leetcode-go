package tests

import (
    "testing"
)

/**
 * [540] Single Element in a Sorted Array
 *
 * Given a sorted array consisting of only integers where every element appears exactly twice except for one element which appears exactly once. Find this single element that appears only once.
 * 
 *  
 * 
 * Example 1:
 * 
 * 
 * Input: [1,1,2,3,3,4,4,8,8]
 * Output: 2
 * 
 * 
 * Example 2:
 * 
 * 
 * Input: [3,3,7,7,10,11,11]
 * Output: 10
 * 
 * 
 *  
 * 
 * Note: Your solution should run in O(log n) time and O(1) space.
 * 
 */

func TestSingleElementinaSortedArray(t *testing.T) {
    var cases = []struct {
        input  []int
        output int
    }{
        {
            input:  []int{1,1,2,3,3,4,4,8,8},
            output: 2,
        },
    }
    for _, c := range cases {
        x := singleNonDuplicate(c.input)
        if x != c.output {
            t.Fail()
        }
    }
}

// submission codes start here

func singleNonDuplicate(nums []int) int {
   x := 0
   for _, v := range nums {
       x ^= v
   }
   return x
}

// submission codes end
