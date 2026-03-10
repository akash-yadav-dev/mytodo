package utils

import (
	"regexp"
	"strings"
)

var nonAlphanumeric = regexp.MustCompile(`[^a-z0-9]+`)

func GenerateSlug(name string) string {

	slug := strings.ToLower(name)

	slug = nonAlphanumeric.ReplaceAllString(slug, "-")

	slug = strings.Trim(slug, "-")

	return slug
}
