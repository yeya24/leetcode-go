package tests

import (
	"reflect"
	"testing"
)

/**
 * [102] Binary Tree Level Order Traversal
 *
 * Given a binary tree, return the level order traversal of its nodes' values. (ie, from left to right, level by level).
 *
 *
 * For example:
 * Given binary tree [3,9,20,null,null,15,7],
 *
 *     3
 *    / \
 *   9  20
 *     /  \
 *    15   7
 *
 *
 *
 * return its level order traversal as:
 *
 * [
 *   [3],
 *   [9,20],
 *   [15,7]
 * ]
 *
 *
 */

func TestBinaryTreeLevelOrderTraversal(t *testing.T) {
	x := &TreeNode{Val: 3, Left: &TreeNode{Val: 9},
		Right: &TreeNode{Val: 20, Left: &TreeNode{Val: 15}, Right: &TreeNode{Val: 7}}}
	var cases = []struct {
		input  *TreeNode
		output [][]int
	}{
		{
			input:  x,
			output: [][]int{{3},{9,20},{15,7}},
		},
	}
	for _, c := range cases {
		x := levelOrder(c.input)
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
// use container.list as queue
func levelOrder(root *TreeNode) [][]int {
	if root == nil {
		return nil
	}
	res := make([][]int, 0)
	cur := []*TreeNode{root}
	for len(cur) != 0 {
		var tmp []int
		var next []*TreeNode
		for _, n := range cur {
			tmp = append(tmp, n.Val)
			if n.Left != nil {
				next = append(next, n.Left)
			}
			if n.Right != nil {
				next = append(next, n.Right)
			}
		}
		res = append(res, tmp)
		cur = next
	}
	return res
}

// submission codes end
