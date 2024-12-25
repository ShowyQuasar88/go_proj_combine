package classical_150

import "strings"

// wordPattern 单词规律
func wordPattern(pattern string, s string) bool {
	mp, diMp := make(map[string]byte), make(map[byte]string)
	ss := strings.Split(s, " ")
	if len(ss) != len(pattern) {
		return false
	}
	for i := 0; i < len(ss); i++ {
		sItem, sok := mp[ss[i]]
		pItem, pok := diMp[pattern[i]]
		if !sok && !pok {
			mp[ss[i]] = pattern[i]
			diMp[pattern[i]] = ss[i]
			continue
		} else if sok && pok && sItem == pattern[i] && pItem == ss[i] {
			continue
		} else {
			return false
		}
	}
	return true
}
