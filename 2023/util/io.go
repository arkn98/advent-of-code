package util

import "strings"

func PrepInput(in string) []string {
	splits := strings.Split(in, "\n")

	return splits[:len(splits)-1]
}
