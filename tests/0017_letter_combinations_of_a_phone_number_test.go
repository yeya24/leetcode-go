package tests

import (
    "reflect"
    "testing"
)

/**
 * [17] Letter Combinations of a Phone Number
 *
 * Given a string containing digits from 2-9 inclusive, return all possible letter combinations that the number could represent.
 * 
 * A mapping of digit to letters (just like on the telephone buttons) is given below. Note that 1 does not map to any letters.
 * 
 * 
 * 
 * Example:
 * 
 * 
 * Input: "23"
 * Output: ["ad", "ae", "af", "bd", "be", "bf", "cd", "ce", "cf"].
 * 
 * 
 * Note:
 * 
 * Although the above answer is in lexicographical order, your answer could be in any order you want.
 * 
 */

func TestLetterCombinationsofaPhoneNumber(t *testing.T) {
    var cases = []struct {
        input  string
        output []string
    }{
        {
            input:  "23",
            output: []string{"ad", "ae", "af", "bd", "be", "bf", "cd", "ce", "cf"},
        },
    }
    for _, c := range cases {
        x := letterCombinations(c.input)
        if !reflect.DeepEqual(x, c.output) {
            t.Fail()
        }
    }
}

// submission codes start here

func letterCombinations(digits string) []string {
    if len(digits) == 0 {
        return nil
    }
    m := [8][]int{}
    m[0]=[]int{0,1,2}
    m[1]=[]int{3,4,5}
    m[2]=[]int{6,7,8}
    m[3]=[]int{9,10,11}
    m[4]=[]int{12,13,14}
    m[5]=[]int{15,16,17,18}
    m[6]=[]int{19,20,21}
    m[7]=[]int{22,23,24,25}
    var res []string
    dfsLetterCombinations(m, 0, digits, "", &res)
    return res
}

func dfsLetterCombinations(m [8][]int, s int, digits, cur string, res *[]string) {
    if len(cur) == len(digits) {
        *res = append(*res, cur)
        return
    }
    for _, v := range m[digits[s]-'0'-2] {
        cur += string('a'+v)
        dfsLetterCombinations(m, s+1, digits, cur, res)
        cur = cur[:len(cur)-1]
    }
}

// submission codes end
