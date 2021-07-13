// 判断链表是否为回文
package list

func isPalindrome(head *ListNode) bool {
	arr := make([]int, 0)
	for cur := head; cur != nil; cur = cur.Next {
		arr = append(arr, cur.Val)
	}

	for i, j := 0, len(arr)-1; i < j; i, j = i+1, j-1 {
		if arr[i] != arr[j] {
			return false
		}
	}

	return true
}

func isPalindrome2(head *ListNode) bool {
	if head == nil {
		return true
	}

	slow, fast := head, head.Next

	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
	}

	second := slow.Next
	slow.Next = nil

	second = reverse(second)

	l1, l2 := head, second

	for ; l1 != nil && l2 != nil; l1, l2 = l1.Next, l2.Next {
		if l1.Val != l2.Val {
			return false
		}
	}

	return true
}

func reverse(head *ListNode) *ListNode {
	dummy := &ListNode{}

	for cur := head; cur != nil; {
		next := cur.Next

		dummy.Next, cur.Next = cur, dummy.Next

		cur = next
	}

	return dummy.Next
}


