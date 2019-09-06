package tests

import (
	"testing"
)

/**
 * [1160] Letter Tile Possibilities
 *
 * You have a set of tiles, where each tile has one letter tiles[i] printed on it.  Return the number of possible non-empty sequences of letters you can make.
 *
 *
 *
 * Example 1:
 *
 *
 * Input: "AAB"
 * Output: 8
 * Explanation: The possible sequences are "A", "B", "AA", "AB", "BA", "AAB", "ABA", "BAA".
 *
 *
 *
 * Example 2:
 *
 *
 * Input: "AAABBC"
 * Output: 188
 *
 *
 *
 *
 *
 * Note:
 *
 *
 * 	1 <= tiles.length <= 7
 * 	tiles consists of uppercase English letters.
 *
 */

func TestLetterTilePossibilities(t *testing.T) {
	var cases = []struct {
		input  string
		output int
	}{
		{
			input:  "AAB",
			output: 8,
		},
		{
			input:  "AAABBC",
			output: 188,
		},
	}
	for _, c := range cases {
		x := numTilePossibilities(c.input)
		if x != c.output {
			t.Fail()
		}
	}
}

// submission codes start here

func numTilePossibilities(tiles string) int {
	count := [26]int{}
	for _, t := range tiles {
		count[t-'A'] += 1
	}

	var dfs func() int
	dfs = func() int {
		sum := 0
		for i := 0; i < 26; i++ {
			if count[i] == 0 {
				continue
			}
			sum++
			count[i]--
			sum += dfs()
			count[i]++
		}
		return sum
	}
	return dfs()
}

// submission codes end
