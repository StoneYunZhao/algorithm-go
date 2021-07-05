// 找出数组中不存在的元素
package array


func findDisappearedNumbers(nums []int) []int {
	m := make([]int, len(nums))

	for _, n := range nums {
		m[n-1]++
	}

	res := make([]int, 0)

	for i, n := range m {
		if n == 0 {
			res = append(res, i+1)
		}
	}

	return res
}

func findDisappearedNumbers2(nums []int) []int {
	for _, n := range nums {
		nums[abs(n)-1] = -abs(nums[abs(n)-1])
	}

	res := make([]int, 0)

	for i, n := range nums {
		if n > 0 {
			res = append(res, i+1)
		}
	}

	return res
}

func abs(i int) int {
	if i < 0 {
		return -i
	}
	return i
}

