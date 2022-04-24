//
//  files.go
//  srm - https://github.com/WestleyR/srm
//
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
	"fmt"
	"os"
	"path/filepath"
	"strconv"

	"golang.org/x/sys/unix"
)

type Manager struct {
	trashPath  string
	trashIndex int32

	options *Options
}

// Options is the remove options for the file
type Options struct {
	Force     bool
	Recursive bool
	DryRun    bool // TODO: not used yet
}

func New(options *Options) (*Manager, error) {
	m := &Manager{}

	m.options = options

	var err error
	m.trashIndex, err = getNextTrashIndex()
	if err != nil {
		return nil, err
	}

	// ie. /home/user/.cache/srm/trash
	m.trashPath = getCachePath()

	return m, nil
}

func (m *Manager) RM(filePath string) error {
	if !doesPathExist(filePath) {
		return fmt.Errorf("%s: does not exist", filePath)
	}

	trashPath := filepath.Join(m.trashPath, strconv.FormatInt(int64(m.trashIndex), 10), filepath.Base(filePath))

	err := os.MkdirAll(filepath.Dir(trashPath), 0700)
	if err != nil {
		return fmt.Errorf("failed to create trash dir: %s", err)
	}

	if isPathADirectory(filePath) {
		// Is a directory

		if !m.options.Recursive {
			return fmt.Errorf("%s: is a directory", filePath)
		}

		if !m.options.Force {
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

		if !m.options.Force {
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
	m.trashIndex++

	return nil
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
