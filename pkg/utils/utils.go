package utils

import (
	"strings"
)

// CheckIfIncludeTag checks if a repository string includes tag
func CheckIfIncludeTag(repository string) bool {
	return strings.Contains(repository, ":")
}
