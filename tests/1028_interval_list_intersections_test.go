package tests

import (
	"math"
	"reflect"
	"testing"
)

/**
 * [1028] Interval List Intersections
 *
 * Given two lists of closed intervals, each list of intervals is pairwise disjoint and in sorted order.
 *
 * Return the intersection of these two interval lists.
 *
 * (Formally, a closed interval [a, b] (with a <= b) denotes the set of real numbers x with a <= x <= b.  The intersection of two closed intervals is a set of real numbers that is either empty, or can be represented as a closed interval.  For example, the intersection of [1, 3] and [2, 4] is [2, 3].)
 *
 *
 *
 *
 * Example 1:
 *
 *
 *
 *
 * Input: A = [[1,5],[8,12],[15,24],[25,26]], B = [[1,5],[8,12],[15,24],[25,26]]
 * Output: [[1,2],[5,5],[8,10],[15,23],[24,24],[25,25]]
 * Reminder: The inputs and the desired output are lists of Interval objects, and not arrays or lists.
 *
 *
 *
 *
 * Note:
 *
 *
 * 	0 <= A.length < 1000
 * 	0 <= B.length < 1000
 * 	0 <= A[i].start, A[i].end, B[i].start, B[i].end < 10^9
 *
 *
 * NOTE: input types have been changed on April 15, 2019. Please reset to default code definition to get new method signature.
 *
 *
 */

func TestIntervalListIntersections(t *testing.T) {
	var cases = []struct {
		inputA [][]int
		inputB [][]int
		output [][]int
	}{
		{
			inputB: [][]int{{1, 5}, {8, 12}, {15, 24}, {25, 26}},
			inputA: [][]int{{0, 2}, {5, 10}, {13, 23}, {24, 25}},
			output: [][]int{{1, 2}, {5, 5}, {8, 10}, {15, 23}, {24, 24}, {25, 25}},
		},
	}
	for _, c := range cases {
		x := intervalIntersection(c.inputA, c.inputB)
		if !reflect.DeepEqual(x, c.output) {
			t.Fail()
		}
	}
}

// submission codes start here

func intervalIntersection(A [][]int, B [][]int) [][]int {
	i, j := 0, 0
	res := make([][]int, 0)
	for i < len(A) && j < len(B) {
		if A[i][1] >= B[j][0] && A[i][0] <= B[j][1] {
			res = append(res, []int{int(math.Max(float64(A[i][0]), float64(B[j][0]))),
				int(math.Min(float64(A[i][1]), float64(B[j][1])))})
		}
		if A[i][1] >= B[j][1] {
			j++
		} else {
			i++
		}
	}
	return res
}

// submission codes end
