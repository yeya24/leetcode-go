package tests

import (
    "strings"
    "testing"
)

/**
 * [125] Valid Palindrome
 *
 * Given a string, determine if it is a palindrome, considering only alphanumeric characters and ignoring cases.
 * 
 * Note: For the purpose of this problem, we define empty string as valid palindrome.
 * 
 * Example 1:
 * 
 * 
 * Input: "A man, a plan, a canal: Panama"
 * Output: true
 * 
 * 
 * Example 2:
 * 
 * 
 * Input: "race a car"
 * Output: false
 * 
 * 
 */

func TestValidPalindrome(t *testing.T) {
    var cases = []struct {
        input  string
        output bool
    }{
        //{
        //    input:  "A man, a plan, a canal: Panama",
        //    output: true,
        //},
        //{
        //    input:  "race a car",
        //    output: false,
        //},
        {
            input: "0P",
            output: false,
        },
    }
    for _, c := range cases {
        x := isPalindrome(c.input)
        if x != c.output {
            t.Fail()
        }
    }
}

// submission codes start here

func isPalindrome(s string) bool {
    x := strings.ToUpper(s)
    left := 0
    right := len(x) - 1

    isNumber := func(a uint8) bool {
        if a >= 48 && a <= 57 {
            return true
        }
        return false
    }
    isUpper := func(a uint8) bool {
        if a >= 65 && a <= 90 {
            return true
        }
        return false
    }

    for left < right {
        for left < right && !isUpper(x[left]) && !isNumber(x[left]) {
            left++
        }
        for left < right && !isUpper(x[right]) && !isNumber(x[right]) {
        	right--
        }
        if x[left] != x[right] {
            return false
        }
        left++
        right--
    }
	return true
}

// submission codes end
