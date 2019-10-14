package tests

import (
	"fmt"
	"sort"
	"testing"
)

/**
 * [49] Group Anagrams
 *
 * Given an array of strings, group anagrams together.
 *
 * Example:
 *
 *
 * Input: ["eat", "tea", "tan", "ate", "nat", "bat"],
 * Output:
 * [
 *   ["ate","eat","tea"],
 *   ["nat","tan"],
 *   ["bat"]
 * ]
 *
 * Note:
 *
 *
 * 	All inputs will be in lowercase.
 * 	The order of your output does not matter.
 *
 *
 */

func TestGroupAnagrams(t *testing.T) {
	var cases = []struct {
		input  []string
		output [][]string
	}{
		{
			input: []string{"eat", "tea", "tan", "ate", "nat", "bat"},
			output: [][]string{{"ate", "eat", "tea"},
				{"nat", "tan"},
				{"bat"}},
		},
	}
	for _, c := range cases {
		x := groupAnagrams2(c.input)
		fmt.Println(x)
	}
}

// submission codes start here

// First sort every string, store the sorted string as key, in a map. Time(NKlogK) Space(NK)
func groupAnagrams(strs []string) [][]string {
	m := make(map[string][]string)
	for _, v := range strs {
		c := sortString(v)
		m[c] = append(m[c], v)
	}
	var res [][]string
	for _, v := range m {
		res = append(res, v)
	}
	return res
}

func sortString(str string) string {
	v := []byte(str)
	sort.Slice(v, func(i, j int) bool {
		return v[i] < v[j]
	})
	return string(v)
}

// Don't sort string, use a strint to represent.
func groupAnagrams2(strs []string) [][]string {
	m := make(map[string][]string)
	for _, v := range strs {
		c := transString(v)
		m[c] = append(m[c], v)
	}
	var res [][]string
	for _, v := range m {
		res = append(res, v)
	}
	return res
}

func transString(str string) string {
	m := []byte("00000000000000000000000000")
	for _, v := range str {
		m[v-'a'] +=1
	}
	return string(m)
}

// submission codes end
