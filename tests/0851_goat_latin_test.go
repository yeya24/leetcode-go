package tests

import (
	"strings"
	"testing"
)

/**
 * [851] Goat Latin
 *
 * A sentence S is given, composed of words separated by spaces. Each word consists of lowercase and uppercase letters only.
 *
 * We would like to convert the sentence to "Goat Latin" (a made-up language similar to Pig Latin.)
 *
 * The rules of Goat Latin are as follows:
 *
 *
 * 	If a word begins with a vowel (a, e, i, o, or u), append "ma" to the end of the word.
 * 	For example, the word 'apple' becomes 'applema'.
 *
 * 	If a word begins with a consonant (i.e. not a vowel), remove the first letter and append it to the end, then add "ma".
 * 	For example, the word "goat" becomes "oatgma".
 *
 * 	Add one letter 'a' to the end of each word per its word index in the sentence, starting with 1.
 * 	For example, the first word gets "a" added to the end, the second word gets "aa" added to the end and so on.
 *
 *
 * Return the final sentence representing the conversion from S to Goat Latin.
 *
 *
 *
 * Example 1:
 *
 *
 * Input: "I speak Goat Latin"
 * Output: "Imaa peaksmaaa oatGmaaaa atinLmaaaaa"
 *
 *
 * Example 2:
 *
 *
 * Input: "The quick brown fox jumped over the lazy dog"
 * Output: "heTmaa uickqmaaa rownbmaaaa oxfmaaaaa umpedjmaaaaaa overmaaaaaaa hetmaaaaaaaa azylmaaaaaaaaa ogdmaaaaaaaaaa"
 *
 *
 *
 *
 * Notes:
 *
 *
 * 	S contains only uppercase, lowercase and spaces. Exactly one space between each word.
 * 	1 <= S.length <= 150.
 *
 *
 */

func TestGoatLatin(t *testing.T) {
	var cases = []struct {
		input  string
		output string
	}{
		{
			input:  "I speak Goat Latin",
			output: "Imaa peaksmaaa oatGmaaaa atinLmaaaaa",
		},
		{
			input:  "The quick brown fox jumped over the lazy dog",
			output: "heTmaa uickqmaaa rownbmaaaa oxfmaaaaa umpedjmaaaaaa overmaaaaaaa hetmaaaaaaaa azylmaaaaaaaaa ogdmaaaaaaaaaa",
		},
	}
	for _, c := range cases {
		x := toGoatLatin(c.input)
		if x != c.output {
			t.Fail()
		}
	}
}

// submission codes start here

func toGoatLatin(S string) string {
	isVowel := func(word string) bool {
		switch word[0] {
		case 'i', 'I', 'a', 'A', 'o', 'O', 'e', 'E', 'u', 'U':
			return true
		}
		return false
	}
	words := strings.Split(S, " ")
	for i := 0; i < len(words); i++ {
		if !isVowel(words[i]) {
			words[i] = words[i][1:] + string(words[i][0])
		}
		words[i] += "ma"
		for j := 0; j <= i; j++ {
			words[i] += "a"
		}
	}
	return strings.Join(words, " ")
}

// submission codes end
