// 数组排序后，元素位置变化的个数
package array

import "sort"

func heightChecker(heights []int) int {
	expected := make([]int, len(heights))
	copy(expected, heights)
	sort.Ints(expected)

	c := 0

	for i := range expected {
		if expected[i] != heights[i] {
			c++
		}
	}

	return c

}