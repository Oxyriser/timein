package main

import (
	"strings"
)

// ToPascalSnakeCase converts an alphanumeric string to Pascal_Snake_Case
func ToPascalSnakeCase(s string) string {
	s = strings.Trim(s, " ")
	n := ""
	capNext := true

	for i, v := range s {
		if isUpperCase(v) {
			n += string(v)
		}

		if isLowerCase(v) {
			if capNext {
				n += strings.ToUpper(string(v))
			} else {
				n += string(v)
			}
		}

		if v == '_' || v == ' ' || v == '-' || v == '.' {
			n += "_"
			capNext = true
		} else {
			capNext = false
			if i+1 < len(s) {
				next := []rune(s)[i+1]
				if isLowerCase(v) && isUpperCase(next) {
					n += "_"
					capNext = true
				}
			}
		}

	}
	return n
}

func isUpperCase(c rune) bool {
	return c >= 'A' && c <= 'Z'
}

func isLowerCase(c rune) bool {
	return c >= 'a' && c <= 'z'
}
