package tests

import (
	"math"
	"testing"
)

/**
 * [124] Binary Tree Maximum Path Sum
 *
 * Given a non-empty binary tree, find the maximum path sum.
 *
 * For this problem, a path is defined as any sequence of nodes from some starting node to any node in the tree along the parent-child connections. The path must contain at least one node and does not need to go through the root.
 *
 * Example 1:
 *
 *
 * Input: [1,2,3]
 *
 *        1
 *       / \
 *      2   3
 *
 * Output: 6
 *
 *
 * Example 2:
 *
 *
 * Input: [-10,9,20,null,null,15,7]
 *
 *    -10
 *    / \
 *   9  20
 *     /  \
 *    15   7
 *
 * Output: 42
 *
 *
 */

func TestBinaryTreeMaximumPathSum(t *testing.T) {
	x := &TreeNode{Val: 1, Right: &TreeNode{Val: 3, Right: &TreeNode{Val: 1}, Left: &TreeNode{Val: 2}}}
	y := &TreeNode{Val: -10, Left: &TreeNode{Val: 9}, Right: &TreeNode{Val: 20, Left: &TreeNode{Val: 15}, Right: &TreeNode{Val: 7}}}
	var cases = []struct {
		input  *TreeNode
		output int
	}{
		{
			input:  x,
			output: 6,
		},
		{
			input: y,
			output: 42,
		},
	}
	for _, c := range cases {
		x := maxPathSum(c.input)
		if x != c.output {
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
func maxPathSum(root *TreeNode) int {
	sum := math.MinInt64
	dfsMaxPathSum(root, &sum)
	return sum
}

func dfsMaxPathSum(root *TreeNode, sum *int) int {
	if root == nil {
		return 0
	}
	left := dfsMaxPathSum(root.Left, sum)
	right := dfsMaxPathSum(root.Right, sum)
	if left < 0 {
		left = 0
	}
	if right < 0 {
		right = 0
	}
	if root.Val+left+right > *sum {
		*sum = root.Val + left + right
	}
	return root.Val + max(left, right)
}

// submission codes end
