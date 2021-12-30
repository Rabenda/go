// Copyright 2021 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// +build loong64

#include "textflag.h"

/*
 * void crosscall2(void (*fn)(void*, int32, uintptr), void*, int32, uintptr)
 * Save registers and call fn with three arguments.
 */
TEXT crosscall2(SB),NOSPLIT|NOFRAME,$0
	/*
	 * We still need to save all callee save register as before, and then
	 *  push 3 args for fn (R5, R6, R7).
	 * Also note that at procedure entry in gc world, 8(R29) will be the
	 *  first arg.
	 */

	ADDV	$(-8*23), R3
	MOVV	R5, (8*1)(R3) // void*
	MOVW	R6, (8*2)(R3) // int32
	MOVV	R7, (8*3)(R3) // uintptr
	MOVV	R23, (8*4)(R3)
	MOVV	R24, (8*5)(R3)
	MOVV	R25, (8*6)(R3)
	MOVV	R26, (8*7)(R3)
	MOVV	R27, (8*8)(R3)
	MOVV	R28, (8*9)(R3)
	MOVV	R29, (8*10)(R3)
	MOVV	R30, (8*11)(R3)
	MOVV	RSB, (8*12)(R3)
	MOVV	g, (8*13)(R3)
	MOVV	R1, (8*14)(R3)
	MOVD	F24, (8*15)(R3)
	MOVD	F25, (8*16)(R3)
	MOVD	F26, (8*17)(R3)
	MOVD	F27, (8*18)(R3)
	MOVD	F28, (8*19)(R3)
	MOVD	F29, (8*20)(R3)
	MOVD	F30, (8*21)(R3)
	MOVD	F31, (8*22)(R3)

	// Initialize Go ABI environment
	// prepare SB register = PC & 0xffffffff00000000
	//JAL	1(PC)
	//SRLV	$32, R1, RSB
	//SLLV	$32, RSB
	JAL	runtime·load_g(SB)
	JAL	(R4)

	MOVV	(8*4)(R3), R23
	MOVV	(8*5)(R3), R24
	MOVV	(8*6)(R3), R25
	MOVV	(8*7)(R3), R26
	MOVV	(8*8)(R3), R27
	MOVV	(8*9)(R3), R28
	MOVV	(8*10)(R3), R29
	MOVV	(8*11)(R3), R30
	MOVV	(8*12)(R3), RSB
	MOVV	(8*13)(R3), g
	MOVV	(8*14)(R3), R1
	MOVD	(8*15)(R3), F24
	MOVD	(8*16)(R3), F25
	MOVD	(8*17)(R3), F26
	MOVD	(8*18)(R3), F27
	MOVD	(8*19)(R3), F28
	MOVD	(8*20)(R3), F29
	MOVD	(8*21)(R3), F30
	MOVD	(8*22)(R3), F31
	ADDV	$(8*23), R3

	RET
