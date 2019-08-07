package tests

import (
	"testing"
)

/**
 * [782] Jewels and Stones
 *
 * You're given strings J representing the types of stones that are jewels, and S representing the stones you have.  Each character in S is a type of stone you have.  You want to know how many of the stones you have are also jewels.
 *
 * The letters in J are guaranteed distinct, and all characters in J and S are letters. Letters are case sensitive, so "a" is considered a different type of stone from "A".
 *
 * Example 1:
 *
 *
 * Input: J = "aA", S = "aAAbbbb"
 * Output: 3
 *
 *
 * Example 2:
 *
 *
 * Input: J = "z", S = "ZZ"
 * Output: 0
 *
 *
 * Note:
 *
 *
 * 	S and J will consist of letters and have length at most 50.
 * 	The characters in J are distinct.
 *
 *
 */

func TestJewelsandStones(t *testing.T) {
	cases := []struct {
		S      string
		J      string
		expect int
	}{
		{
			S:      "aAAbbbb",
			J:      "aA",
			expect: 3,
		},
		{
			S:      "ZZ",
			J:      "z",
			expect: 0,
		},
	}
	for _, c := range cases {
		if numJewelsInStones(c.J, c.S) != c.expect {
			t.Fail()
		}
	}
}

// submission codes start here

func numJewelsInStones(J string, S string) int {
	m := make(map[int32]struct{}, len(J))
	for _, v := range J {
		m[v] = struct{}{}
	}
	var c int
	for _, v := range S {
		if _, ok :=m[v]; ok {
			c++
		}
	}
	return c
}

// submission codes end
