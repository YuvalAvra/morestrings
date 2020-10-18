package morestrings

func RemoveEnding(s string) string {
	len := len(s)
	if len < 0 {
		return s
	}
	return string(s[len-1])
}
