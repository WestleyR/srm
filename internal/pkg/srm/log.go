//
//  log.go
//  srm - https://github.com/WestleyR/srm
//
// Created by WestleyR <westleyr@nym.hush.com> on July 28, 2020
// Source code: https://github.com/WestleyR/srm
//
// Copyright (c) 2020-2021 WestleyR. All rights reserved.
// This software is licensed under a BSD 3-Clause Clear License.
// Consult the LICENSE file that came with this software regarding
// your rights to distribute this software.
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
