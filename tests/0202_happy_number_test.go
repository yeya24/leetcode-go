package tests

import (
    "strconv"
    "testing"
)

/**
 * [202] Happy Number
 *
 * Write an algorithm to determine if a number is "happy".
 * 
 * A happy number is a number defined by the following process: Starting with any positive integer, replace the number by the sum of the squares of its digits, and repeat the process until the number equals 1 (where it will stay), or it loops endlessly in a cycle which does not include 1. Those numbers for which this process ends in 1 are happy numbers.
 * 
 * Example: 
 * 
 * 
 * Input: 19
 * Output: true
 * Explanation: 
 * 1^2 + 9^2 = 82
 * 8^2 + 2^2 = 68
 * 6^2 + 8^2 = 100
 * 1^2 + 0^2 + 0^2 = 1
 * 
 */

func TestHappyNumber(t *testing.T) {
    var cases = []struct {
        input  int
        output bool
    }{
        {
            input: 19,
            output: true,
        },
        {
            input: 2,
            output: false,
        },
        {
            input: 4,
            output: false,
        },
    }
    for _, c := range cases {
        x := isHappy(c.input)
        if x != c.output {
            t.Fail()
        }
    }
}

// submission codes start here

func isHappy(n int) bool {
   m := make(map[int]struct{})
   acc := []int{n}
   for len(acc) > 0 {
       v := acc[len(acc)-1]
       acc = acc[:len(acc)-1]
       if v == 1 {
           return true
       } else {
           var sum int
           for _, x := range strconv.Itoa(v) {
               sum += int((x-'0') * (x-'0'))
           }
           if _, ok := m[sum]; !ok {
               m[sum] = struct{}{}
               acc = append(acc, sum)
           }
       }
   }
   return false
}

// submission codes end
