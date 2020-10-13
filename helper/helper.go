package helper

import "strings"

func InArray(need interface{}, arr []interface{}) bool {
	for _, item := range arr {
		if need == item {
			return true
		}
	}
	return false
}

func StringMultiIndex(need string, arr []string) bool {
	for _, item := range arr {
		if strings.Index(need, item) != -1 {
			return true
		}
	}
	return false
}