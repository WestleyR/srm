// Created by WestleyR on July 28, 2020
// Source code: https://github.com/WestleyR/srm
// Last modified data: 2021-01-28
//
// This file is licensed under the terms of
//
// The Clear BSD License
//
// Copyright (c) 2020-2021 WestleyR
// All rights reserved.
//

// TODO: use a better logging library

package srm

import (
	"fmt"
)

var debug = false

func SetDebug(value bool) {
	debug = value
}

func IsDebug() bool {
	return debug
}

func PrintDebugf(format string) {
	if debug {
		fmt.Println(format)
	}
}
