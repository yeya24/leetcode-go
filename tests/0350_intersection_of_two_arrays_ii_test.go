package tests

import (
	"reflect"
	"testing"
)

/**
 * [350] Intersection of Two Arrays II
 *
 * Given two arrays, write a function to compute their intersection.
 *
 * Example 1:
 *
 *
 * Input: nums1 = [1,2,2,1], nums2 = [2,2]
 * Output: [2,2]
 *
 *
 *
 * Example 2:
 *
 *
 * Input: nums1 = [4,9,5], nums2 = [9,4,9,8,4]
 * Output: [4,9]
 *
 *
 * Note:
 *
 *
 * 	Each element in the result should appear as many times as it shows in both arrays.
 * 	The result can be in any order.
 *
 *
 * Follow up:
 *
 *
 * 	What if the given array is already sorted? How would you optimize your algorithm?
 * 	What if nums1's size is small compared to nums2's size? Which algorithm is better?
 * 	What if elements of nums2 are stored on disk, and the memory is limited such that you cannot load all elements into the memory at once?
 *
 *
 */

func TestIntersectionofTwoArraysII(t *testing.T) {
	var cases = []struct {
		input  []int
		input1 []int
		output []int
	}{
		{
			input:  []int{1, 2, 2, 1},
			input1: []int{2, 2},
			output: []int{2, 2},
		},
		{
			input:  []int{4, 9, 5},
			input1: []int{9, 4, 9, 8, 4},
			output: []int{9, 4},
		},
	}
	for _, c := range cases {
		x := intersect(c.input, c.input1)
		if !reflect.DeepEqual(x, c.output) {
			t.Fail()
		}
	}
}

// submission codes start here

func intersect(nums1 []int, nums2 []int) []int {
	m1 := make(map[int]int)
	for _, n := range nums1 {
		m1[n] += 1
	}
	res := []int{}
	for _, n := range nums2 {
		if m1[n] > 0 {
			res = append(res, n)
			m1[n]--
		}
	}
	return res
}

// submission codes end
