// 盛水最多的容器
package array

func maxArea(height []int) int {
	left, right := 0, len(height)-1
	ret := 0

	for left < right {
		var cur int

		if height[left] < height[right] {
			cur = (right - left) * height[left]
			left++
		} else {
			cur = (right - left) * height[right]
			right--
		}

		if cur > ret {
			ret = cur
		}
	}

	return ret
}
