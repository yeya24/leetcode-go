package tests

import (
	"fmt"
	"reflect"
	"testing"
)

/**
 * [46] Permutations
 *
 * Given a collection of distinct integers, return all possible permutations.
 *
 * Example:
 *
 *
 * Input: [1,2,3]
 * Output:
 * [
 *   [1,2,3],
 *   [1,3,2],
 *   [2,1,3],
 *   [2,3,1],
 *   [3,1,2],
 *   [3,2,1]
 * ]
 *
 *
 */

func TestPermutations(t *testing.T) {
	var cases = []struct {
		input  []int
		output [][]int
	}{
		{
			input:  []int{1, 2, 3},
			output: [][]int{{1, 2, 3}, {1, 3, 2}, {2, 1, 3}, {2, 3, 1}, {3, 1, 2}, {3, 2, 1}},
		},
	}
	for _, c := range cases {
		x := permute(c.input)
		fmt.Println(x)
		if !reflect.DeepEqual(x, c.output) {
			t.Fail()
		}
	}
}

// submission codes start here

func permute(nums []int) [][]int {
	var res [][]int
	var cur []int
	used := make([]bool, len(nums))
	dfsPermute(nums, used, cur, &res)
	return res
}

func dfsPermute(nums []int, used []bool, cur []int, res *[][]int) {
	if len(cur) == len(nums) {
		*res = append(*res, cur)
		return
	}

	for i := 0; i < len(nums); i++ {
		if used[i] {
			continue
		}
		used[i] = true
		cur = append(cur, nums[i])
		dfsPermute(nums, used, append([]int{}, cur...), res)
		cur = cur[:len(cur)-1]
		used[i] = false
	}
}

// submission codes end
