package tests

import (
	"math"
	"testing"
)

/**
 * [633] Sum of Square Numbers
 *
 * Given a non-negative integer c, your task is to decide whether there're two integers a and b such that a^2 + b^2 = c.
 *
 * Example 1:
 *
 *
 * Input: 5
 * Output: True
 * Explanation: 1 * 1 + 2 * 2 = 5
 *
 *
 *
 *
 * Example 2:
 *
 *
 * Input: 3
 * Output: False
 *
 *
 *
 *
 */

func TestSumofSquareNumbers(t *testing.T) {
	var cases = []struct {
		input  int
		output bool
	}{
		{
			input:  5,
			output: true,
		},
		{
			input:  3,
			output: false,
		},
		{
			input: 2,
			output: true,
		},
	}
	for _, c := range cases {
		x := judgeSquareSum(c.input)
		if x != c.output {
			t.Fail()
		}
	}
}

// submission codes start here

func judgeSquareSum(c int) bool {
	i, j := 0, int(math.Sqrt(float64(c)))
	for i <= j {
		t := i * i + j * j
		if t == c {
			return true
		} else if t > c {
			j--
		} else {
			i++
		}
	}
	return false
}

// submission codes end
