// 反转链表
package list


func reverseList(head *ListNode) *ListNode {
	dummy := &ListNode{}

	for cur := head; cur != nil; {
		dummy.Next, cur.Next, cur = cur, dummy.Next, cur.Next
	}

	return dummy.Next
}
