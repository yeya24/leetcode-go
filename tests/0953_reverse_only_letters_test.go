package tests

import (
	"testing"
)

/**
 * [953] Reverse Only Letters
 *
 * Given a string S, return the "reversed" string where all characters that are not a letter stay in the same place, and all letters reverse their positions.
 *
 *
 *
 *
 *
 *
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
 * Input: "ab-cd"
 * Output: "dc-ba"
 *
 *
 *
 * Example 2:
 *
 *
 * Input: "a-bC-dEf-ghIj"
 * Output: "j-Ih-gfE-dCba"
 *
 *
 *
 * Example 3:
 *
 *
 * Input: "Test1ng-Leet=code-Q!"
 * Output: "Qedo1ct-eeLg=ntse-T!"
 *
 *
 *
 *
 *
 * Note:
 *
 *
 * 	S.length <= 100
 * 	33 <= S[i].ASCIIcode <= 122
 * 	S doesn't contain \ or "
 *
 *
 *
 *
 *
 */

func TestReverseOnlyLetters(t *testing.T) {
	var cases = []struct {
		input  string
		output string
	}{
		{
			input:  "ab-cd",
			output: "dc-ba",
		},
		{
			input:  "a-bC-dEf-ghIj",
			output: "j-Ih-gfE-dCba",
		},
		{
			input:  "Test1ng-Leet=code-Q!",
			output: "Qedo1ct-eeLg=ntse-T!",
		},
		{
			input:  "7_28]",
			output: "7_28]",
		},
		{
			input:  "AaW;c?[",
			output: "cWa;A?[",
		},
		{
			input:  "V`me[,Vk_z2#F_PDQZKtmzdz6#JQj*",
			output: "j`QJ[,zd_z2#m_tKZQDPFzkV6#emV*",
		},
	}
	for _, c := range cases {
		x := reverseOnlyLetters(c.input)
		if x != c.output {
			t.Fail()
		}
	}
}

// submission codes start here

func reverseOnlyLetters(S string) string {
	l := len(S)
	start := 0
	end := l - 1
	s := []byte(S)
	for start < end {
		for start < end && !(s[start] >= 65 && s[start] <= 90 || s[start] >= 97 && s[start] <= 122) {
			start += 1
		}

		for start < end && !(s[end] >= 65 && s[end] <= 90 || s[end] >= 97 && s[end] <= 122) {
			end -= 1
		}
		s[start], s[end] = s[end], s[start]
		start += 1
		end -= 1
	}
	return string(s)
}

// submission codes end
