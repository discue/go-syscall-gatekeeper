// Copyright 2018 Google LLC.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

//go:build (linux && arm64) || (linux && amd64) || (linux && riscv64)

package uroot

import (
	"fmt"
	"os"
	"unsafe"
)

// Addr is an address for use in strace I/O
type Addr uintptr

// SyscallArgument is an argument supplied to a syscall implementation. The
// methods used to access the arguments are named after the ***C type name*** and
// they convert to the closest Go type available. For example, Int() refers to a
// 32-bit signed integer argument represented in Go as an int32.
//
// Using the accessor methods guarantees that the conversion between types is
// correct, taking into account size and signedness (i.e., zero-extension vs
// signed-extension).
type SyscallArgument struct {
	// Prefer to use accessor methods instead of 'Value' directly.
	Value uintptr
}

// SyscallArguments represents the set of arguments passed to a syscall.
type SyscallArguments [6]SyscallArgument

// Pointer returns the usermem.Addr representation of a pointer argument.
func (a SyscallArgument) Pointer() Addr {
	return Addr(a.Value)
}

func (a SyscallArgument) String() string {
	strPtr := (*string)(unsafe.Pointer(a.Value))
	return fmt.Sprint(strPtr)
}

// Int returns the int32 representation of a 32-bit signed integer argument.
func (a SyscallArgument) Int() int32 {
	return int32(a.Value)
}

// Uint returns the uint32 representation of a 32-bit unsigned integer argument.
func (a SyscallArgument) Uint() uint32 {
	return uint32(a.Value)
}

// Int64 returns the int64 representation of a 64-bit signed integer argument.
func (a SyscallArgument) Int64() int64 {
	return int64(a.Value)
}

// Uint64 returns the uint64 representation of a 64-bit unsigned integer argument.
func (a SyscallArgument) Uint64() uint64 {
	return uint64(a.Value)
}

// SizeT returns the uint representation of a size_t argument.
func (a SyscallArgument) SizeT() uint {
	return uint(a.Value)
}

// ModeT returns the int representation of a mode_t argument.
func (a SyscallArgument) ModeT() uint {
	return uint(uint16(a.Value))
}

// ModeT returns the int representation of a mode_t argument.
func (a SyscallArgument) Path() string {
	filePath := fmt.Sprintf("/proc/self/fd/%d", a.Int())
	realPath, err := os.Readlink(filePath)
	if err != nil {
		return "<unable to read file descriptor"
	}
	return realPath
}
