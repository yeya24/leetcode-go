package tests

import (
	"reflect"
	"testing"
)

/**
 * [654] Maximum Binary Tree
 *
 *
 * Given an integer array with no duplicates. A maximum tree building on this array is defined as follow:
 *
 * The root is the maximum number in the array.
 * The left subtree is the maximum tree constructed from left part subarray divided by the maximum number.
 * The right subtree is the maximum tree constructed from right part subarray divided by the maximum number.
 *
 *
 *
 *
 * Construct the maximum tree by the given array and output the root node of this tree.
 *
 *
 * Example 1:
 *
 * Input: [3,2,1,6,0,5]
 * Output: return the tree root node representing the following tree:
 *
 *       6
 *     /   \
 *    3     5
 *     \    /
 *      2  0
 *        \
 *         1
 *
 *
 *
 * Note:
 *
 * The size of the given array will be in the range [1,1000].
 *
 *
 */

func TestMaximumBinaryTree(t *testing.T) {
	x := &TreeNode{Val: 6, Left: &TreeNode{Val: 3, Right: &TreeNode{Val: 2, Right: &TreeNode{Val: 1}}},
		Right: &TreeNode{Val: 5, Left: &TreeNode{Val: 0}}}
	var cases = []struct {
		input  []int
		output *TreeNode
	}{
		{
			input:  []int{3, 2, 1, 6, 0, 5},
			output: x,
		},
	}
	for _, c := range cases {
		x := constructMaximumBinaryTree(c.input)
		if !reflect.DeepEqual(x, c.output) {
			t.Fail()
		}
	}
}

// submission codes start here

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

// recursion version
func constructMaximumBinaryTree(nums []int) *TreeNode {
	return constructMax(nums, 0, len(nums)-1)
}

func constructMax(nums []int, start, finish int) *TreeNode {
	if start > finish {
		return nil
	}
	maxIndex := finish
	for i := start; i <= finish; i++ {
		if nums[i] > nums[maxIndex] {
			maxIndex = i
		}
	}
	root := &TreeNode{Val: nums[maxIndex]}
	root.Left = constructMax(nums, start, maxIndex-1)
	root.Right = constructMax(nums, maxIndex+1, finish)
	return root
}

// submission codes end
