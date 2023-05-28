// Created by WestleyR <westleyr@nym.hush.com> on July 28, 2020
// Source code: https://github.com/WestleyR/srm
//
// Copyright (c) 2020-2022 WestleyR. All rights reserved.
// This software is licensed under a BSD 3-Clause Clear License.
// Consult the LICENSE file that came with this software regarding
// your rights to distribute this software.
//

package srm

import (
	"errors"
	"fmt"
	"os"
	"path/filepath"

	"github.com/c2h5oh/datasize"
	"golang.org/x/sys/unix"
)

func doesPathExist(path string) bool {
	err := unix.Access(path, unix.F_OK)
	if err != nil && errors.Is(err, os.ErrNotExist) {
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

func isPathADir(path string) bool {
	fi, err := os.Lstat(path)
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

func getFileSize(path string) datasize.ByteSize {
	fi, err := os.Lstat(path)
	if err != nil {
		return datasize.ByteSize(0)
	}

	return datasize.ByteSize(fi.Size())
}

func getDirSize(path string) datasize.ByteSize {
	var size datasize.ByteSize

	err := filepath.Walk(path, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if !info.IsDir() {
			size += datasize.ByteSize(info.Size())
		}
		return nil
	})
	if err != nil {
		size = 0
	}

	return size
}

func checkIfFileIsWriteProtected(file string) bool {
	err := unix.Access(file, unix.W_OK)
	return !os.IsPermission(err)
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
