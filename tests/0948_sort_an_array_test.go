package tests

import (
	"reflect"
	"testing"
)

/**
 * [948] Sort an Array
 *
 * Given an array of integers nums, sort the array in ascending order.
 *
 *
 *
 *
 *
 *
 * Example 1:
 *
 *
 * Input: [5,2,3,1]
 * Output: [1,2,3,5]
 *
 *
 * Example 2:
 *
 *
 * Input: [5,1,1,2,0,0]
 * Output: [0,0,1,1,2,5]
 *
 *
 *
 *
 * Note:
 *
 *
 * 	1 <= A.length <= 10000
 * 	-50000 <= A[i] <= 50000
 *
 *
 */

func TestSortanArray(t *testing.T) {
	var cases = []struct {
		input  []int
		output []int
	}{
		{
			input:  []int{5, 2, 3, 1},
			output: []int{1, 2, 3, 5},
		},
		{
			input:  []int{5, 1, 1, 2, 0, 0},
			output: []int{0, 0, 1, 1, 2, 5},
		},
	}
	for _, c := range cases {
		x := sortArray(c.input)
		if !reflect.DeepEqual(x, c.output) {
			t.Fail()
		}
	}
}

// submission codes start here

func sortArray(nums []int) []int {
	return  quickSort(nums, 0, len(nums)-1)
}

func quickSort(nums []int, l, r int) []int {
	flag := l
	left := l
	right := r
	if l >= r {
		return nums
	}
	for l < r {
		for l < r && nums[r] >= nums[flag] {
			r--
		}
		nums[r], nums[flag] = nums[flag], nums[r]
		flag = r

		for l < r && nums[l] <= nums[flag] {
			l++
		}
		nums[l], nums[flag] = nums[flag], nums[l]
		flag = l
	}
	nums = quickSort(nums, left, flag-1)
	nums = quickSort(nums, flag+1, right)
	return nums
}

// submission codes end
