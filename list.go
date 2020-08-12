// Created by: WestleyR
// Email: westleyr@nym.hush.com
// Url: https://github.com/WestleyR/srm
// Last modified date: 2020-08-11
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

const (
	colorRed    string = "\x1b[31m"
	colorGreen  string = "\x1b[32m"
	colorYellow string = "\x1b[33m"
	colorBlue   string = "\x1b[34m"
	colorPink   string = "\x1b[35m"
	colorTeal   string = "\x1b[36m"
	colorWhite  string = "\x1b[37m"
	colorReset  string = "\x1b[0m"
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

func getNumberOfFilesInDir(path string) int {
	files := 0
	_ = filepath.Walk(path, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		files++
		return nil
	})

	// Take 1 away from files since we are not including the selected directory
	files--
	return files
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
			trashName := colorRed + "[no items]" + colorReset
			colorDir := false
			if len(files) != 0 {
				trashName = files[0].Name()
				if isPathADir(filepath.Join(trashPath, trashName)) {
					// Trash item is a directory
					colorDir = true
					trashName = fmt.Sprintf("%s%s [%d files]", trashName, colorYellow, getNumberOfFilesInDir(filepath.Join(trashPath, trashName)))
				}
			}
			fileDate := getFileDate(trashPath)
			fmt.Printf("%d: %s(%s)%s %s%s/%s", i, colorTeal, fileDate, colorReset, colorPink, trashPath, colorReset)
			if colorDir {
				fmt.Printf("%s%s%s\n", colorBlue, trashName, colorReset)
			} else {
				fmt.Printf("%s\n", trashName)
			}
		}
	}

	return nil
}
