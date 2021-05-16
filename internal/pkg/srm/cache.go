//
//  cache.go
//  srm - https://github.com/WestleyR/srm
//
// Created by WestleyR <westleyr@nym.hush.com> on July 28, 2020
// Source code: https://github.com/WestleyR/srm
//
// Copyright (c) 2020-2021 WestleyR. All rights reserved.
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

const autocleanSizeLimit = 15024 * 1024 // 5Mbs

// CleanCacheAUTO will remove all cached items above autocleanSizeLimit.
func CleanCacheAUTO() error {
	path := getCachePath()

	cache, _ := getCacheArray(path)

	sortBySize(cache)

	lastCleanedSize := cache[len(cache)-1].size

	for i := len(cache)-1; i > 0; i-- {
		if lastCleanedSize < autocleanSizeLimit {
			break
		}
		fmt.Printf("Would remove: %s -> %s\n", cache[i].name, formatBytesToStr(cache[i].size))
		lastCleanedSize = cache[i].size
	}

	return nil
}

func doesFileExists(path string) bool {
	_, err := os.Stat(path)
	return !os.IsNotExist(err)
}

// cache path should be "~/.cache/srm/trash/{1,2,3...}/item-you-trashed"
func getFileTrashPath(filename string) string {
	cachePath := getCachePath()

	cacheNumber := 0

	items, _ := ioutil.ReadDir(cachePath)
	for _, item := range items {
		if item.IsDir() {
			cacheNumber++
		} else {
			// TODO: handle file there
			fmt.Println(item.Name())
		}
	}

	// Give the extra one
	cacheNumber++

	fullPath := filepath.Join(cachePath, strconv.Itoa(cacheNumber))

	// Make sure it does not already exist, sometimes
	// when one of the cache dirs is removed, it can
	// screw up the last number. Only try 100 times
	// TODO: add debug here...
	for i := 0; i < 100; i++ {
		if doesFileExists(fullPath) {
			// File already exists, then add incremt again...
			cacheNumber++
			fullPath = filepath.Join(cachePath, strconv.Itoa(cacheNumber))
		} else {
			break
		}
	}

	// Write the last cache path number
	lastCachePath := filepath.Join(getSrmCacheDir(), "last.cachepath")
	fp, err := os.OpenFile(lastCachePath, os.O_CREATE|os.O_WRONLY|os.O_TRUNC, 0644)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Failed to open cache path file: %s\n", err)
		return ""
	}
	fmt.Fprintf(fp, "%d", cacheNumber)
	fp.Close()

	// Create the dir
	err = os.MkdirAll(fullPath, 0700)
	if err != nil {
		fmt.Printf("ERROR: %v", err)
	}

	return filepath.Join(fullPath, filename)
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
