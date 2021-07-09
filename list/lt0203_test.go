// 移除链表中值等于某个数值的元素
package list

func removeElements(head *ListNode, val int) *ListNode {
	dummy := &ListNode{Next: head}

	for prev, cur := dummy, head; cur != nil; cur = cur.Next{
		if cur.Val == val {
			prev.Next = cur.Next
		} else {
			prev = cur
		}
	}

	return dummy.Next
}
