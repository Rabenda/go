// Copyright 2021 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// This input was created by taking the ppc64 testcase and modified
// by hand.

#include "../../../../../runtime/textflag.h"

TEXT foo(SB),DUPOK|NOSPLIT,$0
    JAL 1(PC) //CALL 1(PC) //000c0054
    JAL (R4) //CALL (R4) //8100004c
    JAL foo(SB) //CALL foo(SB) //00100054
