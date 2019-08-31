package tests

import (
    "testing"
)

/**
 * [1112] Find Words That Can Be Formed by Characters
 *
 * You are given an array of strings words and a string chars.
 * 
 * A string is good if it can be formed by characters from chars (each character can only be used once).
 * 
 * Return the sum of lengths of all good strings in words.
 * 
 *  
 * 
 * Example 1:
 * 
 * 
 * Input: words = ["cat","bt","hat","tree"], chars = "atach"
 * Output: 6
 * Explanation: 
 * The strings that can be formed are "cat" and "hat" so the answer is 3 + 3 = 6.
 * 
 * 
 * Example 2:
 * 
 * 
 * Input: words = ["hello","world","leetcode"], chars = "welldonehoneyr"
 * Output: 10
 * Explanation: 
 * The strings that can be formed are "hello" and "world" so the answer is 5 + 5 = 10.
 * 
 * 
 *  
 * 
 * Note:
 * 
 * 
 * 	1 <= words.length <= 1000
 * 	1 <= words[i].length, chars.length <= 100
 * 	All strings contain lowercase English letters only.
 * 
 */

func TestFindWordsThatCanBeFormedbyCharacters(t *testing.T) {
    var cases = []struct {
        input  []string
        chars string
        output int
    }{
        {
            input:  []string{"cat","bt","hat","tree"},
            chars: "atach",
            output: 6,
        },
        {
            input:  []string{"hello","world","leetcode"},
            chars: "welldonehoneyr",
            output: 10,
        },
    }
    for _, c := range cases {
        x := countCharacters(c.input, c.chars)
        if x != c.output {
            t.Fail()
        }
    }
}

// submission codes start here

func countCharacters(words []string, chars string) int {
    m := make(map[rune]int)
    res := 0
    for _, v := range chars {
        m[v]+=1
    }
    for _, v := range words {
        z := make(map[rune]int)
        res += len(v)
        for _, j := range v {
            z[j]+=1
            if z[j] > m[j] {
                res -= len(v)
                break
            }
        }

    }
    return res
}

// submission codes end
