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
	"path/filepath"
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

func (m *Manager) ListRecentCache() error {
	l := len(m.Data.Entries)
	for i := l - 10; i < l; i++ {
		if i < 0 {
			continue
		}
		f := m.Data.Entries[i]

		fmt.Printf("%d: %s(%s)%s ", l-i, colorTeal, f.Date.Format("2006-01-02"), colorReset)
		if f.IsDir {
			fmt.Printf("%s%s%s", colorBlue, filepath.Base(f.Path), colorReset)
		} else {
			fmt.Printf("%s", filepath.Base(f.Path))
		}

		fmt.Printf(" %s(%s)%s\n", colorYellow, f.Size.HR(), colorReset)
	}

	return nil
}

func (m *Manager) ListAllCache(bySize, byTime bool) error {
	sortingFunc := func([]*Entry) {}
	if bySize {
		sortingFunc = sortBySize
	} else if byTime {
		sortingFunc = sortByTime
	}

	// TODO: maybe just sort an index, so we dont change the data

	dup := m.Data.Entries
	sortingFunc(dup)

	l := len(dup)
	for i := 0; i < l; i++ {
		f := dup[i]

		fmt.Printf("%d: %s(%s)%s ", l-i, colorTeal, f.Date.Format("2006-01-02"), colorReset)
		if f.IsDir {
			fmt.Printf("%s%s%s", colorBlue, filepath.Base(f.Path), colorReset)
		} else {
			fmt.Printf("%s", filepath.Base(f.Path))
		}

		fmt.Printf(" %s(%s)%s\n", colorYellow, f.Size.HR(), colorReset)
	}

	return nil
}

func sortBySize(cache []*Entry) {
	sorting := true

	for sorting {
		sorting = false
		for i := 0; i < len(cache)-1; i++ {
			if cache[i].Size > cache[i+1].Size {
				tmp := cache[i]
				cache[i] = cache[i+1]
				cache[i+1] = tmp
				sorting = true
			}
		}
	}
}

func sortByTime(cache []*Entry) {
	sorting := true

	for sorting {
		sorting = false
		for i := 0; i < len(cache)-1; i++ {
			if cache[i].Date.Unix() > cache[i+1].Date.Unix() {
				tmp := cache[i]
				cache[i] = cache[i+1]
				cache[i+1] = tmp
				sorting = true
			}
		}
	}
}

//// TODO: cleanup args, like pass the sorting func.
//func ListAllCache(bySize, byTime bool) error {
//	// TODO: fixme
//	path := filepath.Dir(filepath.Dir(paths.GetTrashDirPath()))
//
//	years, err := ioutil.ReadDir(path)
//	if err != nil {
//		return err
//	}
//
//	sortingFunc := sortByName
//	if bySize {
//		sortingFunc = sortBySize
//	} else if byTime {
//		sortingFunc = sortByTime
//	}
//
//	for _, y := range years {
//		months, err := ioutil.ReadDir(filepath.Join(path, y.Name()))
//		if err != nil {
//			return err
//		}
//
//		for _, m := range months {
//			if isPathADir(filepath.Join(path, y.Name(), m.Name())) {
//				formatTrashDirContents(filepath.Join(path, y.Name(), m.Name()), sortingFunc)
//			}
//		}
//	}
//
//	return nil
//}
//
//func formatTrashDirContents(path string, sortingFunc SortingFunc) error {
//	cache, _ := getCacheArray(path)
//
//	sortingFunc(cache)
//
//	for _, f := range cache {
//		trashPath := f.name
//		files, err := ioutil.ReadDir(trashPath)
//		if err != nil {
//			// TODO: if debug is true, then print this
//			//fmt.Fprintf(os.Stderr, "failed to open: %s\n", trashPath)
//		} else {
//			trashName := colorRed + "[no items]" + colorReset
//			colorDir := false
//			if len(files) != 0 {
//				trashName = files[0].Name()
//				if isPathADir(filepath.Join(trashPath, trashName)) {
//					// Trash item is a directory
//					colorDir = true
//				}
//			}
//			fileDate := getFileDate(trashPath)
//			fmt.Printf("%s(%s)%s %s%s/%s", colorTeal, fileDate, colorReset, colorPink, trashPath, colorReset)
//			if colorDir {
//				fmt.Printf("%s%s%s", colorBlue, trashName, colorReset)
//			} else {
//				fmt.Printf("%s", trashName)
//			}
//			fmt.Printf(" %s\n", getTrashFileNumberSizeFormat(filepath.Join(trashPath, trashName)))
//		}
//	}
//
//	return nil
//}
//
//func getFileDate(path string) string {
//	fileDate := ""
//	finfo, err := os.Stat(path)
//	if err != nil {
//		fileDate = "[unknown]"
//		return fileDate
//	}
//
//	fileDate = finfo.ModTime().Format("2006-01-02 15:04:05")
//
//	return fileDate
//}
//
//func getTrashFileNumberSizeFormat(path string) string {
//	files := 0
//	var size int64
//
//	err := filepath.Walk(path, func(path string, info os.FileInfo, err error) error {
//		if err != nil {
//			return err
//		}
//		files++
//		if !info.IsDir() {
//			size += info.Size()
//		}
//
//		return nil
//	})
//	if err != nil {
//		return ""
//	}
//
//	// Take 1 away from files since we are not including the selected directory
//	files--
//
//	fmtStr := colorYellow + "["
//	somethingToPrint := false
//
//	// Not a directory, or empty dir
//	if files != 0 {
//		fmtStr += fmt.Sprintf("%d files", files)
//		somethingToPrint = true
//	}
//	if size != 0 {
//		// Only add the ", " if there was more then one file in the trash item
//		if somethingToPrint {
//			fmtStr += ", "
//		}
//		fmtStr += FormatBytesToStr(size)
//		somethingToPrint = true
//	}
//	fmtStr += "]" + colorReset
//
//	if !somethingToPrint {
//		return ""
//	}
//
//	return fmtStr
//}
//
//func FormatBytesToStr(b int64) string {
//	const unit = 1000
//	div := int64(unit)
//	exp := 0
//	for n := b / unit; n >= unit; n /= unit {
//		div *= unit
//		exp++
//	}
//
//	return fmt.Sprintf("%.1f %cB", float64(b)/float64(div), "kMGTPE"[exp])
//}
//
//func ListRecentCache() error {
//	path := paths.GetTrashDirPath()
//	cache, _ := getCacheArray(path)
//	sortByTime(cache)
//
//	cacheLen := len(cache) - 1
//
//	for i := cacheLen - 9; i <= cacheLen; i++ {
//		if i < 0 {
//			// If less then 10 cached items, then skip the non-existent files
//			continue
//		}
//		f := cache[i]
//		trashPath := f.name
//		files, err := ioutil.ReadDir(trashPath)
//		if err != nil {
//			// TODO: if debug is true, then print this
//			//fmt.Fprintf(os.Stderr, "failed to open: %s\n", trashPath)
//		} else {
//			trashName := colorRed + "[no items]" + colorReset
//			colorDir := false
//			if len(files) != 0 {
//				trashName = files[0].Name()
//				if isPathADir(filepath.Join(trashPath, trashName)) {
//					// Trash item is a directory
//					colorDir = true
//				}
//			}
//			fileDate := getFileDate(trashPath)
//			fmt.Printf("%d: %s(%s)%s %s%s/%s", f.index, colorTeal, fileDate, colorReset, colorPink, trashPath, colorReset)
//			if colorDir {
//				fmt.Printf("%s%s%s", colorBlue, trashName, colorReset)
//			} else {
//				fmt.Printf("%s", trashName)
//			}
//			fmt.Printf(" %s\n", getTrashFileNumberSizeFormat(filepath.Join(trashPath, trashName)))
//		}
//	}
//
//	return nil
//}
