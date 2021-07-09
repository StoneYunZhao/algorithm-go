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


