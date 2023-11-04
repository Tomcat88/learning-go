package roman

import (
	"fmt"
	"testing"
	"testing/quick"
)

type RomanTestingCase struct {
	Arabic uint16
	Roman  string
}

var cases = []RomanTestingCase{
	{1, "I"},
	{2, "II"},
	{3, "III"},
	{4, "IV"},
	{5, "V"},
	{6, "VI"},
	{7, "VII"},
	{8, "VIII"},
	{9, "IX"},
	{10, "X"},
	{14, "XIV"},
	{18, "XVIII"},
	{20, "XX"},
	{39, "XXXIX"},
	{40, "XL"},
	{47, "XLVII"},
	{49, "XLIX"},
	{50, "L"},
	{1984, "MCMLXXXIV"},
	{1988, "MCMLXXXVIII"},
	{2023, "MMXXIII"},
	{1789, "MDCCLXXXIX"},
	{1661, "MDCLXI"},
}

func TestArabic2Roman(t *testing.T) {
	for _, test := range cases {
		t.Run(fmt.Sprintf("%d should be %s", test.Arabic, test.Roman), func(t *testing.T) {
			got := ConvertToRoman(test.Arabic)

			if got != test.Roman {
				t.Errorf("got %s want %s", got, test.Roman)
			}
		})
	}
}

func TestRoman2Arabic(t *testing.T) {
	for _, test := range cases {
		t.Run(fmt.Sprintf("%s should be %d", test.Roman, test.Arabic), func(t *testing.T) {
			got := ConvertToArabic(test.Roman)

			if got != test.Arabic {
				t.Errorf("got %d want %d", got, test.Arabic)
			}
		})
	}
}

func TestWithQuick(t *testing.T) {
	assertion := func (arabic uint16) bool {
		if arabic > 3999 {
			return true
		}
		t.Log(arabic)
		roman := ConvertToRoman(arabic)
		fromRoman := ConvertToArabic(roman)
		return fromRoman == arabic
	}

	if err := quick.Check(assertion, &quick.Config{
		MaxCount: 1000,
	}); err != nil {
		t.Error("failed checks", nil)
	}
}
