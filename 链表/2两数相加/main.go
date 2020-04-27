/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
 //解题思路在微信公众号和收藏中
 func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {

    result := &ListNode{Val:-1}
    var head *ListNode
    head = result
    var sum int
    var carry bool

    for (l1!= nil || l2 != nil){
        sum = 0
        if (l1 != nil){
            sum += l1.Val
            l1 = l1.Next
        }
        if (l2 != nil){
            sum += l2.Val
            l2 = l2.Next
        }
        if carry {
            sum ++
        }
        head.Next = &ListNode{Val:sum%10}
        head=head.Next
        if sum >= 10 {
            carry = true
        }else{
            carry = false
        }
    }
       if carry{
            head.Next= &ListNode{Val:1};
        }
        result = result.Next
    return result;
}