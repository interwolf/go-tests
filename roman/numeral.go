package roman

import (
	"strings"
)

// RomanNumeral combines Value and Symbol of a RomanNumeral
type RomanNumeral struct {
	Value  int
	Symbol string
}

// RomanNumerals is a slice for RomanNumeral
type RomanNumerals []RomanNumeral

// ValueOf gets the int value of the symbol (if available)
func (r RomanNumerals) ValueOf(symbols ...byte) int {
	symbol := string(symbols)
	for _, s := range r {
		if s.Symbol == symbol {
			return s.Value
		}
	}
	return 0
}

var importantRoman = RomanNumerals{
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

// ConvertToArabic converts arabic from Roman numeral
func ConvertToArabic(roman string) int {
	total := 0

	for i := 0; i < len(roman); i++ {
		symbol := roman[i]
		if couldSub(i, symbol, roman) {
			nextSymbol := roman[i+1]
			// potential := []byte{symbol, nextSymbol}
			value := importantRoman.ValueOf(symbol, nextSymbol)

			if value != 0 {
				total += value
				i++
			} else {
				total += importantRoman.ValueOf(symbol)
			}
		} else {
			total += importantRoman.ValueOf(symbol)
		}
	}

	return total
}

func couldSub(index int, curr uint8, roman string) bool {
	return index+1 < len(roman) && (curr == 'I' || curr == 'X' || curr == 'C')
}
