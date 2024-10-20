package utils

import "strings"

func IsSwaggerPath(path string) bool {
	swaggerPaths := []string{"/api/docs", "/swagger", "/swagger/doc.json"}
	for _, swaggerPath := range swaggerPaths {
		if path == swaggerPath || strings.HasPrefix(path, swaggerPath) {
			return true
		}
	}
	return false
}
