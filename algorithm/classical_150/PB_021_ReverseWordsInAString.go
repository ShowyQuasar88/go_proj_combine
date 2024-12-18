package classical_150

// reverseWords 反转字符串中的单词
func reverseWords(s string) string {
	tmp, ans := make([]string, 0), ""
	left, right, length := 0, 0, len(s)
	for left < length && right < length {
		for left < length && s[left] == ' ' {
			left++
		}
		if left >= length {
			break
		}
		right = left + 1
		for right < length && s[right] != ' ' {
			right++
		}
		tmp = append(tmp, s[left:right])
		left = right + 1
	}
	for right = len(tmp) - 1; right >= 0; right-- {
		ans += tmp[right]
		if right != 0 {
			ans += " "
		}
	}
	return ans
}

// reverseWordsDoublePointer 反转字符串中的单词 双指针
func reverseWordsDoublePointer(s string) string {
	left, right, tmp := 0, 0, []byte(s)
	for right < len(tmp) && s[right] == ' ' {
		right++
	}
	for right < len(tmp) {
		if right > 0 && tmp[right] == tmp[right-1] && tmp[right] == ' ' {
			right++
			continue
		}
		tmp[left] = tmp[right]
		right++
		left++
	}
	if tmp[left-1] == ' ' {
		tmp = tmp[:left-1]
	} else {
		tmp = tmp[:left]
	}
	reverse(&tmp, 0, len(tmp)-1)
	for left, right = 0, 1; left < len(tmp) && right < len(tmp); {
		for right < len(tmp) && tmp[right] != ' ' {
			right++
		}
		reverse(&tmp, left, right-1)
		left, right = right+1, right+2
	}
	return string(tmp)
}

func reverse(b *[]byte, left, right int) {
	for left < right {
		tmp := (*b)[left]
		(*b)[left] = (*b)[right]
		(*b)[right] = tmp
		left++
		right--
	}
}
