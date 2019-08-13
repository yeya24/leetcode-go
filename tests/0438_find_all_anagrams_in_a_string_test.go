package tests

import (
	"reflect"
	"testing"
)

/**
 * [438] Find All Anagrams in a String
 *
 * Given a string s and a non-empty string p, find all the start indices of p's anagrams in s.
 *
 * Strings consists of lowercase English letters only and the length of both strings s and p will not be larger than 20,100.
 *
 * The order of output does not matter.
 *
 * Example 1:
 *
 * Input:
 * s: "cbaebabacd" p: "abc"
 *
 * Output:
 * [0, 6]
 *
 * Explanation:
 * The substring with start index = 0 is "cba", which is an anagram of "abc".
 * The substring with start index = 6 is "bac", which is an anagram of "abc".
 *
 *
 *
 * Example 2:
 *
 * Input:
 * s: "abab" p: "ab"
 *
 * Output:
 * [0, 1, 2]
 *
 * Explanation:
 * The substring with start index = 0 is "ab", which is an anagram of "ab".
 * The substring with start index = 1 is "ba", which is an anagram of "ab".
 * The substring with start index = 2 is "ab", which is an anagram of "ab".
 *
 *
 */

func TestFindAllAnagramsinaString(t *testing.T) {
	var cases = []struct {
		s      string
		p      string
		output []int
	}{
		{
			s:      "cbaebabacd",
			p:      "abc",
			output: []int{0, 6},
		},
		{
			s:      "abab",
			p:      "ab",
			output: []int{0, 1, 2},
		},
	}
	for _, c := range cases {
		x := findAnagrams(c.s, c.p)
		if !reflect.DeepEqual(x, c.output) {
			t.Fail()
		}
	}
}

// submission codes start here

func findAnagrams(s string, p string) []int {
	ls := len(s)
	lp := len(p)
	ms := [26]int{}
	mp := [26]int{}
	res := []int{}
	for _, v := range p {
		mp[v-'a'] += 1
	}
	for i := 0; i < ls; i++ {
		if i >= lp {
			ms[s[i-lp]-'a'] -= 1
		}
		ms[s[i]-'a']+=1
		if ms == mp {
			res = append(res, i-lp+1)
		}
	}
	return res
}

// submission codes end
