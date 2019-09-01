package tests

import (
    "testing"
)

/**
 * [461] Hamming Distance
 *
 * The Hamming distance between two integers is the number of positions at which the corresponding bits are different.
 * 
 * Given two integers x and y, calculate the Hamming distance.
 * 
 * Note:
 * 0 <= x, y < 2^31.
 * 
 * 
 * Example:
 * 
 * Input: x = 1, y = 4
 * 
 * Output: 2
 * 
 * Explanation:
 * 1   (0 0 0 1)
 * 4   (0 1 0 0)
 *        &uarr;   &uarr;
 * 
 * The above arrows point to positions where the corresponding bits are different.
 * 
 * 
 */

func TestHammingDistance(t *testing.T) {
    var cases = []struct {
        inputA  int
        inputB  int
        output int
    }{
        {
        	inputA: 1,
        	inputB: 4,
            output: 2,
        },
    }
    for _, c := range cases {
        x := hammingDistance(c.inputA, c.inputB)
        if x != c.output {
            t.Fail()
        }
    }
}

// submission codes start here

func hammingDistance(x int, y int) int {
    a := x^y
    b := []byte{}
    var c int
    for a > 0 {
        c = a % 2
        a = a / 2
        b = append([]byte{byte(c)}, b...)
    }
    var count int
    for _, v := range b {
       if v == 1 {
           count++
       }
    }
    return count
}

// submission codes end
