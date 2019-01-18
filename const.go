// Copyright 2019 Barchart.com, Inc. All rights reserved.
// Use of this source code is governed by a GNU-style
// license that can be found in the LICENSE file.

/*
	Package btrieve contains constants and function that
	help in accessing Btrieve databases.
*/

package btrieve

type OP_CODE uint16

const (
	B_OPEN      OP_CODE = 0
	B_CLOSE     OP_CODE = 1
	B_GET_FIRST OP_CODE = 12
	B_GET_NEXT  OP_CODE = 6
)
