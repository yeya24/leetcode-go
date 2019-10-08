package tests

import (
    "reflect"
    "testing"
)

/**
 * [1079] Sum of Root To Leaf Binary Numbers
 *
 * Given a binary tree, each node has value 0 or 1.  Each root-to-leaf path represents a binary number starting with the most significant bit.  For example, if the path is 0 -> 1 -> 1 -> 0 -> 1, then this could represent 01101 in binary, which is 13.
 * 
 * For all leaves in the tree, consider the numbers represented by the path from the root to that leaf.
 * 
 * Return the sum of these numbers.
 * 
 *  
 * 
 * Example 1:
 * 
 * 
 * 
 * 
 * Input: [1,0,1,0,1,0,1]
 * Output: 22
 * Explanation: (100) + (101) + (110) + (111) = 4 + 5 + 6 + 7 = 22
 * 
 * 
 *  
 * 
 * Note:
 * 
 * 
 * 	The number of nodes in the tree is between 1 and 1000.
 * 	node.val is 0 or 1.
 * 	The answer will not exceed 2^31 - 1.
 * 
 * 
 */

func TestSumofRootToLeafBinaryNumbers(t *testing.T) {
    x := &TreeNode{Val: 1, Left: &TreeNode{Val: 0,Left: &TreeNode{Val: 0}, Right: &TreeNode{Val: 1}}, Right: &TreeNode{Val: 1,Left: &TreeNode{Val: 0}, Right: &TreeNode{Val: 1}}}
    var cases = []struct {
        input  *TreeNode
        sum    int
        output int
    }{
        {
            input:  x,
            output: 22,
        },
    }
    for _, c := range cases {
        x := sumRootToLeaf(c.input)
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
func sumRootToLeaf(root *TreeNode) int {
    var sum int
    dfsSumRootToLeaf(root, 0, &sum)
    return sum
}

func dfsSumRootToLeaf(root *TreeNode, cur int, sum *int) {
    if root == nil {
        return
    }
    cur = cur * 2 + root.Val
    if root.Left == nil && root.Right == nil {
        *sum += cur
    }
    dfsSumRootToLeaf(root.Left, cur, sum)
    dfsSumRootToLeaf(root.Right, cur, sum)
}

// submission codes end
