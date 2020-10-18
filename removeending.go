package morestrings

import "github.com/yuvalavra/evenmorestrings"

func RemoveEnding(s string) string {
	len := len(s)
	if len < 0 {
		return s
	}
	return evenmorestrings.DoubleString(string(s[len-1]))
}
