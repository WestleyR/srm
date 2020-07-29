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

func srmFile(filePath string) error {
	trashPath := getFileTrashPath(filepath.Base(filePath))

	if isPathADirectory(filePath) {
		// Copy the directory
		err := CopyDir(filePath, trashPath)
		if err != nil {
			return err
		}

		err = os.RemoveAll(filePath)
		if err != nil {
			return err
		}
	} else {
		// Its a plain file
		err := CopyFile(filePath, trashPath)
		if err != nil {
			return err
		}

		err = os.Remove(filePath)
		if err != nil {
			return err
		}
	}

	return nil
}
