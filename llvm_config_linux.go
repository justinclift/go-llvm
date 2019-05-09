// +build !byollvm

package llvm

// Automatically generated by `make config BUILDDIR=`, do not edit.

// #cgo CPPFLAGS: -I/usr/lib/llvm-9/include -D_GNU_SOURCE -D__STDC_CONSTANT_MACROS -D__STDC_FORMAT_MACROS -D__STDC_LIMIT_MACROS
// #cgo CXXFLAGS: -std=c++11
// #cgo LDFLAGS: -L/usr/lib/llvm-9/lib  -lLLVM-9
import "C"

type (
	run_build_sh int
)
