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
	"os/exec"
	"path/filepath"
	"strconv"

	"github.com/WestleyR/srm/internal/pkg/paths"
)

type Manager struct {
	trashPath  string
	trashIndex int32

	options   *Options
	DataCache *DataCache // TODO: not currently used, maybe later

	dataCacheFile string
}

type DataCache struct {
	TotalBytes int64 `ini:"total_bytes"`
}

// Options is the remove options for the file
type Options struct {
	Force          bool
	Recursive      bool
	DryRun         bool // TODO: not used yet
	RemoveEmptyDir bool // TODO: not used yet
}

func New(options *Options) (*Manager, error) {
	m := &Manager{}

	m.options = options

	var err error
	m.trashIndex, err = getNextTrashIndex()
	if err != nil {
		return nil, err
	}

	// ie. /home/user/.cache/srm/trash/2022/05
	m.trashPath = paths.GetTrashDirPath()

	m.dataCacheFile = paths.GetDataCachePath()

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

	// Is a directory
	if isPathADir(filePath) && !m.options.Recursive {
		return fmt.Errorf("%s: is a directory", filePath)
	}

	if !m.options.Force && !checkIfFileIsWriteProtected(filePath) {
		return fmt.Errorf("%s: is write protected", filePath)
	}

	// Move the file to srm trash
	err = os.Rename(filePath, trashPath)
	if err != nil {
		cmd := exec.Command("mv", filePath, trashPath)
		err = cmd.Run()
		if err != nil {
			return fmt.Errorf("%s: failed to move to trash directory %s", filePath, trashPath)
		}
	}

	m.trashIndex++

	return nil
}
