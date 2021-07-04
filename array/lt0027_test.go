// 移除数组中给的值的元素
package array

func removeElement(nums []int, val int) int {
	if len(nums) == 0 {
		return 0
	}

	var bound int

	for i, v := range nums {
		if v == val {
			continue
		}

		swap(nums, i, bound)
		bound++
	}

	return bound
}

func swap(nums []int, i, j int) {
	tmp := nums[i]
	nums[i] = nums[j]
	nums[j] = tmp
}

func removeElement2(nums []int, val int) int {
	i := 0
	for j := 0; j < len(nums); j++ {
		if nums[j] != val {
			nums[i] = nums[j]
			i++
		}
	}

	return i
}
