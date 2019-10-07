package tests

import (
    "testing"
)

/**
 * [904] Leaf-Similar Trees
 *
 * Consider all the leaves of a binary tree.  From left to right order, the values of those leaves form a leaf value sequence.
 * 
 * 
 * 
 * For example, in the given tree above, the leaf value sequence is (6, 7, 4, 9, 8).
 * 
 * Two binary trees are considered leaf-similar if their leaf value sequence is the same.
 * 
 * Return true if and only if the two given trees with head nodes root1 and root2 are leaf-similar.
 * 
 *  
 * 
 * Note:
 * 
 * 
 * 	Both of the given trees will have between 1 and 100 nodes.
 * 
 * 
 */

func TestLeafSimilarTrees(t *testing.T) {

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
func leafSimilar(root1 *TreeNode, root2 *TreeNode) bool {
    var arr1, arr2 []int
    arr1 = dfsSimilarLeaf(root1, arr1)
    arr2 = dfsSimilarLeaf(root2, arr2)
    if len(arr1) != len(arr2) {
        return false
    }
    for i := 0; i < len(arr1); i++ {
        if arr1[i] != arr2[i] {
            return false
        }
    }
    return true
}

func dfsSimilarLeaf(root *TreeNode, list []int) []int {
    if root != nil {
        if root.Left == nil && root.Right == nil {
            list = append(list, root.Val)
        }
        if root.Left != nil {
            list = dfsSimilarLeaf(root.Left, list)
        }
        if root.Right != nil {
            list = dfsSimilarLeaf(root.Right, list)
        }
    }
    return list
}
// submission codes end
