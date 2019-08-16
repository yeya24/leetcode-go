package tests

import (
	"testing"
)

/**
 * [742] To Lower Case
 *
 * Implement function ToLowerCase() that has a string parameter str, and returns the same string in lowercase.
 *
 *
 *
 *
 * Example 1:
 *
 *
 * Input: "Hello"
 * Output: "hello"
 *
 *
 *
 * Example 2:
 *
 *
 * Input: "here"
 * Output: "here"
 *
 *
 *
 * Example 3:
 *
 *
 * Input: "LOVELY"
 * Output: "lovely"
 *
 *
 *
 *
 *
 */

func TestToLowerCase(t *testing.T) {
	var cases = []struct {
		input  string
		output string
	}{
		{
			input:  "Hello",
			output: "hello",
		},
		{
			input:  "here",
			output: "here",
		},
		{
			input:  "LOVELY",
			output: "lovely",
		},
	}
	for _, c := range cases {
		x := toLowerCase(c.input)
		if x != c.output {
			t.Fail()
		}
	}
}

// submission codes start here

func toLowerCase(str string) string {
    b := []byte(str)
	for i := 0; i < len(b); i++ {
		if b[i] >= 'A' && b[i] <= 'Z' {
			b[i] = str[i] + 'a' - 'A'
		}
	}
	return string(b)
}

// submission codes end
