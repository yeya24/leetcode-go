package tests

import (
    "math"
    "testing"
)

/**
 * [7] Reverse Integer
 *
 * Given a 32-bit signed integer, reverse digits of an integer.
 * 
 * Example 1:
 * 
 * 
 * Input: 123
 * Output: 321
 * 
 * 
 * Example 2:
 * 
 * 
 * Input: -123
 * Output: -321
 * 
 * 
 * Example 3:
 * 
 * 
 * Input: 120
 * Output: 21
 * 
 * 
 * Note:
 * Assume we are dealing with an environment which could only store integers within the 32-bit signed integer range: [-2^31,  2^31 - 1]. For the purpose of this problem, assume that your function returns 0 when the reversed integer overflows.
 * 
 */

func TestReverseInteger(t *testing.T) {
    var cases = []struct {
        input  int
        output int
    }{
        {
            input: 123,
            output: 321,
        },
        {
            input: -123,
            output: -321,
        },
        {
            input: 1201,
            output: 1021,
        },
        {
            input:  120,
            output: 21,
        },
    }
    for _, c := range cases {
        x := reverse(c.input)
        if x != c.output {
            t.Fail()
        }
    }
}

// submission codes start here

func reverse(x int) int {
    s := 0
    isLast := true
    isNegative := false
    if x < 0 {
        isNegative = true
        x = -x
    }
    for x > 0 {
        v := x % 10
        if isLast && v != 0 || !isLast {
            s = s * 10 + v % 10
            isLast = false
        }
        x = x / 10
    }
    if isNegative {
        if -s < math.MinInt32 {
            return 0
        }
        return -s
    }
    if s > math.MaxInt32 {
        return 0
    }
    return s
}

// submission codes end
