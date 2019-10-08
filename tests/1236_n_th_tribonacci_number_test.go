package tests

import (
    "testing"
)

/**
 * [1236] N-th Tribonacci Number
 *
 * The Tribonacci sequence Tn is defined as follows: 
 * 
 * T0 = 0, T1 = 1, T2 = 1, and Tn+3 = Tn + Tn+1 + Tn+2 for n >= 0.
 * 
 * Given n, return the value of Tn.
 * 
 *  
 * Example 1:
 * 
 * 
 * Input: n = 4
 * Output: 4
 * Explanation:
 * T_3 = 0 + 1 + 1 = 2
 * T_4 = 1 + 1 + 2 = 4
 * 
 * 
 * Example 2:
 * 
 * 
 * Input: n = 25
 * Output: 1389537
 * 
 * 
 *  
 * Constraints:
 * 
 * 
 * 	0 <= n <= 37
 * 	The answer is guaranteed to fit within a 32-bit integer, ie. answer <= 2^31 - 1.
 * 
 */

func TestNthTribonacciNumber(t *testing.T) {
    var cases = []struct {
        input  int
        output int
    }{
        {
            input:  4,
            output: 4,
        },
        {
            input:  25,
            output: 1389537,
        },
    }
    for _, c := range cases {
        x := tribonacci(c.input)
        if x != c.output {
            t.Fail()
        }
    }
}

// submission codes start here

func tribonacci(n int) int {
    if n <= 1 {
        return n
    }
    if n == 2 {
        return 1
    }
    m := make([]int, n+1)
    m[0], m[1], m[2] = 0, 1, 1
    for i := 3; i <=n; i++ {
        m[i] = m[i-1] + m[i-2] + m[i-3]
    }
    return m[n]
}

// submission codes end
