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
