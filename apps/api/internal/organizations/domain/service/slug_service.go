package service

import "fmt"

func GenerateUniqueSlug(baseSlug string, exists func(string) bool) string {

	slug := baseSlug
	counter := 1

	for exists(slug) {
		counter++
		slug = fmt.Sprintf("%s-%d", baseSlug, counter)
	}

	return slug
}
