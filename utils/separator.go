package utils

import (
	"fmt"
	"strings"
)

func ConvertToArrayWithCommaSeparator(String string) []string {
	return strings.Split(String, ", ")
}

func CleanAngleBracketsinString(String string) string {

	String = strings.ReplaceAll(String, "[", "")
	String = strings.ReplaceAll(String, "]", "")
	return String
}

func CleanSingleQuoteinString(String string) string {
	return strings.ReplaceAll(String, "'", "")
}

func ConvertToArrayWithDotSeparator(String string) []string {
	return strings.Split(String, ".. ")
}

func ConvertToArrayWithDoubleLineSeparator(String string) []string {
	return strings.Split(String, "--")
}

func AddNumbering(input []string) []string {
	var result []string
	for i, item := range input {
		result = append(result, fmt.Sprintf("%d. %s", i+1, item))
	}
	return result
}
