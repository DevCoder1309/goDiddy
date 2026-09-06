package utils

import "strings"

func EncodeName(s string) []byte {
	var result []byte

	for _, label := range strings.Split(s , ".") {
		result = append(result, byte(len(label)))
		result = append(result, []byte(label)...)
	}

	result = append(result, 0)

	return result
}