package tests

import (
	"testing"
)

/**
 * [961] Long Pressed Name
 *
 * Your friend is typing his name into a keyboard.  Sometimes, when typing a character c, the key might get long pressed, and the character will be typed 1 or more times.
 *
 * You examine the typed characters of the keyboard.  Return True if it is possible that it was your friends name, with some characters (possibly none) being long pressed.
 *
 *
 *
 * Example 1:
 *
 *
 * Input: name = "alex", typed = "aaleex"
 * Output: true
 * Explanation: 'a' and 'e' in 'alex' were long pressed.
 *
 *
 *
 * Example 2:
 *
 *
 * Input: name = "saeed", typed = "ssaaedd"
 * Output: false
 * Explanation: 'e' must have been pressed twice, but it wasn't in the typed output.
 *
 *
 *
 * Example 3:
 *
 *
 * Input: name = "leelee", typed = "lleeelee"
 * Output: true
 *
 *
 *
 * Example 4:
 *
 *
 * Input: name = "laiden", typed = "laiden"
 * Output: true
 * Explanation: It's not necessary to long press any character.
 *
 *
 *
 *
 *
 *
 *
 * Note:
 *
 *
 * 	name.length <= 1000
 * 	typed.length <= 1000
 * 	The characters of name and typed are lowercase letters.
 *
 *
 *
 *
 *
 *
 *
 *
 *
 *
 *
 */

func TestLongPressedName(t *testing.T) {
	var cases = []struct {
		name   string
		typed  string
		output bool
	}{
		{
			name:   "alex",
			typed:  "aaleex",
			output: true,
		},
		{
			name:   "leelee",
			typed:  "lleeelee",
			output: true,
		},
		{
			name:   "saeed",
			typed:  "ssaaedd",
			output: false,
		},
		{
			name:   "laiden",
			typed:  "laiden",
			output: true,
		},
		{
			name:   "pyplrz",
			typed:  "ppyypllr",
			output: false,
		},
	}
	for _, c := range cases {
		x := isLongPressedName(c.name, c.typed)
		if x != c.output {
			t.Fail()
		}
	}
}

// submission codes start here

func isLongPressedName(name string, typed string) bool {
	i := 0
	j := 0
	for i < len(name) && j < len(typed) {
		c := name[i]
		x := 0
		z := 0
		for name[i] == c && i < len(name) {
			x++
			i++
			if i == len(name) {
				break
			}
		}
		for typed[j] == c && j < len(typed) {
			j++
			z++
			if j == len(typed) {
				if i < len(name) {
					return false
				}
				break
			}
		}
		if z < x {
			return false
		}
	}
	return true
}

// submission codes end
