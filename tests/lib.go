package tests

//Definition for a linkedlist node.
type ListNode struct {
	Val  int
	Next *ListNode
}

//Definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}

func abs(a, b int) int {
	if a > b {
		return a - b
	}
	return b - a
}

func gcd(a, b int) int {
	if a == 0 {
		return b
	}
	return gcd(b%a, a)
}
