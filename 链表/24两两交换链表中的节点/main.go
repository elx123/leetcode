func swapPairs(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	var firstNode *ListNode
	var secondNode *ListNode
	firstNode = head
	secondNode = head.Next
	firstNode.Next = swapPairs(secondNode.Next)
	secondNode.Next = firstNode

	return secondNode
}