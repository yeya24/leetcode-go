package tests

import (
	"reflect"
	"testing"
)

/**
 * [349] Intersection of Two Arrays
 *
 * Given two arrays, write a function to compute their intersection.
 *
 * Example 1:
 *
 *
 * Input: nums1 = [1,2,2,1], nums2 = [2,2]
 * Output: [2]
 *
 *
 *
 * Example 2:
 *
 *
 * Input: nums1 = [4,9,5], nums2 = [9,4,9,8,4]
 * Output: [9,4]
 *
 *
 * Note:
 *
 *
 * 	Each element in the result must be unique.
 * 	The result can be in any order.
 *
 *
 *
 *
 */

func TestIntersectionofTwoArrays(t *testing.T) {
	var cases = []struct {
		input  []int
		input1 []int
		output []int
	}{
		{
			input:  []int{1, 2, 2, 1},
			input1: []int{2, 2},
			output: []int{2},
		},
		{
			input:  []int{4, 9, 5},
			input1: []int{9, 4, 9, 8, 4},
			output: []int{4, 9},
		},
	}
	for _, c := range cases {
		x := intersection(c.input, c.input1)
		if !reflect.DeepEqual(x, c.output) {
			t.Fail()
		}
	}
}

// submission codes start here

func intersection(nums1 []int, nums2 []int) []int {
	m1 := make(map[int]struct{})
	for _, n := range nums1 {
		m1[n] = struct{}{}
	}
	m2 := make(map[int]struct{})
	for _, n := range nums2 {
		m2[n] = struct{}{}
	}
	res := make([]int, 0)
	for n := range m1 {
		if _, ok := m2[n]; ok {
			res = append(res, n)
		}
	}
	return res
}

// submission codes end
