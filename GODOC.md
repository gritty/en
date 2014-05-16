**PACKAGE DOCUMENTATION**

 ```
package en
    import "."

    Package "en" implements function calls to convert floating point 
    numbers to and from Engineering Notation.
 ```

**CONSTANTS**

 ```
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
    Engineering Notation index
 ```

**FUNCTIONS**

 ```
func EnToFloat(mantissa float64, expEnIdx int) (floatVal float64)
    EnToFloat converts a float64 + its "en" code to a pure float64.

	en.EnToFloat(632.5, en.Nano) returns 6.325e-07

func FloatToEn(f float64) (result string)
    FloatToEn converts a float64 to its "en" equivilent rounded to 3
    significant digits.

	en.FloatToEn(6.325e-07) returns "633 n"

func Parse(f float64) (m string, e int, i int, c string)
    Parse breaks out a float64 number into its engineering notation
    components, e.g., mantissa, exponent, index, and code.

	en.Parse(6.325e-07) returns "633.00", -9, 5, "n"
 ```

