// Created by: WestleyR
// Email: westleyr@nym.hush.com
// Url: https://github.com/WestleyR/srm
// Last modified date: 2020-09-22
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
	"io/ioutil"
	"os"
	"path/filepath"
	"strconv"
)

func cleanCacheAUTO() error {
	path := getCachePath()

	// First get the max entities in the cache dir
	// TODO: error check
	files, _ := ioutil.ReadDir(path)
	maxItems := len(files)

	for i := 0; i < maxItems-10; i++ {
		trashPath := filepath.Join(path, strconv.Itoa(i))
		//files, err := ioutil.ReadDir(trashPath)

		fmt.Println("File to remove: ", trashPath)
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

func initCache() {
	cachePath := getCachePath()

	err := os.MkdirAll(cachePath, 0700)
	if err != nil {
		fmt.Printf("ERROR: %v", err)
	}
}
