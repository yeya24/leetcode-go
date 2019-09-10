package tests

import (
	"reflect"
	"testing"
)

/**
 * [979] DI String Match
 *
 * Given a string S that only contains "I" (increase) or "D" (decrease), let N = S.length.
 *
 * Return any permutation A of [0, 1, ..., N] such that for all i = 0, ..., N-1:
 *
 *
 * 	If S[i] == "I", then A[i] < A[i+1]
 * 	If S[i] == "D", then A[i] > A[i+1]
 *
 *
 *
 *
 * Example 1:
 *
 *
 * Input: "IDID"
 * Output: [0,4,1,3,2]
 *
 *
 *
 * Example 2:
 *
 *
 * Input: "III"
 * Output: [0,1,2,3]
 *
 *
 *
 * Example 3:
 *
 *
 * Input: "DDI"
 * Output: [3,2,0,1]
 *
 *
 *
 *
 *
 * Note:
 *
 *
 * 	1 <= S.length <= 10000
 * 	S only contains characters "I" or "D".
 *
 */

func TestDIStringMatch(t *testing.T) {
	var cases = []struct {
		input  string
		output []int
	}{
		{
			input:  "IDID",
			output: []int{0, 4, 1, 3, 2},
		},
		{
			input:  "III",
			output: []int{0, 1, 2, 3},
		},
		{
			input:  "DDI",
			output: []int{3, 2, 0, 1},
		},
	}
	for _, c := range cases {
		x := diStringMatch(c.input)
		if !reflect.DeepEqual(x, c.output) {
			t.Fail()
		}
	}
}

// submission codes start here

func diStringMatch(S string) []int {
	n := len(S)
	res := make([]int, n+1)
	i, j := 0, n
	for k, v := range S {
		if v == 'I' {
			res[k] = i
			i++
		} else {
			res[k] = j
			j--
		}
	}
	res[n] = i
	return res
}

// submission codes end
