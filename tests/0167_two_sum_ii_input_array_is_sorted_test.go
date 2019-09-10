package tests

import (
	"reflect"
	"testing"
)

/**
 * [167] Two Sum II - Input array is sorted
 *
 * Given an array of integers that is already sorted in ascending order, find two numbers such that they add up to a specific target number.
 *
 * The function twoSum should return indices of the two numbers such that they add up to the target, where index1 must be less than index2.
 *
 * Note:
 *
 *
 * 	Your returned answers (both index1 and index2) are not zero-based.
 * 	You may assume that each input would have exactly one solution and you may not use the same element twice.
 *
 *
 * Example:
 *
 *
 * Input: numbers = [2,7,11,15], target = 9
 * Output: [1,2]
 * Explanation: The sum of 2 and 7 is 9. Therefore index1 = 1, index2 = 2.
 *
 */

func TestTwoSumIIInputarrayissorted(t *testing.T) {
	var cases = []struct {
		input  []int
		target int
		output []int
	}{
		{
			input:  []int{2, 7, 11, 15},
			target: 9,
			output: []int{1, 2},
		},
	}
	for _, c := range cases {
		x := twoSumInputarrayissorted(c.input, c.target)
		if !reflect.DeepEqual(x, c.output) {
			t.Fail()
		}
	}
}

// submission codes start here

func twoSumInputarrayissorted(numbers []int, target int) []int {
	if len(numbers) == 0 {
		return nil
	}
	i, j := 0, len(numbers)-1
	for i < j {
		sum := numbers[i] + numbers[j]
		if sum == target {
			return []int{i + 1, j + 1}
		} else if sum > target {
			j--
		} else {
			i++
		}
	}
	return nil
}

// submission codes end
