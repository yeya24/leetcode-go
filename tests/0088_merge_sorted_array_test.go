package tests

import (
	"testing"
)

/**
 * [88] Merge Sorted Array
 *
 * Given two sorted integer arrays nums1 and nums2, merge nums2 into nums1 as one sorted array.
 *
 * Note:
 *
 *
 * 	The number of elements initialized in nums1 and nums2 are m and n respectively.
 * 	You may assume that nums1 has enough space (size that is greater or equal to m + n) to hold additional elements from nums2.
 *
 *
 * Example:
 *
 *
 * Input:
 * nums1 = [1,2,3,0,0,0], m = 3
 * nums2 = [2,5,6],       n = 3
 *
 * Output: [1,2,2,3,5,6]
 *
 *
 */

func TestMergeSortedArray(t *testing.T) {
	var cases = []struct {
		input1 []int
		m      int
		input2 []int
		n      int
		output []int
	}{
		{
			input1: []int{1, 2, 3, 0, 0, 0},
			input2: []int{2, 5, 6},
			m:      3,
			n:      3,
			output: []int{1, 2, 2, 3, 5, 6},
		},
	}
	for _, c := range cases {
		merge(c.input1, c.n, c.input2, c.n)
	}
}

// submission codes start here

func merge(nums1 []int, m int, nums2 []int, n int) {
	if m == 0 {
		copy(nums1, nums2)
		return
	}

	// j is the index to put value in nums1
	j := m + n - 1
	for j >= 0 {
		if m >= 1 && n >= 1 {
			max1 := nums1[m-1]
			max2 := nums2[n-1]
			if max1 >= max2 {
				m--
				nums1[j] = max1
			} else {
				n--
				nums1[j] = max2
			}
		} else if n >= 1 {
			nums1[j] = nums2[n-1]
			n--
		} else if m >= 1 {
			nums1[j] = nums1[m-1]
			m--
		}
		j--
	}
}

// submission codes end
