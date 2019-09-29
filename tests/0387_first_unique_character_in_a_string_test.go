package tests

import (
	"testing"
)

/**
 * [387] First Unique Character in a String
 *
 *
 * Given a string, find the first non-repeating character in it and return it's index. If it doesn't exist, return -1.
 *
 * Examples:
 *
 * s = "leetcode"
 * return 0.
 *
 * s = "loveleetcode",
 * return 2.
 *
 *
 *
 *
 * Note: You may assume the string contain only lowercase letters.
 *
 */

func TestFirstUniqueCharacterinaString(t *testing.T) {
	var cases = []struct {
		input  string
		output int
	}{
		{
			input:  "leetcode",
			output: 0,
		},
		{
			input:  "loveleetcode",
			output: 2,
		},
		{
			input:  "cc",
			output: -1,
		},
		{
			input:  "z",
			output: 0,
		},
	}
	for _, c := range cases {
		x := firstUniqChar(c.input)
		if x != c.output {
			t.Fail()
		}
	}
}

// submission codes start here

func firstUniqChar(s string) int {
	c := [26]int{}
	for i, v := range s {
		if c[v-'a'] == 0 {
			c[v-'a'] = i + 1
		} else if c[v-'a'] > 0 {
			c[v-'a'] = -1
		}
	}
	mm := len(s) + 1
	for _, v := range c {
		if v > 0 {
			if v < mm {
				mm = v
			}
		}
	}
	if mm == len(s) + 1 {
		return -1
	}
	return mm - 1
}

// submission codes end
