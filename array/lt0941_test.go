// 是否为山型数组
package array

func validMountainArray(arr []int) bool {
	if len(arr) < 3 {
		return false
	}

	if arr[0] >= arr[1] {
		return false
	}

	var maxIdx int
	for i := 1; i < len(arr)-1; i++ {
		if arr[i] <= arr[i-1] {
			break
		}

		maxIdx = i
	}

	for i := maxIdx + 1; i < len(arr); i++ {
		if arr[i] >= arr[i-1] {
			return false
		}
	}

	return true
}
