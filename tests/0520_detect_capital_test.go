package tests

import (
	"testing"
)

/**
 * [520] Detect Capital
 *
 * Given a word, you need to judge whether the usage of capitals in it is right or not.
 *
 * We define the usage of capitals in a word to be right when one of the following cases holds:
 *
 *
 * 	All letters in this word are capitals, like "USA".
 * 	All letters in this word are not capitals, like "leetcode".
 * 	Only the first letter in this word is capital, like "Google".
 *
 * Otherwise, we define that this word doesn't use capitals in a right way.
 *
 *
 *
 * Example 1:
 *
 *
 * Input: "USA"
 * Output: True
 *
 *
 *
 *
 * Example 2:
 *
 *
 * Input: "FlaG"
 * Output: False
 *
 *
 *
 *
 * Note: The input will be a non-empty word consisting of uppercase and lowercase latin letters.
 *
 */

func TestDetectCapital(t *testing.T) {
    var cases = []struct {
        input  string
        output bool
    }{
        {
            input: "USA",
            output: true,
        },
        {
            input: "FlaG",
            output: false,
        },
    }
    for _, c := range cases {
        x := detectCapitalUse(c.input)
        if x != c.output {
            t.Fail()
        }
    }
}

// submission codes start here

func detectCapitalUse(word string) bool {
	var c int
	var v int
	isCapital := func(x uint8) bool {
		if x >= 65 && x <= 90 {
			return true
		}
		return false
	}

	if len(word) > 1 {
		for i := 0; i < len(word); i++ {
			if isCapital(word[i]) {
				c++
				if v > 0 {
					return false
				}
			} else {
				if c > 1 {
					return false
				}
				v++
			}
		}
	}
	return true
}

// submission codes end
