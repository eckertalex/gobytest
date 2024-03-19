package main

import (
	"strings"
)

type RomanNumeral struct {
	Symbol string
	Value  uint16
}

var allRomanNumerals = []RomanNumeral{
	{"M", 1000},
	{"CM", 900},
	{"D", 500},
	{"CD", 400},
	{"C", 100},
	{"XC", 90},
	{"L", 50},
	{"XL", 40},
	{"X", 10},
	{"IX", 9},
	{"V", 5},
	{"IV", 4},
	{"I", 1},
}

func ConvertToRoman(arabic uint16) string {
	var roman strings.Builder

	for _, numeral := range allRomanNumerals {
		for arabic >= numeral.Value {
			roman.WriteString(numeral.Symbol)
			arabic -= numeral.Value
		}
	}

	return roman.String()
}

func ConvertToArabic(roman string) uint16 {
	var arabic uint16 = 0

	for _, numeral := range allRomanNumerals {
		for strings.HasPrefix(roman, numeral.Symbol) {
			arabic += numeral.Value
			roman = strings.TrimPrefix(roman, numeral.Symbol)
		}
	}

	return arabic
}
