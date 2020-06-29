// 反转整数
package math

import "math"

func reverse(x int) int {
	if x == 0 {
		return 0
	}

	negative := x < 0
	if negative {
		x = -x
	}

	ret := 0

	for x != 0 {
		digit := x % 10

		if (math.MaxInt32-digit)/10 < ret {
			return 0
		}

		ret = ret*10 + digit
		x /= 10
	}

	if negative {
		return -ret
	}
	return ret
}
