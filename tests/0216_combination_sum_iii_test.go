package tests

import (
	"fmt"
	"testing"
)

/**
 * [216] Combination Sum III
 *
 *
 * Find all possible combinations of k numbers that add up to a number n, given that only numbers from 1 to 9 can be used and each combination should be a unique set of numbers.
 *
 * Note:
 *
 *
 * 	All numbers will be positive integers.
 * 	The solution set must not contain duplicate combinations.
 *
 *
 * Example 1:
 *
 *
 * Input: k = 3, n = 7
 * Output: [[1,2,4]]
 *
 *
 * Example 2:
 *
 *
 * Input: k = 3, n = 9
 * Output: [[1,2,6], [1,3,5], [2,3,4]]
 *
 *
 */

func TestCombinationSumIII(t *testing.T) {
	var cases = []struct {
		n      int
		k      int
		output [][]int
	}{
		{
		   n: 7,
		   k: 3,
		   output: [][]int{{1,2,4}},
		},
		{
			n:      9,
			k:      3,
			output: [][]int{{1, 2, 6}, {1, 3, 5}, {2, 3, 4}},
		},
	}
	for _, c := range cases {
		x := combinationSum3(c.k, c.n)
		fmt.Println(x)
		//if !reflect.DeepEqual(x, c.output) {
		//    t.Fail()
		//}
	}
}

// submission codes start here

func combinationSum3(k int, n int) [][]int {
	var ans [][]int
	var cur []int
	dfsCombinationSum3(k, n, 1, cur, &ans)
	return ans
}

func dfsCombinationSum3(k, n, s int, cur []int, ans *[][]int) {
	if len(cur) == k && n == 0 {
		*ans = append(*ans, cur)
		return
	}

	for i := s; i <= min(n, 9); i++ {
		cur = append(cur, i)
		dfsCombinationSum3(k, n-i, i+1, append([]int{}, cur...), ans)
		cur = cur[:len(cur)-1]
	}
}

// submission codes end
