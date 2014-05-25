**PACKAGE DOCUMENTATION**

 ```
package en
    import "github.com/gritty/en"

    Package "en" implements function calls to convert floating point numbers
    to and from Engineering Notation.
 ```

**CONSTANTS**

 ```
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
    Engineering Notation index
 ```

 ```
const (
    Amp   = "A"
    Volt  = "V"
    Ohm   = "Ω"
    Hertz = "Hz"
    Farad = "F"
    Henry = "H"
    Watt  = "W"
)
    Electronic Unit Abbreviations
 ```

**FUNCTIONS**

 ```
func EntoF(mantissa float64, expEn int) (f float64)
    en.EntoF(number, category) takes a floating point number and adjusts it
    to an engineering notation category. So if the number (It can be any
    valid floating point number.) is say 1.23456 and the category is kilo
    (1.23456 kilos) than the number will be stored as 1.23456e03 which is
    1.23456 kilos or 123.456 units of the item. The returned adjusted number
    can then be used as any other floating point number. en.EntoF(1234.56,
    en.Kilo) returns 123.3456e+06

func FtoEn(f float64) (enNotated string)
    FtoEn(number, category) returns a string of the number in the format of
    an engineering notation category. For the number 2.3456e03, is specified
    with a category Kilo then FtoEn() will returned as "235k" or 123 kilos.
    Note, that FtoEn() rounds the number to three digits.

    FtoEn(number) returns a string of number in the appropriate engineering
    notation category, i.e., If the number 2.3456e07 is specified then the
    string "23.5 M" (23.5 mega) will be returned.

func FtoME(f float64) (m float64, e int)
    Returns a float's mantissa and exponent. en.FtoME(-234.5e-03) returns
    the floating point split into its mantissa and exponent, e.g., -2.345
    and -1

func GetEnCode(exp int) (c string)
    GetEnCode(exponent) returns the Engineering Notation for the specified
    exponent, e.g., en.GetEnCode(en.Micro) returns "µ"
 ```

