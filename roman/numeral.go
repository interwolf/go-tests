package roman

import (
	"strings"
)

// RomanNumeral combines Value and Symbol of a RomanNumeral
type RomanNumeral struct {
	Value  int
	Symbol string
}

var importantRoman = []RomanNumeral{
	{1000, "M"},
	{900, "CM"},
	{500, "D"},
	{400, "CD"},
	{100, "C"},
	{90, "XC"},
	{50, "L"},
	{40, "XL"},
	{10, "X"},
	{9, "IX"},
	{5, "V"},
	{4, "IV"},
	{1, "I"},
}

// ConvertToRoman converts arabic to Roman numeral
func ConvertToRoman(arabic int) string {

	var result strings.Builder

	for _, important := range importantRoman {
		for arabic >= important.Value {
			result.WriteString(important.Symbol)
			arabic -= important.Value
		}
	}

	return result.String()
}

// ConvertFromRoman converts arabic from Roman numeral
func ConvertFromRoman(roman string) int {
	return 0
}
