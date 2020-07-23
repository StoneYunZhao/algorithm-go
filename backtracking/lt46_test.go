// 全排列
package backtracking

import "testing"

var ret [][]int

func TestLT46(t *testing.T) {
	t.Log(permute([]int{1, 2, 3}))
}

func permute(nums []int) [][]int {
	if len(nums) == 0 {
		return nil
	}

	count := 1
	for l := len(nums); l > 1; l-- {
		count *= l
	}

	ret := make([][]int, 0, count)
	return help(nums, 0, ret)
}

func help(nums []int, index int, ret [][]int) [][]int {
	l := len(nums)

	if index == l-1 {
		tmp := make([]int, l)
		copy(tmp, nums)
		return append(ret, tmp)
	}

	for i := index; i < l; i++ {
		swap(nums, i, index)
		ret = help(nums, index+1, ret)
		swap(nums, i, index)
	}
	return ret
}

func swap(nums []int, i, j int) {
	if i == j {
		return
	}

	tmp := nums[i]
	nums[i] = nums[j]
	nums[j] = tmp
}
