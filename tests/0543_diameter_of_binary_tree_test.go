package tests

import (
    "reflect"
    "testing"
)

/**
 * [543] Diameter of Binary Tree
 *
 * 
 * Given a binary tree, you need to compute the length of the diameter of the tree. The diameter of a binary tree is the length of the longest path between any two nodes in a tree. This path may or may not pass through the root.
 * 
 * 
 * 
 * Example:
 * Given a binary tree 
 * 
 *           1
 *          / \
 *         2   3
 *        / \     
 *       4   5    
 * 
 * 
 * 
 * Return 3, which is the length of the path [4,2,1,3] or [5,2,1,3].
 * 
 * 
 * Note:
 * The length of path between two nodes is represented by the number of edges between them.
 * 
 */

func TestDiameterofBinaryTree(t *testing.T) {
    x := &TreeNode{Val: 1, Left: &TreeNode{Val: 2,Left: &TreeNode{Val: 4}, Right: &TreeNode{Val: 5}}, Right: &TreeNode{Val: 3}}
    var cases = []struct {
        input  *TreeNode
        sum    int
        output int
    }{
        {
            input:  x,
            output: 3,
        },
    }
    for _, c := range cases {
        x := diameterOfBinaryTree(c.input)
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
func diameterOfBinaryTree(root *TreeNode) int {
    sum := 1
    dfsDiameterOfBinaryTree(root, &sum)
    return sum - 1
}

func dfsDiameterOfBinaryTree(root *TreeNode, sum *int) int {
    if root == nil {
        return 0
    }
    left := dfsDiameterOfBinaryTree(root.Left, sum)
    right := dfsDiameterOfBinaryTree(root.Right, sum)
    if left + right + 1 > *sum {
        *sum = left + right + 1
    }
    if left > right {
        return left + 1
    }
    return right + 1
}

// submission codes end
