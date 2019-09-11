package tests

import (
	"testing"
)

/**
 * [263] Ugly Number
 *
 * Write a program to check whether a given number is an ugly number.
 *
 * Ugly numbers are positive numbers whose prime factors only include 2, 3, 5.
 *
 * Example 1:
 *
 *
 * Input: 6
 * Output: true
 * Explanation: 6 = 2 * 3
 *
 * Example 2:
 *
 *
 * Input: 8
 * Output: true
 * Explanation: 8 = 2 * 2 * 2
 *
 *
 * Example 3:
 *
 *
 * Input: 14
 * Output: false
 * Explanation: 14 is not ugly since it includes another prime factor 7.
 *
 *
 * Note:
 *
 *
 * 	1 is typically treated as an ugly number.
 * 	Input is within the 32-bit signed integer range: [-2^31,  2^31 - 1].
 *
 */

func TestUglyNumber(t *testing.T) {
	var cases = []struct {
		input  int
		output bool
	}{
		{
			input: 6,
			output: true,
		},
		{
			input: 8,
			output: true,
		},
		{
			input: 14,
			output: false,
		},
	}
	for _, c := range cases {
		x := isUgly(c.input)
		if x != c.output {
			t.Fail()
		}
	}
}

// submission codes start here

func isUgly(num int) bool {
	if num == 1 {
		return true
	}
	if num < 1 {
		return false
	}
	num = reduce(num, 2)
	num = reduce(num, 3)
	num = reduce(num, 5)
	return num == 1
}

func reduce(num, factor int) int {
	for num > 0 && num%factor == 0 {
		num = num / factor
	}
	return num
}

// submission codes end
