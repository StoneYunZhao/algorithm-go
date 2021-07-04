// 合并两个有序数组
package array

func merge(nums1 []int, m int, nums2 []int, n int) {
	for i, j, cur := m-1, n-1, len(nums1)-1; cur >= 0; cur-- {
		if i < 0 {
			nums1[cur] = nums2[j]
			j--
			continue
		}

		if j < 0 {
			nums1[cur] = nums1[i]
			i--
			continue
		}

		if nums1[i] > nums2[j] {
			nums1[cur] = nums1[i]
			i--
		} else {
			nums1[cur] = nums2[j]
			j--
		}
	}
}
