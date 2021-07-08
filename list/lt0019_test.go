// 移除链表中倒数第 n 个元素
package list


func removeNthFromEnd(head *ListNode, n int) *ListNode {
	slow, fast := head, head

	for i := 0; i < n; i++ {
		fast = fast.Next
	}

	if fast == nil {
		return slow.Next
	}

	for fast.Next != nil {
		slow = slow.Next
		fast = fast.Next
	}

	slow.Next = slow.Next.Next

	return head
}
