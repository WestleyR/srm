//
//  list.go
//  srm - https://github.com/WestleyR/srm
//
// Created by WestleyR <westleyr@nym.hush.com> on July 29, 2020
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

//nolint:deadcode,varcheck
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

type cacheList struct {
	name  string
	index int
	size  int64
	time  int64
}

func getDirSize(path string) int64 {
	var size int64

	err := filepath.Walk(path, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if !info.IsDir() {
			size += info.Size()
		}
		return nil
	})
	if err != nil {
		size = -1
	}

	return size
}

func getCacheArray(path string) ([]cacheList, error) {
	var cache []cacheList

	items, _ := ioutil.ReadDir(path)
	for _, item := range items {
		if item.IsDir() {
			cachePath := filepath.Join(path, item.Name())
			index, _ := strconv.Atoi(item.Name())
			item := cacheList{
				name:  cachePath,
				index: index,
				size:  getDirSize(cachePath),
				time:  item.ModTime().UnixNano(),
			}
			cache = append(cache, item)
		} else { //nolint:staticcheck // until TODO is fixed
			// TODO: handle file there
			//fmt.Println(item.Name())
		}
	}

	return cache, nil
}

func sortBySize(cache []cacheList) {
	sorting := true

	for sorting {
		sorting = false
		for i := 0; i < len(cache)-1; i++ {
			if cache[i].size > cache[i+1].size {
				tmp := cache[i]
				cache[i] = cache[i+1]
				cache[i+1] = tmp
				sorting = true
			}
		}
	}
}

func sortByName(cache []cacheList) {
	sorting := true

	for sorting {
		sorting = false
		for i := 0; i < len(cache)-1; i++ {
			if cache[i].index > cache[i+1].index {
				tmp := cache[i]
				cache[i] = cache[i+1]
				cache[i+1] = tmp
				sorting = true
			}
		}
	}
}

func sortByTime(cache []cacheList) {
	sorting := true

	for sorting {
		sorting = false
		for i := 0; i < len(cache)-1; i++ {
			if cache[i].time > cache[i+1].time {
				tmp := cache[i]
				cache[i] = cache[i+1]
				cache[i+1] = tmp
				sorting = true
			}
		}
	}
}

// TODO: cleanup args, like pass the sorting func.
func ListAllCache(bySize, byTime bool) error {
	sortingFunc := sortByName
	if bySize {
		sortingFunc = sortBySize
	} else if byTime {
		sortingFunc = sortByTime
	}

	path := getCachePath()

	cache, _ := getCacheArray(path)

	sortingFunc(cache)

	for _, f := range cache {
		trashPath := f.name
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
			fmt.Printf("%d: %s(%s)%s %s%s/%s", f.index, colorTeal, fileDate, colorReset, colorPink, trashPath, colorReset)
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

// TODO: need to cleanup the following functions

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
		fmtStr += FormatBytesToStr(size)
		somethingToPrint = true
	}
	fmtStr += "]" + colorReset

	if !somethingToPrint {
		return ""
	}

	return fmtStr
}

// TODO: need to have a util package
func FormatBytesToStr(b int64) string {
	const unit = 1000
	//	if b < unit {
	//		return fmt.Sprintf("%d B", b)
	//	}
	div := int64(unit)
	exp := 0
	for n := b / unit; n >= unit; n /= unit {
		div *= unit
		exp++
	}
	return fmt.Sprintf("%.1f %cB", float64(b)/float64(div), "kMGTPE"[exp])
}

func ListRecentCache() error {
	path := getCachePath()
	cache, _ := getCacheArray(path)
	sortByTime(cache)

	cacheLen := len(cache) - 1

	for i := cacheLen - 9; i <= cacheLen; i++ {
		if i < 0 {
			// If less then 10 cached items, then skip the non-existent files
			continue
		}
		f := cache[i]
		trashPath := f.name
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
			fmt.Printf("%d: %s(%s)%s %s%s/%s", f.index, colorTeal, fileDate, colorReset, colorPink, trashPath, colorReset)
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
