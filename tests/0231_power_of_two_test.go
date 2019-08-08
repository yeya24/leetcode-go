package tests

import (
    "testing"
)

/**
 * [231] Power of Two
 *
 * Given an integer, write a function to determine if it is a power of two.
 * 
 * Example 1:
 * 
 * 
 * Input: 1
 * Output: true 
 * Explanation: 2^0 = 1
 * 
 * 
 * Example 2:
 * 
 * 
 * Input: 16
 * Output: true
 * Explanation: 2^4 = 16
 * 
 * Example 3:
 * 
 * 
 * Input: 218
 * Output: false
 * 
 */

func TestPowerofTwo(t *testing.T) {
    cases := []struct {
        input  int
        expect bool
    }{
        {
            input:  38,
            expect: false,
        },
        {
            input:  1,
            expect: true,
        },
        {
            input:  8,
            expect: true,
        },
    }
    for _, c := range cases {
        if isPowerOfTwo(c.input) != c.expect {
            t.Fail()
        }
    }
}

// submission codes start here

func isPowerOfTwo(n int) bool {
    return (n > 0) && (n & (n - 1) == 0)
}

// submission codes end
