package converter

import (
	"strings"
)

var (
	Voievodship = map[string]string{
		"Dolnyśląsk":          "01",
		"Kujawsko-Pomorskie":  "02",
		"Lubelske":            "03",
		"Lubuskie":            "04",
		"Łódzkie":             "05",
		"Małopolskie":         "06",
		"Mazowieckie":         "07",
		"Opolskie":            "08",
		"Podkarpackie":        "09",
		"Podlaskie":           "10",
		"Pomorskie":           "11",
		"Śląskie":             "12",
		"Świętokrzyskie":      "13",
		"Warmińsko-Mazurskie": "14",
		"Wielkopolskie":       "15",
		"Zachodniopomorskie":  "16",
	}
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
	for key, value := range Voievodship {
		if province == value {
			return key
		}
	}
	return ""
}

func CapitalLettersConverter(s string) string {
	return strings.Title(strings.ToLower(s))
}
