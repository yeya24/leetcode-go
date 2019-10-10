package tests

import (
    "reflect"
    "testing"
)

/**
 * [404] Sum of Left Leaves
 *
 * Find the sum of all left leaves in a given binary tree.
 * 
 * Example:
 * 
 *     3
 *    / \
 *   9  20
 *     /  \
 *    15   7
 * 
 * There are two left leaves in the binary tree, with values 9 and 15 respectively. Return 24.
 * 
 * 
 */

func TestSumofLeftLeaves(t *testing.T) {
    x := &TreeNode{Val: 3, Left: &TreeNode{Val: 9}, Right: &TreeNode{Val: 20,Left: &TreeNode{Val: 15}, Right: &TreeNode{Val: 7}}}
    var cases = []struct {
        input  *TreeNode
        sum    int
        output int
    }{
        {
            input:  x,
            output: 24,
        },
    }
    for _, c := range cases {
        x := sumOfLeftLeaves(c.input)
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
func sumOfLeftLeaves(root *TreeNode) int {
    if root == nil {
        return 0
    }
    return dfsLeftLeaves(root.Left, true) + dfsLeftLeaves(root.Right, false)
}

func dfsLeftLeaves(root *TreeNode, isLeft bool) int {
    if root == nil {
        return 0
    }
    if root.Left == nil && root.Right == nil && isLeft {
        return root.Val
    }
    return dfsLeftLeaves(root.Left, true) + dfsLeftLeaves(root.Right, false)
}

// submission codes end
