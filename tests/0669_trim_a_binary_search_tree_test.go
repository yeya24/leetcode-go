package tests

import (
	"fmt"
	"testing"
)

/**
 * [669] Trim a Binary Search Tree
 *
 *
 * Given a binary search tree and the lowest and highest boundaries as L and R, trim the tree so that all its elements lies in [L, R] (R >= L). You might need to change the root of the tree, so the result should return the new root of the trimmed binary search tree.
 *
 *
 * Example 1:
 *
 * Input:
 *     1
 *    / \
 *   0   2
 *
 *   L = 1
 *   R = 2
 *
 * Output:
 *     1
 *       \
 *        2
 *
 *
 *
 * Example 2:
 *
 * Input:
 *     3
 *    / \
 *   0   4
 *    \
 *     2
 *    /
 *   1
 *
 *   L = 1
 *   R = 3
 *
 * Output:
 *       3
 *      /
 *    2
 *   /
 *  1
 *
 *
 */

func TestTrimaBinarySearchTree(t *testing.T) {
	var cases = []struct {
		input  *TreeNode
		L      int
		R      int
		output *TreeNode
	}{
		{
			input:  &TreeNode{Val: 1, Left: &TreeNode{Val: 0}, Right: &TreeNode{Val: 2}},
			L:      1,
			R:      2,
			output: &TreeNode{Val: 1, Right: &TreeNode{Val: 2}},
		},
		{
			input:  &TreeNode{Val: 3, Left: &TreeNode{Val: 0, Right: &TreeNode{Val: 2, Left: &TreeNode{Val: 1}}}, Right: &TreeNode{Val: 4}},
			L:      1,
			R:      3,
			output: &TreeNode{Val: 3, Left: &TreeNode{Val: 2, Left: &TreeNode{Val: 1}}},
		},
	}
	for _, c := range cases {
		x := trimBST(c.input, c.L, c.R)
		fmt.Print(x)
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
func trimBST(root *TreeNode, L int, R int) *TreeNode {
	if root == nil {
		return nil
	}
	if root.Val > R {
		return trimBST(root.Left, L, R)
	}
	if root.Val < L {
		return trimBST(root.Right, L, R)
	}
	root.Left = trimBST(root.Left, L, R)
	root.Right = trimBST(root.Right, L, R)
	return root
}

// submission codes end
