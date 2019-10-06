package tests

import (
    "testing"
)

/**
 * [653] Two Sum IV - Input is a BST
 *
 * Given a Binary Search Tree and a target number, return true if there exist two elements in the BST such that their sum is equal to the given target.
 * 
 * Example 1:
 * 
 * 
 * Input: 
 *     5
 *    / \
 *   3   6
 *  / \   \
 * 2   4   7
 * 
 * Target = 9
 * 
 * Output: True
 * 
 * 
 *  
 * 
 * Example 2:
 * 
 * 
 * Input: 
 *     5
 *    / \
 *   3   6
 *  / \   \
 * 2   4   7
 * 
 * Target = 28
 * 
 * Output: False
 * 
 * 
 *  
 * 
 */

func TestTwoSumIVInputisaBST(t *testing.T) {

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
func findTarget(root *TreeNode, k int) bool {
    array := inorderTraversal(root)
    i, j := 0, len(array) - 1
    for i < j {
        if array[i] + array[j] == k {
            return true
        } else if array[i] + array[j] < k {
            i++
        } else {
            j--
        }
    }
    return false
}

// submission codes end
