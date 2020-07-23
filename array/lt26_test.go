// 删除有序数组中的重复项
package array

func removeDuplicates(nums []int) int {
	if len(nums) == 0 {
		return 0
	}

	bound := 1

	for i, v := range nums {
		if i == 0 || v == nums[bound-1] {
			continue
		}

		nums[bound] = v
		bound++
	}

	return bound
}
