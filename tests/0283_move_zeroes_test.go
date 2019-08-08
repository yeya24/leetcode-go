package tests

import (
	"testing"
)

/**
 * [283] Move Zeroes
 *
 * Given an array nums, write a function to move all 0's to the end of it while maintaining the relative order of the non-zero elements.
 *
 * Example:
 *
 *
 * Input: [0,1,0,3,12]
 * Output: [1,3,12,0,0]
 *
 * Note:
 *
 *
 * 	You must do this in-place without making a copy of the array.
 * 	Minimize the total number of operations.
 *
 */

func TestMoveZeroes(t *testing.T) {
	//var cases = []struct {
	//	input  []int
	//	output []int
	//}{
	//	{
	//		input:  []int{0,1,0,3,12},
	//		output: []int{1,3,12,0,0},
	//	},
	//	{
	//		input:  []int{0,0,1},
	//		output: []int{1,0,0},
	//	},
	//}
	moveZeroes([]int{10,0,1,0,3,12})
}

// submission codes start here

func moveZeroes(nums []int) {
	j := 0
	for i, v := range nums {
		if v != 0 {
			nums[i], nums[j] = nums[j], nums[i]
			j++
		}
	}
}

// submission codes end
