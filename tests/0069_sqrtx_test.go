package tests

import (
    "testing"
)

/**
 * [69] Sqrt(x)
 *
 * Implement int sqrt(int x).
 * 
 * Compute and return the square root of x, where x is guaranteed to be a non-negative integer.
 * 
 * Since the return type is an integer, the decimal digits are truncated and only the integer part of the result is returned.
 * 
 * Example 1:
 * 
 * 
 * Input: 4
 * Output: 2
 * 
 * 
 * Example 2:
 * 
 * 
 * Input: 8
 * Output: 2
 * Explanation: The square root of 8 is 2.82842..., and since 
 *              the decimal part is truncated, 2 is returned.
 * 
 * 
 */

func TestSqrt(t *testing.T) {
    var cases = []struct {
        input  int
        output int
    }{
        {
            input:  8,
            output: 2,
        },
        {
            input:  4,
            output: 2,
        },
    }
    for _, c := range cases {
        x := mySqrt(c.input)
        if x != c.output {
            t.Fail()
        }
    }
}

// submission codes start here

func mySqrt(x int) int {
    i := 0
    j := x
    for i <= j {
        m := (i+j)/2
        if m * m <= x && (m+1) * (m+1) > x {
            return m
        } else if m * m > x {
            j = m - 1
        } else {
            i = m + 1
        }
    }
    return -1
}

// submission codes end
