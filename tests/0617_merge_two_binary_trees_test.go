package tests

import (
	"reflect"
	"testing"
)

/**
 * [617] Merge Two Binary Trees
 *
 * Given two binary trees and imagine that when you put one of them to cover the other, some nodes of the two trees are overlapped while the others are not.
 *
 * You need to merge them into a new binary tree. The merge rule is that if two nodes overlap, then sum node values up as the new value of the merged node. Otherwise, the NOT null node will be used as the node of new tree.
 *
 * Example 1:
 *
 *
 * Input:
 * 	Tree 1                     Tree 2
 *           1                         2
 *          / \                       / \
 *         3   2                     1   3
 *        /                           \   \
 *       5                             4   7
 * Output:
 * Merged tree:
 * 	     3
 * 	    / \
 * 	   4   5
 * 	  / \   \
 * 	 5   4   7
 *
 *
 *
 *
 * Note: The merging process must start from the root nodes of both trees.
 *
 */

func TestMergeTwoBinaryTrees(t *testing.T) {
	x := &TreeNode{Val: 1, Left: &TreeNode{Val: 3, Left: &TreeNode{Val: 5}},
		Right: &TreeNode{Val: 2}}
	y := &TreeNode{Val: 2, Left: &TreeNode{Val: 1, Right: &TreeNode{Val: 4}},
		Right: &TreeNode{Val: 3, Right: &TreeNode{Val: 7}}}
	z := &TreeNode{Val: 3, Left: &TreeNode{Val: 4, Left: &TreeNode{Val: 5}, Right: &TreeNode{Val: 4}},
		Right: &TreeNode{Val: 5, Right: &TreeNode{Val: 7}}}
	var cases = []struct {
		input1 *TreeNode
		input2 *TreeNode
		output *TreeNode
	}{
		{
			input1: x,
			input2: y,
			output: z,
		},
	}
	for _, c := range cases {
		x := mergeTrees(c.input1, c.input2)
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
func mergeTrees(t1 *TreeNode, t2 *TreeNode) *TreeNode {
	if t1 == nil && t2 == nil {
		return nil
	} else if t1 == nil {
		return t2
	} else if t2 == nil {
		return t1
	}
	n := &TreeNode{Val: t1.Val + t2.Val}
	n.Left = mergeTrees(t1.Left, t2.Left)
	n.Right = mergeTrees(t1.Right, t2.Right)
	return n
}

// submission codes end
