// 全排列
package backtracking

import "testing"

var ret [][]int

func TestLT46(t *testing.T) {
	permute([]int{1, 2, 3})
	t.Log(ret)
}

func permute(nums []int) [][]int {
	if len(nums) == 0 {
		return nil
	}

	count := 1
	for l := len(nums); l > 1; l-- {
		count *= l
	}

	ret = make([][]int, 0, count)
	help(nums, 0)
	return ret
}

func help(nums []int, index int) {
	l := len(nums)

	if index == l-1 {
		tmp := make([]int, l)
		copy(tmp, nums)
		ret = append(ret, tmp)
		return
	}

	for i := index; i < l; i++ {
		swap(nums, i, index)
		help(nums, index+1)
		swap(nums, i, index)
	}
}

func swap(nums []int, i, j int) {
	if i == j {
		return
	}

	tmp := nums[i]
	nums[i] = nums[j]
	nums[j] = tmp
}
