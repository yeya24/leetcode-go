package tests

import (
    "reflect"
    "testing"
)

/**
 * [113] Path Sum II
 *
 * Given a binary tree and a sum, find all root-to-leaf paths where each path's sum equals the given sum.
 * 
 * Note: A leaf is a node with no children.
 * 
 * Example:
 * 
 * Given the below binary tree and sum = 22,
 * 
 * 
 *       5
 *      / \
 *     4   8
 *    /   / \
 *   11  13  4
 *  /  \    / \
 * 7    2  5   1
 * 
 * 
 * Return:
 * 
 * 
 * [
 *    [5,4,11,2],
 *    [5,8,4,5]
 * ]
 * 
 * 
 */

func TestPathSumII(t *testing.T) {
    x := &TreeNode{Val: 5, Left:&TreeNode{Val: 4, Left: &TreeNode{Val: 11, Left: &TreeNode{Val:7}, Right:&TreeNode{Val:2}}},
        Right: &TreeNode{Val: 8, Left: &TreeNode{Val: 13}, Right: &TreeNode{Val: 4, Left: &TreeNode{Val:5}, Right:&TreeNode{Val:1}}}}
    var cases = []struct {
        input  *TreeNode
        sum int
        output [][]int
    }{
        {
            input: x,
            sum: 22,
            output: [][]int{{5,4,11,2},{5,8,4,5}},
        },
    }
    for _, c := range cases {
        x := pathSum(c.input, c.sum)
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
func pathSum(root *TreeNode, sum int) [][]int {
    var res [][]int
    var path []int
    dfsPathSumII(root, 0, sum, path, &res)
    return res
}

func dfsPathSumII(root *TreeNode, cur, sum int, path []int, res *[][]int) {
    if root == nil {
        return
    }
    cur += root.Val
    path = append(path, root.Val)
    if root.Left == nil && root.Right == nil {
        if cur == sum {
            *res = append(*res, append([]int{}, path...))
        }
        return
    }
    dfsPathSumII(root.Left, cur, sum, path, res)
    dfsPathSumII(root.Right, cur, sum, path, res)
}

// submission codes end
