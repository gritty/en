package en

import "testing"
import "strconv"

// test EntoF() and FtoEn()
type testSet1 struct {
	fIn   float64
	cIn   int
	fOut  float64
	enOut string
}

var test1 = []testSet1{
	{632.5, Nano, 6.325e-07, "633n"},
	{-632.5, Nano, -6.325e-07, "-633n"},
	{632.5, Kilo, 632500, "633k"},
	{-632.5, Kilo, -632500, "-633k"},
	{632, Nano, 6.32e-07, "632n"},
	{-632, Nano, -6.32e-07, "-632n"},
	{632, Kilo, 632000, "632k"},
	{-632, Kilo, -632000, "-632k"},
}

// test FtoME()
type testSet2 struct {
	fIn  float64
	mOut float64
	eOut int
}

var test2 = []testSet2{
	{0.123, 1.23, -1},
	{123.00, 1.23, 2},
	{1230.00, 1.23, 3},
	{-1230.00, -1.23, 3},
	{-123.00, -1.23, 2},
	{-0.123, -1.23, -1},
}

// test Code(), FtoEn(), FtoME(), and EntoF()
type testSet3 struct {
	mIn   string
	iIn   int
	enOut string
	mOut  float64
	eOut  int
	mcOut string
	ecOut string
}

var test3 = []testSet3{
	{"0.123", Milli, "123µ", 1.23, -4, "12.3", "m"},
	{"123.00", Nano, "123n", 1.23, -7, "12.3", "n"},
	{"1230.00", Kilo, "1.23M", 1.23, 6, "12.3", "k"},
	{"-1230.00", Kilo, "-1.23M", -1.23, 6, "12.3", "k"},
	{"-123.00", Nano, "-123n", -1.23, -7, "12.3", "n"},
	{"-0.123", Milli, "-123µ", -1.23, -4, "12.3", "m"},
}

func TestEn(t *testing.T) {
	for _, test := range test1 {
		fOut := EntoF(test.fIn, test.cIn)
		enOut := FtoEn(fOut)
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
		mOut, eOut := FtoME(test.fIn)
		if mOut != test.mOut || eOut != test.eOut {
			t.Error(
				"For", test.fIn,
				"expected", test.mOut, test.eOut,
				"got", mOut, eOut,
			)
		}
	}
	for _, test := range test3 {
		fIn := 10.0e00
		fOut, _ := strconv.ParseFloat(test.mIn, 64)
		fOut = EntoF(fOut, test.iIn)
		fOut = fOut / fIn
		fOut = fOut * (fIn * fIn)
		fOut = fOut / fIn
		mOut, eOut := FtoME(fOut)
		enOut := FtoEn(fOut)
		mcOut, ecOut := Code(test.iIn-1)
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
		if mcOut != test.mcOut {
			t.Error(
				"For", test.mcOut,
				"expected", test.mcOut,
				"got", mcOut,
			)
		}
		if ecOut != test.ecOut {
			t.Error(
				"For", test.ecOut,
				"expected", test.ecOut,
				"got", ecOut,
			)
		}
	}
}
