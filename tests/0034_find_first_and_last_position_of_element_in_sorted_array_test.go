package tests

import (
    "reflect"
    "testing"
)

/**
 * [34] Find First and Last Position of Element in Sorted Array
 *
 * Given an array of integers nums sorted in ascending order, find the starting and ending position of a given target value.
 * 
 * Your algorithm's runtime complexity must be in the order of O(log n).
 * 
 * If the target is not found in the array, return [-1, -1].
 * 
 * Example 1:
 * 
 * 
 * Input: nums = [5,7,7,8,8,10], target = 8
 * Output: [3,4]
 * 
 * Example 2:
 * 
 * 
 * Input: nums = [5,7,7,8,8,10], target = 6
 * Output: [-1,-1]
 * 
 */

func TestFindFirstandLastPositionofElementinSortedArray(t *testing.T) {
    var cases = []struct {
        input  []int
        target int
        output []int
    }{
        {
            input:  []int{5,7,7,8,8,10},
            target: 8,
            output: []int{3,4},
        },
        {
            input:  []int{5,7,7,8,8,10},
            target: 6,
            output: []int{-1,-1},
        },
    }
    for _, c := range cases {
        x := searchRange(c.input, c.target)
        if !reflect.DeepEqual(x, c.output) {
            t.Fail()
        }
    }
}

// submission codes start here

func searchRange(nums []int, target int) []int {
    res := []int{-1,-1}
    res[0] = findFirst(nums, target)
    res[1] = findLast(nums, target)
    return res
}

func findFirst(nums []int, target int) int {
    index := -1
    i, j := 0, len(nums)-1
    var mid int
    for i <= j {
        mid = i + (j-i)/2
        if nums[mid] >= target {
            j = mid - 1
        } else {
            i = mid + 1
        }
        if nums[mid] == target {
            index = mid
        }
    }

    return index
}

func findLast(nums []int, target int) int {
    index := -1
    i, j := 0, len(nums)-1
    var mid int
    for i <= j {
        mid = i + (j-i)/2
        if nums[mid] <= target {
            i = mid + 1
        } else {
            j = mid - 1
        }
        if nums[mid] == target {
            index = mid
        }
    }

    return index
}

// submission codes end
