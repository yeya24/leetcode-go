package tests

import (
    "testing"
)

/**
 * [680] Valid Palindrome II
 *
 * 
 * Given a non-empty string s, you may delete at most one character.  Judge whether you can make it a palindrome.
 * 
 * 
 * Example 1:
 * 
 * Input: "aba"
 * Output: True
 * 
 * 
 * 
 * Example 2:
 * 
 * Input: "abca"
 * Output: True
 * Explanation: You could delete the character 'c'.
 * 
 * 
 * 
 * Note:
 * 
 * The string will only contain lowercase characters a-z.
 * The maximum length of the string is 50000.
 * 
 * 
 */

func TestValidPalindromeII(t *testing.T) {
    var cases = []struct {
        input  string
        output bool
    }{
        {
            input:  "aba",
            output: true,
        },
        {
            input:  "abca",
            output: true,
        },
    }
    for _, c := range cases {
        x := validPalindrome(c.input)
        if x != c.output {
            t.Fail()
        }
    }
}

// submission codes start here

func validPalindrome(s string) bool {
    i := 0
    j := len(s) - 1

    isPalindrome := func(s string) bool {
        for i := 0; i < len(s); i++ {
            if s[i] != s[len(s) - 1 - i] {
                return false
            }
        }
        return true
    }

    for i < j {
        if s[i] != s[j] {
            return isPalindrome(s[i+1:j+1]) || isPalindrome(s[i:j])
        } else {
            i++
            j--
        }
    }
    return true
}

// submission codes end
