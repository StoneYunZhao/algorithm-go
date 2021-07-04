// 排序数组平方后在排序
package array

func sortedSquares(nums []int) []int {
	res := make([]int, len(nums))

	i, j, cur := 0, len(nums)-1, len(nums)-1

	for cur >= 0 {
		if i2, j2 := nums[i]*nums[i], nums[j]*nums[j]; i2 < j2 {
			res[cur] = j2
			j--
		} else {
			res[cur] = i2
			i++
		}
		cur--
	}

	return res
}
