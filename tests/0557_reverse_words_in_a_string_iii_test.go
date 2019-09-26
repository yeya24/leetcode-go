package tests

import (
    "fmt"
    "testing"
)

/**
 * [557] Reverse Words in a String III
 *
 * Given a string, you need to reverse the order of characters in each word within a sentence while still preserving whitespace and initial word order.
 *
 * Example 1:
 *
 * Input: "Let's take LeetCode contest"
 * Output: "s'teL ekat edoCteeL tsetnoc"
 *
 *
 *
 * Note:
 * In the string, each word is separated by single space and there will not be any extra space in the string.
 *
 */

func TestReverseWordsinaStringIII(t *testing.T) {
	var cases = []struct {
		input  string
		output string
	}{
		{
			input:  "Let's take LeetCode contest",
			output: "s'teL ekat edoCteeL tsetnoc",
		},
	}
	for _, c := range cases {
		x := reverseWords(c.input)
		fmt.Println(x)
		if x != c.output {
			t.Fail()
		}
	}
}

// submission codes start here

func reverseWords(s string) string {
	start := 0
	x := []byte(s)
	for i, v := range x {
		if v == 32 {
			end := i - 1
			for start < end {
				x[start], x[end] = x[end], x[start]
				start++
				end--
			}
			start = i + 1
		}
	}
	end := len(s) - 1
	for start < end {
		x[start], x[end] = x[end], x[start]
		start++
		end--
	}
	return string(x)
}

// submission codes end
