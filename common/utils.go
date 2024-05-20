package common

func IsExsit(list []string, keyword string) bool {
	for _, v := range list {
		if v == keyword {
			return true
		}
	}
	return false
}
