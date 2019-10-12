package tests

import (
	"testing"
)

/**
 * [300] Longest Increasing Subsequence
 *
 * Given an unsorted array of integers, find the length of longest increasing subsequence.
 *
 * Example:
 *
 *
 * Input: [10,9,2,5,3,7,101,18]
 * Output: 4
 * Explanation: The longest increasing subsequence is [2,3,7,101], therefore the length is 4.
 *
 * Note:
 *
 *
 * 	There may be more than one LIS combination, it is only necessary for you to return the length.
 * 	Your algorithm should run in O(n^2) complexity.
 *
 *
 * Follow up: Could you improve it to O(n log n) time complexity?
 *
 */

func TestLongestIncreasingSubsequence(t *testing.T) {
	var cases = []struct {
		input  []int
		output int
	}{
		{
			input:  []int{10, 9, 2, 5, 3, 7, 101, 18},
			output: 4,
		},
		{
			input:  []int{1, 3, 6, 7, 9, 4, 10, 5, 6},
			output: 6,
		},
		{
			input:  []int{2, 2},
			output: 1,
		},
	}
	for _, c := range cases {
		x := lengthOfLIS(c.input)
		if x != c.output {
			t.Fail()
		}
	}
}

// submission codes start here

// dp solution. time O(n^2)
func lengthOfLIS(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	m := make([]int, len(nums))
	for i := 0; i < len(nums); i++ {
		m[i] = 1
	}
	maxV := 1
	for i := 1; i<len(nums); i++ {
		for j := 0; j < i; j++ {
			if nums[j] < nums[i] {
				if m[i] < m[j]+1 {
					m[i] = m[j]+1
				}
			}
			if m[i] > maxV {
				maxV = m[i]
			}
		}
	}
	return maxV
}

// dp solution. Binary search. Copy from leetcode.
func lengthOfLIS2(nums []int) int {
	n:=len(nums)
	if n==0{
		return 0
	}
	s:=[]int{nums[0]}

	for _, v:=range nums{
		if v<s[0]{
			s[0]=v
		}else if v>s[len(s)-1] {
			s=append(s, v)
		}else{
			l, r:=0, len(s)-1
			for ;l<=r;{
				mid:=l+(r-l)/2
				if s[mid] < v{
					l=mid+1
				}else{
					r=mid-1
				}
			}
			s[l]=v
		}
	}

	return len(s)
}

// submission codes end
