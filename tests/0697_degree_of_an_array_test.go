package tests

import (
	"math"
	"testing"
)

/**
 * [697] Degree of an Array
 *
 * Given a non-empty array of non-negative integers nums, the degree of this array is defined as the maximum frequency of any one of its elements.
 * Your task is to find the smallest possible length of a (contiguous) subarray of nums, that has the same degree as nums.
 *
 * Example 1:
 *
 * Input: [1, 2, 2, 3, 1]
 * Output: 2
 * Explanation:
 * The input array has a degree of 2 because both elements 1 and 2 appear twice.
 * Of the subarrays that have the same degree:
 * [1, 2, 2, 3, 1], [1, 2, 2, 3], [2, 2, 3, 1], [1, 2, 2], [2, 2, 3], [2, 2]
 * The shortest length is 2. So return 2.
 *
 *
 *
 *
 * Example 2:
 *
 * Input: [1,2,2,3,1,4,2]
 * Output: 6
 *
 *
 *
 * Note:
 * nums.length will be between 1 and 50,000.
 * nums[i] will be an integer between 0 and 49,999.
 *
 */

func TestDegreeofanArray(t *testing.T) {
	var cases = []struct {
		input  []int
		output int
	}{
		{
			input:  []int{1, 2, 2, 3, 1},
			output: 2,
		},
		{
			input:  []int{1, 2, 2, 3, 1, 4, 2},
			output: 6,
		},
		{
			input:  []int{2, 1, 1, 2, 1, 3, 3, 3, 1, 3, 1, 3, 2},
			output: 7,
		},
	}
	for _, c := range cases {
		x := findShortestSubArray(c.input)
		if x != c.output {
			t.Fail()
		}
	}
}

// submission codes start here

// 3 hashmaps, O(n)
func findShortestSubArray(nums []int) int {
	left := make(map[int]int)
	right := make(map[int]int)
	count := make(map[int]int)
	max := 0
	res := make([]int, 0)
	for i, n := range nums {
		if _, ok := left[n]; !ok {
			left[n] = i
		}
		right[n] = i
		count[n] += 1
	}
	for i, v := range count {
		if v > max {
			max = count[i]
			res = []int{i}
		} else if v == max {
			res = append(res, i)
		}
	}
	ans := len(nums)
	for _, v := range res {
		a := right[v] - left[v] + 1
		if a < ans {
			ans = a
		}
	}
	return ans
}

// Use one hashmap, time O(n*k), k is the number of keys that appear most times.
func findShortestSubArray2(nums []int) int {
	m := make(map[int]int)
	maxValue := 0
	for _, v := range nums {
		m[v] += 1
		if m[v] > maxValue {
			maxValue = m[v]
		}
	}
	res := []int{}
	for k, v := range m {
		if v == maxValue {
			res = append(res, k)
		}
	}
	ans := math.MaxInt64
	for _, r := range res {
		left := len(nums)
		right := 0
		for i, v := range nums {
			if v == r {
				if i < left {
					left = i
				}
				if i > right {
					right = i
				}
			}
		}
		if right - left + 1 < ans {
			ans = right-left+1
		}
	}
	return ans
}

// submission codes end
