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

	"golang.org/x/sys/unix"
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

func checkIfFileIsWriteProtected(file string) bool {
	return unix.Access(file, unix.W_OK) == nil
}

func checkForWriteProtectedFileIn(path string) error {
	err := filepath.Walk(path, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if !checkIfFileIsWriteProtected(path) {
			return fmt.Errorf("%s: is write protected", path)
		}
		return nil
	})

	return err
}

func srmFile(filePath string, options srmOptions) error {
	trashPath := getFileTrashPath(filepath.Base(filePath))

	if isPathADirectory(filePath) {
		// Is a directory

		if !options.recursive {
			return fmt.Errorf("%s: is a directory", filePath)
		}

		if !options.force {
			err := checkForWriteProtectedFileIn(filePath)
			if err != nil {
				return err
			}
		}

		// Move the file to srm trash
		err := os.Rename(filePath, trashPath)
		if err != nil {
			return err
		}
	} else {
		// Its a plain file

		if !options.force {
      if !checkIfFileIsWriteProtected(filePath) {
  			return fmt.Errorf("%s: is write protected", filePath)
      }
    }

		// Move the file to srm trash
		err := os.Rename(filePath, trashPath)
		if err != nil {
			return err
		}
	}

	return nil
}
