package tests

import (
    "math"
    "testing"
)

/**
 * [899] Binary Gap
 *
 * Given a positive integer N, find and return the longest distance between two consecutive 1's in the binary representation of N.
 * 
 * If there aren't two consecutive 1's, return <font face="monospace">0</font>.
 * 
 *  
 * 
 * 
 * 
 * 
 * 
 * 
 * 
 * 
 * 
 * 
 * 
 * Example 1:
 * 
 * 
 * Input: 22
 * Output: 2
 * Explanation: 
 * 22 in binary is 0b10110.
 * In the binary representation of 22, there are three ones, and two consecutive pairs of 1's.
 * The first consecutive pair of 1's have distance 2.
 * The second consecutive pair of 1's have distance 1.
 * The answer is the largest of these two distances, which is 2.
 * 
 * 
 * 
 * Example 2:
 * 
 * 
 * Input: 5
 * Output: 2
 * Explanation: 
 * 5 in binary is 0b101.
 * 
 * 
 * 
 * Example 3:
 * 
 * 
 * Input: 6
 * Output: 1
 * Explanation: 
 * 6 in binary is 0b110.
 * 
 * 
 * 
 * Example 4:
 * 
 * 
 * Input: 8
 * Output: 0
 * Explanation: 
 * 8 in binary is 0b1000.
 * There aren't any consecutive pairs of 1's in the binary representation of 8, so we return 0.
 * 
 * 
 *  
 * 
 * 
 * 
 * 
 * Note:
 * 
 * 
 * 	1 <= N <= 10^9
 * 
 * 
 * 
 * 
 * 
 * 
 * 
 * 
 * 
 */

func TestBinaryGap(t *testing.T) {
    var cases = []struct {
        input  int
        output int
    }{
        {
            input:  22,
            output: 2,
        },
        {
            input:  5,
            output: 2,
        },
        {
            input:  6,
            output: 1,
        },
        {
            input:  8,
            output: 0,
        },
    }
    for _, c := range cases {
        x := binaryGap(c.input)
        if x != c.output {
            t.Fail()
        }
    }
}

// submission codes start here

func binaryGap(N int) int {
    var res int
    for d := -32; N > 0 ; d++ {
        if N % 2 == 1 {
            res = int(math.Max(float64(res), float64(d)))
            d = 0
        }
        N /= 2
    }
    return res
}

// submission codes end
