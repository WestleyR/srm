// Created by: WestleyR
// Email: westleyr@nym.hush.com
// Url: https://github.com/WestleyR/srm
// Last modified date: 2020-07-29
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

func getFileDate(path string) string {
	fileDate := ""
	finfo, err := os.Stat(path)
	if err != nil {
		fileDate = "[unknown]"
		return fileDate
	}

	fileDate = finfo.ModTime().Format("2006-01-02 15:04:05")

	return fileDate
}

func listRecentCache() error {
	path := getCachePath()

	// First get the max entities in the cache dir
	// TODO: error check
	files, _ := ioutil.ReadDir(path)
	maxItems := len(files)

	for i := maxItems - 10; i <= maxItems; i++ {
		trashPath := filepath.Join(path, strconv.Itoa(i))
		files, err := ioutil.ReadDir(trashPath)
		if err != nil {
			fmt.Println(err)
		} else {
			trashName := "[no items]"
			if len(files) != 0 {
				trashName = files[0].Name()
			}
			fileDate := getFileDate(trashPath)
			fmt.Printf("%d: (%s) %s/%s\n", i, fileDate, trashPath, trashName)
		}
	}

	return nil
}
