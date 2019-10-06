package tests

import (
    "testing"
)

/**
 * [1297] Maximum Number of Balloons
 *
 * Given a string text, you want to use the characters of text to form as many instances of the word "balloon" as possible.
 * 
 * You can use each character in text at most once. Return the maximum number of instances that can be formed.
 * 
 *  
 * Example 1:
 * 
 * 
 * 
 * 
 * Input: text = "nlaebolko"
 * Output: 1
 * 
 * 
 * Example 2:
 * 
 * 
 * 
 * 
 * Input: text = "loonbalxballpoon"
 * Output: 2
 * 
 * 
 * Example 3:
 * 
 * 
 * Input: text = "leetcode"
 * Output: 0
 * 
 * 
 *  
 * Constraints:
 * 
 * 
 * 	1 <= text.length <= 10^4
 * 	text consists of lower case English letters only.
 * 
 */

func TestMaximumNumberofBalloons(t *testing.T) {

}

// submission codes start here

func maxNumberOfBalloons(text string) int {
    m := [26]int{}
    for _, v := range text {
        m[v-'a'] += 1
    }

    minBalloon := 10001
    for _, v := range []int{m[0],m[1],m[11]/2,m[14]/2,m[13]} {
        minBalloon = min(minBalloon, v)
    }
    return minBalloon
}

// submission codes end
