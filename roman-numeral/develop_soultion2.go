package main

import (
	"errors"
)

type romanNumeral struct {
	roman   string
	numeral int
}

func ToRomanNumeral(num int) (result string, err error) {
	if num > 3999 || num <= 0 {
		return "", errors.New("invalid input")
	}

	romanNum := []romanNumeral{
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
	for _, value := range romanNum {
		if num >= value.numeral {
			for num >= value.numeral {
				result += value.roman
				num -= value.numeral
			}
		}

	}
	return result, nil
}
