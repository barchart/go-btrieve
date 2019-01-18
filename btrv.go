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

var keylen = uint8(255)

type DB struct {
	btrcall uintptr
	btrvdll syscall.Handle
}

// BTRV is the golang equivalent of the BTRV call in the
// Btrieve C API.
func (db *DB) BTRV(op OP_CODE, posBlock []uint16, dataBuffer []byte, dataLength uint32, keyBuffer []byte, keyNumber int8) uint16 {
	status, _, _ := syscall.Syscall9(
		db.btrcall,
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

	return uint16(status)
}

// InitDB is called to initalize the dll and create
// a placeholder for the db structure.
func InitDB() (*DB, error) {
	var (
		db  DB
		err error
	)

	db.btrvdll, err = syscall.LoadLibrary("WBTRV32.DLL")
	if err != nil {
		return nil, err
	}

	db.btrcall, err = syscall.GetProcAddress(db.btrvdll, "BTRCALL@28")
	if err != nil {
		syscall.FreeLibrary(db.btrvdll)
		return nil, err
	}

	return &db, nil
}

// ReleaseDB releases the dll and the resources associtated
// with the dll.
func ReleaseDB(db *DB) {
	syscall.FreeLibrary(db.btrvdll)
}
