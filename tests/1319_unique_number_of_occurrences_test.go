package tests

import (
	"testing"
)

/**
 * [1319] Unique Number of Occurrences
 *
 * Given an array of integers arr, write a function that returns true if and only if the number of occurrences of each value in the array is unique.
 *
 *
 * Example 1:
 *
 *
 * Input: arr = [1,2,2,1,1,3]
 * Output: true
 * Explanation: The value 1 has 3 occurrences, 2 has 2 and 3 has 1. No two values have the same number of occurrences.
 *
 * Example 2:
 *
 *
 * Input: arr = [1,2]
 * Output: false
 *
 *
 * Example 3:
 *
 *
 * Input: arr = [-3,0,1,-3,1,1,1,-3,10,0]
 * Output: true
 *
 *
 *
 * Constraints:
 *
 *
 * 	1 <= arr.length <= 1000
 * 	-1000 <= arr[i] <= 1000
 *
 *
 */

func TestUniqueNumberofOccurrences(t *testing.T) {
	var cases = []struct {
		input  []int
		output bool
	}{
		{
			input:  []int{1, 2, 2, 1, 1, 3},
			output: true,
		},
		{
			input:  []int{-3, 0, 1, -3, 1, 1, 1, -3, 10, 0},
			output: true,
		},
		{
			input:  []int{1, 2},
			output: false,
		},
	}
	for _, c := range cases {
		x := uniqueOccurrences(c.input)
		if x != c.output {
			t.Fail()
		}
	}
}

// submission codes start here

func uniqueOccurrences(arr []int) bool {
	m := make([]int, 2001)
	for _, v := range arr {
		m[v+1000] += 1
	}
	x := make([]bool, len(arr))
	for _, v := range m {
		if v != 0 {
			if !x[v] {
				x[v] = true
			} else {
				return false
			}
		}
	}
	return true
}

// submission codes end
