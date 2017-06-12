// Copyright 2017 The go-hyperscan Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// WARNING: This file has automatically been generated on Mon, 12 Jun 2017 17:47:08 JST.
// By https://git.io/c-for-go. DO NOT EDIT.

package hyperscan

/*
#cgo CFLAGS: -I${SRCDIR}/include
#cgo LDFLAGS: -lhs -lc++
#include "hs/hs.h"
#include "hs/hs_common.h"
#include "hs/hs_compile.h"
#include "hs/hs_runtime.h"
#include <stdlib.h>
#include "cgo_helpers.h"
*/
import "C"
import "unsafe"

// CompileError as declared in hs/hs_compile.h:124
type CompileError struct {
	Message        []byte
	Expression     int32
	ref3a3e0619    *C.hs_compile_error_t
	allocs3a3e0619 interface{}
}

// PlatformInfo as declared in hs/hs_compile.h:163
type PlatformInfo struct {
	Tune           uint32
	CpuFeatures    uint64
	Reserved1      uint64
	Reserved2      uint64
	ref549e6d63    *C.hs_platform_info_t
	allocs549e6d63 interface{}
}

// ExprInfo as declared in hs/hs_compile.h:216
type ExprInfo struct {
	MinWidth         uint32
	MaxWidth         uint32
	UnorderedMatches byte
	MatchesAtEod     byte
	MatchesOnlyAtEod byte
	ref16cc495d      *C.hs_expr_info_t
	allocs16cc495d   interface{}
}

// ExprExt as declared in hs/hs_compile.h:261
type ExprExt struct {
	Flags          uint64
	MinOffset      uint64
	MaxOffset      uint64
	MinLength      uint64
	EditDistance   uint32
	ref73363562    *C.hs_expr_ext_t
	allocs73363562 interface{}
}

// Database as declared in hs/hs_common.h:64
type Database C.hs_database_t

// Error type as declared in hs/hs_common.h:69
type Error int32

// Alloc type as declared in hs/hs_common.h:288
type Alloc func(size uint) unsafe.Pointer

// Free type as declared in hs/hs_common.h:297
type Free func(ptr unsafe.Pointer)

// Stream as declared in hs/hs_runtime.h:59
type Stream C.hs_stream_t

// Scratch as declared in hs/hs_runtime.h:66
type Scratch C.hs_scratch_t

// MatchEventHandler type as declared in hs/hs_runtime.h:125
type MatchEventHandler func(id uint32, from uint64, to uint64, flags uint32, context unsafe.Pointer) int32
