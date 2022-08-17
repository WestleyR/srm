//
//  cache.go
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
	"io/ioutil"
	"os"
	"path/filepath"
	"strconv"

	"github.com/WestleyR/srm/internal/pkg/paths"
	"github.com/wildwest-productions/goini"
)

// TODO: the following two functions are not used. This may be used later...

func (m *Manager) loadDataCache() error {
	dc := &DataCache{}

	b, err := os.ReadFile(m.dataCacheFile)
	if err != nil {
		// TODO: need to check os.IsNotExist
		return nil
	}

	err = goini.Unmarshal(b, &dc)
	if err != nil {
		return err
	}

	m.DataCache = dc

	return nil
}

func (m *Manager) Write(path string) error {
	b, err := goini.Marshal(m.DataCache)
	if err != nil {
		return fmt.Errorf("failed to marshal ini: %s", err)
	}

	err = os.WriteFile(path, b, 0700)
	if err != nil {
		return fmt.Errorf("failed to write file: %s", err)
	}

	return nil
}

func getNextTrashIndex() (int32, error) {
	cachePath := paths.GetTrashDirPath()

	cacheNumber := int32(0)

	files, err := ioutil.ReadDir(cachePath)
	if err != nil {
		return 0, fmt.Errorf("failed to read dir: %s", err)
	}
	cacheNumber = int32(len(files))

	// Give the extra one
	cacheNumber++

	fullPath := filepath.Join(cachePath, strconv.FormatInt(int64(cacheNumber), 10))

	// Make sure it does not already exist, sometimes
	// when one of the cache dirs is removed, it can
	// screw up the last number. Only try 100 times
	for true {
		if doesPathExist(fullPath) {
			// File already exists, then add incremt again...
			cacheNumber++
			fullPath = filepath.Join(cachePath, strconv.FormatInt(int64(cacheNumber), 10))
		} else {
			break
		}
	}

	return cacheNumber, nil
}

func InitCache() {
	cachePath := paths.GetTrashDirPath()

	err := os.MkdirAll(cachePath, 0700)
	if err != nil {
		panic(fmt.Sprintf("failed to create dir: %s", err))
	}
}
