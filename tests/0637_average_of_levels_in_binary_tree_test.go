package tests

import (
	"reflect"
	"testing"
)

/**
 * [637] Average of Levels in Binary Tree
 *
 * Given a non-empty binary tree, return the average value of the nodes on each level in the form of an array.
 *
 * Example 1:
 *
 * Input:
 *     3
 *    / \
 *   9  20
 *     /  \
 *    15   7
 * Output: [3, 14.5, 11]
 * Explanation:
 * The average value of nodes on level 0 is 3,  on level 1 is 14.5, and on level 2 is 11. Hence return [3, 14.5, 11].
 *
 *
 *
 * Note:
 *
 * The range of node's value is in the range of 32-bit signed integer.
 *
 *
 */

func TestAverageofLevelsinBinaryTree(t *testing.T) {
	x := &TreeNode{Val: 3, Left: &TreeNode{Val: 9},
		Right: &TreeNode{Val: 20, Left: &TreeNode{Val: 15}, Right: &TreeNode{Val: 7}}}
	var cases = []struct {
		input  *TreeNode
		output []float64
	}{
		{
			input:  x,
			output: []float64{3, 14.5, 11},
		},
	}
	for _, c := range cases {
		x := averageOfLevels(c.input)
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
func averageOfLevels(root *TreeNode) []float64 {
	if root == nil {
		return nil
	}
	res := []float64{}
	cur := []*TreeNode{root}
	for len(cur) > 0 {
		levelSum := 0
		next := []*TreeNode{}
		var i float64
		for _, n := range cur {
			levelSum += n.Val
			i++
			if n.Left != nil {
				next = append(next, n.Left)
			}
			if n.Right != nil {
				next = append(next, n.Right)
			}
		}
		cur = next
		res = append(res, float64(levelSum)/i)
	}
	return res
}

// submission codes end
