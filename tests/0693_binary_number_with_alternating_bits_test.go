package tests

import (
	"testing"
)

/**
 * [693] Binary Number with Alternating Bits
 *
 * Given a positive integer, check whether it has alternating bits: namely, if two adjacent bits will always have different values.
 *
 * Example 1:
 *
 * Input: 5
 * Output: True
 * Explanation:
 * The binary representation of 5 is: 101
 *
 *
 *
 * Example 2:
 *
 * Input: 7
 * Output: False
 * Explanation:
 * The binary representation of 7 is: 111.
 *
 *
 *
 * Example 3:
 *
 * Input: 11
 * Output: False
 * Explanation:
 * The binary representation of 11 is: 1011.
 *
 *
 *
 * Example 4:
 *
 * Input: 10
 * Output: True
 * Explanation:
 * The binary representation of 10 is: 1010.
 *
 *
 */

func TestBinaryNumberwithAlternatingBits(t *testing.T) {
	var cases = []struct {
		input  int
		output bool
	}{
		{
			input:  5,
			output: true,
		},
		{
			input:  7,
			output: false,
		},
		{
			input:  11,
			output: false,
		},
		{
			input:  10,
			output: true,
		},
	}
	for _, c := range cases {
		x := hasAlternatingBits(c.input)
		if x != c.output {
			t.Fail()
		}
	}
}

// submission codes start here

func hasAlternatingBits(n int) bool {
	flag := -1
    for n > 0 {
        x := n & 1
        if flag == -1 {
			flag = x
		} else {
			if x == flag {
				return false
			}
			flag = x
		}
        n >>= 1
    }
	return true
}

// submission codes end
