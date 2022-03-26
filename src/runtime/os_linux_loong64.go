// Copyright 2021 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//go:build linux && loong64
// +build linux,loong64

package runtime

func archauxv(tag, val uintptr) {}

func osArchInit() {}

//go:nosplit
func cputicks() int64 {
	// Currently cputicks() is used in blocking profiler and to seed fastrand().
	// nanotime() is a poor approximation of CPU ticks that is enough for the profiler.
	return nanotime()
}

const (
	_SS_DISABLE  = 2
	_NSIG        = 65
	_SI_USER     = 0
	_SIG_BLOCK   = 0
	_SIG_UNBLOCK = 1
	_SIG_SETMASK = 2
)

type sigset [2]uint64

var sigset_all = sigset{^uint64(0), ^uint64(0)}

//go:nosplit
//go:nowritebarrierrec
func sigaddset(mask *sigset, i int) {
	(*mask)[(i-1)/64] |= 1 << ((uint32(i) - 1) & 63)
}

func sigdelset(mask *sigset, i int) {
	(*mask)[(i-1)/64] &^= 1 << ((uint32(i) - 1) & 63)
}

func sigfillset(mask *[2]uint64) {
	(*mask)[0], (*mask)[1] = ^uint64(0), ^uint64(0)
}
