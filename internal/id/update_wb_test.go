// Copyright 2020 Thomas.Hoehenleitner [at] seerose.net
// Use of this source code is governed by a license that can be found in the LICENSE file.
package id

import (
	"fmt"
	"testing"

	"github.com/rokath/trice/pkg/tst"
)

type idCheck struct {
	nbTRICE string
	nbID    string
	id      TriceID
	ok      bool
	tf      TriceFmt
}

// to do: improve parser that this works
func _TestUpdateParamCount1(t *testing.T) {
	text := `
	TRICE8_8( "hi %2d, %13u, %64b, %8x %02d, %013u, %032b, %016x" );
`
	exp := `
	TRICE8_8( Id(0), "hi %2d, %13u, %64b, %8x %02d, %013u, %032b, %016x" );
`
	act := updateParamCount(text)
	tst.Equal(t, exp, act)
}

func TestUpdateParamCount0(t *testing.T) {
	text := `
	TRICE0 ( "hi");
	TRICE0( "hi");
	TRICE8( "hi %d", 5);
	TRICE8( "hi %d, %u", 5, h);
	TRICE8( "hi %d, %u, %b", d, u, b);
	TRICE8( "hi %d, %u, %b, %x", d, u, b, h);
	TRICE8( "hi %d, %u, %b, %x %d, %u, %b, %x", d, u, b, h, d, u, b, h);
	TRICE8( "hi %2d, %13u, %64b, %8x %02d, %013u, %032b, %016x", d, u, b, h, d, u, b, h);
	TRICE16( "hi %d", 5);
	TRICE16( "hi %d, %u", 5, h);
	TRICE16( "hi %d, %u, %b", d, u, b);
	TRICE16( "hi %d, %u, %b, %x", d, u, b, h);
	TRICE32( "hi %d", 5);
	TRICE32( "hi %d, %u", 5, h);
	TRICE32( "hi %d, %u, %b", d, u, b);
	TRICE32( "hi %d, %u, %b, %x", d, u, b, h);
	TRICE64( "hi %d", 5);
	TRICE64( "hi %d, %u", 5, h);
	trice0 ( "hi");
	trice0( "hi");
	trice8( "hi %d", 5);
	trice8( "hi %d, %u", 5, h);
	trice8( "hi %d, %u, %b", d, u, b);
	trice8( "hi %d, %u, %b, %x", d, u, b, h);
	trice8( "hi %d, %u, %b, %x %d, %u, %b, %x", d, u, b, h, d, u, b, h);
	trice8( "hi %2d, %13u, %64b, %8x %02d, %013u, %032b, %016x", d, u, b, h, d, u, b, h);
	trice16( "hi %d", 5);
	trice16( "hi %d, %u", 5, h);
	trice16( "hi %d, %u, %b", d, u, b);
	trice16( "hi %d, %u, %b, %x", d, u, b, h);
	trice32( "hi %d", 5);
	trice32( "hi %d, %u", 5, h);
	trice32( "hi %d, %u, %b", d, u, b);
	trice32( "hi %d, %u, %b, %x", d, u, b, h);
	trice64( "hi %d", 5);
	trice64( "hi %d, %u", 5, h);
	`
	exp := `
	TRICE0 ( Id(0), "hi");
	TRICE0( Id(0), "hi");
	TRICE8_1( Id(0), "hi %d", 5);
	TRICE8_2( Id(0), "hi %d, %u", 5, h);
	TRICE8_3( Id(0), "hi %d, %u, %b", d, u, b);
	TRICE8_4( Id(0), "hi %d, %u, %b, %x", d, u, b, h);
	TRICE8_8( Id(0), "hi %d, %u, %b, %x %d, %u, %b, %x", d, u, b, h, d, u, b, h);
	TRICE8_8( Id(0), "hi %2d, %13u, %64b, %8x %02d, %013u, %032b, %016x", d, u, b, h, d, u, b, h);
	TRICE16_1( Id(0), "hi %d", 5);
	TRICE16_2( Id(0), "hi %d, %u", 5, h);
	TRICE16_3( Id(0), "hi %d, %u, %b", d, u, b);
	TRICE16_4( Id(0), "hi %d, %u, %b, %x", d, u, b, h);
	TRICE32_1( Id(0), "hi %d", 5);
	TRICE32_2( Id(0), "hi %d, %u", 5, h);
	TRICE32_3( Id(0), "hi %d, %u, %b", d, u, b);
	TRICE32_4( Id(0), "hi %d, %u, %b, %x", d, u, b, h);
	TRICE64_1( Id(0), "hi %d", 5);
	TRICE64_2( Id(0), "hi %d, %u", 5, h);
	trice0 ( Id(0), "hi");
	trice0( Id(0), "hi");
	trice8_1( Id(0), "hi %d", 5);
	trice8_2( Id(0), "hi %d, %u", 5, h);
	trice8_3( Id(0), "hi %d, %u, %b", d, u, b);
	trice8_4( Id(0), "hi %d, %u, %b, %x", d, u, b, h);
	trice8_8( Id(0), "hi %d, %u, %b, %x %d, %u, %b, %x", d, u, b, h, d, u, b, h);
	trice8_8( Id(0), "hi %2d, %13u, %64b, %8x %02d, %013u, %032b, %016x", d, u, b, h, d, u, b, h);
	trice16_1( Id(0), "hi %d", 5);
	trice16_2( Id(0), "hi %d, %u", 5, h);
	trice16_3( Id(0), "hi %d, %u, %b", d, u, b);
	trice16_4( Id(0), "hi %d, %u, %b, %x", d, u, b, h);
	trice32_1( Id(0), "hi %d", 5);
	trice32_2( Id(0), "hi %d, %u", 5, h);
	trice32_3( Id(0), "hi %d, %u, %b", d, u, b);
	trice32_4( Id(0), "hi %d, %u, %b, %x", d, u, b, h);
	trice64_1( Id(0), "hi %d", 5);
	trice64_2( Id(0), "hi %d, %u", 5, h);
	`
	act := updateParamCount(text)
	tst.Equal(t, exp, act)
}

func TestUpdateIDsShared(t *testing.T) {
	text := `
	trice8_1( Id(0), "hi %d", 5); // first id
	trice8_1( Id(0), "Hi %d", 5); // different format string needs a new id
	TRICE8_1( Id(       0   ), "Hi %d", 5); // different type case gets same id
	trice8_1( Id(0x0), "Hi %d", 5); // this gets not changed because no decimal number
	TRICE8_1( Id(0), "hi %d", 5); // same format string gets same id
	trice8_1(  Id( 0       ),  "hi %d", 5); // same format string gets same id
	trice8_1( Id(-0), "hi %d", 5); // minus id's are not touched
	`
	exp := `
	trice8_1( Id(   99), "hi %d", 5); // first id
	trice8_1( Id(   98), "Hi %d", 5); // different format string needs a new id
	TRICE8_1( Id(   98), "Hi %d", 5); // different type case gets same id
	trice8_1( Id(0x0), "Hi %d", 5); // this gets not changed because no decimal number
	TRICE8_1( Id(   99), "hi %d", 5); // same format string gets same id
	trice8_1(  Id(   99),  "hi %d", 5); // same format string gets same id
	trice8_1( Id(-0), "hi %d", 5); // minus id's are not touched
	`
	SearchMethod = "downward"
	Min = 10
	Max = 99
	lu := make(TriceIDLookUp)
	tflu := lu.reverse()
	modified := false

	act := updateIDsShared(text, lu, tflu, &modified)
	tst.Equal(t, exp, act)
}

var tryOkSet = []idCheck{
	{`TRICE0(Id(   59), "tt" )`, "Id(   59)", 59, true, TriceFmt{"TRICE0", "tt"}},
	{`TRICE0(Id(59   ), "tt" )`, "Id(59   )", 59, true, TriceFmt{"TRICE0", "tt"}},
	{`TRICE0(Id(59), "tt" )`, "Id(59)", 59, true, TriceFmt{"TRICE0", "tt"}},
	{`TRICE0(Id( 59 ), "tt" )`, "Id( 59 )", 59, true, TriceFmt{"TRICE0", "tt"}},
	{`trice0(Id(59), "tt" )`, "Id(59)", 59, true, TriceFmt{"trice0", "tt"}},
	{`trice64_2(Id(59), "%d,%x", -3, -4 )`, "Id(59)", 59, true, TriceFmt{"trice64_2", "%d,%x"}},
}

var tryNotOkSetID = []idCheck{
	{`TRICE0(Id(0x5), "tt" )`, "Id(0x5)", 5, false, TriceFmt{"TRICE0", "tt"}},
	{`TRICE0(id( 59 ), "tt" )`, "id( 59  )", 59, false, TriceFmt{"TRICE0", "tt"}},
	{`TRICE0(id(0x5 ), "tt" )`, "id(0x59)", 0x59, false, TriceFmt{"TRICE0", "tt"}},
}

func TestTriceParseOK(t *testing.T) {
	set := tryOkSet
	for i := range set {
		nbID, id, tf, ok := triceParse(set[i].nbTRICE)
		tst.AssertTrue(t, ok == set[i].ok)
		if ok {
			checkID(t, set, i, id)
			checkNbID(t, set, i, nbID)
			tst.AssertTrueElseInfo(t, tf == tryOkSet[i].tf, fmt.Sprint(tf, tryOkSet[i].tf))
		}
	}
}

func TestTriceIDParseOK(t *testing.T) {
	set := tryOkSet
	for i := range set {
		nbID, id, ok := triceIDParse(set[i].nbTRICE)
		tst.AssertTrue(t, ok == set[i].ok)
		if ok {
			checkID(t, set, i, id)
			checkNbID(t, set, i, nbID)
		}
	}
}

func TestTriceIDParseNotOK(t *testing.T) {
	set := tryNotOkSetID
	for i := range set {
		nbID, id, ok := triceIDParse(set[i].nbTRICE)
		tst.AssertTrueElseInfo(t, ok == set[i].ok, fmt.Sprint(i))
		if ok {
			checkID(t, set, i, id)
			checkNbID(t, set, i, nbID)
		}
	}
}

func TestTriceFmtParse(t *testing.T) {
	for i := range tryOkSet {
		tf, ok := triceFmtParse(tryOkSet[i].nbTRICE)
		tst.AssertTrue(t, ok == tryOkSet[i].ok)
		tst.AssertTrueElseInfo(t, tf == tryOkSet[i].tf, fmt.Sprint(tf, tryOkSet[i].tf))
	}
}

func checkID(t *testing.T, set []idCheck, i int, id TriceID) {
	tst.AssertTrueElseInfo(t, id == set[i].id, fmt.Sprint(i, id))
}

func checkNbID(t *testing.T, set []idCheck, i int, nbID string) {
	tst.AssertTrueElseInfo(t, nbID == set[i].nbID, fmt.Sprint(i, nbID))
}