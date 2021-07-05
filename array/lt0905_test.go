// 把数组汇总偶数放前面，奇数放后面
package array

func sortArrayByParity(nums []int) []int {

	for i, j := 0, len(nums)-1; i < j; {
		for ; i < j && nums[i]%2 == 0; i++ { }
		for ; i < j && nums[j]%2 != 0; j-- { }

		nums[i], nums[j] = nums[j], nums[i]
		i++
		j--
	}

	return nums
}
