package tests

import (
	"testing"
)

/**
 * [792] Binary Search
 *
 * Given a sorted (in ascending order) integer array nums of n elements and a target value, write a function to search target in nums. If target exists, then return its index, otherwise return -1.
 *
 *
 * Example 1:
 *
 *
 * Input: nums = [-1,0,3,5,9,12], target = 9
 * Output: 4
 * Explanation: 9 exists in nums and its index is 4
 *
 *
 *
 * Example 2:
 *
 *
 * Input: nums = [-1,0,3,5,9,12], target = 2
 * Output: -1
 * Explanation: 2 does not exist in nums so return -1
 *
 *
 *
 *
 * Note:
 *
 *
 * 	You may assume that all elements in nums are unique.
 * 	n will be in the range [1, 10000].
 * 	The value of each element in nums will be in the range [-9999, 9999].
 *
 *
 */

func TestBinarySearch(t *testing.T) {
	var cases = []struct {
		input  []int
		target int
		output int
	}{
		{
			input:  []int{-1, 0, 3, 5, 9, 12},
			target: 9,
			output: 4,
		},
		{
			input:  []int{-1, 0, 3, 5, 9, 12},
			target: 2,
			output: -1,
		},
		{
			input:  []int{2},
			target: 2,
			output: 0,
		},
	}
	for _, c := range cases {
		x := search(c.input, c.target)
		if x != c.output {
			t.Fail()
		}
	}
}

// submission codes start here

func search(nums []int, target int) int {
	start := 0
	end := len(nums) - 1
	// <= for len(nums) == 1
	for start <= end {
		mid := start + (end-start)/2
		if nums[mid] == target {
			return mid
		} else if nums[mid] < target {
			start = mid + 1
		} else {
			end = mid - 1
		}
	}
	return -1
}

// submission codes end
