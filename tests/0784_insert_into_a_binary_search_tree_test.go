package tests

import (
	"reflect"
	"testing"
)

/**
 * [784] Insert into a Binary Search Tree
 *
 * Given the root node of a binary search tree (BST) and a value to be inserted into the tree, insert the value into the BST. Return the root node of the BST after the insertion. It is guaranteed that the new value does not exist in the original BST.
 *
 * Note that there may exist multiple valid ways for the insertion, as long as the tree remains a BST after insertion. You can return any of them.
 *
 * For example,
 *
 *
 * Given the tree:
 *         4
 *        / \
 *       2   7
 *      / \
 *     1   3
 * And the value to insert: 5
 *
 *
 * You can return this binary search tree:
 *
 *
 *          4
 *        /   \
 *       2     7
 *      / \   /
 *     1   3 5
 *
 *
 * This tree is also valid:
 *
 *
 *          5
 *        /   \
 *       2     7
 *      / \
 *     1   3
 *          \
 *           4
 *
 *
 */

func TestInsertintoaBinarySearchTree(t *testing.T) {
	x := &TreeNode{Val: 4, Left: &TreeNode{Val: 2, Left: &TreeNode{Val: 1}, Right: &TreeNode{Val: 3}},
		Right: &TreeNode{Val: 7}}
	var cases = []struct {
		input  *TreeNode
		val    int
		output *TreeNode
	}{
		{
			input:  x,
			val: 5,
			output: &TreeNode{Val: 4, Left: &TreeNode{Val: 2, Left: &TreeNode{Val: 1}, Right: &TreeNode{Val: 3}},
				Right: &TreeNode{Val: 7, Left: &TreeNode{Val: 5}}},
		},
	}
	for _, c := range cases {
		x := insertIntoBST(c.input, c.val)
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
func insertIntoBST(root *TreeNode, val int) *TreeNode {
	if root == nil {
		return &TreeNode{Val: val}
	}
	if root.Val < val {
		if root.Right == nil {
			root.Right = &TreeNode{Val: val}
		} else {
			root.Right = insertIntoBST(root.Right, val)
		}
	} else {
		if root.Left == nil {
			root.Left = &TreeNode{Val: val}
		} else {
			root.Left = insertIntoBST(root.Left, val)
		}
	}
	return root
}

// submission codes end
