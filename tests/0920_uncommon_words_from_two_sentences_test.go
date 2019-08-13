package tests

import (
	"reflect"
    "strings"
    "testing"
)

/**
 * [920] Uncommon Words from Two Sentences
 *
 * We are given two sentences A and B.  (A sentence is a string of space separated words.  Each word consists only of lowercase letters.)
 *
 * A word is uncommon if it appears exactly once in one of the sentences, and does not appear in the other sentence.
 *
 * Return a list of all uncommon words.
 *
 * You may return the list in any order.
 *
 *
 *
 *
 *
 *
 *
 * Example 1:
 *
 *
 * Input: A = "this apple is sweet", B = "this apple is sour"
 * Output: ["sweet","sour"]
 *
 *
 *
 * Example 2:
 *
 *
 * Input: A = "apple apple", B = "banana"
 * Output: ["banana"]
 *
 *
 *
 *
 * Note:
 *
 *
 * 	0 <= A.length <= 200
 * 	0 <= B.length <= 200
 * 	A and B both contain only spaces and lowercase letters.
 *
 *
 *
 *
 */

func TestUncommonWordsfromTwoSentences(t *testing.T) {
	var cases = []struct {
		inputA string
		inputB string
		output []string
	}{
		{
			inputA: "this apple is sweet",
			inputB: "this apple is sour",
			output: []string{"sweet", "sour"},
		},
		{
			inputA: "apple apple",
			inputB: "banana",
			output: []string{"banana"},
		},
	}
	for _, c := range cases {
		x := uncommonFromSentences(c.inputA, c.inputB)
		if !reflect.DeepEqual(x, c.output) {
			t.Fail()
		}
	}
}

// submission codes start here

func uncommonFromSentences(A string, B string) []string {
    ss := strings.Split(A + " " + B, " ")
	words := make([]string, 0)
	m := make(map[string]int)
	for _, a := range ss {
		m[a] += 1
	}
	for i, a := range m	{
		if a == 1 {
			words = append(words, i)
		}
	}
	return words
}

// submission codes end
