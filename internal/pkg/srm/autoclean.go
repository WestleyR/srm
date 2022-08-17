//
//  cache.go
//  srm - https://github.com/WestleyR/srm
//
// Created by WestleyR <westleyr@nym.hush.com> on July 28, 2020
// Source code: https://github.com/WestleyR/srm
//
// Copyright (c) 2020-2022 WestleyR. All rights reserved.
// This software is licensed under a BSD 3-Clause Clear License.
// Consult the LICENSE file that came with this software regarding
// your rights to distribute this software.
//

package srm

import (
	"fmt"
	"os"

	"github.com/WestleyR/srm/internal/pkg/paths"
)

const AutocleanSizeLimit = 15024 * 1024 // 15Mbs
const AutocleanSizeLower = 100          // 100b

// CleanCacheAUTO will remove all cached items above autocleanSizeLimit,
// and below autocleanSizeLower.
func CleanCacheAUTO(dryRun bool) error {
	path := paths.GetTrashDirPath()

	cache, err := getCacheArray(path)
	if err != nil {
		return fmt.Errorf("failed to get cache array: %s", err)
	}

	savedSize := int64(0)

	for _, c := range cache {
		if c.size > AutocleanSizeLimit || c.size < AutocleanSizeLower {
			if dryRun {
				fmt.Printf("Would remove: %s -> %s\n", c.name, FormatBytesToStr(c.size))
			} else {
				fmt.Printf("Removing: %s (%s) ...\n", c.name, FormatBytesToStr(c.size))
				err := os.RemoveAll(c.name)
				if err != nil {
					return fmt.Errorf("failed to autoclean: %s", err)
				}
			}
			savedSize += c.size
		}
	}

	spaceSavedFmt := "Would save %s of space\n"
	if !dryRun {
		spaceSavedFmt = "Saved %s of space\n"
	}
	fmt.Printf(spaceSavedFmt, FormatBytesToStr(savedSize))

	return nil
}
