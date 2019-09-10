package tests

import (
    "reflect"
    "testing"
)

/**
 * [144] Binary Tree Preorder Traversal
 *
 * Given a binary tree, return the preorder traversal of its nodes' values.
 * 
 * Example:
 * 
 * 
 * Input: [1,null,2,3]
 *    1
 *     \
 *      2
 *     /
 *    3
 * 
 * Output: [1,2,3]
 * 
 * 
 * Follow up: Recursive solution is trivial, could you do it iteratively?
 * 
 */

func TestBinaryTreePreorderTraversal(t *testing.T) {
    x := &TreeNode{Val: 1, Right: &TreeNode{Val: 2, Left: &TreeNode{Val: 3}}}
    var cases = []struct {
        input  *TreeNode
        output []int
    }{
        {
            input: x,
            output: []int{1, 2, 3},
        },
    }
    for _, c := range cases {
        x := preorderTraversal(c.input)
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
func preorderTraversal(root *TreeNode) []int {
    if root == nil {
        return nil
    }
    res := []int{}
    stack := []*TreeNode{root}
    for len(stack) > 0 {
        s := stack[len(stack)-1]
        stack = stack[:len(stack)-1]
        res = append(res, s.Val)
        if s.Right != nil {
            stack = append(stack, s.Right)
        }
        if s.Left != nil {
            stack = append(stack, s.Left)
        }
    }
    return res
}

// submission codes end
