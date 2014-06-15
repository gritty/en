en: Engineering Notation for the go language
============================================

version 0.2, 2014-06-14

Package "en" implements function calls to convert floating point numbers to and from Engineering Notation.

Copyright © 2014, Kent Loobey <kent@uoregon.edu>

Licensed under GNU General Public License 3.0.

See LICENSE.md file for details.

**_en contains four public functions._**

1. **_EntoF()_** takes a floating point number and an engineering notation category and returns the appropriate standard float64 value.

2. **_FtoEn()_** takes a standard float64 value and returns an engineering notated string.  Note" FtoEn() rounds the result to three significant digits.

3. **_FtoME()_** splits a float64 number into its engineering notation mantissa and exponent.

4. **_Code()_** returns the period pattern and exponent for the specified engineering notation code/prefix.

#**Example:**

 ```go
package main

import (
	"fmt"
	"math"
	"github.com/gritty/en"
)

func main() {
	fmt.Printf("\n%s\n",
		"Thevenin Equivalence Circuit: Calculate voltage and resistance:")

	//         r1         r3
	//     +---6kΩ----+---4kΩ---+
	//     |          |         |
	// v1 72Vdc   r2 3kΩ      RloadΩ
	//     |          |         |
	//     +----------+---------+

	v1 := float64(72)          // 72V dc
	r1 := en.EntoF(6, en.Kilo) // 6kΩ
	r2 := en.EntoF(3, en.Kilo) // 3kΩ
	r3 := en.EntoF(4, en.Kilo) // 4kΩ

	// calculate Vth
	i := v1 / (r1 + r2)
	Vth := i * r2

	// calculate Rth using product over sum
	Rth := r3 + (r2*r1)/(r2+r1)

	// Thevenin Equivalence Circuit:
	//            Rth
	//       +----6kΩ----+
	//       |           |
	// Vth 24Vdc       RloadΩ
	//       |           |
	//       +-----------+

	fmt.Println("I   =", en.FtoEn(i)+en.Amp)
	fmt.Println("Vth =", en.FtoEn(Vth)+en.Volt)
	fmt.Println("Rth =", en.FtoEn(Rth)+en.Ohm)

	fmt.Printf("\n%s\n",
		"Series RL Circuit: Calculate voltage, resistance, and impedance:")

	//         L
	//    +---330mH--+
	//    |          |
	// E 120Vac   R 68Ω
  // f 60Hz        |
	//    |          |
	//    +----------+

	// Xl = 2πfL
	// Zt = sqrt(Xl**2 + R**2)
	// I  = E / Zt
	// Vl = I * Xl
	// Vr = I * R
	// θ  = tan**-1(Vl / Vr)

	E := en.EntoF(120, en.Unit)
	f := en.EntoF(60, en.Unit)
	L := en.EntoF(330, en.Milli)
	R := en.EntoF(68, en.Unit)

	Xl    := 2.0 * math.Pi * f * L
	Zt    := math.Sqrt(Xl*Xl + R*R)
	I     := E / Zt
	Vl    := E * (Xl / Zt)
	Vr    := E * (R / Zt)
	theta := en.RadToDeg * math.Atan(Vl / Vr)

	fmt.Printf("Xl                  = %s\n", en.FtoEn(Xl)+en.Ohm)
	fmt.Printf("circuit impedance   = %s\n", en.FtoEn(Zt)+en.Ohm)
	fmt.Printf("circuit current     = %s\n", en.FtoEn(I)+en.Amp)
	fmt.Printf("magnitude of Vl     = %s\n", en.FtoEn(Vl)+en.Volt)
	fmt.Printf("magnitude of Vr     = %s\n", en.FtoEn(Vr)+en.Volt)
	fmt.Printf("circuit phase angle = %s\n", en.FtoEn(theta)+en.Degree)
}
 ```

#**Output:**

 ```
Thevenin Equivalence Circuit: Calculate voltage and resistance:
I   = 8.00mA
Vth = 24.0V
Rth = 6.00kΩ

Series RL Circuit: Calculate voltage, resistance, and impedance:
Xl                  = 124Ω
circuit impedance   = 142Ω
circuit current     = 846mA
magnitude of Vl     = 105V
magnitude of Vr     = 57.6V
circuit phase angle = 61.3°
 ```

#**Documentation:**

You can display the "en" documentation by entering the following command within a terminal window or by just looking at the source code:

 ```
godoc "github.com/gritty/en"
 ```

#**Access:**

"en" is gogetable.

 ```
go get "github.com/gritty/en"
 ```

#**Limitations:**

"en" can only return engineering notation for values between yocto (1.0e-24) and yotta(1.0e24), e.g., 0.000,000,000,000,000,000,000,001 through 1,000,000,000,000,000,000,000,000

en.FtoEn() rounds the return value to three significant digits.

