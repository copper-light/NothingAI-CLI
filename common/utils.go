package common

func IsExsit(list []string, keyword string) bool {
	for _, v := range list {
		if v == keyword {
			return true
		}
	}
	return false
}

func IsHangul(r rune) bool {
	return r >= 0xAC00 && r <= 0xD7A3
}

func CountingHangul(s string) int {
	var cnt int = 0
	for _, c := range s {
		if IsHangul(c) {
			cnt++
		}
	}
	return cnt
}
