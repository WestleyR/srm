// Created by: WestleyR
// Email: westleyr@nym.hush.com
// Url: https://github.com/WestleyR/srm
// Last modified date: 2021-01-09
//
// This file is licensed under the terms of
//
// The Clear BSD License
//
// Copyright (c) 2019-2021 WestleyR
// All rights reserved.
//
// This software is licensed under a Clear BSD License.
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
