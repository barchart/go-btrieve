// Copyright 2019 Barchart.com, Inc. All rights reserved.
// Use of this source code is governed by a GNU-style
// license that can be found in the LICENSE file.

/*
	Package btrieve contains constants and function that
	help in accessing Btrieve databases.
*/

package btrieve

import (
	"syscall"
	"unsafe"
)

var (
	btrcall uintptr
	btrvdll syscall.Handle
)

// Init loads the dll and finds the pointer to the
// BTRCALL function. Call this once before calling
// BTRV.
func Init() error {
	var err error
	btrvdll, err = syscall.LoadLibrary("WBTRV32.DLL")
	if err != nil {
		return err
	}

	btrcall, err = syscall.GetProcAddress(btrvdll, "BTRCALL@28")
	if err != nil {
		syscall.FreeLibrary(btrvdll)
		return err
	}

	return nil
}

// BTRV is the golang equivalent of the BTRV call in the
// Btrieve C API.
func BTRV(op OP_CODE, posBlock []uint16, dataBuffer []byte, dataLength uint32, keyBuffer []byte, keyNumber int8) (uint16, error) {
	if btrcall == 0 {
		err := Init()
		if err != nil {
			return 0, err
		}
	}

	var keylen = MAX_KEY_SIZE

	status, _, _ := syscall.Syscall9(
		btrcall,
		7,
		uintptr(op),
		uintptr(unsafe.Pointer(&posBlock[0])),
		uintptr(unsafe.Pointer(&dataBuffer[0])),
		uintptr(unsafe.Pointer(&dataLength)),
		uintptr(unsafe.Pointer(&keyBuffer[0])),
		uintptr(unsafe.Pointer(&keylen)),
		uintptr(keyNumber),
		0,
		0)

	return uint16(status), nil
}

// Release releases the dll and the resources associtated
// with the dll. Call this once at the end of your program.
func Release() {
	syscall.FreeLibrary(btrvdll)
}
