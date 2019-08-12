package tests

import (
	"testing"
)

/**
 * [409] Longest Palindrome
 *
 * Given a string which consists of lowercase or uppercase letters, find the length of the longest palindromes that can be built with those letters.
 *
 * This is case sensitive, for example "Aa" is not considered a palindrome here.
 *
 * Note:
 * Assume the length of given string will not exceed 1,010.
 *
 *
 * Example:
 *
 * Input:
 * "abccccdd"
 *
 * Output:
 * 7
 *
 * Explanation:
 * One longest palindrome that can be built is "dccaccd", whose length is 7.
 *
 *
 */

func TestLongestPalindrome(t *testing.T) {
	var cases = []struct {
		s      string
		output int
	}{
		{
			s:      "abccccdd",
			output: 7,
		},
		{
			s:      "aabbbccc",
			output: 7,
		},
	}
	for _, c := range cases {
		x := longestPalindrome(c.s)
		if x != c.output {
			t.Fail()
		}
	}
}

// submission codes start here

func longestPalindrome(s string) int {
	m := make(map[int32]int)
	c := 0
	var maxOdd int
	for _, v := range s {
		m[v] += 1
	}
	for _, v := range m {
		if v%2 == 0 {
			c += v
		} else if v > maxOdd {
			if maxOdd != 0 {
				c -= 1
			}
			c += v
			maxOdd = v
		} else {
			c += v - 1
		}
	}
	return c
}

// submission codes end
