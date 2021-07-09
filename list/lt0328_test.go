// 链表按照奇偶数索引分组
package list

import (
	"fmt"
	"testing"
)

func Test328(t *testing.T) {
	head := &ListNode{
		Val: 1,
		Next: &ListNode{
			Val: 2,
			Next: &ListNode{
				Val: 3,
				Next: &ListNode{
					Val:  4,
					Next: nil,
				},
			},
		},
	}

	head = oddEvenList(head)
	fmt.Printf("head=%#v", head)
}
func oddEvenList(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}

	odd, even, evenHeader := head, head.Next, head.Next

	for cur, isOdd := head.Next.Next, true; cur != nil; cur, isOdd = cur.Next, !isOdd {
		if isOdd {
			odd.Next, odd = cur, cur
		} else {
			even.Next, even = cur, cur
		}
	}

	odd.Next, even.Next = evenHeader, nil

	return head

}
