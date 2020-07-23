// 两数之和
package math

func twoSum(nums []int, target int) []int {
	if nums == nil {
		return nil
	}

	var cache = make(map[int]int)
	for i, v := range nums {
		if i1, ok := cache[target-v]; ok {
			return []int{i, i1}
		}

		cache[v] = i
	}

	return nil
}
