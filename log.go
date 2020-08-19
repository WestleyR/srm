// Created by: WestleyR
// Email: westleyr@nym.hush.com
// Url: https://github.com/WestleyR/srm
// Last modified date: 2020-08-18
//
// This file is licensed under the terms of
//
// The Clear BSD License
//
// Copyright (c) 2019-2020 WestleyR
// All rights reserved.
//
// This software is licensed under a Clear BSD License.
//

// TODO: use a better logging library

package main

import (
	"fmt"
)

var debug = false

func setDebug(value bool) {
	debug = value
}

func isDebug() bool {
	return debug
}

func printDebugf(format string) {
	if debug {
		fmt.Println(format)
	}
}
