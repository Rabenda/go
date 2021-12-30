// Copyright 2021 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// +build loong64

#include "go_asm.h"
#include "textflag.h"

// func memclrNoHeapPointers(ptr unsafe.Pointer, n uintptr)
TEXT runtime·memclrNoHeapPointers(SB),NOSPLIT,$0-16
	MOVV	ptr+0(FP), R6
	MOVV	n+8(FP), R7
	ADDV	R6, R7, R4

	// if less than 16 bytes or no MSA, do words check
//	SGTU	$16, R2, R3
//	BNE	R3, no_msa
//	MOVBU	internal∕cpu·LOONG64X+const_offsetLOONG64XHasMSA(SB), R3
//	BEQ	R3, R0, no_msa

//	VMOVB	$0, W0

//	SGTU	$128, R2, R3
//	BEQ	R3, msa_large

//	AND	$15, R2, R5
//	XOR	R2, R5, R6
//	ADDVU	R1, R6

//msa_small:
//	VMOVB	W0, (R1)
//	ADDVU	$16, R1
//	SGTU	R6, R1, R3
//	BNE	R3, R0, msa_small
//	BEQ	R5, R0, done
//	VMOVB	W0, -16(R4)
//	JMP	done

//msa_large:
//	AND	$127, R2, R5
//	XOR	R2, R5, R6
//	ADDVU	R1, R6

//msa_large_loop:
//	VMOVB	W0, (R1)
//	VMOVB	W0, 16(R1)
//	VMOVB	W0, 32(R1)
//	VMOVB	W0, 48(R1)
//	VMOVB	W0, 64(R1)
//	VMOVB	W0, 80(R1)
//	VMOVB	W0, 96(R1)
//	VMOVB	W0, 112(R1)

//	ADDVU	$128, R1
//	SGTU	R6, R1, R3
//	BNE	R3, R0, msa_large_loop
//	BEQ	R5, R0, done
//	VMOVB	W0, -128(R4)
//	VMOVB	W0, -112(R4)
//	VMOVB	W0, -96(R4)
//	VMOVB	W0, -80(R4)
//	VMOVB	W0, -64(R4)
//	VMOVB	W0, -48(R4)
//	VMOVB	W0, -32(R4)
//	VMOVB	W0, -16(R4)
//	JMP	done

//no_msa:
	// if less than 8 bytes, do one byte at a time
	SGTU	$8, R7, R8
	BNE	R8, out

	// do one byte at a time until 8-aligned
	AND	$7, R6, R8
	BEQ	R8, words
	MOVB	R0, (R6)
	ADDV	$1, R6
	JMP	-4(PC)

words:
	// do 8 bytes at a time if there is room
	ADDV	$-7, R4, R7

	SGTU	R7, R6, R8
	BEQ	R8, out
	MOVV	R0, (R6)
	ADDV	$8, R6
	JMP	-4(PC)

out:
	BEQ	R6, R4, done
	MOVB	R0, (R6)
	ADDV	$1, R6
	JMP	-3(PC)
done:
	RET
