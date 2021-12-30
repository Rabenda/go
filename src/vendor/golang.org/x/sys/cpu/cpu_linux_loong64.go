// Copyright 2021 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// +build linux,loong64

package cpu

// HWCAP bits. These are exposed by the Linux kernel 5.4.
const (
	// CPU features
	hwcap_LOONG64_MSA = 1 << 1
)

func doinit() {
	// HWCAP feature bits
	LOONG64.HasMSA = isSet(hwCap, hwcap_LOONG64_MSA)
}

func isSet(hwc uint, value uint) bool {
	return hwc&value != 0
}
