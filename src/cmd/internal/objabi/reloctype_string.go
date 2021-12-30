// Code generated by "stringer -type=RelocType"; DO NOT EDIT.

package objabi

import "strconv"

func _() {
	// An "invalid array index" compiler error signifies that the constant values have changed.
	// Re-run the stringer command to generate them again.
	var x [1]struct{}
	_ = x[R_ADDR-1]
	_ = x[R_ADDRPOWER-2]
	_ = x[R_ADDRARM64-3]
	_ = x[R_ADDRMIPS-4]
	_ = x[R_ADDRLOONG64-5]
	_ = x[R_ADDROFF-6]
	_ = x[R_SIZE-7]
	_ = x[R_CALL-8]
	_ = x[R_CALLARM-9]
	_ = x[R_CALLARM64-10]
	_ = x[R_CALLIND-11]
	_ = x[R_CALLPOWER-12]
	_ = x[R_CALLMIPS-13]
	_ = x[R_CALLLOONG64-14]
	_ = x[R_CONST-15]
	_ = x[R_PCREL-16]
	_ = x[R_TLS_LE-17]
	_ = x[R_TLS_IE-18]
	_ = x[R_GOTOFF-19]
	_ = x[R_PLT0-20]
	_ = x[R_PLT1-21]
	_ = x[R_PLT2-22]
	_ = x[R_USEFIELD-23]
	_ = x[R_USETYPE-24]
	_ = x[R_USEIFACE-25]
	_ = x[R_USEIFACEMETHOD-26]
	_ = x[R_USEGENERICIFACEMETHOD-27]
	_ = x[R_METHODOFF-28]
	_ = x[R_KEEP-29]
	_ = x[R_POWER_TOC-30]
	_ = x[R_GOTPCREL-31]
	_ = x[R_JMPMIPS-32]
	_ = x[R_JMPLOONG64-33]
	_ = x[R_DWARFSECREF-34]
	_ = x[R_DWARFFILEREF-35]
	_ = x[R_ARM64_TLS_LE-36]
	_ = x[R_ARM64_TLS_IE-37]
	_ = x[R_ARM64_GOTPCREL-38]
	_ = x[R_ARM64_GOT-39]
	_ = x[R_ARM64_PCREL-40]
	_ = x[R_ARM64_LDST8-41]
	_ = x[R_ARM64_LDST16-42]
	_ = x[R_ARM64_LDST32-43]
	_ = x[R_ARM64_LDST64-44]
	_ = x[R_ARM64_LDST128-45]
	_ = x[R_POWER_TLS_LE-46]
	_ = x[R_POWER_TLS_IE-47]
	_ = x[R_POWER_TLS-48]
	_ = x[R_ADDRPOWER_DS-49]
	_ = x[R_ADDRPOWER_GOT-50]
	_ = x[R_ADDRPOWER_PCREL-51]
	_ = x[R_ADDRPOWER_TOCREL-52]
	_ = x[R_ADDRPOWER_TOCREL_DS-53]
	_ = x[R_RISCV_CALL-54]
	_ = x[R_RISCV_CALL_TRAMP-55]
	_ = x[R_RISCV_PCREL_ITYPE-56]
	_ = x[R_RISCV_PCREL_STYPE-57]
	_ = x[R_RISCV_TLS_IE_ITYPE-58]
	_ = x[R_RISCV_TLS_IE_STYPE-59]
	_ = x[R_PCRELDBL-60]
	_ = x[R_ADDRMIPSU-61]
	_ = x[R_ADDRLOONG64U-62]
	_ = x[R_ADDRMIPSTLS-63]
	_ = x[R_ADDRLOONG64TLS-64]
	_ = x[R_ADDRLOONG64TLSU-65]
	_ = x[R_ADDRCUOFF-66]
	_ = x[R_WASMIMPORT-67]
	_ = x[R_XCOFFREF-68]
}

const _RelocType_name = "R_ADDRR_ADDRPOWERR_ADDRARM64R_ADDRMIPSR_ADDRLOONG64R_ADDROFFR_SIZER_CALLR_CALLARMR_CALLARM64R_CALLINDR_CALLPOWERR_CALLMIPSR_CALLLOONG64R_CONSTR_PCRELR_TLS_LER_TLS_IER_GOTOFFR_PLT0R_PLT1R_PLT2R_USEFIELDR_USETYPER_USEIFACER_USEIFACEMETHODR_USEGENERICIFACEMETHODR_METHODOFFR_KEEPR_POWER_TOCR_GOTPCRELR_JMPMIPSR_JMPLOONG64R_DWARFSECREFR_DWARFFILEREFR_ARM64_TLS_LER_ARM64_TLS_IER_ARM64_GOTPCRELR_ARM64_GOTR_ARM64_PCRELR_ARM64_LDST8R_ARM64_LDST16R_ARM64_LDST32R_ARM64_LDST64R_ARM64_LDST128R_POWER_TLS_LER_POWER_TLS_IER_POWER_TLSR_ADDRPOWER_DSR_ADDRPOWER_GOTR_ADDRPOWER_PCRELR_ADDRPOWER_TOCRELR_ADDRPOWER_TOCREL_DSR_RISCV_CALLR_RISCV_CALL_TRAMPR_RISCV_PCREL_ITYPER_RISCV_PCREL_STYPER_RISCV_TLS_IE_ITYPER_RISCV_TLS_IE_STYPER_PCRELDBLR_ADDRMIPSUR_ADDRLOONG64UR_ADDRMIPSTLSR_ADDRLOONG64TLSR_ADDRLOONG64TLSUR_ADDRCUOFFR_WASMIMPORTR_XCOFFREF"

var _RelocType_index = [...]uint16{0, 6, 17, 28, 38, 51, 60, 66, 72, 81, 92, 101, 112, 122, 135, 142, 149, 157, 165, 173, 179, 185, 191, 201, 210, 220, 236, 259, 270, 276, 287, 297, 306, 318, 331, 345, 359, 373, 389, 400, 413, 426, 440, 454, 468, 483, 497, 511, 522, 536, 551, 568, 586, 607, 619, 637, 656, 675, 695, 715, 725, 736, 750, 763, 779, 796, 807, 819, 829}

func (i RelocType) String() string {
	i -= 1
	if i < 0 || i >= RelocType(len(_RelocType_index)-1) {
		return "RelocType(" + strconv.FormatInt(int64(i+1), 10) + ")"
	}
	return _RelocType_name[_RelocType_index[i]:_RelocType_index[i+1]]
}
