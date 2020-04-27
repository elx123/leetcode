/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
 //这道题重点是理解哑节点，为了解决长度为1的情况
 func removeNthFromEnd(head *ListNode, n int) *ListNode {
	var fakePoint *ListNode
	fakePoint = head
	var len int
	for fakePoint != nil {
		len++
		fakePoint = fakePoint.Next
	}
	var fake ListNode
	fake.Next = head
	fakePoint = &fake
	key := len - n
	for key > 0 {
		key--
		fakePoint = fakePoint.Next
	}
	fakePoint.Next = fakePoint.Next.Next
	return fake.Next
}