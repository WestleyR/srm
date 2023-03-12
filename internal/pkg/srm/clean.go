//
//  cache.go
//  srm - https://github.com/WestleyR/srm
//
// Created by WestleyR <westleyr@nym.hush.com> on July 28, 2020
// Source code: https://github.com/WestleyR/srm
//
// Copyright (c) 2020-2023 WestleyR. All rights reserved.
// This software is licensed under a BSD 3-Clause Clear License.
// Consult the LICENSE file that came with this software regarding
// your rights to distribute this software.
//

package srm

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/c2h5oh/datasize"
)

var (
	// Auto clean size limit for all files and directories
	AutocleanSizeUpper = datasize.MB * 15 // 15Mbs
	AutocleanSizeLower = datasize.B * 50  // 50 bytes
)

func (m *Manager) CleanCache(all, dryRun bool) error {
	spacesSaved := datasize.ByteSize(0)

	var dup []*Entry

	for i := 0; i < len(m.Data.Entries); i++ {
		size := m.Data.Entries[i].Size.Bytes()

		remove := all

		if size > AutocleanSizeUpper.Bytes() || size < AutocleanSizeLower.Bytes() {
			remove = true
		}

		if remove {
			spacesSaved += datasize.ByteSize(size)

			if !dryRun {
				err := SafeRMF(filepath.Dir(m.Data.Entries[i].Path))
				if err != nil {
					return fmt.Errorf("failed to clean path: %w", err)
				}
			}
		} else {
			dup = append(dup, m.Data.Entries[i])
		}
	}

	if !dryRun {
		m.Data.Entries = dup
		m.Data.LastTrashIndex = m.getNextTrashIndex(0)
	}

	if dryRun {
		fmt.Printf("Would restored %s from cache.\n", spacesSaved.HumanReadable())
	} else {
		fmt.Printf("Restored %s from cache.\n", spacesSaved.HumanReadable())
	}

	return nil
}

// SafeRMF is a wrapper for os.RemoveAll, but **would** apply some safety
// checks before removing the file/dir.
func SafeRMF(path string) error {
	// TODO: add safety checks

	return os.RemoveAll(path)
}
