package tests

import (
	"reflect"
	"testing"
)

/**
 * [238] Product of Array Except Self
 *
 * Given an array nums of n integers where n > 1,  return an array output such that output[i] is equal to the product of all the elements of nums except nums[i].
 *
 * Example:
 *
 *
 * Input:  [1,2,3,4]
 * Output: [24,12,8,6]
 *
 *
 * Note: Please solve it without division and in O(n).
 *
 * Follow up:
 * Could you solve it with constant space complexity? (The output array does not count as extra space for the purpose of space complexity analysis.)
 *
 */

func TestProductofArrayExceptSelf(t *testing.T) {
	var cases = []struct {
		input  []int
		output []int
	}{
		{
			input:  []int{1, 2, 3, 4},
			output: []int{24, 12, 8, 6},
		},
	}
	for _, c := range cases {
		x := productExceptSelf(c.input)
		if !reflect.DeepEqual(x, c.output) {
			t.Fail()
		}
	}
}

// submission codes start here

func productExceptSelf(nums []int) []int {
	l := len(nums)
	res := make([]int, l)
	res[0] = 1
	for i := 1; i < l; i++ {
		res[i] = res[i-1] * nums[i-1]
	}
	R := 1
	for i := l - 2; i >= 0; i-- {
		R *= nums[i+1]
		res[i] = res[i] * R
	}
	return res
}

// submission codes end
