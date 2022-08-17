//
//  recover.go
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
	"os/exec"
	"path/filepath"
	"strconv"

	"github.com/WestleyR/srm/internal/pkg/paths"
)

func RecoverFileFromTrashIndex(trashIndex int) error {
	path := paths.GetTrashDirPath()

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
	    cmd := exec.Command("mv", fullTrashFile, fileName)
	    err = cmd.Run()
            if err != nil {
                return fmt.Errorf("%s: failed to recover the current directory from %s", fileName, fullTrashFile)
            }
	}

	fmt.Printf("File %s has been recovered to the current directory as: %s\n", files[0].Name(), fileName)

	return nil
}
