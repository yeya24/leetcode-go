package tests

import (
	"testing"
)

/**
 * [20] Valid Parentheses
 *
 * Given a string containing just the characters '(', ')', '{', '}', '[' and ']', determine if the input string is valid.
 *
 * An input string is valid if:
 *
 *
 * 	Open brackets must be closed by the same type of brackets.
 * 	Open brackets must be closed in the correct order.
 *
 *
 * Note that an empty string is also considered valid.
 *
 * Example 1:
 *
 *
 * Input: "()"
 * Output: true
 *
 *
 * Example 2:
 *
 *
 * Input: "()[]{}"
 * Output: true
 *
 *
 * Example 3:
 *
 *
 * Input: "(]"
 * Output: false
 *
 *
 * Example 4:
 *
 *
 * Input: "([)]"
 * Output: false
 *
 *
 * Example 5:
 *
 *
 * Input: "{[]}"
 * Output: true
 *
 *
 */

func TestValidParentheses(t *testing.T) {
	var cases = []struct {
		input  string
		output bool
	}{
		{
			input:  "()",
			output: true,
		},
		{
			input:  "()[]{}",
			output: true,
		},
		{
			input:  "([)]",
			output: false,
		},
		{
			input:  "(]",
			output: false,
		},
		{
			input:  "{[]}",
			output: true,
		},
		{
			input:  "",
			output: true,
		},
		{
			input:  "[",
			output: false,
		},
	}
	for _, c := range cases {
		x := isValid(c.input)
		if x != c.output {
			t.Fail()
		}
	}
}

// submission codes start here

func isValid(s string) bool {
	x := []byte(s)
	stack := []byte{}
	for _, v := range x {
		if len(stack) == 0 {
			stack = append(stack, v)
		} else {
			p := stack[len(stack)-1]
			switch p {
			case '(':
				if v == ')' {
					stack = stack[:len(stack)-1]
				} else if v == '}' || v == ']' {
					return false
				} else {
					stack = append(stack, v)
				}
			case '[':
				if v == ']' {
					stack = stack[:len(stack)-1]
				} else if v == '}' || v == ')' {
					return false
				} else {
					stack = append(stack, v)
				}
			case '{':
				if v == '}' {
					stack = stack[:len(stack)-1]
				} else if v == ']' || v == ')' {
					return false
				} else {
					stack = append(stack, v)
				}
			}
		}
	}
	if len(stack) > 0 {
		return false
	}
	return true
}

// submission codes end
