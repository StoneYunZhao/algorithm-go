// 找出链表中环的起点位置
package list

func detectCycle(head *ListNode) *ListNode {
	cache := make(map[*ListNode]struct{})

	for cur := head; cur != nil; cur = cur.Next {
		if _, ok := cache[cur]; ok {
			return cur
		}

		cache[cur] = struct{}{}
	}

	return nil
}

func detectCycle2(head *ListNode) *ListNode {
	for slow, fast := head, head; fast != nil && fast.Next != nil; {
		slow = slow.Next
		fast = fast.Next.Next

		if slow == fast {
			p1, p2 := slow, head
			for p1 != p2 {
				p1 = p1.Next
				p2 = p2.Next
			}
			return p1
		}
	}

	return nil
}