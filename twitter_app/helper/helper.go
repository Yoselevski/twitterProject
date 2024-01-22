package helper

import (
	"time"
)

func RemoveString(slice []string, s string) []string {
	index := -1
	for i, val := range slice {
		if val == s {
			index = i
			break
		}
	}
	if index == -1 {
		return slice
	}
	return append(slice[:index], slice[index+1:]...)
}

func GetLatestDate(date1 time.Time, date2 time.Time) time.Time {
	if date1.After(date2) {
		return date1
	}
	return date2
}

func Contains(slice []string, s string) bool {
	for _, val := range slice {
		if val == s {
			return true
		}
	}
	return false
}

