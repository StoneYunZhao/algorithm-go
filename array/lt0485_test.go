// 最长连续 1
package array

func findMaxConsecutiveOnes(nums []int) int {
	max, cur := 0, 0

	for _, n := range nums {
		if n == 0 {
			if cur > max {
				max = cur
			}
			cur = 0
		} else if n == 1 {
			cur++
		}
	}

	if cur > max {
		max = cur
	}

	return max
}
