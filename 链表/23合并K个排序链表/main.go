func mergeKLists(lists []*ListNode) *ListNode {
	result := make([]int, 0)
	var pointer *ListNode
	for i := 0; i < len(lists); i++ {
		pointer = lists[i]
		for pointer != nil {
			result = append(result, pointer.Val)
			pointer = pointer.Next
		}
	}
	sort.Ints(result)
	return InsertLink(result)
}

func InsertLink(values []int) *ListNode {
	var root ListNode
	if len(values) == 0 {
		return root.Next
	}
	var Itpointer *ListNode
	Itpointer = &root
	for k, _ := range values {
		temp := ListNode{
			Val: values[k],
		}
		Itpointer.Next = &temp
		Itpointer = Itpointer.Next
	}
	return root.Next
}