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
import (
	"runtime"
	"sync"
	"unsafe"
)

// cgoAllocMap stores pointers to C allocated memory for future reference.
type cgoAllocMap struct {
	mux sync.RWMutex
	m   map[unsafe.Pointer]struct{}
}

var cgoAllocsUnknown = new(cgoAllocMap)

func (a *cgoAllocMap) Add(ptr unsafe.Pointer) {
	a.mux.Lock()
	if a.m == nil {
		a.m = make(map[unsafe.Pointer]struct{})
	}
	a.m[ptr] = struct{}{}
	a.mux.Unlock()
}

func (a *cgoAllocMap) IsEmpty() bool {
	a.mux.RLock()
	isEmpty := len(a.m) == 0
	a.mux.RUnlock()
	return isEmpty
}

func (a *cgoAllocMap) Borrow(b *cgoAllocMap) {
	if b == nil || b.IsEmpty() {
		return
	}
	b.mux.Lock()
	a.mux.Lock()
	for ptr := range b.m {
		if a.m == nil {
			a.m = make(map[unsafe.Pointer]struct{})
		}
		a.m[ptr] = struct{}{}
		delete(b.m, ptr)
	}
	a.mux.Unlock()
	b.mux.Unlock()
}

func (a *cgoAllocMap) Free() {
	a.mux.Lock()
	for ptr := range a.m {
		C.free(ptr)
		delete(a.m, ptr)
	}
	a.mux.Unlock()
}

// allocCompileErrorMemory allocates memory for type C.hs_compile_error_t in C.
// The caller is responsible for freeing the this memory via C.free.
func allocCompileErrorMemory(n int) unsafe.Pointer {
	mem, err := C.calloc(C.size_t(n), (C.size_t)(sizeOfCompileErrorValue))
	if err != nil {
		panic("memory alloc error: " + err.Error())
	}
	return mem
}

const sizeOfCompileErrorValue = unsafe.Sizeof([1]C.hs_compile_error_t{})

type sliceHeader struct {
	Data uintptr
	Len  int
	Cap  int
}

// Ref returns the underlying reference to C object or nil if struct is nil.
func (x *CompileError) Ref() *C.hs_compile_error_t {
	if x == nil {
		return nil
	}
	return x.ref3a3e0619
}

// Free invokes alloc map's free mechanism that cleanups any allocated memory using C free.
// Does nothing if struct is nil or has no allocation map.
func (x *CompileError) Free() {
	if x != nil && x.allocs3a3e0619 != nil {
		x.allocs3a3e0619.(*cgoAllocMap).Free()
		x.ref3a3e0619 = nil
	}
}

// NewCompileErrorRef creates a new wrapper struct with underlying reference set to the original C object.
// Returns nil if the provided pointer to C object is nil too.
func NewCompileErrorRef(ref unsafe.Pointer) *CompileError {
	if ref == nil {
		return nil
	}
	obj := new(CompileError)
	obj.ref3a3e0619 = (*C.hs_compile_error_t)(unsafe.Pointer(ref))
	return obj
}

// PassRef returns the underlying C object, otherwise it will allocate one and set its values
// from this wrapping struct, counting allocations into an allocation map.
func (x *CompileError) PassRef() (*C.hs_compile_error_t, *cgoAllocMap) {
	if x == nil {
		return nil, nil
	} else if x.ref3a3e0619 != nil {
		return x.ref3a3e0619, nil
	}
	mem3a3e0619 := allocCompileErrorMemory(1)
	ref3a3e0619 := (*C.hs_compile_error_t)(mem3a3e0619)
	allocs3a3e0619 := new(cgoAllocMap)
	var cmessage_allocs *cgoAllocMap
	ref3a3e0619.message, cmessage_allocs = (*C.char)(unsafe.Pointer((*sliceHeader)(unsafe.Pointer(&x.Message)).Data)), cgoAllocsUnknown
	allocs3a3e0619.Borrow(cmessage_allocs)

	var cexpression_allocs *cgoAllocMap
	ref3a3e0619.expression, cexpression_allocs = (C.int)(x.Expression), cgoAllocsUnknown
	allocs3a3e0619.Borrow(cexpression_allocs)

	x.ref3a3e0619 = ref3a3e0619
	x.allocs3a3e0619 = allocs3a3e0619
	return ref3a3e0619, allocs3a3e0619

}

// PassValue does the same as PassRef except that it will try to dereference the returned pointer.
func (x CompileError) PassValue() (C.hs_compile_error_t, *cgoAllocMap) {
	if x.ref3a3e0619 != nil {
		return *x.ref3a3e0619, nil
	}
	ref, allocs := x.PassRef()
	return *ref, allocs
}

// Deref uses the underlying reference to C object and fills the wrapping struct with values.
// Do not forget to call this method whether you get a struct for C object and want to read its values.
func (x *CompileError) Deref() {
	if x.ref3a3e0619 == nil {
		return
	}
	hxfc4425b := (*sliceHeader)(unsafe.Pointer(&x.Message))
	hxfc4425b.Data = uintptr(unsafe.Pointer(x.ref3a3e0619.message))
	hxfc4425b.Cap = 0x7fffffff
	// hxfc4425b.Len = ?

	x.Expression = (int32)(x.ref3a3e0619.expression)
}

// allocPlatformInfoMemory allocates memory for type C.hs_platform_info_t in C.
// The caller is responsible for freeing the this memory via C.free.
func allocPlatformInfoMemory(n int) unsafe.Pointer {
	mem, err := C.calloc(C.size_t(n), (C.size_t)(sizeOfPlatformInfoValue))
	if err != nil {
		panic("memory alloc error: " + err.Error())
	}
	return mem
}

const sizeOfPlatformInfoValue = unsafe.Sizeof([1]C.hs_platform_info_t{})

// Ref returns the underlying reference to C object or nil if struct is nil.
func (x *PlatformInfo) Ref() *C.hs_platform_info_t {
	if x == nil {
		return nil
	}
	return x.ref549e6d63
}

// Free invokes alloc map's free mechanism that cleanups any allocated memory using C free.
// Does nothing if struct is nil or has no allocation map.
func (x *PlatformInfo) Free() {
	if x != nil && x.allocs549e6d63 != nil {
		x.allocs549e6d63.(*cgoAllocMap).Free()
		x.ref549e6d63 = nil
	}
}

// NewPlatformInfoRef creates a new wrapper struct with underlying reference set to the original C object.
// Returns nil if the provided pointer to C object is nil too.
func NewPlatformInfoRef(ref unsafe.Pointer) *PlatformInfo {
	if ref == nil {
		return nil
	}
	obj := new(PlatformInfo)
	obj.ref549e6d63 = (*C.hs_platform_info_t)(unsafe.Pointer(ref))
	return obj
}

// PassRef returns the underlying C object, otherwise it will allocate one and set its values
// from this wrapping struct, counting allocations into an allocation map.
func (x *PlatformInfo) PassRef() (*C.hs_platform_info_t, *cgoAllocMap) {
	if x == nil {
		return nil, nil
	} else if x.ref549e6d63 != nil {
		return x.ref549e6d63, nil
	}
	mem549e6d63 := allocPlatformInfoMemory(1)
	ref549e6d63 := (*C.hs_platform_info_t)(mem549e6d63)
	allocs549e6d63 := new(cgoAllocMap)
	var ctune_allocs *cgoAllocMap
	ref549e6d63.tune, ctune_allocs = (C.uint)(x.Tune), cgoAllocsUnknown
	allocs549e6d63.Borrow(ctune_allocs)

	var ccpu_features_allocs *cgoAllocMap
	ref549e6d63.cpu_features, ccpu_features_allocs = (C.ulonglong)(x.CpuFeatures), cgoAllocsUnknown
	allocs549e6d63.Borrow(ccpu_features_allocs)

	var creserved1_allocs *cgoAllocMap
	ref549e6d63.reserved1, creserved1_allocs = (C.ulonglong)(x.Reserved1), cgoAllocsUnknown
	allocs549e6d63.Borrow(creserved1_allocs)

	var creserved2_allocs *cgoAllocMap
	ref549e6d63.reserved2, creserved2_allocs = (C.ulonglong)(x.Reserved2), cgoAllocsUnknown
	allocs549e6d63.Borrow(creserved2_allocs)

	x.ref549e6d63 = ref549e6d63
	x.allocs549e6d63 = allocs549e6d63
	return ref549e6d63, allocs549e6d63

}

// PassValue does the same as PassRef except that it will try to dereference the returned pointer.
func (x PlatformInfo) PassValue() (C.hs_platform_info_t, *cgoAllocMap) {
	if x.ref549e6d63 != nil {
		return *x.ref549e6d63, nil
	}
	ref, allocs := x.PassRef()
	return *ref, allocs
}

// Deref uses the underlying reference to C object and fills the wrapping struct with values.
// Do not forget to call this method whether you get a struct for C object and want to read its values.
func (x *PlatformInfo) Deref() {
	if x.ref549e6d63 == nil {
		return
	}
	x.Tune = (uint32)(x.ref549e6d63.tune)
	x.CpuFeatures = (uint64)(x.ref549e6d63.cpu_features)
	x.Reserved1 = (uint64)(x.ref549e6d63.reserved1)
	x.Reserved2 = (uint64)(x.ref549e6d63.reserved2)
}

// allocExprInfoMemory allocates memory for type C.hs_expr_info_t in C.
// The caller is responsible for freeing the this memory via C.free.
func allocExprInfoMemory(n int) unsafe.Pointer {
	mem, err := C.calloc(C.size_t(n), (C.size_t)(sizeOfExprInfoValue))
	if err != nil {
		panic("memory alloc error: " + err.Error())
	}
	return mem
}

const sizeOfExprInfoValue = unsafe.Sizeof([1]C.hs_expr_info_t{})

// Ref returns the underlying reference to C object or nil if struct is nil.
func (x *ExprInfo) Ref() *C.hs_expr_info_t {
	if x == nil {
		return nil
	}
	return x.ref16cc495d
}

// Free invokes alloc map's free mechanism that cleanups any allocated memory using C free.
// Does nothing if struct is nil or has no allocation map.
func (x *ExprInfo) Free() {
	if x != nil && x.allocs16cc495d != nil {
		x.allocs16cc495d.(*cgoAllocMap).Free()
		x.ref16cc495d = nil
	}
}

// NewExprInfoRef creates a new wrapper struct with underlying reference set to the original C object.
// Returns nil if the provided pointer to C object is nil too.
func NewExprInfoRef(ref unsafe.Pointer) *ExprInfo {
	if ref == nil {
		return nil
	}
	obj := new(ExprInfo)
	obj.ref16cc495d = (*C.hs_expr_info_t)(unsafe.Pointer(ref))
	return obj
}

// PassRef returns the underlying C object, otherwise it will allocate one and set its values
// from this wrapping struct, counting allocations into an allocation map.
func (x *ExprInfo) PassRef() (*C.hs_expr_info_t, *cgoAllocMap) {
	if x == nil {
		return nil, nil
	} else if x.ref16cc495d != nil {
		return x.ref16cc495d, nil
	}
	mem16cc495d := allocExprInfoMemory(1)
	ref16cc495d := (*C.hs_expr_info_t)(mem16cc495d)
	allocs16cc495d := new(cgoAllocMap)
	var cmin_width_allocs *cgoAllocMap
	ref16cc495d.min_width, cmin_width_allocs = (C.uint)(x.MinWidth), cgoAllocsUnknown
	allocs16cc495d.Borrow(cmin_width_allocs)

	var cmax_width_allocs *cgoAllocMap
	ref16cc495d.max_width, cmax_width_allocs = (C.uint)(x.MaxWidth), cgoAllocsUnknown
	allocs16cc495d.Borrow(cmax_width_allocs)

	var cunordered_matches_allocs *cgoAllocMap
	ref16cc495d.unordered_matches, cunordered_matches_allocs = (C.char)(x.UnorderedMatches), cgoAllocsUnknown
	allocs16cc495d.Borrow(cunordered_matches_allocs)

	var cmatches_at_eod_allocs *cgoAllocMap
	ref16cc495d.matches_at_eod, cmatches_at_eod_allocs = (C.char)(x.MatchesAtEod), cgoAllocsUnknown
	allocs16cc495d.Borrow(cmatches_at_eod_allocs)

	var cmatches_only_at_eod_allocs *cgoAllocMap
	ref16cc495d.matches_only_at_eod, cmatches_only_at_eod_allocs = (C.char)(x.MatchesOnlyAtEod), cgoAllocsUnknown
	allocs16cc495d.Borrow(cmatches_only_at_eod_allocs)

	x.ref16cc495d = ref16cc495d
	x.allocs16cc495d = allocs16cc495d
	return ref16cc495d, allocs16cc495d

}

// PassValue does the same as PassRef except that it will try to dereference the returned pointer.
func (x ExprInfo) PassValue() (C.hs_expr_info_t, *cgoAllocMap) {
	if x.ref16cc495d != nil {
		return *x.ref16cc495d, nil
	}
	ref, allocs := x.PassRef()
	return *ref, allocs
}

// Deref uses the underlying reference to C object and fills the wrapping struct with values.
// Do not forget to call this method whether you get a struct for C object and want to read its values.
func (x *ExprInfo) Deref() {
	if x.ref16cc495d == nil {
		return
	}
	x.MinWidth = (uint32)(x.ref16cc495d.min_width)
	x.MaxWidth = (uint32)(x.ref16cc495d.max_width)
	x.UnorderedMatches = (byte)(x.ref16cc495d.unordered_matches)
	x.MatchesAtEod = (byte)(x.ref16cc495d.matches_at_eod)
	x.MatchesOnlyAtEod = (byte)(x.ref16cc495d.matches_only_at_eod)
}

// allocExprExtMemory allocates memory for type C.hs_expr_ext_t in C.
// The caller is responsible for freeing the this memory via C.free.
func allocExprExtMemory(n int) unsafe.Pointer {
	mem, err := C.calloc(C.size_t(n), (C.size_t)(sizeOfExprExtValue))
	if err != nil {
		panic("memory alloc error: " + err.Error())
	}
	return mem
}

const sizeOfExprExtValue = unsafe.Sizeof([1]C.hs_expr_ext_t{})

// Ref returns the underlying reference to C object or nil if struct is nil.
func (x *ExprExt) Ref() *C.hs_expr_ext_t {
	if x == nil {
		return nil
	}
	return x.ref73363562
}

// Free invokes alloc map's free mechanism that cleanups any allocated memory using C free.
// Does nothing if struct is nil or has no allocation map.
func (x *ExprExt) Free() {
	if x != nil && x.allocs73363562 != nil {
		x.allocs73363562.(*cgoAllocMap).Free()
		x.ref73363562 = nil
	}
}

// NewExprExtRef creates a new wrapper struct with underlying reference set to the original C object.
// Returns nil if the provided pointer to C object is nil too.
func NewExprExtRef(ref unsafe.Pointer) *ExprExt {
	if ref == nil {
		return nil
	}
	obj := new(ExprExt)
	obj.ref73363562 = (*C.hs_expr_ext_t)(unsafe.Pointer(ref))
	return obj
}

// PassRef returns the underlying C object, otherwise it will allocate one and set its values
// from this wrapping struct, counting allocations into an allocation map.
func (x *ExprExt) PassRef() (*C.hs_expr_ext_t, *cgoAllocMap) {
	if x == nil {
		return nil, nil
	} else if x.ref73363562 != nil {
		return x.ref73363562, nil
	}
	mem73363562 := allocExprExtMemory(1)
	ref73363562 := (*C.hs_expr_ext_t)(mem73363562)
	allocs73363562 := new(cgoAllocMap)
	var cflags_allocs *cgoAllocMap
	ref73363562.flags, cflags_allocs = (C.ulonglong)(x.Flags), cgoAllocsUnknown
	allocs73363562.Borrow(cflags_allocs)

	var cmin_offset_allocs *cgoAllocMap
	ref73363562.min_offset, cmin_offset_allocs = (C.ulonglong)(x.MinOffset), cgoAllocsUnknown
	allocs73363562.Borrow(cmin_offset_allocs)

	var cmax_offset_allocs *cgoAllocMap
	ref73363562.max_offset, cmax_offset_allocs = (C.ulonglong)(x.MaxOffset), cgoAllocsUnknown
	allocs73363562.Borrow(cmax_offset_allocs)

	var cmin_length_allocs *cgoAllocMap
	ref73363562.min_length, cmin_length_allocs = (C.ulonglong)(x.MinLength), cgoAllocsUnknown
	allocs73363562.Borrow(cmin_length_allocs)

	var cedit_distance_allocs *cgoAllocMap
	ref73363562.edit_distance, cedit_distance_allocs = (C.uint)(x.EditDistance), cgoAllocsUnknown
	allocs73363562.Borrow(cedit_distance_allocs)

	x.ref73363562 = ref73363562
	x.allocs73363562 = allocs73363562
	return ref73363562, allocs73363562

}

// PassValue does the same as PassRef except that it will try to dereference the returned pointer.
func (x ExprExt) PassValue() (C.hs_expr_ext_t, *cgoAllocMap) {
	if x.ref73363562 != nil {
		return *x.ref73363562, nil
	}
	ref, allocs := x.PassRef()
	return *ref, allocs
}

// Deref uses the underlying reference to C object and fills the wrapping struct with values.
// Do not forget to call this method whether you get a struct for C object and want to read its values.
func (x *ExprExt) Deref() {
	if x.ref73363562 == nil {
		return
	}
	x.Flags = (uint64)(x.ref73363562.flags)
	x.MinOffset = (uint64)(x.ref73363562.min_offset)
	x.MaxOffset = (uint64)(x.ref73363562.max_offset)
	x.MinLength = (uint64)(x.ref73363562.min_length)
	x.EditDistance = (uint32)(x.ref73363562.edit_distance)
}

// Ref returns a reference to C object as it is.
func (x *Database) Ref() *C.hs_database_t {
	if x == nil {
		return nil
	}
	return (*C.hs_database_t)(unsafe.Pointer(x))
}

// Free cleanups the referenced memory using C free.
func (x *Database) Free() {
	if x != nil {
		C.free(unsafe.Pointer(x))
	}
}

// NewDatabaseRef converts the C object reference into a raw struct reference without wrapping.
func NewDatabaseRef(ref unsafe.Pointer) *Database {
	return (*Database)(ref)
}

// NewDatabase allocates a new C object of this type and converts the reference into
// a raw struct reference without wrapping.
func NewDatabase() *Database {
	return (*Database)(allocDatabaseMemory(1))
}

// allocDatabaseMemory allocates memory for type C.hs_database_t in C.
// The caller is responsible for freeing the this memory via C.free.
func allocDatabaseMemory(n int) unsafe.Pointer {
	mem, err := C.calloc(C.size_t(n), (C.size_t)(sizeOfDatabaseValue))
	if err != nil {
		panic("memory alloc error: " + err.Error())
	}
	return mem
}

const sizeOfDatabaseValue = unsafe.Sizeof([1]C.hs_database_t{})

// PassRef returns a reference to C object as it is or allocates a new C object of this type.
func (x *Database) PassRef() *C.hs_database_t {
	if x == nil {
		x = (*Database)(allocDatabaseMemory(1))
	}
	return (*C.hs_database_t)(unsafe.Pointer(x))
}

func (x Alloc) PassRef() (ref *C.hs_alloc_t, allocs *cgoAllocMap) {
	if x == nil {
		return nil, nil
	}
	if alloc29384626Func == nil {
		alloc29384626Func = x
	}
	return (*C.hs_alloc_t)(C.hs_alloc_t_29384626), nil
}

func (x Alloc) PassValue() (ref C.hs_alloc_t, allocs *cgoAllocMap) {
	if x == nil {
		return nil, nil
	}
	if alloc29384626Func == nil {
		alloc29384626Func = x
	}
	return (C.hs_alloc_t)(C.hs_alloc_t_29384626), nil
}

func NewAllocRef(ref unsafe.Pointer) *Alloc {
	return (*Alloc)(ref)
}

//export alloc29384626
func alloc29384626(csize C.size_t) unsafe.Pointer {
	if alloc29384626Func != nil {
		size29384626 := (uint)(csize)
		ret29384626 := alloc29384626Func(size29384626)
		ret, _ := (unsafe.Pointer)(unsafe.Pointer(ret29384626)), cgoAllocsUnknown
		return ret
	}
	panic("callback func has not been set (race?)")
}

var alloc29384626Func Alloc

func (x Free) PassRef() (ref *C.hs_free_t, allocs *cgoAllocMap) {
	if x == nil {
		return nil, nil
	}
	if free83E09D93Func == nil {
		free83E09D93Func = x
	}
	return (*C.hs_free_t)(C.hs_free_t_83e09d93), nil
}

func (x Free) PassValue() (ref C.hs_free_t, allocs *cgoAllocMap) {
	if x == nil {
		return nil, nil
	}
	if free83E09D93Func == nil {
		free83E09D93Func = x
	}
	return (C.hs_free_t)(C.hs_free_t_83e09d93), nil
}

func NewFreeRef(ref unsafe.Pointer) *Free {
	return (*Free)(ref)
}

//export free83E09D93
func free83E09D93(cptr unsafe.Pointer) {
	if free83E09D93Func != nil {
		ptr83e09d93 := (unsafe.Pointer)(unsafe.Pointer(cptr))
		free83E09D93Func(ptr83e09d93)
		return
	}
	panic("callback func has not been set (race?)")
}

var free83E09D93Func Free

// Ref returns a reference to C object as it is.
func (x *Stream) Ref() *C.hs_stream_t {
	if x == nil {
		return nil
	}
	return (*C.hs_stream_t)(unsafe.Pointer(x))
}

// Free cleanups the referenced memory using C free.
func (x *Stream) Free() {
	if x != nil {
		C.free(unsafe.Pointer(x))
	}
}

// NewStreamRef converts the C object reference into a raw struct reference without wrapping.
func NewStreamRef(ref unsafe.Pointer) *Stream {
	return (*Stream)(ref)
}

// NewStream allocates a new C object of this type and converts the reference into
// a raw struct reference without wrapping.
func NewStream() *Stream {
	return (*Stream)(allocStreamMemory(1))
}

// allocStreamMemory allocates memory for type C.hs_stream_t in C.
// The caller is responsible for freeing the this memory via C.free.
func allocStreamMemory(n int) unsafe.Pointer {
	mem, err := C.calloc(C.size_t(n), (C.size_t)(sizeOfStreamValue))
	if err != nil {
		panic("memory alloc error: " + err.Error())
	}
	return mem
}

const sizeOfStreamValue = unsafe.Sizeof([1]C.hs_stream_t{})

// PassRef returns a reference to C object as it is or allocates a new C object of this type.
func (x *Stream) PassRef() *C.hs_stream_t {
	if x == nil {
		x = (*Stream)(allocStreamMemory(1))
	}
	return (*C.hs_stream_t)(unsafe.Pointer(x))
}

// Ref returns a reference to C object as it is.
func (x *Scratch) Ref() *C.hs_scratch_t {
	if x == nil {
		return nil
	}
	return (*C.hs_scratch_t)(unsafe.Pointer(x))
}

// Free cleanups the referenced memory using C free.
func (x *Scratch) Free() {
	if x != nil {
		C.free(unsafe.Pointer(x))
	}
}

// NewScratchRef converts the C object reference into a raw struct reference without wrapping.
func NewScratchRef(ref unsafe.Pointer) *Scratch {
	return (*Scratch)(ref)
}

// NewScratch allocates a new C object of this type and converts the reference into
// a raw struct reference without wrapping.
func NewScratch() *Scratch {
	return (*Scratch)(allocScratchMemory(1))
}

// allocScratchMemory allocates memory for type C.hs_scratch_t in C.
// The caller is responsible for freeing the this memory via C.free.
func allocScratchMemory(n int) unsafe.Pointer {
	mem, err := C.calloc(C.size_t(n), (C.size_t)(sizeOfScratchValue))
	if err != nil {
		panic("memory alloc error: " + err.Error())
	}
	return mem
}

const sizeOfScratchValue = unsafe.Sizeof([1]C.hs_scratch_t{})

// PassRef returns a reference to C object as it is or allocates a new C object of this type.
func (x *Scratch) PassRef() *C.hs_scratch_t {
	if x == nil {
		x = (*Scratch)(allocScratchMemory(1))
	}
	return (*C.hs_scratch_t)(unsafe.Pointer(x))
}

func (x MatchEventHandler) PassRef() (ref *C.match_event_handler, allocs *cgoAllocMap) {
	if x == nil {
		return nil, nil
	}
	if matchEventHandler6C9F61FBFunc == nil {
		matchEventHandler6C9F61FBFunc = x
	}
	return (*C.match_event_handler)(C.match_event_handler_6c9f61fb), nil
}

func (x MatchEventHandler) PassValue() (ref C.match_event_handler, allocs *cgoAllocMap) {
	if x == nil {
		return nil, nil
	}
	if matchEventHandler6C9F61FBFunc == nil {
		matchEventHandler6C9F61FBFunc = x
	}
	return (C.match_event_handler)(C.match_event_handler_6c9f61fb), nil
}

func NewMatchEventHandlerRef(ref unsafe.Pointer) *MatchEventHandler {
	return (*MatchEventHandler)(ref)
}

//export matchEventHandler6C9F61FB
func matchEventHandler6C9F61FB(cid C.uint, cfrom C.ulonglong, cto C.ulonglong, cflags C.uint, ccontext unsafe.Pointer) C.int {
	if matchEventHandler6C9F61FBFunc != nil {
		id6c9f61fb := (uint32)(cid)
		from6c9f61fb := (uint64)(cfrom)
		to6c9f61fb := (uint64)(cto)
		flags6c9f61fb := (uint32)(cflags)
		context6c9f61fb := (unsafe.Pointer)(unsafe.Pointer(ccontext))
		ret6c9f61fb := matchEventHandler6C9F61FBFunc(id6c9f61fb, from6c9f61fb, to6c9f61fb, flags6c9f61fb, context6c9f61fb)
		ret, _ := (C.int)(ret6c9f61fb), cgoAllocsUnknown
		return ret
	}
	panic("callback func has not been set (race?)")
}

var matchEventHandler6C9F61FBFunc MatchEventHandler

// unpackPCharString represents the data from Go string as *C.char and avoids copying.
func unpackPCharString(str string) (*C.char, *cgoAllocMap) {
	h := (*stringHeader)(unsafe.Pointer(&str))
	return (*C.char)(unsafe.Pointer(h.Data)), cgoAllocsUnknown
}

type stringHeader struct {
	Data uintptr
	Len  int
}

const sizeOfPtr = unsafe.Sizeof(&struct{}{})

// unpackArgSPlatformInfo transforms a sliced Go data structure into plain C format.
func unpackArgSPlatformInfo(x []PlatformInfo) (unpacked *C.hs_platform_info_t, allocs *cgoAllocMap) {
	if x == nil {
		return nil, nil
	}
	allocs = new(cgoAllocMap)
	defer runtime.SetFinalizer(&unpacked, func(**C.hs_platform_info_t) {
		go allocs.Free()
	})

	len0 := len(x)
	mem0 := allocPlatformInfoMemory(len0)
	allocs.Add(mem0)
	h0 := &sliceHeader{
		Data: uintptr(mem0),
		Cap:  len0,
		Len:  len0,
	}
	v0 := *(*[]C.hs_platform_info_t)(unsafe.Pointer(h0))
	for i0 := range x {
		v0[i0], _ = x[i0].PassValue()
	}
	h := (*sliceHeader)(unsafe.Pointer(&v0))
	unpacked = (*C.hs_platform_info_t)(unsafe.Pointer(h.Data))
	return
}

// packSPlatformInfo reads sliced Go data structure out from plain C format.
func packSPlatformInfo(v []PlatformInfo, ptr0 *C.hs_platform_info_t) {
	const m = 0x7fffffff
	for i0 := range v {
		ptr1 := (*(*[m / sizeOfPlatformInfoValue]C.hs_platform_info_t)(unsafe.Pointer(ptr0)))[i0]
		v[i0] = *NewPlatformInfoRef(unsafe.Pointer(&ptr1))
	}
}

// allocPDatabaseMemory allocates memory for type *C.hs_database_t in C.
// The caller is responsible for freeing the this memory via C.free.
func allocPDatabaseMemory(n int) unsafe.Pointer {
	mem, err := C.calloc(C.size_t(n), (C.size_t)(sizeOfPDatabaseValue))
	if err != nil {
		panic("memory alloc error: " + err.Error())
	}
	return mem
}

const sizeOfPDatabaseValue = unsafe.Sizeof([1]*C.hs_database_t{})

// unpackArgSSDatabase transforms a sliced Go data structure into plain C format.
func unpackArgSSDatabase(x [][]Database) (unpacked **C.hs_database_t, allocs *cgoAllocMap) {
	if x == nil {
		return nil, nil
	}
	allocs = new(cgoAllocMap)
	defer runtime.SetFinalizer(&unpacked, func(***C.hs_database_t) {
		go allocs.Free()
	})

	len0 := len(x)
	mem0 := allocPDatabaseMemory(len0)
	allocs.Add(mem0)
	h0 := &sliceHeader{
		Data: uintptr(mem0),
		Cap:  len0,
		Len:  len0,
	}
	v0 := *(*[]*C.hs_database_t)(unsafe.Pointer(h0))
	for i0 := range x {
		h := (*sliceHeader)(unsafe.Pointer(&x[i0]))
		v0[i0] = (*C.hs_database_t)(unsafe.Pointer(h.Data))
	}
	h := (*sliceHeader)(unsafe.Pointer(&v0))
	unpacked = (**C.hs_database_t)(unsafe.Pointer(h.Data))
	return
}

// packSSDatabase reads sliced Go data structure out from plain C format.
func packSSDatabase(v [][]Database, ptr0 **C.hs_database_t) {
	const m = 0x7fffffff
	for i0 := range v {
		ptr1 := (*(*[m / sizeOfPtr]*C.hs_database_t)(unsafe.Pointer(ptr0)))[i0]
		hxf95e7c8 := (*sliceHeader)(unsafe.Pointer(&v[i0]))
		hxf95e7c8.Data = uintptr(unsafe.Pointer(ptr1))
		hxf95e7c8.Cap = 0x7fffffff
		// hxf95e7c8.Len = ?
	}
}

// allocPCompileErrorMemory allocates memory for type *C.hs_compile_error_t in C.
// The caller is responsible for freeing the this memory via C.free.
func allocPCompileErrorMemory(n int) unsafe.Pointer {
	mem, err := C.calloc(C.size_t(n), (C.size_t)(sizeOfPCompileErrorValue))
	if err != nil {
		panic("memory alloc error: " + err.Error())
	}
	return mem
}

const sizeOfPCompileErrorValue = unsafe.Sizeof([1]*C.hs_compile_error_t{})

// unpackArgSSCompileError transforms a sliced Go data structure into plain C format.
func unpackArgSSCompileError(x [][]CompileError) (unpacked **C.hs_compile_error_t, allocs *cgoAllocMap) {
	if x == nil {
		return nil, nil
	}
	allocs = new(cgoAllocMap)
	defer runtime.SetFinalizer(&unpacked, func(***C.hs_compile_error_t) {
		go allocs.Free()
	})

	len0 := len(x)
	mem0 := allocPCompileErrorMemory(len0)
	allocs.Add(mem0)
	h0 := &sliceHeader{
		Data: uintptr(mem0),
		Cap:  len0,
		Len:  len0,
	}
	v0 := *(*[]*C.hs_compile_error_t)(unsafe.Pointer(h0))
	for i0 := range x {
		len1 := len(x[i0])
		mem1 := allocCompileErrorMemory(len1)
		allocs.Add(mem1)
		h1 := &sliceHeader{
			Data: uintptr(mem1),
			Cap:  len1,
			Len:  len1,
		}
		v1 := *(*[]C.hs_compile_error_t)(unsafe.Pointer(h1))
		for i1 := range x[i0] {
			v1[i1], _ = x[i0][i1].PassValue()
		}
		h := (*sliceHeader)(unsafe.Pointer(&v1))
		v0[i0] = (*C.hs_compile_error_t)(unsafe.Pointer(h.Data))
	}
	h := (*sliceHeader)(unsafe.Pointer(&v0))
	unpacked = (**C.hs_compile_error_t)(unsafe.Pointer(h.Data))
	return
}

// packSSCompileError reads sliced Go data structure out from plain C format.
func packSSCompileError(v [][]CompileError, ptr0 **C.hs_compile_error_t) {
	const m = 0x7fffffff
	for i0 := range v {
		ptr1 := (*(*[m / sizeOfPtr]*C.hs_compile_error_t)(unsafe.Pointer(ptr0)))[i0]
		for i1 := range v[i0] {
			ptr2 := (*(*[m / sizeOfCompileErrorValue]C.hs_compile_error_t)(unsafe.Pointer(ptr1)))[i1]
			v[i0][i1] = *NewCompileErrorRef(unsafe.Pointer(&ptr2))
		}
	}
}

// allocPCharMemory allocates memory for type *C.char in C.
// The caller is responsible for freeing the this memory via C.free.
func allocPCharMemory(n int) unsafe.Pointer {
	mem, err := C.calloc(C.size_t(n), (C.size_t)(sizeOfPCharValue))
	if err != nil {
		panic("memory alloc error: " + err.Error())
	}
	return mem
}

const sizeOfPCharValue = unsafe.Sizeof([1]*C.char{})

// unpackArgSString transforms a sliced Go data structure into plain C format.
func unpackArgSString(x []string) (unpacked **C.char, allocs *cgoAllocMap) {
	if x == nil {
		return nil, nil
	}
	allocs = new(cgoAllocMap)
	defer runtime.SetFinalizer(&unpacked, func(***C.char) {
		go allocs.Free()
	})

	len0 := len(x)
	mem0 := allocPCharMemory(len0)
	allocs.Add(mem0)
	h0 := &sliceHeader{
		Data: uintptr(mem0),
		Cap:  len0,
		Len:  len0,
	}
	v0 := *(*[]*C.char)(unsafe.Pointer(h0))
	for i0 := range x {
		v0[i0], _ = unpackPCharString(x[i0])
	}
	h := (*sliceHeader)(unsafe.Pointer(&v0))
	unpacked = (**C.char)(unsafe.Pointer(h.Data))
	return
}

// packSString reads sliced Go data structure out from plain C format.
func packSString(v []string, ptr0 **C.char) {
	const m = 0x7fffffff
	for i0 := range v {
		ptr1 := (*(*[m / sizeOfPtr]*C.char)(unsafe.Pointer(ptr0)))[i0]
		v[i0] = packPCharString(ptr1)
	}
}

// packPCharString creates a Go string backed by *C.char and avoids copying.
func packPCharString(p *C.char) (raw string) {
	if p != nil && *p != 0 {
		h := (*stringHeader)(unsafe.Pointer(&raw))
		h.Data = uintptr(unsafe.Pointer(p))
		for *p != 0 {
			p = (*C.char)(unsafe.Pointer(uintptr(unsafe.Pointer(p)) + 1)) // p++
		}
		h.Len = int(uintptr(unsafe.Pointer(p)) - h.Data)
	}
	return
}

// RawString reperesents a string backed by data on the C side.
type RawString string

// Copy returns a Go-managed copy of raw string.
func (raw RawString) Copy() string {
	if len(raw) == 0 {
		return ""
	}
	h := (*stringHeader)(unsafe.Pointer(&raw))
	return C.GoStringN((*C.char)(unsafe.Pointer(h.Data)), C.int(h.Len))
}

// allocPExprExtMemory allocates memory for type *C.hs_expr_ext_t in C.
// The caller is responsible for freeing the this memory via C.free.
func allocPExprExtMemory(n int) unsafe.Pointer {
	mem, err := C.calloc(C.size_t(n), (C.size_t)(sizeOfPExprExtValue))
	if err != nil {
		panic("memory alloc error: " + err.Error())
	}
	return mem
}

const sizeOfPExprExtValue = unsafe.Sizeof([1]*C.hs_expr_ext_t{})

// unpackArgSSExprExt transforms a sliced Go data structure into plain C format.
func unpackArgSSExprExt(x [][]ExprExt) (unpacked **C.hs_expr_ext_t, allocs *cgoAllocMap) {
	if x == nil {
		return nil, nil
	}
	allocs = new(cgoAllocMap)
	defer runtime.SetFinalizer(&unpacked, func(***C.hs_expr_ext_t) {
		go allocs.Free()
	})

	len0 := len(x)
	mem0 := allocPExprExtMemory(len0)
	allocs.Add(mem0)
	h0 := &sliceHeader{
		Data: uintptr(mem0),
		Cap:  len0,
		Len:  len0,
	}
	v0 := *(*[]*C.hs_expr_ext_t)(unsafe.Pointer(h0))
	for i0 := range x {
		len1 := len(x[i0])
		mem1 := allocExprExtMemory(len1)
		allocs.Add(mem1)
		h1 := &sliceHeader{
			Data: uintptr(mem1),
			Cap:  len1,
			Len:  len1,
		}
		v1 := *(*[]C.hs_expr_ext_t)(unsafe.Pointer(h1))
		for i1 := range x[i0] {
			v1[i1], _ = x[i0][i1].PassValue()
		}
		h := (*sliceHeader)(unsafe.Pointer(&v1))
		v0[i0] = (*C.hs_expr_ext_t)(unsafe.Pointer(h.Data))
	}
	h := (*sliceHeader)(unsafe.Pointer(&v0))
	unpacked = (**C.hs_expr_ext_t)(unsafe.Pointer(h.Data))
	return
}

// packSSExprExt reads sliced Go data structure out from plain C format.
func packSSExprExt(v [][]ExprExt, ptr0 **C.hs_expr_ext_t) {
	const m = 0x7fffffff
	for i0 := range v {
		ptr1 := (*(*[m / sizeOfPtr]*C.hs_expr_ext_t)(unsafe.Pointer(ptr0)))[i0]
		for i1 := range v[i0] {
			ptr2 := (*(*[m / sizeOfExprExtValue]C.hs_expr_ext_t)(unsafe.Pointer(ptr1)))[i1]
			v[i0][i1] = *NewExprExtRef(unsafe.Pointer(&ptr2))
		}
	}
}

// unpackArgSCompileError transforms a sliced Go data structure into plain C format.
func unpackArgSCompileError(x []CompileError) (unpacked *C.hs_compile_error_t, allocs *cgoAllocMap) {
	if x == nil {
		return nil, nil
	}
	allocs = new(cgoAllocMap)
	defer runtime.SetFinalizer(&unpacked, func(**C.hs_compile_error_t) {
		go allocs.Free()
	})

	len0 := len(x)
	mem0 := allocCompileErrorMemory(len0)
	allocs.Add(mem0)
	h0 := &sliceHeader{
		Data: uintptr(mem0),
		Cap:  len0,
		Len:  len0,
	}
	v0 := *(*[]C.hs_compile_error_t)(unsafe.Pointer(h0))
	for i0 := range x {
		v0[i0], _ = x[i0].PassValue()
	}
	h := (*sliceHeader)(unsafe.Pointer(&v0))
	unpacked = (*C.hs_compile_error_t)(unsafe.Pointer(h.Data))
	return
}

// packSCompileError reads sliced Go data structure out from plain C format.
func packSCompileError(v []CompileError, ptr0 *C.hs_compile_error_t) {
	const m = 0x7fffffff
	for i0 := range v {
		ptr1 := (*(*[m / sizeOfCompileErrorValue]C.hs_compile_error_t)(unsafe.Pointer(ptr0)))[i0]
		v[i0] = *NewCompileErrorRef(unsafe.Pointer(&ptr1))
	}
}

// allocPExprInfoMemory allocates memory for type *C.hs_expr_info_t in C.
// The caller is responsible for freeing the this memory via C.free.
func allocPExprInfoMemory(n int) unsafe.Pointer {
	mem, err := C.calloc(C.size_t(n), (C.size_t)(sizeOfPExprInfoValue))
	if err != nil {
		panic("memory alloc error: " + err.Error())
	}
	return mem
}

const sizeOfPExprInfoValue = unsafe.Sizeof([1]*C.hs_expr_info_t{})

// unpackArgSSExprInfo transforms a sliced Go data structure into plain C format.
func unpackArgSSExprInfo(x [][]ExprInfo) (unpacked **C.hs_expr_info_t, allocs *cgoAllocMap) {
	if x == nil {
		return nil, nil
	}
	allocs = new(cgoAllocMap)
	defer runtime.SetFinalizer(&unpacked, func(***C.hs_expr_info_t) {
		go allocs.Free()
	})

	len0 := len(x)
	mem0 := allocPExprInfoMemory(len0)
	allocs.Add(mem0)
	h0 := &sliceHeader{
		Data: uintptr(mem0),
		Cap:  len0,
		Len:  len0,
	}
	v0 := *(*[]*C.hs_expr_info_t)(unsafe.Pointer(h0))
	for i0 := range x {
		len1 := len(x[i0])
		mem1 := allocExprInfoMemory(len1)
		allocs.Add(mem1)
		h1 := &sliceHeader{
			Data: uintptr(mem1),
			Cap:  len1,
			Len:  len1,
		}
		v1 := *(*[]C.hs_expr_info_t)(unsafe.Pointer(h1))
		for i1 := range x[i0] {
			v1[i1], _ = x[i0][i1].PassValue()
		}
		h := (*sliceHeader)(unsafe.Pointer(&v1))
		v0[i0] = (*C.hs_expr_info_t)(unsafe.Pointer(h.Data))
	}
	h := (*sliceHeader)(unsafe.Pointer(&v0))
	unpacked = (**C.hs_expr_info_t)(unsafe.Pointer(h.Data))
	return
}

// packSSExprInfo reads sliced Go data structure out from plain C format.
func packSSExprInfo(v [][]ExprInfo, ptr0 **C.hs_expr_info_t) {
	const m = 0x7fffffff
	for i0 := range v {
		ptr1 := (*(*[m / sizeOfPtr]*C.hs_expr_info_t)(unsafe.Pointer(ptr0)))[i0]
		for i1 := range v[i0] {
			ptr2 := (*(*[m / sizeOfExprInfoValue]C.hs_expr_info_t)(unsafe.Pointer(ptr1)))[i1]
			v[i0][i1] = *NewExprInfoRef(unsafe.Pointer(&ptr2))
		}
	}
}

// unpackArgSExprExt transforms a sliced Go data structure into plain C format.
func unpackArgSExprExt(x []ExprExt) (unpacked *C.hs_expr_ext_t, allocs *cgoAllocMap) {
	if x == nil {
		return nil, nil
	}
	allocs = new(cgoAllocMap)
	defer runtime.SetFinalizer(&unpacked, func(**C.hs_expr_ext_t) {
		go allocs.Free()
	})

	len0 := len(x)
	mem0 := allocExprExtMemory(len0)
	allocs.Add(mem0)
	h0 := &sliceHeader{
		Data: uintptr(mem0),
		Cap:  len0,
		Len:  len0,
	}
	v0 := *(*[]C.hs_expr_ext_t)(unsafe.Pointer(h0))
	for i0 := range x {
		v0[i0], _ = x[i0].PassValue()
	}
	h := (*sliceHeader)(unsafe.Pointer(&v0))
	unpacked = (*C.hs_expr_ext_t)(unsafe.Pointer(h.Data))
	return
}

// packSExprExt reads sliced Go data structure out from plain C format.
func packSExprExt(v []ExprExt, ptr0 *C.hs_expr_ext_t) {
	const m = 0x7fffffff
	for i0 := range v {
		ptr1 := (*(*[m / sizeOfExprExtValue]C.hs_expr_ext_t)(unsafe.Pointer(ptr0)))[i0]
		v[i0] = *NewExprExtRef(unsafe.Pointer(&ptr1))
	}
}

// unpackArgSSByte transforms a sliced Go data structure into plain C format.
func unpackArgSSByte(x [][]byte) (unpacked **C.char, allocs *cgoAllocMap) {
	if x == nil {
		return nil, nil
	}
	allocs = new(cgoAllocMap)
	defer runtime.SetFinalizer(&unpacked, func(***C.char) {
		go allocs.Free()
	})

	len0 := len(x)
	mem0 := allocPCharMemory(len0)
	allocs.Add(mem0)
	h0 := &sliceHeader{
		Data: uintptr(mem0),
		Cap:  len0,
		Len:  len0,
	}
	v0 := *(*[]*C.char)(unsafe.Pointer(h0))
	for i0 := range x {
		h := (*sliceHeader)(unsafe.Pointer(&x[i0]))
		v0[i0] = (*C.char)(unsafe.Pointer(h.Data))
	}
	h := (*sliceHeader)(unsafe.Pointer(&v0))
	unpacked = (**C.char)(unsafe.Pointer(h.Data))
	return
}

// packSSByte reads sliced Go data structure out from plain C format.
func packSSByte(v [][]byte, ptr0 **C.char) {
	const m = 0x7fffffff
	for i0 := range v {
		ptr1 := (*(*[m / sizeOfPtr]*C.char)(unsafe.Pointer(ptr0)))[i0]
		hxfa9955c := (*sliceHeader)(unsafe.Pointer(&v[i0]))
		hxfa9955c.Data = uintptr(unsafe.Pointer(ptr1))
		hxfa9955c.Cap = 0x7fffffff
		// hxfa9955c.Len = ?
	}
}

// allocPStreamMemory allocates memory for type *C.hs_stream_t in C.
// The caller is responsible for freeing the this memory via C.free.
func allocPStreamMemory(n int) unsafe.Pointer {
	mem, err := C.calloc(C.size_t(n), (C.size_t)(sizeOfPStreamValue))
	if err != nil {
		panic("memory alloc error: " + err.Error())
	}
	return mem
}

const sizeOfPStreamValue = unsafe.Sizeof([1]*C.hs_stream_t{})

// unpackArgSSStream transforms a sliced Go data structure into plain C format.
func unpackArgSSStream(x [][]Stream) (unpacked **C.hs_stream_t, allocs *cgoAllocMap) {
	if x == nil {
		return nil, nil
	}
	allocs = new(cgoAllocMap)
	defer runtime.SetFinalizer(&unpacked, func(***C.hs_stream_t) {
		go allocs.Free()
	})

	len0 := len(x)
	mem0 := allocPStreamMemory(len0)
	allocs.Add(mem0)
	h0 := &sliceHeader{
		Data: uintptr(mem0),
		Cap:  len0,
		Len:  len0,
	}
	v0 := *(*[]*C.hs_stream_t)(unsafe.Pointer(h0))
	for i0 := range x {
		h := (*sliceHeader)(unsafe.Pointer(&x[i0]))
		v0[i0] = (*C.hs_stream_t)(unsafe.Pointer(h.Data))
	}
	h := (*sliceHeader)(unsafe.Pointer(&v0))
	unpacked = (**C.hs_stream_t)(unsafe.Pointer(h.Data))
	return
}

// packSSStream reads sliced Go data structure out from plain C format.
func packSSStream(v [][]Stream, ptr0 **C.hs_stream_t) {
	const m = 0x7fffffff
	for i0 := range v {
		ptr1 := (*(*[m / sizeOfPtr]*C.hs_stream_t)(unsafe.Pointer(ptr0)))[i0]
		hxf69fe70 := (*sliceHeader)(unsafe.Pointer(&v[i0]))
		hxf69fe70.Data = uintptr(unsafe.Pointer(ptr1))
		hxf69fe70.Cap = 0x7fffffff
		// hxf69fe70.Len = ?
	}
}

// allocPScratchMemory allocates memory for type *C.hs_scratch_t in C.
// The caller is responsible for freeing the this memory via C.free.
func allocPScratchMemory(n int) unsafe.Pointer {
	mem, err := C.calloc(C.size_t(n), (C.size_t)(sizeOfPScratchValue))
	if err != nil {
		panic("memory alloc error: " + err.Error())
	}
	return mem
}

const sizeOfPScratchValue = unsafe.Sizeof([1]*C.hs_scratch_t{})

// unpackArgSSScratch transforms a sliced Go data structure into plain C format.
func unpackArgSSScratch(x [][]Scratch) (unpacked **C.hs_scratch_t, allocs *cgoAllocMap) {
	if x == nil {
		return nil, nil
	}
	allocs = new(cgoAllocMap)
	defer runtime.SetFinalizer(&unpacked, func(***C.hs_scratch_t) {
		go allocs.Free()
	})

	len0 := len(x)
	mem0 := allocPScratchMemory(len0)
	allocs.Add(mem0)
	h0 := &sliceHeader{
		Data: uintptr(mem0),
		Cap:  len0,
		Len:  len0,
	}
	v0 := *(*[]*C.hs_scratch_t)(unsafe.Pointer(h0))
	for i0 := range x {
		h := (*sliceHeader)(unsafe.Pointer(&x[i0]))
		v0[i0] = (*C.hs_scratch_t)(unsafe.Pointer(h.Data))
	}
	h := (*sliceHeader)(unsafe.Pointer(&v0))
	unpacked = (**C.hs_scratch_t)(unsafe.Pointer(h.Data))
	return
}

// packSSScratch reads sliced Go data structure out from plain C format.
func packSSScratch(v [][]Scratch, ptr0 **C.hs_scratch_t) {
	const m = 0x7fffffff
	for i0 := range v {
		ptr1 := (*(*[m / sizeOfPtr]*C.hs_scratch_t)(unsafe.Pointer(ptr0)))[i0]
		hxf3b8dbd := (*sliceHeader)(unsafe.Pointer(&v[i0]))
		hxf3b8dbd.Data = uintptr(unsafe.Pointer(ptr1))
		hxf3b8dbd.Cap = 0x7fffffff
		// hxf3b8dbd.Len = ?
	}
}
