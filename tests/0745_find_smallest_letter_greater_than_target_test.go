package tests

import (
	"testing"
)

/**
 * [745] Find Smallest Letter Greater Than Target
 *
 *
 * Given a list of sorted characters letters containing only lowercase letters, and given a target letter target, find the smallest element in the list that is larger than the given target.
 *
 * Letters also wrap around.  For example, if the target is target = 'z' and letters = ['a', 'b'], the answer is 'a'.
 *
 *
 * Examples:
 *
 * Input:
 * letters = ["c", "f", "j"]
 * target = "a"
 * Output: "c"
 *
 * Input:
 * letters = ["c", "f", "j"]
 * target = "c"
 * Output: "f"
 *
 * Input:
 * letters = ["c", "f", "j"]
 * target = "d"
 * Output: "f"
 *
 * Input:
 * letters = ["c", "f", "j"]
 * target = "g"
 * Output: "j"
 *
 * Input:
 * letters = ["c", "f", "j"]
 * target = "j"
 * Output: "c"
 *
 * Input:
 * letters = ["c", "f", "j"]
 * target = "k"
 * Output: "c"
 *
 *
 *
 * Note:
 *
 * letters has a length in range [2, 10000].
 * letters consists of lowercase letters, and contains at least 2 unique letters.
 * target is a lowercase letter.
 *
 *
 */

func TestFindSmallestLetterGreaterThanTarget(t *testing.T) {
	var cases = []struct {
		input  []byte
		target byte
		output byte
	}{
		{
			input:  []byte{'c', 'f', 'j'},
			target: 'a',
			output: 'c',
		},
		{
			input:  []byte{'c', 'f', 'j'},
			target: 'c',
			output: 'f',
		},
		{
			input: []byte{'a','b'},
			target: 'z',
			output: 'a',
		},
	}
	for _, c := range cases {
		x := nextGreatestLetter(c.input, c.target)
		if x != c.output {
			t.Fail()
		}
	}
}

// submission codes start here

func nextGreatestLetter(letters []byte, target byte) byte {
	i, j := 0, len(letters)
	for i < j {
		m := (i + j) / 2
		if letters[m] <= target {
			i = m + 1
		} else {
			j = m
		}
	}
	return letters[j % len(letters)]
}

// submission codes end
