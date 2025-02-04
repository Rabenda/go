// Copyright 2021 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// This file encapsulates some of the odd characteristics of the
// LOONG64 (LOONG64) instruction set, to minimize its interaction
// with the core of the assembler.

package arch

import (
	"cmd/internal/obj"
	"cmd/internal/obj/loong64"
)

func jumpLOONG64(word string) bool {
	switch word {
	case "BEQ", "BFPF", "BFPT", "BLTZ", "BGEZ", "BLEZ", "BGTZ", "BLT", "BLTU", "JIRL", "BNE", "BGE", "BGEU", "JMP", "JAL", "CALL":
		return true
	}
	return false
}

// IsLOONG64CMP reports whether the op (as defined by an loong64.A* constant) is
// one of the CMP instructions that require special handling.
func IsLOONG64CMP(op obj.As) bool {
	switch op {
	case loong64.ACMPEQF, loong64.ACMPEQD, loong64.ACMPGEF, loong64.ACMPGED,
		loong64.ACMPGTF, loong64.ACMPGTD:
		return true
	}
	return false
}

// IsLOONG64MUL reports whether the op (as defined by an loong64.A* constant) is
// one of the MUL/DIV/REM/MADD/MSUB instructions that require special handling.
func IsLOONG64MUL(op obj.As) bool {
	switch op {
	case loong64.AMUL, loong64.AMULU, loong64.AMULV, loong64.AMULVU,
		loong64.ADIV, loong64.ADIVU, loong64.ADIVV, loong64.ADIVVU,
		loong64.AREM, loong64.AREMU, loong64.AREMV, loong64.AREMVU:
		return true
	}
	return false
}

func loong64RegisterNumber(name string, n int16) (int16, bool) {
	switch name {
	case "F":
		if 0 <= n && n <= 31 {
			return loong64.REG_F0 + n, true
		}
	case "FCSR":
		if 0 <= n && n <= 31 {
			return loong64.REG_FCSR0 + n, true
		}
	case "FCC":
		if 0 <= n && n <= 31 {
			return loong64.REG_FCC0 + n, true
		}
	case "R":
		if 0 <= n && n <= 31 {
			return loong64.REG_R0 + n, true
		}
		//MSA
		//	case "W":
		//		if 0 <= n && n <= 31 {
		//			return loong64.REG_W0 + n, true
		//		}
	}
	return 0, false
}
