package roman

import "strings"

type RomanNumeral struct {
	Value   uint16
	Numeral string
}

var knownCases = []RomanNumeral{
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

func ConvertToRoman(num uint16) string {
	var result strings.Builder

	for _, numeral := range knownCases {
		for num >= numeral.Value {
			result.WriteString(numeral.Numeral)
			num -= numeral.Value
		}
	}

	return result.String()
}

func ConvertToArabic(roman string) (result uint16) {
	for _, numeral := range knownCases {
		for strings.HasPrefix(roman, numeral.Numeral) {
			result += numeral.Value
			roman = strings.Replace(roman, numeral.Numeral, "", 1)
		}
	}
	return
}
