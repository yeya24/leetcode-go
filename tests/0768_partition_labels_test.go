package tests

import (
	"reflect"
	"testing"
)

/**
 * [768] Partition Labels
 *
 *
 * A string S of lowercase letters is given.  We want to partition this string into as many parts as possible so that each letter appears in at most one part, and return a list of integers representing the size of these parts.
 *
 *
 * Example 1:
 *
 * Input: S = "ababcbacadefegdehijhklij"
 * Output: [9,7,8]
 * Explanation:
 * The partition is "ababcbaca", "defegde", "hijhklij".
 * This is a partition so that each letter appears in at most one part.
 * A partition like "ababcbacadefegde", "hijhklij" is incorrect, because it splits S into less parts.
 *
 *
 *
 * Note:<br>
 * S will have length in range [1, 500].
 * S will consist of lowercase letters ('a' to 'z') only.
 *
 */

func TestPartitionLabels(t *testing.T) {
	var cases = []struct {
		input  string
		output []int
	}{
		{
			input:  "ababcbacadefegdehijhklij",
			output: []int{9, 7, 8},
		},
	}
	for _, c := range cases {
		x := partitionLabels(c.input)
		if !reflect.DeepEqual(x, c.output) {
			t.Fail()
		}
	}
}

// submission codes start here

func partitionLabels(S string) []int {
	char := [26]int{}
	for i, v := range S {
		char[v-'a'] = i
	}
	start, last := 0, 0
	res := make([]int, 0)
	for i := 0; i < len(S); i++ {
		if char[S[i]-'a'] > last {
			last = char[S[i]-'a']
		}
		if last == i {
			res = append(res, last-start+1)
			start = last + 1
		}
	}
	return res
}

// submission codes end
