package tests

import (
	"reflect"
	"testing"
)

/**
 * [857] Positions of Large Groups
 *
 * In a string S of lowercase letters, these letters form consecutive groups of the same character.
 *
 * For example, a string like S = "abbxxxxzyy" has the groups "a", "bb", "xxxx", "z" and "yy".
 *
 * Call a group large if it has 3 or more characters.  We would like the starting and ending positions of every large group.
 *
 * The final answer should be in lexicographic order.
 *
 *
 *
 * Example 1:
 *
 *
 * Input: "abbxxxxzzy"
 * Output: [[3,6]]
 * Explanation: "xxxx" is the single large group with starting  3 and ending positions 6.
 *
 *
 * Example 2:
 *
 *
 * Input: "abc"
 * Output: []
 * Explanation: We have "a","b" and "c" but no large group.
 *
 *
 * Example 3:
 *
 *
 * Input: "abcdddeeeeaabbbcd"
 * Output: [[3,5],[6,9],[12,14]]
 *
 *
 *
 * Note:  1 <= S.length <= 1000
 *
 */

func TestPositionsofLargeGroups(t *testing.T) {
	var cases = []struct {
		input  string
		output [][]int
	}{
		//{
		//	input:  "abbxxxxzzy",
		//	output: [][]int{{3, 6}},
		//},
		//{
		//	input:  "abc",
		//	output: [][]int{},
		//},
		//{
		//	input:  "abcdddeeeeaabbbcd",
		//	output: [][]int{{3, 5}, {6, 9}, {12, 14}},
		//},
		{
			input:  "aaa",
			output: [][]int{{0, 2}},
		},
	}
	for _, c := range cases {
		x := largeGroupPositions(c.input)
		if !reflect.DeepEqual(x, c.output) {
			t.Fail()
		}
	}
}

// submission codes start here

func largeGroupPositions(S string) [][]int {
	res := [][]int{}
	i := 0
	l := len(S)
	for ; i < l; {
		c := 1
		for i+c < l && S[i+c] == S[i] {
			c++
		}
		if c >= 3 {
			res = append(res, []int{i, i + c - 1})
			i += c
		} else {
			i++
		}
	}
	return res
}

// submission codes end
