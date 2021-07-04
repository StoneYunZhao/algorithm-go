// 偶数位数字个数
package array

func findNumbers(nums []int) int {
	res := 0

	for _, n := range nums {
		if digits(n)%2 == 0 {
			res++
		}
	}

	return res
}

func digits(num int) int {
	res := 0

	for num != 0 {
		num = num / 10
		res++
	}

	return res
}
