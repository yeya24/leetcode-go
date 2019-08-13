package tests

import (
	"reflect"
	"testing"
)

/**
 * [1072] Next Greater Node In Linked List
 *
 * We are given a linked list with head as the first node.  Let's number the nodes in the list: node_1, node_2, node_3, ... etc.
 *
 * Each node may have a next larger value: for node_i, next_larger(node_i) is the node_j.val such that j > i, node_j.val > node_i.val, and j is the smallest possible choice.  If such a j does not exist, the next larger value is 0.
 *
 * Return an array of integers answer, where answer[i] = next_larger(node_{i+1}).
 *
 * Note that in the example inputs (not outputs) below, arrays such as [2,1,5] represent the serialization of a linked list with a head node value of 2, second node value of 1, and third node value of 5.
 *
 *
 *
 *
 * Example 1:
 *
 *
 * Input: [2,1,5]
 * Output: [5,5,0]
 *
 *
 *
 * Example 2:
 *
 *
 * Input: [2,7,4,3,5]
 * Output: [7,0,5,5,0]
 *
 *
 *
 * Example 3:
 *
 *
 * Input: [1,7,5,1,9,2,5,1]
 * Output: [7,9,9,9,0,5,0,0]
 *
 *
 *
 *
 * Note:
 *
 *
 * 	1 <= node.val <= 10^9 for each node in the linked list.
 * 	The given list has length in the range [0, 10000].
 *
 *
 *
 *
 */

func TestNextGreaterNodeInLinkedList(t *testing.T) {
	a := &ListNode{Val: 2, Next: &ListNode{Val: 7, Next: &ListNode{Val: 4, Next: &ListNode{Val: 3, Next: &ListNode{Val: 5, Next: nil}}}}}
	b := &ListNode{Val: 2, Next: &ListNode{Val: 1, Next: &ListNode{Val: 5, Next: nil}}}
	var cases = []struct {
		input  *ListNode
		output []int
	}{
		{
			input:  b,
			output: []int{5, 5, 0},
		},
		{
			input:  a,
			output: []int{7, 0, 5, 5, 0},
		},
	}
	for _, c := range cases {
		x := nextLargerNodes(c.input)
		if !reflect.DeepEqual(x, c.output) {
			t.Fail()
		}
	}
}

// submission codes start here

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func nextLargerNodes(head *ListNode) []int {
	if head == nil {
		return nil
	}
	vals := make([]int, 0)
	for r := head; r != nil; r = r.Next {
		vals = append(vals, r.Val)
	}
	stack := []int{}
	res := make([]int, len(vals))
	for i, v := range vals {
		if len(stack) != 0 {
			for len(stack) > 0 {
				l := len(stack)
				if v > vals[stack[l-1]] {
					res[stack[l-1]] = v
					stack = stack[:l-1]
				} else {
					break
				}
			}
		}
		stack = append(stack, i)
		res[i] = 0
	}
	return res
}

// submission codes end
