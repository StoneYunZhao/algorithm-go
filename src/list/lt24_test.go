// 链表中节点两两交换
package list

func swapPairs(head *ListNode) *ListNode {
	dummy := &ListNode{
		Next: head,
	}

	cur := dummy

	for cur.Next != nil && cur.Next.Next != nil {
		l1, l2 := cur.Next, cur.Next.Next
		l3 := l2.Next

		cur.Next = l2
		l2.Next = l1
		l1.Next = l3

		cur = l1
	}

	return dummy.Next
}
