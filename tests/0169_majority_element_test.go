package tests

import (
	"testing"
)

/**
 * [169] Majority Element
 *
 * Given an array of size n, find the majority element. The majority element is the element that appears more than &lfloor; n/2 &rfloor; times.
 *
 * You may assume that the array is non-empty and the majority element always exist in the array.
 *
 * Example 1:
 *
 *
 * Input: [3,2,3]
 * Output: 3
 *
 * Example 2:
 *
 *
 * Input: [2,2,1,1,1,2,2]
 * Output: 2
 *
 *
 */

func TestMajorityElement(t *testing.T) {
	var cases = []struct {
		input  []int
		output int
	}{
		{
			input:  []int{3, 2, 3},
			output: 3,
		},
		{
			input:  []int{2, 2, 1, 1, 1, 2, 2},
			output: 2,
		},
	}
	for _, c := range cases {
		x := majorityElement(c.input)
		if x != c.output {
			t.Fail()
		}
	}
}

// submission codes start here

func majorityElement(nums []int) int {
	major, count := nums[0], 0
	for _, v := range nums {
		if count == 0 {
			major = v
			count++
		} else if major == v {
			count++
		} else {
			count--
		}
	}
	return major
}

// submission codes end
