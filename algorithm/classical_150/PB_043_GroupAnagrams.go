package classical_150

import (
	"slices"
)

// groupAnagrams 字母异位词分组
func groupAnagrams(strs []string) [][]string {
	mp, res := make(map[string][]string), make([][]string, 0)
	for i := 0; i < len(strs); i++ {
		cur := []byte(strs[i])
		slices.Sort(cur)
		str := string(cur)
		mp[str] = append(mp[str], strs[i])
	}
	for _, v := range mp {
		res = append(res, v)
	}
	return res
}
