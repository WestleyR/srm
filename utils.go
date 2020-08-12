// Created by: WestleyR
// Email: westleyr@nym.hush.com
// Url: https://github.com/WestleyR/srm
// Last modified date: 2020-08-11
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

package main

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

