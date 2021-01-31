// Copyright 2020 Thomas.Hoehenleitner [at] seerose.net
// Use of this source code is governed by a license that can be found in the LICENSE file.

package decoder

import (
	"testing"
)

func TestBareL(t *testing.T) {
	doTableTest(t, NewBareDecoder, littleEndian, bareLTestTable, "unwrapped")
}

func TestBare(t *testing.T) {
	doTableTest(t, NewBareDecoder, bigEndian, bareTestTable, "unwrapped")
}

func Test2BareL(t *testing.T) {
	doTableTest(t, NewBareDecoder, littleEndian, bareLTestTable2, "unwrapped")
}

func Test2Bare(t *testing.T) {
	doTableTest(t, NewBareDecoder, bigEndian, bareTestTable2, "unwrapped")
}

var bareLTestTable = testTable{
	{[]byte{239, 205, 171, 137, 152, 1, 0, 0, 0, 0, 69, 70}, `MSG: triceFifoMaxDepth = 408, select = 0`},
	{[]byte{0, 0, 161, 95}, `--------------------------------------------------`},
	{[]byte{0, 0, 116, 52}, `--------------------------------------------------`},
	{[]byte{57, 48, 213, 70}, `dbg:12345 as 16bit is 0b0011000000111001`},
	{[]byte{0, 0, 207, 14}, `--------------------------------------------------`},
	{[]byte{0, 0, 229, 66}, `sig:This ASSERT error is just a demo and no real error:`},
	{[]byte{0, 0, 173, 104}, `--------------------------------------------------`},
	{[]byte{239, 205, 171, 137, 152, 1, 0, 0, 1, 0, 69, 70}, `MSG: triceFifoMaxDepth = 408, select = 1`},
	{[]byte{176, 180, 193, 70}, `ERR:error       message, SysTick is -19280`},
	{[]byte{158, 179, 129, 110}, `WRN:warning     message, SysTick is -19554`},
	{[]byte{147, 178, 56, 209}, `ATT:attension   message, SysTick is -19821`},
	{[]byte{136, 177, 32, 65}, `DIA:diagnostics message, SysTick is -20088`},
	{[]byte{125, 176, 222, 164}, `TIM:timing      message, SysTick is -20355`},
	{[]byte{114, 175, 165, 93}, `DBG:debug       message, SysTick is -20622`},
	{[]byte{239, 205, 171, 137, 152, 1, 0, 0, 2, 0, 69, 70}, `MSG: triceFifoMaxDepth = 408, select = 2`},
	{[]byte{177, 180, 162, 53}, `SIG:signal      message, SysTick is -19279`},
	{[]byte{159, 179, 238, 94}, `RD:read         message, SysTick is -19553`},
	{[]byte{148, 178, 58, 55}, `WR:write        message, SysTick is -19820`},
	{[]byte{137, 177, 149, 216}, `ISR:interrupt   message, SysTick is -20087`},
	{[]byte{126, 176, 23, 162}, `TST:test        message, SysTick is -20354`},
	{[]byte{115, 175, 129, 178}, `MSG:normal      message, SysTick is -20621`},
	{[]byte{104, 174, 13, 96}, `INFO:informal   message, SysTick is -20888`},
	{[]byte{0, 0, 0, 0, 88, 172, 128, 216}, `TST:test        message, SysTick is  44120`},
	{[]byte{0, 0, 0, 0, 73, 170, 128, 216}, `TST:test        message, SysTick is  43593`},
	{[]byte{0, 0, 0, 0, 61, 168, 128, 216}, `TST:test        message, SysTick is  43069`},
	{[]byte{0, 0, 0, 0, 49, 166, 128, 216}, `TST:test        message, SysTick is  42545`},
	{[]byte{0, 0, 0, 0, 37, 164, 128, 216}, `TST:test        message, SysTick is  42021`},
	{[]byte{0, 0, 0, 0, 25, 162, 128, 216}, `TST:test        message, SysTick is  41497`},
	{[]byte{0, 0, 0, 0, 13, 160, 128, 216}, `TST:test        message, SysTick is  40973`},
	{[]byte{239, 205, 171, 137, 152, 1, 0, 0, 3, 0, 69, 70}, `MSG: triceFifoMaxDepth = 408, select = 3`},
	{[]byte{127, 1, 0, 0, 255, 128, 223, 9}, `tst:TRICE8  %03x ->  001  07f  -80  -01`},
	{[]byte{127, 1, 0, 0, 255, 128, 191, 247}, `tst:TRICE8   %4d ->    1  127 -128   -1`},
	{[]byte{127, 1, 0, 0, 255, 128, 127, 156}, `tst:TRICE8   %4o ->    1  177 -200   -1`},
	{[]byte{1, 0, 0, 0, 255, 127, 0, 0, 0, 128, 0, 0, 255, 255, 226, 33}, `tst:TRICE16  %05x ->   00001   07fff   -8000   -0001`},
	{[]byte{1, 0, 0, 0, 255, 127, 0, 0, 0, 128, 0, 0, 255, 255, 221, 233}, `tst:TRICE16   %6d ->       1   32767  -32768      -1`},
	{[]byte{1, 0, 0, 0, 255, 127, 0, 0, 0, 128, 0, 0, 255, 255, 205, 96}, `tst:TRICE16   %7o ->       1   77777 -100000      -1`},
	{[]byte{239, 205, 171, 137, 152, 1, 0, 0, 4, 0, 69, 70}, `MSG: triceFifoMaxDepth = 408, select = 4`},
	{[]byte{0, 0, 0, 0, 1, 0, 0, 0, 255, 127, 0, 0, 255, 255, 0, 0, 0, 128, 0, 0, 0, 0, 0, 0, 255, 255, 0, 0, 255, 255, 171, 65}, `tst:TRICE32_4 %09x ->      000000001      07fffffff       -80000000     -00000001`},
	{[]byte{0, 0, 0, 0, 1, 0, 0, 0, 255, 127, 0, 0, 255, 255, 0, 0, 0, 128, 0, 0, 0, 0, 0, 0, 255, 255, 0, 0, 255, 255, 16, 113}, `tst:TRICE32_4 %10d ->              1     2147483647     -2147483648            -1`},
	{[]byte{34, 17, 0, 0, 68, 51, 0, 0, 102, 85, 0, 0, 136, 119, 105, 74}, `att:64bit 0b1000100100010001100110100010001010101011001100111011110001000`},
	{[]byte{239, 205, 171, 137, 152, 1, 0, 0, 5, 0, 69, 70}, `MSG: triceFifoMaxDepth = 408, select = 5`},
	{[]byte{255, 0, 51, 161}, `tst:TRICE8_1 -1`},
	{[]byte{254, 255, 201, 112}, `tst:TRICE8_2 -1 -2`},
	{[]byte{254, 255, 0, 0, 253, 0, 69, 72}, `tst:TRICE8_3 -1 -2 -3`},
	{[]byte{254, 255, 0, 0, 252, 253, 76, 136}, `tst:TRICE8_4 -1 -2 -3 -4`},
	{[]byte{254, 255, 0, 0, 252, 253, 0, 0, 251, 0, 251, 7}, `tst:TRICE8_5 -1 -2 -3 -4 -5`},
	{[]byte{254, 255, 0, 0, 252, 253, 0, 0, 250, 251, 146, 141}, `tst:TRICE8_6 -1 -2 -3 -4 -5 -6`},
	{[]byte{254, 255, 0, 0, 252, 253, 0, 0, 250, 251, 0, 0, 249, 0, 147, 90}, `tst:TRICE8_7 -1 -2 -3 -4 -5 -6 -7`},
	{[]byte{254, 255, 0, 0, 252, 253, 0, 0, 250, 251, 0, 0, 248, 249, 159, 20}, `tst:TRICE8_8 -1 -2 -3 -4 -5 -6 -7 -8`},
	{[]byte{239, 205, 171, 137, 152, 1, 0, 0, 6, 0, 69, 70}, `MSG: triceFifoMaxDepth = 408, select = 6`},
	{[]byte{239, 205, 171, 137, 152, 1, 0, 0, 7, 0, 69, 70}, `MSG: triceFifoMaxDepth = 408, select = 7`},
	{[]byte{255, 255, 33, 199}, `tst:TRICE16_1 -1`},
	{[]byte{255, 255, 0, 0, 254, 255, 112, 254}, `tst:TRICE16_2 -1 -2`},
	{[]byte{255, 255, 0, 0, 254, 255, 0, 0, 253, 255, 35, 230}, `tst:TRICE16_3 -1 -2 -3`},
	{[]byte{255, 255, 0, 0, 254, 255, 0, 0, 253, 255, 0, 0, 252, 255, 61, 176}, `tst:TRICE16_4 -1 -2 -3 -4`},
	{[]byte{239, 205, 171, 137, 152, 1, 0, 0, 8, 0, 69, 70}, `MSG: triceFifoMaxDepth = 408, select = 8`},
	{[]byte{255, 255, 0, 0, 255, 255, 240, 24}, `tst:TRICE32_1 -1`},
	{[]byte{255, 255, 0, 0, 255, 255, 240, 24}, `tst:TRICE32_1 -1`},
	{[]byte{255, 255, 0, 0, 255, 255, 0, 0, 255, 255, 0, 0, 254, 255, 241, 44}, `tst:TRICE32_2 -1 -2`},
	{[]byte{255, 255, 0, 0, 255, 255, 0, 0, 255, 255, 0, 0, 254, 255, 0, 0, 255, 255, 0, 0, 253, 255, 125, 219}, `tst:TRICE32_3 -1 -2 -3`},
	{[]byte{255, 255, 0, 0, 255, 255, 0, 0, 255, 255, 0, 0, 254, 255, 0, 0, 255, 255, 0, 0, 253, 255, 0, 0, 255, 255, 0, 0, 252, 255, 221, 142}, `tst:TRICE32_4 -1 -2 -3 -4`},
	{[]byte{239, 205, 171, 137, 152, 1, 0, 0, 9, 0, 69, 70}, `MSG: triceFifoMaxDepth = 408, select = 9`},
	{[]byte{255, 255, 0, 0, 255, 255, 0, 0, 255, 255, 0, 0, 255, 255, 163, 168}, `tst:TRICE64_1 -1`},
	{[]byte{255, 255, 0, 0, 255, 255, 0, 0, 255, 255, 0, 0, 255, 255, 0, 0, 255, 255, 0, 0, 255, 255, 0, 0, 255, 255, 0, 0, 254, 255, 210, 32}, `tst:TRICE64_2 -1 -2`},
	{[]byte{239, 205, 171, 137, 152, 1, 0, 0, 10, 0, 69, 70}, `MSG: triceFifoMaxDepth = 408, select = 10`},
	{[]byte{0, 0, 146, 174, 0, 0, 191, 5, 0, 0, 83, 205, 0, 0, 164, 240, 0, 0, 163, 154}, `e:Aw:Ba:cwr:drd:e`},
	{[]byte{0, 0, 188, 208, 0, 0, 166, 127, 0, 0, 19, 53, 0, 0, 123, 31, 0, 0, 245, 102, 0, 0, 252, 10}, `diag:fd:Gt:Htime:imessage:Jdbg:k`},
	{[]byte{239, 205, 171, 137, 152, 1, 0, 0, 11, 0, 69, 70}, `MSG: triceFifoMaxDepth = 408, select = 11`},
	{[]byte{0, 0, 6, 171, 0, 0, 201, 124, 0, 0, 33, 225, 0, 0, 117, 125, 0, 0, 156, 15, 0, 0, 159, 25, 0, 0, 15, 119}, `1234e:7m:12m:123`},
	{[]byte{239, 205, 171, 137, 152, 1, 0, 0, 12, 0, 69, 70}, `MSG: triceFifoMaxDepth = 408, select = 12`},
	{[]byte{1, 0, 51, 161}, `tst:TRICE8_1 1`},
	{[]byte{2, 1, 201, 112}, `tst:TRICE8_2 1 2`},
	{[]byte{2, 1, 0, 0, 3, 0, 69, 72}, `tst:TRICE8_3 1 2 3`},
	{[]byte{2, 1, 0, 0, 4, 3, 76, 136}, `tst:TRICE8_4 1 2 3 4`},
	{[]byte{2, 1, 0, 0, 4, 3, 0, 0, 5, 0, 251, 7}, `tst:TRICE8_5 1 2 3 4 5`},
	{[]byte{2, 1, 0, 0, 4, 3, 0, 0, 6, 5, 146, 141}, `tst:TRICE8_6 1 2 3 4 5 6`},
	{[]byte{2, 1, 0, 0, 4, 3, 0, 0, 6, 5, 0, 0, 7, 0, 147, 90}, `tst:TRICE8_7 1 2 3 4 5 6 7`},
	{[]byte{2, 1, 0, 0, 4, 3, 0, 0, 6, 5, 0, 0, 8, 7, 159, 20}, `tst:TRICE8_8 1 2 3 4 5 6 7 8`},
	{[]byte{239, 205, 171, 137, 152, 1, 0, 0, 13, 0, 69, 70}, `MSG: triceFifoMaxDepth = 408, select = 13`},
	{[]byte{110, 97, 0, 0, 101, 95, 0, 0, 97, 120, 0, 0, 112, 109, 26, 207, 101, 108, 0, 0, 115, 95, 0, 0, 114, 116, 0, 0, 110, 105, 26, 207, 10, 103, 198, 194}, `an_example_string`},
	{[]byte{10, 0, 56, 141}, ``},
	{[]byte{10, 97, 198, 194}, `a`},
	{[]byte{110, 97, 0, 0, 10, 0, 226, 237}, `an`},
	{[]byte{110, 97, 0, 0, 10, 95, 114, 226}, `an_`},
	{[]byte{110, 97, 0, 0, 101, 95, 0, 0, 10, 0, 149, 194}, `an_e`},
	{[]byte{110, 97, 0, 0, 101, 95, 0, 0, 10, 120, 217, 39}, `an_ex`},
	{[]byte{110, 97, 0, 0, 101, 95, 0, 0, 97, 120, 0, 0, 10, 0, 95, 224}, `an_exa`},
	{[]byte{110, 97, 0, 0, 101, 95, 0, 0, 97, 120, 0, 0, 10, 109, 26, 207}, `an_exam`},
	{[]byte{110, 97, 0, 0, 101, 95, 0, 0, 97, 120, 0, 0, 112, 109, 26, 207, 10, 0, 56, 141}, `an_examp`},
	{[]byte{110, 97, 0, 0, 101, 95, 0, 0, 97, 120, 0, 0, 112, 109, 26, 207, 10, 108, 198, 194}, `an_exampl`},
	{[]byte{110, 97, 0, 0, 101, 95, 0, 0, 97, 120, 0, 0, 112, 109, 26, 207, 101, 108, 0, 0, 10, 0, 226, 237}, `an_example`},
	{[]byte{110, 97, 0, 0, 101, 95, 0, 0, 97, 120, 0, 0, 112, 109, 26, 207, 101, 108, 0, 0, 10, 95, 114, 226}, `an_example_`},
	{[]byte{110, 97, 0, 0, 101, 95, 0, 0, 97, 120, 0, 0, 112, 109, 26, 207, 101, 108, 0, 0, 115, 95, 0, 0, 10, 0, 149, 194}, `an_example_s`},
	{[]byte{110, 97, 0, 0, 101, 95, 0, 0, 97, 120, 0, 0, 112, 109, 26, 207, 101, 108, 0, 0, 115, 95, 0, 0, 10, 116, 217, 39}, `an_example_st`},
	{[]byte{110, 97, 0, 0, 101, 95, 0, 0, 97, 120, 0, 0, 112, 109, 26, 207, 101, 108, 0, 0, 115, 95, 0, 0, 114, 116, 0, 0, 10, 0, 95, 224}, `an_example_str`},
	{[]byte{110, 97, 0, 0, 101, 95, 0, 0, 97, 120, 0, 0, 112, 109, 26, 207, 101, 108, 0, 0, 115, 95, 0, 0, 114, 116, 0, 0, 10, 105, 26, 207}, `an_example_stri`},
	{[]byte{110, 97, 0, 0, 101, 95, 0, 0, 97, 120, 0, 0, 112, 109, 26, 207, 101, 108, 0, 0, 115, 95, 0, 0, 114, 116, 0, 0, 110, 105, 26, 207, 10, 0, 56, 141}, `an_example_strin`},
	{[]byte{110, 97, 0, 0, 101, 95, 0, 0, 97, 120, 0, 0, 112, 109, 26, 207, 101, 108, 0, 0, 115, 95, 0, 0, 114, 116, 0, 0, 110, 105, 26, 207, 10, 103, 198, 194}, `an_example_string`},
	{[]byte{239, 205, 171, 137, 152, 1, 0, 0, 14, 0, 69, 70}, `MSG: triceFifoMaxDepth = 408, select = 14`},
}

var bareTestTable = testTable{
	{[]byte{137, 171, 205, 239, 0, 0, 1, 152, 70, 69, 0, 0}, `MSG: triceFifoMaxDepth = 408, select = 0`},
	{[]byte{95, 161, 0, 0}, `--------------------------------------------------`},
	{[]byte{52, 116, 0, 0}, `--------------------------------------------------`},
	{[]byte{70, 213, 48, 57}, `dbg:12345 as 16bit is 0b0011000000111001`},
	{[]byte{14, 207, 0, 0}, `--------------------------------------------------`},
	{[]byte{66, 229, 0, 0}, `sig:This ASSERT error is just a demo and no real error:`},
	{[]byte{104, 173, 0, 0}, `--------------------------------------------------`},
	{[]byte{137, 171, 205, 239, 0, 0, 1, 152, 70, 69, 0, 1}, `MSG: triceFifoMaxDepth = 408, select = 1`},
	{[]byte{70, 193, 180, 116}, `ERR:error       message, SysTick is -19340`},
	{[]byte{110, 129, 179, 22}, `WRN:warning     message, SysTick is -19690`},
	{[]byte{209, 56, 177, 189}, `ATT:attension   message, SysTick is -20035`},
	{[]byte{65, 32, 176, 99}, `DIA:diagnostics message, SysTick is -20381`},
	{[]byte{164, 222, 175, 10}, `TIM:timing      message, SysTick is -20726`},
	{[]byte{93, 165, 173, 176}, `DBG:debug       message, SysTick is -21072`},
	{[]byte{137, 171, 205, 239, 0, 0, 1, 152, 70, 69, 0, 2}, `MSG: triceFifoMaxDepth = 408, select = 2`},
	{[]byte{53, 162, 180, 201}, `SIG:signal      message, SysTick is -19255`},
	{[]byte{94, 238, 179, 103}, `RD:read         message, SysTick is -19609`},
	{[]byte{55, 58, 178, 13}, `WR:write        message, SysTick is -19955`},
	{[]byte{216, 149, 176, 180}, `ISR:interrupt   message, SysTick is -20300`},
	{[]byte{162, 23, 175, 90}, `TST:test        message, SysTick is -20646`},
	{[]byte{178, 129, 173, 3}, `MSG:normal      message, SysTick is -21245`},
	{[]byte{96, 13, 172, 169}, `INFO:informal   message, SysTick is -21335`},
	{[]byte{0, 0, 0, 0, 216, 128, 170, 48}, `TST:test        message, SysTick is  43568`},
	{[]byte{0, 0, 0, 0, 216, 128, 167, 183}, `TST:test        message, SysTick is  42935`},
	{[]byte{0, 0, 0, 0, 216, 128, 165, 61}, `TST:test        message, SysTick is  42301`},
	{[]byte{0, 0, 0, 0, 216, 128, 162, 200}, `TST:test        message, SysTick is  41672`},
	{[]byte{0, 0, 0, 0, 216, 128, 160, 82}, `TST:test        message, SysTick is  41042`},
	{[]byte{0, 0, 0, 0, 216, 128, 157, 221}, `TST:test        message, SysTick is  40413`},
	{[]byte{0, 0, 0, 0, 216, 128, 155, 103}, `TST:test        message, SysTick is  39783`},
	{[]byte{137, 171, 205, 239, 0, 0, 1, 152, 70, 69, 0, 3}, `MSG: triceFifoMaxDepth = 408, select = 3`},
	{[]byte{0, 0, 1, 127, 9, 223, 128, 255}, `tst:TRICE8  %03x ->  001  07f  -80  -01`},
	{[]byte{0, 0, 1, 127, 247, 191, 128, 255}, `tst:TRICE8   %4d ->    1  127 -128   -1`},
	{[]byte{0, 0, 1, 127, 156, 127, 128, 255}, `tst:TRICE8   %4o ->    1  177 -200   -1`},
	{[]byte{0, 0, 0, 1, 0, 0, 127, 255, 0, 0, 128, 0, 33, 226, 255, 255}, `tst:TRICE16  %05x ->   00001   07fff   -8000   -0001`},
	{[]byte{0, 0, 0, 1, 0, 0, 127, 255, 0, 0, 128, 0, 233, 221, 255, 255}, `tst:TRICE16   %6d ->       1   32767  -32768      -1`},
	{[]byte{0, 0, 0, 1, 0, 0, 127, 255, 0, 0, 128, 0, 96, 205, 255, 255}, `tst:TRICE16   %7o ->       1   77777 -100000      -1`},
	{[]byte{137, 171, 205, 239, 0, 0, 1, 152, 70, 69, 0, 4}, `MSG: triceFifoMaxDepth = 408, select = 4`},
	{[]byte{0, 0, 0, 0, 0, 0, 0, 1, 0, 0, 127, 255, 0, 0, 255, 255, 0, 0, 128, 0, 0, 0, 0, 0, 0, 0, 255, 255, 65, 171, 255, 255}, `tst:TRICE32_4 %09x ->      000000001      07fffffff       -80000000     -00000001`},
	{[]byte{0, 0, 0, 0, 0, 0, 0, 1, 0, 0, 127, 255, 0, 0, 255, 255, 0, 0, 128, 0, 0, 0, 0, 0, 0, 0, 255, 255, 113, 16, 255, 255}, `tst:TRICE32_4 %10d ->              1     2147483647     -2147483648            -1`},
	{[]byte{0, 0, 17, 34, 0, 0, 51, 68, 0, 0, 85, 102, 74, 105, 119, 136}, `att:64bit 0b1000100100010001100110100010001010101011001100111011110001000`},
	{[]byte{137, 171, 205, 239, 0, 0, 1, 152, 70, 69, 0, 5}, `MSG: triceFifoMaxDepth = 408, select = 5`},
	{[]byte{161, 51, 0, 255}, `tst:TRICE8_1 -1`},
	{[]byte{112, 201, 255, 254}, `tst:TRICE8_2 -1 -2`},
	{[]byte{0, 0, 255, 254, 72, 69, 0, 253}, `tst:TRICE8_3 -1 -2 -3`},
	{[]byte{0, 0, 255, 254, 136, 76, 253, 252}, `tst:TRICE8_4 -1 -2 -3 -4`},
	{[]byte{0, 0, 255, 254, 0, 0, 253, 252, 7, 251, 0, 251}, `tst:TRICE8_5 -1 -2 -3 -4 -5`},
	{[]byte{0, 0, 255, 254, 0, 0, 253, 252, 141, 146, 251, 250}, `tst:TRICE8_6 -1 -2 -3 -4 -5 -6`},
	{[]byte{0, 0, 255, 254, 0, 0, 253, 252, 0, 0, 251, 250, 90, 147, 0, 249}, `tst:TRICE8_7 -1 -2 -3 -4 -5 -6 -7`},
	{[]byte{0, 0, 255, 254, 0, 0, 253, 252, 0, 0, 251, 250, 20, 159, 249, 248}, `tst:TRICE8_8 -1 -2 -3 -4 -5 -6 -7 -8`},
	{[]byte{137, 171, 205, 239, 0, 0, 1, 152, 70, 69, 0, 6}, `MSG: triceFifoMaxDepth = 408, select = 6`},
	{[]byte{137, 171, 205, 239, 0, 0, 1, 152, 70, 69, 0, 7}, `MSG: triceFifoMaxDepth = 408, select = 7`},
	{[]byte{199, 33, 255, 255}, `tst:TRICE16_1 -1`},
	{[]byte{0, 0, 255, 255, 254, 112, 255, 254}, `tst:TRICE16_2 -1 -2`},
	{[]byte{0, 0, 255, 255, 0, 0, 255, 254, 230, 35, 255, 253}, `tst:TRICE16_3 -1 -2 -3`},
	{[]byte{0, 0, 255, 255, 0, 0, 255, 254, 0, 0, 255, 253, 176, 61, 255, 252}, `tst:TRICE16_4 -1 -2 -3 -4`},
	{[]byte{137, 171, 205, 239, 0, 0, 1, 152, 70, 69, 0, 8}, `MSG: triceFifoMaxDepth = 408, select = 8`},
	{[]byte{0, 0, 255, 255, 24, 240, 255, 255}, `tst:TRICE32_1 -1`},
	{[]byte{0, 0, 255, 255, 24, 240, 255, 255}, `tst:TRICE32_1 -1`},
	{[]byte{0, 0, 255, 255, 0, 0, 255, 255, 0, 0, 255, 255, 44, 241, 255, 254}, `tst:TRICE32_2 -1 -2`},
	{[]byte{0, 0, 255, 255, 0, 0, 255, 255, 0, 0, 255, 255, 0, 0, 255, 254, 0, 0, 255, 255, 219, 125, 255, 253}, `tst:TRICE32_3 -1 -2 -3`},
	{[]byte{0, 0, 255, 255, 0, 0, 255, 255, 0, 0, 255, 255, 0, 0, 255, 254, 0, 0, 255, 255, 0, 0, 255, 253, 0, 0, 255, 255, 142, 221, 255, 252}, `tst:TRICE32_4 -1 -2 -3 -4`},
	{[]byte{137, 171, 205, 239, 0, 0, 1, 152, 70, 69, 0, 9}, `MSG: triceFifoMaxDepth = 408, select = 9`},
	{[]byte{0, 0, 255, 255, 0, 0, 255, 255, 0, 0, 255, 255, 168, 163, 255, 255}, `tst:TRICE64_1 -1`},
	{[]byte{0, 0, 255, 255, 0, 0, 255, 255, 0, 0, 255, 255, 0, 0, 255, 255, 0, 0, 255, 255, 0, 0, 255, 255, 0, 0, 255, 255, 32, 210, 255, 254}, `tst:TRICE64_2 -1 -2`},
	{[]byte{137, 171, 205, 239, 0, 0, 1, 152, 70, 69, 0, 10}, `MSG: triceFifoMaxDepth = 408, select = 10`},
	{[]byte{174, 146, 0, 0, 5, 191, 0, 0, 205, 83, 0, 0, 240, 164, 0, 0, 154, 163, 0, 0}, `e:Aw:Ba:cwr:drd:e`},
	{[]byte{208, 188, 0, 0, 127, 166, 0, 0, 53, 19, 0, 0, 31, 123, 0, 0, 102, 245, 0, 0, 10, 252, 0, 0}, `diag:fd:Gt:Htime:imessage:Jdbg:k`},
	{[]byte{137, 171, 205, 239, 0, 0, 1, 152, 70, 69, 0, 11}, `MSG: triceFifoMaxDepth = 408, select = 11`},
	{[]byte{171, 6, 0, 0, 124, 201, 0, 0, 225, 33, 0, 0, 125, 117, 0, 0, 15, 156, 0, 0, 25, 159, 0, 0, 119, 15, 0, 0}, `1234e:7m:12m:123`},
	{[]byte{137, 171, 205, 239, 0, 0, 1, 152, 70, 69, 0, 12}, `MSG: triceFifoMaxDepth = 408, select = 12`},
	{[]byte{161, 51, 0, 1}, `tst:TRICE8_1 1`},
	{[]byte{112, 201, 1, 2}, `tst:TRICE8_2 1 2`},
	{[]byte{0, 0, 1, 2, 72, 69, 0, 3}, `tst:TRICE8_3 1 2 3`},
	{[]byte{0, 0, 1, 2, 136, 76, 3, 4}, `tst:TRICE8_4 1 2 3 4`},
	{[]byte{0, 0, 1, 2, 0, 0, 3, 4, 7, 251, 0, 5}, `tst:TRICE8_5 1 2 3 4 5`},
	{[]byte{0, 0, 1, 2, 0, 0, 3, 4, 141, 146, 5, 6}, `tst:TRICE8_6 1 2 3 4 5 6`},
	{[]byte{0, 0, 1, 2, 0, 0, 3, 4, 0, 0, 5, 6, 90, 147, 0, 7}, `tst:TRICE8_7 1 2 3 4 5 6 7`},
	{[]byte{0, 0, 1, 2, 0, 0, 3, 4, 0, 0, 5, 6, 20, 159, 7, 8}, `tst:TRICE8_8 1 2 3 4 5 6 7 8`},
	{[]byte{137, 171, 205, 239, 0, 0, 1, 152, 70, 69, 0, 13}, `MSG: triceFifoMaxDepth = 408, select = 13`},
	{[]byte{0, 0, 97, 110, 0, 0, 95, 101, 0, 0, 120, 97, 207, 26, 109, 112, 0, 0, 108, 101, 0, 0, 95, 115, 0, 0, 116, 114, 207, 26, 105, 110, 194, 198, 103, 10}, `an_example_string`},
	{[]byte{141, 56, 0, 10}, ``},
	{[]byte{194, 198, 97, 10}, `a`},
	{[]byte{0, 0, 97, 110, 237, 226, 0, 10}, `an`},
	{[]byte{0, 0, 97, 110, 226, 114, 95, 10}, `an_`},
	{[]byte{0, 0, 97, 110, 0, 0, 95, 101, 194, 149, 0, 10}, `an_e`},
	{[]byte{0, 0, 97, 110, 0, 0, 95, 101, 39, 217, 120, 10}, `an_ex`},
	{[]byte{0, 0, 97, 110, 0, 0, 95, 101, 0, 0, 120, 97, 224, 95, 0, 10}, `an_exa`},
	{[]byte{0, 0, 97, 110, 0, 0, 95, 101, 0, 0, 120, 97, 207, 26, 109, 10}, `an_exam`},
	{[]byte{0, 0, 97, 110, 0, 0, 95, 101, 0, 0, 120, 97, 207, 26, 109, 112, 141, 56, 0, 10}, `an_examp`},
	{[]byte{0, 0, 97, 110, 0, 0, 95, 101, 0, 0, 120, 97, 207, 26, 109, 112, 194, 198, 108, 10}, `an_exampl`},
	{[]byte{0, 0, 97, 110, 0, 0, 95, 101, 0, 0, 120, 97, 207, 26, 109, 112, 0, 0, 108, 101, 237, 226, 0, 10}, `an_example`},
	{[]byte{0, 0, 97, 110, 0, 0, 95, 101, 0, 0, 120, 97, 207, 26, 109, 112, 0, 0, 108, 101, 226, 114, 95, 10}, `an_example_`},
	{[]byte{0, 0, 97, 110, 0, 0, 95, 101, 0, 0, 120, 97, 207, 26, 109, 112, 0, 0, 108, 101, 0, 0, 95, 115, 194, 149, 0, 10}, `an_example_s`},
	{[]byte{0, 0, 97, 110, 0, 0, 95, 101, 0, 0, 120, 97, 207, 26, 109, 112, 0, 0, 108, 101, 0, 0, 95, 115, 39, 217, 116, 10}, `an_example_st`},
	{[]byte{0, 0, 97, 110, 0, 0, 95, 101, 0, 0, 120, 97, 207, 26, 109, 112, 0, 0, 108, 101, 0, 0, 95, 115, 0, 0, 116, 114, 224, 95, 0, 10}, `an_example_str`},
	{[]byte{0, 0, 97, 110, 0, 0, 95, 101, 0, 0, 120, 97, 207, 26, 109, 112, 0, 0, 108, 101, 0, 0, 95, 115, 0, 0, 116, 114, 207, 26, 105, 10}, `an_example_stri`},
	{[]byte{0, 0, 97, 110, 0, 0, 95, 101, 0, 0, 120, 97, 207, 26, 109, 112, 0, 0, 108, 101, 0, 0, 95, 115, 0, 0, 116, 114, 207, 26, 105, 110, 141, 56, 0, 10}, `an_example_strin`},
	{[]byte{0, 0, 97, 110, 0, 0, 95, 101, 0, 0, 120, 97, 207, 26, 109, 112, 0, 0, 108, 101, 0, 0, 95, 115, 0, 0, 116, 114, 207, 26, 105, 110, 194, 198, 103, 10}, `an_example_string`},
	{[]byte{137, 171, 205, 239, 0, 0, 1, 152, 70, 69, 0, 14}, `MSG: triceFifoMaxDepth = 408, select = 14`},
}

var bareLTestTable2 = testTable{
	{[]byte{239, 205, 171, 137, 152, 1, 0, 0, 0, 0, 69, 70}, `MSG: triceFifoMaxDepth = 408, select = 0`},
	{[]byte{0, 0, 161, 95}, `--------------------------------------------------`},
	{[]byte{0, 0, 116, 52}, `--------------------------------------------------`},
	{[]byte{57, 48, 213, 70}, `dbg:12345 as 16bit is 0b0011000000111001`},
	{[]byte{0, 0, 207, 14}, `--------------------------------------------------`},
	{[]byte{0, 0, 229, 66}, `sig:This ASSERT error is just a demo and no real error:`},
	{[]byte{0, 0, 173, 104}, `--------------------------------------------------`},
	{[]byte{239, 205, 171, 137, 152, 1, 0, 0, 1, 0, 69, 70}, `MSG: triceFifoMaxDepth = 408, select = 1`},
	{[]byte{109, 181, 193, 70}, `ERR:error       message, SysTick is -19091`},
	{[]byte{200, 180, 129, 110}, `WRN:warning     message, SysTick is -19256`},
	{[]byte{42, 180, 56, 209}, `ATT:attension   message, SysTick is -19414`},
	{[]byte{140, 179, 32, 65}, `DIA:diagnostics message, SysTick is -19572`},
	{[]byte{238, 178, 222, 164}, `TIM:timing      message, SysTick is -19730`},
	{[]byte{80, 178, 165, 93}, `DBG:debug       message, SysTick is -19888`},
	{[]byte{239, 205, 171, 137, 152, 1, 0, 0, 2, 0, 69, 70}, `MSG: triceFifoMaxDepth = 408, select = 2`},
	{[]byte{193, 181, 162, 53}, `SIG:signal      message, SysTick is -19007`},
	{[]byte{28, 181, 238, 94}, `RD:read         message, SysTick is -19172`},
	{[]byte{126, 180, 58, 55}, `WR:write        message, SysTick is -19330`},
	{[]byte{224, 179, 149, 216}, `ISR:interrupt   message, SysTick is -19488`},
	{[]byte{66, 179, 23, 162}, `TST:test        message, SysTick is -19646`},
	{[]byte{164, 178, 129, 178}, `MSG:normal      message, SysTick is -19804`},
	{[]byte{6, 178, 13, 96}, `INFO:informal   message, SysTick is -19962`},
	{[]byte{0, 0, 0, 0, 209, 176, 128, 216}, `TST:test        message, SysTick is  45265`},
	{[]byte{0, 0, 0, 0, 159, 175, 128, 216}, `TST:test        message, SysTick is  44959`},
	{[]byte{0, 0, 0, 0, 109, 174, 128, 216}, `TST:test        message, SysTick is  44653`},
	{[]byte{0, 0, 0, 0, 59, 173, 128, 216}, `TST:test        message, SysTick is  44347`},
	{[]byte{0, 0, 0, 0, 9, 172, 128, 216}, `TST:test        message, SysTick is  44041`},
	{[]byte{0, 0, 0, 0, 215, 170, 128, 216}, `TST:test        message, SysTick is  43735`},
	{[]byte{0, 0, 0, 0, 165, 169, 128, 216}, `TST:test        message, SysTick is  43429`},
	{[]byte{239, 205, 171, 137, 152, 1, 0, 0, 3, 0, 69, 70}, `MSG: triceFifoMaxDepth = 408, select = 3`},
	{[]byte{127, 1, 0, 0, 255, 128, 223, 9}, `tst:TRICE8  %03x ->  001  07f  -80  -01`},
	{[]byte{127, 1, 0, 0, 255, 128, 191, 247}, `tst:TRICE8   %4d ->    1  127 -128   -1`},
	{[]byte{127, 1, 0, 0, 255, 128, 127, 156}, `tst:TRICE8   %4o ->    1  177 -200   -1`},
	{[]byte{1, 0, 0, 0, 255, 127, 0, 0, 0, 128, 0, 0, 255, 255, 226, 33}, `tst:TRICE16  %05x ->   00001   07fff   -8000   -0001`},
	{[]byte{1, 0, 0, 0, 255, 127, 0, 0, 0, 128, 0, 0, 255, 255, 221, 233}, `tst:TRICE16   %6d ->       1   32767  -32768      -1`},
	{[]byte{1, 0, 0, 0, 255, 127, 0, 0, 0, 128, 0, 0, 255, 255, 205, 96}, `tst:TRICE16   %7o ->       1   77777 -100000      -1`},
	{[]byte{239, 205, 171, 137, 152, 1, 0, 0, 4, 0, 69, 70}, `MSG: triceFifoMaxDepth = 408, select = 4`},
	{[]byte{0, 0, 0, 0, 1, 0, 0, 0, 255, 127, 0, 0, 255, 255, 0, 0, 0, 128, 0, 0, 0, 0, 0, 0, 255, 255, 0, 0, 255, 255, 171, 65}, `tst:TRICE32_4 %09x ->      000000001      07fffffff       -80000000     -00000001`},
	{[]byte{0, 0, 0, 0, 1, 0, 0, 0, 255, 127, 0, 0, 255, 255, 0, 0, 0, 128, 0, 0, 0, 0, 0, 0, 255, 255, 0, 0, 255, 255, 16, 113}, `tst:TRICE32_4 %10d ->              1     2147483647     -2147483648            -1`},
	{[]byte{34, 17, 0, 0, 68, 51, 0, 0, 102, 85, 0, 0, 136, 119, 105, 74}, `att:64bit 0b1000100100010001100110100010001010101011001100111011110001000`},
	{[]byte{239, 205, 171, 137, 152, 1, 0, 0, 5, 0, 69, 70}, `MSG: triceFifoMaxDepth = 408, select = 5`},
	{[]byte{255, 0, 51, 161}, `tst:TRICE8_1 -1`},
	{[]byte{254, 255, 201, 112}, `tst:TRICE8_2 -1 -2`},
	{[]byte{254, 255, 0, 0, 253, 0, 69, 72}, `tst:TRICE8_3 -1 -2 -3`},
	{[]byte{254, 255, 0, 0, 252, 253, 76, 136}, `tst:TRICE8_4 -1 -2 -3 -4`},
	{[]byte{254, 255, 0, 0, 252, 253, 0, 0, 251, 0, 251, 7}, `tst:TRICE8_5 -1 -2 -3 -4 -5`},
	{[]byte{254, 255, 0, 0, 252, 253, 0, 0, 250, 251, 146, 141}, `tst:TRICE8_6 -1 -2 -3 -4 -5 -6`},
	{[]byte{254, 255, 0, 0, 252, 253, 0, 0, 250, 251, 0, 0, 249, 0, 147, 90}, `tst:TRICE8_7 -1 -2 -3 -4 -5 -6 -7`},
	{[]byte{254, 255, 0, 0, 252, 253, 0, 0, 250, 251, 0, 0, 248, 249, 159, 20}, `tst:TRICE8_8 -1 -2 -3 -4 -5 -6 -7 -8`},
	{[]byte{239, 205, 171, 137, 152, 1, 0, 0, 6, 0, 69, 70}, `MSG: triceFifoMaxDepth = 408, select = 6`},
	{[]byte{239, 205, 171, 137, 152, 1, 0, 0, 7, 0, 69, 70}, `MSG: triceFifoMaxDepth = 408, select = 7`},
	{[]byte{255, 255, 33, 199}, `tst:TRICE16_1 -1`},
	{[]byte{255, 255, 0, 0, 254, 255, 112, 254}, `tst:TRICE16_2 -1 -2`},
	{[]byte{255, 255, 0, 0, 254, 255, 0, 0, 253, 255, 35, 230}, `tst:TRICE16_3 -1 -2 -3`},
	{[]byte{255, 255, 0, 0, 254, 255, 0, 0, 253, 255, 0, 0, 252, 255, 61, 176}, `tst:TRICE16_4 -1 -2 -3 -4`},
	{[]byte{239, 205, 171, 137, 152, 1, 0, 0, 8, 0, 69, 70}, `MSG: triceFifoMaxDepth = 408, select = 8`},
	{[]byte{255, 255, 0, 0, 255, 255, 240, 24}, `tst:TRICE32_1 -1`},
	{[]byte{255, 255, 0, 0, 255, 255, 240, 24}, `tst:TRICE32_1 -1`},
	{[]byte{255, 255, 0, 0, 255, 255, 0, 0, 255, 255, 0, 0, 254, 255, 241, 44}, `tst:TRICE32_2 -1 -2`},
	{[]byte{255, 255, 0, 0, 255, 255, 0, 0, 255, 255, 0, 0, 254, 255, 0, 0, 255, 255, 0, 0, 253, 255, 125, 219}, `tst:TRICE32_3 -1 -2 -3`},
	{[]byte{255, 255, 0, 0, 255, 255, 0, 0, 255, 255, 0, 0, 254, 255, 0, 0, 255, 255, 0, 0, 253, 255, 0, 0, 255, 255, 0, 0, 252, 255, 221, 142}, `tst:TRICE32_4 -1 -2 -3 -4`},
	{[]byte{239, 205, 171, 137, 152, 1, 0, 0, 9, 0, 69, 70}, `MSG: triceFifoMaxDepth = 408, select = 9`},
	{[]byte{255, 255, 0, 0, 255, 255, 0, 0, 255, 255, 0, 0, 255, 255, 163, 168}, `tst:TRICE64_1 -1`},
	{[]byte{255, 255, 0, 0, 255, 255, 0, 0, 255, 255, 0, 0, 255, 255, 0, 0, 255, 255, 0, 0, 255, 255, 0, 0, 255, 255, 0, 0, 254, 255, 210, 32}, `tst:TRICE64_2 -1 -2`},
	{[]byte{239, 205, 171, 137, 152, 1, 0, 0, 10, 0, 69, 70}, `MSG: triceFifoMaxDepth = 408, select = 10`},
	{[]byte{0, 0, 146, 174, 0, 0, 191, 5, 0, 0, 83, 205, 0, 0, 164, 240, 0, 0, 163, 154}, `e:Aw:Ba:cwr:drd:e`},
	{[]byte{0, 0, 188, 208, 0, 0, 166, 127, 0, 0, 19, 53, 0, 0, 123, 31, 0, 0, 245, 102, 0, 0, 252, 10}, `diag:fd:Gt:Htime:imessage:Jdbg:k`},
	{[]byte{239, 205, 171, 137, 152, 1, 0, 0, 11, 0, 69, 70}, `MSG: triceFifoMaxDepth = 408, select = 11`},
	{[]byte{0, 0, 6, 171, 0, 0, 201, 124, 0, 0, 33, 225, 0, 0, 117, 125, 0, 0, 156, 15, 0, 0, 159, 25, 0, 0, 15, 119}, `1234e:7m:12m:123`},
	{[]byte{239, 205, 171, 137, 152, 1, 0, 0, 12, 0, 69, 70}, `MSG: triceFifoMaxDepth = 408, select = 12`},
	{[]byte{1, 0, 51, 161}, `tst:TRICE8_1 1`},
	{[]byte{2, 1, 201, 112}, `tst:TRICE8_2 1 2`},
	{[]byte{2, 1, 0, 0, 3, 0, 69, 72}, `tst:TRICE8_3 1 2 3`},
	{[]byte{2, 1, 0, 0, 4, 3, 76, 136}, `tst:TRICE8_4 1 2 3 4`},
	{[]byte{2, 1, 0, 0, 4, 3, 0, 0, 5, 0, 251, 7}, `tst:TRICE8_5 1 2 3 4 5`},
	{[]byte{2, 1, 0, 0, 4, 3, 0, 0, 6, 5, 146, 141}, `tst:TRICE8_6 1 2 3 4 5 6`},
	{[]byte{2, 1, 0, 0, 4, 3, 0, 0, 6, 5, 0, 0, 7, 0, 147, 90}, `tst:TRICE8_7 1 2 3 4 5 6 7`},
	{[]byte{2, 1, 0, 0, 4, 3, 0, 0, 6, 5, 0, 0, 8, 7, 159, 20}, `tst:TRICE8_8 1 2 3 4 5 6 7 8`},
	{[]byte{239, 205, 171, 137, 152, 1, 0, 0, 13, 0, 69, 70}, `MSG: triceFifoMaxDepth = 408, select = 13`},
	{[]byte{110, 97, 0, 0, 101, 95, 0, 0, 97, 120, 0, 0, 112, 109, 26, 207, 101, 108, 0, 0, 115, 95, 0, 0, 114, 116, 0, 0, 110, 105, 26, 207, 10, 103, 198, 194}, `an_example_string`},
	{[]byte{10, 0, 56, 141}, ``},
	{[]byte{10, 97, 198, 194}, `a`},
	{[]byte{110, 97, 0, 0, 10, 0, 226, 237}, `an`},
	{[]byte{110, 97, 0, 0, 10, 95, 114, 226}, `an_`},
	{[]byte{110, 97, 0, 0, 101, 95, 0, 0, 10, 0, 149, 194}, `an_e`},
	{[]byte{110, 97, 0, 0, 101, 95, 0, 0, 10, 120, 217, 39}, `an_ex`},
	{[]byte{110, 97, 0, 0, 101, 95, 0, 0, 97, 120, 0, 0, 10, 0, 95, 224}, `an_exa`},
	{[]byte{110, 97, 0, 0, 101, 95, 0, 0, 97, 120, 0, 0, 10, 109, 26, 207}, `an_exam`},
	{[]byte{110, 97, 0, 0, 101, 95, 0, 0, 97, 120, 0, 0, 112, 109, 26, 207, 10, 0, 56, 141}, `an_examp`},
	{[]byte{110, 97, 0, 0, 101, 95, 0, 0, 97, 120, 0, 0, 112, 109, 26, 207, 10, 108, 198, 194}, `an_exampl`},
	{[]byte{110, 97, 0, 0, 101, 95, 0, 0, 97, 120, 0, 0, 112, 109, 26, 207, 101, 108, 0, 0, 10, 0, 226, 237}, `an_example`},
	{[]byte{110, 97, 0, 0, 101, 95, 0, 0, 97, 120, 0, 0, 112, 109, 26, 207, 101, 108, 0, 0, 10, 95, 114, 226}, `an_example_`},
	{[]byte{110, 97, 0, 0, 101, 95, 0, 0, 97, 120, 0, 0, 112, 109, 26, 207, 101, 108, 0, 0, 115, 95, 0, 0, 10, 0, 149, 194}, `an_example_s`},
	{[]byte{110, 97, 0, 0, 101, 95, 0, 0, 97, 120, 0, 0, 112, 109, 26, 207, 101, 108, 0, 0, 115, 95, 0, 0, 10, 116, 217, 39}, `an_example_st`},
	{[]byte{110, 97, 0, 0, 101, 95, 0, 0, 97, 120, 0, 0, 112, 109, 26, 207, 101, 108, 0, 0, 115, 95, 0, 0, 114, 116, 0, 0, 10, 0, 95, 224}, `an_example_str`},
	{[]byte{110, 97, 0, 0, 101, 95, 0, 0, 97, 120, 0, 0, 112, 109, 26, 207, 101, 108, 0, 0, 115, 95, 0, 0, 114, 116, 0, 0, 10, 105, 26, 207}, `an_example_stri`},
	{[]byte{110, 97, 0, 0, 101, 95, 0, 0, 97, 120, 0, 0, 112, 109, 26, 207, 101, 108, 0, 0, 115, 95, 0, 0, 114, 116, 0, 0, 110, 105, 26, 207, 10, 0, 56, 141}, `an_example_strin`},
	{[]byte{110, 97, 0, 0, 101, 95, 0, 0, 97, 120, 0, 0, 112, 109, 26, 207, 101, 108, 0, 0, 115, 95, 0, 0, 114, 116, 0, 0, 110, 105, 26, 207, 10, 103, 198, 194}, `an_example_string`},
	{[]byte{239, 205, 171, 137, 152, 1, 0, 0, 14, 0, 69, 70}, `MSG: triceFifoMaxDepth = 408, select = 14`},
}

var bareTestTable2 = testTable{
	{[]byte{137, 171, 205, 239, 0, 0, 1, 152, 70, 69, 0, 0}, `MSG: triceFifoMaxDepth = 408, select = 0`},
	{[]byte{95, 161, 0, 0}, `--------------------------------------------------`},
	{[]byte{52, 116, 0, 0}, `--------------------------------------------------`},
	{[]byte{70, 213, 48, 57}, `dbg:12345 as 16bit is 0b0011000000111001`},
	{[]byte{14, 207, 0, 0}, `--------------------------------------------------`},
	{[]byte{66, 229, 0, 0}, `sig:This ASSERT error is just a demo and no real error:`},
	{[]byte{104, 173, 0, 0}, `--------------------------------------------------`},
	{[]byte{137, 171, 205, 239, 0, 0, 1, 152, 70, 69, 0, 1}, `MSG: triceFifoMaxDepth = 408, select = 1`},
	{[]byte{70, 193, 181, 10}, `ERR:error       message, SysTick is -19190`},
	{[]byte{110, 129, 180, 19}, `WRN:warning     message, SysTick is -19437`},
	{[]byte{209, 56, 179, 33}, `ATT:attension   message, SysTick is -19679`},
	{[]byte{65, 32, 178, 46}, `DIA:diagnostics message, SysTick is -19922`},
	{[]byte{164, 222, 177, 60}, `TIM:timing      message, SysTick is -20164`},
	{[]byte{93, 165, 176, 73}, `DBG:debug       message, SysTick is -20407`},
	{[]byte{137, 171, 205, 239, 0, 0, 1, 152, 70, 69, 0, 2}, `MSG: triceFifoMaxDepth = 408, select = 2`},
	{[]byte{53, 162, 181, 94}, `SIG:signal      message, SysTick is -19106`},
	{[]byte{94, 238, 180, 99}, `RD:read         message, SysTick is -19357`},
	{[]byte{55, 58, 179, 112}, `WR:write        message, SysTick is -19600`},
	{[]byte{216, 149, 178, 126}, `ISR:interrupt   message, SysTick is -19842`},
	{[]byte{162, 23, 177, 139}, `TST:test        message, SysTick is -20085`},
	{[]byte{178, 129, 176, 155}, `MSG:normal      message, SysTick is -20325`},
	{[]byte{96, 13, 175, 168}, `INFO:informal   message, SysTick is -20568`},
	{[]byte{0, 0, 0, 0, 216, 128, 173, 253}, `TST:test        message, SysTick is  44541`},
	{[]byte{0, 0, 0, 0, 216, 128, 172, 86}, `TST:test        message, SysTick is  44118`},
	{[]byte{0, 0, 0, 0, 216, 128, 170, 170}, `TST:test        message, SysTick is  43690`},
	{[]byte{0, 0, 0, 0, 216, 128, 168, 2}, `TST:test        message, SysTick is  43010`},
	{[]byte{0, 0, 0, 0, 216, 128, 167, 91}, `TST:test        message, SysTick is  42843`},
	{[]byte{0, 0, 0, 0, 216, 128, 165, 179}, `TST:test        message, SysTick is  42419`},
	{[]byte{0, 0, 0, 0, 216, 128, 164, 12}, `TST:test        message, SysTick is  41996`},
	{[]byte{137, 171, 205, 239, 0, 0, 1, 152, 70, 69, 0, 3}, `MSG: triceFifoMaxDepth = 408, select = 3`},
	{[]byte{0, 0, 1, 127, 9, 223, 128, 255}, `tst:TRICE8  %03x ->  001  07f  -80  -01`},
	{[]byte{0, 0, 1, 127, 247, 191, 128, 255}, `tst:TRICE8   %4d ->    1  127 -128   -1`},
	{[]byte{0, 0, 1, 127, 156, 127, 128, 255}, `tst:TRICE8   %4o ->    1  177 -200   -1`},
	{[]byte{0, 0, 0, 1, 0, 0, 127, 255, 0, 0, 128, 0, 33, 226, 255, 255}, `tst:TRICE16  %05x ->   00001   07fff   -8000   -0001`},
	{[]byte{0, 0, 0, 1, 0, 0, 127, 255, 0, 0, 128, 0, 233, 221, 255, 255}, `tst:TRICE16   %6d ->       1   32767  -32768      -1`},
	{[]byte{0, 0, 0, 1, 0, 0, 127, 255, 0, 0, 128, 0, 96, 205, 255, 255}, `tst:TRICE16   %7o ->       1   77777 -100000      -1`},
	{[]byte{137, 171, 205, 239, 0, 0, 1, 152, 70, 69, 0, 4}, `MSG: triceFifoMaxDepth = 408, select = 4`},
	{[]byte{0, 0, 0, 0, 0, 0, 0, 1, 0, 0, 127, 255, 0, 0, 255, 255, 0, 0, 128, 0, 0, 0, 0, 0, 0, 0, 255, 255, 65, 171, 255, 255}, `tst:TRICE32_4 %09x ->      000000001      07fffffff       -80000000     -00000001`},
	{[]byte{0, 0, 0, 0, 0, 0, 0, 1, 0, 0, 127, 255, 0, 0, 255, 255, 0, 0, 128, 0, 0, 0, 0, 0, 0, 0, 255, 255, 113, 16, 255, 255}, `tst:TRICE32_4 %10d ->              1     2147483647     -2147483648            -1`},
	{[]byte{0, 0, 17, 34, 0, 0, 51, 68, 0, 0, 85, 102, 74, 105, 119, 136}, `att:64bit 0b1000100100010001100110100010001010101011001100111011110001000`},
	{[]byte{137, 171, 205, 239, 0, 0, 1, 152, 70, 69, 0, 5}, `MSG: triceFifoMaxDepth = 408, select = 5`},
	{[]byte{161, 51, 0, 255}, `tst:TRICE8_1 -1`},
	{[]byte{112, 201, 255, 254}, `tst:TRICE8_2 -1 -2`},
	{[]byte{0, 0, 255, 254, 72, 69, 0, 253}, `tst:TRICE8_3 -1 -2 -3`},
	{[]byte{0, 0, 255, 254, 136, 76, 253, 252}, `tst:TRICE8_4 -1 -2 -3 -4`},
	{[]byte{0, 0, 255, 254, 0, 0, 253, 252, 7, 251, 0, 251}, `tst:TRICE8_5 -1 -2 -3 -4 -5`},
	{[]byte{0, 0, 255, 254, 0, 0, 253, 252, 141, 146, 251, 250}, `tst:TRICE8_6 -1 -2 -3 -4 -5 -6`},
	{[]byte{0, 0, 255, 254, 0, 0, 253, 252, 0, 0, 251, 250, 90, 147, 0, 249}, `tst:TRICE8_7 -1 -2 -3 -4 -5 -6 -7`},
	{[]byte{0, 0, 255, 254, 0, 0, 253, 252, 0, 0, 251, 250, 20, 159, 249, 248}, `tst:TRICE8_8 -1 -2 -3 -4 -5 -6 -7 -8`},
	{[]byte{137, 171, 205, 239, 0, 0, 1, 152, 70, 69, 0, 6}, `MSG: triceFifoMaxDepth = 408, select = 6`},
	{[]byte{137, 171, 205, 239, 0, 0, 1, 152, 70, 69, 0, 7}, `MSG: triceFifoMaxDepth = 408, select = 7`},
	{[]byte{199, 33, 255, 255}, `tst:TRICE16_1 -1`},
	{[]byte{0, 0, 255, 255, 254, 112, 255, 254}, `tst:TRICE16_2 -1 -2`},
	{[]byte{0, 0, 255, 255, 0, 0, 255, 254, 230, 35, 255, 253}, `tst:TRICE16_3 -1 -2 -3`},
	{[]byte{0, 0, 255, 255, 0, 0, 255, 254, 0, 0, 255, 253, 176, 61, 255, 252}, `tst:TRICE16_4 -1 -2 -3 -4`},
	{[]byte{137, 171, 205, 239, 0, 0, 1, 152, 70, 69, 0, 8}, `MSG: triceFifoMaxDepth = 408, select = 8`},
	{[]byte{0, 0, 255, 255, 24, 240, 255, 255}, `tst:TRICE32_1 -1`},
	{[]byte{0, 0, 255, 255, 24, 240, 255, 255}, `tst:TRICE32_1 -1`},
	{[]byte{0, 0, 255, 255, 0, 0, 255, 255, 0, 0, 255, 255, 44, 241, 255, 254}, `tst:TRICE32_2 -1 -2`},
	{[]byte{0, 0, 255, 255, 0, 0, 255, 255, 0, 0, 255, 255, 0, 0, 255, 254, 0, 0, 255, 255, 219, 125, 255, 253}, `tst:TRICE32_3 -1 -2 -3`},
	{[]byte{0, 0, 255, 255, 0, 0, 255, 255, 0, 0, 255, 255, 0, 0, 255, 254, 0, 0, 255, 255, 0, 0, 255, 253, 0, 0, 255, 255, 142, 221, 255, 252}, `tst:TRICE32_4 -1 -2 -3 -4`},
	{[]byte{137, 171, 205, 239, 0, 0, 1, 152, 70, 69, 0, 9}, `MSG: triceFifoMaxDepth = 408, select = 9`},
	{[]byte{0, 0, 255, 255, 0, 0, 255, 255, 0, 0, 255, 255, 168, 163, 255, 255}, `tst:TRICE64_1 -1`},
	{[]byte{0, 0, 255, 255, 0, 0, 255, 255, 0, 0, 255, 255, 0, 0, 255, 255, 0, 0, 255, 255, 0, 0, 255, 255, 0, 0, 255, 255, 32, 210, 255, 254}, `tst:TRICE64_2 -1 -2`},
	{[]byte{137, 171, 205, 239, 0, 0, 1, 152, 70, 69, 0, 10}, `MSG: triceFifoMaxDepth = 408, select = 10`},
	{[]byte{174, 146, 0, 0, 5, 191, 0, 0, 205, 83, 0, 0, 240, 164, 0, 0, 154, 163, 0, 0}, `e:Aw:Ba:cwr:drd:e`},
	{[]byte{208, 188, 0, 0, 127, 166, 0, 0, 53, 19, 0, 0, 31, 123, 0, 0, 102, 245, 0, 0, 10, 252, 0, 0}, `diag:fd:Gt:Htime:imessage:Jdbg:k`},
	{[]byte{137, 171, 205, 239, 0, 0, 1, 152, 70, 69, 0, 11}, `MSG: triceFifoMaxDepth = 408, select = 11`},
	{[]byte{171, 6, 0, 0, 124, 201, 0, 0, 225, 33, 0, 0, 125, 117, 0, 0, 15, 156, 0, 0, 25, 159, 0, 0, 119, 15, 0, 0}, `1234e:7m:12m:123`},
	{[]byte{137, 171, 205, 239, 0, 0, 1, 152, 70, 69, 0, 12}, `MSG: triceFifoMaxDepth = 408, select = 12`},
	{[]byte{161, 51, 0, 1}, `tst:TRICE8_1 1`},
	{[]byte{112, 201, 1, 2}, `tst:TRICE8_2 1 2`},
	{[]byte{0, 0, 1, 2, 72, 69, 0, 3}, `tst:TRICE8_3 1 2 3`},
	{[]byte{0, 0, 1, 2, 136, 76, 3, 4}, `tst:TRICE8_4 1 2 3 4`},
	{[]byte{0, 0, 1, 2, 0, 0, 3, 4, 7, 251, 0, 5}, `tst:TRICE8_5 1 2 3 4 5`},
	{[]byte{0, 0, 1, 2, 0, 0, 3, 4, 141, 146, 5, 6}, `tst:TRICE8_6 1 2 3 4 5 6`},
	{[]byte{0, 0, 1, 2, 0, 0, 3, 4, 0, 0, 5, 6, 90, 147, 0, 7}, `tst:TRICE8_7 1 2 3 4 5 6 7`},
	{[]byte{0, 0, 1, 2, 0, 0, 3, 4, 0, 0, 5, 6, 20, 159, 7, 8}, `tst:TRICE8_8 1 2 3 4 5 6 7 8`},
	{[]byte{137, 171, 205, 239, 0, 0, 1, 152, 70, 69, 0, 13}, `MSG: triceFifoMaxDepth = 408, select = 13`},
	{[]byte{0, 0, 97, 110, 0, 0, 95, 101, 0, 0, 120, 97, 207, 26, 109, 112, 0, 0, 108, 101, 0, 0, 95, 115, 0, 0, 116, 114, 207, 26, 105, 110, 194, 198, 103, 10}, `an_example_string`},
	{[]byte{141, 56, 0, 10}, ``},
	{[]byte{194, 198, 97, 10}, `a`},
	{[]byte{0, 0, 97, 110, 237, 226, 0, 10}, `an`},
	{[]byte{0, 0, 97, 110, 226, 114, 95, 10}, `an_`},
	{[]byte{0, 0, 97, 110, 0, 0, 95, 101, 194, 149, 0, 10}, `an_e`},
	{[]byte{0, 0, 97, 110, 0, 0, 95, 101, 39, 217, 120, 10}, `an_ex`},
	{[]byte{0, 0, 97, 110, 0, 0, 95, 101, 0, 0, 120, 97, 224, 95, 0, 10}, `an_exa`},
	{[]byte{0, 0, 97, 110, 0, 0, 95, 101, 0, 0, 120, 97, 207, 26, 109, 10}, `an_exam`},
	{[]byte{0, 0, 97, 110, 0, 0, 95, 101, 0, 0, 120, 97, 207, 26, 109, 112, 141, 56, 0, 10}, `an_examp`},
	{[]byte{0, 0, 97, 110, 0, 0, 95, 101, 0, 0, 120, 97, 207, 26, 109, 112, 194, 198, 108, 10}, `an_exampl`},
	{[]byte{0, 0, 97, 110, 0, 0, 95, 101, 0, 0, 120, 97, 207, 26, 109, 112, 0, 0, 108, 101, 237, 226, 0, 10}, `an_example`},
	{[]byte{0, 0, 97, 110, 0, 0, 95, 101, 0, 0, 120, 97, 207, 26, 109, 112, 0, 0, 108, 101, 226, 114, 95, 10}, `an_example_`},
	{[]byte{0, 0, 97, 110, 0, 0, 95, 101, 0, 0, 120, 97, 207, 26, 109, 112, 0, 0, 108, 101, 0, 0, 95, 115, 194, 149, 0, 10}, `an_example_s`},
	{[]byte{0, 0, 97, 110, 0, 0, 95, 101, 0, 0, 120, 97, 207, 26, 109, 112, 0, 0, 108, 101, 0, 0, 95, 115, 39, 217, 116, 10}, `an_example_st`},
	{[]byte{0, 0, 97, 110, 0, 0, 95, 101, 0, 0, 120, 97, 207, 26, 109, 112, 0, 0, 108, 101, 0, 0, 95, 115, 0, 0, 116, 114, 224, 95, 0, 10}, `an_example_str`},
	{[]byte{0, 0, 97, 110, 0, 0, 95, 101, 0, 0, 120, 97, 207, 26, 109, 112, 0, 0, 108, 101, 0, 0, 95, 115, 0, 0, 116, 114, 207, 26, 105, 10}, `an_example_stri`},
	{[]byte{0, 0, 97, 110, 0, 0, 95, 101, 0, 0, 120, 97, 207, 26, 109, 112, 0, 0, 108, 101, 0, 0, 95, 115, 0, 0, 116, 114, 207, 26, 105, 110, 141, 56, 0, 10}, `an_example_strin`},
	{[]byte{0, 0, 97, 110, 0, 0, 95, 101, 0, 0, 120, 97, 207, 26, 109, 112, 0, 0, 108, 101, 0, 0, 95, 115, 0, 0, 116, 114, 207, 26, 105, 110, 194, 198, 103, 10}, `an_example_string`},
	{[]byte{137, 171, 205, 239, 0, 0, 1, 152, 70, 69, 0, 14}, `MSG: triceFifoMaxDepth = 408, select = 14`},
}