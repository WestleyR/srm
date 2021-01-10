// Created by: WestleyR
// Email: westleyr@nym.hush.com
// Url: https://github.com/WestleyR/srm
// Last modified date: 2021-01-09
//
// This file is licensed under the terms of
//
// The Clear BSD License
//
// Copyright (c) 2019-2021 WestleyR
// All rights reserved.
//
// This software is licensed under a Clear BSD License.
//

package srm

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"strconv"
)

func RecoverFileFromTrashIndex(trashIndex int) error {
	path := getCachePath()

	trashPath := filepath.Join(path, strconv.Itoa(trashIndex))
	files, err := ioutil.ReadDir(trashPath)
	if err != nil {
		fmt.Println(err)
		return err
	}

	if len(files) == 0 {
		return fmt.Errorf("no files at index: %d", trashIndex)
	}

	fileName := files[0].Name()
	fullTrashFile := filepath.Join(trashPath, fileName)

	// Check to see if it already exists
	// Only try 10 times
	tmpName := fileName
	for i := 1; i < 11; i++ {
		if _, err := os.Stat(tmpName); err == nil {
			tmpName = fileName + "." + strconv.Itoa(i)
		} else {
			fileName = tmpName
			break
		}
	}

	err = os.Rename(fullTrashFile, fileName)
	if err != nil {
		return err
	}

	fmt.Printf("File %s has been recovered to the current directory as: %s\n", files[0].Name(), fileName)

	return nil
}
