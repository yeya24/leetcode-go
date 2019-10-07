package tests

import (
	"math"
	"testing"
)

/**
 * [1307] Ugly Number III
 *
 * Write a program to find the n-th ugly number.
 * Ugly numbers are positive integers which are divisible by a or b or c.
 *
 * Example 1:
 *
 * Input: n = 3, a = 2, b = 3, c = 5
 * Output: 4
 * Explanation: The ugly numbers are 2, 3, 4, 5, 6, 8, 9, 10... The 3rd is 4.
 * Example 2:
 *
 * Input: n = 4, a = 2, b = 3, c = 4
 * Output: 6
 * Explanation: The ugly numbers are 2, 3, 4, 6, 8, 9, 10, 12... The 4th is 6.
 *
 * Example 3:
 *
 * Input: n = 5, a = 2, b = 11, c = 13
 * Output: 10
 * Explanation: The ugly numbers are 2, 4, 6, 8, 10, 11, 12, 13... The 5th is 10.
 *
 * Example 4:
 *
 * Input: n = 1000000000, a = 2, b = 217983653, c = 336916467
 * Output: 1999999984
 *
 *
 * Constraints:
 *
 * 	1 <= n, a, b, c <= 10^9
 * 	1 <= a * b * c <= 10^18
 * 	It's guaranteed that the result will be in range [1, 2 * 10^9]
 *
 *
 */

func TestUglyNumberIII(t *testing.T) {
	var cases = []struct {
		n, a, b, c int
		output     int
	}{
		{
			n:      3,
			a:      2,
			b:      3,
			c:      5,
			output: 4,
		},
		{
			n:      4,
			a:      2,
			b:      3,
			c:      4,
			output: 6,
		},
		{
			n:      5,
			a:      2,
			b:      11,
			c:      13,
			output: 10,
		},
		{
			n:      1000000000,
			a:      2,
			b:      217983653,
			c:      336916467,
			output: 1999999984,
		},
	}
	for _, c := range cases {
		x := nthUglyNumberIII(c.n, c.a, c.b, c.c)
		if x != c.output {
			t.Fail()
		}
	}
}

// submission codes start here

func nthUglyNumberIII(n int, a int, b int, c int) int {
	var left, right, ab, ac, bc, abc int64
	left = 1
	right = math.MaxInt32
	ab = int64(a / gcd(a, b) * b)
	ac = int64(a / gcd(a, c) * c)
	bc = int64(b / gcd(b, c) * c)
	abc = ab / int64(gcd(int(ab), c)) * int64(c)
	for left < right {
		mid := left + (right-left)/2
		count := mid/int64(a) + mid/int64(b) + mid/int64(c) - mid/ab - mid/bc - mid/ac + mid/abc
		if count >= int64(n) {
			right = mid
		} else {
			left = mid + 1
		}
	}
	return int(left)
}

// submission codes end
