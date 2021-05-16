//
//  recover.go
//  srm - https://github.com/WestleyR/srm
//
// Created by WestleyR <westleyr@nym.hush.com> on Aug 11, 2020
// Source code: https://github.com/WestleyR/srm
//
// Copyright (c) 2020-2021 WestleyR. All rights reserved.
// This software is licensed under a BSD 3-Clause Clear License.
// Consult the LICENSE file that came with this software regarding
// your rights to distribute this software.
//

package srm

import (
	"os"
)

func isPathADir(path string) bool {
	fi, err := os.Stat(path)
	if err != nil {
		return false
	}
	switch mode := fi.Mode(); {
	case mode.IsDir():
		// path is a directory
		return true
	}
	return false
}
