package en

import "testing"
import "strconv"

type testSet1 struct {
	fIn   float64
	cIn   int
	fOut  float64
	enOut string
}

var test1 = []testSet1{
	{632.5, Nano, 6.325e-07, "633 n"},
	{-632.5, Nano, -6.325e-07, "-633 n"},
	{632.5, Kilo, 632500, "633 k"},
	{-632.5, Kilo, -632500, "-633 k"},
	{632, Nano, 6.32e-07, "632 n"},
	{-632, Nano, -6.32e-07, "-632 n"},
	{632, Kilo, 632000, "632 k"},
	{-632, Kilo, -632000, "-632 k"},
}

type testSet2 struct {
	fIn  float64
	mOut string
	eOut int
	iOut int
	cOut string
}

var test2 = []testSet2{
	{0.123, "123.00", -3, 7, "m"},
	{123.00, "123.00", 0, 8, " "},
	{1230.00, "  1.23", 3, 9, "k"},
	{-1230.00, "  -1.23", 3, 9, "k"},
	{-123.00, "-123.00", 0, 8, " "},
	{-0.123, "-123.00", -3, 7, "m"},
}

type testSet3 struct {
	mIn   string
	iIn   int
	enOut string
	mOut  string
	eOut  int
	iOut  int
	cOut  string
}

var test3 = []testSet3{
	{"0.123", Milli, "123 u", "123.00", -6, 6, "u"},
	{"123.00", Nano, "123 n", "123.00", -9, 5, "n"},
	{"1230.00", Kilo, "1.23 M", "  1.23", 6, 10, "M"},
	{"-1230.00", Kilo, "-1.23 M", "  -1.23", 6, 10, "M"},
	{"-123.00", Nano, "-123 n", "-123.00", -9, 5, "n"},
	{"-0.123", Milli, "-123 u", "-123.00", -6, 6, "u"},
}

func TestEn(t *testing.T) {
	for _, test := range test1 {
		fOut := EnToFloat(test.fIn, test.cIn)
		enOut := FloatToEn(fOut)
		if fOut != test.fOut {
			t.Error(
				"For", test.fIn,
				"and", test.cIn,
				"expected", test.fOut,
				"got", fOut,
			)
		}
		if enOut != test.enOut {
			t.Error(
				"For", test.fIn,
				"and", test.cIn,
				"expected", test.enOut,
				"got", enOut,
			)
		}
	}
	for _, test := range test2 {
		mOut, eOut, iOut, cOut := Parse(test.fIn)
		if mOut != test.mOut || eOut != test.eOut || iOut != test.iOut || cOut != test.cOut {
			t.Error(
				"For", test.fIn,
				"expected", test.mOut, test.eOut, test.iOut, test.cOut,
				"got", mOut, eOut, iOut, cOut,
			)
		}
	}
	for _, test := range test3 {
		fIn := 10.0e00
		fOut, _ := strconv.ParseFloat(test.mIn, 64)
		fOut = EnToFloat(fOut, test.iIn)
		fOut = fOut / fIn
		fOut = fOut * (fIn * fIn)
		fOut = fOut / fIn
		mOut, eOut, iOut, cOut := Parse(fOut)
		enOut := FloatToEn(fOut)
		if enOut != test.enOut {
			t.Error(
				"For", test.enOut,
				"expected", test.enOut,
				"got", enOut,
			)
		}
		if mOut != test.mOut {
			t.Error(
				"For", test.mOut,
				"expected", test.mOut,
				"got", mOut,
			)
		}
		if eOut != test.eOut {
			t.Error(
				"For", test.eOut,
				"expected", test.eOut,
				"got", eOut,
			)
		}
		if iOut != test.iOut {
			t.Error(
				"For", test.iOut,
				"expected", test.iOut,
				"got", iOut,
			)
		}
		if cOut != test.cOut {
			t.Error(
				"For", test.cOut,
				"expected", test.cOut,
				"got", cOut,
			)
		}
	}
}
