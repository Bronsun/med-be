package converter

import "gin-boilerplate/pkg/helpers"

func BoolConverter(pseudoBool string) bool {
	if pseudoBool == "Y" || pseudoBool == "y" {
		return true
	}
	if pseudoBool == "N" || pseudoBool == "n" || pseudoBool == "" {
		return false
	}

	return false

}

func VoivodeshipConverter(province string) string {
	for key, value := range helpers.Voievodship {
		if province == value {
			return key
		}
	}
	return ""
}
