package util

import "strings"

func FetchNodeId(prefix string) string {
	split := strings.Split(prefix, "/")
	if len(split) < 4 {
		return ""
	}
	return split[4]
}

func FetchKni(list []string) string {
	if len(list) < 5 {
		return ""
	}
	return list[5]
}

func FetchFStack(list []string) string {
	if len(list) < 5 {
		return ""
	}
	return list[5]
}
