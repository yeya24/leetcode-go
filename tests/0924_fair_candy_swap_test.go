package tests

import (
	"reflect"
	"testing"
)

/**
 * [924] Fair Candy Swap
 *
 * Alice and Bob have candy bars of different sizes: A[i] is the size of the i-th bar of candy that Alice has, and B[j] is the size of the j-th bar of candy that Bob has.
 *
 * Since they are friends, they would like to exchange one candy bar each so that after the exchange, they both have the same total amount of candy.  (The total amount of candy a person has is the sum of the sizes of candy bars they have.)
 *
 * Return an integer array ans where ans[0] is the size of the candy bar that Alice must exchange, and ans[1] is the size of the candy bar that Bob must exchange.
 *
 * If there are multiple answers, you may return any one of them.  It is guaranteed an answer exists.
 *
 *
 *
 *
 * Example 1:
 *
 *
 * Input: A = [1, 1], B = [2,2]
 * Output: [1,2]
 *
 *
 *
 * Example 2:
 *
 *
 * Input: A = [1,2], B = [2,3]
 * Output: [1,2]
 *
 *
 *
 * Example 3:
 *
 *
 * Input: A = [2], B = [1,3]
 * Output: [2,3]
 *
 *
 *
 * Example 4:
 *
 *
 * Input: A = [1,2,5], B = [2,4]
 * Output: [5,4]
 *
 *
 *
 *
 * Note:
 *
 *
 * 	1 <= A.length <= 10000
 * 	1 <= B.length <= 10000
 * 	1 <= A[i] <= 100000
 * 	1 <= B[i] <= 100000
 * 	It is guaranteed that Alice and Bob have different total amounts of candy.
 * 	It is guaranteed there exists an answer.
 *
 *
 *
 *
 *
 *
 */

func TestFairCandySwap(t *testing.T) {
	var cases = []struct {
		inputA []int
		inputB []int
		output []int
	}{
		{
			inputA: []int{1, 1},
			inputB: []int{2, 2},
			output: []int{1, 2},
		},
		{
			inputA: []int{1, 2},
			inputB: []int{2, 3},
			output: []int{1, 2},
		},
		{
			inputA: []int{1, 2, 5},
			inputB: []int{2, 4},
			output: []int{5, 4},
		},
	}
	for _, c := range cases {
		x := fairCandySwap(c.inputA, c.inputB)
		if !reflect.DeepEqual(x, c.output) {
			t.Fail()
		}
	}
}

// submission codes start here

func fairCandySwap(A []int, B []int) []int {
	sumA := 0
	sumB := 0
	for _, v := range A {
		sumA += v
	}
	m := make(map[int]struct{})
	for _, v := range B {
		sumB += v
		m[v] = struct{}{}
	}
	delta := (sumB - sumA) / 2
	for _, v := range A {
		if _, ok := m[v+delta]; ok {
			return []int{v, v + delta}
		}
	}
	return []int{}
}

// submission codes end
