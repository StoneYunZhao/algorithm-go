package math

import "testing"

func Test9(t *testing.T) {
	ret := isPalindrome(121)
	t.Log(ret)
}

func isPalindrome(x int) bool {
	if x < 0 {
		return false
	}

	if x == 0 {
		return true
	}

	var digits []int

	for x != 0 {
		digits = append(digits, x%10)
		x /= 10
	}

	i, j := 0, len(digits)-1

	for i < j {
		if digits[i] != digits[j] {
			return false
		}
		i++
		j--
	}

	return true
}
