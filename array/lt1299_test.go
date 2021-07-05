// 数组中元素值改为右边最大元素
package array

func replaceElements(arr []int) []int {
	curMax := -1
	for i := len(arr) - 1; i >= 0; i-- {
		tmp := curMax
		if arr[i] > curMax {
			curMax = arr[i]
		}

		arr[i] = tmp
	}

	return arr
}
