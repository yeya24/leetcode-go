package tests

import (
	"testing"
)

/**
 * [1279] Prime Arrangements
 *
 * Return the number of permutations of 1 to n so that prime numbers are at prime indices (1-indexed.)
 * (Recall that an integer is prime if and only if it is greater than 1, and cannot be written as a product of two positive integers both smaller than it.)
 * Since the answer may be large, return the answer modulo 10^9 + 7.
 *
 * Example 1:
 *
 * Input: n = 5
 * Output: 12
 * Explanation: For example [1,2,5,4,3] is a valid permutation, but [5,2,3,4,1] is not because the prime number 5 is at index 1.
 *
 * Example 2:
 *
 * Input: n = 100
 * Output: 682289015
 *
 *
 * Constraints:
 *
 * 	1 <= n <= 100
 *
 *
 */

func TestPrimeArrangements(t *testing.T) {
	var cases = []struct {
		input  int
		output int
	}{
		{
			input:  5,
			output: 12,
		},
		{
			input:  100,
			output: 682289015,
		},
	}
	for _, c := range cases {
		x := numPrimeArrangements(c.input)
		if x != c.output {
			t.Fail()
		}
	}
}

// submission codes start here

func numPrimeArrangements(n int) int {
	res := 0
	for i := 1; i <= n; i++ {
		if isPrime(i) {
			res++
		}
	}
	return factorial(n-res) % 1000000007 * factorial(res) % 1000000007
}

func isPrime(n int) bool {
	if n == 2 {
		return true
	}
	if n == 1 || n%2 == 0 {
		return false
	}
	for i := 3; i*i <= n; i += 2 {
		if n%i == 0 {
			return false
		}
	}
	return true
}

func factorial(n int) int {
	if n <= 0 {
		return 1
	}
	if n == 1 {
		return 1
	}
	res := 1
	for i := 2; i <= n; i++ {
		res = res * i % 1000000007
	}
	return res
}

// submission codes end
