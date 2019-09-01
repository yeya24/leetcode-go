package tests

import (
	"testing"
)

/**
 * [990] Verifying an Alien Dictionary
 *
 * In an alien language, surprisingly they also use english lowercase letters, but possibly in a different order. The order of the alphabet is some permutation of lowercase letters.
 *
 * Given a sequence of words written in the alien language, and the order of the alphabet, return true if and only if the given words are sorted lexicographicaly in this alien language.
 *
 *
 *
 *
 * Example 1:
 *
 *
 * Input: words = ["hello","leetcode"], order = "hlabcdefgijkmnopqrstuvwxyz"
 * Output: true
 * Explanation: As 'h' comes before 'l' in this language, then the sequence is sorted.
 *
 *
 *
 * Example 2:
 *
 *
 * Input: words = ["word","world","row"], order = "worldabcefghijkmnpqstuvxyz"
 * Output: false
 * Explanation: As 'd' comes after 'l' in this language, then words[0] > words[1], hence the sequence is unsorted.
 *
 *
 *
 * Example 3:
 *
 *
 * Input: words = ["apple","app"], order = "abcdefghijklmnopqrstuvwxyz"
 * Output: false
 * Explanation: The first three characters "app" match, and the second string is shorter (in size.) According to lexicographical rules "apple" > "app", because 'l' > '&empty;', where '&empty;' is defined as the blank character which is less than any other character (More info).
 *
 *
 *
 *
 * Note:
 *
 *
 * 	1 <= words.length <= 100
 * 	1 <= words[i].length <= 20
 * 	order.length == 26
 * 	All characters in words[i] and order are english lowercase letters.
 *
 *
 *
 *
 *
 */

func TestVerifyinganAlienDictionary(t *testing.T) {
	var cases = []struct {
		input  []string
		order  string
		output bool
	}{
		{
			input:  []string{"hello", "leetcode"},
			order:  "hlabcdefgijkmnopqrstuvwxyz",
			output: true,
		},
		{
			input:  []string{"word", "world", "row"},
			order:  "worldabcefghijkmnpqstuvxyz",
			output: false,
		},
		{
			input:  []string{"apple", "app"},
			order:  "abcdefghijklmnopqrstuvwxyz",
			output: false,
		},
	}
	for _, c := range cases {
		x := isAlienSorted(c.input, c.order)
		if x != c.output {
			t.Fail()
		}
	}
}

// submission codes start here

func isAlienSorted(words []string, order string) bool {
	m := make([]int, 26)
	for i, o := range order {
		m[o-'a'] = i
	}
	for i := 0; i< len(words)-1;i++{
		if !isWordInOrder(m, words[i], words[i+1]) {
			return false
		}
	}
	return true
}

func isWordInOrder(m []int, word1, word2 string) bool {
	i := 0
	for i < len(word1) && i < len(word2) && word1[i] == word2[i] {
		i++
	}
	if i == len(word1) {
		return true
	}
	if i == len(word2) {
		return false
	}
	if m[word1[i]-'a'] > m[word2[i]-'a'] {
		return false
	}
	return true
}

// submission codes end
