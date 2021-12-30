// Copyright 2021 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// +build loong64

package cpu

const CacheLinePadSize = 32

// This is initialized by archauxv and should not be changed after it is
// initialized.
var HWCap uint

// HWCAP bits. These are exposed by the Linux kernel 5.4.
const (
	// CPU features
	hwcap_LOONG_MSA = 1 << 1
)

func doinit() {
	options = []option{
		{Name: "msa", Feature: &LOONG64.HasMSA},
	}

	// HWCAP feature bits
	LOONG64.HasMSA = isSet(HWCap, hwcap_LOONG_MSA)
}

func isSet(hwc uint, value uint) bool {
	return hwc&value != 0
}
