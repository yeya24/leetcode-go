package tests

import (
	"testing"
)

/**
 * [162] Find Peak Element
 *
 * A peak element is an element that is greater than its neighbors.
 *
 * Given an input array nums, where nums[i] &ne; nums[i+1], find a peak element and return its index.
 *
 * The array may contain multiple peaks, in that case return the index to any one of the peaks is fine.
 *
 * You may imagine that nums[-1] = nums[n] = -&infin;.
 *
 * Example 1:
 *
 *
 * Input: nums = [1,2,3,1]
 * Output: 2
 * Explanation: 3 is a peak element and your function should return the index number 2.
 *
 * Example 2:
 *
 *
 * Input: nums = [1,2,1,3,5,6,4]
 * Output: 1 or 5
 * Explanation: Your function can return either index number 1 where the peak element is 2,
 *              or index number 5 where the peak element is 6.
 *
 *
 * Note:
 *
 * Your solution should be in logarithmic complexity.
 *
 */

func TestFindPeakElement(t *testing.T) {
	var cases = []struct {
		input  []int
		output int
	}{
		{
			input:  []int{1, 2, 3, 1},
			output: 2,
		},
		{
			input:  []int{1, 2, 1, 3, 5, 6, 4},
			output: 5,
		},
	}
	for _, c := range cases {
		x := findPeakElement(c.input)
		if x != c.output {
			t.Fail()
		}
	}
}

// submission codes start here

func findPeakElement(nums []int) int {
	l := len(nums)
	if l == 0 {
		return -1
	}
	i := 0
	j := l - 1
	for i < j {
		m := i + (j-i)/2
		if nums[m] > nums[m+1] {
			j = m
		} else {
			i = m + 1
		}
	}
	return i
}

func findPeakElementRecursive(nums []int) int {
	return findPeakRecursive(nums, 0, len(nums)-1)
}

func findPeakRecursive(nums []int, l, r int) int {
	if l == r {
		return l
	}
	m := l + (r-l)/2
	if nums[m] > nums[m+1] {
		return findPeakRecursive(nums, l, m)
	}
	return findPeakRecursive(nums, m+1, r)
}

// submission codes end
