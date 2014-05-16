en: Engineering Notation for the go language
==

Package "en" implements function calls to convert floating point numbers to and from Engineering Notation.

Copyright © 2014, Kent Loobey <kent@uoregon.edu>

Licensed under GNU General Public License 3.0.

See LICENSE.md file for details.

**_en contains three functions._**

1. **_EnToFloat()_** encodes a float64 value with an associated engineering notation constant to a standard float64 value.

2. **_FloatToEn()_** creates an engineering notation string from a standard float64 value.

3. **_Parse()_** breaks out a float64 number into its engineering notation components.

#**Example:**

 ```go
package main

import (
  "github.com/gritty/en"
  "fmt"
)

func main() {
  fmt.Printf("Testing Engineering Notation.\n\n")

  // EnToFloat converts an engineering notation number to float64.
  f64, err := en.EnToFloat(632.5, en.Nano) // returns 6.325e-07
  if err != nil {
    // Returns err string if out-of-range or invalid exponet detected.
    fmt.Println(err)
  }
  fmt.Println("en.EnToFloat(632.5, en.Nano)               returns:",
    f64)

  // FloatToEn converts a float64 number to its engineering notation.
  sEn := en.FloatToEn(6.325e-07)      // returns "633 n"
  fmt.Println("en.FloatToEn(6.325e-07)                    returns:",
    sEn)

  // Parse breaks out a float64 number into its engineering notation.
  m, e, x, c := en.Parse(6.325e-07)   // returns "633.00", -12, 3, "n"
  // m = mantissa, e = exponent, x = index, and  c = code
  fmt.Println("en.Parse(6.325e-07)                        returns:",
    m, e, x, c)

  fmt.Println()

  // Format returned of out-of-range Engineering Notation value.
  fmt.Println("OUT-OF-RANGE value en.FloatToEn(6.325e-32) returns:",
    en.FloatToEn(6.325e-32))  // returns 63.3e-33

  fmt.Println()

  // Four ways to convert "4.83 k" to float64.
  f1, _ := en.EnToFloat(4.83, en.Kilo)
  fmt.Printf("en.EnToFloat(4.83, en.Kilo)     returns: %.3e\n", f1)
  fmt.Println("en.FloatToEn(f1) returns:", en.FloatToEn(f1))
  f2, _ := en.EnToFloat(0.00483, en.Mega)
  fmt.Printf("en.EnToFloat(0.00483, en.Mega)  returns: %.3e\n", f2)
  fmt.Println("en.FloatToEn(f2) returns:", en.FloatToEn(f2))
  f3, _ := en.EnToFloat(4830.0, en.Unit)
  fmt.Printf("en.EnToFloat(4830.0, en.Unit)   returns: %.3e\n", f3)
  fmt.Println("en.FloatToEn(f3) returns:", en.FloatToEn(f3))
  f4, _ := en.EnToFloat(4830000, en.Milli)
  fmt.Printf("en.EnToFloat(4830000, en.Milli) returns: %.3e\n", f4)
  fmt.Println("en.FloatToEn(f4) returns:", en.FloatToEn(f4))

  // Calculate the Thevenin voltage and resistance for a circuit.
  fmt.Printf("\n%s\n",
    "Calculating Thevenin voltage and resistance for a circuit.")
  //   +---6kohm--+--4kOhm-A-+
  //   |          |          |
  //  72v       3kohm      Rlohm
  //   |          |          |
  //   +----------+--------B-+

  v1       := float64(72) // volts dc
  r1, _  := en.EnToFloat(6, en.Kilo)  // ohms
  r2, _  := en.EnToFloat(3, en.Kilo)  // ohms
  r3, _  := en.EnToFloat(4, en.Kilo)  // ohms

  // calculate Vth
  i := v1 / (r1 + r2)
  fmt.Println("i   =", en.FloatToEn(i), "amps")
  Vth := i * r2
  fmt.Println("Vth =", en.FloatToEn(Vth), "volts")

  // calculate Rth using product over sum
  Rth := r3 + (r2 * r1) / (r2 + r1)

  // Thevemom circuit:
  //   +---6kohm---A-+
  //   |             |
  //  24v          Rlohm
  //   |             |
  //   +-----------B-+

  fmt.Println("Rth =", en.FloatToEn(Rth), "ohms")
}
 ```

#**Output:**

 ```
Testing Engineering Notation.

en.EnToFloat(632.5, en.Nano)               returns: 6.325e-07
en.FloatToEn(6.325e-07)                    returns: 633 n
en.Parse(6.325e-07)                        returns: 633.00 -9 5 n

OUT-OF-RANGE value en.FloatToEn(6.325e-32) returns: 63.3e-33

en.EnToFloat(4.83, en.Kilo)     returns: 4.830e+03
en.FloatToEn(f1) returns: 4.83 k
en.EnToFloat(0.00483, en.Mega)  returns: 4.830e+03
en.FloatToEn(f2) returns: 4.83 k
en.EnToFloat(4830.0, en.Unit)   returns: 4.830e+03
en.FloatToEn(f3) returns: 4.83 k
en.EnToFloat(4830000, en.Milli) returns: 4.830e+03
en.FloatToEn(f4) returns: 4.83 k

Calculating Thevenin voltage and resistance for a circuit.
i   = 8.00 m amps
Vth = 24.0   volts
Rth = 6.00 k ohms
 ```

#**Documentation:**

You display the en documentation in Linux by entering the following 
within a terminal window or by just looking at the source code:

 ```
$ godoc "github.com/gritty/en"
 ```

#**Access:**

 ```
$ go get "github.com/gritty/en"
 ```

#**Limitations:**

en notation is valid only between Yocto (1.0e-24) and Yotta(1.0e24), e.g., 1/1,000,000,000,000,000,000,000 to 1,000,000,000,000,000,000,000,000

en.FloatToEn() rounds its value to three significant digits.

