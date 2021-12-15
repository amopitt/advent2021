package util

import "strconv"

func GetInt(val string) int {
	if i, err := strconv.Atoi(val); err == nil {
		return i
	}
	return -1
}
