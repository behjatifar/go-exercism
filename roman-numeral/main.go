package main

import (
	"errors"
	"fmt"
)

//   romanNumStructure := []romanNumerial{
// 	{roman:M,numerical:1000},

//   }

const (
	M  = 1000
	D  = 500
	C  = 100
	L  = 50
	X  = 10
	V  = 5
	I  = 1
	CM = 900
	CD = 400
	XC = 90
	XL = 40
	IX = 9
	IV = 4
)
const stringLimit = 4

// XCIII
//  1400 MCD
// 1240 -> MCCXL
// 1550 -> MDL
// 3555 MMMDLV

// var userInput int

func main() {
	str, error := ConvertDTR(2400) // solution
	fmt.Println(str, error)
	fmt.Println(ToRomanNumeral(2400)) // solution two
}

func ConvertDTR(digit int) (string, error) {
	var romanStr string
	var stringCounter int
	if digit > 3999 || digit <= 0 {
		return "", errors.New("invalid digit Input")
	}
	if digit >= 1000 {
		// CHECKING M
		for digit >= M {
			digit -= M
			stringCounter += 1
			romanStr += "M"
			if stringCounter > stringLimit {
				stringCounter = 0
				break
			}
		}

	}
	if digit >= CM {
		// CHECKING CM
		for digit >= CM {
			digit -= CM
			stringCounter += 1
			romanStr += "CM"
			if stringCounter > stringLimit {
				stringCounter = 0
				break
			}
		}
	}

	if digit >= D {
		// CHECKING D
		for digit >= D {
			digit -= D
			stringCounter += 1
			romanStr += "D"
			if stringCounter > stringLimit {
				stringCounter = 0
				break
			}
		}
	}
	if digit >= CD {
		for digit >= CD {
			digit -= CD
			stringCounter += 1
			romanStr += "CD"
			if stringCounter > stringLimit {
				stringCounter = 0
				break
			}
		}
	}
	if digit >= C {
		// CHECKING C
		for digit >= C {
			digit -= C
			stringCounter += 1
			romanStr += "C"
			if stringCounter > stringLimit {
				stringCounter = 0
				break
			}
		}

	}
	if digit >= XC {
		for digit >= XC {
			digit -= XC
			stringCounter += 1
			romanStr += "XC"
			if stringCounter > stringLimit {
				stringCounter = 0
				break
			}
		}
	}

	if digit >= L {
		// CHECKING L
		for digit >= L {
			digit -= L
			stringCounter += 1
			romanStr += "L"
			if stringCounter > stringLimit {
				stringCounter = 0
				break
			}
		}
	}
	if digit >= XL {
		// checking XL
		for digit >= XL {
			digit -= XL
			stringCounter += 1
			romanStr += "XL"
			if stringCounter > stringLimit {
				stringCounter = 0
				break
			}
		}
	}
	if digit >= X {
		// CHECKING X
		for digit >= X {
			stringCounter += 1
			digit -= X
			romanStr += "X"
			if stringCounter > stringLimit {
				stringCounter = 0
				break
			}
		}
	}
	if digit >= IX {
		for digit >= IX {
			digit -= IX
			stringCounter += 1
			romanStr += "IX"
			if stringCounter > stringLimit {
				stringCounter = 0
				break
			}
		}
	}
	if digit >= V {
		// CHECKING V
		for digit >= V {
			digit -= V
			stringCounter += 1
			romanStr += "V"
			if stringCounter > stringLimit {
				stringCounter = 0
				break
			}
		}
	}
	if digit >= IV {
		for digit >= IV {
			digit -= IV
			stringCounter += 1
			romanStr += "IV"
			if stringCounter > stringLimit {
				stringCounter = 0
				break
			}
		}
	}
	// CHECKING I
	stringCounter = 0
	if digit >= I {
		for digit >= I {
			digit -= I
			stringCounter += 1
			romanStr += "I"
			// fmt.Println(stringCounter)
			if stringCounter > stringLimit {
				// fmt.Println(stringCounter)
				stringCounter = 0
				break
			}
		}

	}
	return romanStr, nil
}
