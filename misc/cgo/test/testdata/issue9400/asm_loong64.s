// Copyright 2021 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// +build loong64
// +build !gccgo

#include "textflag.h"

#define DBAR	WORD $0x38720000

TEXT ·RewindAndSetgid(SB),NOSPLIT|NOFRAME,$0-0
	// Rewind stack pointer so anything that happens on the stack
	// will clobber the test pattern created by the caller
	ADDV	$(1024*8), R3

	// Ask signaller to setgid
	MOVW	$1, R12
	DBAR
	MOVW	R12, ·Baton(SB)
	DBAR

	// Wait for setgid completion
loop:
	DBAR
	MOVW	·Baton(SB), R12
	OR	R13, R13, R13	// hint that we're in a spin loop
	BNE	R12, loop
	DBAR

	// Restore stack
	ADDV	$(-1024*8), R3
	RET
