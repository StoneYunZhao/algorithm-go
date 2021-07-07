// 实现链表
package list

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func Test0707(t *testing.T) {
	list := Constructor()
	list.AddAtHead(1)
	list.AddAtTail(3)
	list.AddAtIndex(1, 2)

	got := list.Get(1)
	require.Equal(t, got, 2)

	list.DeleteAtIndex(0)

	got = list.Get(0)
	require.Equal(t, got, 2)
}

type MyLinkedList struct {
	val  int
	next *MyLinkedList
}

/** Initialize your data structure here. */
func Constructor() MyLinkedList {
	return MyLinkedList{}
}

/** Get the value of the index-th node in the linked list. If the index is invalid, return -1. */
func (this *MyLinkedList) Get(index int) int {
	if index < 0 {
		return -1
	}

	cur := this

	for i := 0; i <= index; i++ {
		cur = cur.next

		if cur == nil {
			return -1
		}
	}

	return cur.val
}

/** Add a node of value val before the first element of the linked list. After the insertion, the new node will be the first node of the linked list. */
func (this *MyLinkedList) AddAtHead(val int) {
	node := MyLinkedList{val: val}

	this.next, node.next = &node, this.next
}

/** Append a node of value val to the last element of the linked list. */
func (this *MyLinkedList) AddAtTail(val int) {
	cur := this
	for ; cur.next != nil; cur = cur.next {
	}

	cur.next = &MyLinkedList{val: val}
}

/** Add a node of value val before the index-th node in the linked list. If index equals to the length of linked list, the node will be appended to the end of linked list. If index is greater than the length, the node will not be inserted. */
func (this *MyLinkedList) AddAtIndex(index int, val int) {
	if index < 0 {
		return
	}

	cur := this

	for i := 0; i < index; i++ {
		cur = cur.next

		if cur == nil {
			return
		}
	}

	node := &MyLinkedList{val: val}
	cur.next, node.next = node, cur.next
}

/** Delete the index-th node in the linked list, if the index is valid. */
func (this *MyLinkedList) DeleteAtIndex(index int) {
	if index < 0 {
		return
	}

	cur := this

	for i := 0; i < index; i++ {
		cur = cur.next

		if cur == nil {
			return
		}
	}

	if cur.next != nil {
		cur.next = cur.next.next
	}
}
