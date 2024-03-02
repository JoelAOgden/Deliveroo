package cron

import (
	"strconv"
)

func isAsterisk(in string) bool {
	return in == "*"
}

func isNumber(in string) bool {
	if _, err := strconv.Atoi(in); err != nil {
		return false
	}
	return true
}
