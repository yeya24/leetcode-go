package tests

import (
	"testing"
)

/**
 * [1128] Remove All Adjacent Duplicates In String
 *
 * Given a string S of lowercase letters, a duplicate removal consists of choosing two adjacent and equal letters, and removing them.
 *
 * We repeatedly make duplicate removals on S until we no longer can.
 *
 * Return the final string after all such duplicate removals have been made.  It is guaranteed the answer is unique.
 *
 *
 *
 * Example 1:
 *
 *
 * Input: "abbaca"
 * Output: "ca"
 * Explanation:
 * For example, in "abbaca" we could remove "bb" since the letters are adjacent and equal, and this is the only possible move.  The result of this move is that the string is "aaca", of which only "aa" is possible, so the final string is "ca".
 *
 *
 *
 *
 * Note:
 *
 *
 * 	1 <= S.length <= 20000
 * 	S consists only of English lowercase letters.
 *
 */

func TestRemoveAllAdjacentDuplicatesInString(t *testing.T) {
	var cases = []struct {
		input  string
		output string
	}{
		{
			input:  "abbaca",
			output: "ca",
		},
	}
	for _, c := range cases {
		x := removeDuplicates(c.input)
		if x != c.output {
			t.Fail()
		}
	}
}

// submission codes start here

func removeDuplicates(S string) string {
	s := []byte{}
	x := []byte(S)
	for i := 0; i < len(S); i++ {
		l := len(s)
		if l > 0 {
			if x[i] == s[l-1] {
				s = s[:l-1]
			} else {
				s = append(s, x[i])
			}
		} else {
			s = append(s, x[i])
		}
	}
	return string(s)
}

// submission codes end
