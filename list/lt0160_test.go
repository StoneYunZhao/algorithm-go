// 找到两个链表相交的点
package list

func getIntersectionNode(headA, headB *ListNode) *ListNode {
	numA := getNodesCount(headA)
	numB := getNodesCount(headB)

	skip := numB - numA
	short, long := headA, headB
	if numA > numB {
		short, long = long, short
		skip = -skip
	}

	for i := 0; i < skip; i++ {
		long = long.Next
	}

	for ; short != long ; {
		short = short.Next
		long = long.Next
	}

	return short
}

func getNodesCount(head *ListNode) int {
	c := 0
	for cur := head; cur != nil; cur = cur.Next {
		c++
	}
	return c
}
