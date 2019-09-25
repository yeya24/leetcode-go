package tests

import (
	"fmt"
	"reflect"
	"testing"
)

/**
 * [1273] Compare Strings by Frequency of the Smallest Character
 *
 * Let's define a function f(s) over a non-empty string s, which calculates the frequency of the smallest character in s. For example, if s = "dcce" then f(s) = 2 because the smallest character is "c" and its frequency is 2.
 * Now, given string arrays queries and words, return an integer array answer, where each answer[i] is the number of words such that f(queries[i]) < f(W), where W is a word in words.
 *
 * Example 1:
 *
 * Input: queries = ["cbd"], words = ["zaaaz"]
 * Output: [1]
 * Explanation: On the first query we have f("cbd") = 1, f("zaaaz") = 3 so f("cbd") < f("zaaaz").
 *
 * Example 2:
 *
 * Input: queries = ["bbb","cc"], words = ["a","aa","aaa","aaaa"]
 * Output: [1,2]
 * Explanation: On the first query only f("bbb") < f("aaaa"). On the second query both f("aaa") and f("aaaa") are both > f("cc").
 *
 *
 * Constraints:
 *
 * 	1 <= queries.length <= 2000
 * 	1 <= words.length <= 2000
 * 	1 <= queries[i].length, words[i].length <= 10
 * 	queries[i][j], words[i][j] are English lowercase letters.
 *
 *
 */

func TestCompareStringsbyFrequencyoftheSmallestCharacter(t *testing.T) {
	var cases = []struct {
		input  []string
		words  []string
		output []int
	}{
		{
			input:  []string{"cbd"},
			words:  []string{"zaaaz"},
			output: []int{1},
		},
		{
			input:  []string{"bbb", "cc"},
			words:  []string{"a", "aa", "aaa", "aaaa"},
			output: []int{1, 2},
		},
		{
			input:  []string{"bba", "abaaaaaa", "aaaaaa", "bbabbabaab", "aba", "aa", "baab", "bbbbbb", "aab", "bbabbaabb"},
			words:  []string{"aaabbb", "aab", "babbab", "babbbb", "b", "bbbbbbbbab", "a", "bbbbbbbbbb", "baaabbaab", "aa"},
			output: []int{6, 1, 1, 2, 3, 3, 3, 1, 3, 2},
		},
	}
	for _, c := range cases {
		x := numSmallerByFrequency(c.input, c.words)
		fmt.Println(x)
		if !reflect.DeepEqual(x, c.output) {
			t.Fail()
		}
	}
}

// submission codes start here

func numSmallerByFrequency(queries []string, words []string) []int {
	ww := [11]int{}
	for _, v := range words {
		n := counter(v)
		for j := n; j > 0; j-- {
			ww[j-1] += 1
		}
	}
	res := make([]int, len(queries))
	for i := 0; i < len(queries); i++ {
		n := counter(queries[i])
		res[i] = ww[n]
	}
	return res
}

func counter(word string) int {
	m := [26]int{}
	min := 26
	for _, v := range word {
		if int(v-'a') < min {
			min = int(v - 'a')
			m[min] += 1
		} else if int(v-'a') == min {
			m[min] += 1
		}
	}
	return m[min]
}

// submission codes end
