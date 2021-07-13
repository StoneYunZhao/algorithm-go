// 一个链表代表一个整数，两个整数相加
package list

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	dummy := &ListNode{}

	cur := dummy
	carry := 0

	for ; l1 != nil || l2 != nil; cur = cur.Next {
		n := carry

		if l1 != nil {
			n += l1.Val
			l1 = l1.Next
		}

		if l2 != nil {
			n += l2.Val
			l2 = l2.Next
		}

		if n > 9 {
			carry, n = 1, n-10
		} else {
			carry = 0
		}

		cur.Next = &ListNode{Val: n}
	}

	if carry > 0 {
		cur.Next = &ListNode{Val: 1}
	}

	return dummy.Next
}
