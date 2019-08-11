package tests

import (
    "testing"
)

/**
 * [136] Single Number
 *
 * Given a non-empty array of integers, every element appears twice except for one. Find that single one.
 * 
 * Note:
 * 
 * Your algorithm should have a linear runtime complexity. Could you implement it without using extra memory?
 * 
 * Example 1:
 * 
 * 
 * Input: [2,2,1]
 * Output: 1
 * 
 * 
 * Example 2:
 * 
 * 
 * Input: [4,1,2,1,2]
 * Output: 4
 * 
 * 
 */

func TestSingleNumber(t *testing.T) {
    var cases = []struct {
        input  []int
        output int
    }{
        {
            input:  []int{2,2,1},
            output: 1,
        },
        {
            input:  []int{4,2,2,1,1},
            output: 4,
        },
    }
    for _, c := range cases {
        x := singleNumber(c.input)
        if x != c.output {
            t.Fail()
        }
    }
}

// submission codes start here

func singleNumber(nums []int) int {
   l := 0
   for _, n := range nums {
       l = l ^ n
   }
   return l
}

// submission codes end
