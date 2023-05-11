package stringutil

import (
	"strconv"
)

func ParseInt(s string) int {
	i, err := strconv.ParseInt(s, 10, 64)
	if err != nil {
		return -1
	}
	return int(i)
}

func ParseFloat(s string) float64 {
	f, err := strconv.ParseFloat(s, 64)
	if err != nil {
		return -1.0
	}
	return f
}
