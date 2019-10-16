package setSchedule

import (
	"strconv"
)

func ValidateMethod(method string) bool  {
	return method == "POST"
}

func ValidateId(id string) bool {
	if _, err := strconv.Atoi(id); err == nil {
		return true
	}
	return false
}