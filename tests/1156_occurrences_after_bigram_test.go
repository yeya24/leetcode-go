package tests

import (
	"reflect"
	"strings"
	"testing"
)

/**
 * [1156] Occurrences After Bigram
 *
 * Given words first and second, consider occurrences in some text of the form "first second third", where second comes immediately after first, and third comes immediately after second.
 *
 * For each such occurrence, add "third" to the answer, and return the answer.
 *
 *
 *
 * Example 1:
 *
 *
 * Input: text = "alice is a good girl she is a good student", first = "a", second = "good"
 * Output: ["girl","student"]
 *
 *
 *
 * Example 2:
 *
 *
 * Input: text = "we will we will rock you", first = "we", second = "will"
 * Output: ["we","rock"]
 *
 *
 *
 *
 * Note:
 *
 *
 * 	1 <= text.length <= 1000
 * 	text consists of space separated words, where each word consists of lowercase English letters.
 * 	1 <= first.length, second.length <= 10
 * 	first and second consist of lowercase English letters.
 *
 *
 */

func TestOccurrencesAfterBigram(t *testing.T) {
	var cases = []struct {
		input  string
		f      string
		s      string
		output []string
	}{
		{
			input:  "alice is a good girl she is a good student",
			f:      "a",
			s:      "good",
			output: []string{"girl", "student"},
		},
		{
			input:  "we will we will rock you",
			f:      "we",
			s:      "will",
			output: []string{"we", "rock"},
		},
	}
	for _, c := range cases {
		x := findOcurrences(c.input, c.f, c.s)
		if !reflect.DeepEqual(x, c.output) {
			t.Fail()
		}
	}
}

// submission codes start here

func findOcurrences(text string, first string, second string) []string {
	res := []string{}
	ss := strings.Split(text, " ")
	for i := 0; i < len(ss)-2; i++ {
		if strings.Compare(ss[i], first) == 0 {
			if strings.Compare(ss[i+1], second) == 0 {
				res = append(res, ss[i+2])
			}
		}
	}
	return res
}

// submission codes end
