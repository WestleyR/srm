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
)

const AutocleanSizeLimit = 15024 * 1024 // 15Mbs
const AutocleanSizeLower = 100          // 100b

// CleanCacheAUTO will remove all cached items above autocleanSizeLimit,
// and below autocleanSizeLower.
func CleanCacheAUTO(dryRun bool) error {
	path := getCachePath()

	cache, _ := getCacheArray(path)

	var totalCleanedMbs float32

	for _, c := range cache {
		if c.size > AutocleanSizeLimit || c.size < AutocleanSizeLower {
			if dryRun {
				fmt.Printf("Would remove: %s -> %s\n", c.name, FormatBytesToStr(c.size))
			} else {
				fmt.Printf("Removing: %s (%s) ...\n", c.name, FormatBytesToStr(c.size))
				err := os.RemoveAll(c.name)
				if err != nil {
					return fmt.Errorf("failed to autoclean: %s", err)
				}
			}
			totalCleanedMbs += float32(c.size / 1000024)
		}
	}

	spaceSavedFmt := "Would save %dMbs of space\n"
	if !dryRun {
		spaceSavedFmt = "Saved %dMbs of space\n"
	}
	fmt.Printf(spaceSavedFmt, int(totalCleanedMbs))

	return nil
}

func doesFileExists(path string) bool {
	_, err := os.Stat(path)
	return !os.IsNotExist(err)
}

func getNextTrashIndex() (int32, error) {
	cachePath := getCachePath()

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
		if doesFileExists(fullPath) {
			// File already exists, then add incremt again...
			cacheNumber++
			fullPath = filepath.Join(cachePath, strconv.FormatInt(int64(cacheNumber), 10))
		} else {
			break
		}
	}

	return cacheNumber, nil
}

func getSrmCacheDir() string {
	home := os.Getenv("HOME")
	return filepath.Join(home, ".cache/srm2")
}

func getCachePath() string {
	home := getSrmCacheDir()

	return filepath.Join(home, "trash")
}

func InitCache() {
	cachePath := getCachePath()

	err := os.MkdirAll(cachePath, 0700)
	if err != nil {
		fmt.Printf("ERROR: %v", err)
	}
}
