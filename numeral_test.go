package main

import (
	"fmt"
	"strings"
	"testing"
	"testing/quick"
)

type RomanNumeral struct {
	Value  uint16
	Symbol string
}

var allRomanNumerals = []RomanNumeral{
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

var cases = []struct {
	Arabic uint16
	Roman  string
}{
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
	{Arabic: 100, Roman: "C"},
	{Arabic: 90, Roman: "XC"},
	{Arabic: 400, Roman: "CD"},
	{Arabic: 500, Roman: "D"},
	{Arabic: 900, Roman: "CM"},
	{Arabic: 1000, Roman: "M"},
	{Arabic: 1984, Roman: "MCMLXXXIV"},
	{Arabic: 3999, Roman: "MMMCMXCIX"},
	{Arabic: 2014, Roman: "MMXIV"},
	{Arabic: 1006, Roman: "MVI"},
	{Arabic: 798, Roman: "DCCXCVIII"},
	{Arabic: 1984, Roman: "MCMLXXXIV"},
}

func TestRomanNumerals(t *testing.T) {
	for _, v := range cases {
		t.Run(fmt.Sprintf("%d gets converted to %q", v.Arabic, v.Roman), func(t *testing.T) {
			got := ConvertToRoman(v.Arabic)

			if got != v.Roman {
				t.Errorf("got %q, want %q", got, v.Roman)
			}
		})
	}
}

func TestConvertingToArabic(t *testing.T) {
	for _, v := range cases {
		t.Run(fmt.Sprintf("%q gets converted to %d", v.Roman, v.Arabic), func(t *testing.T) {
			got := ConvertToArabic(v.Roman)

			if got != v.Arabic {
				t.Errorf("got %d, want %d", got, v.Arabic)
			}
		})
	}
}

func TestPropertiesOfConversion(t *testing.T) {
	assertion := func(arabic uint16) bool {
		if arabic > 3999 {
			return true
		}
		t.Log("testing", arabic)
		roman := ConvertToRoman(arabic)
		fromRoman := ConvertToArabic(roman)
		return fromRoman == arabic
	}

	if err := quick.Check(assertion, &quick.Config{MaxCount: 1000}); err != nil {
		t.Error("failed checks", err)
	}
}

func ConvertToRoman(arabic uint16) string {
	var result strings.Builder

	for _, v := range allRomanNumerals {
		for arabic >= v.Value {
			result.WriteString(v.Symbol)
			arabic -= v.Value
		}
	}

	return result.String()
}

func ConvertToArabic(roman string) uint16 {
	arabic := uint16(0)

	for _, v := range allRomanNumerals {
		for strings.HasPrefix(roman, v.Symbol) {
			arabic += v.Value
			roman = strings.TrimPrefix(roman, v.Symbol)
		}
	}

	return arabic
}
