/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
 func reverseList(head *ListNode) *ListNode {
	var current *ListNode
	var preview *ListNode
	var next *ListNode
	current = head
	preview = nil
	for current != nil {
		next = current.Next
		current.Next = preview
		preview = current
		current = next
	}
	return preview
}