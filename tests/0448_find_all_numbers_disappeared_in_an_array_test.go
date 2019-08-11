package tests

import (
	"math"
	"reflect"
	"testing"
)

/**
 * [448] Find All Numbers Disappeared in an Array
 *
 * Given an array of integers where 1 <= a[i] <= n (n = size of array), some elements appear twice and others appear once.
 *
 * Find all the elements of [1, n] inclusive that do not appear in this array.
 *
 * Could you do it without extra space and in O(n) runtime? You may assume the returned list does not count as extra space.
 *
 * Example:
 *
 * Input:
 * [4,3,2,7,8,2,3,1]
 *
 * Output:
 * [5,6]
 *
 *
 */

func TestFindAllNumbersDisappearedinanArray(t *testing.T) {
	var cases = []struct {
		input  []int
		output []int
	}{
		{
			input:  []int{4, 3, 2, 7, 8, 2, 3, 1},
			output: []int{5, 6},
		},
	}
	for _, c := range cases {
		x := findDisappearedNumbers(c.input)
		if !reflect.DeepEqual(x, c.output) {
			t.Fail()
		}
	}
}

// submission codes start here

func findDisappearedNumbers(nums []int) []int {
	res := make([]int, 0)
	for i := 0; i < len(nums); i++ {
		v := int(math.Abs(float64(nums[i]))) - 1
		if nums[v] > 0 {
			nums[v] = -nums[v]
		}
	}
	for i := 0; i < len(nums); i++ {
		if nums[i] > 0 {
			res = append(res, i+1)
		}
	}
	return res
}

// submission codes end
