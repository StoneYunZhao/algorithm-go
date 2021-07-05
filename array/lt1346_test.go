// 数组中元素是否有两倍关系
package array

func checkIfExist(arr []int) bool {
	res := false

	cache := make(map[int]int)
	for i, n := range arr {
		cache[n] = i
	}

	for j, n := range arr {
		if i, ok := cache[n*2]; ok && i != j {
			res = true
			break
		}
	}

	return res
}
