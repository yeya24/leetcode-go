package tests

import (
	"fmt"
	"sort"
	"testing"
)

/**
 * [15] 3Sum
 *
 * Given an array nums of n integers, are there elements a, b, c in nums such that a + b + c = 0? Find all unique triplets in the array which gives the sum of zero.
 *
 * Note:
 *
 * The solution set must not contain duplicate triplets.
 *
 * Example:
 *
 *
 * Given array nums = [-1, 0, 1, 2, -1, -4],
 *
 * A solution set is:
 * [
 *   [-1, 0, 1],
 *   [-1, -1, 2]
 * ]
 *
 *
 */

func Test3Sum(t *testing.T) {
	var cases = []struct {
		input  []int
		output [][]int
	}{
		{
			input:  []int{-1, 0, 1, 2, -1, -4},
			output: [][]int{{-1, 0, 1}, {-1, -1, 2}},
		},
		{
			input:  []int{0, 0, 0, 0},
			output: [][]int{{0, 0, 0}},
		},
		{
			input:  []int{-2, 0, 0, 2, 2},
			output: [][]int{{-2, 0, 2}},
		},
	}
	for _, c := range cases {
		x := threeSum(c.input)
		fmt.Println(x)
	}
}

// submission codes start here

func threeSum(nums []int) [][]int {
	sort.Ints(nums)
	var res [][]int
	for i := 0; i < len(nums); i++ {
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}
		target := -nums[i]
		left := i + 1
		right := len(nums) - 1
		for left < right {
			sum := nums[left] + nums[right]
			if sum == target {
				res = append(res, []int{nums[i], nums[left], nums[right]})
				for left < right && nums[left] == nums[left+1] {
					left++
				}
				for left < right && nums[right] == nums[right-1] {
					right--
				}
				left++
				right--
			} else if sum < target {
				left++
			} else {
				right--
			}
		}
	}
	return res
}

// submission codes end
