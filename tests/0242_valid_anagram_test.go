package tests

import (
	"testing"
)

/**
 * [242] Valid Anagram
 *
 * Given two strings s and t , write a function to determine if t is an anagram of s.
 * 
 * Example 1:
 * 
 * 
 * Input: s = "anagram", t = "nagaram"
 * Output: true
 * 
 * 
 * Example 2:
 * 
 * 
 * Input: s = "rat", t = "car"
 * Output: false
 * 
 * 
 * Note:
 * You may assume the string contains only lowercase alphabets.
 * 
 * Follow up:
 * What if the inputs contain unicode characters? How would you adapt your solution to such case?
 * 
 */

func TestValidAnagram(t *testing.T) {
    var cases = []struct {
    	s string
    	t string
        output bool
    }{
        {
        	s: "anagram",
        	t: "nagaram",
        	output: true,
        },
        {
        	s: "rat",
        	t: "car",
        	output: false,
        },
    }
    for _, c := range cases {
        x := isAnagram(c.s, c.t)
        if x != c.output {
            t.Fail()
        }
    }
}

// submission codes start here

func isAnagram(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}
	charArr := make([]int, 26)
   	for _, v := range s {
   		charArr[v-'a'] += 1
   	}
   	for _, v := range t {
   		charArr[v-'a'] -= 1
	}
   	for _, v := range charArr {
   		if v != 0 {
   			return false
		}
	}
   	return true
}

// submission codes end
