// Package "en" implements function calls to convert floating
// point numbers to and from Engineering Notation.
package en

/*
  Copyright 2014, Kent Loobey, all rights reserved.

  This program is free software: you can redistribute it and/or modify
  it under the terms of the GNU General Public License as published by
  the Free Software Foundation, either version 3 of the License, or
  (at your option) any later version.

  This program is distributed in the hope that it will be useful,
  but WITHOUT ANY WARRANTY; without even the implied warranty of
  MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
  GNU General Public License for more details.

  You should have received a copy of the GNU General Public License
  along with this program.  If not, see <http://www.gnu.org/licenses/>.
*/

import "fmt"
import "strconv"
import "math"

// Engineering Notation uses powers of ten that are multiples of 3
// e.g., 10x-06, 10x-03, 10x+00, 10x+03, 10x+06, ...
//
// In the floating point contimuium of numbers the middle range of
// number from 10x-24 to 10x+24 have been categorized into
// 17 divisions or prefixes.  Each category of 3 (ex: 10x-21, 10x-20,
// 10x-19) are named (ex: milli, kilo, etc.) and each assigned a
// character to represent the category (ex: M for mega, G for giga,
// etc.).
// For example: The kilo category has the prefix "k" for the range
// 10x+03 to 10x+05.
//
// This categoration is called Engineering Notation.
//
// This package, en, facilitates the entry and display of numbers with
// this categorization into and out of your floating point
// calculations.
//
// For an explanation of Engineering Notation see:
//   http://en.wikipedia.org/wiki/Engineering_notation

// Engineering Notation index
const (
	Yotta = 24  // Y 1,000,000,000,000,000,000,000,000
	Zetta = 21  // Z 1,000,000,000,000,000,000,000
	Exa   = 18  // E 1,000,000,000,000,000,000
	Peta  = 15  // P 1,000,000,000,000,000
	Tera  = 12  // T 1,000,000,000,000
	Giga  = 9   // G 1,000,000,000
	Mega  = 6   // M 1,000,000
	Kilo  = 3   // k 1,000
	Unit  = 0   //   1
	Milli = -3  // m 0.001
	Micro = -6  // u 0.000,001
	Nano  = -9  // n 0.000,000,001
	Pico  = -12 // p 0.000,000,000,001
	Femto = -15 // f 0.000,000,000,000,001
	Atto  = -18 // a 0.000,000,000,000,000,001
	Zepto = -21 // z 0.000,000,000,000,000,000,001
	Yocto = -24 // y 0.000,000,000,000,000,000,000,001
)

// Electronic Unit Abbreviations
const (
	Amp   = "A"
	Volt  = "V"
	Ohm   = "Ω"
	Hertz = "Hz"
	Farad = "F"
	Henry = "H"
	Watt  = "W"
)

// Returns a float's mantissa and exponent.
// en.FtoME(-234.5e-03) returns the mantissa and exponent for
// the specified floating point, e.g., -2.345 and -1
func FtoME(f float64) (m float64, e int) {
	s, _, fstE := floatToDigits(f) // -m.mmme-ee
	// Break out the mantissa (m.mmm).
	m, _ = strconv.ParseFloat(s[:fstE-1], 64)
	// Break out the exponent (ee)
	e, _ = strconv.Atoi(s[fstE:])
	return
}

// en.EntoF(number, category) takes a floating point number and
// adjusts it to an engineering notation category.  So if the number
// (It can be any valid floating point number.) is say 1.23456 and the
// category is kilo (1.23456 kilos) then EntoF() will return the
// number as
// 1.23456e03 which is 1.23456 kilos or 123.456 units of the item.
// The returned adjusted number can then be used as any other floating
// point number.  en.EntoF(1234.56, en.Kilo) returns 123.3456e+06
func EntoF(mantissa float64, expEn int) (f float64) {
	// Identify parts location and size, e.g., -m.mmmx-ee.
	str, _, fstE := floatToDigits(mantissa)

	// Break out the mantissa (m.mmm).
	man, _ := strconv.ParseFloat(str[:fstE-1], 64)

	// Break out the exponent (ee)
	exp, _ := strconv.Atoi(str[fstE:])
	adjExp := exp + expEn

	// Build internal float from the mantissa and the supplied
	// Engineering Notation entry.
	f = man * math.Pow(10.0, float64(adjExp))
	return
}

// FtoEn(number) returns a string of the number in the appropriate
// engineering notation category, i.e., If the number 2.3456e07 is
// specified then the string "23.5M" (23.5 mega) will be returned.
func FtoEn(f float64) (enNotated string) {
	enNotated = floatToEn(f)
	return
}

// GetEnCode(exponent) returns the Engineering Notation code for the
// specified exponent, e.g., en.GetEnCode(en.Micro) returns "µ"
func GetEnCode(exp int) (c string) {
	if float64(exp) >= Yocto && float64(exp) <= Yotta {
		// look up the en code using our en index
		enExpCode := []string{
			"y", "z", "a", "f", "p", "n", "µ", "m",
			"",
			"k", "M", "G", "T", "P", "E", "Z", "Y"}
		// shift e to enExpCode index range
		cIdx := (exp / 3) + (len(enExpCode) / 2)
		c = enExpCode[cIdx]
	} else {
		c = "" // out-of-range for engineering notation
	}
	return
}

func floatToDigits(f float64) (str string, fstM, fstE int) {
	// Convert the float to a string so it can be picked apart.
	str = fmt.Sprintf("%e", f) // -M.MMMe-EE
	if str[:1] == "-" {
		fstM = 1
	} else {
		fstM = 0
	}
	for {
		if str[fstE:fstE+1] == "e" {
			break
		}
		fstE++
	}
	fstE++
	return
}

func getDigits(f float64) (d [4]string, e float64) {
	minus := false
	// pick off the digits
	more := 2 // set for 2 passes just in case we need to round.
	for more > 0 {
		// Get the start points.
		str, fstM, fstE := floatToDigits(f) // -M.MMMe-EE

		// pull the exponent
		e, _ = strconv.ParseFloat(str[fstE:], 64)

		// Determine if the mantissa is minus or not.
		if str[0:1] == "-" {
			minus = true
		}
		// pull the mantissa  M.MMMe-EE
		d[0] = str[0 : fstM+1]
		for i := 1; i <= 3; i++ { // skip over the period
			d[i] = str[fstM+i+1 : fstM+i+2]
		}
		// Round if needed.
		if d[3] >= "5" && d[3] <= "9" {
			// we need to round
			roundVal := 0.01
			if minus {
				roundVal = -roundVal
			}
			f = (roundVal * math.Pow(10.0, e)) + f
			more -= 1
		} else {
			more = 0 // round not needed
		}
	}
	return
}

func floatToEn(f float64) (s string) {
	var c string
	d, fExp := getDigits(f)
	exp := int(fExp)
	// Determine where the period should go.
	pFmt := []string{
		// mx-02     mx-01     mx+00      mx+01     mx+02
		"%s%s.%s", "%s%s%s", "%s.%s%s", "%s%s.%s", "%s%s%s"}
	pIdx := (exp % 3) + 2 // shift value -2,-1,0,1,2 => 0,1,2,3,4
	e := []int{exp - 1, exp - 2, exp, exp - 1, exp - 2}[pIdx]
	// return number in engineering notation
	m := fmt.Sprintf(pFmt[pIdx], d[0], d[1], d[2])
	c = GetEnCode(e)
	// GetEnCode returns a "" string for both the out-of-range
	// condition and en.Unit.
	if c != "" || e == Unit {
		s = fmt.Sprintf("%s%s", m, c) // MMMC, e.g., "123k"
	} else { // number is out of the engineering notation range
		s = fmt.Sprintf("%se%d", m, e) // MMMeEE, e.g., "1.23e30"
	}
	return
}
