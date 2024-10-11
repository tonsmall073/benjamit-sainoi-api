package utils

import "strings"

//Example return นายวิศรุต รูปเขีบน (ต้น)
func ConcatFullname(prefixName, firstname, lastname, nickname string) string {
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
