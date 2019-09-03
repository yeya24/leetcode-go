package tests

import (
	"testing"
)

/**
 * [835] Linked List Components
 *
 * We are given head, the head node of a linked list containing unique integer values.
 *
 * We are also given the list G, a subset of the values in the linked list.
 *
 * Return the number of connected components in G, where two values are connected if they appear consecutively in the linked list.
 *
 * Example 1:
 *
 *
 * Input:
 * head: 0->1->2->3
 * G = [0, 1, 3]
 * Output: 2
 * Explanation:
 * 0 and 1 are connected, so [0, 1] and [3] are the two connected components.
 *
 *
 * Example 2:
 *
 *
 * Input:
 * head: 0->1->2->3->4
 * G = [0, 3, 1, 4]
 * Output: 2
 * Explanation:
 * 0 and 1 are connected, 3 and 4 are connected, so [0, 1] and [3, 4] are the two connected components.
 *
 *
 * Note:
 *
 *
 * 	If N is the length of the linked list given by head, 1 <= N <= 10000.
 * 	The value of each node in the linked list will be in the range [0, N - 1].
 * 	1 <= G.length <= 10000.
 * 	G is a subset of all values in the linked list.
 *
 *
 */

func TestLinkedListComponents(t *testing.T) {
	node1 := &ListNode{Val: 0, Next: &ListNode{Val: 1, Next: &ListNode{Val: 2, Next: &ListNode{Val: 3}}}}
	node2 := &ListNode{Val: 0, Next: &ListNode{Val: 1, Next: &ListNode{Val: 2, Next: &ListNode{Val: 3, Next: &ListNode{Val: 4}}}}}
	var cases = []struct {
		input  *ListNode
		g      []int
		output int
	}{
		{
			input:  node1,
			g:      []int{0, 1, 3},
			output: 2,
		},
		{
			input:  node2,
			g:      []int{0, 3, 1, 4},
			output: 2,
		},
	}
	for _, c := range cases {
		x := numComponents(c.input, c.g)
		if x != c.output {
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
func numComponents(head *ListNode, G []int) int {
	m := make(map[int]struct{})
	for _, g := range G {
		m[g] = struct{}{}
	}
	res := 0
	var connected bool
	for head != nil {
		if _, ok := m[head.Val]; ok {
			if !connected {
				res++
			}
			connected = true
		} else {
			connected = false
		}
		head = head.Next
	}
	return res
}

// submission codes end
