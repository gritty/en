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
import "errors"

// Engineering Notation index
const (
	Yotta = 16 // Y  24 1,000,000,000,000,000,000,000,000
	Zetta = 15 // Z  21 1,000,000,000,000,000,000,000
	Exa   = 14 // E  18 1,000,000,000,000,000,000
	Peta  = 13 // P  15 1,000,000,000,000,000
	Tera  = 12 // T  12 1,000,000,000,000
	Giga  = 11 // G   9 1,000,000,000
	Mega  = 10 // M   6 1,000,000
	Kilo  = 9  // k   3 1,000
	Unit  = 8  //     0 1
	Milli = 7  // m - 3 0.001
	Micro = 6  // u - 6 0.000,001
	Nano  = 5  // n - 9 0.000,000,001
	Pico  = 4  // p -12 0.000,000,000,001
	Femto = 3  // f -15 0.000,000,000,000,001
	Atto  = 2  // a -18 0.000,000,000,000,000,001
	Zepto = 1  // z -21 0.000,000,000,000,000,000,001
	Yocto = 0  // y -24 0.000,000,000,000,000,000,000,001
)

var enExpCode = []string{
	"y", "z", "a", "f", "p", "n", "u", "m",
	" ",
	"k", "M", "G", "T", "P", "E", "Z", "Y"}

// EnToFloat converts a float64 + its "en" code
// to a pure float64.
//   en.EnToFloat(632.5, en.Nano) returns 6.325e-07
func EnToFloat(mantissa float64, expEnIdx int) (floatVal float64,
err error) {
	// is exponent within range?
	if expEnIdx < 0 || expEnIdx >= len(enExpCode) {
    err = errors.New("en: Invalid exponent code or out-of-range.")
    return
	}
	// convert the mantissa to a string so we can pick it apart
	str := fmt.Sprintf("%.3e", mantissa) // -n.nnne-nn
	_, mSize, _, eIdx := parseMantissa(str)
	man, _ := strconv.ParseFloat(str[:mSize], 64)
	expVal, _ := strconv.Atoi(str[eIdx:])
	// adjust exponent to a new en value
	adjExpVal := (expEnIdx-8)*3 + expVal
	floatVal = man * math.Pow(10.0, float64(adjExpVal))
	return
}

// FloatToEn converts a float64 to its "en" equivilent
// rounded to 3 significant digits.
//   en.FloatToEn(6.325e-07) returns "633 n"
func FloatToEn(f float64) (result string) {
  str, fstDigit, eIdx := roundMantissa(f)  // round if needed
	// pick off the digits
	var d = []string{ // skip over the period
		str[0 : fstDigit+1],
		str[fstDigit+2 : fstDigit+3],
		str[fstDigit+3 : fstDigit+4],
	}
	// get the exponent
	exp, _ := strconv.Atoi(str[eIdx:])
	// determine where the period should go and get the en code
	var pFmt = []string{
		"%s%s.%s", "%s%s%s", "%s.%s%s", "%s%s.%s", "%s%s%s"}
	pIdx := (exp % 3) + 2 // shift value -2,-1,0,1,2 => 0,1,2,3,4
	enExp := []int{exp - 1, exp - 2, exp, exp - 1, exp - 2}[pIdx]
	// return number in engineering notation
	enMan := fmt.Sprintf(pFmt[pIdx], d[0], d[1], d[2])
	// look up the en code using our en index
	enIdx := (enExp / 3) + 8 // shift enExp to 0-len(enExpCode) range
	if enIdx >= 0 && enIdx < len(enExpCode) {
		result = enMan + " " + enExpCode[enIdx]
	} else {
		// Out of Range
		result = fmt.Sprintf("%se%s", enMan, strconv.Itoa(enExp))
	}
	return
}

// Parse breaks out a float64 number into its engineering notation
// components, e.g., mantissa, exponent, index, and code.
//    en.Parse(6.325e-07) returns "633.00", -9, 5, "n"
func Parse(f float64) (m string, e int, i int, c string) {
  str, fstDigit, eIdx := roundMantissa(f) // round if needed
	// pick off the digits
	var d = []string{ // skip over the period
		str[0 : fstDigit+1],
		str[fstDigit+2 : fstDigit+3],
		str[fstDigit+3 : fstDigit+4],
	}
	// get the exponent
	exp, _ := strconv.Atoi(str[eIdx:])
	// determine where the period should go and get the en code
	var pFmt = []string{
		" %s%s.%s0", "%s%s%s.00", "  %s.%s%s", " %s%s.%s0", "%s%s%s.00"}
	pIdx := (exp % 3) + 2 // shift value -2,-1,0,1,2 => 0,1,2,3,4
	enExp := []int{exp - 1, exp - 2, exp, exp - 1, exp - 2}[pIdx]
	// return number in engineering notation
	enMan := fmt.Sprintf(pFmt[pIdx], d[0], d[1], d[2])
	// look up the en code using our en index
	enIdx := (enExp / 3) + 8 // shift enExp to 0-len(enExpCode) range
	m = enMan
	e = enExp
	i = enIdx
	if enIdx >= 0 && enIdx < len(enExpCode) {
		c = enExpCode[i]
	} else {
		c = fmt.Sprintf("%se%s", enMan, strconv.Itoa(e))  // Out of Range
	}
	return
}

func roundMantissa(f float64) (str string, fstDigit, eIdx int) {
	str = fmt.Sprintf("%.3e", f)
  var negSign bool
	negSign, _, fstDigit, eIdx = parseMantissa(str)
	tstDig := str[eIdx-2] - 48  // adjust test digit to 0-9
	if tstDig >= 5 && tstDig <= 9 { // ck the last digit of the mantissa
		exp, _ := strconv.Atoi(str[eIdx:])
		rndVal := 0.01 ; if negSign {	rndVal = -rndVal }
		ff := (rndVal * math.Pow(10.0, float64(exp))) + f
		str = fmt.Sprintf("%.3e", ff)
		negSign, _, fstDigit, eIdx = parseMantissa(str)
	}
  return
}

func parseMantissa(s string) (neg bool, size, first, exp int) {
	neg = false
	size = 5
	first = 0
	exp = 6
	if s[:1] == "-" {
		neg = true
		size += 1  // include the minus sign
		first += 1 // first actual digit
		exp += 1   // move the exponent start position past the "e"
	}
	return
}

/*
	str := fmt.Sprintf("%.3e", f)
	negSign, _, fstDigit, eIdx := parseMantissa(str)
	// round the float if needed
//	if str[eIdx-2] != 0 { // ck the last digit of the mantissa
tstDig := str[eIdx-2] - 48  // adjust test digit to 0-9
fmt.Println("*** str[",eIdx-2, "]=", str[eIdx-2],
", tstDig=", tstDig, ", neg=", negSign)
if tstDig >= 5 && tstDig <= 9 { // ck the last digit of the mantissa
		exp, _ := strconv.Atoi(str[eIdx:])
//		rndVal := 0.005
rndVal := 0.01
		if negSign {
			rndVal = -rndVal
		}
		ff := (rndVal * math.Pow(10.0, float64(exp))) + f
		str = fmt.Sprintf("%.3e", ff)
		negSign, _, fstDigit, eIdx = parseMantissa(str)
	}
*/

