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

package srm

import (
	"fmt"
	"os"
	"path/filepath"

	"golang.org/x/sys/unix"
)

type SrmOptions struct {
	Force     bool
	Recursive bool

	// filePath not currently used
	FilePath string
}

func doesPathExist(path string) bool {
	err := unix.Access(path, unix.F_OK)
	if err != nil && !os.IsNotExist(err) {
		return false
	}

	if os.IsNotExist(err) {
		_, err := os.Lstat(path)
		if err != nil {
			return false
		}
	}
	return true
}

func isPathADirectory(path string) bool {
	fi, err := os.Stat(path)
	if err != nil {
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
	err := unix.Access(file, unix.W_OK)
	if os.IsPermission(err) {
		return false
	}
	return true
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

func SrmFile(filePath string, options SrmOptions) error {
	if !doesPathExist(filePath) {
		return fmt.Errorf("%s: does not exist", filePath)
	}

	trashPath := getFileTrashPath(filepath.Base(filePath))

	if isPathADirectory(filePath) {
		// Is a directory

		if !options.Recursive {
			return fmt.Errorf("%s: is a directory", filePath)
		}

		if !options.Force {
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

		if !options.Force {
			if !checkIfFileIsWriteProtected(filePath) {
				return fmt.Errorf("%s: is write protected", filePath)
			}
		}

		// Move the file to srm trash
		err := os.Rename(filePath, trashPath)
		if err != nil && !os.IsNotExist(err) {
			return err
		}
	}

	return nil
}
