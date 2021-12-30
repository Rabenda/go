// Copyright 2021 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// +build linux,loong64

package cpu

const cacheLineSize = 32

func initOptions() {
	options = []option{
		{Name: "msa", Feature: &LOONG64.HasMSA},
	}
}