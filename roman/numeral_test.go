package roman

import (
	"fmt"
	"testing"
)

func TestRoman(t *testing.T) {
	testCases := []struct {
		Description string
		Arabic      int
		Roman       string
	}{
		{"1 - I", 1, "I"},
		{"2 - III", 2, "II"},
		{"3 - III", 3, "III"},
		{"4 - IV", 4, "IV"},
		{"5 - V", 5, "V"},
		{"6 - VI", 6, "VI"},
		{"7 - VII", 7, "VII"},
		{"8 - VIII", 8, "VIII"},
		{"9 - IX", 9, "IX"},
		{"10 - X", 10, "X"},
		{"14 - XIV", 14, "XIV"},
		{"18 - XVIII", 18, "XVIII"},
		{"20 - XX", 20, "XX"},
		{"39 - XXXIX", 39, "XXXIX"},
		{"40 - XL", 40, "XL"},
		{"47 - XLVII", 47, "XLVII"},
		{"49 - XLIX", 49, "XLIX"},
		{"50 - L", 50, "L"},
		{"100 - C", 100, "C"},
		{"90 - XC", 90, "XC"},
		{"400 - CD", 400, "CD"},
		{"500 - D", 500, "D"},
		{"900 - CM", 900, "CM"},
		{"1000 - M", 1000, "M"},
		{"1984 - MCMLXXXIV", 1984, "MCMLXXXIV"},
		{"3999 - MMMCMXCIX", 3999, "MMMCMXCIX"},
		{"2014 - MMXIV", 2014, "MMXIV"},
		{"1006 - MVI", 1006, "MVI"},
		{"798 - DCCXCVIII", 798, "DCCXCVIII"},
	}

	fmt.Println("============= ConvertToRoman ==============")
	for _, test := range testCases {
		t.Run(test.Description, func(t *testing.T) {
			got := ConvertToRoman(test.Arabic)
			want := test.Roman

			if got != want {
				t.Errorf("got: %q, want: %q\n", got, want)
			}
		})
	}

	fmt.Println("============= ConvertFromRoman ==============")
	for _, test := range testCases {
		t.Run(test.Description, func(t *testing.T) {
			got := ConvertToArabic(test.Roman)
			want := test.Arabic

			if got != want {
				t.Errorf("got: %d, want: %d\n", got, want)
			}
		})
	}

}
