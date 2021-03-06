//
//  list.go
//  srm - https://github.com/WestleyR/srm
//
// Created by WestleyR on July 29, 2020
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

func getTrashFileNumberSizeFormat(path string) string {
	files := 0
	var size int64

	err := filepath.Walk(path, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		files++
		if !info.IsDir() {
			size += info.Size()
		}

		return nil
	})
	if err != nil {
		return ""
	}

	// Take 1 away from files since we are not including the selected directory
	files--

	fmtStr := colorYellow + "["
	somethingToPrint := false

	// Not a directory, or empty dir
	if files != 0 {
		fmtStr += fmt.Sprintf("%d files", files)
		somethingToPrint = true
	}
	if size != 0 {
		// Only add the ", " if there was more then one file in the trash item
		if somethingToPrint {
			fmtStr += ", "
		}
		fmtStr += formatBytesToStr(size)
		somethingToPrint = true
	}
	fmtStr += "]" + colorReset

	if !somethingToPrint {
		return ""
	}

	return fmtStr
}

func formatBytesToStr(b int64) string {
	const unit = 1000
	if b < unit {
		return fmt.Sprintf("%d B", b)
	}
	div, exp := int64(unit), 0
	for n := b / unit; n >= unit; n /= unit {
		div *= unit
		exp++
	}
	return fmt.Sprintf("%.1f %cB", float64(b)/float64(div), "kMGTPE"[exp])
}

func ListRecentCache() error {
	path := getCachePath()

	// Get the last cache number
	maxItems := 0
	items, _ := ioutil.ReadDir(path)
	for _, item := range items {
		if item.IsDir() {
			maxItems++
		} else {
			// TODO: handle file there
			fmt.Println(item.Name())
		}
	}

	fullPath := filepath.Join(path, strconv.Itoa(maxItems))

	// Make sure no more exist, sometimes when one of the cache dirs is
	// removed, it can screw up the last number. Only try 100 times
	// TODO: add debug here...
	for i := 0; i < 100; i++ {
		if doesFileExists(fullPath) {
			// File already exists, then add incremt again...
			maxItems++
			fullPath = filepath.Join(path, strconv.Itoa(maxItems))
		} else {
			break
		}
	}

	for i := maxItems - 10; i <= maxItems; i++ {
		trashPath := filepath.Join(path, strconv.Itoa(i))
		files, err := ioutil.ReadDir(trashPath)
		if err != nil {
			// TODO: if debug is true, then print this
			//fmt.Fprintf(os.Stderr, "failed to open: %s\n", trashPath)
		} else {
			trashName := colorRed + "[no items]" + colorReset
			colorDir := false
			if len(files) != 0 {
				trashName = files[0].Name()
				if isPathADir(filepath.Join(trashPath, trashName)) {
					// Trash item is a directory
					colorDir = true
				}
			}
			fileDate := getFileDate(trashPath)
			fmt.Printf("%d: %s(%s)%s %s%s/%s", i, colorTeal, fileDate, colorReset, colorPink, trashPath, colorReset)
			if colorDir {
				fmt.Printf("%s%s%s", colorBlue, trashName, colorReset)
			} else {
				fmt.Printf("%s", trashName)
			}
			fmt.Printf(" %s\n", getTrashFileNumberSizeFormat(filepath.Join(trashPath, trashName)))
		}
	}

	return nil
}
