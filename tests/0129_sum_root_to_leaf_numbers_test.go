package tests

import (
	"reflect"
	"testing"
)

/**
 * [129] Sum Root to Leaf Numbers
 *
 * Given a binary tree containing digits from 0-9 only, each root-to-leaf path could represent a number.
 *
 * An example is the root-to-leaf path 1->2->3 which represents the number 123.
 *
 * Find the total sum of all root-to-leaf numbers.
 *
 * Note: A leaf is a node with no children.
 *
 * Example:
 *
 *
 * Input: [1,2,3]
 *     1
 *    / \
 *   2   3
 * Output: 25
 * Explanation:
 * The root-to-leaf path 1->2 represents the number 12.
 * The root-to-leaf path 1->3 represents the number 13.
 * Therefore, sum = 12 + 13 = 25.
 *
 * Example 2:
 *
 *
 * Input: [4,9,0,5,1]
 *     4
 *    / \
 *   9   0
 *  / \
 * 5   1
 * Output: 1026
 * Explanation:
 * The root-to-leaf path 4->9->5 represents the number 495.
 * The root-to-leaf path 4->9->1 represents the number 491.
 * The root-to-leaf path 4->0 represents the number 40.
 * Therefore, sum = 495 + 491 + 40 = 1026.
 *
 */

func TestSumRoottoLeafNumbers(t *testing.T) {
	x := &TreeNode{Val: 1, Left: &TreeNode{Val: 2}, Right: &TreeNode{Val: 3}}
	y := &TreeNode{Val: 4, Left: &TreeNode{Val: 9, Left: &TreeNode{Val: 5}, Right: &TreeNode{Val: 1}}, Right: &TreeNode{Val: 0}}
	var cases = []struct {
		input  *TreeNode
		sum    int
		output int
	}{
		{
			input:  x,
			output: 25,
		},
		{
			input:  y,
			output: 1026,
		},
	}
	for _, c := range cases {
		x := sumNumbers(c.input)
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
func sumNumbers(root *TreeNode) int {
	var cur, sum int
	dfsSumNumbers(root, cur, &sum)
	return sum
}

func dfsSumNumbers(root *TreeNode, cur int, sum *int) {
	if root == nil {
		return
	}
	cur = cur * 10 + root.Val
	if root.Left == nil && root.Right == nil {
		*sum += cur
	}
	dfsSumNumbers(root.Left, cur, sum)
	dfsSumNumbers(root.Right, cur, sum)
}

// submission codes end
