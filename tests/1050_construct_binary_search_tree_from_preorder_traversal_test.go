package tests

import (
    "testing"
)

/**
 * [1050] Construct Binary Search Tree from Preorder Traversal
 *
 * Return the root node of a binary search tree that matches the given preorder traversal.
 * 
 * (Recall that a binary search tree is a binary tree where for every  node.val.  Also recall that a preorder traversal displays the value of the node first, then traverses node.left, then traverses node.right.)
 * 
 *  
 * 
 * Example 1:
 * 
 * 
 * Input: [8,5,1,7,10,12]
 * Output: [8,5,10,1,7,null,12]
 * 
 * 
 * 
 *  
 * 
 * Note: 
 * 
 * 
 * 	1 <= preorder.length <= 100
 * 	The values of preorder are distinct.
 * 
 * 
 */

func TestConstructBinarySearchTreefromPreorderTraversal(t *testing.T) {

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
func bstFromPreorder(preorder []int) *TreeNode {
	var root *TreeNode
	for _, v := range preorder {
		root = insertIntoBST(root, v)
	}
	return root
}


// submission codes end
