package tests

import (
    "reflect"
    "testing"
)

/**
 * [107] Binary Tree Level Order Traversal II
 *
 * Given a binary tree, return the bottom-up level order traversal of its nodes' values. (ie, from left to right, level by level from leaf to root).
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
 * return its bottom-up level order traversal as:
 * 
 * [
 *   [15,7],
 *   [9,20],
 *   [3]
 * ]
 * 
 * 
 */

func TestBinaryTreeLevelOrderTraversalII(t *testing.T) {
    x := &TreeNode{Val: 3, Left: &TreeNode{Val: 9},
        Right: &TreeNode{Val: 20, Left: &TreeNode{Val: 15}, Right: &TreeNode{Val: 7}}}
    var cases = []struct {
        input  *TreeNode
        output [][]int
    }{
        {
            input:  x,
            output: [][]int{{15,7},{9,20},{3}},
        },
    }
    for _, c := range cases {
        x := levelOrderBottom(c.input)
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
func levelOrderBottom(root *TreeNode) [][]int {
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
        res = append([][]int{tmp}, res...)
        cur = next
    }
    return res
}

// submission codes end
