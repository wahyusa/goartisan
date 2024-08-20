package utils

import (
	"strings"
	"unicode"
)

func ToCamelCase(s string) string {
	s = strings.ToLower(s)
	words := strings.FieldsFunc(s, func(r rune) bool {
		return r == '_' || r == '-'
	})
	for i := range words {
		words[i] = strings.Title(words[i])
	}
	return strings.Join(words, "")
}

func ToSnakeCase(s string) string {
	var result []rune
	for i, r := range s {
		if unicode.IsUpper(r) {
			if i > 0 {
				result = append(result, '_')
			}
			result = append(result, unicode.ToLower(r))
		} else {
			result = append(result, r)
		}
	}
	return string(result)
}

func ToLowerCase(s string) string {
	return strings.ToLower(s)
}
