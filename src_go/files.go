// Created by: WestleyR
// Email: westleyr@nym.hush.com
// Url: https://github.com/WestleyR/srm
// Last modified date: 2020-07-28
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
	"fmt"
	"os"
	"path/filepath"
)

type srmOptions struct {
	force     bool
	recursive bool

	// filePath not currently used
	filePath string
}

func isPathADirectory(path string) bool {
	fi, err := os.Stat(path)
	if err != nil {
		fmt.Println(err)
		return false
	}
	switch mode := fi.Mode(); {
	case mode.IsDir():
		return true
	case mode.IsRegular():
		return false
	}
	return false
}

func srmFile(filePath string, options srmOptions) error {
	trashPath := getFileTrashPath(filepath.Base(filePath))

	if isPathADirectory(filePath) {
		// Is a directory

		if !options.recursive {
			return fmt.Errorf("%s: is a directory", filePath)
		}

		// Move the file to srm trash
		err := os.Rename(filePath, trashPath)
		if err != nil {
			return err
		}
	} else {
		// Its a plain file

		// Move the file to srm trash
		err := os.Rename(filePath, trashPath)
		if err != nil {
			return err
		}
	}

	return nil
}
