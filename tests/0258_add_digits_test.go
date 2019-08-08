package tests

import (
	"testing"
)

/**
 * [258] Add Digits
 *
 * Given a non-negative integer num, repeatedly add all its digits until the result has only one digit.
 *
 * Example:
 *
 *
 * Input: 38
 * Output: 2
 * Explanation: The process is like: 3 + 8 = 11, 1 + 1 = 2.
 *              Since 2 has only one digit, return it.
 *
 *
 * Follow up:
 * Could you do it without any loop/recursion in O(1) runtime?
 */

func TestAddDigits(t *testing.T) {
	cases := []struct {
		input  int
		expect int
	}{
		{
			input:  38,
			expect: 2,
		},
		{
			input:  9,
			expect: 9,
		},
		{
			input:  103,
			expect: 4,
		},
	}
	for _, c := range cases {
		if addDigits(c.input) != c.expect {
			t.Fail()
		}
	}
}

// submission codes start here

func addDigits(num int) int {
	return (num-1)%9 + 1
}

// submission codes end
