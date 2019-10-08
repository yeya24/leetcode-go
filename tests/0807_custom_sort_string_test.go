package tests

import (
	"sort"
	"testing"
)

/**
 * [807] Custom Sort String
 *
 * S and T are strings composed of lowercase letters. In S, no letter occurs more than once.
 *
 * S was sorted in some custom order previously. We want to permute the characters of T so that they match the order that S was sorted. More specifically, if x occurs before y in S, then x should occur before y in the returned string.
 *
 * Return any permutation of T (as a string) that satisfies this property.
 *
 *
 * Example :
 * Input:
 * S = "cba"
 * T = "abcd"
 * Output: "cbad"
 * Explanation:
 * "a", "b", "c" appear in S, so the order of "a", "b", "c" should be "c", "b", and "a".
 * Since "d" does not appear in S, it can be at any position in T. "dcba", "cdba", "cbda" are also valid outputs.
 *
 *
 *
 *
 * Note:
 *
 *
 * 	S has length at most 26, and no character is repeated in S.
 * 	T has length at most 200.
 * 	S and T consist of lowercase letters only.
 *
 *
 */

func TestCustomSortString(t *testing.T) {
	var cases = []struct {
		S      string
		T      string
		output string
	}{
		{
			S:      "cba",
			T:      "abcd",
			output: "cbad",
		},
		{
			S: "test",
			T: "abtggsccedt",
			output: "",
		},
	}
	for _, c := range cases {
		x := customSortString(c.S, c.T)
		if x != c.output {
			t.Fail()
		}
	}
}

// submission codes start here

func customSortString(S string, T string) string {
    m := [26]int{}
    for i, v := range S {
    	if m[v-'a'] == 0 {
    		m[v-'a'] = i+1
		}
    }
    x := []byte(T)
    sort.Slice(x, func(i, j int) bool {
    	return m[x[i]-'a'] < m[x[j]-'a']
	})
    return string(x)
}

// submission codes end
