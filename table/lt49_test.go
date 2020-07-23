// 字符串按照异位词分组
package table

func groupAnagrams(strs []string) [][]string {
	if strs == nil {
		return nil
	}

	var ret [][]string
	var tbs [][]int

	for _, str := range strs {
		tb1 := getTable(str)

		var found = false
		for i, tb2 := range tbs {
			if tableEqual(tb1, tb2) {
				found = true
				ret[i] = append(ret[i], str)
			}
		}

		if !found {
			ret = append(ret, []string{str})
			tbs = append(tbs, tb1)
		}
	}

	return ret

}

func tableEqual(a, b []int) bool {
	for i := 0; i < 26; i++ {
		if a[i] != b[i] {
			return false
		}
	}

	return true
}

func getTable(str string) []int {
	ret := make([]int, 26)

	for _, c := range str {
		ret[c-'a']++
	}

	return ret
}
