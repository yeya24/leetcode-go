package tests

import (
	"testing"
)

/**
 * [344] Reverse String
 *
 * Write a function that reverses a string. The input string is given as an array of characters char[].
 *
 * Do not allocate extra space for another array, you must do this by modifying the input array in-place with O(1) extra memory.
 *
 * You may assume all the characters consist of printable ascii characters.
 *
 *
 *
 *
 * Example 1:
 *
 *
 * Input: ["h","e","l","l","o"]
 * Output: ["o","l","l","e","h"]
 *
 *
 *
 * Example 2:
 *
 *
 * Input: ["H","a","n","n","a","h"]
 * Output: ["h","a","n","n","a","H"]
 *
 *
 *
 *
 */

func TestReverseString(t *testing.T) {
	cases := []struct {
		input  string
		expect string
	}{
		{
			input:  "hello",
			expect: "olleh",
		},
		{
			input:  "test",
			expect: "tset",
		},
	}
	for _, tc := range cases {
		reverseString([]byte(tc.input))
	}
}

// submission codes start here

func reverseString(s []byte) {
	end := len(s) - 1
	start := 0
	for start < end {
		s[end], s[start] = s[start], s[end]
		start++
		end--
	}
}

// submission codes end
