package tests

import (
	"testing"
)

/**
 * [797] Rabbits in Forest
 *
 * In a forest, each rabbit has some color. Some subset of rabbits (possibly all of them) tell you how many other rabbits have the same color as them. Those answers are placed in an array.
 *
 * Return the minimum number of rabbits that could be in the forest.
 *
 *
 * Examples:
 * Input: answers = [1, 1, 2]
 * Output: 5
 * Explanation:
 * The two rabbits that answered "1" could both be the same color, say red.
 * The rabbit than answered "2" can't be red or the answers would be inconsistent.
 * Say the rabbit that answered "2" was blue.
 * Then there should be 2 other blue rabbits in the forest that didn't answer into the array.
 * The smallest possible number of rabbits in the forest is therefore 5: 3 that answered plus 2 that didn't.
 *
 * Input: answers = [10, 10, 10]
 * Output: 11
 *
 * Input: answers = []
 * Output: 0
 *
 *
 * Note:
 *
 *
 * 	answers will have length at most 1000.
 * 	Each answers[i] will be an integer in the range [0, 999].
 *
 *
 */

func TestRabbitsinForest(t *testing.T) {
	var cases = []struct {
		input  []int
		output int
	}{
		{
			input:  []int{1, 1, 2},
			output: 5,
		},
		{
			input:  []int{10, 10, 10},
			output: 11,
		},
		{
			input:  []int{},
			output: 0,
		},
		{
			input:  []int{1, 0, 1, 0, 0},
			output: 5,
		},
		{
			input:  []int{0, 0, 1, 1, 1},
			output: 6,
		},
		{
			input:  []int{2, 1, 2, 2, 2, 2, 2, 2, 1, 1},
			output: 13,
		},
	}
	for _, c := range cases {
		x := numRabbits(c.input)
		if x != c.output {
			t.Fail()
		}
	}
}

// submission codes start here

func numRabbits(answers []int) int {
	if len(answers) == 0 {
		return 0
	}
	res := 0
	m := make(map[int]int)
	for _, v := range answers {
		if v == 0 {
			res += 1
		} else {
			if _, ok := m[v]; !ok {
				m[v] = 1
				res += v + 1
			} else {
				m[v] += 1
				if m[v] > v+1 {
					m[v] = 1
					res += v + 1
				}
			}
		}
	}
	return res
}

// submission codes end
