package tests

import (
    "testing"
)

/**
 * [367] Valid Perfect Square
 *
 * Given a positive integer num, write a function which returns True if num is a perfect square else False.
 * 
 * Note: Do not use any built-in library function such as sqrt.
 * 
 * Example 1:
 * 
 * 
 * 
 * Input: 16
 * Output: true
 * 
 * 
 * 
 * Example 2:
 * 
 * 
 * Input: 14
 * Output: false
 * 
 * 
 * 
 */

func TestValidPerfectSquare(t *testing.T) {
    var cases = []struct {
        input  int
        output bool
    }{
        {
            input:  16,
            output: true,
        },
        {
            input:  14,
            output: false,
        },
    }
    for _, c := range cases {
        x := isPerfectSquare(c.input)
        if x != c.output {
            t.Fail()
        }
    }
}

// submission codes start here

func isPerfectSquare(num int) bool {
   i, sum := 1, 0
   for sum < num {
       sum += i
       i += 2
   }
   if sum == num {
       return true
   }
   return false
}

// submission codes end
