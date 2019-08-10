package tests

import (
	"fmt"
	"testing"
)

/**
 * [142] Linked List Cycle II
 *
 * Given a linked list, return the node where the cycle begins. If there is no cycle, return null.
 * To represent a cycle in the given linked list, we use an integer pos which represents the position (0-indexed) in the linked list where tail connects to. If pos is -1, then there is no cycle in the linked list.
 * Note: Do not modify the linked list.
 *
 * Example 1:
 *
 * Input: head = [3,2,0,-4], pos = 1
 * Output: tail connects to node index 1
 * Explanation: There is a cycle in the linked list, where tail connects to the second node.
 *
 *
 * Example 2:
 *
 * Input: head = [1,2], pos = 0
 * Output: tail connects to node index 0
 * Explanation: There is a cycle in the linked list, where tail connects to the first node.
 *
 *
 * Example 3:
 *
 * Input: head = [1], pos = -1
 * Output: no cycle
 * Explanation: There is no cycle in the linked list.
 *
 *
 *
 * Follow-up:
 * Can you solve it without using extra space?
 *
 */

func TestLinkedListCycleII(t *testing.T) {
	node1 := &ListNode{Val: 1}
	node2 := &ListNode{Val: 2, Next: node1}
	node1.Next = node2
	a := detectCycle(node1)
	fmt.Println(a)
}

// submission codes start here

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func detectCycle(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return nil
	}
	//assume at least 3 nodes
	dumyHead := &ListNode{
		Next:head,
	}
	fast, slow := dumyHead, dumyHead
	for {
		if fast == nil {
			return nil
		}
		if slow == fast && fast != dumyHead {
			slow = dumyHead
			break
		}
		slow = slow.Next
		fast = fast.Next
		if fast != nil {
			fast = fast.Next
		}
	}

	for slow != fast {
		slow = slow.Next
		fast = fast.Next
	}
	return slow
}

// submission codes end
