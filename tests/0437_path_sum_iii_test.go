package tests

import (
	"testing"
)

/**
 * [437] Path Sum III
 *
 * You are given a binary tree in which each node contains an integer value.
 *
 * Find the number of paths that sum to a given value.
 *
 * The path does not need to start or end at the root or a leaf, but it must go downwards
 * (traveling only from parent nodes to child nodes).
 *
 * The tree has no more than 1,000 nodes and the values are in the range -1,000,000 to 1,000,000.
 *
 * Example:
 *
 * root = [10,5,-3,3,2,null,11,3,-2,null,1], sum = 8
 *
 *       10
 *      /  \
 *     5   -3
 *    / \    \
 *   3   2   11
 *  / \   \
 * 3  -2   1
 *
 * Return 3. The paths that sum to 8 are:
 *
 * 1.  5 -> 3
 * 2.  5 -> 2 -> 1
 * 3. -3 -> 11
 *
 *
 */

func TestPathSumIII(t *testing.T) {
	x := &TreeNode{Val: 10, Left:&TreeNode{Val: 5, Right: &TreeNode{Val: 2, Right: &TreeNode{Val:1}}, Left: &TreeNode{Val: 3, Left: &TreeNode{Val:3}, Right:&TreeNode{Val:-2}}},
		Right: &TreeNode{Val: -3, Right: &TreeNode{Val: 11}}}
	var cases = []struct {
		input  *TreeNode
		sum int
		output int
	}{
		{
			input: x,
			sum: 8,
			output: 3,
		},
	}
	for _, c := range cases {
		x := pathSumIII(c.input, c.sum)
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
func pathSumIII(root *TreeNode, sum int) int {
	path := []int{}
	var c int
	dfsPathSumIII(root, path, &c, sum)
	return c
}

func dfsPathSumIII(root *TreeNode, path []int, count *int, sum int) {
	if root == nil {
		return
	}
	cur := 0
	path = append(path, root.Val)
	i := len(path) - 1
	for i >= 0 {
		cur += path[i]
		if cur == sum {
			*count++
		}
		i--
	}
	dfsPathSumIII(root.Left, path, count, sum)
	dfsPathSumIII(root.Right, path, count, sum)
}

// submission codes end
