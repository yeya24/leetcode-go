package tests

import (
	"reflect"
	"testing"
)

/**
 * [1005] Univalued Binary Tree
 *
 * A binary tree is univalued if every node in the tree has the same value.
 *
 * Return true if and only if the given tree is univalued.
 *
 *
 *
 * Example 1:
 *
 *
 * Input: [1,1,1,1,1,null,1]
 * Output: true
 *
 *
 *
 * Example 2:
 *
 *
 * Input: [2,2,2,5,2]
 * Output: false
 *
 *
 *
 *
 *
 * Note:
 *
 *
 * 	The number of nodes in the given tree will be in the range [1, 100].
 * 	Each node's value will be an integer in the range [0, 99].
 *
 *
 */

func TestUnivaluedBinaryTree(t *testing.T) {
	x := &TreeNode{Val: 2, Left: &TreeNode{Val: 2, Left: &TreeNode{Val: 5}, Right: &TreeNode{Val: 2}},
		Right: &TreeNode{Val: 2}}
	var cases = []struct {
		input  *TreeNode
		output bool
	}{
		{
			input:  x,
			output: false,
		},
	}
	for _, c := range cases {
		x := isUnivalTree(c.input)
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
func isUnivalTree(root *TreeNode) bool {
	if root == nil {
		return true
	}
	if root.Left != nil {
		if root.Val != root.Left.Val {
			return false
		}
	}
	if root.Right != nil {
		if root.Val != root.Right.Val {
			return false
		}
	}
	return isUnivalTree(root.Left) && isUnivalTree(root.Right)
}

// submission codes end
