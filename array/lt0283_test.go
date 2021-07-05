// 把数组中 0 元素都移至右边。保持非 0 元素的相对顺序
package array

func moveZeroes(nums []int) {
	bound := 0
	for ; bound < len(nums) && nums[bound] != 0; bound++ {
	}

	for j := bound + 1; j < len(nums); j++ {
		if nums[j] == 0 {
			continue
		}

		nums[bound], nums[j] = nums[j], nums[bound]
		bound++
	}

}

func moveZeroes2(nums []int) {
	if len(nums) == 0 {
		return
	}
	j := 0
	for i := 0; i < len(nums); i++ {
		if nums[i] != 0 {
			if i != j {
				nums[i], nums[j] = nums[j], nums[i]
			}
			j++
		}
	}
}
