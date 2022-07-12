package converter

import (
	"gin-boilerplate/pkg/helpers"
	"strings"
)

// BoolConverter converts Y and N to bool statetment
func BoolConverter(pseudoBool string) bool {
	if pseudoBool == "Y" || pseudoBool == "y" {
		return true
	}
	if pseudoBool == "N" || pseudoBool == "n" || pseudoBool == "" {
		return false
	}

	return false
}

// VoivodeshipConverter map voivodeships to
func VoivodeshipConverter(province string) string {
	for key, value := range helpers.Voievodship {
		if province == value {
			return key
		}
	}
	return ""
}

func CapitalLettersConverter(s string) string {
	return strings.Title(strings.ToLower(s))
}
