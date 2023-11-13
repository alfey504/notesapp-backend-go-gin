package utils

import (
	"fmt"
	"strings"
	"time"
)

func ParseDateFromString(dateString string) (time.Time, error) {
	return time.Parse(time.RFC3339, dateString)
}

func ParseBoolFromString(str string) (bool, error) {
	if str == "1" || strings.ToLower(str) == "true" {
		return true, nil
	} else if str == "0" || strings.ToLower(str) == "false" {
		return false, nil
	} else {
		return false, fmt.Errorf("Invalid value for favorite")
	}
}
