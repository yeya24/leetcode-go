package tests

import (
	"testing"
)

/**
 * [1320] Remove All Adjacent Duplicates in String II
 *
 * Given a string s, a k duplicate removal consists of choosing k adjacent and equal letters from s and removing them causing the left and the right side of the deleted substring to concatenate together.
 * We repeatedly make k duplicate removals on s until we no longer can.
 * Return the final string after all such duplicate removals have been made.
 * It is guaranteed that the answer is unique.
 *
 * Example 1:
 *
 * Input: s = "abcd", k = 2
 * Output: "abcd"
 * Explanation: There's nothing to delete.
 * Example 2:
 *
 * Input: s = "deeedbbcccbdaa", k = 3
 * Output: "aa"
 * Explanation:
 * First delete "eee" and "ccc", get "ddbbbdaa"
 * Then delete "bbb", get "dddaa"
 * Finally delete "ddd", get "aa"
 * Example 3:
 *
 * Input: s = "pbbcggttciiippooaais", k = 2
 * Output: "ps"
 *
 *
 * Constraints:
 *
 * 	1 <= s.length <= 10^5
 * 	2 <= k <= 10^4
 * 	s only contains lower case English letters.
 *
 *
 */

func TestRemoveAllAdjacentDuplicatesinStringII(t *testing.T) {
	var cases = []struct {
		input  string
		k      int
		output string
	}{
		{
			input:  "abcd",
			k: 2,
			output: "abcd",
		},
		{
			input:  "deeedbbcccbdaa",
			k: 3,
			output: "aa",
		},
		{
			input: "pbbcggttciiippooaais",
			k: 2,
			output: "ps",
		},
	}
	for _, c := range cases {
		x := removeDuplicatesII(c.input, c.k)
		if x != c.output {
			t.Fail()
		}
	}
}

// submission codes start here

func removeDuplicatesII(s string, k int) string {
	var stack []rune
	var count []int
	for _, v := range s {
		if len(stack) == 0 {
			stack = append(stack, v)
			count = append(count, 1)
		} else {
			l := len(stack)
			if v == stack[l-1] {
				stack = append(stack, v)
				count = append(count, count[l-1]+1)
			} else {
				stack = append(stack, v)
				count = append(count, 1)
			}
		}
		for len(count) > 0 && count[len(count)-1] == k {
			stack = stack[:len(count)-k]
			count = count[:len(count)-k]
		}
	}
	return string(stack)
}

// submission codes end
