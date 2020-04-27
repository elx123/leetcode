/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

 func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	result := make([]int,0)

    for ( l1!=nil ){
        result = append(result,l1.Val)
        l1 = l1.Next
    }

    for ( l2 != nil){
        result = append(result,l2.Val)
        l2 = l2.Next
    }

    sort.Ints(result)

    var root *ListNode
    var fake ListNode
    root = &fake

    for _,v := range result{
        temp := &ListNode{
            Val:v,
        }

        root.Next = temp
        root = temp
    }

    return fake.Next
}

/*
func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	if l1 == nil {
		return l2
	}

	if l2 == nil {
		return l1
	}

	if l1.Val < l2.Val {
		l1.Next = mergeTwoLists(l1.Next, l2)
		return l1
	} else {
		l2.Next = mergeTwoLists(l1, l2.Next)
		return l2
	}
}
*/