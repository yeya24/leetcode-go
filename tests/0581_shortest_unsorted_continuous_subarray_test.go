package tests

import (
    "sort"
    "testing"
)

/**
 * [581] Shortest Unsorted Continuous Subarray
 *
 * Given an integer array, you need to find one continuous subarray that if you only sort this subarray in ascending order, then the whole array will be sorted in ascending order, too.  
 * 
 * You need to find the shortest such subarray and output its length.
 * 
 * Example 1:
 * 
 * Input: [2, 6, 4, 8, 10, 9, 15]
 * Output: 5
 * Explanation: You need to sort [6, 4, 8, 10, 9] in ascending order to make the whole array sorted in ascending order.
 * 
 * 
 * 
 * Note:
 * 
 * Then length of the input array is in range [1, 10,000].
 * The input array may contain duplicates, so ascending order here means <=. 
 * 
 * 
 */

func TestShortestUnsortedContinuousSubarray(t *testing.T) {
    var cases = []struct {
        input  []int
        output int
    }{
        //{
        //    input:  []int{2, 6, 4, 8, 10, 9, 15},
        //    output: 5,
        //},
        {
            input:  []int{1,3,5,4,2},
            output: 4,
        },
    }
    for _, c := range cases {
        x := findUnsortedSubarray2(c.input)
        if x != c.output {
            t.Fail()
        }
    }
}

// submission codes start here

// use sort. Time O(NlogN) Space O(N)
func findUnsortedSubarray(nums []int) int {
    x := make([]int, len(nums))
    for i:=0; i<len(nums);i++{
        x[i]=nums[i]
    }
    sort.Ints(x)
    start, end := len(nums), 0
    for i := 0; i < len(nums); i++ {
        if nums[i] != x[i] {
            if i < start {
                start = i
            }
            if i > end {
                end = i
            }
        }
    }
    if end > start {
        return end - start + 1
    }
    return 0
}

// Monotonic stack Time O(N) Space O(N)
func findUnsortedSubarray2(nums []int) int {
    stack := []int{}
    l, r := len(nums), 0
    for i := 0; i < len(nums); i++ {
        for len(stack) > 0 && nums[stack[len(stack)-1]] > nums[i] {
            p := stack[len(stack)-1]
            stack = stack[:len(stack)-1]
            if p < l {
                l = p
            }
        }
        stack = append(stack, i)
    }

    stack = []int{}
    for i := len(nums)-1; i >= 0; i-- {
        for len(stack) > 0 && nums[stack[len(stack)-1]] < nums[i] {
            p := stack[len(stack)-1]
            stack = stack[:len(stack)-1]
            if p > r {
                r = p
            }
        }
        stack = append(stack, i)
    }
    if r > l {
        return r-l+1
    }
    return 0
}

// submission codes end
