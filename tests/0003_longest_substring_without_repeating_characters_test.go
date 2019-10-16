package tests

import (
    "testing"
)

/**
 * [3] Longest Substring Without Repeating Characters
 *
 * Given a string, find the length of the longest substring without repeating characters.
 * 
 * 
 * Example 1:
 * 
 * 
 * Input: "abcabcbb"
 * Output: 3 
 * Explanation: The answer is "abc", with the length of 3. 
 * 
 * 
 * 
 * Example 2:
 * 
 * 
 * Input: "bbbbb"
 * Output: 1
 * Explanation: The answer is "b", with the length of 1.
 * 
 * 
 * 
 * Example 3:
 * 
 * 
 * Input: "pwwkew"
 * Output: 3
 * Explanation: The answer is "wke", with the length of 3. 
 *              Note that the answer must be a substring, "pwke" is a subsequence and not a substring.
 * 
 * 
 * 
 * 
 * 
 */

func TestLongestSubstringWithoutRepeatingCharacters(t *testing.T) {
    var cases = []struct {
        input  string
        output int
    }{
        {
            input:  "abcabcbb",
            output: 3,
        },
        {
            input:  "bbbbb",
            output: 1,
        },
        {
            input:  "pwwkew",
            output: 3,
        },
    }
    for _, c := range cases {
        x := lengthOfLongestSubstring(c.input)
        if x != c.output {
            t.Fail()
        }
    }
}

// submission codes start here

func lengthOfLongestSubstring(s string) int {
    n := len(s)
    m := make(map[byte]bool)
    ans, i, j := 0, 0, 0
    for i < n && j < n {
        if !m[s[j]] {
            m[s[j]]=true
            j++
            if j - i > ans {
                ans = j - i
            }
        } else {
            m[s[i]] = false
            i++
        }
    }
    return ans
}
// submission codes end
