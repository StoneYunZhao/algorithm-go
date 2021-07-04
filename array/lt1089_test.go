// 复制数组中的 0
package array

func duplicateZeros(arr []int) {
	for i := 0; i < len(arr); i++ {
		if arr[i] == 0 && i+1 < len(arr) {
			arr = append(arr[:i+1], arr[i:len(arr)-1]...)
			i++
		}
	}
}

func duplicateZeros2(arr []int) {
	curLen, last := 0, -1
	odd := false

	for i := range arr {
		curLen++
		if arr[i] == 0 {
			curLen++
		}

		if curLen == len(arr) {
			last = i
			break
		}

		if curLen > len(arr) {
			odd = true
			last = i
			break
		}
	}

	j, i := last, len(arr)-1
	if odd {
		arr[i] = arr[j]
		i--
		j--
	}

	for ; j >= 0; j-- {
		if arr[j] == 0 {
			arr[i] = 0
			arr[i-1] = 0
			i -= 2
		} else {
			arr[i] = arr[j]
			i--
		}
	}
}
