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
// 10x+01, 10x+02, 10x+03.
//
// This categoration is called Engineering Notation.
//
// This package, en, facilitates the entry and display of numbers with
// this categorization into and out of your floating point
// calculations.

// Engineering Notation index
const (
	Yotta =  24 // 16 Y 1,000,000,000,000,000,000,000,000
	Zetta =  21 // 15 Z 1,000,000,000,000,000,000,000
	Exa   =  18 // 14 E 1,000,000,000,000,000,000
	Peta  =  15 // 13 P 1,000,000,000,000,000
	Tera  =  12 // 12 T 1,000,000,000,000
	Giga  =   9 // 11 G 1,000,000,000
	Mega  =   6 // 10 M 1,000,000
	Kilo  =   3 //  9 k 1,000
	Unit  =   0 //  8   1
	Milli =  -3 //  7 m 0.001
	Micro =  -6 //  6 µ 0.000,001
	Nano  =  -9 //  5 n 0.000,000,001
	Pico  = -12 //  4 p 0.000,000,000,001
	Femto = -15 //  3 f 0.000,000,000,000,001
	Atto  = -18 //  2 a 0.000,000,000,000,000,001
	Zepto = -21 //  1 z 0.000,000,000,000,000,000,001
	Yocto = -24 //  0 y 0.000,000,000,000,000,000,000,001
)

var enCode = []string{
	"y", "z", "a", "f", "p", "n", "µ", "m",
	"",
	"k", "M", "G", "T", "P", "E", "Z", "Y",
}

// Electronic Unit Abbreviations
const (
	About      = "≈"
	NotEq      = "≠"
	Amp        = "A"
	Volt       = "V"
	Ohm        = "Ω"
	Hertz      = "Hz"
	Farad      = "F"
	Henry      = "H"
	Watt       = "W"
	Reluctance = "R"
	Alpha      = "α"
	Beta       = "β"
	Delta      = "δ"
	Pi         = "π"
	Tau        = "τ"
	Theta      = "θ"
	Phi        = "Φ"
	Lambda     = "λ"
	Degree     = "°"
)

// Conversion Factors
const (
	RadToDeg  = 180/math.Pi
	DegToRad  = math.Pi/180
	RadToGrad = 200/math.Pi
	GradToDeg = math.Pi/200
)

// FtoMe(float64) returns a float's mantissa and exponent as it will
// be normalized as a float64 number.
// Therefore FtoME(-234.5e-03) will return (-2.345,-1)
func FtoME(f float64) (m float64, e int) {
	s := fmt.Sprintf("%e", f) // -M.MMMe-EE
	for i := len(s)-1; i >= 0; i-- {
		if string(s[i]) == "e" {
			m, _ = strconv.ParseFloat(s[:i], 64)
			e, _ = strconv.Atoi(s[i+1:])
			break
		}
	}
	return
}

// EntoF(float64, en_category) takes a floating point number and
// adjusts it to an engineering notation category.  So if the number
// (It can be any valid floating point number.) is 1.23456 and the
// category is kilo (en.EntoF(1.23456, en.Kilo)) then it will return a
// floating point number of 1.23456e+03 which is 1.23456 kilos.
func EntoF(mantissa float64, en int) (f float64) {
	m, e := FtoME(mantissa)
	f     = m * math.Pow(10.0, float64(e+en))
	return
}

// FtoEn(float64) returns a string of the number in the appropriate
// engineering notation category.  If the number 2.3456e07 is
// specified then the string "23.5M" (23.5 mega) will be returned.
// Note: The returned value is rounded to three significant digits.
func FtoEn(f float64) (en string) {
	m, e := encode(f)
	en = fmt.Sprintf("%s%s", m, e)
	return
}

// Code(exponent) returns the Engineering Notation
// period pattern and code for the specified exponent,
// e.g., en.Code(en.Micro-1) returns ("12.3","µ")
// Note: The period pattern will be "1.23", "12.3", or "123"
func Code(en int) (p, e string) {
	f := 123.0 * math.Pow(10.0, float64(en))
	p, e = encode(f)
	return
}

// encode(float64) returns a mantissa and exponent in Engineering
// Notation, e.g., encode(1.235e-12) returns ("1.24","p")
func encode(f float64) (am, ce string) {
	var d [4]string ; s := "" ; e := 0
	// pick off the digits
	more := 2 // set for a 2nd passes just in case we need to round.
	for more > 0 {
		// convert mantissa and exponent to a string
		s = fmt.Sprintf("%e", f)
		for i := len(s)-1; i >= 0; i-- {  // pull the exponent
			if string(s[i]) == "e" { e, _ = strconv.Atoi(s[i+1:]) ; break }
		}
		// Determine if the mantissa is minus or not.
		minus := false ; fstDigit := 0
		if s[:1] == "-" { minus = true ; fstDigit = 1 }
		// pull the mantissa  M.MMMe-EE
		d[0] = s[:fstDigit+1]
		for i := 1; i <= 3; i++ { // skip over the period
			if string(s[fstDigit+i+1:fstDigit+i+2]) == "r" { break }
			d[i] = s[fstDigit+i+1:fstDigit+i+2]
		}
		// Round if needed.
		if d[3] >= "5" && d[3] <= "9" {
			roundVal := 0.01   // we need to round
			if minus { roundVal = -roundVal }
			f = (roundVal * math.Pow(10.0, float64(e))) + f
			more -= 1  // since we rounded we need to redo this process
		} else { more = 0 }  // round not needed
	}
	// Determine where the period should go.
	pIdx := (e % 3) + 2 // shift value -2,-1,0,1,2 => 0,1,2,3,4
	// return mantissa in engineering notation
	pFmt := []string{
		"%s%s.%s", "%s%s%s", "%s.%s%s", "%s%s.%s", "%s%s%s"}
	am = fmt.Sprintf(pFmt[pIdx], d[0], d[1], d[2])
	// adjust the exponent to the change we made to the period
	ec := []int{e - 1, e - 2, e, e - 1, e - 2}[pIdx]
	switch {
	case float64(ec) < (Yocto-2) :  // below the lower limit
		ce = fmt.Sprintf("e%3d", ec)
	case float64(ec) > (Yotta+2) :  // above the upper limit
		ce = fmt.Sprintf("e+%2d", ec)
	default:
		// shift ec to the enCode index range
		cIdx := (ec / 3) + (len(enCode) / 2)
		ce = enCode[cIdx]  // look up the en code using our en index
	}
	return
}

