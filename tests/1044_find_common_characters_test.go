package tests

import (
	"math"
	"reflect"
	"testing"
)

/**
 * [1044] Find Common Characters
 *
 * Given an array A of strings made only from lowercase letters, return a list of all characters that show up in all strings within the list (including duplicates).  For example, if a character occurs 3 times in all strings but not 4 times, you need to include that character three times in the final answer.
 *
 * You may return the answer in any order.
 *
 *
 *
 *
 * Example 1:
 *
 *
 * Input: ["bella","label","roller"]
 * Output: ["e","l","l"]
 *
 *
 *
 * Example 2:
 *
 *
 * Input: ["cool","lock","cook"]
 * Output: ["c","o"]
 *
 *
 *
 *
 * Note:
 *
 *
 * 	1 <= A.length <= 100
 * 	1 <= A[i].length <= 100
 * 	A[i][j] is a lowercase letter
 *
 *
 *
 */

func TestFindCommonCharacters(t *testing.T) {
	var cases = []struct {
		input  []string
		output []string
	}{
		{
			input:  []string{"cool", "lock", "cook"},
			output: []string{"c", "o"},
		},
		{
			input:  []string{"bella", "label", "roller"},
			output: []string{"e", "l", "l"},
		},
		{
			// one word might have multiple same char
			input:  []string{"acabcddd", "bcbdbcbd", "baddbadb", "cbdddcac", "aacbcccd", "ccccddda", "cababaab", "addcaccd"},
			output: []string{},
		},
	}
	for _, c := range cases {
		x := commonChars(c.input)
		if !reflect.DeepEqual(x, c.output) {
			t.Fail()
		}
	}
}

// submission codes start here

func commonChars(A []string) []string {
	m := [26]int{}
	for i := 0; i < 26; i++ {
		m[i] = math.MaxInt64
	}
	res := []string{}
	for _, s := range A {
		m1 := [26]int{}
		for _, v := range s {
			m1[v-97] += 1
		}
		for i := 0; i < 26; i++ {
			m[i] = int(math.Min(float64(m1[i]), float64(m[i])))
		}
	}
	for i := 0; i < 26; i++ {
		for j := 0; j < m[i]; j++ {
			res = append(res, string(i+'a'))
		}
	}
	return res
}

// submission codes end
