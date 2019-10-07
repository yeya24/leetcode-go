package tests

import (
    "testing"
)

/**
 * [110] Balanced Binary Tree
 *
 * Given a binary tree, determine if it is height-balanced.
 * 
 * For this problem, a height-balanced binary tree is defined as:
 * 
 * 
 * a binary tree in which the depth of the two subtrees of every node never differ by more than 1.
 * 
 * 
 * Example 1:
 * 
 * Given the following tree [3,9,20,null,null,15,7]:
 * 
 * 
 *     3
 *    / \
 *   9  20
 *     /  \
 *    15   7
 * 
 * Return true.
 * 
 * Example 2:
 * 
 * Given the following tree [1,2,2,3,3,null,null,4,4]:
 * 
 * 
 *        1
 *       / \
 *      2   2
 *     / \
 *    3   3
 *   / \
 *  4   4
 * 
 * 
 * Return false.
 * 
 */

func TestBalancedBinaryTree(t *testing.T) {

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
func isBalanced(root *TreeNode) bool {
    if root != nil {
        left := getHeight(root.Left)
        right := getHeight(root.Right)
        if left > right + 1 || right > left + 1 {
            return false
        }
        return isBalanced(root.Left) && isBalanced(root.Right)
    }
    return true
}

func getHeight(root *TreeNode) int {
    if root != nil {
        left := getHeight(root.Left)
        right := getHeight(root.Right)
        if left > right {
            return left + 1
        } else {
            return right + 1
        }
    }
    return 0
}

// submission codes end
