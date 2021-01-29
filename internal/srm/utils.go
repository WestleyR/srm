// Created by WestleyR on Aug 11, 2020
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
