package tests

import (
    "testing"
)

/**
 * [206] Reverse Linked List
 *
 * Reverse a singly linked list.
 * 
 * Example:
 * 
 * 
 * Input: 1->2->3->4->5->NULL
 * Output: 5->4->3->2->1->NULL
 * 
 * 
 * Follow up:
 * 
 * A linked list can be reversed either iteratively or recursively. Could you implement both?
 * 
 */

func TestReverseLinkedList(t *testing.T) {

}

// submission codes start here

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func reverseList(head *ListNode) *ListNode {
    var prev *ListNode
    var next *ListNode
    node := head

    for node != nil {
        next = node.Next
        node.Next = prev
        prev = node
        node = next
    }
    return prev
}

// submission codes end
