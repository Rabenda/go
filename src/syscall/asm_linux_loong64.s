// Copyright 2021 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// +build linux
// +build loong64

#include "textflag.h"

//
// System calls for loong64, Linux
//

// func Syscall(trap int64, a1, a2, a3 int64) (r1, r2, err int64);
// Trap # in R11, args in R4-R9, return in R4
TEXT ·Syscall(SB),NOSPLIT,$0-56
	JAL	runtime·entersyscall(SB)
	MOVV	a1+8(FP), R4
	MOVV	a2+16(FP), R5
	MOVV	a3+24(FP), R6
	MOVV	R0, R7
	MOVV	R0, R8
	MOVV	R0, R9
	MOVV	trap+0(FP), R11	// syscall entry
	SYSCALL
	MOVW	$-4096, R12
	BGEU	R12, R4, ok
	MOVV	$-1, R12
	MOVV	R12, r1+32(FP)	// r1
	MOVV	R0, r2+40(FP)	// r2
	SUBVU	R4, R0, R4
	MOVV	R4, err+48(FP)	// errno
	JAL	runtime·exitsyscall(SB)
	RET
ok:
	MOVV	R4, r1+32(FP)	// r1
	MOVV	R5, r2+40(FP)	// r2
	MOVV	R0, err+48(FP)	// errno
	JAL	runtime·exitsyscall(SB)
	RET

TEXT ·Syscall6(SB),NOSPLIT,$0-80
	JAL	runtime·entersyscall(SB)
	MOVV	a1+8(FP), R4
	MOVV	a2+16(FP), R5
	MOVV	a3+24(FP), R6
	MOVV	a4+32(FP), R7
	MOVV	a5+40(FP), R8
	MOVV	a6+48(FP), R9
	MOVV	trap+0(FP), R11	// syscall entry
	SYSCALL
	MOVW	$-4096, R12
	BGEU	R12, R4, ok
	MOVV	$-1, R12
	MOVV	R12, r1+56(FP)	// r1
	MOVV	R0, r2+64(FP)	// r2
	SUBVU	R4, R0, R4
	MOVV	R4, err+72(FP)	// errno
	JAL	runtime·exitsyscall(SB)
	RET
ok:
	MOVV	R4, r1+56(FP)	// r1
	MOVV	R5, r2+64(FP)	// r2
	MOVV	R0, err+72(FP)	// errno
	JAL	runtime·exitsyscall(SB)
	RET

TEXT ·RawSyscall(SB),NOSPLIT,$0-56
	MOVV	a1+8(FP), R4
	MOVV	a2+16(FP), R5
	MOVV	a3+24(FP), R6
	MOVV	R0, R7
	MOVV	R0, R8
	MOVV	R0, R9
	MOVV	trap+0(FP), R11	// syscall entry
	SYSCALL
	MOVW	$-4096, R12
	BGEU	R12, R4, ok
	MOVV	$-1, R12
	MOVV	R12, r1+32(FP)	// r1
	MOVV	R0, r2+40(FP)	// r2
	SUBVU	R4, R0, R4
	MOVV	R4, err+48(FP)	// errno
	RET
ok:
	MOVV	R4, r1+32(FP)	// r1
	MOVV	R5, r2+40(FP)	// r2
	MOVV	R0, err+48(FP)	// errno
	RET

TEXT ·RawSyscall6(SB),NOSPLIT,$0-80
	MOVV	a1+8(FP), R4
	MOVV	a2+16(FP), R5
	MOVV	a3+24(FP), R6
	MOVV	a4+32(FP), R7
	MOVV	a5+40(FP), R8
	MOVV	a6+48(FP), R9
	MOVV	trap+0(FP), R11	// syscall entry
	SYSCALL
	MOVW	$-4096, R12
	BGEU	R12, R4, ok
	MOVV	$-1, R12
	MOVV	R12, r1+56(FP)	// r1
	MOVV	R0, r2+64(FP)	// r2
	SUBVU	R4, R0, R4
	MOVV	R4, err+72(FP)	// errno
	RET
ok:
	MOVV	R4, r1+56(FP)	// r1
	MOVV	R5, r2+64(FP)	// r2
	MOVV	R0, err+72(FP)	// errno
	RET

// func rawVforkSyscall(trap, a1 uintptr) (r1, err uintptr)
TEXT ·rawVforkSyscall(SB),NOSPLIT,$0-32
	MOVV	a1+8(FP), R4
	MOVV	$0, R5
	MOVV	$0, R6
	MOVV	$0, R7
	MOVV	$0, R8
	MOVV	$0, R9
	MOVV	trap+0(FP), R11	// syscall entry
	SYSCALL
	MOVW	$-4096, R12
	BGEU	R12, R4, ok
	MOVV	$-1, R12
	MOVV	R12, r1+16(FP)		// r1
	SUBVU	R4, R0, R4
	MOVV	R4, err+24(FP)		// errno
	RET
ok:
	MOVV	R4, r1+16(FP)	// r1
	MOVV	R0, err+24(FP)	// errno
	RET

TEXT ·rawSyscallNoError(SB),NOSPLIT,$0-48
	MOVV	a1+8(FP), R4
	MOVV	a2+16(FP), R5
	MOVV	a3+24(FP), R6
	MOVV	R0, R7
	MOVV	R0, R8
	MOVV	R0, R9
	MOVV	trap+0(FP), R11	// syscall entry
	SYSCALL
	MOVV	R4, r1+32(FP)
	MOVV	R5, r2+40(FP)
	RET
