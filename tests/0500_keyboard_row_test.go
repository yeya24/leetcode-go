package tests

import (
	"reflect"
	"testing"
)

/**
 * [500] Keyboard Row
 *
 * Given a List of words, return the words that can be typed using letters of alphabet on only one row's of American keyboard like the image below.
 *
 *
 *
 *
 *
 *
 * Example:
 *
 *
 * Input: ["Hello", "Alaska", "Dad", "Peace"]
 * Output: ["Alaska", "Dad"]
 *
 *
 *
 *
 * Note:
 *
 *
 * 	You may use one character in the keyboard more than once.
 * 	You may assume the input string will only contain letters of alphabet.
 *
 *
 */

func TestKeyboardRow(t *testing.T) {
	var cases = []struct {
		input  []string
		output []string
	}{
		{
			input:  []string{"Hello", "Alaska", "Dad", "Peace"},
			output: []string{"Alaska", "Dad"},
		},
	}
	for _, c := range cases {
		x := findWords(c.input)
		if !reflect.DeepEqual(x, c.output) {
			t.Fail()
		}
	}
}

// submission codes start here

func findWords(words []string) []string {
	ma := map[rune]struct{}{
		'q': {}, 'w': {}, 'e': {}, 'r': {}, 't': {}, 'y': {}, 'u': {}, 'i': {}, 'o': {}, 'p': {},
	}
	mb := map[rune]struct{}{
		'a': {}, 's': {}, 'd': {}, 'f': {}, 'g': {}, 'h': {}, 'j': {}, 'k': {}, 'l': {},
	}
	mc := map[rune]struct{}{
		'z': {}, 'x': {}, 'c': {}, 'v': {}, 'b': {}, 'n': {}, 'm': {},
	}
	res := make([]string, 0)

	for _, w := range words {
		f := -1
		flag := false
		for _, v := range w {
			if v <= 'Z' {
				v = v + ('a' - 'A')
			}
			if _, ok := ma[v]; ok {
				if f == 1 {
					continue
				} else if f != -1 {
					flag = true
					break
				}
				f = 1
			} else if _, ok := mb[v]; ok {
				if f == 2 {
					continue
				} else if f != -1 {
					flag = true
					break
				}
				f = 2
			} else if _, ok := mc[v]; ok {
				if f == 3 {
					continue
				} else if f != -1 {
					flag = true
					break
				}
				f = 3
			}
		}
		if !flag {
			res = append(res, w)
		}
	}
	return res
}

// submission codes end
