package utils

import (
	"strings"
	"unicode"

	"golang.org/x/text/cases"
	"golang.org/x/text/language"
)

func ToCamelCase(s string) string {
	s = strings.ToLower(s)
	words := strings.FieldsFunc(s, func(r rune) bool {
		return r == '_' || r == '-'
	})
	caser := cases.Title(language.Und)
	for i := range words {
		words[i] = caser.String(words[i])
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
