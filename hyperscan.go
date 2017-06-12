// Copyright 2017 The go-hyperscan Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// WARNING: This file has automatically been generated on Mon, 12 Jun 2017 15:00:10 JST.
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

// Compile function as declared in hs/hs_compile.h:346
func Compile(expression string, flags uint32, mode uint32, platform []PlatformInfo, db [][]Database, error [][]CompileError) int32 {
	cexpression, _ := unpackPCharString(expression)
	cflags, _ := (C.uint)(flags), cgoAllocsUnknown
	cmode, _ := (C.uint)(mode), cgoAllocsUnknown
	cplatform, _ := unpackArgSPlatformInfo(platform)
	cdb, _ := unpackArgSSDatabase(db)
	cerror, _ := unpackArgSSCompileError(error)
	__ret := C.hs_compile(cexpression, cflags, cmode, cplatform, cdb, cerror)
	packSSCompileError(error, cerror)
	packSSDatabase(db, cdb)
	packSPlatformInfo(platform, cplatform)
	__v := (int32)(__ret)
	return __v
}

// CompileMulti function as declared in hs/hs_compile.h:425
func CompileMulti(expressions []string, flags []uint32, ids []uint32, elements uint32, mode uint32, platform []PlatformInfo, db [][]Database, error [][]CompileError) int32 {
	cexpressions, _ := unpackArgSString(expressions)
	cflags, _ := (*C.uint)(unsafe.Pointer((*sliceHeader)(unsafe.Pointer(&flags)).Data)), cgoAllocsUnknown
	cids, _ := (*C.uint)(unsafe.Pointer((*sliceHeader)(unsafe.Pointer(&ids)).Data)), cgoAllocsUnknown
	celements, _ := (C.uint)(elements), cgoAllocsUnknown
	cmode, _ := (C.uint)(mode), cgoAllocsUnknown
	cplatform, _ := unpackArgSPlatformInfo(platform)
	cdb, _ := unpackArgSSDatabase(db)
	cerror, _ := unpackArgSSCompileError(error)
	__ret := C.hs_compile_multi(cexpressions, cflags, cids, celements, cmode, cplatform, cdb, cerror)
	packSSCompileError(error, cerror)
	packSSDatabase(db, cdb)
	packSPlatformInfo(platform, cplatform)
	packSString(expressions, cexpressions)
	__v := (int32)(__ret)
	return __v
}

// CompileExtMulti function as declared in hs/hs_compile.h:512
func CompileExtMulti(expressions []string, flags []uint32, ids []uint32, ext [][]ExprExt, elements uint32, mode uint32, platform []PlatformInfo, db [][]Database, error [][]CompileError) int32 {
	cexpressions, _ := unpackArgSString(expressions)
	cflags, _ := (*C.uint)(unsafe.Pointer((*sliceHeader)(unsafe.Pointer(&flags)).Data)), cgoAllocsUnknown
	cids, _ := (*C.uint)(unsafe.Pointer((*sliceHeader)(unsafe.Pointer(&ids)).Data)), cgoAllocsUnknown
	cext, _ := unpackArgSSExprExt(ext)
	celements, _ := (C.uint)(elements), cgoAllocsUnknown
	cmode, _ := (C.uint)(mode), cgoAllocsUnknown
	cplatform, _ := unpackArgSPlatformInfo(platform)
	cdb, _ := unpackArgSSDatabase(db)
	cerror, _ := unpackArgSSCompileError(error)
	__ret := C.hs_compile_ext_multi(cexpressions, cflags, cids, cext, celements, cmode, cplatform, cdb, cerror)
	packSSCompileError(error, cerror)
	packSSDatabase(db, cdb)
	packSPlatformInfo(platform, cplatform)
	packSSExprExt(ext, cext)
	packSString(expressions, cexpressions)
	__v := (int32)(__ret)
	return __v
}

// FreeCompileError function as declared in hs/hs_compile.h:531
func FreeCompileError(error []CompileError) int32 {
	cerror, _ := unpackArgSCompileError(error)
	__ret := C.hs_free_compile_error(cerror)
	packSCompileError(error, cerror)
	__v := (int32)(__ret)
	return __v
}

// ExpressionInfo function as declared in hs/hs_compile.h:590
func ExpressionInfo(expression string, flags uint32, info [][]ExprInfo, error [][]CompileError) int32 {
	cexpression, _ := unpackPCharString(expression)
	cflags, _ := (C.uint)(flags), cgoAllocsUnknown
	cinfo, _ := unpackArgSSExprInfo(info)
	cerror, _ := unpackArgSSCompileError(error)
	__ret := C.hs_expression_info(cexpression, cflags, cinfo, cerror)
	packSSCompileError(error, cerror)
	packSSExprInfo(info, cinfo)
	__v := (int32)(__ret)
	return __v
}

// ExpressionExtInfo function as declared in hs/hs_compile.h:657
func ExpressionExtInfo(expression string, flags uint32, ext []ExprExt, info [][]ExprInfo, error [][]CompileError) int32 {
	cexpression, _ := unpackPCharString(expression)
	cflags, _ := (C.uint)(flags), cgoAllocsUnknown
	cext, _ := unpackArgSExprExt(ext)
	cinfo, _ := unpackArgSSExprInfo(info)
	cerror, _ := unpackArgSSCompileError(error)
	__ret := C.hs_expression_ext_info(cexpression, cflags, cext, cinfo, cerror)
	packSSCompileError(error, cerror)
	packSSExprInfo(info, cinfo)
	packSExprExt(ext, cext)
	__v := (int32)(__ret)
	return __v
}

// PopulatePlatform function as declared in hs/hs_compile.h:673
func PopulatePlatform(platform []PlatformInfo) int32 {
	cplatform, _ := unpackArgSPlatformInfo(platform)
	__ret := C.hs_populate_platform(cplatform)
	packSPlatformInfo(platform, cplatform)
	__v := (int32)(__ret)
	return __v
}

// FreeDatabase function as declared in hs/hs_common.h:84
func FreeDatabase(db []Database) int32 {
	cdb, _ := (*C.hs_database_t)(unsafe.Pointer((*sliceHeader)(unsafe.Pointer(&db)).Data)), cgoAllocsUnknown
	__ret := C.hs_free_database(cdb)
	__v := (int32)(__ret)
	return __v
}

// SerializeDatabase function as declared in hs/hs_common.h:108
func SerializeDatabase(db []Database, bytes [][]byte, length []uint) int32 {
	cdb, _ := (*C.hs_database_t)(unsafe.Pointer((*sliceHeader)(unsafe.Pointer(&db)).Data)), cgoAllocsUnknown
	cbytes, _ := unpackArgSSByte(bytes)
	clength, _ := (*C.size_t)(unsafe.Pointer((*sliceHeader)(unsafe.Pointer(&length)).Data)), cgoAllocsUnknown
	__ret := C.hs_serialize_database(cdb, cbytes, clength)
	packSSByte(bytes, cbytes)
	__v := (int32)(__ret)
	return __v
}

// DeserializeDatabase function as declared in hs/hs_common.h:137
func DeserializeDatabase(bytes string, length uint, db [][]Database) int32 {
	cbytes, _ := unpackPCharString(bytes)
	clength, _ := (C.size_t)(length), cgoAllocsUnknown
	cdb, _ := unpackArgSSDatabase(db)
	__ret := C.hs_deserialize_database(cbytes, clength, cdb)
	packSSDatabase(db, cdb)
	__v := (int32)(__ret)
	return __v
}

// DeserializeDatabaseAt function as declared in hs/hs_common.h:169
func DeserializeDatabaseAt(bytes string, length uint, db []Database) int32 {
	cbytes, _ := unpackPCharString(bytes)
	clength, _ := (C.size_t)(length), cgoAllocsUnknown
	cdb, _ := (*C.hs_database_t)(unsafe.Pointer((*sliceHeader)(unsafe.Pointer(&db)).Data)), cgoAllocsUnknown
	__ret := C.hs_deserialize_database_at(cbytes, clength, cdb)
	__v := (int32)(__ret)
	return __v
}

// StreamSize function as declared in hs/hs_common.h:187
func StreamSize(database []Database, streamSize []uint) int32 {
	cdatabase, _ := (*C.hs_database_t)(unsafe.Pointer((*sliceHeader)(unsafe.Pointer(&database)).Data)), cgoAllocsUnknown
	cstreamSize, _ := (*C.size_t)(unsafe.Pointer((*sliceHeader)(unsafe.Pointer(&streamSize)).Data)), cgoAllocsUnknown
	__ret := C.hs_stream_size(cdatabase, cstreamSize)
	__v := (int32)(__ret)
	return __v
}

// DatabaseSize function as declared in hs/hs_common.h:203
func DatabaseSize(database []Database, databaseSize []uint) int32 {
	cdatabase, _ := (*C.hs_database_t)(unsafe.Pointer((*sliceHeader)(unsafe.Pointer(&database)).Data)), cgoAllocsUnknown
	cdatabaseSize, _ := (*C.size_t)(unsafe.Pointer((*sliceHeader)(unsafe.Pointer(&databaseSize)).Data)), cgoAllocsUnknown
	__ret := C.hs_database_size(cdatabase, cdatabaseSize)
	__v := (int32)(__ret)
	return __v
}

// SerializedDatabaseSize function as declared in hs/hs_common.h:230
func SerializedDatabaseSize(bytes string, length uint, deserializedSize []uint) int32 {
	cbytes, _ := unpackPCharString(bytes)
	clength, _ := (C.size_t)(length), cgoAllocsUnknown
	cdeserializedSize, _ := (*C.size_t)(unsafe.Pointer((*sliceHeader)(unsafe.Pointer(&deserializedSize)).Data)), cgoAllocsUnknown
	__ret := C.hs_serialized_database_size(cbytes, clength, cdeserializedSize)
	__v := (int32)(__ret)
	return __v
}

// DatabaseInfo function as declared in hs/hs_common.h:249
func DatabaseInfo(database []Database, info [][]byte) int32 {
	cdatabase, _ := (*C.hs_database_t)(unsafe.Pointer((*sliceHeader)(unsafe.Pointer(&database)).Data)), cgoAllocsUnknown
	cinfo, _ := unpackArgSSByte(info)
	__ret := C.hs_database_info(cdatabase, cinfo)
	packSSByte(info, cinfo)
	__v := (int32)(__ret)
	return __v
}

// SerializedDatabaseInfo function as declared in hs/hs_common.h:271
func SerializedDatabaseInfo(bytes string, length uint, info [][]byte) int32 {
	cbytes, _ := unpackPCharString(bytes)
	clength, _ := (C.size_t)(length), cgoAllocsUnknown
	cinfo, _ := unpackArgSSByte(info)
	__ret := C.hs_serialized_database_info(cbytes, clength, cinfo)
	packSSByte(info, cinfo)
	__v := (int32)(__ret)
	return __v
}

// SetAllocator function as declared in hs/hs_common.h:325
func SetAllocator(allocFunc Alloc, freeFunc Free) int32 {
	callocFunc, _ := allocFunc.PassValue()
	cfreeFunc, _ := freeFunc.PassValue()
	__ret := C.hs_set_allocator(callocFunc, cfreeFunc)
	__v := (int32)(__ret)
	return __v
}

// SetDatabaseAllocator function as declared in hs/hs_common.h:358
func SetDatabaseAllocator(allocFunc Alloc, freeFunc Free) int32 {
	callocFunc, _ := allocFunc.PassValue()
	cfreeFunc, _ := freeFunc.PassValue()
	__ret := C.hs_set_database_allocator(callocFunc, cfreeFunc)
	__v := (int32)(__ret)
	return __v
}

// SetMiscAllocator function as declared in hs/hs_common.h:385
func SetMiscAllocator(allocFunc Alloc, freeFunc Free) int32 {
	callocFunc, _ := allocFunc.PassValue()
	cfreeFunc, _ := freeFunc.PassValue()
	__ret := C.hs_set_misc_allocator(callocFunc, cfreeFunc)
	__v := (int32)(__ret)
	return __v
}

// SetScratchAllocator function as declared in hs/hs_common.h:412
func SetScratchAllocator(allocFunc Alloc, freeFunc Free) int32 {
	callocFunc, _ := allocFunc.PassValue()
	cfreeFunc, _ := freeFunc.PassValue()
	__ret := C.hs_set_scratch_allocator(callocFunc, cfreeFunc)
	__v := (int32)(__ret)
	return __v
}

// SetStreamAllocator function as declared in hs/hs_common.h:439
func SetStreamAllocator(allocFunc Alloc, freeFunc Free) int32 {
	callocFunc, _ := allocFunc.PassValue()
	cfreeFunc, _ := freeFunc.PassValue()
	__ret := C.hs_set_stream_allocator(callocFunc, cfreeFunc)
	__v := (int32)(__ret)
	return __v
}

// Version function as declared in hs/hs_common.h:450
func Version() string {
	__ret := C.hs_version()
	__v := packPCharString(__ret)
	return __v
}

// ValidPlatform function as declared in hs/hs_common.h:467
func ValidPlatform() int32 {
	__ret := C.hs_valid_platform()
	__v := (int32)(__ret)
	return __v
}

// OpenStream function as declared in hs/hs_runtime.h:148
func OpenStream(db []Database, flags uint32, stream [][]Stream) int32 {
	cdb, _ := (*C.hs_database_t)(unsafe.Pointer((*sliceHeader)(unsafe.Pointer(&db)).Data)), cgoAllocsUnknown
	cflags, _ := (C.uint)(flags), cgoAllocsUnknown
	cstream, _ := unpackArgSSStream(stream)
	__ret := C.hs_open_stream(cdb, cflags, cstream)
	packSSStream(stream, cstream)
	__v := (int32)(__ret)
	return __v
}

// ScanStream function as declared in hs/hs_runtime.h:188
func ScanStream(id []Stream, data string, length uint32, flags uint32, scratch []Scratch, onEvent MatchEventHandler, ctxt unsafe.Pointer) int32 {
	cid, _ := (*C.hs_stream_t)(unsafe.Pointer((*sliceHeader)(unsafe.Pointer(&id)).Data)), cgoAllocsUnknown
	cdata, _ := unpackPCharString(data)
	clength, _ := (C.uint)(length), cgoAllocsUnknown
	cflags, _ := (C.uint)(flags), cgoAllocsUnknown
	cscratch, _ := (*C.hs_scratch_t)(unsafe.Pointer((*sliceHeader)(unsafe.Pointer(&scratch)).Data)), cgoAllocsUnknown
	conEvent, _ := onEvent.PassValue()
	cctxt, _ := (unsafe.Pointer)(unsafe.Pointer(ctxt)), cgoAllocsUnknown
	__ret := C.hs_scan_stream(cid, cdata, clength, cflags, cscratch, conEvent, cctxt)
	__v := (int32)(__ret)
	return __v
}

// CloseStream function as declared in hs/hs_runtime.h:226
func CloseStream(id []Stream, scratch []Scratch, onEvent MatchEventHandler, ctxt unsafe.Pointer) int32 {
	cid, _ := (*C.hs_stream_t)(unsafe.Pointer((*sliceHeader)(unsafe.Pointer(&id)).Data)), cgoAllocsUnknown
	cscratch, _ := (*C.hs_scratch_t)(unsafe.Pointer((*sliceHeader)(unsafe.Pointer(&scratch)).Data)), cgoAllocsUnknown
	conEvent, _ := onEvent.PassValue()
	cctxt, _ := (unsafe.Pointer)(unsafe.Pointer(ctxt)), cgoAllocsUnknown
	__ret := C.hs_close_stream(cid, cscratch, conEvent, cctxt)
	__v := (int32)(__ret)
	return __v
}

// ResetStream function as declared in hs/hs_runtime.h:267
func ResetStream(id []Stream, flags uint32, scratch []Scratch, onEvent MatchEventHandler, context unsafe.Pointer) int32 {
	cid, _ := (*C.hs_stream_t)(unsafe.Pointer((*sliceHeader)(unsafe.Pointer(&id)).Data)), cgoAllocsUnknown
	cflags, _ := (C.uint)(flags), cgoAllocsUnknown
	cscratch, _ := (*C.hs_scratch_t)(unsafe.Pointer((*sliceHeader)(unsafe.Pointer(&scratch)).Data)), cgoAllocsUnknown
	conEvent, _ := onEvent.PassValue()
	ccontext, _ := (unsafe.Pointer)(unsafe.Pointer(context)), cgoAllocsUnknown
	__ret := C.hs_reset_stream(cid, cflags, cscratch, conEvent, ccontext)
	__v := (int32)(__ret)
	return __v
}

// CopyStream function as declared in hs/hs_runtime.h:285
func CopyStream(toId [][]Stream, fromId []Stream) int32 {
	ctoId, _ := unpackArgSSStream(toId)
	cfromId, _ := (*C.hs_stream_t)(unsafe.Pointer((*sliceHeader)(unsafe.Pointer(&fromId)).Data)), cgoAllocsUnknown
	__ret := C.hs_copy_stream(ctoId, cfromId)
	packSSStream(toId, ctoId)
	__v := (int32)(__ret)
	return __v
}

// ResetAndCopyStream function as declared in hs/hs_runtime.h:318
func ResetAndCopyStream(toId []Stream, fromId []Stream, scratch []Scratch, onEvent MatchEventHandler, context unsafe.Pointer) int32 {
	ctoId, _ := (*C.hs_stream_t)(unsafe.Pointer((*sliceHeader)(unsafe.Pointer(&toId)).Data)), cgoAllocsUnknown
	cfromId, _ := (*C.hs_stream_t)(unsafe.Pointer((*sliceHeader)(unsafe.Pointer(&fromId)).Data)), cgoAllocsUnknown
	cscratch, _ := (*C.hs_scratch_t)(unsafe.Pointer((*sliceHeader)(unsafe.Pointer(&scratch)).Data)), cgoAllocsUnknown
	conEvent, _ := onEvent.PassValue()
	ccontext, _ := (unsafe.Pointer)(unsafe.Pointer(context)), cgoAllocsUnknown
	__ret := C.hs_reset_and_copy_stream(ctoId, cfromId, cscratch, conEvent, ccontext)
	__v := (int32)(__ret)
	return __v
}

// Scan function as declared in hs/hs_runtime.h:359
func Scan(db []Database, data string, length uint32, flags uint32, scratch []Scratch, onEvent MatchEventHandler, context unsafe.Pointer) int32 {
	cdb, _ := (*C.hs_database_t)(unsafe.Pointer((*sliceHeader)(unsafe.Pointer(&db)).Data)), cgoAllocsUnknown
	cdata, _ := unpackPCharString(data)
	clength, _ := (C.uint)(length), cgoAllocsUnknown
	cflags, _ := (C.uint)(flags), cgoAllocsUnknown
	cscratch, _ := (*C.hs_scratch_t)(unsafe.Pointer((*sliceHeader)(unsafe.Pointer(&scratch)).Data)), cgoAllocsUnknown
	conEvent, _ := onEvent.PassValue()
	ccontext, _ := (unsafe.Pointer)(unsafe.Pointer(context)), cgoAllocsUnknown
	__ret := C.hs_scan(cdb, cdata, clength, cflags, cscratch, conEvent, ccontext)
	__v := (int32)(__ret)
	return __v
}

// ScanVector function as declared in hs/hs_runtime.h:402
func ScanVector(db []Database, data []string, length []uint32, count uint32, flags uint32, scratch []Scratch, onEvent MatchEventHandler, context unsafe.Pointer) int32 {
	cdb, _ := (*C.hs_database_t)(unsafe.Pointer((*sliceHeader)(unsafe.Pointer(&db)).Data)), cgoAllocsUnknown
	cdata, _ := unpackArgSString(data)
	clength, _ := (*C.uint)(unsafe.Pointer((*sliceHeader)(unsafe.Pointer(&length)).Data)), cgoAllocsUnknown
	ccount, _ := (C.uint)(count), cgoAllocsUnknown
	cflags, _ := (C.uint)(flags), cgoAllocsUnknown
	cscratch, _ := (*C.hs_scratch_t)(unsafe.Pointer((*sliceHeader)(unsafe.Pointer(&scratch)).Data)), cgoAllocsUnknown
	conEvent, _ := onEvent.PassValue()
	ccontext, _ := (unsafe.Pointer)(unsafe.Pointer(context)), cgoAllocsUnknown
	__ret := C.hs_scan_vector(cdb, cdata, clength, ccount, cflags, cscratch, conEvent, ccontext)
	packSString(data, cdata)
	__v := (int32)(__ret)
	return __v
}

// AllocScratch function as declared in hs/hs_runtime.h:435
func AllocScratch(db []Database, scratch [][]Scratch) int32 {
	cdb, _ := (*C.hs_database_t)(unsafe.Pointer((*sliceHeader)(unsafe.Pointer(&db)).Data)), cgoAllocsUnknown
	cscratch, _ := unpackArgSSScratch(scratch)
	__ret := C.hs_alloc_scratch(cdb, cscratch)
	packSSScratch(scratch, cscratch)
	__v := (int32)(__ret)
	return __v
}

// CloneScratch function as declared in hs/hs_runtime.h:456
func CloneScratch(src []Scratch, dest [][]Scratch) int32 {
	csrc, _ := (*C.hs_scratch_t)(unsafe.Pointer((*sliceHeader)(unsafe.Pointer(&src)).Data)), cgoAllocsUnknown
	cdest, _ := unpackArgSSScratch(dest)
	__ret := C.hs_clone_scratch(csrc, cdest)
	packSSScratch(dest, cdest)
	__v := (int32)(__ret)
	return __v
}

// ScratchSize function as declared in hs/hs_runtime.h:473
func ScratchSize(scratch []Scratch, scratchSize []uint) int32 {
	cscratch, _ := (*C.hs_scratch_t)(unsafe.Pointer((*sliceHeader)(unsafe.Pointer(&scratch)).Data)), cgoAllocsUnknown
	cscratchSize, _ := (*C.size_t)(unsafe.Pointer((*sliceHeader)(unsafe.Pointer(&scratchSize)).Data)), cgoAllocsUnknown
	__ret := C.hs_scratch_size(cscratch, cscratchSize)
	__v := (int32)(__ret)
	return __v
}

// FreeScratch function as declared in hs/hs_runtime.h:489
func FreeScratch(scratch []Scratch) int32 {
	cscratch, _ := (*C.hs_scratch_t)(unsafe.Pointer((*sliceHeader)(unsafe.Pointer(&scratch)).Data)), cgoAllocsUnknown
	__ret := C.hs_free_scratch(cscratch)
	__v := (int32)(__ret)
	return __v
}
