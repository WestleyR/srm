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
	"os"
	"os/exec"
	"path/filepath"
	"strconv"
)

// Recover recovers a entry to the current working directory. index should
// 1 = newest, 10 = 10 before the newest.
func (m *Manager) Recover(index int) error {
	entry := m.Data.Entries[len(m.Data.Entries)-index]

	fileName := filepath.Base(entry.Path)

	// Check to see if it already exists
	// Only try 10 times
	tmpName := fileName
	for i := 1; i < 11; i++ {
		if _, err := os.Lstat(tmpName); err == nil {
			tmpName = fileName + "." + strconv.Itoa(i)
		} else {
			fileName = tmpName
			break
		}
	}

	err := os.Rename(entry.Path, fileName)
	if err != nil {
		cmd := exec.Command("mv", entry.Path, fileName)
		err = cmd.Run()
		if err != nil {
			return fmt.Errorf("%s: failed to recover the current directory from %s", fileName, entry.Path)
		}
	}

	m.Data.Remove(len(m.Data.Entries) - index)

	fmt.Printf("File %s has been recovered to the current directory as: %s\n", filepath.Base(entry.Path), fileName)

	return nil
}
