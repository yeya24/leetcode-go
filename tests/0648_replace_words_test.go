package tests

import (
    "strings"
    "testing"
)

/**
 * [648] Replace Words
 *
 * In English, we have a concept called root, which can be followed by some other words to form another longer word - let's call this word successor. For example, the root an, followed by other, which can form another word another.
 * 
 * Now, given a dictionary consisting of many roots and a sentence. You need to replace all the successor in the sentence with the root forming it. If a successor has many roots can form it, replace it with the root with the shortest length.
 * 
 * You need to output the sentence after the replacement.
 * 
 * Example 1:
 * 
 * 
 * Input: dict = ["cat", "bat", "rat"]
 * sentence = "the cattle was rattled by the battery"
 * Output: "the cat was rat by the bat"
 * 
 * 
 *  
 * 
 * Note:
 * 
 * 
 * 	The input will only have lower-case letters.
 * 	1 <= dict words number <= 1000
 * 	1 <= sentence words number <= 1000
 * 	1 <= root length <= 100
 * 	1 <= sentence words length <= 1000
 * 
 * 
 *  
 * 
 */

func TestReplaceWords(t *testing.T) {
    var cases = []struct {
        input  []string
        sentence string
        output string
    }{
        {
            input:  []string{"cat", "bat", "rat"},
            sentence: "the cattle was rattled by the battery",
            output: "the cat was rat by the bat",
        },
        {
          input: []string{"a", "b", "c"},
          sentence: "aadsfasf absbs bbab cadsfafs",
          output: "a a b c",
        },
    }
    for _, c := range cases {
        x := replaceWords(c.input, c.sentence)
        if x != c.output {
            t.Fail()
        }
    }
}

// submission codes start here

func replaceWords(dict []string, sentence string) string {
    strs := strings.Split(sentence, " ")
    for i := 0; i < len(strs); i++ {
    	strs[i] = traverseDict(strs[i], dict)
    }
    return strings.Join(strs, " ")
}

func traverseDict(word string, dict []string) string {
    l := len(word)
    w := word
    for _, v := range dict {
        if strings.HasPrefix(word, v) {
            if len(v) < l {
                l = len(v)
                w = v
            }
        }
    }
    return w
}

// submission codes end
