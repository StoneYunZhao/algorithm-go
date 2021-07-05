// 找出数组中第 3 大元素
package array

import (
	"math"
	"sort"
)

func thirdMax(nums []int) int {
	max, max2, max3 := nums[0], nums[0], nums[0]

	for _, n := range nums {
		if max == max2 {
			if n > max {
				max, max2, max3 = n, max, max2
			} else if n < max {
				max2, max3 = n, n
			}

			continue
		}

		if max2 == max3 {
			if n > max {
				max, max2, max3 = n, max, max2
			} else if n < max && n > max2 {
				max2, max3 = n, max2
			} else if n < max3 {
				max3 = n
			}

			continue
		}

		if n > max {
			max, max2, max3 = n, max, max2
		} else if n < max && n > max2 {
			max2, max3 = n, max2
		} else if n < max2 && n > max3 {
			max3 = n
		}
	}

	if max != max2 && max2 != max3 {
		return max3
	}

	return max

}

func thirdMax2(nums []int) int {
	max, max2, max3 := math.MinInt64, math.MinInt64, math.MinInt64

	for _, n := range nums {
		if n > max {
			max, max2, max3 = n, max, max2
		} else if n < max && n > max2 {
			max2, max3 = n, max2
		} else if n < max2 && n > max3 {
			max3 = n
		}
	}

	if max3 == math.MinInt64 {
		return max
	}

	return max3
}

func thirdMax3(nums []int) int {
	sort.Ints(nums)

	found2 := false
	for i := len(nums) - 1; i > 0; i-- {
		if nums[i] > nums[i - 1] {
			if !found2 {
				found2 = true
				continue
			}

			return nums[i - 1]
		}
	}

	return nums[len(nums)-1]
}