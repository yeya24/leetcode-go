package tests

import (
	"testing"
)

/**
 * [204] Count Primes
 *
 * Count the number of prime numbers less than a non-negative number, n.
 *
 * Example:
 *
 *
 * Input: 10
 * Output: 4
 * Explanation: There are 4 prime numbers less than 10, they are 2, 3, 5, 7.
 *
 *
 */

func TestCountPrimes(t *testing.T) {
	var cases = []struct {
		input  int
		output int
	}{
		{
			input:  10,
			output: 4,
		},
	}
	for _, c := range cases {
		x := countPrimes(c.input)
		if x != c.output {
			t.Fail()
		}
	}
}

// submission codes start here

func countPrimes(n int) int {
	primes := make([]bool, n)
	for i := range primes {
		if i == 0 || i == 1 {
			primes[i] = false
		} else {
			primes[i] = true
		}
	}
	p := 2
	for p * p <= n {
	    if primes[p] == true {
	        for i := p * p; i < n ; i+=p {
	            primes[i] = false
            }
        }
	    p++
    }
	res := 0
	for _, v := range primes {
	    if v == true {
	        res++
        }
    }
	return res
}

// submission codes end
