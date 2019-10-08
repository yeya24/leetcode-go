package tests

import (
    "testing"
)

/**
 * [442] Find All Duplicates in an Array
 *
 * Given an array of integers, 1 <= a[i] <= n (n = size of array), some elements appear twice and others appear once.
 * 
 * Find all the elements that appear twice in this array.
 * 
 * Could you do it without extra space and in O(n) runtime?
 * 
 * Example:
 * 
 * Input:
 * [4,3,2,7,8,2,3,1]
 * 
 * Output:
 * [2,3]
 * 
 */

func TestFindAllDuplicatesinanArray(t *testing.T) {

}

// submission codes start here

func findDuplicates(nums []int) []int {
    res := []int{}
    for i := 0; i< len(nums); i++ {
        index := abs(nums[i]) - 1
        if nums[index] < 0 {
            res = append(res, index + 1)
        } else {
            nums[index] = -nums[index]
        }
    }
    return res
}

// submission codes end
