package tests

import (
	"math"
	"testing"
)

/**
 * [507] Perfect Number
 *
 * We define the Perfect Number is a positive integer that is equal to the sum of all its positive divisors except itself.
 *
 * Now, given an integer n, write a function that returns true when it is a perfect number and false when it is not.
 *
 *
 * Example:
 *
 * Input: 28
 * Output: True
 * Explanation: 28 = 1 + 2 + 4 + 7 + 14
 *
 *
 *
 * Note:
 * The input number n will not exceed 100,000,000. (1e8)
 *
 */

func TestPerfectNumber(t *testing.T) {
	var cases = []struct {
		input  int
		output bool
	}{
		{
			input:  28,
			output: true,
		},
	}
	for _, c := range cases {
		x := checkPerfectNumber(c.input)
		if x != c.output {
			t.Fail()
		}
	}
}

// submission codes start here

func checkPerfectNumber(num int) bool {
	if num <= 2 {
		return false
	}
	sum := 1
	root := int(math.Sqrt(float64(num)))
	for i := 2; i <= root; i++ {
		if num%i == 0 {
			sum += i + (num/i)
		}
	}
	return sum == num
}

// submission codes end
