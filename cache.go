// Created by: WestleyR
// Email: westleyr@nym.hush.com
// Url: https://github.com/WestleyR/srm
// Last modified date: 2020-08-18
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

// cache path should be "~/.cache/srm/trash/{1,2,3...}/item-you-trashed"
func getFileTrashPath(filename string) string {
	cachePath := getCachePath()

	cacheNumber := 0

	items, _ := ioutil.ReadDir(cachePath)
	for _, item := range items {
		if item.IsDir() {
			cacheNumber++
		} else {
			// handle file there
			fmt.Println(item.Name())
		}
	}

	// Give the extra one
	cacheNumber++

	fullPath := filepath.Join(cachePath, strconv.Itoa(cacheNumber))

	err := os.MkdirAll(fullPath, 0700)
	if err != nil {
		fmt.Printf("ERROR: %v", err)
	}

	return filepath.Join(fullPath, filename)
}

func getCachePath() string {
	home := os.Getenv("HOME")

	return filepath.Join(home, ".cache/srm2/trash")
}

func initCache() {
	cachePath := getCachePath()

	err := os.MkdirAll(cachePath, 0700)
	if err != nil {
		fmt.Printf("ERROR: %v", err)
	}
}
