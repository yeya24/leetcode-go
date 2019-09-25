package tests

import (
	"fmt"
	"testing"
)

/**
 * [264] Ugly Number II
 *
 * Write a program to find the n-th ugly number.
 *
 * Ugly numbers are positive numbers whose prime factors only include 2, 3, 5.
 *
 * Example:
 *
 *
 * Input: n = 10
 * Output: 12
 * Explanation: 1, 2, 3, 4, 5, 6, 8, 9, 10, 12 is the sequence of the first 10 ugly numbers.
 *
 * Note:
 *
 *
 * 	1 is typically treated as an ugly number.
 * 	n does not exceed 1690.
 *
 */

func TestUglyNumberII(t *testing.T) {
	var cases = []struct {
		input  int
		output int
	}{
		{
			input:  10,
			output: 12,
		},
	}
	for _, c := range cases {
		x := nthUglyNumber(c.input)
		fmt.Print(x)
		if x != c.output {
			t.Fail()
		}
	}
}

// submission codes start here

func nthUglyNumber(n int) int {
	count := 1
	res := make([]int, n)
	res[0] = 1
	indexA, indexB, indexC := 0, 0, 0
	for count < n {
		res[count] = uglyMin(res[indexA]*2, res[indexB]*3, res[indexC]*5)
		if res[count] == res[indexA]*2 {
			indexA += 1
		}
		if res[count] == res[indexB]*3 {
			indexB += 1
		}
		if res[count] == res[indexC]*5 {
			indexC += 1
		}
		count++
	}
	return res[n-1]
}

func uglyMin(a, b, c int) int {
	if a <= b && a <= c {
		return a
	}
	if a >= b && b <= c {
		return b
	}
	return c
}

// submission codes end
