package utils

import "strings"

func ConcatFullName(prefixName, firstname, lastname, nickname string) string {
	var parts []string

	if prefixName != "" && firstname != "" {
		parts = append(parts, prefixName+firstname)
	} else {
		if firstname != "" {
			parts = append(parts, firstname)
		}
	}

	if lastname != "" {
		parts = append(parts, lastname)
	}

	if nickname != "" {
		parts = append(parts, "("+nickname+")")
	}

	return strings.Join(parts, " ")
}
