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

const (
	// ExtFlagMinOffset as defined in hs/hs_compile.h:273
	ExtFlagMinOffset = uint64(1)
	// ExtFlagMaxOffset as defined in hs/hs_compile.h:276
	ExtFlagMaxOffset = uint64(2)
	// ExtFlagMinLength as defined in hs/hs_compile.h:279
	ExtFlagMinLength = uint64(4)
	// ExtFlagEditDistance as defined in hs/hs_compile.h:282
	ExtFlagEditDistance = uint64(8)
	// FlagCaseless as defined in hs/hs_compile.h:688
	FlagCaseless = 1
	// FlagDotall as defined in hs/hs_compile.h:698
	FlagDotall = 2
	// FlagMultiline as defined in hs/hs_compile.h:709
	FlagMultiline = 4
	// FlagSinglematch as defined in hs/hs_compile.h:730
	FlagSinglematch = 8
	// FlagAllowempty as defined in hs/hs_compile.h:742
	FlagAllowempty = 16
	// FlagUtf8 as defined in hs/hs_compile.h:752
	FlagUtf8 = 32
	// FlagUcp as defined in hs/hs_compile.h:762
	FlagUcp = 64
	// FlagPrefilter as defined in hs/hs_compile.h:790
	FlagPrefilter = 128
	// FlagSomLeftmost as defined in hs/hs_compile.h:802
	FlagSomLeftmost = 256
	// CpuFeaturesAvx2 as defined in hs/hs_compile.h:818
	CpuFeaturesAvx2 = (uint64(1) << 2)
	// CpuFeaturesAvx512 as defined in hs/hs_compile.h:826
	CpuFeaturesAvx512 = (uint64(1) << 3)
	// TuneFamilyGeneric as defined in hs/hs_compile.h:842
	TuneFamilyGeneric = 0
	// TuneFamilySnb as defined in hs/hs_compile.h:850
	TuneFamilySnb = 1
	// TuneFamilyIvb as defined in hs/hs_compile.h:858
	TuneFamilyIvb = 2
	// TuneFamilyHsw as defined in hs/hs_compile.h:866
	TuneFamilyHsw = 3
	// TuneFamilySlm as defined in hs/hs_compile.h:874
	TuneFamilySlm = 4
	// TuneFamilyBdw as defined in hs/hs_compile.h:882
	TuneFamilyBdw = 5
	// TuneFamilySkl as defined in hs/hs_compile.h:890
	TuneFamilySkl = 6
	// TuneFamilySkx as defined in hs/hs_compile.h:898
	TuneFamilySkx = 7
	// TuneFamilyGlm as defined in hs/hs_compile.h:906
	TuneFamilyGlm = 8
	// ModeBlock as defined in hs/hs_compile.h:928
	ModeBlock = 1
	// ModeNostream as defined in hs/hs_compile.h:933
	ModeNostream = 1
	// ModeStream as defined in hs/hs_compile.h:938
	ModeStream = 2
	// ModeVectored as defined in hs/hs_compile.h:943
	ModeVectored = 4
	// ModeSomHorizonLarge as defined in hs/hs_compile.h:956
	ModeSomHorizonLarge = (uint32(1) << 24)
	// ModeSomHorizonMedium as defined in hs/hs_compile.h:969
	ModeSomHorizonMedium = (uint32(1) << 25)
	// ModeSomHorizonSmall as defined in hs/hs_compile.h:982
	ModeSomHorizonSmall = (uint32(1) << 26)
	// Success as defined in hs/hs_common.h:478
	Success = 0
	// Invalid as defined in hs/hs_common.h:483
	Invalid = (-1)
	// Nomem as defined in hs/hs_common.h:488
	Nomem = (-2)
	// ScanTerminated as defined in hs/hs_common.h:497
	ScanTerminated = (-3)
	// CompilerError as defined in hs/hs_common.h:503
	CompilerError = (-4)
	// DbVersionError as defined in hs/hs_common.h:508
	DbVersionError = (-5)
	// DbPlatformError as defined in hs/hs_common.h:513
	DbPlatformError = (-6)
	// DbModeError as defined in hs/hs_common.h:520
	DbModeError = (-7)
	// BadAlign as defined in hs/hs_common.h:525
	BadAlign = (-8)
	// BadAlloc as defined in hs/hs_common.h:532
	BadAlloc = (-9)
	// ScratchInUse as defined in hs/hs_common.h:551
	ScratchInUse = (-10)
	// ArchError as defined in hs/hs_common.h:562
	ArchError = (-11)
	// OffsetPastHorizon as defined in hs/hs_runtime.h:495
	OffsetPastHorizon = (^uint64(0))
)
