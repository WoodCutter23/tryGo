package getSchedule

import (
	"strconv"
)

func ValidateMethod(method string) bool  {
	return method == "GET"
}

func ValidateData(id string) bool {
	if _, err := strconv.Atoi(id); err == nil {
		return true
	}

	return false
}