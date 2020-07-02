// 最长公共前缀
package string

import (
	"strings"
	"testing"
)

func TestLT14(t *testing.T) {
	var strs = []string{"", "a"}
	t.Log("ret1: ", longestCommonPrefix(strs))

	strs = []string{"aa", "a"}
	t.Log("ret2: ", longestCommonPrefix(strs))

	var ret strings.Builder

	t.Log("empty builder: ", ret.String())
}

func longestCommonPrefix(strs []string) string {
	if strs == nil || len(strs) == 0 {
		return ""
	}

	l := len(strs)

	var ret strings.Builder
	index := 0

	for true {
		str := strs[0]

		if len(str) <= index {
			return ret.String()
		}

		var c = str[index]

		for i := 1; i < l; i++ {
			str := strs[i]

			if len(str) <= index {
				return ret.String()
			}

			if str[index] != c {
				return ret.String()
			}
		}
		ret.WriteByte(c)
		index++
	}

	return ret.String()
}
