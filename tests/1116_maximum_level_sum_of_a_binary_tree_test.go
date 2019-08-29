package tests

import (
    "math"
    "testing"
)

/**
 * [1116] Maximum Level Sum of a Binary Tree
 *
 * Given the root of a binary tree, the level of its root is 1, the level of its children is 2, and so on.
 * 
 * Return the smallest level X such that the sum of all the values of nodes at level X is maximal.
 * 
 *  
 * 
 * Example 1:
 * 
 * 
 * 
 * 
 * Input: [1,7,0,7,-8,null,null]
 * Output: 2
 * Explanation: 
 * Level 1 sum = 1.
 * Level 2 sum = 7 + 0 = 7.
 * Level 3 sum = 7 + -8 = -1.
 * So we return the level with the maximum sum which is level 2.
 * 
 * 
 *  
 * 
 * Note:
 * 
 * 
 * 	The number of nodes in the given tree is between 1 and 10^4.
 * 	-10^5 <= node.val <= 10^5
 * 
 * 
 */

func TestMaximumLevelSumofaBinaryTree(t *testing.T) {
    x := &TreeNode{Val: 1, Left: &TreeNode{Val: 7, Left: &TreeNode{Val: 7}, Right: &TreeNode{Val: -8}},
        Right: &TreeNode{Val: 0}}
    var cases = []struct {
        input  *TreeNode
        output int
    }{
        {
            input:  x,
            output: 2,
        },
    }
    for _, c := range cases {
        if c.output != maxLevelSum(c.input) {
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
func maxLevelSum(root *TreeNode) int {
    if root == nil {
        return 0
    }
    index := 0
    curLevel := 1
    maxValue := math.MinInt64
    cur := []*TreeNode{root}
    for len(cur) > 0 {
        levelSum := 0
        next := []*TreeNode{}
        for _, n := range cur {
            levelSum += n.Val
            if n.Left != nil {
                next = append(next, n.Left)
            }
            if n.Right != nil {
                next = append(next, n.Right)
            }
        }
        cur = next
        if levelSum > maxValue {
            maxValue = levelSum
            index = curLevel
        }
        curLevel++
    }
    return index
}

// submission codes end
